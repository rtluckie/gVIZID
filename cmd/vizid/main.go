package main

import "github.com/ryanl/vizid/cmd/vizid/commands"

var version = "dev"

func main() {
	commands.SetVersion(version)
	commands.Execute()
}
