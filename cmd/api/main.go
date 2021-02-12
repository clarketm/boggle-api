package main

import (
	"fmt"

	"github.com/clarketm/boggle-api/pkg/api"
	"github.com/clarketm/boggle-api/pkg/util"
)

func main() {
	cfg := util.NewConfig()
	addr := fmt.Sprintf(":%d", cfg.HttpPort)

	logger := util.NewLogger()

	svc, err := api.NewService(cfg, logger)
	if err != nil {
		logger.Error.Fatal(err)
	}

	server := api.NewServer(addr, logger.Error, svc)

	logger.Error.Fatal(server.ListenAndServe())
}
