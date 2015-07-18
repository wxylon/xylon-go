//http://blog.5d13.cn/resources/goweb/07.3.html
//http://bbs.studygolang.com/thread-11-1-1.html
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "ABB"
	re, _ := regexp.Compile(`^[A-Z]{1,}$`)
	macth := re.Find([]byte(str))
	fmt.Println(string(macth))
}
