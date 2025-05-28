package configs

type Configs struct {
	Env     string         `mapstructure:"env"`
	BitPin  BitPinConfigs  `mapstructure:"bitpin"`
	Nobitex NobitexConfigs `mapstructure:"nobitex"`
}

type BitPinConfigs struct {
	ApiKey    string `mapstructure:"apiKey"`
	SecretKey string `mapstructure:"secretKey"`
}

type NobitexConfigs struct {
	ApiKey string `mapstructure:"apiKey"`
}
