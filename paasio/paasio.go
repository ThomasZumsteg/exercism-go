package paasio

import (
	"io"
)

//TestVersion is the unit tests that this will pass
const TestVersion = 1

//readCounter counts the read accesses and how any bytes are transfered.
type counter struct {
	iofunc func([]byte) (int, error)
	bytes  chan []int
}

/*newCounter creates new counter to record io operations.*/
func newCounter(iofunc func([]byte) (int, error)) *counter {
	bytes := make(chan []int, 1)
	bytes <- []int{}
	return &counter{iofunc: iofunc, bytes: bytes}
}

/*doIoFunc preforms the io function and stores the result.*/
func (c *counter) doIoFunc(p []byte) (n int, err error) {
	n, err = c.iofunc(p)
	c.bytes <- append(<-c.bytes, n)
	return
}

/*count counts the number of bytes and calls to the io operations.*/
func (c *counter) count() (int64, int) {
	bytesTotal := int64(0)
	bytes := <-c.bytes
	c.bytes <- bytes
	var i, b int
	for i, b = range bytes {
		bytesTotal += int64(b)
	}
	return bytesTotal, i + 1
}

//readCounter counts read io operations.
type readCounter struct{ counter }

/*NewReadCounter create a new readCounter.*/
func NewReadCounter(source io.Reader) *readCounter { return &readCounter{*newCounter(source.Read)} }

/*Read read from the data source.*/
func (reader *readCounter) Read(p []byte) (n int, err error) { return reader.doIoFunc(p) }

/*ReadCount report how many bytes and accesses have been read.*/
func (reader *readCounter) ReadCount() (int64, int) { return reader.count() }

//writeCounter counts write accesses and how many bytes are transfered.
type writeCounter struct{ counter }

/*NewWriteCounter create a new write writeCounter.*/
func NewWriteCounter(source io.Writer) *writeCounter { return &writeCounter{*newCounter(source.Write)} }

/*Write write to the data source.*/
func (writer *writeCounter) Write(p []byte) (n int, err error) { return writer.doIoFunc(p) }

/*WriteCount report how many bytes and accesses have been written.*/
func (writer *writeCounter) WriteCount() (int64, int) { return writer.count() }
