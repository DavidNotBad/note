## 获取程序的进程号
```bash
@echo off
for /f "tokens=2 " %%i in ('tasklist /nh /fi "imagename eq ngrok.exe"') do (
    set pid=%%i
)
echo,%pid%
pause
```
## 隐藏命令行窗口

```python
if "%1"=="hide" goto CmdBegin
start mshta vbscript:createobject("wscript.shell").run("""%~0"" hide",0)(window.close)&&exit
:CmdBegin
```

## 结束进程

```shell
taskkill /f /fi "IMAGENAME eq ngrok.exe"
```

## 进入当前目录

```shell
cd %cd%
```

## 删除当前目录的文件

```bash
del %cd%\ngrok.log
```

