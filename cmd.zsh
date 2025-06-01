
# 部署数据库 MySQL 
docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql



# kind is a tool for running local Kubernetes clusters using Docker container “nodes”.
brew install kind
brew install kubectl 

brew install podman
podman machine init
podman machine start

# 部署集群
kind create cluster --config=./deploy/gomall-dev-cluster.yaml

docker pull nginx 
kind load docker-image --name gomall-dev nginx:latest
kubectl run nginx --image=nginx --port=80 --image-pull-policy=IfNotPresent
kubectl get pod  


kubectl apply -f ./deploy/gomall-dev-base.yaml
## 如果代理失败，尝试下面本地加入
kind load docker-image --name gomall-dev mysql:latest
kind load docker-image --name gomall-dev redis:latest
kind load docker-image --name gomall-dev nats:latest
kind load docker-image --name gomall-dev jaegertracing/all-in-one:latest
kind load docker-image --name gomall-dev prom/prometheus:latest

kind load docker-image --name gomall-dev hashicorp/consul:latest
kind load docker-image --name gomall-dev docker.io/bitnami/etcd:latest
kind load docker-image --name gomall-dev grafana/grafana:latest
kind load docker-image --name gomall-dev grafana/loki:latest
kind load docker-image --name gomall-dev grafana/promtail:latest



kubectl apply -f ./deploy/gomall-dev-app.yaml

kind load docker-image --name gomall-dev email:v1.1.1 frontend:v1.1.1 cart:v1.1.1 checkout:v1.1.1 order:v1.1.1 payment:v1.1.1 product:v1.1.1 user:v1.1.1 

kubectl apply -f ./deploy/gomall-dev-base.yaml

# 暴露服务：kubectl 端口转发
kubectl port-forward pod/frontend-747669c8bd-mdz27 8080:8080




http://localhost:8080/

service discovery error: Get "http://consul-svc:8500/v1/health/service/product?passing=1": dial tcp 10.96.97.120:8500: connect: connection refused

➜  ~ kubectl get pod
NAME                          READY   STATUS             RESTARTS         AGE
cart-7c8cb65d9b-45zf5         0/1     CrashLoopBackOff   11 (3m28s ago)   38m
checkout-757f6c5789-pgnxt     0/1     CrashLoopBackOff   11 (55s ago)     38m
consul-85d6cc886f-pmb5t       0/1     ErrImagePull       0                45m
email-99f58b4b5-b7rqc         0/1     CrashLoopBackOff   9 (2m12s ago)    38m
frontend-747669c8bd-mdz27     1/1     Running            0                38m
jaeger-56dcb4f99-25rkp        1/1     Running            0                45m
mysql-7758879967-rh7kp        1/1     Running            0                45m
nats-758c499dd6-zhqtj         1/1     Running            0                45m
nginx                         1/1     Running            0                45m
order-59dcb7cccd-dhn72        0/1     CrashLoopBackOff   11 (3m10s ago)   38m
payment-5fb6c457c6-9dqj5      0/1     CrashLoopBackOff   11 (2m25s ago)   38m
product-d8694bdfc-bp7tx       0/1     CrashLoopBackOff   11 (86s ago)     38m
prometheus-7cbdf8f994-ccnz5   1/1     Running            0                45m
redis-7c7dfbc86b-89fhn        1/1     Running            0                45m
user-84677d649f-m4rr8         0/1     CrashLoopBackOff   11 (2m43s ago)   38m


2025/05/28 10:18:51 /usr/src/gomall/app/product/biz/dal/mysql/init.go:22

[error] failed to initialize database, got error dial tcp :3306: connect: connection refused

panic: dial tcp :3306: connect: connection refused



