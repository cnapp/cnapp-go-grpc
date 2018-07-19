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

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	// "github.com/cnapp/cnapp-go-grpc/pkg/config"
	"github.com/cnapp/cnapp-go-grpc/pkg/transport"
)

var (
	ErrGrpcAddressNotFound = errors.New(fmt.Sprintf(
		"gRPC address not found. Setup environment variable %s", GrpcAddr))
)

// GRPCClient define a client using gRPC protocol
type GRPCClient struct {
	ServerAddress string
}

// NewGRPCClient creates a new gRPC client
func NewGRPCClient(cmd *cobra.Command) (*GRPCClient, error) {
	setupFromEnvironmentVariables()
	if len(ServerAddress) == 0 {
		return nil, ErrGrpcAddressNotFound
	}
	// conf := &config.Configuration{}
	glog.V(2).Infof("gRPC client created: %s %s", ServerAddress)
	return &GRPCClient{
		ServerAddress: ServerAddress,
	}, nil
}

func (client *GRPCClient) GetConn() (*grpc.ClientConn, error) {
	return grpc.Dial(
		client.ServerAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
}

func (client *GRPCClient) GetContext(cliName string) (context.Context, error) {
	ctx := context.Background()
	headers := map[string]string{}
	if host, err := os.Hostname(); err != nil {
		headers[transport.UserHostname] = host
	}
	md := metadata.New(headers)
	glog.V(2).Infof("Transport metadata: %s", md)
	return metadata.NewIncomingContext(ctx, md), nil
}
