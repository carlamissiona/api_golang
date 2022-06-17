package main

import (
    "log"
    "os"
    _ "net/http"
    _ "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    _ "database/sql"

    _ "github.com/gin-gonic/gin"

    "gomorganexam/clean/adapters/in"
    "gomorganexam/clean/adapters/out/db"

    "gomorganexam/clean/ports"
    "github.com/joho/godotenv"
)

 
func main() {

    if (len(os.Args) == 2) {
        if (os.Args[1] == "minimal") {
            log.Println(os.Args[1])

            err:= godotenv.Load(".environ_development")

            if err != nil {
                log.Fatalf("failed reading env file: %v", err)
            }

            db_info:= os.Getenv("POSTGRES_URL")

            db:= db.NewAdapter("postgres", db_info)
            server:= in.NewServer()

            // csv record
            apisvc:= ports.NewAPIService(db, nil, server)
            apisvc.Start(checkInitData(db))
 
        }
    }
    log.Println("Hi Postg API")
}



func checkInitData(db_adapter * db.Adapter) bool {
   
    var count int32 = 0
    err:= db_adapter.Data.QueryRow(`SELECT Count(*) FROM covid_observations`).Scan(&count)
    if err != nil {
        log.Println("=>ERROR!! Check PG DATABASE DATA")
        log.Println(err)
        log.Println("=>END OF ERROR!! Check PG DATABASE DATA")
        return false
    }

    log.Printf("=>DEBUG!! Number of rows are %s\n", count)
    log.Printf("=>DEBUG!! Number of rows are %s\n", count)

    return true
}