package main

import (
	"fmt"
	"os"
	"testing"
)

func TestWriteData(t *testing.T) {
	tmpData := settingData{
		Enviroment: make([]keyValueItem, 0),
	}

	item := keyValueItem{
		Key:   "GOPATH",
		Value: "D:/1.txt",
	}
	tmpData.Enviroment = append(tmpData.Enviroment, item)
	item = keyValueItem{
		Key:   "GOPATH",
		Value: "D:/1.txt",
	}
	tmpData.Enviroment = append(tmpData.Enviroment, item)

	fmt.Println(os.Getwd())
	er := writeSettingToFile(tmpData)
	if er != nil {
		t.Log(er.Error())
		t.Fail()
	} else {
		t.Log("测试Ok")
	}

}
