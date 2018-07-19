# Copyright (C) 2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APP="cnapp-go-grpc"

VERSION=$(shell \
	grep "const Version" pkg/version/version.go \
	|awk -F'=' '{print $$2}' \
	|sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

MAKE_COLOR=\033[33;01m%-20s\033[0m

DB_ENGINE = "cockroachdb://cnapps@192.168.99.100:32007/cnapps?sslmode=disable"

SHELL = /bin/bash
DOCKER = docker

GO = go
GOX = gox -os="linux darwin windows freebsd openbsd netbsd"
GOX_ARGS = "-output={{.Dir}}-$(VERSION)_{{.OS}}_{{.Arch}}"

IMAGE=$(APP)

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) : %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Cleanup
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup environnement$(NO_COLOR)"
	@rm ./cnappadm ./cnappctl ./cnappd

.PHONY: tools
tools:
	@echo -e "$(OK_COLOR)[$(APP)] Initialize environnement$(NO_COLOR)"
	@go get -u github.com/golang/glog
	@go get -u github.com/kardianos/govendor
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/kisielk/errcheck
	@go get -u github.com/mitchellh/gox
	@wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip

.PHONY: proto
proto: ## Install protocol buffer tools
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u github.com/golang/protobuf/proto
	@go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

init: tools proto ## Install requirements

.PHONY: deps
deps: ## Install dependencies
	@echo -e "$(OK_COLOR)[$(APP)] Update dependencies$(NO_COLOR)"
	@govendor update

.PHONY: pb
pb: ## Generate Protobuf
	@go generate pb/api.go

.PHONY: swagger
swagger: ## Generate Swagger
	go-bindata-assetfs -pkg swagger third_party/swagger-ui/... && mv bindata_assetfs.go pkg/ui/swagger/

.PHONY: changelog
changelog:
	@$(GO) generate -x ./pkg/static/


.PHONY: webdoc
webdoc: doc
	@$(GO) generate -x ./pkg/webdoc/

.PHONY: build
build: ## Make binary
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@$(GO) build -o cnappd github.com/cnapp/cnapp-go-grpc/cmd/cnappd
	@$(GO) build -o cnappctl github.com/cnapp/cnapp-go-grpc/cmd/cnappctl
	@$(GO) build -o cnappadm github.com/cnapp/cnapp-go-grpc/cmd/cnappadm

.PHONY: test
test: ## Launch unit tests
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@govendor test +local

.PHONY: lint
lint: ## Launch golint
	@$(foreach file,$(SRCS),golint $(file) || exit;)

.PHONY: vet
vet: ## Launch go vet
	@$(foreach file,$(SRCS),$(GO) vet $(file) || exit;)

.PHONY: errcheck
errcheck: ## Launch go errcheck
	@echo -e "$(OK_COLOR)[$(APP)] Go Errcheck $(NO_COLOR)"
	@$(foreach pkg,$(PKGS),errcheck $(pkg) $(glide novendor) || exit;)

.PHONY: coverage
coverage: ## Launch code coverage
	@$(foreach pkg,$(PKGS),$(GO) test -cover $(pkg) $(glide novendor) || exit;)

gox: ## Make all binaries
	@echo -e "$(OK_COLOR)[$(APP)] Create binaries $(NO_COLOR)"
	$(GOX) -output=cnappctl-$(VERSION)_{{.OS}}_{{.Arch}} -osarch="linux/amd64 darwin/amd64 windows/amd64" github.com/cnapp/cnapp-go-grpc/cmd/cnappctl
	$(GOX) -output=cnappadm-$(VERSION)_{{.OS}}_{{.Arch}} -osarch="linux/amd64 darwin/amd64 windows/amd64" github.com/cnapp/cnapp-go-grpc/cmd/cnappadm
	$(GOX) -output=cnappd-$(VERSION)_{{.OS}}_{{.Arch}} -osarch="linux/amd64 darwin/amd64 windows/amd64" github.com/cnapp/cnapp-go-grpc/cmd/cnappd

#
# Docker
#

.PHONY: docker-build
docker-build: ## Build Docker image for application
	@echo -e "$(OK_COLOR)[$(APP)] Build Docker Image$(NO_COLOR)"
	$(DOCKER) build \
		--build-arg http_proxy=$$http_proxy \
		--build-arg https_proxy=$$https_proxy \
		-t $(NAMESPACE)/$(IMAGE):$(VERSION) .

.PHONY: docker-debug
docker-debug: ## Run a shell into the Docker image
	@echo -e "$(OK_COLOR)[$(APP)] Build Docker Image$(NO_COLOR)"
	@$(DOCKER) run --rm \
		-it $(NAMESPACE)/$(IMAGE):$(VERSION) /bin/bash

.PHONY: docker-run
docker-run: ## Run application using Docker image with ip=x.x.x.x
	@echo -e "$(OK_COLOR)[$(APP)] Run Docker Container$(NO_COLOR)"
	@$(DOCKER) run --rm -p 9191:9191 \
		--name $(APP) $(NAMESPACE)/$(IMAGE):$(VERSION)

#
# Kubernetes
#

.PHONY: minikube-build
minikube-build: ## Build Docker image into Minikube
	@echo -e "$(OK_COLOR)[$(APP)] Deploy application to local Kubernetes$(NO_COLOR)"
	@eval $$(KUBECONFIG=../deploy/minikube/kube-config minikube docker-env -p cnapps); \
		$(DOCKER) build \
			--build-arg http_proxy=$$http_proxy \
			--build-arg https_proxy=$$https_proxy \
			-t $(IMAGE):$(VERSION) .

.PHONY: minikube-deploy
minikube-deploy: minikube-build ## Deploy application into Minikube
	@echo -e "$(OK_COLOR)[$(APP)] Deploy application to local Kubernetes$(NO_COLOR)"
	@./scripts/kubernetes.sh -c minikube -e local -p create -a $(APP) -d deploy/k8s/ -t $(VERSION)

.PHONY: minikube-undeploy
minikube-undeploy: ## Undeploy application into Minikube
	@echo -e "$(OK_COLOR)[$(APP)] Deploy application to local Kubernetes$(NO_COLOR)"
	@./scripts/kubernetes.sh -c minikube -e local -p destroy -a $(APP) -d deploy/k8s/ -t $(VERSION)