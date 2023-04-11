package main

import (
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/middleware/httpServer"
)

func main() {
	httpServer.Run(config.ServerURL)
}
