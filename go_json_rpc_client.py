import json
import socket


request = {
    "method": "HelloService.Hello",  # 只保留一个 method 键
    "params": ["bobby"]
}

# 创建到 localhost 1234 端口的连接
client = socket.create_connection(("localhost", 1234))

# 发送请求数据给服务器
client.sendall(json.dumps(request).encode())

# 获取服务器的响应
rsp = client.recv(4096)
rsp = json.loads(rsp.decode())

# 打印响应内容
print(rsp["result"])
