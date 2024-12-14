package config

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		AWSCognito
	}
	App struct {
		Name    string `yaml:"name" env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	// HTTP
	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL"`
	}

	AWSCognito struct {
		AppID  string `env-required:"true" env:"AWS_COGNITO_APP_ID"`
		Region string `env-required:"true" env:"AWS_COGNITO_REGION"`
	}
)
