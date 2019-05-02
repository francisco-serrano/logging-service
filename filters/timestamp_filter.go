package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func InsertMetadata(ctx *context.Context) {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	ctx.Request.Header.Add("correlation-id", uuid.New().String())
	ctx.Request.Header.Add("initial-timestamp", timestamp)
}

func LogRequestMetadata(ctx *context.Context) {
	correlationId := ctx.Request.Header.Get("correlation-id")
	initialTimestamp := ctx.Request.Header.Get("initial-timestamp")
	currentTimestamp, err := strconv.ParseInt(initialTimestamp, 10, 64)

	if err != nil {
		panic(err)
	}

	difference := time.Now().UnixNano() - currentTimestamp

	log.WithFields(log.Fields{
		"method":            ctx.Request.Method,
		"path":              ctx.Request.URL.Path,
		"statusCode":        ctx.ResponseWriter.Status,
		"correlation-id":    correlationId,
		"initial-timestamp": initialTimestamp,
		"current-timestamp": currentTimestamp,
		"difference":        difference,
	}).Info("REQUEST_INFO")
}
