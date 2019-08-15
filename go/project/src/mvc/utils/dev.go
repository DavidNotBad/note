package utils

import (
	"bufio"
	"github.com/ChimeraCoder/gojson"
	"os"
	"strings"
)

func JsonToStruct(jsonStr string, structName string, pkgName string, targetFile string)(err error)  {
	//将json转化成struct字符串
	i := strings.NewReader(jsonStr)
	res, err := gojson.Generate(i, gojson.ParseJson, structName, pkgName, []string{"json"}, false, true)

	//将struct字符串保存到文件里
	file, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE, 0666)
	defer func() {
		err = file.Close()
	}()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(res))
	err = writer.Flush()

	return
}
