package bloc4_config

func GetDefaultConfig() Config {

	return Config{
		Server: serverCnf{
			Host:              "127.0.0.1",
			Port:              "8080",
			Https:             false,
			CertPath:          "",
			ReadTimeout:       5,
			ReadHeaderTimeout: 3,
			WriteTimeout:      5,
			IdleTimeout:       30,
			MaxHeaderBytes:    1048576,
		},
	}
}
