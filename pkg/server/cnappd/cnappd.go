// Copyright (C) 2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cnappd

import (
	"fmt"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"

	"github.com/cnapp/cnapp-go-grpc/pkg/api"
	"github.com/cnapp/cnapp-go-grpc/pkg/config"
)

const (
	apiVersion = "v1"
)

func StartServer(configFilename string) {
	conf, err := config.LoadFileConfig(configFilename)
	if err != nil {
		glog.Fatalf("failed to load configuration: %v", err)
	}

	glog.V(0).Infoln("Create the gRPC servers")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAddr := fmt.Sprintf(":%d", conf.API.GrpcPort)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	glog.V(0).Infof("Listen on %s", grpcAddr)

	grpcServer, err := registerServer(conf, grpcAddr)
	if err != nil {
		glog.Fatalf("Failed to register gRPC server: %s", err.Error())
	}

	gwmux, err := registerGateway(ctx, fmt.Sprintf("localhost:%d", conf.API.GrpcPort))
	if err != nil {
		glog.Fatalf("Failed to register JSON gateway: %s", err.Error())
	}

	httpmux := http.NewServeMux()
	httpmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>cnapp</title></head>
             <body>
             <h1>cnapp</h1>
             </body>
             </html>`))
	})
	httpmux.Handle(fmt.Sprintf("/%s/", apiVersion), gwmux)
	httpmux.Handle("/metrics", prometheus.Handler())
	httpmux.HandleFunc("/version", api.VersionHandler)
	api.ServeStaticFile(httpmux)
	api.ServeSwagger(httpmux)

	glog.V(0).Infof("Start gRPC server on %s", grpcAddr)
	go grpcServer.Serve(lis)

	gwAddr := fmt.Sprintf(":%d", conf.API.RestPort)
	srv := &http.Server{
		Addr:    gwAddr,
		Handler: grpcHandlerFunc(grpcServer, httpmux),
	}
	glog.V(0).Infof("Start HTTP server on %s", gwAddr)
	glog.Fatal(srv.ListenAndServe())

}
