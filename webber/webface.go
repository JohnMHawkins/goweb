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
	"crypto/rand"
)
 
// WebHandler is the base interface for all web server types.
// It only supports the common methods GET and POST.  If you 
// need to support other methods, use WebFullHandler instead
//
type WebHandler interface {
	//Handler(w http.ResponseWriter, r *http.Request)
	HandleGet(w http.ResponseWriter, r *http.Request)
	HandlePost(w http.ResponseWriter, r *http.Request)
	GetBasePath() string
	GetName() string
}

// WebFullHanlder is the same as WebHandler, but implements all
// http methods.  Use this handler if you need to support PUT, PATCH,
// TRACE, etc.  
//
type WebFullHandler interface {
	//Handler(w http.ResponseWriter, r *http.Request)
	HandleGet(w http.ResponseWriter, r *http.Request)
	HandlePost(w http.ResponseWriter, r *http.Request)
	HandlePut(w http.ResponseWriter, r *http.Request)
	HandlePatch(w http.ResponseWriter, r *http.Request)
	HandleHead(w http.ResponseWriter, r *http.Request)
	HandleOptions(w http.ResponseWriter, r *http.Request)
	HandleDelete(w http.ResponseWriter, r *http.Request)
	HandleTrace(w http.ResponseWriter, r *http.Request)
	HandleConnect(w http.ResponseWriter, r *http.Request)
	GetBasePath() string
	GetName() string
}


// root dispatcher called by all WebHandlers to determine Method and dispatch to appropriate case handler
func DispatchMethod(h WebHandler, w http.ResponseWriter, r *http.Request) {
	
	switch {
		case r.Method == "GET":
			fmt.Println("Get called")
			h.HandleGet(w, r)
		case r.Method == "POST":
			fmt.Println("Post called")
			h.HandlePost(w, r)

		case r.Method == "HEAD":
			fmt.Println("Head called")
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		case r.Method == "TRACE":
			fmt.Println("Trace called")
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		case r.Method == "OPTIONS":
			fmt.Println("Options called")
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		case r.Method == "PUT":
			fmt.Println("Put called")
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		case r.Method == "PATCH":
			fmt.Println("Patch called")
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		case r.Method == "DELETE":
			fmt.Println("Delete called")
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		case r.Method == "CONNECT":
			fmt.Println("Connect called")
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}

}

// root dispatcher called by all WebHandlers to determine Method and dispatch to appropriate case handler
func DispatchFullMethod(h WebFullHandler, w http.ResponseWriter, r *http.Request) {
	
	switch {
		case r.Method == "GET":
			fmt.Println("Get called")
			h.HandleGet(w, r)
		case r.Method == "HEAD":
			fmt.Println("Head called")
			h.HandleHead(w, r)
		case r.Method == "TRACE":
			fmt.Println("Trace called")
			h.HandleTrace(w, r)
		case r.Method == "OPTIONS":
			fmt.Println("Options called")
			h.HandleOptions(w, r)
		case r.Method == "POST":
			fmt.Println("Post called")
			h.HandlePost(w, r)
		case r.Method == "PUT":
			fmt.Println("Put called")
			h.HandlePut(w, r)
		case r.Method == "PATCH":
			fmt.Println("Patch called")
			h.HandlePatch(w, r)
		case r.Method == "DELETE":
			fmt.Println("Delete called")
			h.HandleDelete(w, r)
		case r.Method == "CONNECT":
			fmt.Println("Connect called")
			h.HandleConnect(w, r)
	}

}

// SetSession adds a session header, generating a random key if needed.
//
// Parameters:
//	w http.ResponseWriter :	the response to add the session header to
//	sessionKey string : the session key, if one already exists, or an empty string to generate a new one
//	sessionValue string : the value written to the header
//
// Returns:
//	nothing
//
func SetSession (w http.ResponseWriter, sessionKey string, sessionValue string) {
	if len(sessionKey) == 0 {
		// generate a random session sessionKey
		var keyBuffer [16]byte

		_, err := rand.Read(keyBuffer[:])
		if err != nil {
			// huh?  
			// TODO : what is our recovery strategy here if we couldn't generate a session key?
			return
		}
		sessionKey = fmt.Sprintf("%x", keyBuffer)
	}

	w.Header().Add(sessionKey, sessionValue)

}
