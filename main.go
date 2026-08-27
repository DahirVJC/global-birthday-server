package main

import "global-birthday-server/startup"

func main() {
	startup.LoadEnvironment()

	startup.InitializeDatabase()

	router := startup.SetupRouter()
	startup.StartRouter(router)
}
