package dao

type AuthConfig struct {
	Id         string           `bson:"_id,omitempty"`
	Type       string           `bson:"type,omitempty"`
	AuthConfig CreateAuthConfig `bson:"authConfig,omitempty"`
	CreatedOn  int64            `bson:"createdOn"`
	UpdatedOn  int64            `bson:"updatedOn"`
}

type CreateAuthConfig struct {
	Endpoint        string          `bson:"endpoint"`
	Headers         []string        `bson:"headers"`
	RequestMethod   string          `bson:"requestMethod"`
	ForwardHeaders  []ForwardHeader `bson:"forwardHeaders"`
	RolesKey        string          `bson:"rolesKey"`
	IsCache         bool            `bson:"isCache"`
	CacheExpiryMins int             `bson:"cacheExpiryMins"`
}

// ForwardHeader represents the structure of the ForwardHeaders array
type ForwardHeader struct {
	Key     string `bson:"key"`
	Address string `bson:"address"`
}
