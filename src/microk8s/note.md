### in each node, run below
```
snap install microk8s --classic
sudo usermod -a -G microk8s $USER
```
### enable addons in control panel
```
microk8s.enable dns storage dashboard fluentd ingress istio jaeger metrics-server prometheus rbac registry
```
### each time when you need to add new node, run below to get token in master node
```
microk8s.add-node

Join node with: microk8s.join 192.168.64.2:25000/GmfMyZSTGCVBavCpnhqNrdXyeZWcKtNN

If the node you are adding is not reachable through the default interface you can use one of the following:
 microk8s.join 192.168.64.2:25000/GmfMyZSTGCVBavCpnhqNrdXyeZWcKtNN
 microk8s.join 10.1.60.0:25000/GmfMyZSTGCVBavCpnhqNrdXyeZWcKtNN
```
### create config file
```
microk8s.config view > ~/.kube/config
```
### run tkn
```
tkn taskrun describe echo-hello-world-task-run --kubeconfig=/home/ubuntu/.kube/config
```
### list images in cluster
```
microk8s.ctr images ls
```
### copy image to cluster from docker
```
docker save mynginx > myimage.tar
microk8s.ctr image import myimage.tar
```
### create docker registry secret and uses it in service account
```
kubectl apply secret docker-registry regcred --docker-username=dzwu --docker-password=dzwu --docker-email=dzwu@gmail.com --docker-server=https://index.docker.io/v1/
```
