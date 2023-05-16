package routers

import (
	"ta_microservice_auth/app/controllers"
	"ta_microservice_auth/app/middleware"

	"github.com/gin-gonic/gin"
)

type API struct {
	RepoAuth   controllers.AuthRepo
	RepoSeeder controllers.DbSeeder
}

func SetupRouter(RepoAuth controllers.AuthRepo, RepoSeeder controllers.DbSeeder) *gin.Engine {
	r := gin.New()
	api := API{
		RepoAuth, RepoSeeder,
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	protectedRouter := r.Group("/api/auth")
	protectedRouter.Use(middleware.ApiKey(), middleware.ReqJson())
	protectedRouter.POST("/login", api.RepoAuth.Login)
	protectedRouter.POST("/refresh", api.RepoAuth.RefreshToken)
	// protectedRouter.POST("/seeder", api.RepoSeeder.RunSeeder)

	SeederRoter := r.Group("/api")
	SeederRoter.Use(middleware.ApiKey(), middleware.ReqJson(), middleware.Jwt())
	SeederRoter.POST("/register", api.RepoAuth.CreateAnggota)
	SeederRoter.GET("/get-all", api.RepoAuth.GetAllAnggota)
	SeederRoter.GET("/total", api.RepoAuth.TotalAnggota)

	protectedRouter.POST("/seeder", api.RepoSeeder.RunSeeder)
	return r
}
