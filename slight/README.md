最小化容器
=======


> 2019.05.17



## 使用方式


**建立容器：**

```
# Shell
docker build -t local/k8s-test:slight-sh-alpine --network host -f ./slight/shell/alpine/Dockerfile .
# Node
docker build -t local/k8s-test:slight-js-node-alpine           -f ./slight/node/node_alpine/Dockerfile .
# Golang
docker build -t local/k8s-test:slight-go-golang-alpine         -f ./slight/golang/golang_alpine/Dockerfile .
docker build -t local/k8s-test:slight-go-alpine                -f ./slight/golang/alpine/Dockerfile .
docker build -t local/k8s-test:slight-go-scratch               -f ./slight/golang/scratch/Dockerfile .
```


**查看容器大小：**

```
$ docker images
REPOSITORY          TAG                       IMAGE ID            SIZE
local/k8s-test      slight-sh-alpine          b886d30e78c9        9.57MB
local/k8s-test      slight-js-node-alpine     eb520d9c229e        77.7MB
local/k8s-test      slight-go-golang-alpine   63b734a0ef67        350MB
local/k8s-test      slight-go-alpine          0454ccc2e143        8.35MB
local/k8s-test      slight-go-scratch         7a00e86d70f1        2.81MB
alpine              latest                    cdf98d1859c1        5.53MB
archlinux/base      latest                    337c35b7ef88        445MB
node                latest                    80121c35659a        906MB
node                alpine                    f391dabf9dce        77.7MB
golang              alpine                    c7330979841b        350MB
golang              latest                    b860ab44e93e        774MB
```



## 參考


* [Kubernetes best practices: How and why to build small container images | Google Cloud Blog](https://cloud.google.com/blog/products/gcp/kubernetes-best-practices-how-and-why-to-build-small-container-images)
* [容器与云|创建尽可能小的 Docker 容器](https://linux.cn/article-5597-1.html)

