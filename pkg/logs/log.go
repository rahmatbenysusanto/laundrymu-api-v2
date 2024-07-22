package logs

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"time"
)

func InitLog() {
	logFile, err := os.OpenFile("./app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Info(err)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			return
		}
	}(logFile)

	log.Info("TES")

	log.SetOutput(logFile)
}

func Sentry(ctx *fiber.Ctx) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://54e5615db7191c187c1d6d1d63018f78@o4507555736977408.ingest.de.sentry.io/4507555740319824",
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	defer sentry.Flush(2 * time.Second)

	go func() {
		transaction := sentry.TransactionFromContext(ctx.Context())
		if transaction != nil {
			transaction.Name = "Pelanggan"
		}

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelError)
			scope.SetContext("error", map[string]interface{}{
				"message": "Sentry error",
				"data":    transaction,
			})
		})

		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Level:    sentry.LevelError,
			Message:  "Sentry error",
			Data:     nil,
			Category: "pelanggan",
		})
		sentry.CaptureMessage("Error Get Pelanggan")
	}()
}
