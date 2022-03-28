package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	counter := initTotalRequestsCounter() // create vector counter

	router := gin.Default()

	sneedGroup := router.Group("/sneed")

	sneedGroup.Use(promCounterMiddleware(counter)) // attach counter as middleware

	sneedGroup.GET("/metrics", promMetricsHandler()) // expose metrics endpoint

	sneedGroup.GET("/feed", func(c *gin.Context) {
		c.String(200, "feed")
		c.FullPath()
	})

	sneedGroup.GET("/seed", func(c *gin.Context) {
		c.String(200, "seed")
		c.FullPath()
	})

	router.Run(":8080") // run on 0.0.0.0:8080
}

func promCounterMiddleware(counter *prometheus.CounterVec) func(*gin.Context) {
	return func(c *gin.Context) {
		route := c.FullPath()
		method := c.Request.Method
		// detect method and path of request, using them as labels for counter increment
		counter.WithLabelValues(route, method).Inc()
	}
}

func promMetricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func initTotalRequestsCounter() *prometheus.CounterVec {
	// Use Vec version as it allows multiple routes to use the same counter
	// by using the vector of labels specified.
	// If you don't do this. you have to have multiple counters with different names
	totalRequests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"route", "method"},
	)
	prometheus.Register(totalRequests)
	return totalRequests
}
