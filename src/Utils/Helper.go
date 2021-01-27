package Utils

import (
	"bytes"
	"strings"
)

func Join(strs ...string) string {
	ret :=""
	for _,s:=range strs{
		ret +=s
	}
	return ret
}
func Join1(strs ...string) string {

	return strings.Join(strs,"")
}
func Join2(strs ...string) string {
	var bf bytes.Buffer
	for _,s :=range strs {
		bf.WriteString(s)
	}
	return strings.Join(strs,"")
}