## 免登陆安装sourcetree

```python
# 安装后跳转到目录
%LocalAppData%\Atlassian\SourceTree\

# 新建文件accounts.json
accounts.json

# 添加
[
  {
    "$id": "1",
    "$type": "SourceTree.Api.Host.Identity.Model.IdentityAccount, SourceTree.Api.Host.Identity",
    "Authenticate": true,
    "HostInstance": {
      "$id": "2",
      "$type": "SourceTree.Host.Atlassianaccount.AtlassianAccountInstance, SourceTree.Host.AtlassianAccount",
      "Host": {
        "$id": "3",
        "$type": "SourceTree.Host.Atlassianaccount.AtlassianAccountHost, SourceTree.Host.AtlassianAccount",
        "Id": "atlassian account"
      },
      "BaseUrl": "https://id.atlassian.com/"
    },
    "Credentials": {
      "$id": "4",
      "$type": "SourceTree.Model.BasicAuthCredentials, SourceTree.Api.Account",
      "Username": "",
      "Email": null
    },
    "IsDefault": false
  }
]
```

## 配置使用本地的ssh key

```python
# 工具 -> 选项 -> 一般 -> ssh客户端配置 -> ssh客户端 -> 选择openssh
```

## 修改提交规则

```python
# 拉取 -> (去掉)立即提交合并的改动
```





