package config

import (
	"flag"
	"strings"
)

var Config = struct {
	AllowedUsersList []string
}{}

func init() {
	allowedUsers := flag.String("allowedUsers", "user", "Comma separated allowed user list")
	Config.AllowedUsersList = strings.Split(*allowedUsers, ",")
}
