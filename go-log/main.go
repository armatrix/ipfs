package main

import logging "github.com/ipfs/go-log/v2"

func main() {
	var log = logging.Logger("main")

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.DPanic("dpanic")
	log.Panic("panic")
	log.Fatal("fatal")
}
