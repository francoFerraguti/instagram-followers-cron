package main

type Config struct {
	ENV         string
	PORT        string
	DB_TYPE     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		config := newConfig()
		instance = &config
	}
	return instance
}

func newConfig() Config {
	return Config{
		ENV:         "develop",
		PORT:        "8000",
		DB_TYPE:     "mysql",
		DB_USERNAME: "b30424f2194078",
		DB_PASSWORD: "0cd09c77",
		DB_HOST:     "us-cdbr-iron-east-05.cleardb.net",
		DB_PORT:     "3306",
		DB_NAME:     "heroku_496add156cc831f",
	}
}
