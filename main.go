package main

import (
    "log"
    "os"
  "io"
  
	   "encoding/csv"
     
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



func checkInitData(db_adapter *db.Adapter) bool {
   
    var count int32 = -1
    db_adapter.Connect_DB()
    err:= db_adapter.Data.QueryRow(`SELECT Count(*) FROM covid_observations`).Scan(&count)
    if err != nil {
        log.Printf("=>ERROR!! Check Initial Data Error in Query DB, %v ", err)
        return false
    }
    if (count == 0){
         log.Printf("= data_todb data_todb ")
      data_todb(db_adapter)
      
    }


    return true
}
 

func data_todb(db *db.Adapter) {

	csvfile, err := os.Open("./data.csv")
  
 

	if err != nil {
		log.Printf("=>ERROR!! NO CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE , %v ", err, csvfile)
		// api.Server.Msg = "No Routes Created \nPls. Input CSV File To Continue."
		// ports.handle_error_csv_file()
	} else {

  	db.Execute(`COPY covid_observations(SNo,ObservationDate,ProvinceState,CountryRegion,LastUpdate,Confirmed,Deaths,Recovered) FROM '~/goexammorgan/data.csv' DELIMITER ' 'CSV HEADER;`)
    
	if err != nil {
		log.Printf("=>ERROR!!DB Copy Error %v", err)
		// api.Server.Msg = "No Routes Created \nPls. Input CSV File To Continue."
		// ports.handle_error_csv_file()
	}
    }
//  	log.Println("CSV NO ERROR CSV NO ERROR , ")
//    	// remember to close the file at the end of the program
	
//   // read csv values using csv.Reader
// 	 csv_reader := csv.NewReader(csvfile)
     
//   // _ = csv.NewReader(csvfile)
// 	// log.Println("API data")
// 	// log.Printf("=>DEBUG!! API DATA")
// 	// log.Println(data)
// 	// log.Printf("=>END OF DEBUG!! API DATA")

   
//     err:= db.Connect_DB()
//     if err != nil {
//         log.Println("@ a.Connect_DB ->", err)
//         err = db.Data.Ping()
//         if err != nil {
//             log.Println("@DB Connection ping error ->", err)
    
//         }

//     }
//      csv_len, err := csv_reader.ReadAll()
//     if err != nil {
//         log.Println("Unknown lengrh of csv ")
//         // put logic here
//     }
//     csv_rows_length := len(csv_len)  
//     log.Printf("rows lengthhhhhh each_record %v", csv_rows_length)
     
    
//        var sqlstring = "INSERT INTO covid_observations (the_serial_id, observation_date, the_state_name, the_country_name,confirmed_no,deaths_no,recovered_no) VALUES( "
        
//             // db.Execute(`INSERT INTO covid_observations (the_serial_id, observation_date, the_state_name, the_country_name,confirmed_no,deaths_no,recovered_no)`+
//             //            `VALUES('` + Mainland China','Anhui',102,10,2,timestamp '2015-01-10 17:00:00' );`)
// log.Print(sqlstring)
//     var incr = 0 
//     log.Println("800 % incr")  ; 
//   //   for {
//   //       each_record, err := csv_reader.Read()
      
      
//   //       if err!=nil{
          
//   //           log.Println("@CSV Error End of File each_record------->", err)
            	
//   //           retry_csv()
             
//   //           break
//   //       }
//   //     	log.Printf("each_record %v", each_record)
//   //       sqlstring += ""
//   //       log.Println("Incr line no ", incr)
//   //       if ( ( incr >= 0 &&  800 % incr != 0) || incr <= 2){
//   //          index := 0
//   //          for value := range each_record {
//   //                if( index == 7 ){
//   //                   sqlstring +=  each_record[value] + " ) ,  \n"
                    
//   //                 }
//   //                 if(index == 0 ||index == 5 || index == 6 ){
//   //                   sqlstring +=  each_record[value] + ", "
                    
//   //                 }
//   //                 if(index == 1 ||index == 4  ){
                      
//   //                    sqlstring += " timestamp '" + each_record[value] + "', "
//   //                 } 
//   //                 if(index == 2 ||index == 3  ){
//   //                    sqlstring += " '" + each_record[value] + "', "
//   //                 }
//   //                 if (err != nil || err == io.EOF ) {
                  
//   //                   log.Printf("This is a record of matrix %s  %v", each_record[value] , incr)
                
//   //                }
//   //            index++
//   //          }
//   //        log.Print("sqlstring")    
    
//   //       }
//   //     incr++
        
 
//   // }


    

// 	// remember to close the file at the end of the program
// 	// defer csvfile.Close()

// 	// // // read csv values using csv.Reader
// 	// // csv_reader := csv.NewReader(csvfile)
//  // _ = csv.NewReader(csvfile)
// 	// log.Println("API data")
// 	// log.Printf("=>DEBUG!! API DATA")
// 	// log.Println(data)
// 	// log.Printf("=>END OF DEBUG!! API DATA")
	    
  
// 	db.Execute(`INSERT INTO covid_observations (the_country_name, the_state_name,confirmed_no,deaths_no,recovered_no,observation_date) VALUES('Mainland China','Anhui',102,10,2,timestamp '2015-01-10 17:00:00' );`)

// 	log.Println("api.Data.Query(`Select * from covid_observations`)")
// log.Println(db.Query(`Select * from covid_observations`))
//  tofile_sql(sqlstring)    
// } // error catche else

  }


func tofile_sql(str string){



    f, err := os.Create("sqlfile_to_db")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(str)

    if err2 != nil {
        log.Fatal(err2)
    }

    log.Println("done writing sql to file")


  
}
func retry_csv(){

csvfile_retry, err := os.Open("./data.csv")
  csv_reader := csv.NewReader(csvfile_retry)
  if err != nil {
		log.Printf("=>ERROR!! NO CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE , %v ", err, csvfile_retry)
		// api.Server.Msg = "No Routes Created \nPls. Input CSV File To Continue."
		// ports.handle_error_csv_file()
	}
  
       var sqlstring = "INSERT INTO covid_observations (the_serial_id, observation_date, the_state_name, the_country_name,confirmed_no,deaths_no,recovered_no) VALUES( "
        
            // db.Execute(`INSERT INTO covid_observations (the_serial_id, observation_date, the_state_name, the_country_name,confirmed_no,deaths_no,recovered_no)`+
            //            `VALUES('` + Mainland China','Anhui',102,10,2,timestamp '2015-01-10 17:00:00' );`)
log.Print(sqlstring)
    var incr = 0 
    log.Println("800 % incr")  ; 
    for {
        each_record, err := csv_reader.Read()
      
      
        if err!=nil{
          
            log.Println("@CSV each_record each_recordError End of File Finals------->", err) 
             
            break
        }
      	log.Printf("each_record %v", each_record)
        sqlstring += ""
        log.Println("Incr line no ", incr)
        if ( ( incr >0 &&  800 % incr != 0) || incr <= 2){
           index := 0
           for value := range each_record {
                 if( index == 7 ){
                    sqlstring +=  each_record[value] + " ) ,  \n"
                    
                  }
                  if(index == 0 ||index == 5 || index == 6 ){
                    sqlstring +=  each_record[value] + ", "
                    
                  }
                  if(index == 1 ||index == 4  ){
                      
                     sqlstring += " timestamp '" + each_record[value] + "', "
                  } 
                  if(index == 2 ||index == 3  ){
                     sqlstring += " '" + each_record[value] + "', "
                  }
                  if (err != nil || err == io.EOF ) {
                  
                    log.Printf("This is a record of matrix %s  %v", each_record[value] , incr)
                
                 }
             index++
           }
         log.Print("sqlstring")    
          // log.Print(sqlstring)
        }
      incr++
        

}

  
  tofile_sql(sqlstring)    

  }