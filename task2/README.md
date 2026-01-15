# Go Hello World – Kubernetes Deployment (Task 2)

## Overview

This task deploys a simple Go **Hello World** HTTP service to Kubernetes using **raw manifests** and **Helm**, while intentionally identifying and fixing common issues across Go code, Docker, and Kubernetes.

The application exposes three endpoints:

* `/` → Hello World response
* `/health` → Health probe
* `/metrics` → Pod and runtime metadata

---

## What Was Fixed

### 1. Go Application Issues

*  Missing imports (`time`, `os`) causing build failure
*  HTTP handlers returned before writing responses
*  Incorrect `Content-Type` headers (`image/png` for JSON)
*  Go module missing (`go.mod`)

 **Fixes**:

* Added required imports
* Removed premature `return`
* Corrected `Content-Type` to `application/json`
* Initialized Go modules with `go mod init`

---

### 2. Docker Issues

*  `go build` failed due to missing `go.mod`
*  Binary name mismatch between build and runtime

 **Fixes**:

* Added Go modules
* Built correct binary and copied it into distroless image

---

### 3. Kubernetes (Raw Manifests)

*  Container port mismatch (8081 vs app listening on 8080)
*  Service selector mismatch with Deployment labels
*  Missing `imagePullSecrets` for private registry

 **Fixes**:

* Unified ports to `8080`
* Fixed labels/selectors
* Added `imagePullSecrets` for private image pull

---

### 4. Helm Chart Issues

*  Broken template syntax (`{{ .Release.Name -hello-world`)`
*  Wrong values key (`image.version` vs `image.tag`)

 **Fixes**:

* Corrected Helm templating syntax
* Aligned values with templates
* Added configurable `imagePullSecrets` via `values.yaml`


---

## How to Deploy

### Raw Manifests

```bash
kubectl apply -f k8s/
```

### Helm

```bash
helm upgrade --install hello ./helm/hello-world -n hello-app --create-namespace
```

---

## Validation Commands

```bash
kubectl get pods -n hello-app
kubectl exec -n hello-app deploy/hello-world -- curl localhost:8080/
kubectl exec -n hello-app deploy/hello-world -- curl localhost:8080/health
kubectl exec -n hello-app deploy/hello-world -- curl localhost:8090/metrics
```

---

## Improvements (Next Steps)

* Add Ingress for external access
* Add Prometheus annotations
* Add CI pipeline for image build & push
* Add HorizontalPodAutoscaler

---

## Result

 Application builds correctly
 Container runs in Kubernetes
 Health & metrics probes work
 Deployment works via **raw YAML** and **Helm**

This demonstrates understanding of **Go**, **Docker**, **Kubernetes**, and **Helm** with real-world debugging and fixes.

