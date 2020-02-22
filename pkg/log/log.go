package log

import (
	"log"
	"net/http"
	"time"
)

var (
	logger *log.Logger
)

// Logger is interface ofr logging.
type Logger interface {
	Print()
}

/*
Print displays logs in console
*/
func Print(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}
