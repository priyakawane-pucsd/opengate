package dao

type ServiceConfig struct {
	Name     string        `bson:"name"`
	Endpoint string        `bson:"endpoint"`
	Regex    string        `bson:"regex"`
	Apis     []ServiceApis `bson:"apis"`
}

type ServiceApis struct {
	Regex         string   `bson:"regex"`
	Authorization bool     `bson:"authorization"`
	Roles         []string `bson:"roles"`
}
