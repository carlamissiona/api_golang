package ports

import (
	"gomorganexam/clean/adapters/in"
	"gomorganexam/clean/adapters/out/db"
	"gomorganexam/clean/core/domain"

	_ "encoding/csv"
	_ "github.com/gorilla/sessions"
	"log"
  _ "io"
	_"os"
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
func (svc *Service) Start(okStart bool) {
	var s Service
	s = *svc
	log.Printf("No Implementation %v", s)
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
