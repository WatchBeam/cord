package cord

import (
	"encoding/json"

	"github.com/WatchBeam/cord/events"
)

// The Socket represents a connection to a Discord server. All methods on
// the socket are safe for concurrent use.
type Socket interface {
	// Send dispatches an event down the Discord socket. It returns an error
	// if there was any issue in sending it.
	Send(op Operation, data json.Marshaler) error

	// On attaches a handler to an event.
	On(h events.Handler)

	// On attaches a handler that's called once when an event happens.
	Once(h events.Handler)

	// Off detaches a previously-attached handler from an event.
	Off(h events.Handler)

	// Errs returns a channel of errors which may occur asynchronously
	// on the websocket.
	Errs() <-chan error

	// Frees resources associated with the socket.
	Close() error
}

// New creates a connection to the Discord servers. Options may be nil if
// you want to use the defaults.
func New(token string, options *WsOptions) Socket {
	if options == nil {
		options = &WsOptions{}
	}
	options.fillDefaults(token)

	ws := &Websocket{
		opts:   options,
		events: newEmitter(),
		errs:   make(chan error),
	}

	ws.start()

	return ws
}
