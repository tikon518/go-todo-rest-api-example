package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "103.195.49.199",
			Port:     5589,
			Username: "dbadmin",
			Password: "aXa#iMH*US%*Jzfkr",
			Name:     "shopify",
			Charset:  "utf8",
		},
	}
}
