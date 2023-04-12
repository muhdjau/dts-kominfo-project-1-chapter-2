package main

import (
	"project-1-chapter-2/config"
	"project-1-chapter-2/routers"
)

func main() {
	config.ConnectDB()

	routers.StartServer().Run(":80")
}
