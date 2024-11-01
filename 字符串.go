package main
import "fmt"
func main(){
	k:=0
    s:="hello，世界"
	fmt.Println(len(s))
	for o:=range s{
		fmt.Println(o)
		k++
	}
	fmt.Println(k)
}