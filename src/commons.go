package src

import "regexp"

// Constants name in uppercase and snake_case
const (
	AGE_SESSION = 60 * 20 // In seconds
)

// Global variables with first letter in uppercase
var (
	AllData  = []TemplateData{}
	Accounts = []Account{}
	Sessions = []Session{}
	Likes    = []Like{}

	EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)
