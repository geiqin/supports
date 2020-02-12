package config

import (
	"github.com/micro/go-micro/config"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"strings"
)

var conf map[string]interface{}

func Load(pathDir string)  {
	files :=LoadConfList(pathDir)
	conf =make(map[string]interface{})
	for k, v := range files {
		file  := pathDir +"/"+v
		err := config.LoadFile(file)
		if err !=nil{
			log.Println("load config name : ",err)
			log.Println("load configs failed")
			return
		}
		conf[k] = config.Map()
	}
	log.Println("load configs succeed")
}

func GetConfig(name string , keyPaths ...string) map[string]interface{} {

	var oneMap map[string]interface{}
	var oneVal interface{}
	oneOk :=false
	channel,ok :=conf[name]
	if !ok {
		log.Println("error: config name not exits")
		return nil
	}

	oneMap =channel.(map[string]interface{})
	plen :=len(keyPaths)

	if plen >0 {
		for i, k :=range keyPaths{
			oneVal,oneOk =oneMap[k]
			if !oneOk {
				return nil
			}
			kid:=reflect.TypeOf(oneVal)
			if kid.Kind().String() !="map"{
				log.Println("不能取节点为值类型，节点必须集合型, pathkey:",k)
				return nil
			}
			if i < plen{
				oneMap =oneVal.(map[string]interface{})
			}
		}
		oneMap =oneVal.(map[string]interface{})
	}
	return oneMap
}

func LoadConfList(dirname string) map[string]string {
	suffix :=".json"
	fileList :=map[string]string{}
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Println("not configs :",err)
		return fileList
	}

	for _, file := range files {
		fileName :=file.Name()
		fileSuffix := path.Ext(fileName)

		if fileSuffix ==suffix{
			filePrefix := strings.TrimSuffix(fileName, fileSuffix)
			fileList[filePrefix] =fileName
		}
	}
	return fileList
}