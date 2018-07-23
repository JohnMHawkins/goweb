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
)
 
// WebHandler is the base interface for all web server types.
// It only supports the common methods GET and POST.  
//
type WebHandler interface {
	HandleGet(w http.ResponseWriter, r *http.Request)
	HandlePost(w http.ResponseWriter, r *http.Request)
	BasePath() string
	Name() string
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


