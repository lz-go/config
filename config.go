package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Load() error {
	path, err := lookupConfigPath()
	if err != nil {
		return err
	}

	_ = godotenv.Load(fmt.Sprintf("%s/.env", path))

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}

func lookupConfigPath() (string, error) {
	pathHasConfigFile := func(path string) (bool, error) {
		filePath := filepath.Join(path, "config.yaml")

		_, err := os.Stat(filePath)
		if err == nil {
			return true, nil
		} else if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}

	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		existed, err := pathHasConfigFile(currentPath)
		if err != nil {
			return "", err
		}
		if existed {
			return currentPath, nil
		}

		parentPath := filepath.Dir(currentPath)
		if parentPath == "/" || parentPath == "." {
			return currentPath, nil
		}

		currentPath = parentPath
	}
}
