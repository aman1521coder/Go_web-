package controllers
import(
	
"encoding/json"
"io"
"log"
"net/http"
"runners_go/models"
"runners_go/services"
"github.com/gin-gonic/gin")
type RunnerController struct{
runnersService *services.RunnerService

}
func NewRunnersController(runnersService *services.RunnerService) *RunnerController{
	return &RunnerController{
		runnersService:runnersService,
	}
}

func (rh RunnersController) CreateRunner(ctx *gin.Context) {
	body,err:=io.ReadAll(ctx.Request.body)
	if err!=nil{
      ctx.AbortwithError(http.StatusInternalServerError,err)
	return
	}
	var runner models.Runner
	err:=json.Unmarshall(body,&runner)
	if err!=nil{
		log.Println("error while unmarshalling")
		ctx.AbortwithError(http.StatusInternalServerError,err)
		return
	}
response, responseErr := rh.runnersService.CreateRunner(&runner)
if responseErr != nil {
ctx.AbortWithStatusJSON(responseErr.Status,
responseErr)
return
}
ctx.JSON(http.StatusOK, response)

}
func (rh RunnersController) UpdateRunner(ctx *gin.Context) {
	//updating
	body,err:=io.ReadAll(ctx.Request.body)
	if err!=nil{
      ctx.AbortwithError(http.StatusInternalServerError,err)
	return
	}
	response, responseErr := rh.runnersService.UpdateRunner(&runner)
if responseErr != nil {
ctx.AbortWithStatusJSON(responseErr.Status,
responseErr)
return
}
ctx.JSON(http.StatusNoContent)

	


}
func (rh RunnersController) DeleteRunner(ctx *gin.Context) {
	params:=ctx.Request.URL.Query()
	id,err:=params.Get("id")
	if err!=nil{
		log.Fatal("error while reading id ",err)
		ctx.AbortwithError(http.StatusInternalServerError,err)
		return
	}
	response,responseErr:=rh.runnersService.DeleteRunner(id)
	if responseErr!=nil{
      log.Fatal("something wrong while deleting ")
		ctx.AbortwithError(http.StatusInternalServerError,responseErr)
		return
	}
	ctx.JSON(http.StatusNoContent)

}
func (rh RunnersController) GetRunner(ctx *gin.Context) {
 
}
func (rh RunnersController) GetRunnersBatch(ctx *gin.Context)
{

}