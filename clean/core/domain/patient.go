package domain

type PatientObservation struct {
  Data Observation
}


func NewPatientObservation() PatientObservation {
	return PatientObservation{}
}



// func (pobv PatientObservation) GetTopConfirmed(observation_date string, max_result int32) Observation {
// 	return Observation 
// }

// func (pobv PatientObservation) GetAllConfirmed(orderby string) Observation {
// 	return Observation 
// }
