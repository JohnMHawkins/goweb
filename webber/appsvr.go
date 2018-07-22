// webber - WebServer package
//
// Copyright (c) 2018 - John M. Hawkins <jmhawkins@msn.com>
//
// All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and 
// associated documentation files (the "Software"), to deal in the Software without restriction, 
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, 
// and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, 
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial 
// portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
// NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//

package webber

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type MethodHandler map[string]WebHandler

// AppServer is a webserver intended to support web applications
type AppServer struct {
	configFileName string
	Config ServerConfig
	BasePath string
	FileServerInst* FileServer
	Handlers map[string]WebHandler
	MethodHandlers map[string]map[string]WebHandler
}

// NewAppServer creates a new appserver with configuration information supplied by a file.  Will
// also create a FileServer that handles any paths not handled by the api server
//
// Parameters:
//	configFileName string : path to a file on the server to pull configuraiton information from.
//	                        If empty string, then the default of "file_server_config.json" will be used
//
//		configuration parameters:
//		root :  relative path from the executable to the start of the file root
//		apibase : path prefix for all handlers
//		DefaultFile : If no path is provided, this is the file returns (e.g. index.html)
//
// Returns:
//	*AppServer : the app server created
//
func NewAppServer(configFileName string) *AppServer {
	f := new(AppServer)
	f.MethodHandlers = make(map[string]map[string]WebHandler)

	// load configuration from file
	if (len(configFileName) == 0) {
		configFileName = "file_server_config.json"
	}

	f.configFileName = configFileName

	config, err := ioutil.ReadFile(configFileName)
	if err != nil {
		fmt.Println("Failed to load app server config file " + configFileName + ":", err)
	}

	jsonerr := json.Unmarshal(config, &f.Config )
	if jsonerr != nil {
		fmt.Println("Failed to parse file server config file : ", jsonerr)
	} else {
		fmt.Println("root = " + f.Config.Root)
		fmt.Println("apibase = " + f.Config.ApiBase)
	}

	f.FileServerInst = NewFileServerWithConfig("/", f.Config)

	f.Handlers = make(map[string]WebHandler)
	return f
}

// GetName returns the name of the server (AppServers are always 
// named "AppServer")
//
// Parmaeters: 
//	none
//
// Returns:
//	string : Always just returns "AppServer"
//
func (h AppServer) GetName() string {
	// AppServers don't have a unique name
	return "AppServer"
}

// GetBasePath returns the base path of the server (AppServers don't have base paths)
// they start at the root and have other handlers set underneath them
//
// Parmaeters: 
//	none
//
// Returns:
//	string : Always just returns "/"
//
func (h AppServer) GetBasePath() string {
	// AppServers don't have a base path othert than "/""
	return "/"
}

// Handler - the base handler for the AppServer.
//
func (h AppServer) Handler (w http.ResponseWriter, r *http.Request) {
	wasHandled := false
	urlPath := r.URL.Path
	l := len(urlPath)
	if l > 0 {
		if urlPath[l-1:l] != "/" {
			// tack on a trailing slash
			urlPath = urlPath + "/"
		}
		fmt.Println("appServer handler path=", urlPath)
		
		for p := range h.Handlers {
			fmt.Println("...checking ", p)
			if len(urlPath) >= len(p) &&	urlPath[:len(p)] == p {
				fmt.Println("...found it")
				wasHandled = true
				phf := h.Handlers[p]
				fmt.Println("...is type ",phf.GetName())
				DispatchMethod(phf, w, r)
				//phf.Handler(w, r)
			} 
		}
	}
	if !wasHandled {
		// not specific handler, assume it's a file
		DispatchMethod(h.FileServerInst, w, r)
	}

}


// RegisterHandler will add a new handler to the appServer
//
// Parameters:
//	handler WebHandler : the handler to add, it should implement the WebHandler interface
//
// Returns:
//	none
//
func (h AppServer) RegisterHandler(handler WebHandler) {
	basePath := handler.GetBasePath()
	h.Handlers[basePath] = handler
}



func (h AppServer) HandleGet (w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unknown path", http.StatusNotFound)

}

func (h AppServer) HandlePost (w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unknown path", http.StatusNotFound)

}
