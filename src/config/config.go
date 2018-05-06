// Package config toml语法配置组件
/*关于配置:
1.各组件功能的配置往往有三种方式:
  a.默认值;
  b.配置文件;
  c.使用者代码直接设置
2.因为配置文件是最灵活的,按理三个配置优先级b>c>a这样应该是比较理想的.
3.但因为使用者代码设置较难控制,需要每个组件都仔细封装设置的代码,避免代码设置覆盖配置文件,有时甚至是不可能的任务.
4.所以gobase约定/实现的优先级顺序是: c>b>a, 也就是说一旦有代码设置,配置文件就无效了.
5.建议:能用配置文件控制的,使用者不要用代码再设置.
*/
package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mcuadros/defaults"
)

// ConfPath 默认配置文件路径，可修改
var ConfPath = flag.String("confpath", "./conf/config.toml", "config file path")

// Duration duration for config parse
type Duration time.Duration

func init() {
	flag.Parse()
}

func (d Duration) String() string {
	dd := time.Duration(d)
	return dd.String()
}

// GoString  duration go string
func (d Duration) GoString() string {
	dd := time.Duration(d)
	return dd.String()
}

// Duration duration
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

// UnmarshalText 字符串解析时间
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	dd, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(dd)
	}
	return err
}

// Parse parse config with default and config file ../conf/config.toml
func Parse(c interface{}) error {
	defaults.SetDefaults(c)
	fmt.Printf("\ndefault config: %+v", c)
	return ParseConfigWithoutDefaults(c)
}

// ParseConfig same as Parse
func ParseConfig(c interface{}) error {
	return Parse(c)
}

// ParseConfigWithPath 自己定义配置文件路径
func ParseConfigWithPath(c interface{}, path string) error {
	defaults.SetDefaults(c)
	fmt.Printf("\ndefault config: %+v", c)
	fmt.Printf("\nconfig file:%s", path)
	if _, err := toml.DecodeFile(path, c); err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Printf("toml config: %+v\n", c)
	return nil
}

// ParseConfigWithoutDefaults no default value
func ParseConfigWithoutDefaults(c interface{}) error {
	fmt.Printf("\nconfig file:%s", *ConfPath)
	if _, err := toml.DecodeFile(*ConfPath, c); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("\ntoml config: %+v\n", c)
	return nil
}
