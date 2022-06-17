package ports
import (
   _ "github.com/gin-gonic/gin"

  )


//Errors

func handle_api_top_confirmed(api *Service, date_observed, max_no ){
  api.Server.Route.("/api/get/top" , func(c *gin.Context) {
     c.String( 301 , api.Server.Msg  )
  })
}
func handle_api_error(api *Service, err Error){
  api.Server.Route.("/server-error" , func(c *gin.Context) {
       c.String( 301 , api.Server.Msg  )
  })
}

func handle_error_csv_read(){
   c.Redirect(301, "/server-error/")
}
func handle_error_csv_file(){
   c.Redirect(301, "/server-error/")
}