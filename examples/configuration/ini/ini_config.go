package main

import (
	"fmt"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

type Section struct {
	Enabled bool
	Path    string
}

func main() {
	config := Section{}

	err := gcfg.ReadFileInto(&config, "conf.ini")

	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
		os.Exit(1)
	}

	fmt.Println(config.Enabled)
	fmt.Println(config.Path)
}
