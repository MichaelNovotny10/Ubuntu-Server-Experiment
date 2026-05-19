package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	batteryCache   string
	batteryCacheAt time.Time
	batteryMu      sync.Mutex
)

func batteryHandler(w http.ResponseWriter, r *http.Request) {
	batteryMu.Lock()
	defer batteryMu.Unlock()

	if time.Since(batteryCacheAt) > time.Minute {
		out, err := exec.Command("acpi", "-b").Output()
		if err != nil {
			batteryCache = "error running acpi: " + err.Error()
		} else {
			batteryCache = string(out)
		}
		batteryCacheAt = time.Now()
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(batteryCache))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/battery", batteryHandler)
	mux.Handle("/changelog/", http.StripPrefix("/changelog/", http.FileServer(http.Dir("./changelog"))))
	mux.Handle("/projects/", http.StripPrefix("/projects/", http.FileServer(http.Dir("./projects"))))
	mux.Handle("/progress/", http.StripPrefix("/progress/", http.FileServer(http.Dir("./progress"))))
	mux.Handle("/", http.FileServer(http.Dir("./index")))

	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	go func() {
		log.Printf("Server starting on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}
	log.Println("Server stopped")
}
