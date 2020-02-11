// automatically generated by stateify.

package pipe

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *buffer) beforeSave() {}
func (x *buffer) save(m state.Map) {
	x.beforeSave()
	m.Save("data", &x.data)
	m.Save("read", &x.read)
	m.Save("write", &x.write)
	m.Save("bufferEntry", &x.bufferEntry)
}

func (x *buffer) afterLoad() {}
func (x *buffer) load(m state.Map) {
	m.Load("data", &x.data)
	m.Load("read", &x.read)
	m.Load("write", &x.write)
	m.Load("bufferEntry", &x.bufferEntry)
}

func (x *bufferList) beforeSave() {}
func (x *bufferList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *bufferList) afterLoad() {}
func (x *bufferList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *bufferEntry) beforeSave() {}
func (x *bufferEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *bufferEntry) afterLoad() {}
func (x *bufferEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *inodeOperations) beforeSave() {}
func (x *inodeOperations) save(m state.Map) {
	x.beforeSave()
	m.Save("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Save("p", &x.p)
}

func (x *inodeOperations) afterLoad() {}
func (x *inodeOperations) load(m state.Map) {
	m.Load("InodeSimpleAttributes", &x.InodeSimpleAttributes)
	m.Load("p", &x.p)
}

func (x *Pipe) beforeSave() {}
func (x *Pipe) save(m state.Map) {
	x.beforeSave()
	m.Save("isNamed", &x.isNamed)
	m.Save("atomicIOBytes", &x.atomicIOBytes)
	m.Save("readers", &x.readers)
	m.Save("writers", &x.writers)
	m.Save("data", &x.data)
	m.Save("max", &x.max)
	m.Save("size", &x.size)
	m.Save("hadWriter", &x.hadWriter)
}

func (x *Pipe) afterLoad() {}
func (x *Pipe) load(m state.Map) {
	m.Load("isNamed", &x.isNamed)
	m.Load("atomicIOBytes", &x.atomicIOBytes)
	m.Load("readers", &x.readers)
	m.Load("writers", &x.writers)
	m.Load("data", &x.data)
	m.Load("max", &x.max)
	m.Load("size", &x.size)
	m.Load("hadWriter", &x.hadWriter)
}

func (x *Reader) beforeSave() {}
func (x *Reader) save(m state.Map) {
	x.beforeSave()
	m.Save("ReaderWriter", &x.ReaderWriter)
}

func (x *Reader) afterLoad() {}
func (x *Reader) load(m state.Map) {
	m.Load("ReaderWriter", &x.ReaderWriter)
}

func (x *ReaderWriter) beforeSave() {}
func (x *ReaderWriter) save(m state.Map) {
	x.beforeSave()
	m.Save("Pipe", &x.Pipe)
}

func (x *ReaderWriter) afterLoad() {}
func (x *ReaderWriter) load(m state.Map) {
	m.Load("Pipe", &x.Pipe)
}

func (x *Writer) beforeSave() {}
func (x *Writer) save(m state.Map) {
	x.beforeSave()
	m.Save("ReaderWriter", &x.ReaderWriter)
}

func (x *Writer) afterLoad() {}
func (x *Writer) load(m state.Map) {
	m.Load("ReaderWriter", &x.ReaderWriter)
}

func init() {
	state.Register("pkg/sentry/kernel/pipe.buffer", (*buffer)(nil), state.Fns{Save: (*buffer).save, Load: (*buffer).load})
	state.Register("pkg/sentry/kernel/pipe.bufferList", (*bufferList)(nil), state.Fns{Save: (*bufferList).save, Load: (*bufferList).load})
	state.Register("pkg/sentry/kernel/pipe.bufferEntry", (*bufferEntry)(nil), state.Fns{Save: (*bufferEntry).save, Load: (*bufferEntry).load})
	state.Register("pkg/sentry/kernel/pipe.inodeOperations", (*inodeOperations)(nil), state.Fns{Save: (*inodeOperations).save, Load: (*inodeOperations).load})
	state.Register("pkg/sentry/kernel/pipe.Pipe", (*Pipe)(nil), state.Fns{Save: (*Pipe).save, Load: (*Pipe).load})
	state.Register("pkg/sentry/kernel/pipe.Reader", (*Reader)(nil), state.Fns{Save: (*Reader).save, Load: (*Reader).load})
	state.Register("pkg/sentry/kernel/pipe.ReaderWriter", (*ReaderWriter)(nil), state.Fns{Save: (*ReaderWriter).save, Load: (*ReaderWriter).load})
	state.Register("pkg/sentry/kernel/pipe.Writer", (*Writer)(nil), state.Fns{Save: (*Writer).save, Load: (*Writer).load})
}
