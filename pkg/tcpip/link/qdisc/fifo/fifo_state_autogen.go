// automatically generated by stateify.

package fifo

import (
	"context"

	"github.com/sagernet/gvisor/pkg/state"
)

func (d *discipline) StateTypeName() string {
	return "pkg/tcpip/link/qdisc/fifo.discipline"
}

func (d *discipline) StateFields() []string {
	return []string{
		"dispatchers",
		"closed",
	}
}

func (d *discipline) beforeSave() {}

// +checklocksignore
func (d *discipline) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.dispatchers)
	stateSinkObject.Save(1, &d.closed)
}

func (d *discipline) afterLoad(context.Context) {}

// +checklocksignore
func (d *discipline) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.dispatchers)
	stateSourceObject.Load(1, &d.closed)
}

func (qd *queueDispatcher) StateTypeName() string {
	return "pkg/tcpip/link/qdisc/fifo.queueDispatcher"
}

func (qd *queueDispatcher) StateFields() []string {
	return []string{
		"lower",
		"queue",
	}
}

func (qd *queueDispatcher) beforeSave() {}

// +checklocksignore
func (qd *queueDispatcher) StateSave(stateSinkObject state.Sink) {
	qd.beforeSave()
	stateSinkObject.Save(0, &qd.lower)
	stateSinkObject.Save(1, &qd.queue)
}

func (qd *queueDispatcher) afterLoad(context.Context) {}

// +checklocksignore
func (qd *queueDispatcher) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &qd.lower)
	stateSourceObject.Load(1, &qd.queue)
}

func (pl *packetBufferCircularList) StateTypeName() string {
	return "pkg/tcpip/link/qdisc/fifo.packetBufferCircularList"
}

func (pl *packetBufferCircularList) StateFields() []string {
	return []string{
		"pbs",
		"head",
		"size",
	}
}

func (pl *packetBufferCircularList) beforeSave() {}

// +checklocksignore
func (pl *packetBufferCircularList) StateSave(stateSinkObject state.Sink) {
	pl.beforeSave()
	stateSinkObject.Save(0, &pl.pbs)
	stateSinkObject.Save(1, &pl.head)
	stateSinkObject.Save(2, &pl.size)
}

func (pl *packetBufferCircularList) afterLoad(context.Context) {}

// +checklocksignore
func (pl *packetBufferCircularList) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &pl.pbs)
	stateSourceObject.Load(1, &pl.head)
	stateSourceObject.Load(2, &pl.size)
}

func init() {
	state.Register((*discipline)(nil))
	state.Register((*queueDispatcher)(nil))
	state.Register((*packetBufferCircularList)(nil))
}
