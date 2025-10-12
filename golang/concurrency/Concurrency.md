## Concurrency and Parallelism in Golang

Concurrency is one of the most important topic while developing scalable applications in Golang.

This documentation along with the code examples helps you in covering all the necessary topics for a quick revision. 

For Concurrency exercises refer to this [repo](https://github.com/loong/go-concurrency-exercises) in Github it has some of the problems which helps you in getting a good understanding of concurrency in Golang.


#### Understanding the difference between Concurrency and Parallism
Concurrency basically means doing many things at a time where as parallism means doing many things at the same time.

_Let's take the example of a chef in a kitchen_ - 

- A chef preparing multiple dishes at the same time not necessarily simultaneously is called Concurrency.
- Multiple chefs preparing different dishes in the kitchen at the same time can be termed as Parallism.

You have tools like Go routines and channels to support **concurrency** in Golang and when you spin off multiple go routines and the go scheduler maps them into multiple OS threads across the CPU cores, in that scenario **parallism** is achieved in golang

You can set the `GOMAXPROCS` variable value to have that control.


```go
runtime.GOMAXPROCS(n)
```

**Note:**
> _n refers to the number of CPU cores you wish to use_

#### Go routines - 
Go routines are light weighted threads scheduled by the Go runtime. They differ from the threads in other high level langauges because they are not scheduled by the OS, where there is a cost for context switching.

The default size for a thread scheduled by the os can be somewhere between **~1mb** where as the go routines sizes range somewhere between **~1kb** a much smaller stack size.

Go routines follow **M:N model**.

#### M:N Model - 
The M:N model specifies that `M` number of Go routines will be scheduled onto `N` OS threads, which run on `P` logical cores, this differs from the traditional threads which follows **1:1** model, Each task is mapped to Each thread.


#### Channels - 
Channels are custom types in Golang, which are used to transmit data between Go routines. Basically, Go routines communicate between each other through channels.

There are of two types channels:
- **Bufferred channels-** It has a fixed size they are created using the syntax, `ch:=make(chan int, 2)` They help in achieving async communication which means the channel is only blocked when the buffer is full or the buffer is empty.

- **Unbufferred channels-** Unbufferred channels have no fixed size, ideal for sync communication between go routines.


#### Main Goroutines Exit - 
When the main go routine exits all the go routines terminate regardless of anything.


#### Mutex - 
Mutex is a synchronization primitive in Golang which makes sure one go routine has access to a data member at a time. It is used to tackle race conditions.

```go
var mx sync.mutex
mx.Lock()
// data member or business operation
mx.Unlock()
```
A mutex has less overhead than channels in Golang. It is used in critical scenarios where shared access needs to be synchronized.

#### Mutex and Channels: When to use what?
A mutex is supposed to protect access for shared memory, where as channels are used for communication message passing. They both operate in different scope catering different business requirement.

**Mutex-** Mutex can be used when the write contention is more then read operations.

- It helps in attaining a single lock.
- It is more simple to implement.

**RWMutex-** When you need to implement dual lock both on read and write modes, use RWmutex.
    
- Ideal for scenarios where you have more read contention.
- You can put lock on read opeations.
- A bit complex to implement.

**TryLock and TryRLock-** They are non-blocking versions of **Lock()** and **RLock()**. They return false if the lock is not successful.

*Classic Scenario for Deadlock-* If you forget to unlock a mutex, the hold will be eternal. Thus, resulting in dead lock.

*Unlocking the same mutex twice-* If you unlock the same mutex twice a panic will be thrown.

