package main

import (
	"ta_microservice_auth/app/controllers"
	"ta_microservice_auth/app/routers"
)

func main() {

	packageAuth := controllers.NewAuth()
	packageSeeder := controllers.NewDbSeerder()

	r := routers.SetupRouter(*packageAuth, *packageSeeder)
	_ = r.Run(":8080")
}
