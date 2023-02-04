package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	if err := utils.InitConfig(); err != nil {
		panic(err.Error())
	}

	if err := utils.InitMySQL(); err != nil {
		panic(err.Error())
	}

	r := router.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8081")
}
