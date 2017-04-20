// Copyright 2017 Kevin Bayes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"flag"
	"log"
)

type Config struct {
	Header string
	Workdir string
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

	return c.Workdir + "/tmp"
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