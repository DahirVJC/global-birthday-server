package main

import "global-birthday-server/startup"

func main() {
	startup.LoadEnvironment()

	router := startup.SetupRouter()
	startup.StartRouter(router)
}
