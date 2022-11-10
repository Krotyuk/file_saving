package sequential_logger

import . "file_saving/interface"

type SequentialLogger struct {
	Logger
	messageChan chan string
	done        chan struct{}
	ErrCh       chan error
}

const buffer = 1000

func New(wrppedLogger Logger) *SequentialLogger {
	sl := &SequentialLogger{
		Logger:      wrppedLogger,
		messageChan: make(chan string, buffer),
		done:        make(chan struct{}),
		ErrCh:       make(chan error),
	}

	go func() {
		er := sl.save()
		if er != nil {
			sl.ErrCh <- er
		}
	}()

	return sl
}

func (sl *SequentialLogger) Log(message string) error {
	sl.messageChan <- message
	return nil
}

func (sl *SequentialLogger) save() error {
	for {
		select {
		case m := <-sl.messageChan:
			err := sl.Logger.Log(m)
			if err != nil {
				return err
			}
		case <-sl.done:
			return nil
		}
	}

}

func (sl *SequentialLogger) Close() error {
	close(sl.ErrCh)
	close(sl.done)
	close(sl.messageChan)

	for m := range sl.messageChan {
		err := sl.Logger.Log(m)
		if err != nil {
			return err
		}
	}
	err := sl.Logger.Close()
	if err != nil {
		return err
	}
	return nil
}
