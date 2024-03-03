package dependency

import "go.uber.org/zap"

func ErrorLog(){
	// intance logger
	var logger = *zap.NewExample()

	// use logger
	logger.Error("example error", zap.Int("example", 1))
}