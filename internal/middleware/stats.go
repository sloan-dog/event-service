package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	metrics "github.com/rcrowley/go-metrics"
)

const (
	latencyMetric = "latency"
	requestMetric = "request"
)

func StatsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		url := ctx.Request.URL.RequestURI()

		ctx.Next() // call the subsequential handlers (eventually we hit actual handler)

		latency := metrics.GetOrRegisterTimer(fmt.Sprintf("%s.%s.%s", latencyMetric, url, ctx.Request.Method), nil)
		latency.UpdateSince(start)

		// captures latency by url and method

		req := metrics.GetOrRegisterMeter(fmt.Sprintf("%s.%s.%s.%d", requestMetric, url, ctx.Request.Method, ctx.Writer.Status()), nil)
		req.Mark(1)

		// track count of requests by url, method, and status

	}
}

func GetStats() map[string]map[string]interface{} {
	return metrics.DefaultRegistry.GetAll()
}
