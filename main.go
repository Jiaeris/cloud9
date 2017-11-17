package main

import (
	_ "cloud9/config"
	_ "cloud9/models"

	"cloud9/ginserver"
)

func main() {
	ginserver.Start()
}
