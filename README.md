# 环境

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