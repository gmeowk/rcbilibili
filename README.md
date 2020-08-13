# B站录播姬

## 简单说明

建议运行在NAS、路由器、电视盒子等不需要关机的设备中

云服务器好像会被B站屏蔽IP，自行测试

## 使用

根据对应的操作系统及架构，从[Release](https://github.com/zhu6feng/rcbilibili/releases) 页面下载最新版本的程序。

- 录制111房间的直播，文件保存在程序启动位置
```
rcbilibili -r 111
```

- 录制111和222房间的直播，文件保存在/mnts路径下，并转码MP4
```
rcbilibili -r 111,222 -v /mnts -t
```


参数说明

- -r：房间号，多房间用逗号分隔
- -v：录播文件位置，未指定的话就是当前程序启动位置
- -t：转码MP4，需要依赖[FFmpeg](https://ffmpeg.org/download.html)