package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	g = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "todoist",
		Name:      "filter_task_items",
		Help:      "todoist item's amount by unique filter",
	}, []string{"name", "filter"})
)

func recordMetrics() {
	go func() {
		for {
			token := "test"
			res := todoist.search_query_by_filter(token)
			fmt.Println(res)
			g.With(prometheus.Labels{"name": "hogehoge", "filter": "hoge"}).Set(123)
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	r := prometheus.NewRegistry()
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	r.Register(g)

	recordMetrics()

	http.Handle("/metrics", handler)
	http.ListenAndServe(":2112", nil)
}
