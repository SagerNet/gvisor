// automatically generated by stateify.

package tpuproxy

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (dev *vfioDevice) StateTypeName() string {
	return "pkg/sentry/devices/tpuproxy.vfioDevice"
}

func (dev *vfioDevice) StateFields() []string {
	return []string{
		"mu",
		"minor",
	}
}

func (dev *vfioDevice) beforeSave() {}

// +checklocksignore
func (dev *vfioDevice) StateSave(stateSinkObject state.Sink) {
	dev.beforeSave()
	stateSinkObject.Save(0, &dev.mu)
	stateSinkObject.Save(1, &dev.minor)
}

func (dev *vfioDevice) afterLoad() {}

// +checklocksignore
func (dev *vfioDevice) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &dev.mu)
	stateSourceObject.Load(1, &dev.minor)
}

func init() {
	state.Register((*vfioDevice)(nil))
}
