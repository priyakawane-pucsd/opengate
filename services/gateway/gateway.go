package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"opengate/models/dao"
	"opengate/models/dto"
	"opengate/utils"
	"regexp"
	"strings"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repo       Repository
	srvConfigs []*dao.ServiceConfig
	authConfig *dao.AuthConfig
}

type Repository interface {
	GetAllConfigs(ctx context.Context) ([]*dao.ServiceConfig, error)
	GetAuthConfig(ctx context.Context) (*dao.AuthConfig, error)
}

func NewService(ctx context.Context, repo Repository) *Service {
	configs, err := repo.GetAllConfigs(ctx)
	if err != nil {
		logger.Panic(ctx, "getting service configs")
		return nil
	}

	auth, err := repo.GetAuthConfig(ctx)
	if err != nil {
		logger.Panic(ctx, "getting service configs")
		return nil
	}
	return &Service{repo: repo, srvConfigs: configs, authConfig: auth}
}

// completed this function
func (s *Service) getConfig(ctx context.Context, urlPath string) *dao.ServiceConfig {
	for _, c := range s.srvConfigs {
		r, err := regexp.Compile(*&c.ServiceConfig.Regex)
		if err != nil {
			logger.Error(ctx, "invalid regular expression in config: %v", c)
			continue
		}

		match := r.FindString(urlPath)
		if match != "" {
			return c
		}
	}
	return nil
}

func (s *Service) authorize(ctx *gin.Context) (*dto.AuthorizationResponse, error) {
	// URL of the service you want to send a request to
	apiURL := s.authConfig.AuthConfig.Endpoint

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	request, err := http.NewRequest(s.authConfig.AuthConfig.RequestMethod, apiURL, nil)
	if err != nil {
		logger.Error(ctx, "Error creating request: %s\n", err.Error())
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error creating request")
	}

	// Add headers to the request
	for _, h := range s.authConfig.AuthConfig.Headers {
		hd := ctx.GetHeader(h)
		request.Header.Add(h, hd)
	}

	// Send the request
	response, err := client.Do(request)
	if err != nil {
		logger.Error(ctx, "Error sending request: %s\n", err.Error())
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error sending request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "Unathorized")
	}
	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(ctx, "Error reading response body: %s\n", err.Error())
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error reading response body")
	}

	var res dto.AuthorizationResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		logger.Error(ctx, "invalid response from auth service %v", err.Error())
		return nil, utils.NewCustomError(http.StatusInternalServerError, "invalid response from auth service")
	}
	// Print the response body
	logger.Info(ctx, "Response from %s:\n%s\n", apiURL, body)
	return &res, nil
}

func getValue(res *dto.AuthorizationResponse, address string) (any, error) {
	// Split the address into keys
	keys := strings.Split(address, ".")

	// Traverse the map using the keys
	currentMap := *res
	end := len(keys) - 1
	for i, key := range keys {
		value, ok := currentMap[key]
		if !ok {
			return nil, fmt.Errorf("Key '%s' not found in AuthorizationResponse", key)
		}
		if i == end {
			return value, nil
		}
		// Check if the value is a nested map
		if nestedMap, ok := value.(map[string]interface{}); ok {
			currentMap = nestedMap
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "invalid response from authservice")
		}
	}

	return nil, utils.NewCustomError(http.StatusInternalServerError, "invalid response from authservice")
}

func (s *Service) modifyRequest(ctx *gin.Context, res *dto.AuthorizationResponse) error {
	forwardHeadersConfig := s.authConfig.AuthConfig.ForwardHeaders
	for _, h := range forwardHeadersConfig {
		value, err := getValue(res, h.Address)
		if err != nil {
			return err
		}
		str := fmt.Sprintf("%v", value)
		ctx.Request.Header.Set(h.Key, str)
	}
	return nil
}

func (s *Service) authorizeAndModifyRequest(ctx *gin.Context) error {
	res, err := s.authorize(ctx)
	if err != nil {
		return err
	}
	logger.Info(ctx, "Response of authorization>>>> %v", res)

	//TODO MODIFY REQUEST
	err = s.modifyRequest(ctx, res)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) HandleRequest(ctx *gin.Context) error {
	cfg := s.getConfig(ctx, ctx.Param("path"))
	if cfg == nil {
		return utils.NewCustomError(http.StatusNotFound, "unknown service")
	}

	//If authorization true then verify authorization
	if cfg.ServiceConfig.Authorization {
		err := s.authorizeAndModifyRequest(ctx)
		if err != nil {
			return err
		}
	}

	remote, err := url.Parse(cfg.ServiceConfig.Endpoint)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "invalid endpoint config in db")
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.Method = ctx.Request.Method
		req.URL.Path = ctx.Param("path")
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
	return nil
}
