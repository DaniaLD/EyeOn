package configs

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var configs Configs

func LoadConfigs() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../configs")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading common-structs file, %s", err)
	}

	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func GetConfigs() Configs {
	return configs
}
