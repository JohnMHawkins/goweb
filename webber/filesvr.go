package webber

import (
	"fmt"
	"net/http"
	"io/ioutil"	
	"encoding/json"
	"path"
)

type FileServer struct {
	ConfigFileName string
	Config ServerConfig
	BasePath string
}

func NewFileServerWithConfig(basePathToHere string, config ServerConfig) *FileServer {
	f := new(FileServer)

	f.BasePath = basePathToHere
	f.Config = config

	return f
}

func NewFileServer(basePathToHere string, configFileName string) *FileServer {
	f := new(FileServer)

	f.BasePath = basePathToHere 

	// load configuration from file
	if (len(configFileName) == 0) {
		configFileName = "file_server_config.json"
	}

	f.ConfigFileName = configFileName

	config, err := ioutil.ReadFile(configFileName)
	if err != nil {
		fmt.Println("Failed to load file server config file " + configFileName + ":", err)
	}
	
	jsonerr := json.Unmarshal(config, &f.Config )
	if jsonerr != nil {
		fmt.Println("Failed to parse file server config file : ", jsonerr)
	} else {
		fmt.Println("root = " + f.Config.Root)
	}
	

	return f
}


func (h FileServer) GetName() string {
	return "FileServer"
}

func (h FileServer) GetBasePath() string {
	return h.BasePath;
}


func (h FileServer) HandleGet (w http.ResponseWriter, r *http.Request) {
	ourPath := r.URL.Path[len(h.BasePath):]
	fmt.Println("fileserver handleGet of ", ourPath)
	SetSession(w, "", "hello")
	if len(ourPath) == 0 {
		ourPath = h.Config.DefaultFile
	}
	filename := path.Join(h.Config.Root,  ourPath)
	fmt.Println("...fileserver handleGet looking for ", filename)
	body, err := ioutil.ReadFile(filename)
	if err == nil {
		w.Write(body)
	} else {
		// try adding html
		filename = filename + ".html"
		body, err := ioutil.ReadFile(filename)
		if err == nil {
			w.Write(body)
		} else {
			fmt.Println("Error: ", err)
			w.WriteHeader(404)
		}

	}

}


func (h FileServer) HandlePost (w http.ResponseWriter, r *http.Request) {

	// todo:  add uploader support
}

