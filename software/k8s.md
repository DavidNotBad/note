## 集群架构与组件

### 单master架构图

![image-20190619194635594](../image/k8s/image-20190619194635594.png)![image-20190619194635594](../image/k8s/WX20190619-200224.png)

### 用户访问流程图

![WX20190619-195751](../image/k8s/WX20190619-195751.png)

### 模型对象

![WX20190619-195751](../image/k8s/WX20190619-200848.png)
![WX20190619-201213](../image/k8s/WX20190619-201213.png)
![WX20190619-201449](../image/k8s/WX20190619-201449.png)

## 高可用集群部署

### 部署方式与集群规划

* minikube

  minikube是一个工具, 可以在本地快速运行一个单点的k8s, 仅用于尝试或日常开发的用户使用

  部署地址: https://kubernetes.io/docs/setup/minikube

* kubeadm

  kubeadm也是一个工具, 提供kubeadm和kubeadm join, 用于快速部署k8s集群

  部署地址: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm

* 二进制包

  推荐, 从官方下载你发行版的二进制包, 手动部署每个组件, 组成k8s集群

  下载地址: https://github.com/kubernetes/kubernetes/releases

### 平台规划

![WX20190619-201449](../image/k8s/20190620095503.png)

### 单节点集群架构图

![WX20190619-201449](../image/k8s/20190620095757.png)

### 多节点集群架构图

![WX20190619-201449](../image/k8s/20190620095915.png)

### 环境搭建并初始化(单节点)

1. 使用vagrant创建3台机器

   boxes和bridge自行修改, bridge不清楚可以删掉
   
   vagrantfile内容见附录: vagrantfile

2. 初始化系统

   ```shell
   # 以下操作均在master机器上
   
   # 关闭selinux
   sudo vi /etc/selinux/config
   SELINUX=disabled
   sudo setenforce 0
   # 关闭防火墙(可选)
   sudo systemctl stop firewalld
   # 修改主机名(重新进入终端)
   sudo vi /etc/hostname
   sudo hostname k8s-master01
   # 同步时间
   sudo yum -y install ntpdate && sudo ntpdate time.windows.com
   ```

3. 生成https证书

   ![image-20190619194635594](../image/k8s/20190620110036.png)

4. 为etcd自签证书

```shell
# 创建目录
cd 
mkdir k8s && cd k8s
mkdir k8s-cert
mkdir etcd-cert && cd etcd-cert

# 搭建证书生成工具
vi cfssl.sh
# 内容见附录: cfssl.sh
sudo bash cfssl.sh

# 生成证书
vi etcd-cert.sh
# 内容见附录: etcd-cert.sh
sudo bash etcd-cert.sh
```

5. 安装etcd

