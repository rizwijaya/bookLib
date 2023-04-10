package config

type LoadConfig struct {
	App struct {
		Mode       string `env:"APP_MODE"`
		Name       string `env:"APP_NAME"`
		Port       string `env:"PORT"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET_KEY"`
	}

	Database struct {
		Host     string `env:"PGHOST"`
		Name     string `env:"PGDATABASE"`
		Username string `env:"PGUSER"`
		Password string `env:"PGPASSWORD"`
		Port     string `env:"PGPORT"`
	}
}
