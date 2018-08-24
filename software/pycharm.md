## GLOBAL

### 使用Python环境

```shell
Pycharm->Preferences->Project:名称->Project Interpreter->点击齿轮->add->Existing environment->选择Python版本路径即可
```

## MAC

## 补充依赖

```python
option + enter
```

## 设置文件编码

```python
File->settings->Editor->File Encodings->Global Encoding和Project Encoding改成UTF-8
```

## 修改快捷键

```python
# Main menu -> Edit -> Macros -> maohao
# Main menu -> View ->Toggle Full Screen mode -> F11
# Tools -> Python Console -> f12
```

## 设置命令行初始化

```python
cmd.exe "/k" activate "F:\www\python\proxy_pool\venv" 
```
## 设置文件夹属性

```python
phpstorm 为了方便用户管理项目目录，目前可以将项目文件夹设置为 4 类 Test,Sources,Excluded,Resource Root。

1. Test (颜色为绿色)
       > 测试主目录，如 `Laravel` 的 `tests` 目录
2. Sources (颜色为蓝色)
       > 项目主代码目录，如 `Laravel` 的 `app` 目录
3. Excluded (颜色为红色)
       > 第三方扩展依赖(不会修改代码)，不建立索引，不由`phpstorm`管理，如 `Laravel` 的 `vendor` 目录
4. Resource Root (颜色为紫色)
       > 前端资源，如 `Laravel` 的 `public` 目录
   合理设置项目的目录是有作用的，如

- 设置 Test 目录，可以在project勾选只显示 Test,方便测试时查看
- 设置 Excluded 目录，可以减少 phpstorm 建立索引的时间
- 设置 Resource Root 目录，以 Laravel 为例，可以帮助检测模板文件的资源路径

```