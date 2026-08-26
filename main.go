package main

import "global-birthday-server/startup"

func main() {
	router := startup.SetupRouter()
	startup.StartRouter(router)
}
