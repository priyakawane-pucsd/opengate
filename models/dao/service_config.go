package dao

type ServiceConfig struct {
	Endpoint      string   `bson:"endpoint"`
	Regex         string   `bson:"regex"`
	Authorization bool     `bson:"authorization"`
	Roles         []string `bson:"roles"`
}
