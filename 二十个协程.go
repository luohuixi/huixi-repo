package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
type Num struct {
	Id int
	Value int
}
func main(){
	var wait sync.WaitGroup
	wait.Add(20)
	rand.Seed(time.Now().UnixNano())
    ch:=make(chan Num,20)
	for i:=0;i<20;i++{
		go func(i int) {
			a:=rand.Intn(114514)
			ch<-Num{Id:i,Value:a}
			wait.Done()
		}(i)
	}
	wait.Wait()
	close(ch)
	m:=make(map[int]int,20)
	fmt.Println("排序前")
	for ch2:=range ch{
		fmt.Printf("编号: %d, 值: %d\n", ch2.Id, ch2.Value)
	    m[ch2.Id]=ch2.Value
	}
	fmt.Println("排序后")
	for i:=0;i<20;i++{
        fmt.Printf("编号: %d, 值: %d\n", i, m[i])
	}
}