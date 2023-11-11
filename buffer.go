package tsbuffer

import (
	"bytes"
	"sync"

	"github.com/pkg/errors"
)

var ErrBufferClosed = errors.New("buffer closed")

func New() *TSBuffer {
	return &TSBuffer{
		buffer: &bytes.Buffer{},
		closed: false,
		mu:     &sync.RWMutex{},
		cursor: 0,
	}
}

type TSBuffer struct {
	buffer *bytes.Buffer
	closed bool
	mu     *sync.RWMutex
	cursor int
}

func (sb *TSBuffer) Len() int {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.buffer.Len()
}

func (sb *TSBuffer) Read(p []byte) (n int, err error) {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.buffer.Read(p)
}
func (sb *TSBuffer) String() string {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.buffer.String()
}

func (sb *TSBuffer) Bytes() []byte {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	return sb.buffer.Bytes()
}

func (sb *TSBuffer) Close() error {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.closed = true
	return nil
}

func (sb *TSBuffer) Write(p []byte) (n int, err error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	if sb.closed {
		err = ErrBufferClosed
		return
	}
	return sb.buffer.Write(p)

}
func (sb *TSBuffer) WriteString(str string) (n int, err error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	if sb.closed {
		err = ErrBufferClosed
		return
	}
	return sb.buffer.WriteString(str)
}

func (sb *TSBuffer) Truncate(n int) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	if sb.closed {
		return
	}
	sb.buffer.Truncate(n)
}
