package configs

import (
	"encoding/json"
	"os"
)

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
	Timezone string `json:"timezone"`
}

type CloudinaryConfig struct {
	CloudName string `json:"cloud_name"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

type Config struct {
	Database     DatabaseConfig   `json:"database"`
	JWTSecretKey string           `json:"jwt_secret_key"`
	Cloudinary   CloudinaryConfig `json:"cloudinary"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("configs/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		panic(err)
	}
}