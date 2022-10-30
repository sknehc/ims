# ims
socket通信
##构建
go build -o app main.go
##作为客户端
./app -mode client -ip 127.0.0.1 -port 8523
##作为服务端
./app -mode server -ip 127.0.0.1 -port 8523
