# CNAPPS / Golang gRPC-Gateway

Features:

* [x] Logging
* [ ] Metrics
* [ ] Tracing

## Local

Execute the server in local:

    $ ./cnappd run --config cnappd.toml --alsologtostderr -v 9

Check info:

    $ CNAPPD_GRPC_SERVER=127.0.0.1:8210 ./cnappadm info --alsologtostderr -v 9
    I0719 09:40:49.954184   27408 env.go:34] Env: 127.0.0.1:8210
    I0719 09:40:49.954452   27408 grpc.go:50] gRPC client created: 127.0.0.1:8210 %!s(MISSING)
    I0719 09:40:49.954475   27408 info.go:57] Retrieve informations
    I0719 09:40:49.954571   27408 grpc.go:72] Transport metadata: map[]
    +---------+----------------+---------+--------+
    | SERVICE |      URI       | VERSION | STATUS |
    +---------+----------------+---------+--------+
    | cnappd  | 127.0.0.1:8210 | 0.1.0   | OK     |
    +---------+----------------+---------+--------+

Check info:

    $ CNAPPD_GRPC_SERVER=127.0.0.1:8210 ./cnappadm health --alsologtostderr -v 9
    I0719 09:42:19.185068   27762 env.go:34] Env: 127.0.0.1:8210
    I0719 09:42:19.185206   27762 grpc.go:50] gRPC client created: 127.0.0.1:8210 %!s(MISSING)
    I0719 09:42:19.185221   27762 health.go:53] Check health
    I0719 09:42:19.185281   27762 grpc.go:72] Transport metadata: map[]
    +---------+--------+------+
    | SERVICE | STATUS | TEXT |
    +---------+--------+------+


## Local with Docker

Build the Docker image:

    $ make minikube-build

Run a container:

    $ make docker-run

## Minikube

Build the Docker image into minikube:

    $ make minikube-build

Deploy the application into minikube:

    $ make minikube-deploy

Add to your `/etc/hosts` the URI :

    $ echo $(KUBECONFIG=./deploy/minikube-kube-config minikube ip) gogrpcgw.cnapps.minikube | sudo tee -a /etc/hosts

Then check the service on URL: http://gogrpcgw.cnapps.minikube/

You could see the API documentation on: http://grgpcgw.cnapps.minikube/swagger-ui/

Undeploy the application from minikube:

    $ make minikube-undeploy