package model

type ApplicationConfig struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	RedisConfig    RedisConfig    `mapstructure:"redis"`
	DatabaseConfig DatabaseConfig `mapstructure:"database"`
	RouterConfig   RouterConfig   `mapstructure:"router"`
	AwsConfig      AwsConfig      `mapstructure:"aws_config"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	URI        string `mapstructure:"uri"`
	Protocol   string `mapstructure:"protocol"`
	URL        string `mapstructure:"url"`
	ReplicaURL string `mapstructure:"replica_url"`
	Name       string `mapstructure:"name"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Port       int    `mapstructure:"port"`

	MaxDBConns      int `mapstructure:"max_db_conns"`
	MaxConnLifetime int `mapstructure:"max_conn_lifetime"`
	MaxConnIdleTime int `mapstructure:"max_conn_idle_time"`
}

type RedisConfig struct {
	Hosts           string `mapstructure:"hosts"`
	PoolSize        int    `mapstructure:"pool_size"`
	MinIdleConns    int    `mapstructure:"min_idle_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	WriteTimeout    int    `mapstructure:"write_timeout"`
	ReadTimeout     int    `mapstructure:"read_timeout"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type RouterConfig struct {
	AllowedOrigins string `mapstructure:"allowed_origins"`
	AllowedHeaders string `mapstructure:"allowed_headers"`
}

type AwsConfig struct {
	K8sRoleArn    string   `mapstructure:"k8s_role_arn"`
	UseLocalstack bool     `mapstructure:"use_localstack"`
	S3Config      S3Config `mapstructure:"s3"`
}

type S3Config struct {
	DataBucketName string `mapstructure:"data_bucket_name"`
}
