package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

// 配置数据
type settingData struct {
	XMLName    xml.Name       `xml:"GoLang,element"`
	Enviroment []keyValueItem `xml:"EnviromentVar,element"`
}

// 键值对
type keyValueItem struct {
	Key   string `xml:"Name,attr"`
	Value string `xml:",innerxml"`
}

const (
	con_FileName string = "GoSetting.xml"
)

var goSettingData settingData

// 包初始化
func init() {
	er := initSettingData()
	if er != nil {
		panic(er)
	}
}

// 从xml文件加载配置数据
func initSettingData() error {
	content, err := ioutil.ReadFile(con_FileName)
	if err != nil {
		return err
	}

	var result settingData
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return err
	}

	goSettingData = result

	return nil
}

func writeSettingToFile(data settingData) error {
	byteData, err := xml.Marshal(&data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("GoSetting2.xml", byteData, os.ModeAppend)
}

// 读取配置项
func getSetting(key string) string {

	for i := 0; i < len(goSettingData.Enviroment); i++ {
		if strings.ToLower(goSettingData.Enviroment[i].Key) == strings.ToLower(key) {
			return goSettingData.Enviroment[i].Value
		}
	}

	return ""
}

// 填充配置项到envData
// envData：等待被填充的配置项
func fill(envData []string) []string {

	// 标记是否存在对应的配置项
	ifFind := false

	for _, val := range goSettingData.Enviroment {

		// 重置是否存在
		ifFind = false

		// 把以key开头的字符替换成key=valu的形式
		for index, item := range envData {
			if strings.HasPrefix(strings.ToLower(item), strings.ToLower(val.Key)) {

				// 值更新
				envData[index] = strings.ToUpper(val.Key) + "=" + val.Value
				ifFind = true

				break
			}
		}

		// 不存在，则添加
		if ifFind == false {
			envData = append(envData, strings.ToUpper(val.Key)+"="+val.Value)
		}
	}

	return envData
}
