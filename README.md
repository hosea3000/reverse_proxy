# reverse_proxy

懒得装nginx。用golang实现了一个反向代理工具。

### 使用方式

去 release 界面下载对应系统和架构的二进制文件

在服务上执行二进制文件

### 准备配置文件
目录下增加 config.yaml 配置文件

```
port: 9999
proxies:
  - router_path: "/v2"
    target_url: "https://httpbin.org/anything/v2"
  - router_path: "/v1"
    target_url: "https://httpbin.org/anything/v1"
  - router_path: "/"
    target_url: "https://httpbin.org/anything/default"
```

### 启动

```shell
./reverse_proxy
```


