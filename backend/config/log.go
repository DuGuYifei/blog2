package config

import (
	"fmt"
	log "github.com/cihub/seelog"
)

func InitLog() {
	logger, err := log.LoggerFromConfigAsFile("log.xml")
	if err != nil {
		panic(fmt.Sprintf("error parsing log config file: %v", err))
	}
	err = log.ReplaceLogger(logger)
	if err != nil {
		panic(fmt.Sprintf("error replacing logger: %v", err))
	}
	defer log.Flush()
	err = log.Critical("Log 启动成功")
	if err != nil {
		panic(fmt.Sprintf("error writing critical log message: %v", err))
	}
}
