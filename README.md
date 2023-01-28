# douyin

go run main.go root.go router.go router_gen.go

## 发布视频接口

详细信息：[发布视频](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707520)  
说明：使用minio将链接中包含的data视频数据保存到服务器本地

### 安装minio

使用以下命令安装运行minio,`/data`替换为自己想要保存视频的目录，根据自己本地环境修改`pkg/minio/init.go`中的变量值  
    
    wget https://dl.min.io/server/minio/release/linux-amd64/minio  
    chmod +x minio  
    ./minio server /data  

### TODO
- [] 鉴权
- [] 保存视频路径到数据库