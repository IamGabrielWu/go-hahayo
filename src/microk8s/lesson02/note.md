### create secret
```
CONTEXT=gke_gabrielhome_asia-east1-a_gke-learn
NAMESPACE=default
kubectl --namespace=$NAMESPACE --context=$CONTEXT \
create secret docker-registry gcr-pull-json-key \
--docker-server=https://us.gcr.io \
--docker-username=_json_key \
--docker-email=sre@aftership.com \
--docker-password="$(cat /Users/dzwu/gitspace/go-hahayo/src/microk8s/lesson02/gcr-pull.json)"

kubectl --namespace=$NAMESPACE --context=$CONTEXT \
patch serviceaccount default \
-p '{"imagePullSecrets": [{"name": "gcr-pull-json-key"}]}'
```
###
