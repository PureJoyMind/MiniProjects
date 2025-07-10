package async

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func CancellationWithContext() {
	// create a context with a timeout and cancel it at the end of every call
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// every context should close itself within the scope of it's creation
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(3)
	// showing the cancellation of a context and how it affects
	go fn_weak_ref(ctx, &wg)
	go fn_mid_ref(ctx, &wg)
	go fn_strong_ref(ctx, &wg)

	go func() {
		<-ctx.Done()
		fmt.Println("done")
	}()
	wg.Wait()
}

func fn_weak_ref(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second)
	}
	fmt.Println("done weak ref")
}

func fn_mid_ref(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 3 {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second)
	}
	fmt.Println("done mid ref")

}

func fn_strong_ref(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 1 {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second)
	}
	fmt.Println("done strong ref")
}
