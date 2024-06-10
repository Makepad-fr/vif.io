package usernameverifierservice

import "strings"

// isTrimmedEmpty returns true if the given string is empty after trimmed
func isTrimmedEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsUserExists(username string) bool {
	return isTrimmedEmpty(username) || username != "kaanyagci"
}
