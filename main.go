package main
import (
	"log"

	"runners_go/config"
"runners_go/server"
_ "github.com/lib/pq"
) 
func main(){
	
config := config.InitConfig("runners")
log.Println("Initializing database")
dbHandler := server.InitDatabase(config)
log.Println("Initializing HTTP server")
httpServer := server.InitHttpServer(config, dbHandler)
httpServer.Start()


}
