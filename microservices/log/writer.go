package log

import "sync"

type Writer struct {
	// properties
	mu sync.Mutex
}

func (w Writer) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// implement writer

	return 0, nil
}

func (w Writer) Close() error {
	// close open channels
	return nil
}
