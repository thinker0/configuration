package configuration

import (
	"fmt"
)

var configText = `
####################################
# Typesafe HOCON                   #
####################################

config {
  # Comment
  version = "0.0.1"
  one-second = 1s
  one-day = 1day
  array = ["one", "two", "three"] #comment
  bar = "bar"
  foo = foo.${config.bar} 
  number = 1
  object {
    a = "a"
    b = "b"
    c = {
            d = ${config.object.a} //comment
        }
    }
}
// fallback
config.object.a="newA"
config.object.c.f="valueF"

// self reference
self-ref=1
self-ref=[${self-ref}][2]

// byte size
byte-size=10MiB

// system envs
home:${HOME}

plus-equal=foo
plus-equal+=bar

plus-equal-array=[foo]
plus-equal-array+=[bar, ${HOME}]

configList = [
	{
		a=1
	}
	{
		b=2
	}
]

`

func main() {
	conf := ParseString(configText)

	fmt.Println("config.one-second:", conf.GetTimeDuration("config.one-second"))
	fmt.Println("config.one-day:", conf.GetTimeDuration("config.one-day"))
	fmt.Println("config.array:", conf.GetStringList("config.array"))
	fmt.Println("config.bar:", conf.GetString("config.bar"))
	fmt.Println("config.foo:", conf.GetString("config.foo"))
	fmt.Println("config.number:", conf.GetInt64("config.number"))
	fmt.Println("config.object.a:", conf.GetString("config.object.a"))
	fmt.Println("config.object.c.d:", conf.GetString("config.object.c.d"))
	fmt.Println("config.object.c.f:", conf.GetString("config.object.c.f"))
	fmt.Println("self-ref:", conf.GetInt64List("self-ref"))
	fmt.Println("byte-size:", conf.GetByteSize("byte-size"))
	fmt.Println("home:", conf.GetString("home"))
	fmt.Println("default:", conf.GetString("none", "default-value"))
	fmt.Println("plus-equal:", conf.GetString("plus-equal"))
	fmt.Println("plus-equal-array:", conf.GetStringList("plus-equal-array"))
	fmt.Println("configList:", conf.GetConfigList("configList"))
}