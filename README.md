# jk-cloudNativeLists
针对极客时间-云原生课程练习作业清单

## 2021-10-07
第一次作业：增加http目录，编写一个 HTTP 服务器，增加http相关知识点练习

## 2021-10-16
第二次作业

- 1、提交docker hub； 
```
docker tag httpserver firehmx/httpserver:v1
docker login
docker push firehmx/httpserver:v1

地址：docker.io/firehmx/httpserver
```
 
-  2、启动容器
```
docker run -p 9093:8888 httpserver
```

- 3、通过 nsenter 进入容器查看 IP 配置。

通过nsenter获取容器ip，可进行docker网络连通性测试
```cgo
# 查看pid
docker inspect 354 -f '{{.State.Pid}}'
30354

# 执行nsenter命令
nsenter -n -t30354

# 查看容器IP配置 ifconfig 
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.3  netmask 255.255.0.0  broadcast 0.0.0.0
        inet6 fe80::42:acff:fe11:3  prefixlen 64  scopeid 0x20<link>
        ether 02:42:ac:11:00:03  txqueuelen 0  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 8  bytes 656 (656.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
# 退出
exit

# 物理机ip配置
docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 0.0.0.0
        ether 02:42:de:5d:bb:55  txqueuelen 0  (Ethernet)
        RX packets 2388  bytes 231674 (226.2 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 2517  bytes 1687771 (1.6 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.27.0.9  netmask 255.255.240.0  broadcast 172.27.15.255
        ether 52:54:00:ee:87:05  txqueuelen 1000  (Ethernet)
        RX packets 11373521  bytes 1780959272 (1.6 GiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 11638601  bytes 3904021709 (3.6 GiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 785897  bytes 389858802 (371.7 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 785897  bytes 389858802 (371.7 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

veth5a433e9: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        ether 02:bd:5e:79:55:c6  txqueuelen 0  (Ethernet)
        RX packets 1056  bytes 136666 (133.4 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 1447  bytes 116315 (113.5 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

vethf366923: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        ether 5a:fc:ae:17:dc:a5  txqueuelen 0  (Ethernet)
        RX packets 8  bytes 656 (656.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

   
```

windows下构建linux可运行的二进制包


```goland
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

go build -o bin/amd64/httpserver httpserver.go


