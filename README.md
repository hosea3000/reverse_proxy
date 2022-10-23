# reverse_proxy

懒得装nginx。用golang实现了一个反向代理工具。

### 使用方式

去 release 界面下载对应系统和架构的二进制文件

在服务上执行二进制文件

./reverse_proxy -h 可以查看支持的参数

- -host 是你需要反向代理的地址
- -port 是服务启动端口

### 启动

```shell
./reverse_proxy -host=https://httpbin.org/anything -port=8080
```


