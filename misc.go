// This file contains bunch of miscful helper functions.
// The functions here are either too rare to be assiociated to some known file
// or are met more or less everywhere across the code.
package main

import (
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/sessions"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := gorm.Open("sqlite3", "./vertigo.db")

	if err != nil {
		panic(err)
	}

	// Here database and tables are created in case they do not exist yet.
	// If database or tables do exist, nothing will happen to the original ones.
	db.CreateTable(&User{})
	db.CreateTable(&Post{})
}

func sessionchecker() martini.Handler {
	return func(session sessions.Session) {
		data := session.Get("user")
		_, exists := data.(int64)
		if exists {
			return
		}
		session.Set("user", -1)
		return
	}
}

// Middleware function hooks the database to be accessible for Martini routes.
func middleware() martini.Handler {
	db, err := gorm.Open("sqlite3", "./vertigo.db")
	if err != nil {
		panic(err)
	}
	return func(c martini.Context) {
		c.Map(&db)
	}
}

// sessionIsAlive checks that session cookie with label "user" exists and is valid.
func sessionIsAlive(session sessions.Session) bool {
	data := session.Get("user")
	_, exists := data.(int64)
	if exists {
		return true
	}
	return false
}

// SessionRedirect in addition to sessionIsAlive makes HTTP redirection to user home.
// SessionRedirect is useful for redirecting from pages which are only visible when logged out,
// for example login and register pages.
func SessionRedirect(res http.ResponseWriter, req *http.Request, session sessions.Session) {
	if sessionIsAlive(session) {
		http.Redirect(res, req, "/user", 302)
	}
}

// ProtectedPage makes sure that the user is logged in. Use on pages which need authentication
// or which have to deal with user structure later on.
func ProtectedPage(res http.ResponseWriter, req *http.Request, session sessions.Session) {
	if !sessionIsAlive(session) {
		session.Delete("user")
		http.Redirect(res, req, "/", 302)
	}
}

// root returns HTTP request "root".
// For example, calling it with http.Request which has URL of /api/user/5348482a2142dfb84ca41085
// would return "api". This function is used to route both JSON API and frontend requests in the same function.
func root(req *http.Request) string {
	return strings.Split(strings.TrimPrefix(req.URL.String(), "/"), "/")[0]
}
