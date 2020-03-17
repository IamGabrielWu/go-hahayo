### Go client
```
https://github.com/kubernetes/client-go
```

### accessing kubernetes api server
API way
* kubectl get --raw /api/v1/namespaces/kube-system

kubectl way
* kubectl get ns


### kubectl proxy: starts a proxy between localhost and api server
```
kubectl proxy --port 8080
curl -s http://localhost:8080/api/v1/namespaces/kube-system
```

### accessing kubernetes api server with service account
For each pod, the following information and credentials related to service accounts are mounted by default
* service account and token: /var/run/secrets/kubernetes.io/serviceaccount/token
* certificate bundle: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
* namespace: /var/run/secrets/kubernetes.io/serviceaccount/namespace

任何pod 里都有这些
/ # ls /var/run/secrets/kubernetes.io/serviceaccount
ca.crt     namespace  token
namespace: 当前的namespace
token: 和 api server 交互的bear authentication token
ca.crt:  和 apiserver 交互的证书

任何pod 里 默认都有以下环境变量
/ # env | grep KUBE
KUBERNETES_PORT=tcp://10.15.240.1:443
KUBERNETES_SERVICE_PORT=443
KUBERNETES_PORT_443_TCP_ADDR=10.15.240.1
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_PROTO=tcp
KUBERNETES_PORT_443_TCP=tcp://10.15.240.1:443
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_SERVICE_HOST=10.15.240.1

以下解释了pod 内如何调用api server 与此进行交互的
CACERT=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
APISERVER="https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}"
curl --header "Authorization: Bearer $TOKEN" --cacert $CACERT $APISERVER/api/v1/namespaces/kube-system
