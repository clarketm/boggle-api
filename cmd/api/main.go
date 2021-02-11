package main

import (
	"fmt"

	"github.com/clarketm/boggle-api/pkg/api"
	"github.com/clarketm/boggle-api/pkg/util"
)

func main() {
	cfg := util.NewConfig()
	addr := fmt.Sprintf(":%d", cfg.Addr)

	svc := api.NewService()
	if err := svc.Configure(); err != nil {
		svc.Log.Error.Fatal(err)
	}

	server := api.NewServer(addr, svc.Log.Error, svc)

	svc.Log.Error.Fatal(server.ListenAndServe())
}
