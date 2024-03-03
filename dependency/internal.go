package dependency

import "go.uber.org/zap"

var logger zap.Logger

func init(){
	logger = *zap.NewExample()
}

func Log(){
	//use logger
}