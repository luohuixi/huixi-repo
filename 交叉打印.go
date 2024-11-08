package main
import (
	"fmt"
	"sync"
)
func main(){
	a:=0
	var lock1,lock2 sync.Mutex
	var wait sync.WaitGroup
	ch:=make(chan string)
	English:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	number:="0123456789"
	wait.Add(1)
	lock2.Lock()
	go func(){	
        for k,value:=range English{
			lock1.Lock()
			ch<-string(value)
			if (k%2==1&&a==0) {lock2.Unlock()} else {lock1.Unlock()}
		}
		wait.Done()
	}()
	go func(){
        for k,value:=range number{
			lock2.Lock()
			ch<-string(value)
			if k==9 {
				a=1
				lock1.Unlock()
				return
			}
			if k%2==1 {lock1.Unlock()} else {lock2.Unlock()}
		}
	}()
	go func(){
	   wait.Wait()
	   close(ch)
	}()  
	for char := range ch {
        fmt.Print(char)
    }
}