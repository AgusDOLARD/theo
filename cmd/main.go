// shell main.go
package main

import (
	"flag"

	"AgusDOLARD/theo/theo"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "$HOME/.config/theo/theorc", "load the config file")
	flag.Parse()
}

func main() {
	s := theo.New(configPath)
	s.BuiltInCommand("cd", theo.Cd)
	s.BuiltInCommand("addpath", theo.AddPath)
	s.BuiltInCommand("setenv", theo.SetEnviromentVariable)
	s.Init()
}
