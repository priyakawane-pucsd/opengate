package dao

type AuthConfigService struct {
	Id         string                   `bson:"_id,omitempty"`
	Type       string                   `bson:"type,omitempty"`
	AuthConfig CreateAuthConfigRequest1 `bson:"authConfig,omitempty"`
	CreatedOn  int64                    `bson:"createdOn"`
	UpdatedOn  int64                    `bson:"updatedOn"`
}

type CreateAuthConfigRequest1 struct {
	Endpoint string   `bson:"endpoint"`
	Headers  []string `bson:"headers"`
}
