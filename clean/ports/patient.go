package ports

import (
	"gomorganexam/clean/adapters/in"
	"gomorganexam/clean/adapters/out/db"
	"gomorganexam/clean/core/domain"

	"encoding/csv"
	_ "github.com/gorilla/sessions"
	"log"
  "io"
	"os"
)

type ServicePatientsAPI interface {
	GetTopConfirmed(observation_date string, max_result int32) domain.PatientObservation
	GetAllRecovered(pagelimit int32) domain.PatientObservation
	StartSvc(hasRecords bool)
}

type Service struct {
	Data    *db.Adapter
	Records []byte
	Server  *in.WebAdapter
}

var InstanceAPI *Service

func NewAPIService(data *db.Adapter, records []byte, engine *in.WebAdapter) *Service {
	InstanceAPI = &Service{
		Data:    data,
		Records: nil,
		Server:  engine,
	}
	return InstanceAPI
}
func (svc *Service) Start(hasRecords bool) {
	var s Service
	s = *svc
	if !hasRecords {
		data_todb(&s)
		*svc = s
		// data_todb(&svc.Server.Msg, svc)
	}
	svc.Server.Start()
}

func (svc *Service) GetTopConfirmed(observation_date string, max_result int32) {
	var s Service
	s = *svc
	var response string = "{'g':'h', 'form':'data{}'}"
	handle_api_top_confirmed(&s, response)

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

func (svc *Service) GetAllConfirmed(pagelimit int32) {
	log.Println("No Implementation")
}

func data_todb(api *Service) {

	// session := sessions.Default(c)

	csvfile, err := os.Open("data.csv")
	if err != nil {
		log.Printf("=>ERROR!! NO CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE ", err)
		api.Server.Msg = "No Routes Created \nPls. Input CSV File To Continue."
		handle_error_csv_file()
	}

	// remember to close the file at the end of the program
	defer csvfile.Close()

	// read csv values using csv.Reader
	csv_reader := csv.NewReader(csvfile)
	// data, err := csv_reader.ReadAll()
	// _, err = csv_reader.ReadAll()
	// if err != nil {
	// 	log.Printf("=>ERROR!! CANT READ CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE ", err)
	// 	handle_error_csv_read()

	// }
	log.Println("API data")
	// log.Printf("=>DEBUG!! API DATA")
	// log.Println(data)
	// log.Printf("=>END OF DEBUG!! API DATA")
	log.Printf("=>END OF DEBUG!! Insert to DB %v", api.Data.Connect_DB())

    for {
        each_record, err := csv_reader.Read()
      	log.Printf("each_record %v", each_record)
      	// log.Println("each_record error %v", err)
        if err != nil || err == io.EOF {

            log.Fatal(err)
            break
        }
         var incr =0
         for value := range each_record {
            
           log.Printf("This is a record of matrix %s\n  %v", each_record[value] , incr)
            incr++
         }
    }
  
  
	api.Data.Execute(`INSERT INTO covid_observations (the_country_name, the_state_name,confirmed_no,deaths_no,recovered_no,observation_date) VALUES('Mainland China','Anhui',102,10,2,timestamp '2015-01-10 17:00:00' );`)

	log.Println(api.Data.Query(`Select * from covid_observations`))

}
