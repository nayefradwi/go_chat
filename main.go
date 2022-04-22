package main

import "gochat/config"

func main() {
	config.SetUpEnvironment()
	config.SetUpDatabaseConnection()
	config.SetupServer()
}
