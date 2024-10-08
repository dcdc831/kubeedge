package clients

import (
	"github.com/kubeedge/beehive/pkg/core/model"
)

// Adapter is a web socket client interface
type Adapter interface {
	Init() error
	UnInit()
	// async mode
	Send(message model.Message) error
	Receive() (model.Message, error)

	// Notify auth info
	Notify(authInfo map[string]string)
}
