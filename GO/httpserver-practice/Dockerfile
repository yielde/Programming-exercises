FROM centos:latest

LABEL maintaienr="Galen Tong<937013596@qq.com>"
WORKDIR /root

ADD ./http_server /root/

EXPOSE 8888

ENTRYPOINT [ "/root/http_server" ]