package main

import (
	"project/libraryManagementServer/server"

	"project/libraryManagementServer/conf"

	"project/libraryManagementServer/util"
)

const version = "0.0.1"

func main() {

	util.Log("LIBRARY SERVER STARTED, VERSION: ", version)

	printConfig()

	server.Run()
}

// PrintConfig prints the conifiguration
func printConfig() {

	var m map[string]string

	m["LogPath"] = conf.Conf.LogPath

	util.Logf("Configuration >> %v", m)
}
