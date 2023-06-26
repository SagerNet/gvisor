// automatically generated by stateify.

package ip

import (
	"github.com/sagernet/gvisor/pkg/state"
)

func (e *ErrMessageTooLong) StateTypeName() string {
	return "pkg/tcpip/network/internal/ip.ErrMessageTooLong"
}

func (e *ErrMessageTooLong) StateFields() []string {
	return []string{}
}

func (e *ErrMessageTooLong) beforeSave() {}

// +checklocksignore
func (e *ErrMessageTooLong) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
}

func (e *ErrMessageTooLong) afterLoad() {}

// +checklocksignore
func (e *ErrMessageTooLong) StateLoad(stateSourceObject state.Source) {
}

func (e *ErrNoMulticastPendingQueueBufferSpace) StateTypeName() string {
	return "pkg/tcpip/network/internal/ip.ErrNoMulticastPendingQueueBufferSpace"
}

func (e *ErrNoMulticastPendingQueueBufferSpace) StateFields() []string {
	return []string{}
}

func (e *ErrNoMulticastPendingQueueBufferSpace) beforeSave() {}

// +checklocksignore
func (e *ErrNoMulticastPendingQueueBufferSpace) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
}

func (e *ErrNoMulticastPendingQueueBufferSpace) afterLoad() {}

// +checklocksignore
func (e *ErrNoMulticastPendingQueueBufferSpace) StateLoad(stateSourceObject state.Source) {
}

func init() {
	state.Register((*ErrMessageTooLong)(nil))
	state.Register((*ErrNoMulticastPendingQueueBufferSpace)(nil))
}
