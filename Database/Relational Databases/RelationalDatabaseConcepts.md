

### Releational Database Concepts


#### What is Relational Database and when should we consider using it? ####

Relational database is a kind of database where data is stored in tabular format in a more structured format.

Usually developer considers using Relational Databases when - 
-  **Schema is consistent:** Whenever the data models are not subjected to flexibility. Relational databases prove to be effective as they offer rigid structure.

- When ACID properties are meant to be enforced:
    - **Atomicity:** All transactions are atomic in nature. It means *all or nothing.*
    - **Consitency:** It means that data in the database will remain consistent even before and after a transaction.
    - **Isolation:** It means that all the concurrent transaction will occur independently and will not interfere. 
    - **Durability:** Once data written in the disk will survive even after crash.
    
        _**For example:**_ **WAL**(Write Ahead Logging) is used in all the mordern databases to ensure data durability even if a crash occurs. Before writing in the database, the service writes sequentially on a log present in the disk, when the WAL is flushed the db can safely tell the client the data has been written to the disk.

- **Data integrity:** If you have strict type/value enforcing strategy the best thing to do over there is to opt for relational databases instead of going for any document based database.

#### Effects of Normailization on Data: ####

Normalization is a technique used in relational databases to organize data. It helps in - 
- It reduces data redunndancy.
- Improves the data integrity.
- Makes the data easy to query.

    ##### Types of Normal form: #####
    - **First normal form:** Each column must hold atomic values. No collections should be allowed in a single column.
    - **Second normal form:** 1NF + All the non-key attributes must be dependent on a single primary key
    - **Third normal form:** 2 NF+ Non key attributes must not be be dependent on other non-key attributes.
    - **BCNF:** 3NF + Every determinant must be a candidate key which will act as a foreign key on the dependent table.
