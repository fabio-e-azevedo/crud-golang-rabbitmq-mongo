package config

import (
	"os"
)

type Config struct {
	ConfigMongo
	ConfigRabbit
}

type ConfigMongo struct {
	MongoURI        string
	MongoDatabase   string
	MongoCollection string
}

func NewConfigMongo() *ConfigMongo {
	return &ConfigMongo{
		MongoURI:        getEnv("MONGODB_URI", "localhost"),
		MongoDatabase:   getEnv("MONGODB_DATABASE", "localhost"),
		MongoCollection: getEnv("RESOURCE_TYPE", "changeit"),
	}
}

type ConfigRabbit struct {
	RabbitURI           string
	RabbitQueueConsumer string
}

func NewConfigRabbit() *ConfigRabbit {
	return &ConfigRabbit{
		RabbitURI:           getEnv("RABBITMQ_URI", "localhost"),
		RabbitQueueConsumer: getEnv("RESOURCE_TYPE", "changeit"),
	}
}

func NewConfig() *Config {
	return &Config{
		ConfigMongo: ConfigMongo{
			MongoURI:        getEnv("MONGODB_URI", "mongodb://root:root@localhost"),
			MongoDatabase:   getEnv("MONGODB_DATABASE", "db"),
			MongoCollection: getEnv("RESOURCE_TYPE", "changeit"),
		},
		ConfigRabbit: ConfigRabbit{
			RabbitURI:           getEnv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/"),
			RabbitQueueConsumer: getEnv("RESOURCE_TYPE", "changeit"),
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
