package middlewares

import (
	"log"

	"github.com/familybook-project/familybook-api-gin/src/internal/db"
	"go.uber.org/zap"
)

func InitDB() {
	logger, zapErr := zap.NewProduction()
	if zapErr != nil {
		log.Fatal(zapErr.Error())
	}
	var err error
	db.DBInstance, err = db.Connect()
	if err != nil {
		logger.Error("Failed to connect to database",
			zap.String("error", err.Error()),
		)
		panic("Failed to connect to database")
	}
	logger.Info("Successfully connected to database")
}
