package contextz

import (
	"context"
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	reqID := r.Header.Get("X-Request-ID")
	ctx := context.WithValue(r.Context(), "requestID", reqID)

	// Pass ctx to functions that need access to the request ID.
	doSomething(ctx)

	fmt.Fprintf(w, "Request processed with ID: %s\n", reqID)
}

func doSomething(ctx context.Context) {
	reqID, ok := ctx.Value("requestID").(string)
	if !ok {
		reqID = "unknown"
	}
	fmt.Printf("Processing request with ID: %s\n", reqID)
}

func request() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

// In this example, when an HTTP request is handled, the handleRequest function extracts the request ID from the HTTP header and creates a new context with that request ID. This context is then passed to the doSomething function, which can access the request ID using ctx.Value.
