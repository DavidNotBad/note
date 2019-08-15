package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

var conf map[string]string

func Get(key string)(value string)  {
	if len(conf) == 0 {
		//读取配置文件
		_, file, _, _ := runtime.Caller(0)
		dir, _ := filepath.Split(file)
		content, err := ioutil.ReadFile(dir + "app.json")
		if err != nil {
			panic(err)
		}

		//解析配置文件
		var c map[string]string
		err = json.Unmarshal(content, &c)
		if err != nil {
			panic(err)
		}
		//把结果解析到conf
		conf = c
	}

	//从map中获取值
	value, ok := conf[key]
	if !ok {
		panic("key \"" + key + "\" not found")
	}

	return
}


