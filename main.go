package main

import (
	"flag"
	"log"
)

var flagAddr string
var flagConfig string

func main() {
	flag.StringVar(&flagAddr, "addr", ":8081", "")
	flag.StringVar(&flagConfig, "config", "config.json", "")

	flag.Parse()

	LoadConfig(flagConfig)

	log.Printf("Starting, version='%s', build='%s', hash='%s'", CfgVersion, CfgBuildStamp, CfgGitHash)

	RunServer()
}
