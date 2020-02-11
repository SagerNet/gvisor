// automatically generated by stateify.

package unix

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *SocketOperations) beforeSave() {}
func (x *SocketOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("SendReceiveTimeout", &x.SendReceiveTimeout)
	m.Save("ep", &x.ep)
	m.Save("stype", &x.stype)
}

func (x *SocketOperations) afterLoad() {}
func (x *SocketOperations) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("SendReceiveTimeout", &x.SendReceiveTimeout)
	m.Load("ep", &x.ep)
	m.Load("stype", &x.stype)
}

func init() {
	state.Register("pkg/sentry/socket/unix.SocketOperations", (*SocketOperations)(nil), state.Fns{Save: (*SocketOperations).save, Load: (*SocketOperations).load})
}
