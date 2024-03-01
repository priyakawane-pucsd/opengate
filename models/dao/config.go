package dao

type Config struct {
	Id            string         `bson:"_id,omitempty"`
	Type          string         `bson:"type,omitempty"`
	CreatedOn     int64          `bson:"createdOn"`
	UpdatedOn     int64          `bson:"updatedOn"`
	AuthConfig    *AuthConfig    `bson:"authConfig,omitempty"`
	ServiceConfig *ServiceConfig `bson:"serviceConfig,omitempty"`
}
