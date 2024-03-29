package circular

import (
	"errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Buffer defintion datastructure
type Buffer struct {
	start int
	end   int
	size  int
	data  []*byte
}

// NewBuffer returns a new buffer
func NewBuffer(size int) *Buffer {
	return &Buffer{
		start: 0,
		end:   0,
		data:  make([]*byte, size),
		size:  size,
	}
}

// ReadByte read a byte from the buffer
func (b *Buffer) ReadByte() (byte, error) {
	returnByte := b.data[b.start]
	if returnByte == nil {
		return 0, errors.New("tried to read from empty buffer")
	}
	b.data[b.start] = nil

	if b.start != b.end {
		b.start = (b.start + 1) % b.size
	}
	return *returnByte, nil
}

// WriteByte writes a byte to the buffer
func (b *Buffer) WriteByte(c byte) error {
	if b.data[b.end] != nil {
		newEnd := (b.end + 1) % b.size

		if newEnd == b.start {
			return errors.New("tried to write to full buffer")
		}

		b.end = newEnd
	}
	b.data[b.end] = &c
	return nil
}

// Overwrite overwrites if the buffer is full
func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		b.data[b.start] = &c
		b.end = b.start
		b.start = (b.start + 1) % b.size
	}
}

// Reset resets the buffer
func (b *Buffer) Reset() {
	b.data = make([]*byte, len(b.data))
	b.start = 0
	b.end = 0
}
