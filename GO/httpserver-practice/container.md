1. docker build -t myserver .
2. docker run -d -p 9999:8888 myserver 测试可用
3. docker tag myserver:latest tonggalen/myserver:latest
4. docker login
5. docker push tonggalen/myserver 推送成功
6. https://hub.docker.com/r/tonggalen/myserver
7. nsenter查看容器网络

```bash
[root@localhost galen]# nsenter -t 759225 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
26: eth0@if27: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:12:00:06 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.18.0.6/16 brd 172.18.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
