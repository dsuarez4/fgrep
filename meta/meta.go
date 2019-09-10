package meta

import "fmt"

var GITCOMMIT string

var VERSION string

// Release is the current release number, should match the value in
// $GOPATH/src/github.com/dsuarez4/fgrep/VERSION
const Release = "0.0.1"

func Meta() string {
	version, commit := VERSION, GITCOMMIT

	if commit == "" || version == "" {
		version, commit = Release, "master"
	}
	return fmt.Sprintf("fgrep version %v, built from %v", version, commit)
}