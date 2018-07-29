// webbertut - Tutorial for using the webber package in Golang
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



//////////////////////////////////////////////////////
// This demonstrates use of the webber library
//
//

package main

import (
	"net/http"
	"jmh/goweb/webber"
//	"encoding/json"
//	"math/rand"
	"fmt"
)


// This is our api/auth handler, which will do a simple login and save session auth info
//

type AuthServer struct {
	basePath string
}

func NewAuthServer(basePath string) *AuthServer {
	f := new(AuthServer)
	f.basePath = "/" + basePath + "/"
	return f 
}

func (h AuthServer) Name() string {
	return "AuthServer"
}

func (h AuthServer) BasePath() string {
	return h.basePath
}

func (h AuthServer) Handler ( w http.ResponseWriter, r *http.Request) { 
	apiPath := r.URL.Path[len(h.basePath):]
	fmt.Println("AuthServer Handler  called for ", apiPath)
	webber.DispatchMethod(h, w, r);
}

func (h AuthServer) HandleGet (w http.ResponseWriter, r *http.Request) {
	apiPath := r.URL.Path[len(h.basePath):]
	pathVars := map[int]string{1:"a"}
	fmt.Println("AuthServer handling ", apiPath)
	pathParts, _ := webber.ParsePathAndQueryFlat(r, apiPath, pathVars )

	switch pathParts[0] {
	case "check":
		// get the session if one exists
		bHasSession, sessionKey := webber.GetSession(r)
		// ordinarily we would use this to look up the session in a db or memcache to 
		// get session information, instead we'll just write the session key back to the caller
		// as our example
		fmt.Println(bHasSession, sessionKey)
		fmt.Fprintf(w, "<html><body>The session key is %s</body></html>", sessionKey)
	case "logout":
		// in this case, we would clear the session entry in the db and the header itself
		webber.ClearSession(w)
		fmt.Fprintf(w, "ok")
	
	}
	
}

func (h AuthServer) HandlePost (w http.ResponseWriter, r *http.Request) {
	parseErr := r.ParseForm()
	if parseErr != nil {
		fmt.Println("error parsing login form: %s", parseErr)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "dog" && password == "bark" {
		// that's our login!  Go ahead and make a session
		// ordinarily, we'd create info about the login and store
		// it in a db or memcache under a key, but for this example,
		// we'll just create the key and set the header
		sk := webber.MakeSessionKey(w);
		fmt.Println("session key is ", sk)
		//w.Header().Add("Session", sk)
		fmt.Fprintf(w, "Success")
		return		
	} else {
		fmt.Println("invalid credentials")
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

}



type HikeInfo struct {
	Name string
	Length int
	Description string
}


// This is our api/hike handler, which will tell us simple things about hikes
//

type HikeServer struct {
	basePath string
}

func NewHikeServer(basePath string) *HikeServer {
	f := new(HikeServer)
	f.basePath = "/" + basePath + "/"
	return f 
}

func (h HikeServer) Name() string {
	return "HikeServer"
}

func (h HikeServer) BasePath() string {
	return h.basePath
}

func (h HikeServer) Handler ( w http.ResponseWriter, r *http.Request) { 
	apiPath := r.URL.Path[len(h.basePath):]
	fmt.Println("HikeServer Handler  called for ", apiPath)
	webber.DispatchMethod(h, w, r);
}

func (h HikeServer) HandleGet (w http.ResponseWriter, r *http.Request) {
	apiPath := r.URL.Path[len(h.basePath):]
	// api is in the form /api/hike/<hikename>/describe
	//                    /api/hike/<hikename>/length
	pathVars := map[int]string{0:"hike_name"}
	fmt.Println("HikeServer handling ", apiPath)
	pathParts, vars := webber.ParsePathAndQueryFlat(r, apiPath, pathVars )
	fmt.Println("pathparts = ", pathParts )

	hikeInfo := new(HikeInfo)
	hikeInfo.Name = vars["hike_name"]
	hikeInfo.Length = len(hikeInfo.Name)  // convenient, the hike is one mile per letter!
	hikeInfo.Description = "It's a nice hike.  Isn't every hike a nice hike?"

	if ( len(pathParts) == 0) {
		// just return the hike info
		webber.ReturnJson(w, hikeInfo)
	} else {
		switch pathParts[0] {
		case "describe":
			webber.ReturnJson(w, hikeInfo.Description)
		case "length":
			webber.ReturnJson(w, hikeInfo.Length)
		default:
			http.Error(w, "NYI", http.StatusNotImplemented)
		}
	
	}
	
}

func (h HikeServer) HandlePost (w http.ResponseWriter, r *http.Request) {
	parseErr := r.ParseForm()
	if parseErr != nil {
		fmt.Println("error parsing login form: %s", parseErr)
	}
	//username := r.FormValue("username")
	//password := r.FormValue("password")
	http.Error(w, "NYI", http.StatusNotImplemented)

}



// main func - we'll load our config, create an AppServer, and add a handler (for auth) to it,
// then start serving.
//
func main() {

	config := webber.LoadConfig("config.json")

	// create an App Server
	as := webber.NewAppServer(config)

	// or
	//defaultConfig := webber.DefaultConfig()
	//as := webber.NewAppServer(defaultConfig)

	// create our auth handler and assign it to <apibase>/auth
	auths := NewAuthServer(config.ApiBase + "/auth")
	as.RegisterHandler(auths)
	hikes := NewHikeServer(config.ApiBase + "/hike")
	as.RegisterHandler(hikes)

	// now start the server
	http.HandleFunc("/", as.Handler)
	http.ListenAndServe(config.Port, nil)
	
}