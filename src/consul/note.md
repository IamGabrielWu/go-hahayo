### dig
https://blog.csdn.net/adparking/article/details/5819725
https://www.cnblogs.com/sunshine-long/p/8862396.html

list service
```
curl http://localhost:8500/v1/catalog/service/web
```
list healthchecks
```
curl 'http://localhost:8500/v1/health/service/web?passing'
```
### socat
https://www.hi-linux.com/posts/61543.html

### create consul datacenter. two nodes, one is server and one is client
cd consul-getting-started-join
vagraunt up
vagrant ssh n1
### in node n1, start the consul server
```
consul agent \
 -server \
 -bootstrap-expect=1 \
 -data-dir=/tmp/consul \
 -node=agent-one \
 -bind=172.20.20.10 \
 -data-dir=/tmp/consul \
 -config-dir=/etc/consul.d
 ```
### in node n2, start the consul client
```
consul agent \
  -node=agent-two \
  -bind=172.20.20.11 \
  -enable-script-checks=true \
  -data-dir=/tmp/consul \
  -config-dir=/etc/consul.d
```
### in node n1, ask node n2 to join
```
consul join 172.20.20.11
```

### leave datacenter
```
consul leave
```
