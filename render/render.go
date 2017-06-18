package render

import (
	"html/template"
	"os"
	"time"

	. "github.com/toldjuuso/vertigo/databases/sqlx"

	"github.com/toldjuuso/timezone"
	unrolled "github.com/unrolled/render"
)

var R = unrolled.New(unrolled.Options{
	Funcs:  []template.FuncMap{helpers},
	Layout: "layout",
})

var helpers = template.FuncMap{
	// unescape unescapes HTML of s.
	// Used in templates such as "/post/display.tmpl"
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	// title renders post's Title as the HTML document's title.
	"title": func(t interface{}) string {
		post, exists := t.(Post)
		if exists {
			return post.Title
		}
		return Settings.Name
	},
	"blogname": func() string {
		if Settings.Name == "" {
			return "Blog in Go"
		}
		return Settings.Name
	},
	// description renders page description.
	// If none is defined, returns "Blog in Go" instead.
	"description": func() string {
		if Settings.Description == "" {
			return "Blog in Go"
		}
		return Settings.Description
	},
	// updated checks if post has been updated.
	"updated": func(p Post) bool {
		if p.Updated > p.Created {
			return true
		}
		return false
	},
	// date calculates unix date from d and offset in format: Monday, January 2, 2006 3:04PM (-0700 GMT)
	"date": func(d int64, offset int) string {
		return time.Unix(d, 0).UTC().In(time.FixedZone("", offset)).Format("Monday, January 2, 2006 3:04PM (-0700 GMT)")
	},
	"shortdate": func(d int64, offset int) string {
		return time.Unix(d, 0).UTC().In(time.FixedZone("", offset)).Format("02 Jan 2006")
	},
	"iso8601": func(d int64, offset int) string {
		return time.Unix(d, 0).UTC().In(time.FixedZone("", offset)).Format("2006-01-02T15:04:05+07:00")
	},
	// env returns environment variable of s.
	"env": func(s string) string {
		return os.Getenv(s)
	},
	// timezones returns all 416 valid IANA timezone locations.
	"timezones": func() []timezone.Timezone {
		return timezone.Locations
	},
	// returns whether registrations are allowed on user/login.tmpl
	"registerationsallowed": func() bool {
		return Settings.AllowRegistrations
	},
}
