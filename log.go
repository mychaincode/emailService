package service

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger ...
var Logger *zap.Logger

// InitLogger ...
func InitLogger() {
	lp := Conf.Common.LogPath
	lv := Conf.Common.LogLevel
	isDebug := true
	if Conf.Common.IsDebug != true {
		isDebug = false
	}
	initLogger(lp, lv, isDebug)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			checkLog(lp, lv, isDebug)
		}
	}()
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func checkLog(lp string, lv string, isDebug bool) {
	if !fileExist(lp) {
		initLogger(lp, lv, isDebug)
		log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	}
}

func initLogger(lp string, lv string, isDebug bool) {
	js := fmt.Sprintf(`{
		"level": "%s",
		"encoding": "json",
		"outputPaths": ["stdout","%s"],
		"errorOutputPaths": ["stderr","%s"]
	}`, lv, lp, lp)

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoderScc
	// cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}
