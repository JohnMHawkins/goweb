package webber

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type ServerConfig struct {
	Port string
	WWWRoot string			// path to wwwroot on server for fileserver.  If empty, no file server
	DefaultFile string		// default html file, e.g. index.html
	ApiBase string			// base url path to start of api, e.g. /api
	FileBase string			// base url path to files, e.g. /files
}


// DefaultConfig creates a Server config with all defaults
//
func DefaultConfig () *ServerConfig {
	config := new(ServerConfig)

	config.Port = "80"
	config.WWWRoot = "wwwroot"
	config.DefaultFile = "index.html"
	config.ApiBase = "api"
	config.FileBase = "/"

	return config
}

// LoadConfig reads a json config file, parses it, and fills in any defaults.  
//
// Parameters:
//	configFile string : path to file on the server to pull config information from.  if empty, will
//						use the default of "config.json"
//
// Returns:
//	*ServerConfig : the server config created
//
func LoadConfig (configFile string) *ServerConfig {
	config := new(ServerConfig)

	// load configuration from file
	if (len(configFile) == 0) {
		configFile = "file_server_config.json"
	}

	configJson, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Failed to load app server config file " + configFile + ":", err)
	}

	jsonerr := json.Unmarshal(configJson, config )
	if jsonerr != nil {
		fmt.Println("Failed to parse file server config file : ", jsonerr)
	} else {
		fmt.Println("root = " + config.WWWRoot)
		fmt.Println("apibase = " + config.ApiBase)
	}

	return config
}