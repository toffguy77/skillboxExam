package main

import (
	"flag"
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/middleware/httpServer"
)

func init() {
	var path string
	flag.StringVar(&path, "path", "internal/config/config.json", "path to app configuration file")
	flag.Parse()
	config.LoadConfig(path)
}

func main() {
	httpServer.Run(config.Conf.ServerURL)
}
