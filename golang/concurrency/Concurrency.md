## Concurrency and Parallelism in Golang
###### Last Draft Date: 14th December, 2025

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

