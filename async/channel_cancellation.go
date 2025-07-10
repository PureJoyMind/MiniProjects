package async

import (
	"fmt"
	"sync"
	"time"
)

func CancellationWithChannel() {
	// we can basically imitate a context using a channel which sends an empty struct
	// the main difference it that we can only close a channel once
	// but contexts can be closed many times without a panic
	ctx := make(chan struct{})
	go func() {
		defer close(ctx)
		time.Sleep(10 * time.Second)
	}()
	var wg sync.WaitGroup
	wg.Add(3)
	go fn_weak_ref_c(ctx, &wg)
	go fn_mid_ref_c(ctx, &wg)
	go fn_strong_ref_c(ctx, &wg)
	go func() {
		<-ctx
		fmt.Println("done")
	}()
	wg.Wait()
}

func fn_weak_ref_c(ctx <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		select {
		case <-ctx:
			return
		default:
		}
		time.Sleep(time.Second)
	}
	fmt.Println("done weak ref")
}

func fn_mid_ref_c(ctx <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 3 {
		select {
		case <-ctx:
			return
		default:
		}
		time.Sleep(time.Second)
	}
	fmt.Println("done mid ref")

}
func fn_strong_ref_c(ctx <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 1 {
		select {
		case <-ctx:
			return
		default:
		}
		time.Sleep(time.Second)
	}
	fmt.Println("done strong ref")
}
