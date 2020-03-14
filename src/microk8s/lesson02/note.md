### create secret
```
CONTEXT=gke_gabrielhome_asia-east1-a_gke-learn
NAMESPACE=default
kubectl --namespace=$NAMESPACE --context=$CONTEXT \
create secret docker-registry gcr-pull-json-key \
--docker-server=https://gcr.io \
--docker-username=_json_key \
--docker-email=iamgabrielwu@gmail.com \
--docker-password="$(cat /Users/dzwu/gitspace/go-hahayo/src/microk8s/secrets/gcr-pull.json)"
```
###
