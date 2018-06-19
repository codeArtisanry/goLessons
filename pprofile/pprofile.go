package main

import (
	"net/http"
	"net/http/pprof"
)

func main() {

	go func() {
		// logger := log.With(logger, "transport", "debug")

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
		// m.Handle("/metrics", stdprometheus.Handler())

		// logger.Log("addr", ":6060")
		http.ListenAndServe(":6060", m)
	}()
	for {
	}
}
