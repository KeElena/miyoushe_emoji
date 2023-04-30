# Go版

**基于Go标准库，通过API接口获取表情包资源url然后下载到本地，速度非常快（有对表情包分类）**

**一、使用方式**

* 安装Go编译器[安装包地址](https://golang.google.cn/)，window用户推荐下载`.msi`，linux用户参考[golang的安装与配置](https://gitee.com/fish_keqing/environment-configuration/blob/master/operation/GoLang的安装与环境配置.md)

```go
go run main.go
```

# Python版

**使用selenium模拟用户浏览页面爬取内容**

## 环境

**一、拉取库**

```bash
pip install selenium
```

```bash
pip install urllib
```

**二、设置浏览器driver**

* [下载对应chrome浏览器版本的driver](http://chromedriver.storage.googleapis.com/index.html)

* chrome浏览器处(`~/Google/Chrome/Application`)放一个driver
* 解释器处(`~/anaconda3`)放一个driver

**三、运行**

* `addr`设置emoji存储位置
* 由于用的是`XPATH`定位修改了目的地址可能需要同步修改`点击表情图标`处的`XPATH`