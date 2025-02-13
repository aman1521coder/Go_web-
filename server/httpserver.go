package server
import (
"database/sql"
"log"
"runner_go/controllers"
"runner_go/repositories"
"github.com/spf13/viper"
"github.com/gin-gonic/gin"
)
type httpServer struct{
	config *viper.Viper
	router *gin.Engine
	runnersController  *controllers.runnersController
	resultController *controllers.resultController
}
func InitHttpserver(config *viper.Viper,dbHandler *sql.DB) httpServer{
runnersrepository:=repositories.NewRunnersRepository(dbHandler)
Resultrepository:=repositories.NewResultsRepository(dbHandler)
runnersService := services.NewRunnersService(
	runnersRepository, resultRepository)
	resultsService := services.NewResultsService(
	resultRepository, runnersRepository)
	runnersController := controllers.NewRunnersController(
	runnersService)
	resultsController := controllers.NewResultsController(
	resultsService)
	router := gin.Default()
	router.POST("/runner", runnersController.CreateRunner)
	router.PUT("/runner", runnersController.UpdateRunner)
	router.DELETE("/runner/:id",runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runner", runnersController.GetRunnersBatch)
	router.POST("/result", resultsController.CreateResult)
	router.DELETE("/result/:id",
	resultsController.DeleteResult)
	return HttpServer{
	config: config,
	router: router,
	runnersController: runnersController,
	resultsController: resultsController,
	}
	}



func (hs httpServer) start{
	err:=hs.router.Run(hs.config.GetString("http.server_address"))
	if err !=nil{
		log.Fatal("error while staring while  Http server:%v",err)
	}

}