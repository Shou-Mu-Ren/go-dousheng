# douyin

## 抖音项目服务端

## 项目环境依赖



## 项目运行

```shell
#运行项目
go run serve.go router.go
#打包
go build -ldflags "-w -s" serve.go router.go
```

### 功能说明


* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

