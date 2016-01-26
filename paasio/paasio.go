package paasio

import (
	"io"
)

//TestVersion is the unit tests that this will pass
const TestVersion = 1

//readCounter counts the read accesses and how any bytes are transfered.
type readCounter struct {
	source io.Reader
	bytes  chan []int
}

//writeCounter counts write accesses and how many bytes are transfered.
type writeCounter struct {
	source io.Writer
	bytes  chan []int
}

/*NewReadCounter create a new readCounter.*/
func NewReadCounter(source io.Reader) *readCounter {
	c := make(chan []int, 1)
	c <- []int{}
	return &readCounter{source: source, bytes: c}
}

/*NewWriteCounter create a new write writeCounter.*/
func NewWriteCounter(source io.Writer) *writeCounter {
	c := make(chan []int, 1)
	c <- []int{}
	return &writeCounter{source: source, bytes: c}
}

/*WriteCount report how many bytes and accesses have been written.*/
func (writer *writeCounter) WriteCount() (int64, int) {
	bytesTotal := int64(0)
	bytes := <-writer.bytes
	writer.bytes <- bytes
	var i, b int
	for i, b = range bytes {
		bytesTotal += int64(b)
	}
	return bytesTotal, i + 1
}

/*ReadCount report how many bytes and accesses have been read.*/
func (reader *readCounter) ReadCount() (int64, int) {
	bytesTotal := int64(0)
	bytes := <-reader.bytes
	reader.bytes <- bytes
	var i, b int
	for i, b = range bytes {
		bytesTotal += int64(b)
	}
	return bytesTotal, i + 1
}

/*Read read from the data source.*/
func (reader *readCounter) Read(p []byte) (n int, err error) {
	n, err = reader.source.Read(p)
	reader.bytes <- append(<-reader.bytes, n)
	return
}

/*Write write to the data source.*/
func (writer *writeCounter) Write(p []byte) (n int, err error) {
	n, err = writer.source.Write(p)
	writer.bytes <- append(<-writer.bytes, n)
	return
}
