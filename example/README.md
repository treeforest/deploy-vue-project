# 示例

使用时将 dist 目录**替换**即可。

* dist： 编译好的vuejs项目静态资源
* server： 编译好的Go静态文件资源服务

## 二进制部署

* 直接运行

  ```shell
  $ ./server
  ```

* 以守护进程的方式运行

  ```shell
  $ (./server &)
  ```

## Docker部署

1. 构建镜像

   ```shell
   $ docker build -t webserver:v1 .
   ```

2. 启动容器

   ```shell
   $ docker run -v ./dist:/dist -p 80:80 --name webserver -d webserver:v1
   ```

   