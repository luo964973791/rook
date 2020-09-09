### rook-ceph安装部署

```javascript
git clone https://github.com/luo964973791/rook.git
cd rook/cluster/examples/kubernetes/ceph
kubectl create -f common.yaml
kubectl create -f operator.yaml

#注意需要更改cluster.yaml里面的磁盘信息
kubectl create -f cluster.yaml
kubectl create -f toolbox.yaml
```

### cephfs安装

```javascript
cd rook/cluster/examples/kubernetes/ceph
kubectl create -f filesystem.yaml
cd rook/cluster/examples/kubernetes/ceph/csi/cephfs
kubectl create -f storageclass.yaml
kubectl create -f pvc.yaml
kubectl create -f pod.yaml
```

### 部署ceph客户端

```javascript
kubectl apply -f dashboard-external-https.yaml
#获取端口
kubectl get svc --all-namespaces
#账号.
admin
#登录密码
kubectl -n rook-ceph get secret rook-ceph-dashboard-password -o jsonpath="{['data']['password']}" | base64 --decode && echo
```

### 格式化ceph挂载盘
```javascript
yum install gdisk -y
sgdisk --zap-all /dev/sdb
```

