package dao

type ServiceConfig struct {
	Id        string `bson:"_id,omitempty"`
	Endpoint  string `bson:"endpoint"`
	Regex     string `bson:"regex"`
	CreatedOn int64  `bson:"createdOn"`
	UpdatedOn int64  `bson:"updatedOn"`
}
