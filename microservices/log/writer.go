package log

import "sync"

type Writer struct {
	// properties
	me sync.Mutex
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.me.Lock()
	// exclusive writes
	defer w.me.Unlock()

	// implement writer

	return 0, nil
}

func (w *Writer) Close() error {
	// close open channels
	return nil
}
