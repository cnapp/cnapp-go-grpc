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
	"os"

	"github.com/golang/glog"
)

const (
	// GrpcAddr is the environment variable that points to the gRPC server
	GrpcAddr = "CNAPPD_GRPC_SERVER"
)

var (
	ServerAddress string
)

func setupFromEnvironmentVariables() {
	ServerAddress = os.Getenv(GrpcAddr)
	glog.V(2).Infof("Env: %s", ServerAddress)
}
