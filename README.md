<p align="center"><img src="https://i.loli.net/2020/04/07/nAzjDJlX7oc5qEw.png" width="400"></p>

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-blue" alt="license MIT"></a>
<a href="https://github.com/assimon/dujiaoka-tools/releases/tag/v1.0.0"><img src="https://img.shields.io/badge/version-1.0.0-red" alt="version 1.0.0"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/Golang-1.16-lightgrey" alt="golang"></a>
<a href="https://shang.qq.com/wpa/qunwpa?idkey=37b6b06f7c941dae20dcd5784088905d6461064d7f33478692f0c4215546cee0"><img src="https://img.shields.io/badge/QQ%E7%BE%A4-568679748-green" alt="QQ群：568679748"></a>
</p>


## 独角数卡-工具箱

采用Golang编写，支持全平台编译运行，主要用于独角数卡`迁移`，`更新`，`帮助`等等一系列功能。

## 目前实现的功能

- [x] 1.8.x版本数据迁移至2.x版本


## 安装方式

### 一、自行编译
此安装方式多用于开发者，需电脑上安装`go语言`环境。   
[go语言官网](https://golang.org/)  
下载源代码：  
```shell
# 克隆仓库
git clone https://github.com/assimon/dujiaoka-tools
# 进入源码目录
cd dujiaoka-tools
```
编译命令：   
```go
go build -o  djtools
```
执行：   
```shell
./djtools -h
```

### 二、下载已经编译好的二进制程序
此方式多用于小白，用完就走的那种。  
进入打包好的版本列表，下载程序：[https://github.com/assimon/dujiaoka-tools/releases](https://github.com/assimon/dujiaoka-tools/releases) 
运行：   
```shell
./dujiaoka-tools -h
```
