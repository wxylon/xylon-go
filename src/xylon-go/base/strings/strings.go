package main

import (
	"fmt"
	"strings"
)

func main() {
	/**byte转字符串*/
	char := []byte{'a', 'b', 'c'}
	char2str := string(char)
	fmt.Println(char2str)

	/**字符串转byte*/
	str2char := []byte(char2str)
	fmt.Println(str2char)

	/**字符串长度，数组长度*/
	charlen, strlen := len(str2char), len(str2char)
	fmt.Println(charlen, strlen)

	/**字符串截取*/
	fmt.Println(char[0:len(char)-1], char2str[0:len(char2str)-1])

	/**字符串比较*/
	fmt.Println(char2str == "abc")

	/**字符串拆分 strings.Split() */
	fmt.Printf("%q\n", strings.Split(",a,b,c,d,e,", ","))

	/**	
		试图修改字符串的做法都是被禁止的
		char2str[0] = 'x';
		var p *string = &s
		(*p)[1] = 'y';
	*/
	
	
}
