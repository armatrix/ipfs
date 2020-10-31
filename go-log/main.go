package main

import (
	"github.com/armatrix/ipfs/go-log/level"
	logging "github.com/ipfs/go-log/v2"
)

func main() {
	var log = logging.Logger("main")
	level.Lv()
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.DPanic("dpanic")
	log.Panic("panic")
	log.Fatal("fatal")
}
