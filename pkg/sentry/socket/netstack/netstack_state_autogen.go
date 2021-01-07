// automatically generated by stateify.

package netstack

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *SocketOperations) StateTypeName() string {
	return "pkg/sentry/socket/netstack.SocketOperations"
}

func (s *SocketOperations) StateFields() []string {
	return []string{
		"socketOpsCommon",
	}
}

func (s *SocketOperations) beforeSave() {}

func (s *SocketOperations) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.socketOpsCommon)
}

func (s *SocketOperations) afterLoad() {}

func (s *SocketOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.socketOpsCommon)
}

func (s *socketOpsCommon) StateTypeName() string {
	return "pkg/sentry/socket/netstack.socketOpsCommon"
}

func (s *socketOpsCommon) StateFields() []string {
	return []string{
		"SendReceiveTimeout",
		"Queue",
		"family",
		"Endpoint",
		"skType",
		"protocol",
		"readCM",
		"sender",
		"linkPacketInfo",
		"sockOptTimestamp",
		"timestampValid",
		"timestampNS",
		"sockOptInq",
	}
}

func (s *socketOpsCommon) beforeSave() {}

func (s *socketOpsCommon) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.SendReceiveTimeout)
	stateSinkObject.Save(1, &s.Queue)
	stateSinkObject.Save(2, &s.family)
	stateSinkObject.Save(3, &s.Endpoint)
	stateSinkObject.Save(4, &s.skType)
	stateSinkObject.Save(5, &s.protocol)
	stateSinkObject.Save(6, &s.readCM)
	stateSinkObject.Save(7, &s.sender)
	stateSinkObject.Save(8, &s.linkPacketInfo)
	stateSinkObject.Save(9, &s.sockOptTimestamp)
	stateSinkObject.Save(10, &s.timestampValid)
	stateSinkObject.Save(11, &s.timestampNS)
	stateSinkObject.Save(12, &s.sockOptInq)
}

func (s *socketOpsCommon) afterLoad() {}

func (s *socketOpsCommon) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.SendReceiveTimeout)
	stateSourceObject.Load(1, &s.Queue)
	stateSourceObject.Load(2, &s.family)
	stateSourceObject.Load(3, &s.Endpoint)
	stateSourceObject.Load(4, &s.skType)
	stateSourceObject.Load(5, &s.protocol)
	stateSourceObject.Load(6, &s.readCM)
	stateSourceObject.Load(7, &s.sender)
	stateSourceObject.Load(8, &s.linkPacketInfo)
	stateSourceObject.Load(9, &s.sockOptTimestamp)
	stateSourceObject.Load(10, &s.timestampValid)
	stateSourceObject.Load(11, &s.timestampNS)
	stateSourceObject.Load(12, &s.sockOptInq)
}

func (s *SocketVFS2) StateTypeName() string {
	return "pkg/sentry/socket/netstack.SocketVFS2"
}

func (s *SocketVFS2) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"LockFD",
		"socketOpsCommon",
	}
}

func (s *SocketVFS2) beforeSave() {}

func (s *SocketVFS2) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.vfsfd)
	stateSinkObject.Save(1, &s.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &s.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &s.LockFD)
	stateSinkObject.Save(4, &s.socketOpsCommon)
}

func (s *SocketVFS2) afterLoad() {}

func (s *SocketVFS2) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.vfsfd)
	stateSourceObject.Load(1, &s.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &s.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &s.LockFD)
	stateSourceObject.Load(4, &s.socketOpsCommon)
}

func (s *Stack) StateTypeName() string {
	return "pkg/sentry/socket/netstack.Stack"
}

func (s *Stack) StateFields() []string {
	return []string{}
}

func (s *Stack) beforeSave() {}

func (s *Stack) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
}

func (s *Stack) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.AfterLoad(s.afterLoad)
}

func init() {
	state.Register((*SocketOperations)(nil))
	state.Register((*socketOpsCommon)(nil))
	state.Register((*SocketVFS2)(nil))
	state.Register((*Stack)(nil))
}
