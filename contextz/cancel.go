package contextz

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Work canceled")
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed")
	}
}

func cancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel is called to release resources.

	go doWork(ctx)

	time.Sleep(1 * time.Second)
	cancel() // This will cancel the work.
	time.Sleep(1 * time.Second) // Give some time for cancellation to take effect.
}

// In this example, we create a context with a cancellation function using context.WithCancel. The doWork function watches for the cancellation signal using ctx.Done(). When the context is canceled (by calling cancel() in the main function), doWork will detect it and exit gracefully.