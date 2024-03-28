package model

type ChannelMsgRequest struct {
	Data       any
	ReturnChan chan any
}

type ChannelMsgResponse struct {
	Error error
	Data  any
}
