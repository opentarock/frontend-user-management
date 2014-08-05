package logutil

import (
	"fmt"
	"log"
	"net/http"
)

func Print(r *http.Request, v ...interface{}) {
	logWithRequestInfo(r, fmt.Sprint(v))
}

func Printf(r *http.Request, format string, v ...interface{}) {
	logWithRequestInfo(r, fmt.Sprintf(format, v))
}

func logWithRequestInfo(r *http.Request, v string) {
	log.Printf("[%s] '%s' %s - %s", r.RemoteAddr, r.Method, r.RequestURI, v)
}
