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
}
func InitHttpserver(config *viper.Viper,dbHandler *sql.DB) httpServer{
runnersrepository:=repositories.NewRunnersRepository(dbHandler)
Resultrepository:=repositories.NewResultsRepository(dbHandler)

}
func (hs httpServer) start{

}