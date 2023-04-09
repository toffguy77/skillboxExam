package main

import (
	"github.com/toffguy77/statusPage/middleware/httpServer"
)

func main() {
	httpServer.Run("127.0.0.1:8888")
}
