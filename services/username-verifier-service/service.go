package usernameverifierservice

import (
	"log"
	"os"
	"strings"

	"github.com/Makepad-fr/vif.io/model"
)

var dbURL, dbUsername, dbPassword string

const dbURLEnvVarName = "DATABASE_URL"
const dbUsernameEnvVarName = "DATABASE_USER"
const dbPasswordEnvVarName = "DATABASE_PASSWORD"

func init() {
	v, exists := os.LookupEnv(dbURLEnvVarName)
	if !exists {
		log.Fatalf("Environment variable %s does not exists\n", dbURLEnvVarName)
	}
	dbURL = v
	v, exists = os.LookupEnv(dbUsernameEnvVarName)
	if !exists {
		log.Fatalf("Environment variable %s does not exists\n", dbUsernameEnvVarName)
	}
	dbUsername = v
	v, exists = os.LookupEnv(dbPasswordEnvVarName)
	if !exists {
		log.Fatalf("Environment variable %s does not exists\n", dbPasswordEnvVarName)
	}
	dbPassword = v
}

// isTrimmedEmpty returns true if the given string is empty after trimmed
func IsTrimmedEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsUserExists(username string) bool {
	// TODO: Complete function implementation
	return !IsTrimmedEmpty(username) && username == "kaanyagci"
}

func GetUserDetails(username string) (model.UserDetails, bool) {
	// TODO: Complete function definition
	return model.UserDetails{
		Username:       "kaanyagci",
		Fullname:       "Kaan Yagci",
		ProfilePicture: "https://via.placeholder.com/150",
		Description: `Lorem ipsum dolor sit amet consectetur adipisicing elit. Mollitia
              hic, porro nemo sint ipsa numquam, velit a explicabo dicta iste
              recusandae architecto voluptate repellat. Tempore quaerat impedit
              dicta explicabo fugit.`,
	}, true
}

func GetUserLinks(username string) []model.Link {
	// TODO: Complete user links
	return []model.Link{
		{
			Text: "My GitHub",
			Url:  "/kaanyagci/gh",
			Icon: "/static/images/github-mark.png",
		},
		{
			Text: "My GitHub",
			Url:  "/kaanyagci/gh",
			Icon: "/static/images/github-mark.png",
		},
	}
}
