package in

import (
  // "net/http"
  "log"
  "github.com/gin-gonic/gin"
  // "github.com/gorilla/sessions"
  // "gomorganexam/clean/adapters/out/db"
)  

var Server *gin.Engine
var StatusMsg = "Status Ok"
type WebServerAdapter interface {
    CheckServer()
    Start()
}


type WebAdapter struct {
    HttpEngine *gin.Engine
    Msg string
    Response string
}

func NewServer()( *WebAdapter) {
    
    Server = gin.Default()

    Server.GET("/error", func(c * gin.Context) {
        c.String(503, StatusMsg)

    })

    return &WebAdapter {
        HttpEngine: Server,
        Msg: StatusMsg,
    }

}

func (h *WebAdapter)Start() {
    h.HttpEngine.Run() 
 
} 
 
func (h *WebAdapter)Route(route string, err_no int, http_hndl gin.HandlerFunc) {

    if(err_no == 1000) {
      
      h.HttpEngine.GET(route, http_hndl) 
    } else {
      log.Println("Error Print from session")
    }


}

 


// err := db.QueryRow("SELECT COUNT(*) FROM main_table").Scan(&count)
// switch {    
// case err != nil:
//     log.Fatal(err)
// default:
//     fmt.Printf("Number of rows are %s\n", count)
// }



