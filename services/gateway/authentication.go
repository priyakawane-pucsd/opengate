package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"opengate/models/dto"
	"opengate/utils"
	"strings"
	"time"

	"github.com/bappaapp/goutils/logger"
	"github.com/gin-gonic/gin"
)

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
	body, err := io.ReadAll(response.Body)
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
			return nil, fmt.Errorf("key '%s' not found in AuthorizationResponse", key)
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

func (s *Service) modifyRequest(ctx *gin.Context, forwordHeadersMap map[string]string) error {
	for k, v := range forwordHeadersMap {
		ctx.Request.Header.Set(k, v)
	}
	return nil
}

func (s *Service) getForwardAuthHeaders(ctx *gin.Context, res *dto.AuthorizationResponse) (map[string]string, error) {
	var forwardHeaders map[string]string = make(map[string]string)
	forwardHeadersConfig := s.authConfig.AuthConfig.ForwardHeaders
	for _, h := range forwardHeadersConfig {
		value, err := getValue(res, h.Address)
		if err != nil {
			return nil, err
		}
		str := fmt.Sprintf("%v", value)
		forwardHeaders[h.Key] = str
	}
	return forwardHeaders, nil
}

func (s *Service) getAuthHeaderCacheKey(ctx *gin.Context) string {
	var authHeaderKey string
	for _, h := range s.authConfig.AuthConfig.Headers {
		value := ctx.GetHeader(h)
		authHeaderKey += value + "_"
	}
	return authHeaderKey
}

func (s *Service) authorizeAndModifyRequest(ctx *gin.Context) error {

	var forwordHeadersMap map[string]string
	var authHeaderCacheKey string
	if s.authConfig.AuthConfig.IsCache {
		authHeaderCacheKey = s.getAuthHeaderCacheKey(ctx)
		err := s.cache.GetV(ctx, authHeaderCacheKey, &forwordHeadersMap)
		if err == nil {
			return s.modifyRequest(ctx, forwordHeadersMap)
		}
	}

	res, err := s.authorize(ctx)
	if err != nil {
		return err
	}
	logger.Info(ctx, "Response of authorization>>>> %v", res)
	forwordHeadersMap, err = s.getForwardAuthHeaders(ctx, res)
	if err != nil {
		return err
	}

	if s.authConfig.AuthConfig.IsCache {
		s.cache.SetWithTimeout(ctx, authHeaderCacheKey, forwordHeadersMap, time.Duration(s.authConfig.AuthConfig.CacheExpiryMins))
	}
	//TODO MODIFY REQUEST
	return s.modifyRequest(ctx, forwordHeadersMap)

}
