package main

import (
	"boiler/pkg/conf"
	"boiler/pkg/factory"
	"boiler/pkg/logging"
	"context"
	"flag"
	"log"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yml", "config file path")
	flag.Parse()

	config, err := conf.NewFileReader(configPath).ReadConf()
	if err != nil {
		log.Fatal(err)
	}

	var logger logging.Logger = logging.Logrus()

	controller, err := factory.CreateController(config, logger)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = controller.Execute(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			logger.Warn("budget timeout expired")
			return
		}

		log.Fatal(err)
	}

}
