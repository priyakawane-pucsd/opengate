package dao

// Newly Added
type ServiceConfig struct {
	Id            string                     `bson:"_id,omitempty"`
	Type          string                     `bson:"type,omitempty"`
	ServiceConfig CreateServiceConfigRequest `bson:"serviceConfig,omitempty"`
	CreatedOn     int64                      `bson:"createdOn"`
	UpdatedOn     int64                      `bson:"updatedOn"`
}

// newly added
type CreateServiceConfigRequest struct {
	Endpoint      string   `bson:"endpoint"`
	Regex         string   `bson:"regex"`
	Authorization bool     `bson:"authorization"`
	Roles         []string `bson:"roles"`
}
