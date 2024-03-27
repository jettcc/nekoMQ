package model

type RequestChannel struct {
	Var        interface{}
	ReturnChan chan interface{}
}

type ResponseChannel struct {
	Err      error
	Variable interface{}
}
