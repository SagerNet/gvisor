// automatically generated by stateify.

package ports

import (
	"context"

	"github.com/sagernet/gvisor/pkg/state"
)

func (f *Flags) StateTypeName() string {
	return "pkg/tcpip/ports.Flags"
}

func (f *Flags) StateFields() []string {
	return []string{
		"MostRecent",
		"LoadBalanced",
		"TupleOnly",
	}
}

func (f *Flags) beforeSave() {}

// +checklocksignore
func (f *Flags) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.MostRecent)
	stateSinkObject.Save(1, &f.LoadBalanced)
	stateSinkObject.Save(2, &f.TupleOnly)
}

func (f *Flags) afterLoad(context.Context) {}

// +checklocksignore
func (f *Flags) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.MostRecent)
	stateSourceObject.Load(1, &f.LoadBalanced)
	stateSourceObject.Load(2, &f.TupleOnly)
}

func (c *FlagCounter) StateTypeName() string {
	return "pkg/tcpip/ports.FlagCounter"
}

func (c *FlagCounter) StateFields() []string {
	return []string{
		"refs",
	}
}

func (c *FlagCounter) beforeSave() {}

// +checklocksignore
func (c *FlagCounter) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.refs)
}

func (c *FlagCounter) afterLoad(context.Context) {}

// +checklocksignore
func (c *FlagCounter) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.refs)
}

func (p *portDescriptor) StateTypeName() string {
	return "pkg/tcpip/ports.portDescriptor"
}

func (p *portDescriptor) StateFields() []string {
	return []string{
		"network",
		"transport",
		"port",
	}
}

func (p *portDescriptor) beforeSave() {}

// +checklocksignore
func (p *portDescriptor) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.network)
	stateSinkObject.Save(1, &p.transport)
	stateSinkObject.Save(2, &p.port)
}

func (p *portDescriptor) afterLoad(context.Context) {}

// +checklocksignore
func (p *portDescriptor) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.network)
	stateSourceObject.Load(1, &p.transport)
	stateSourceObject.Load(2, &p.port)
}

func (d *destination) StateTypeName() string {
	return "pkg/tcpip/ports.destination"
}

func (d *destination) StateFields() []string {
	return []string{
		"addr",
		"port",
	}
}

func (d *destination) beforeSave() {}

// +checklocksignore
func (d *destination) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.addr)
	stateSinkObject.Save(1, &d.port)
}

func (d *destination) afterLoad(context.Context) {}

// +checklocksignore
func (d *destination) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.addr)
	stateSourceObject.Load(1, &d.port)
}

func (pm *PortManager) StateTypeName() string {
	return "pkg/tcpip/ports.PortManager"
}

func (pm *PortManager) StateFields() []string {
	return []string{
		"allocatedPorts",
		"firstEphemeral",
		"numEphemeral",
	}
}

func (pm *PortManager) beforeSave() {}

// +checklocksignore
func (pm *PortManager) StateSave(stateSinkObject state.Sink) {
	pm.beforeSave()
	stateSinkObject.Save(0, &pm.allocatedPorts)
	stateSinkObject.Save(1, &pm.firstEphemeral)
	stateSinkObject.Save(2, &pm.numEphemeral)
}

func (pm *PortManager) afterLoad(context.Context) {}

// +checklocksignore
func (pm *PortManager) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &pm.allocatedPorts)
	stateSourceObject.Load(1, &pm.firstEphemeral)
	stateSourceObject.Load(2, &pm.numEphemeral)
}

func init() {
	state.Register((*Flags)(nil))
	state.Register((*FlagCounter)(nil))
	state.Register((*portDescriptor)(nil))
	state.Register((*destination)(nil))
	state.Register((*PortManager)(nil))
}