```shell
# 下载etcd
# 其它版本: https://github.com/etcd-io/etcd/releases
curl -L -O https://github.com/etcd-io/etcd/releases/download/v3.3.13/etcd-v3.3.13-linux-amd64.tar.gz
# 解压
tar -zxvf etcd-v3.3.13-linux-amd64.tar.gz
# 移动到指定目录
sudo mkdir -p /opt/etcd/{cfg,bin,ssl}
cd etcd-v3.3.13-linux-amd64
sudo mv etcd etcdctl /opt/etcd/bin/
# 部署启动脚本
cd ~/k8s/
vi etcd.sh
## 内容见附录: etcd.sh
chmod +x etcd.sh

# 拷贝证书
sudo cp ~/k8s/etcd-cert/{ca,server-key,server}.pem /opt/etcd/ssl/

# 执行部署脚本
sudo ./etcd.sh etcd01 192.168.1.30 etcd02=https://192.168.1.31:2380,etcd03=https://192.168.1.32:2380

# 查看日志
## 发现正在等待连接其它两个节点, 而此时其它节点我们还没有部署
sudo tail /var/log/messages -f

# 部署另外连个节点(在node1和node2上执行)
## 因为虚拟机是vagrant创建的, 外部是无法通过密码访问ssh的, 需要修改配置
## vagrant用户的默认密码是vagrant
sudo vi /etc/ssh/sshd_config
PasswordAuthentication yes
sudo systemctl restart sshd

# master上执行(分别发送到node1和node2)
sudo scp -r /opt/etcd/ vagrant@192.168.1.31:~/etcd/
sudo scp -r /usr/lib/systemd/system/etcd.service vagrant@192.168.1.31:~
# node1和node2
vagrant ssh node1
sudo cp -r etcd /opt/etcd/
sudo cp etcd.service /usr/lib/systemd/system/etcd.service
sudo vi /opt/etcd/cfg/etcd
## 修改ip地址成本机ip: ETCD_NAME ETCD_LISTEN_PEER_URLS ETCD_LISTEN_CLIENT_URLS ETCD_INITIAL_ADVERTISE_PEER_URLS ETCD_ADVERTISE_CLIENT_URLS
sudo systemctl daemon-reload
sudo systemctl enable etcd
sudo systemctl restart etcd

# master验证集群是否搭建成功
## 创建命令别名
sudo touch /opt/etcd/bin/my_etcdctl
sudo chown vagrant /opt/etcd/bin/my_etcdctl
sudo chmod 0744 /opt/etcd/bin/my_etcdctl
cat > /opt/etcd/bin/my_etcdctl <<EOF
#!/bin/bash
/opt/etcd/bin/etcdctl --ca-file=/opt/etcd/ssl/ca.pem --cert-file=/opt/etcd/ssl/server.pem --key-file=/opt/etcd/ssl/server-key.pem --endpoints="https://192.168.1.30:2379,https://192.168.1.31:2379,https://192.168.1.32:2379" \$@
EOF
sudo /opt/etcd/bin/my_etcdctl cluster-health
```

### node安装docker

```shell
# 安装必要的系统工具
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
# 添加软件源信息
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# 更新并安装 Docker-CE
sudo yum makecache fast
sudo yum -y install docker-ce
# 镜像加速器(https://www.daocloud.io/mirror)
curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://f1361db2.m.daocloud.io
# 重启docker
sudo systemctl restart docker
```

### 网络模型(CNI)

```shell
#常用的网络模型: flannel 或 calico 或 Contiv
# 我们使用Flannel ---- 基于overlay网络

# 部署Flannel
## 给master的flannel分配子网
sudo /opt/etcd/bin/my_etcdctl set /coreos.com/network/config '{"Network": "172.17.0.0/16", "Backend": {"Type": "vxlan"}}'
# 查看分配情况
sudo /opt/etcd/bin/my_etcdctl get /coreos.com/network/config
# 下载二进制包(node节点上)
## https://github.com/coreos/flannel/releases
cd /home/vagrant
curl -OL https://github.com/coreos/flannel/releases/download/v0.11.0/flannel-v0.11.0-linux-amd64.tar.gz
mkdir flannel-v0.11.0-linux-amd64
cd flannel-v0.11.0-linux-amd64
tar -zxvf ../flannel-v0.11.0-linux-amd64.tar.gz
sudo mkdir -p /opt/kubernetes/{bin,cfg,ssl}
sudo mv flanneld mk-docker-opts.sh /opt/kubernetes/bin/
# 部署与配置flannel(node节点上)
cd /home/vagrant
vi flannel.sh
## 内容见附录: flannel.sh
## systemd管理Flannel
## 配置Docker使用Flannel生成子网
## 启动Flannel
chmod a+x flannel.sh
sudo ./flannel.sh https://192.168.1.30:2379,https://192.168.1.31:2379,https://192.168.1.32:2379

# 验证安装
ps -ef |grep doc
## 出现： --bip=172.17.59.1/24
ip addr
## docker0与flannal处于同一网段
sudo /opt/etcd/bin/my_etcdctl ls /coreos.com/network
```

