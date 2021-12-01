package mongodb

// Config 是 mongo 客户端配置。
type Config struct {
	URI          string              `json:"uri"`
	Database     string              `json:"database"`
	Auth         *AuthConfig         `json:"auth"`
	TimeoutMs    int64               `json:"timeout_ms"`
	WriteConcern *WriteConcernConfig `json:"write_concern"`
	ReadConcern  string              `json:"read_concern"`
}

// WriteConcernConfig 描述了 mongo 中 write concern 的配置，具体见：https://docs.mongodb.com/manual/reference/write-concern/。
type WriteConcernConfig struct {
	Majority  bool  `json:"majority"`
	TimeoutMs int64 `json:"timeout_ms"`
}

// AuthConfig 描述了 mongo 中 认证的相关信息
type AuthConfig struct {
	AuthMechanism           string            `json:"auth_mechanism"`
	AuthMechanismProperties map[string]string `json:"auth_mechanism_properties"`
	AuthSource              string            `json:"auth_source"`
	Username                string            `json:"username"`
	Password                *string           `json:"password"`
}
