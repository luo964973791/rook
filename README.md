### rook-ceph安装部署

```javascript
git clone https://github.com/luo964973791/rook.git
cd rook/cluster/examples/kubernetes/ceph
kubectl create -f crds.yaml -f common.yaml -f operator.yaml

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
#挂载,cp容器里ceph.repo到主机上.
yum install ceph-fuse -y
chmod 755 /etc/rc.d/rc.local
echo "ceph-fuse -m 10.233.39.242:6789,10.233.47.20:6789,10.233.2.92:6789 /data -n client.admin --keyring=/etc/ceph/keyring" >> /etc/rc.local
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
