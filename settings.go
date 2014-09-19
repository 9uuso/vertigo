// Settings.go includes everything you would think site-wide settings need. It also contains a installation wizard
// route at the bottom of the file. You generally should not need to change anything in here.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"code.google.com/p/go-uuid/uuid"
	"github.com/martini-contrib/render"
)

// Vertigo struct is used as a site wide settings structure. Different from posts and person
// it is saved on local disk in JSON format.
// Firstrun and CookieHash are generated and controlled by the application and should not be
// rendered or made editable anywhere on the site.
type Vertigo struct {
	Name        string          `json:"name" form:"name" binding:"required"`
	Hostname    string          `json:"hostname" form:"hostname" binding:"required"`
	Firstrun    bool            `json:"firstrun"`
	CookieHash  string          `json:"cookiehash"`
	Description string          `json:"description" form:"description" binding:"required"`
	Mailer      MailgunSettings `json:"mailgun"`
}

// MailgunSettings holds the API keys necessary to send account recovery email.
// You can find the necessary values for these structures in https://mailgun.com/cp
type MailgunSettings struct {
	Domain     string `json:"mgdomain" form:"mgdomain" binding:"required"`
	PrivateKey string `json:"mgprikey" form:"mgprikey" binding:"required"`
	PublicKey  string `json:"mgpubkey" form:"mgpubkey" binding:"required"`
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // defining gomaxprocs is proven to add performance by few percentages
}

// Settings is a global variable which holds settings stored in the settings.json file.
// You can call it globally anywhere by simply using the Settings keyword. For example
// fmt.Println(Settings.Name) will print out your site's name.
// As mentioned in the Vertigo struct, be careful when dealing with the Firstun and CookieHash values.
var Settings *Vertigo = VertigoSettings()

// VertigoSettings populates the global namespace with data from settings.json.
// If the file does not exist, it creates it.
func VertigoSettings() *Vertigo {
	_, err := os.OpenFile("settings.json", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("settings.json")
	if err != nil {
		panic(err)
	}

	// If settings file is empty, we presume its a first run.
	if len(data) == 0 {
		var settings Vertigo
		settings.CookieHash = uuid.New()
		settings.Firstrun = true
		jsonconfig, err := json.Marshal(settings)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("settings.json", jsonconfig, 0600)
		if err != nil {
			panic(err)
		}
		return VertigoSettings()
	}

	var settings *Vertigo
	if err := json.Unmarshal(data, &settings); err != nil {
		panic(err)
	}
	return settings
}

// Save() or Settings.Save() is a method which replaces the global Settings structure with the structure is is called with.
// It has builtin variable declaration which prevents you from overwriting CookieHash field.
func (settings *Vertigo) Save() error {
	var old Vertigo
	data, err := ioutil.ReadFile("settings.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &old); err != nil {
		panic(err)
	}
	Settings = settings
	settings.CookieHash = old.CookieHash // this to assure that cookiehash cannot be overwritten even if system is hacked
	jsonconfig, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("settings.json", jsonconfig, 0600)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSettings is a route which updates the local .json settings file.
// It is supposed to be disabled after the first run. Therefore the JSON API route is not available for now.
func UpdateSettings(req *http.Request, res render.Render, settings Vertigo) {
	if Settings.Firstrun == false {
		log.Println("Somebody tried to change your local settings...")
		res.JSON(406, map[string]interface{}{"error": "You are not allowed to change the settings this time. :)"})
		return
	}
	settings.Firstrun = false
	err := settings.Save()
	if err != nil {
		log.Println(err)
		res.JSON(500, map[string]interface{}{"error": "Internal server error"})
		return
	}
	switch root(req) {
	case "api":
		res.JSON(200, map[string]interface{}{"success": "Settings were successfully saved"})
		return
	case "user":
		res.Redirect("/user/register", 302)
		return
	}
	res.JSON(500, map[string]interface{}{"error": "Internal server error"})
}
