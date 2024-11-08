#### 1.

```go
  if len(consumeMSG) > 0 {
             //进行异步消费 
             go func() {
                m := consumeMSG[:ConsumeNum]
                fn(m)
             }()
```

如果切片长度小于5会越界，应该改为

```go
m:=consumeMSG[:len(consumeMSG)]

consumeMSG=consumeMSG[len(consumeMSG):]
```

#### 2.

```go
if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute
```

如果这个if成立后快速产生了多条数据存入切片中会导致前一个if语句判断执行，这样两个if下面的协程同时进行，会导致冲突，下一行if应该改为

```go
if len(consumeMSG)>0&&len(consumeMSG)<ConsumeNum 
```

并在执行是加锁防止上一个if下的协程执行

#### 3.

for循环下面的协程可能还没进行完就清除插入数据了，同时也可能上一次协程还没结束下一次协程就开始了，可以加个锁和waitgroup

```go
 if len(consumeMSG) >= ConsumeNum {
          //进行异步消费
     lock1.Lock()
       if len(consumeMSG) >= ConsumeNum {//再次判断更新后是否还大于
          var wait sync.WatiGroup
           wait.Add(1)
           go func() {
             m := consumeMSG[:ConsumeNum]
             fn(m)
             wait.Done()
          }()
          // 更新上次消费时间
           wait.Wait()
          lastConsumeTime = time.Now()
          // 清除插入的数据
          consumeMSG = consumeMSG[ConsumeNum:]
       }
     lock1.Unlock()
       } 
```

下面那个协程也这样改。
