package main

import (
	"github.com/Kaebraza/gopportunities/config"
	"github.com/Kaebraza/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// iniciallize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config inicialization error: %v", err)
		return
	}
	// iniciallize router
	router.Initialize()
}
