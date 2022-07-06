/*
 * @Author: ww
 * @Date: 2022-07-03 17:35:13
 * @Description:
 * @FilePath: /goconf/conf.go
 */
package goconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/uerax/goconf/category"
)

type CfgFile struct {
	//采用替代方式更新数据,重载的时候刚好查询数据的概率低
	//无需加锁
	//mx 		sync.Mutex
	Path	string
	File	string
	Data	map[string]interface{}
	pre		map[string]interface{}
}

func NewCfgFile() *CfgFile {
	return &CfgFile{Data: make(map[string]interface{}, 0), pre: make(map[string]interface{}, 0)}
}

func (t *CfgFile) New() *CfgFile {
	return &CfgFile{Data: make(map[string]interface{}, 0), pre: make(map[string]interface{}, 0)}
}

/*
* read all configuration file on this path
* @param path string 文件路径
*/
func (t *CfgFile) ReadAll(path string) error {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("文件目录打开出错 : %v", err)
	}
	
	for _, file := range dir {
		p := path + "/" + file.Name()
		if file.IsDir() {
			t.ReadAll(p)
		} else {
			err = t.ReadConfig(p)
			if err != nil {
				log.Println(err)
			}
		}
	}
	t.Data = t.pre
	t.Path = path
	return nil
}

/*
* if you just have one configuration file, that you can use this method
* @param file string 配置文件绝对路径或者相对路径
*/
func (t *CfgFile) ReadConfig(file string) error {
	
	f, err := os.Stat(file)
	
	if err != nil || f.IsDir() {
		return fmt.Errorf("配置文件或者文件夹不存在 : %v", err)
	}
	
	ext := strings.ToLower(path.Ext(file))
	
	switch ext {
	case ".yaml":
		data, err := category.ReadYaml(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.pre[f.Name()] = data

	case ".yml":
		data, err := category.ReadYaml(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.pre[f.Name()] = data
	case ".json":
		data, err := category.ReadJson(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.pre[f.Name()] = data
	case ".toml":
		data, err := category.ReadToml(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.pre[f.Name()] = data
	default:
		return fmt.Errorf("配置文件类型[%s]不支持,目前仅支持yaml,yml,json,toml格式", f.Name())
	}

	t.File = file
	return nil
}

/**
* @param section string "section:如果没有可以不填"
* @param key	string	"key:必填的参数key"
* @return value interface{} "value: 自行转换成 string|slice|map 等类型"
*/
func (t *CfgFile) GetValue(in... string) (interface{}, error) {
	if len(in) < 1 || len(in) > 2 {
		return nil, fmt.Errorf("参数有误,只能是 (section, key) 或则 (key) ")
	}
	if len(in) == 1 {
		for _, d := range t.Data {
			if v, ok := d.(map[string]interface{})[in[0]]; ok {
				return v, nil
			}
		}
	}
	if len(in) == 2 {
		for _, d := range t.Data {
			if v, ok := d.(map[string]interface{})[in[0]]; ok {
				if v1, ok := v.(map[string]interface{})[in[1]]; ok {
					return v1, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("没有找到对应的配置")
}

/*
* @param filename string 文件名
* @param obj interface{} 转换的结构体
*/
func (t *CfgFile) Unmarshal4Name(name string, obj interface{}) error {
	if v, ok := t.Data[name]; ok {
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("%s 文件不存在", name) 
}

func (t *CfgFile) Reload() error {
	if len(t.Path) != 0 {
		t.ReadAll(t.Path)
	} else if len(t.File) != 0 {
		t.ReadConfig(t.File)
	} else {
		return fmt.Errorf("未读取过配置文件无法重载")
	}
	return nil
}
