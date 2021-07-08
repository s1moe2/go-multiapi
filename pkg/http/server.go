package http

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartServer takes an http.Server instance and takes care of starting it
// and handling shutdown routines
func StartServer(server *http.Server) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()
	log.Infof("Server started on %s", server.Addr)

	<-done
	log.Info("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Info("Server exited gracefully")
}
