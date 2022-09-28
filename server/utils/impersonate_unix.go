// A demonstration example for http://stackoverflow.com/a/26124494
// It runs a goroutine locked to an OS thread on Windows
// then impersonates that thread as another user using its name
// and plaintext password, then reverts to the default security
// context before detaching from its OS thread.

// +build !windows

package utils

func WindowsLogin(user string, pass string) bool {
	return false
}
