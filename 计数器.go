package main
import "fmt"
func main(){
	j:=0
	count:=func(jishuqi int) int{return jishuqi+1}
	for i:=1;i<=10;i++{
		j=count(j)
		fmt.Println(j)
	}
}