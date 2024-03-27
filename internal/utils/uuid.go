package utils

import (
	"sesm/global"

	"github.com/google/uuid"
)

var uuidChan = make(chan []byte, 1024)

func uuidFactory() {
	for {
		uuidChan <- cuuid()
	}
}

func cuuid() []byte {
	uid, err := uuid.NewUUID()
	if err != nil {
		global.Sugar.Warn("uuid.NewUUID() failed: ", err.Error())
		uid = uuid.New()
	}
	return uid[:]
}
