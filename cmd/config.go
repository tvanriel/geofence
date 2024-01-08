package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigServiceRule struct {
	If     string `yaml:"if"`
	Action string `yaml:"action"`
}

type ConfigService struct {
	Name           string              `mapstructure:"name"`
	Listen         string              `mapstructure:"listen"`
	Upstream       string              `mapstructure:"upstream"`
	StandardAction string              `mapstructure:"standard_action"`
	Rules          []ConfigServiceRule `mapstructure:"rules"`
}

type ConfigReport struct {
	BotToken  string `mapstructure:"token"`
	ChannelID string `mapstructure:"channel"`
}

type Configuration struct {
	Db       string          `mapstructure:"db"`
	Report   ConfigReport    `mapstructure:"report"`
	Services []ConfigService `mapstructure:"services"`
}

func NewConfiguration() (*Configuration, error) {
	config := &Configuration{}
	x := viper.GetString("token")
	fmt.Println(x)
	err := viper.Unmarshal(config)
	return config, err
}
