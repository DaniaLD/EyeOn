package configs

type Configs struct {
	Env    string        `mapstructure:"env"`
	BitPin BitPinConfigs `mapstructure:"bitpin"`
}

type BitPinConfigs struct {
	ApiKey    string `mapstructure:"apiKey"`
	SecretKey string `mapstructure:"secretKey"`
}
