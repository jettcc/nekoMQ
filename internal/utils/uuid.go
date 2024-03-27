package utils

import (
	"sesm/global"

	"github.com/google/uuid"
)

var UuidChan = make(chan []byte, 1024)

func UuidFactory() {
	for {
		UuidChan <- cuuid()
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
