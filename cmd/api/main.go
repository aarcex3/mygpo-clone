package main

import (
	"github.com/aarcex3/mygpo-clone/internals"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	internals.SetUpApp(app)
	app.Run("127.0.0.1:8000")
}
