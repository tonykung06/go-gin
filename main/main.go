package main

import ()

func main() {
	r := registerRoutes()
	r.Run(":3000")
}
