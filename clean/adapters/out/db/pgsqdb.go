package db
 
import (
    "database/sql"
    _ "encoding/json"
    _ "fmt"
    _ "os"
    "log"
    _ "net/http"
    _ "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

type DataAdapter interface {
    NewAdapter(driver_db, conn_str string)( *Adapter)
    Desc_DB()                    string
    Query(sqlStatement string)   *sql.Rows
    Execute(sqlStatement string) *sql.Result
    Connect_DB()              
}
 
type Adapter struct {
    Data  *sql.DB
    Driver string
    ConnUrl string
}
 
var Database_Instance  *sql.DB

func NewAdapter(driver_db, conn_str string)( *Adapter) {
    
    Database_Instance, err := sql.Open(driver_db, conn_str)
    if err != nil {
        panic(err)
    }
  
    defer Database_Instance.Close()

    err = Database_Instance.Ping()
    if err != nil {
            log.Fatalf("failed No DB connection %v", err)
        }
      

    return &Adapter {
        Data: Database_Instance,
      Driver: driver_db,
      ConnUrl: conn_str,
    }

}

func (a *Adapter)Desc_DB() string {
    return "Database_Instance"

}


func (a *Adapter)Connect_DB() (error) {
    var err error
    a.Data , err = sql.Open(a.Driver, a.ConnUrl)
    if err != nil {
        return err
    }
    return nil
}
 
func (a *Adapter)Query(sqlStatement string) *sql.Rows {
    var rows  *sql.Rows = nil
    var err error
     log.Printf("@Connection error -> \n %v", a.Connect_DB())
    err = a.Data.Ping()
    if err != nil {
  		log.Printf("Error In DB Connection: %v", err)
  	
    }
  
    rows, err = a.Data.Query(sqlStatement)
    if err != nil {
         log.Println("@Query error ->", err)
    }
    return rows
}
func(a *Adapter) Execute(sqlStatement string) sql.Result {
    var rows sql.Result = nil
        // sqlStatement := `SELECT * FROM TBL_PATIENTS)`
    log.Printf("@Connection error -> \n %v", a.Connect_DB())
    err := a.Data.Ping()
    rows, err = a.Data.Exec(sqlStatement)
    if err != nil {
        log.Println("@Execute error ->", err)

    }
    return rows
}
