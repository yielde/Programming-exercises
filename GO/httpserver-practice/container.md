1. docker build -t myserver .
2. docker run -d -p 9999:8888 myserver 测试可用
3. docker tag myserver:latest tonggalen/myserver:latest
4. docker login
5. docker push tonggalen/myserver 推送成功
6. https://hub.docker.com/r/tonggalen/myserver
