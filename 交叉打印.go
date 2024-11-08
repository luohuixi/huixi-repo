package main
import (
	"fmt"
	"sync"
)
func main(){
	a:=0
	var lock1,lock2,lock3,lock4 sync.Mutex
	var wait sync.WaitGroup
	ch:=make(chan string)
	English:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	number:="0123456789"
	wait.Add(1)
	lock4.Lock()
	go func(){	
        for k,value:=range English{
			lock1.Lock()
			lock3.Lock()
			ch<-string(value)
			if (k%2==1&&a==0) {lock4.Unlock()} else {lock3.Unlock()}
			lock1.Unlock()
		}
		wait.Done()
	}()
	go func(){
        for k,value:=range number{
			lock2.Lock()
			lock4.Lock()
			ch<-string(value)
			if k==9 {
				a=1
				lock3.Unlock()
				return
			}
			if k%2==1 {lock3.Unlock()} else {lock4.Unlock()}
			lock2.Unlock()
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