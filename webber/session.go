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
	"math/rand"
	"time"
)


var sessionHeader string = "Session"
var sessionLength int = 0				// 0 - no expiration, otherwise seconds

// MakeSessionKey creates a session key, adds it as a cookie, and returns it
// so it can be saved in a session cache or db
//
func MakeSessionKey (w http.ResponseWriter) string {
	sessionKey := fmt.Sprintf("%d", rand.Uint64())
	cookie := http.Cookie{Name: sessionHeader, Value: sessionKey, Path: "/", HttpOnly: true, MaxAge: sessionLength}
	http.SetCookie(w, &cookie)
	return sessionKey
}

// GetSession returns a session if one exists
//
func GetSession ( r *http.Request) (bool, string) {
	session, err := r.Cookie(sessionHeader)
	if err == nil  {
		// check if expired
		if session.Expires.IsZero() || session.Expires.After(time.Now()) {
			return true, session.Value
		} else {
			fmt.Println ("expired session")
			return false, ""
		}
	} else {
		return false, ""
	}
}

// Clears any session
//
func ClearSession(w http.ResponseWriter) {
	cookie := http.Cookie{Name: sessionHeader, Value: "", Path: "/", HttpOnly: true, MaxAge: 0}
	http.SetCookie(w, &cookie)

}