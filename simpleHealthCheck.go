package k8sHealthCheck

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type SimpleHealthCheck struct {
	srv *http.Server
}

func NewSimpleHealthCheck(port string) *SimpleHealthCheck {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		if r.URL.Query().Get("full") != "1" {
			w.Write([]byte("OK\n")) //nolint: errcheck
			return
		}
	})
	return &SimpleHealthCheck{srv: &http.Server{Addr: port}}
}

func (shc *SimpleHealthCheck) Run() error {
	var err error
	go func() {
		err = shc.srv.ListenAndServe()
	}()

	if err != nil {
		return err
	}

	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
	<-kill

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := shc.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
