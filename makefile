# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

#RSA keys. To generate a private/public key PEM file.
# openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# openssl rsa -pubout -in private.pem -out public.pem


run:
	go run api/cmd/services/sales/main.go | go run api/cmd/tooling/logfmt/main.go
# 	go run api/cmd/services/auth/main.go | go run api/cmd/tooling/logfmt/main.go

run-auth:
	go run api/cmd/services/auth/main.go | go run api/cmd/tooling/logfmt/main.go

admin: 
	go run apis/tooling/admin/main.go

help:
	go run api/cmd/services/sales/main.go --help

version:
	go run api/cmd/services/sales/main.go --version

tidy:
	go mod tidy
	go mod vendor

curl-hello:
	curl -i -X GET http://127.0.0.1:3000/hello

curl-liveness: 
	curl -i -X GET http://localhost:3000/liveness


curl-readiness: 
	curl -i -X GET http://127.0.0.1:3000/readiness


curl-testerr: 
	curl -i -X GET http://127.0.0.1:3000/testerr

curl-panic: 
	curl -i -X GET http://127.0.0.1:3000/testpanic


# ==============================================================================
# Define dependencies

GOLANG          := golang:1.22
ALPINE          := alpine:3.19
KIND            := kindest/node:v1.29.2
POSTGRES        := postgres:16.2
GRAFANA         := grafana/grafana:10.4.0
PROMETHEUS      := prom/prometheus:v2.51.0
TEMPO           := grafana/tempo:2.4.0
LOKI            := grafana/loki:2.9.0
PROMTAIL        := grafana/promtail:2.9.0

KIND_CLUSTER    := lab-starter-cluster
NAMESPACE       := sales-system
SALES_APP       := sales
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/lab
VERSION         := 0.0.1
SALES_IMAGE     := $(BASE_IMAGE_NAME)/$(SALES_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)


# ==============================================================================
# Building containers

build: sales 

sales: 
	docker build \
	-f zarf/docker/dockerfile.sales \
	-t $(SALES_IMAGE) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
	.


#=====================================================================
#Running from within k8s/kind

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl cluster-info --context kind-$(KIND_CLUSTER)

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

# 	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	kubectl get pods -o wide -A --watch

# ---------------------------------------------------------------------
dev-load:
	kind load docker-image $(SALES_IMAGE) --name $(KIND_CLUSTER)
# 	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-load-db:
	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/sales | kubectl apply -f -

dev-restart:
	kubectl rollout restart deployment $(SALES_APP) --namespace=$(NAMESPACE)

dev-update: build dev-load dev-restart

dev-update-apply: build dev-load dev-apply
	

dev-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) --all-containers=true -f --tail=1000 --max-log-requests=6 | go run api/cmd/tooling/logfmt/main.go

# ---------------------------------------------------------------------
dev-describe-deployment:
	kubectl describe deployment $(SALES_APP) --namespace=$(NAMESPACE)

dev-describe-sales:
	kubectl describe pod $(SALES_APP) --namespace=$(NAMESPACE)

dev-start-db:
	docker run -d -p 5432:5432 --name lab_db -e POSTGRES_PASSWORD=lab_db -e POSTGRES_DB=lab_db postgres

  

# ==============================================================================
# Metrics and Tracing

metrics: 
	expvarmon -ports="localhost:3010" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"

statsviz:
	open -a "Google Chrome" http://localhost:3010/debug/statsviz