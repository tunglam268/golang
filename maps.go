package main
import (
	"golang.org/x/tour/wc"
	"strings"
)
func WordCount(s string) map[string]int{
	m:=make(map[String]int)
	a: = string.Fields(s)
	for _;v :=range a{
		m[v]++
	}
	return m
}
func main(){
	wc.Test(WordCount)
}