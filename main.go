package main

import "github.com/kofeebrian/short-url-server/routes"

func main() {
	r := routes.SetupRouters()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
