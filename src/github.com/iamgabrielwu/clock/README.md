### 参考
https://www.jianshu.com/p/0c79d08d63a4
### 目标
```
钟表的服务，当每一分钟，更新自己pod name当前时间的字符， clock
```
### 设置环境
cd /Users/dzwu/gitspace/go-hahayo
export GOPATH=$PWD
mkdir -p $GOPATH/src/github.com/iamgabrielwu
cd $GOPATH/src/github.com/iamgabrielwu
### 启动开发目录
```
operator-sdk new clock
cd clock
operator-sdk add api --api-version=clock.iamthat.com/v1 --kind=Clock
operator-sdk add controller --api-version=clock.iamthat.com/v1 --kind=Clock
```
### 修改代码
修改 /Users/dzwu/gitspace/go-hahayo/src/github.com/iamgabrielwu/clock/pkg/apis/clock/v1/clock_types.go
1. ClockSpec
2. ClockStatus
至少
### 生成代码
everytime that file is modified, run ```operator-sdk generate k8s```
### 编写控制器业务逻辑
add business logic to /Users/dzwu/gitspace/go-hahayo/src/github.com/iamgabrielwu/clock/pkg/controller/clock/clock_controller.go
modify method "Reconcile"

### 本地调试， 可以在控制台上看到调试日志
operator-sdk run --local

### 如果没有问题，打包成镜像
```
operator-sdk build gcr.io/gabrielhome/crd/clock:latest

```
### 上传镜像， 修改image
```
gcloud docker -- push build gcr.io/gabrielhome/crd/clock:latest
修改文件 /Users/dzwu/gitspace/go-hahayo/src/github.com/iamgabrielwu/clock/deploy/operator.yaml
```
### 在集群上创建 crd，部署operator
```
kubectl create -f deploy/crds/clock.iamthat.com_clocks_crd.yaml
kubectl create -f deploy/operator.yaml

```
### 发布资源清单来发布应用
```
kubectl apply -f deploy/crds/clock.iamthat.com_v1_clock_cr.yaml
```
