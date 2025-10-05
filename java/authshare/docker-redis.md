
To start the Redis container locally using docker use the following command:

```bash
docker run -d --name local-redis -p 6379:6379 redis:latest redis-server --requirepass "authshare$001"
```