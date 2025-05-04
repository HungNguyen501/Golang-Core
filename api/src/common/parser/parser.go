package parser

import (
	"fmt"

	"github.com/spf13/viper"

	"golang-core/api/src/common/parser/model"
)

func ParseAppConfig(env string) (cfg model.ApplicationConfig, err error) {
	configPath := "./config"
	if env == "test" {
		configPath = "../../../config"
	}
	viper.SetConfigName("app-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("error reading config file: %v", err)
	} else {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
	viper.AutomaticEnv()

	// server
	viper.BindEnv("server.port", "SERVER_PORT")

	// Database
	viper.BindEnv("database.url", "DATABASE_URL")
	viper.BindEnv("database.replica_url", "DATABASE_REPLICA_URL")
	viper.BindEnv("database.name", "DATABASE_NAME")
	viper.BindEnv("database.port", "DATABASE_PORT")
	viper.BindEnv("database.username", "DATABASE_USERNAME")
	viper.BindEnv("database.password", "DATABASE_PASSWORD")

	// Redis
	viper.BindEnv("redis.hosts", "REDIS_HOSTS")

	// Router
	viper.BindEnv("router.allowed_origins", "ROUTER_ALLOWED_ORIGINS")
	viper.BindEnv("router.allowed_headers", "ROUTER_ALLOWED_HEADERS")

	// AWS arn
	viper.BindEnv("aws_config.k8s_role_arn", "K8S_ROLE_ARN")

	// AWS S3
	viper.BindEnv("aws_config.s3.data_bucket_name", "AWS_CONFIG_S3_DATA_BUCKET_NAME")
	viper.BindEnv("aws_config.use_localstack", "AWS_CONFIG_USE_LOCALSTACK")

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("Unable to decode into struct, %v", err)
	}
	return cfg, err
}
