package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"flag"
	"log"
)

type Config struct {
	Header string
	WorkDir string
	Server struct {
		Host string
		Port string
		Filepath string
	}
	Clair struct {
		Protocol string
		Host string
		Port string
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

func (c *Config) TmpDir() string {

	return c.WorkDir + "/tmp"
}

var config *Config;



func GetConfig() *Config {

	if(config == nil) {

		configFile := flag.String("config", "./src/default.yml", "Location of the config file.")

		log.Printf("Loading config from " + *configFile)

		data, err := ioutil.ReadFile(*configFile);

		if (err != nil) {

			log.Panic("Could not load config: %s \n", err)
			panic(err);
		}

		log.Printf("Loaded file data:" + string(data[:]))

		_config := &Config{}

		parseErr := yaml.Unmarshal(data, _config)

		if (parseErr != nil) {
			log.Panic("Could not parse config: %s \n", parseErr)
			panic(err);
		}

		config = _config
	}

	return config
}