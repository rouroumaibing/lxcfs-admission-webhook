lxcfs增强容器隔离top查看容器资源

原代码库地址：
https://github.com/ymping/lxcfs-admission-webhook

改造点： 

1) 证书验证使用cert-manager

2) 节点使用ubuntu 18.04

3) 根据实际情况 修改目录和挂载文件

1、部署前检查API接口是否支持

kubectl api-versions | grep admissionregistration.k8s.io/v1

2、 每个节点安装lxcfs,查看lxcfs安装的库的都有哪些，根据实际情况修改代码库的挂载文件及对应的目录,修改代码/cmd/volume.go

apt install -y lxcfs

#也可以通过源码安装https://github.com/lxc/lxcfs

2.1 当前测试的集群节点有以下文件，需要删除一些代码中多挂载的存储：

/proc/ 

cpuinfo  diskstats  meminfo  stat  swaps  uptime

2.2 节点对应的lxcfs目录为/var/lib/lxcfs

根据实际情况修改cmd/volume.go

3、 构建镜像

docker build -t lxcfs-admission-webhook:v1.0 . --network=host


4、安装cert-manager

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.3/cert-manager.yaml

5、创建命名空间

kubectl create ns lxcfs

6、生成证书

cd deploy/certs/ ; kubectl apply -f ./ ; cd -


7、部署lxcfs-admission-webhook

修改./deploy/webhook/deployment.yaml镜像地址

cd deploy/webhook/ ; kubectl apply -f ./ ; cd -


8、给命名空间test打标签

kubectl create ns test

kubectl label namespaces  test lxcfs-admission-webhook=enabled

9、在test命名空间创建工作负载，查看pod yaml注入情况

kubectl -n test  create deploy nginx --image=nginx

kubectl -n test describe pod nginx-85b98978db-7286b

    Mounts:
    
      /proc/cpuinfo from lxcfs (ro,path="lxcfs/proc/cpuinfo")
      
      /proc/diskstats from lxcfs (ro,path="lxcfs/proc/diskstats")
      
      /proc/meminfo from lxcfs (ro,path="lxcfs/proc/meminfo")
      
      /proc/stat from lxcfs (ro,path="lxcfs/proc/stat")
      
      /proc/swaps from lxcfs (ro,path="lxcfs/proc/swaps")
      
      /proc/uptime from lxcfs (ro,path="lxcfs/proc/uptime")

10、卸载

 cd deploy/certs/ ; kubectl delete -f ./ ; cd -

cd deploy/webhook/ ; kubectl delete -f ./ ; cd -

kubectl delete ns lxcfs test
