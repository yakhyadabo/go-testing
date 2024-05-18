package contextz

import (
	"context"
	"fmt"
	"time"
)

func deadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel() // Ensure cancel is called to release resources.

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work canceled due to deadline exceeded")
	}
}


// In this example, we create a context with a deadline of 3 seconds. If the work doesn’t complete within that time, the context will be canceled, and the “Work canceled due to deadline exceeded” message will be printed.