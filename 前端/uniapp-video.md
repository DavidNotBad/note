## HBuilderX安装

```shell
# 官网下载安装即可
https://www.dcloud.io/
```

## 创建第一个程序

```shell
# 新建项目 -> uni-app
# 默认首页
pages/index/index.vue

# {{title}}是数据绑定, 绑定的内容为data()模型?返回的结果
<text class="title">{{title}}</text>
data() {
  return {
  	title: 'Hello Next学院'
  }
}

```

