package domain

type Country struct {
	country   string
	confirmed int32
	deaths    int32
	recovered int32
}

type Observation struct {
	date      string
	countries Country
}
