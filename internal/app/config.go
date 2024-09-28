package app

import (
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	SPAPath          string   `mapstructure:"SPA_PATH"`
	SPACacheDisabled bool     `mapstructure:"SPA_CACHE_DISABLED"`
	ClientOrigins    []string `mapstructure:"CLIENT_ORIGINS"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetTypeByDefaultValue(true)
	viper.SetConfigType("env")
	if envFile := os.Getenv("ENV_FILE"); envFile != "" {
		viper.SetConfigFile(envFile)
	} else {
		viper.SetConfigFile(".env")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}

	// Silly bug in viper, it will only read ENV variables if they are defined in a .env file.
	// So we manually tell viper to read all variables defined in the Config type.
	envKeysMap := &map[string]interface{}{}
	if err := mapstructure.Decode(config, &envKeysMap); err != nil {
		return nil, err
	}
	for k := range *envKeysMap {
		if bindErr := viper.BindEnv(k); bindErr != nil {
			return nil, err
		}
	}

	err = viper.Unmarshal(config)

	return config, nil
}
