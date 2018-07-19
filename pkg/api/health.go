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

package api

import (
	"fmt"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"

	"github.com/cnapp/cnapp-go-grpc/pb/health"
	"github.com/cnapp/cnapp-go-grpc/pkg/config"
	"github.com/cnapp/cnapp-go-grpc/pkg/rbac"
	"github.com/cnapp/cnapp-go-grpc/pkg/transport"
)

type HealthService struct {
	HealthUser string
	HealthKey  string
	URI        string
	Services   []string
}

func NewHealthService(conf *config.Configuration, uri string, services []string) (*HealthService, error) {
	glog.V(2).Info("Create the health service")
	rbac.AddRoles("health", "HealthService", map[string][]string{
		"Status": []string{rbac.AdminRole},
	})
	return &HealthService{
		// Conf:     conf,
		URI:      uri,
		Services: services,
	}, nil
}

func (service *HealthService) Status(ctx context.Context, req *health.StatusRequest) (*health.StatusResponse, error) {
	glog.V(1).Info("Check Health services")

	conn, err := grpc.Dial(service.URI, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := healthpb.NewHealthClient(conn)
	md := metadata.New(map[string]string{
		transport.UserID: service.HealthUser,
	})
	newCtx := metadata.NewIncomingContext(ctx, md)

	servicesStatus := []*health.ServiceStatus{}
	for _, service := range service.Services {
		glog.V(2).Infof("Check health service: %s", service)
		resp, err := client.Check(newCtx, &healthpb.HealthCheckRequest{
			Service: service,
		})
		if err != nil {
			servicesStatus = append(servicesStatus, &health.ServiceStatus{
				Name:   service,
				Status: "KO",
				Text:   err.Error(),
			})
		} else {
			servicesStatus = append(servicesStatus, &health.ServiceStatus{
				Name:   service,
				Status: "OK",
				Text:   fmt.Sprintf("%s", resp.Status),
			})
		}
	}

	resp := &health.StatusResponse{}
	resp.Services = servicesStatus

	glog.V(0).Infof("Health response: %s", resp)
	return resp, nil
}
