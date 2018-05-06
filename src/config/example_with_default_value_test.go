package config_test

import (
	"config"
	"fmt"
)

func init() {
	//仅用于单元测试: 修改默认的配置路径,以便测试用例可以一键通过
	*config.ConfPath = "./config.toml"
}

//config.toml
//[section]
//k1=11
//k2="v2"
//k3="1s"
//k4=[1, 2, 3]
//k5=100

var confWithDefaultValue = struct {
	Section struct {
		K1 int             `default:"1"` //默认值都必须用""引起来,即使是数值,这是golang的struct tag的要求
		K2 string          `default:"val2"`
		K3 config.Duration `default:"800ms"`
		K4 []uint64        //数组不支持设置默认值
	}
}{}

var myConf struct {
	MySection struct {
		K5 int `default:"2"`
	}
}

func Example() {
	config.Parse(&confWithDefaultValue)
	fmt.Printf("final conf:%+v\n", confWithDefaultValue)

	//不同的模块可以定义不同的结构,解析同一个toml文件,获取自己想要的配置
	config.Parse(&myConf)
	fmt.Printf("my conf:%+v\n", myConf)
	//output:
	//default config: &{Section:{K1:1 K2:val2 K3:800ms K4:[]}}
	//
	//config file:./config.toml
	//final conf:{Section:{K1:2 K2:vvv K3:1s K4:[1 2 3]}}
	//
	//default config: &{MySection:{K5:2}}
	//
	//config file:./config.toml
	//my conf:{MySection:{K5:100}}
}
