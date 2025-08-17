package main

import "fmt"

var (
	AppName   = "GoApp"
	Version   = "none"
	Commit    = "none"
	BuildTime = "unknown"
)

func PrintVersion() string {
	return fmt.Sprintf("AppName: %s, Version: %s, Commit: %s, BuildTime: %s", AppName, Version, Commit, BuildTime)
}
