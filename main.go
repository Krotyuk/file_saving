package main

import (
	"fmt"
	"log"
	"sync"

	"file_saving/file_logger"
	"file_saving/sequential_logger"
)

func main() {
	f, err := file_logger.New("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	sl := sequential_logger.New(f)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			_ = sl.Log(fmt.Sprintf("log %d", i))
			wg.Done()
			return
		}(i, wg)
	}
	wg.Wait()

	err = sl.Close()
	if err != nil {
		fmt.Println(err)
	}
	if errCn := <-sl.ErrCh; errCn != nil {
		fmt.Println(errCn)
	}

}
