package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"hw09/orders/internal/tracer"
	"net/http"
	"time"
)

func RootHandler() func (c *gin.Context) {
	return func (c *gin.Context) {
		ctx := context.Background()

		ctx, span := tracer.NewSpan(ctx, "GET /")
		defer span.End()

		doWork(ctx)
		doWork3(ctx)
		c.JSON(http.StatusOK, "Hello to order service!")
	}
}

func doWork(context context.Context) {
	ctx, span := tracer.NewSpan(context, "doWork")
	defer span.End()

	time.Sleep(time.Second * 2)

	doWork2(ctx)
}

func doWork2(context context.Context) {
	_, span := tracer.NewSpan(context, "doWork2")
	defer span.End()

	time.Sleep(time.Second * 1)
}

func doWork3(context context.Context) {
	_, span := tracer.NewSpan(context, "doWork3")
	defer span.End()

	time.Sleep(time.Second * 1)
}
