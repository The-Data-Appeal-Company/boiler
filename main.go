package main

import (
	"boiler/pkg/conf"
	"boiler/pkg/factory"
	"context"
	"flag"
	"log"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "config.yml", "config file path")
	flag.Parse()

	config, err := conf.NewFileReader(configPath).ReadConf()
	if err != nil {
		log.Fatal(err)
	}

	controller, err := factory.CreateController(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = controller.Execute(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