<https://edu.51cto.com//center/course/lesson/index?id=235852>

   















































## 附录: etcd-cert.sh

```shell
export PATH=$PATH:/usr/local/bin
cat > ca-config.json <<EOF
{
  "signing": {
    "default": {
      "expiry": "87600h"
    },
    "profiles": {
      "www": {
         "expiry": "87600h",
         "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
        ]
      }
    }
  }
}
EOF

cat > ca-csr.json <<EOF
{
    "CN": "etcd CA",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "Beijing",
            "ST": "Beijing"
        }
    ]
}
EOF

cfssl gencert -initca ca-csr.json | cfssljson -bare ca -

#-----------------------

cat > server-csr.json <<EOF
{
    "CN": "etcd",
    "hosts": [
    "192.168.1.30",
    "192.168.1.31",
    "192.168.1.32"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "BeiJing",
            "ST": "BeiJing"
        }
    ]
}
EOF

cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=www server-csr.json | cfssljson -bare server

```



```shell
export PATH=$PATH:/usr/local/bin
cat > ca-config.json <<EOF
{
  "signing": {
    "default": {
      "expiry": "87600h"
    },
    "profiles": {
      "www": {
         "expiry": "87600h",
         "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
        ]
      }
    }
  }
}
EOF

cat > ca-csr.json <<EOF
{
    "CN": "etcd CA",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "Beijing",
            "ST": "Beijing"
        }
    ]
}
EOF

cfssl gencert -initca ca-csr.json | cfssljson -bare ca -

#-----------------------

cat > server-csr.json <<EOF
{
    "CN": "etcd",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "BeiJing",
            "ST": "BeiJing"
        }
    ]
}
EOF

cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=www server-csr.json | cfssljson -bare server

```

## 附录: vagrantfile

```shell
# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.require_version ">= 1.6.0"

boxes = [
    {
        :name => "swarm-manager",
        :eth1 => "192.168.1.30",
        :mem => "1024",
        :cpu => "1"
    },
    {
        :name => "swarm-worker1",
        :eth1 => "192.168.1.31",
        :mem => "1024",
        :cpu => "1"
    },
    {
        :name => "swarm-worker2",
        :eth1 => "192.168.1.32",
        :mem => "1024",
        :cpu => "1"
    }
]

Vagrant.configure(2) do |config|

  config.vm.box = "centos/7"
  boxes.each do |opts|
      config.vm.define opts[:name] do |config|
        config.vm.hostname = opts[:name]
        config.vm.provider "vmware_fusion" do |v|
          v.vmx["memsize"] = opts[:mem]
          v.vmx["numvcpus"] = opts[:cpu]
        end

        config.vm.provider "virtualbox" do |v|
          v.customize ["modifyvm", :id, "--memory", opts[:mem]]
          v.customize ["modifyvm", :id, "--cpus", opts[:cpu]]
        end

        config.vm.network "public_network", ip: opts[:eth1], bridge: "Realtek PCIe GBE Family Controller"
      end
  end

  config.vm.synced_folder "./labs", "/home/vagrant/labs"
  config.vm.provision "shell", privileged: true, path: "./setup.sh"

end
```

## 附录: cfssl.sh

```shell
curl -L https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -o /usr/local/bin/cfssl
curl -L https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -o /usr/local/bin/cfssljson
curl -L https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64 -o /usr/local/bin/cfssl-certinfo
chmod +x /usr/local/bin/cfssl /usr/local/bin/cfssljson /usr/local/bin/cfssl-certinfo
```

## 附录: etcd.sh

