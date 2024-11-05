package bloc4_config

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-yaml"
)

//Todo: ADD VALIDATION SCHEMES FOR DIFFERENT CONFIG TYPES

/*
* This code responsability is threefold, search for the appropriate config files, read them, and load them into structures that represent various configs.
* As outlined in types.go, config files should be located within the "/config" subfolder, available from the current working directory
* /config/API/ hold the different API config files
 */
func ConfigIntoServer(config Config) *http.Server {
	var s http.Server
	s.Addr = config.Server.Host + ":" + config.Server.Port

	if config.Server.Https {
		//handle https and cert files
	}

	s.ReadTimeout = time.Duration(config.Server.ReadTimeout) * time.Second
	s.ReadHeaderTimeout = time.Duration(config.Server.ReadHeaderTimeout) * time.Second
	s.WriteTimeout = time.Duration(config.Server.WriteTimeout) * time.Second
	s.IdleTimeout = time.Duration(config.Server.IdleTimeout) * time.Second
	s.MaxHeaderBytes = config.Server.MaxHeaderBytes

	return &s
}

// Should add validation of configuration
func Load() (Config, error) {
	var config Config

	found_configs, err := SearchConfigFiles(DEFAULT_CONFIG_PATH)
	if err != nil {
		return config, err
	}

	if len(found_configs) == 0 {
		config.Server = GetDefaultConfig().Server
		return config, nil
	}

	for _, v := range found_configs {

		file, err := loadFile(v.path)
		yamlDecoder := yaml.NewDecoder(bytes.NewReader(*file), yaml.Strict())

		if err != nil {
			return config, err
		}

		switch v.kind {
		case cAPI:
			var api_config ApiCnf

			if err := yamlDecoder.Decode(&api_config); err != nil {
				return config, fmt.Errorf("error trying to parse %v :\n %w", v.path, err)
			}
			config.Apis = append(config.Apis, api_config)

		case cSERVER:
			var server_config ServerCnf

			if err := yamlDecoder.Decode(&server_config); err != nil {
				return config, fmt.Errorf("error trying to parse %v :\n %w", v.path, err)
			}

			config.Server = server_config
		}
	}

	return config, nil
}
