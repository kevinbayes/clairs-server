package config

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
	"flag"
)

type Config struct {
	Header string
	Server struct {
		Host string
		Port string
		Filepath string
	}
	Database struct {
		  Vendor string
		  Database string
		  Host string
		  Port string
		  Schema string
		  Username string
		  Password string
	}
}

var config *Config;


func GetConfig() *Config {

	if(config == nil) {

		configFile := flag.String("config", "./src/default.yml", "Location of the config file.")

		fmt.Println("Loading config from " + *configFile)

		data, err := ioutil.ReadFile(*configFile);

		if (err != nil) {

			fmt.Printf("Could not load config: %s \n", err)
			panic(err);
		}

		fmt.Println("Loaded file data:" + string(data[:]))

		_config := &Config{}

		parseErr := yaml.Unmarshal(data, _config)

		if (parseErr != nil) {
			fmt.Printf("Could not parse config: %s \n", parseErr)
			panic(err);
		}

		config = _config
	}

	return config
}