```shell
#!/bin/bash
# example: ./etcd.sh etcd01 192.168.1.10 etcd02=https://192.168.1.11:2380,etcd03=https://192.168.1.12:2380

ETCD_NAME=$1
ETCD_IP=$2
ETCD_CLUSTER=$3

WORK_DIR=/opt/etcd

cat <<EOF >$WORK_DIR/cfg/etcd
#[Member]
ETCD_NAME="${ETCD_NAME}"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://${ETCD_IP}:2380"
ETCD_LISTEN_CLIENT_URLS="https://${ETCD_IP}:2379"

#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://${ETCD_IP}:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://${ETCD_IP}:2379"
ETCD_INITIAL_CLUSTER="etcd01=https://${ETCD_IP}:2380,${ETCD_CLUSTER}"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_INITIAL_CLUSTER_STATE="new"
EOF

cat <<EOF >/usr/lib/systemd/system/etcd.service
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile=${WORK_DIR}/cfg/etcd
ExecStart=${WORK_DIR}/bin/etcd \
--name=\${ETCD_NAME} \
--data-dir=\${ETCD_DATA_DIR} \
--listen-peer-urls=\${ETCD_LISTEN_PEER_URLS} \
--listen-client-urls=\${ETCD_LISTEN_CLIENT_URLS},http://127.0.0.1:2379 \
--advertise-client-urls=\${ETCD_ADVERTISE_CLIENT_URLS} \
--initial-advertise-peer-urls=\${ETCD_INITIAL_ADVERTISE_PEER_URLS} \
--initial-cluster=\${ETCD_INITIAL_CLUSTER} \
--initial-cluster-token=\${ETCD_INITIAL_CLUSTER_TOKEN} \
--initial-cluster-state=new \
--cert-file=${WORK_DIR}/ssl/server.pem \
--key-file=${WORK_DIR}/ssl/server-key.pem \
--peer-cert-file=${WORK_DIR}/ssl/server.pem \
--peer-key-file=${WORK_DIR}/ssl/server-key.pem \
--trusted-ca-file=${WORK_DIR}/ssl/ca.pem \
--peer-trusted-ca-file=${WORK_DIR}/ssl/ca.pem
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable etcd
systemctl restart etcd


```

## 附录: flannel.sh

```shell
#!/bin/bash

ETCD_ENDPOINTS=${1:-"http://127.0.0.1:2379"}

cat <<EOF >/opt/kubernetes/cfg/flanneld

FLANNEL_OPTIONS="--etcd-endpoints=${ETCD_ENDPOINTS} \
-etcd-cafile=/opt/etcd/ssl/ca.pem \
-etcd-certfile=/opt/etcd/ssl/server.pem \
-etcd-keyfile=/opt/etcd/ssl/server-key.pem"

EOF

cat <<EOF >/usr/lib/systemd/system/flanneld.service
[Unit]
Description=Flanneld overlay address etcd agent
After=network-online.target network.target
Before=docker.service

[Service]
Type=notify
EnvironmentFile=/opt/kubernetes/cfg/flanneld
ExecStart=/opt/kubernetes/bin/flanneld --ip-masq \$FLANNEL_OPTIONS
ExecStartPost=/opt/kubernetes/bin/mk-docker-opts.sh -k DOCKER_NETWORK_OPTIONS -d /run/flannel/subnet.env
Restart=on-failure

[Install]
WantedBy=multi-user.target

EOF

cat <<EOF >/usr/lib/systemd/system/docker.service

[Unit]
Description=Docker Application Container Engine
Documentation=https://docs.docker.com
After=network-online.target firewalld.service
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile=/run/flannel/subnet.env
ExecStart=/usr/bin/dockerd \$DOCKER_NETWORK_OPTIONS
ExecReload=/bin/kill -s HUP \$MAINPID
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity
TimeoutStartSec=0
Delegate=yes
KillMode=process
Restart=on-failure
StartLimitBurst=3
StartLimitInterval=60s

[Install]
WantedBy=multi-user.target

EOF

systemctl daemon-reload
systemctl enable flanneld
systemctl restart flanneld
systemctl restart docker


```








