package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "healthy\n")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		numCPUs := runtime.NumCPU()
		runtime.GOMAXPROCS(numCPUs)

		// Busy wait to emulate heavy load on all CPUs
		var wg sync.WaitGroup
		wg.Add(numCPUs)
		end := time.Now().Add(5 * time.Second)
		for i := 0; i < numCPUs; i++ {
			go func() {
				defer wg.Done()
				for time.Now().Before(end) {
					_ = 1 // Minimal operation to keep CPU occupied
				}
			}()
		}

		wg.Wait()

		fmt.Fprintf(w, "hello world\n")
	})

	http.ListenAndServe(":80", nil)
}
