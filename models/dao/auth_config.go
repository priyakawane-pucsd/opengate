package dao

type AuthConfig struct {
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
