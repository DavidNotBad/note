## 搭建git服务器并自动同步到站点目录

https://blog.csdn.net/glx490676405/article/details/78329004

## 把文件移出版本控制
```shell
# 当我们需要删除暂存区或分支上的文件, 但本地又需要使用, 只是不希望这个文件被版本控制, 可以使用
git rm --cached file_path
git commit -m 'delete remote somefile'
git push origin master
```




