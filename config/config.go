package config

type (
	Config struct {
		Databases             []Database              `yaml:"databases"`
		Translator            Translator              `yaml:"translator"`
		Logging               Logging                 `yaml:"logging"`
		Gateway               Microservice            `yaml:"gateway"`
		Microservices         map[string]Microservice `yaml:"microservices"`
		Debug                 bool                    `yaml:"debug"`
		Domain                string                  `yaml:"domain"`
		PWD                   string                  `yaml:"pwd"`
		AllowOrigins          string                  `yaml:"allow_origins"`
		AllowHeaders          string                  `yaml:"allow_headers"`
		MaxAge                int                     `yaml:"max_age"`
		Timeout               int64                   `yaml:"timeout"`
		MaxConcurrentRequests int                     `yaml:"max_concurrent_requests"`
		SecretKey             string                  `yaml:"secret_key"`
	}

	Database struct {
		Name     string `yaml:"name"`
		Type     string `yaml:"type"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}

	Translator struct {
		Path string `yaml:"path"`
	}

	Logging struct {
		Path         string `yaml:"path"`
		Pattern      string `yaml:"pattern"`
		MaxAge       string `yaml:"max_age"`
		RotationTime string `yaml:"rotation_time"`
		RotationSize string `yaml:"rotation_size"`
	}

	Microservice struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
)
