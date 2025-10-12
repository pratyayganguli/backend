## Mongo DB: A Guide for scaling Mongo DB to Production level

###### Last Draft Date: 14th December, 2025

**Mongo DB:** Mongo db is a document based database primarily no-sql type. Some of the managed services of Mongo in Cloud plarforms are Mongo DB atlas with options for GCP and AWS deployment.

*Docker command to start the Mongo Conatiner:*
```bash
docker run --name CONTAINER_NAME -p 27017:27017 \ 
-e MONGO_INITDB_ROOT_USERNAME=YOUR_USER \
-e MONGO_INITDB_ROOT_PASSWORD=YOUR_SECREY_PASSWORD \
-e MONGO_INITDB_ROOT_DATABASE=YOUR_DATABASE_NAME \
-d mongo:latest
```

#### Sharding in Mongo DB - 

Sharding is the technique of distributing the database into uniform chunks which are stored accross multiple cluster based on the shard key.

Usually it is a good practice to use the shard key on something which is auto incremenented like `userId`

#### Mongos - 

Mongos is an stateless interface which basically routes the query across the clusters.

Configuring Mongos - 

- Setting up the config servers for metadata.
- Each replica set will need atleast 3-4 config servers for high availability.
- Then start the mongos process and add the shards to the cluster.

**Note:** Instead of connecting to the parent mongo db instance, we connect to the mongos query router. Apart from that there should not be any changes in the driver code in Go.

#### Connection pooling in Mongo - 

Instead of using a single connection in Mongo, what we do is mantain a pool of mongo connections and reuse them. Having a upper bound and lower bound is necessary in such cases for Agility.

```go
clientOpts := options.Client()
.ApplyURI("")
.SetMaxPoolSize()
.SetMinPoolSize()
.SetMaxConnIdleTime(1 * time.Minute)
```

#### High availability - 

Replication helps in mantaining high availabilty across all the regions and also enhances latency.

- Primary Node - (AZ-1)
- Read Node 1 - (AZ-2)
- Read Node 2 - (AZ-3)

If the application is read intensive then we can set up multiple read replicas across the regions to make read faster.


