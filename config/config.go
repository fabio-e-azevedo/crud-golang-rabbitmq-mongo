package config

import (
	"os"
)

type Config struct {
	ConfigMongo
	ConfigRabbit
}

type ConfigMongo struct {
	MongoURI      string
	MongoDatabase string
}

func NewConfigMongo() *ConfigMongo {
	return &ConfigMongo{
		MongoURI:      getEnv("MONGODB_URI", "localhost"),
		MongoDatabase: getEnv("MONGODB_DATABASE", "localhost"),
	}
}

type ConfigRabbit struct {
	RabbitURI string
}

func NewConfigRabbit() *ConfigRabbit {
	return &ConfigRabbit{
		RabbitURI: getEnv("RABBITMQ_URI", "localhost"),
	}
}

func NewConfig() *Config {
	return &Config{
		ConfigMongo: ConfigMongo{
			MongoURI:      getEnv("MONGODB_URI", "mongodb://root:root@localhost"),
			MongoDatabase: getEnv("MONGODB_DATABASE", "db"),
		},
		ConfigRabbit: ConfigRabbit{
			RabbitURI: getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
