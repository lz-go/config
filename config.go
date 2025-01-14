package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Load() {
	path := lookupConfigPath()

	if err := godotenv.Load(fmt.Sprintf("%s/.env", path)); err != nil {
		log.Printf("godotenv.Load returns error: %s\n", err.Error())
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper.ReadInConfig returns error: %s\n", err.Error())
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func lookupConfigPath() string {
	pathHasConfigFile := func(path string) bool {
		filePath := filepath.Join(path, "config.yaml")

		_, err := os.Stat(filePath)
		if err == nil {
			return true
		} else if os.IsNotExist(err) {
			return false
		} else {
			log.Fatalf("os.Stat returns error: %s (%s)", err.Error(), filePath)
			return false
		}
	}

	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd returns error: %s (%s)", err.Error(), currentPath)
	}

	for {
		if pathHasConfigFile(currentPath) {
			return currentPath
		}

		parentPath := filepath.Dir(currentPath)
		if parentPath == currentPath {
			return currentPath
		}

		currentPath = parentPath
	}
}
