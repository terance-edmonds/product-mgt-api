package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

const (
	DefaultPort            = 8080
	DefaultHostname        = "localhost"
	DefaultInitialDataPath = "configs/initial_data.json"
	DefaultBasePath        = "/api/v1"
)

var (
	EnvName         = "ENV"
	Hostname        = "HOSTNAME"
	Port            = "PORT"
	basePath        = "BASE_PATH"
	initialDataPath = "INIT_DATA_PATH"
)

var config Config

func GetConfig() *Config {
	return &config
}

func LoadConfig() (*Config, error) {
	config = Config{
		Hostname:        getEnvString(Hostname, DefaultHostname),
		Port:            getEnvInt(Port, DefaultPort),
		Env:             os.Getenv(EnvName),
		InitialDataPath: getEnvString(initialDataPath, DefaultInitialDataPath),
		BasePath:        getEnvString(basePath, DefaultBasePath),
	}
	return &config, nil
}

func LoadInitialData() (data InitialData) {
	log.Printf("Attempting to read initial data from: %s", config.InitialDataPath)
	if config.InitialDataPath == "" {
		return
	}
	contents, err := os.ReadFile(config.InitialDataPath)
	if err != nil {
		log.Fatalf("failed to read initial data at [%s]: %s", config.InitialDataPath, err)
	}
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Fatalf("failed to unmarshal initial data at [%s]: %s", config.InitialDataPath, err)
	}
	return
}

func getEnvInt(key string, defaultVal int) int {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func getEnvString(key string, defaultVal string) string {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	return s
}
