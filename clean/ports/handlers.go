package ports

import (
   "github.com/gin-gonic/gin"
    "net/http"

)
  
    
func handle_api_top_confirmed(api *Service, response string ){
  api.Server.Route("/api/get/top" , 1000 , func(c *gin.Context) {
     c.String( 301 , api.Server.Msg) }) 
}

   
func handle_api_error(api *Service, response string ){
  api.Server.Route("/server-error" , 500 , func(c *gin.Context) {
     c.String( 301 , api.Server.Msg) }) 
}


func handle_error_csv_read(){ 
  http.RedirectHandler("/server-error/", 301)
}
func handle_error_csv_file(){
    http.RedirectHandler("/server-error/", 301)
}