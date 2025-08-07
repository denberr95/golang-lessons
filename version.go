package main

import "fmt"

var (
	AppName   = "GoApp"
	Version   = "none"
	Commit    = "none"
	BuildTime = "unknown"
)

func FullVersion() string {
	return fmt.Sprintf("appname: %s, version:%s, commit: %s, builtTime: %s)", AppName, Version, Commit, BuildTime)
}

func ShortVersion() string {
	return fmt.Sprintf("%s %s", AppName, Version)
}
