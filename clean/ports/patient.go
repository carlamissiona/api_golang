package ports

import (
    "gomorganexam/clean/core/domain"
    "gomorganexam/clean/adapters/in"
    "gomorganexam/clean/adapters/out/db"
  
    _ "github.com/gorilla/sessions"
    "log"
    "os" 
    "encoding/csv"
)


 
type ServicePatientsAPI interface {
    GetTopConfirmed(observation_date string, max_result int32) domain.PatientObservation
    GetAllConfirmed(pagelimit int32) domain.PatientObservation
    StartSvc(hasRecords bool)
}

type Service struct {
    Data *db.Adapter
    Records[] byte
    Server *in.WebAdapter

}
var InstanceAPI *Service
func NewAPIService(data *db.Adapter, records[] byte, engine *in.WebAdapter)( *Service) {
     InstanceAPI =  &Service {
        Data: data,
        Records: nil,
        Server: engine, 
    }
 return InstanceAPI
}
func(svc *Service) Start(hasRecords bool) {
    var s Service; s = *svc
    if !hasRecords {
      // data_todb(&s); *svc = s ;
      // data_todb(&svc.Server.Msg, svc)
    }
    svc.Server.Start() 
}
  
func(svc *Service) GetTopConfirmed(observation_date string, max_result int32) {
    // svc.Database 


} 


// Redirect error
// parse csv
// save csv
// serve api
// docker
// fix HEX
// Finalize

// Grafana + Kubernetes
// upload http redirect 

func(svc *Service) GetAllConfirmed(pagelimit int32) {
    log.Println("No Implementation")
}

func data_todb(api *Service) {
    
	  // session := sessions.Default(c)
    
    csvfile, err := os.Open("data.csv")
    if err != nil {
        log.Printf("=>ERROR!! NO CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE " , err );   
        api.Server.Msg = "No Routes Created \nPls. Input CSV File To Continue."
        handle_error_csv_file(&api)
    }
     
    // remember to close the file at the end of the program
    defer csvfile.Close()
  
    // read csv values using csv.Reader
    csv_reader := csv.NewReader(csvfile)

    data, err := csv_reader.ReadAll()
    if err != nil {
        log.Printf("=>ERROR!! CANT READ CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE " , err ); 
        handle_error_csv_read(&api)
        
    }
    log.Println("API data")
    log.Printf("=>DEBUG!! API DATA")
    log.Println(data)
    log.Printf("=>END OF DEBUG!! API DATA")

}
