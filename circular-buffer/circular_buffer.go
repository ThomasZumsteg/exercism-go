package circular

import "errors"

//TestVersion is the unit tests that pass.
const TestVersion = 2

//Buffer is a circular ring buffer.
type Buffer struct {
	read, used int
	data       []byte
}

/*NewBuffer creates a new ring buffer of a certain size.*/
func NewBuffer(size int) *Buffer {
	var buff = Buffer{0, 0, make([]byte, size)}
	return &buff
}

/*ReadByte reads the oldest byte in the buffer,
cant read an empty buffer.*/
func (b *Buffer) ReadByte() (byte, error) {
	if b.used == 0 {
		return 0, errors.New("Buffer is empty")
	}
	read := b.data[b.read]
	b.read = (b.read + 1) % len(b.data)
	b.used--
	return read, nil
}

/*WriteByte writes to the buffer and woun't write to a full buffer.*/
func (b *Buffer) WriteByte(c byte) error {
	if b.used == len(b.data) {
		return errors.New("Buffer is full")
	}
	write := (b.read + b.used) % len(b.data)
	b.data[write] = c
	b.used++
	return nil
}

/*Overwrite writes to buffer even if it's full*/
func (b *Buffer) Overwrite(c byte) {
	if b.used == len(b.data) {
		b.ReadByte()
	}
	b.WriteByte(c)
}

/*Reset clears the buffer of all data.*/
func (b *Buffer) Reset() {
	//Doesn't actually clear the buffer
	b.used = 0
}
