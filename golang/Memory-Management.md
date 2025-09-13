### Garbage Collection - 

Garbage collection is the process of reclaiming memory which is not used by the program anymore, so that it can be used for any future allocations needed. They help in memory leaks at the cost of runtime overheads.

#### Reachable and Unreachable objects - 
Those objects which are accessible from the program roots such as global variables, stack variables are considered as **Reachable Objects**, which means they are not eligible for garbage collection.

Those objects are basically unused after execution full or partial are called **Unreachable objects**.

#### Garbage collection in Golang - 

The garbage collector in Golang relies on **Tri Color Mark & Sweep Algorithm**

- High level overview:

    Classification works on three levels:
    
    - White: Mostly unreachable objects eligible for garbage collection.
    - Grey: These objects are reachable but not fully scanned.
    - Black: Reachable and fully scanned, they are not eligble for garbage collection.

- Low level overview:

    Garbage collector starts from root level objects such as global objects, stack variables. It marks them grey and then after scanning marks them as black.

    Then comes the sweep phase, any object which is unreachable (white) will be freed.
    Memory will be returned to the Go runtime and then back to the OS.

    The env variable `GOGC` helps in tuning the garbage collector, higher the value the less often it will be performing the garbage collection. The default value is `GOGC=100`, which means only collect when the heap size has growth of 100%.


#### Escape Analysis in Golang Garbage Collector - 

The garbage collector decides on the collection when the object escapes from the stack frame and goes to the heap section.




