// automatically generated by stateify.

package waiter

import (
	"context"

	"github.com/sagernet/gvisor/pkg/state"
)

func (e *Entry) StateTypeName() string {
	return "pkg/waiter.Entry"
}

func (e *Entry) StateFields() []string {
	return []string{
		"waiterEntry",
		"eventListener",
		"mask",
	}
}

func (e *Entry) beforeSave() {}

// +checklocksignore
func (e *Entry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.waiterEntry)
	stateSinkObject.Save(1, &e.eventListener)
	stateSinkObject.Save(2, &e.mask)
}

func (e *Entry) afterLoad(context.Context) {}

// +checklocksignore
func (e *Entry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.waiterEntry)
	stateSourceObject.Load(1, &e.eventListener)
	stateSourceObject.Load(2, &e.mask)
}

func (q *Queue) StateTypeName() string {
	return "pkg/waiter.Queue"
}

func (q *Queue) StateFields() []string {
	return []string{
		"list",
	}
}

func (q *Queue) beforeSave() {}

// +checklocksignore
func (q *Queue) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.list)
}

func (q *Queue) afterLoad(context.Context) {}

// +checklocksignore
func (q *Queue) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.list)
}

func (l *waiterList) StateTypeName() string {
	return "pkg/waiter.waiterList"
}

func (l *waiterList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *waiterList) beforeSave() {}

// +checklocksignore
func (l *waiterList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *waiterList) afterLoad(context.Context) {}

// +checklocksignore
func (l *waiterList) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *waiterEntry) StateTypeName() string {
	return "pkg/waiter.waiterEntry"
}

func (e *waiterEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *waiterEntry) beforeSave() {}

// +checklocksignore
func (e *waiterEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *waiterEntry) afterLoad(context.Context) {}

// +checklocksignore
func (e *waiterEntry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*Entry)(nil))
	state.Register((*Queue)(nil))
	state.Register((*waiterList)(nil))
	state.Register((*waiterEntry)(nil))
}
