### Redis Guide: (Basic to Advanced Concepts) ###

#### Redis: ####
Redis is an in-memory (key-value) data store primarily used as database. The reason for fast retrieval is because of in-memory storage and not disk based storage.


#### Persistent storage in Redis: ####

- **RDB**
    - Redis offers persistent storage but in the form of RDB (Redis Database Backup), creates snapshot of data in specific time intervals.
    - For eg. At every regular intervals you can have the changes saved if a certain number of keys are saved.
    - The data is stored in `.rdb` file.

- **Append Only File**
    - It is a mechanism offered by Redis that logs every write operations which the server processes to a file called `appendOnly.aof`
    - If the server redis restarts it replays the AOF log to rebuild the exact state of the database.




