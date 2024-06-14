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
func isTrimmedEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsUserExists(username string) bool {
	// TODO: Complete function implementation
	return !isTrimmedEmpty(username) && username == "kaanyagci"
}
