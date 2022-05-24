package main

import (
	"flag"
	"log"
	"os"

	"Hybrid-blog/helpers/config"
	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/server"
)

func main() {
	os.Setenv("SESSION_KEY", utility.RandStringBytes(9))
	log.Println("Byte Key Session", os.Getenv("SESSION_KEY"))
	env := flag.String("env", "local", "application runtime environment")
	flag.Parse()
	config.Loads("config.yml")
	config.SetEnv(*env)
	server.Route()
}
