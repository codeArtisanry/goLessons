package main

import (
	"log"
	"net/http"
	"net/http/pprof"
)

func main() {
	// 如果使用了默认的 http.DefaultServeMux（通常是代码直接使用 http.ListenAndServe("0.0.0.0:8000", nil)），只需要添加一行：
	// net/http/pprof 已经在 init()函数中通过 import 副作用完成默认 Handler 的注册
	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()

	// 如果你使用自定义的 Mux，则需要手动注册一些路由规则
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
