package paasio

import (
	"io"
	"sync"
)

type readWriteCounted struct {
	n     int64
	nops  int
	read  func(d []byte) (n int, err error)
	write func(d []byte) (n int, err error)
	m     sync.RWMutex
}

func (w *readWriteCounted) WriteCount() (n int64, nops int) {
	w.m.RLock()
	defer w.m.RUnlock()
	return w.n, w.nops
}

func (w *readWriteCounted) ReadCount() (n int64, nops int) {
	w.m.RLock()
	defer w.m.RUnlock()
	return w.n, w.nops
}

func (w *readWriteCounted) Read(d []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	n, err = w.read(d)
	if err != nil {
		return
	}
	w.n += int64(n)
	w.nops += 1
	return
}

func (w *readWriteCounted) Write(d []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()

	n, err = w.write(d)
	if err != nil {
		return
	}
	w.n += int64(n)
	w.nops += 1
	return
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &readWriteCounted{n: 0, nops: 0, write: writer.Write}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readWriteCounted{n: 0, nops: 0, read: reader.Read}
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounted{n: 0, nops: 0, read: readWriter.Read, write: readWriter.Write}
}
