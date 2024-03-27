package model

type ChannelMsgRequest struct {
	Data       interface{}
	ReturnChan chan interface{}
}

type ChannelMsgResponse struct {
	Error error
	Data  interface{}
}
