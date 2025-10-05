package com.github.pratyayganguli.config;

import org.junit.jupiter.api.Test;
import redis.clients.jedis.*;
import redis.clients.jedis.csc.CacheConfig;

public class RedisConfigTest {
    // todo: run this test when you have a local docker container.
    @Test
    void testRedisConnection() {
        HostAndPort endPoint = new HostAndPort("localhost", 6379);
        RedisCredentials redisCredentials = new DefaultRedisCredentials("default", "authshare$001".toCharArray());
        DefaultJedisClientConfig defaultJedisClientConfig = DefaultJedisClientConfig.builder().credentials(redisCredentials).protocol(RedisProtocol.RESP3).build();
        CacheConfig cacheConfig = CacheConfig.builder().maxSize(1000).build();
        try {
            UnifiedJedis client = new UnifiedJedis(endPoint, defaultJedisClientConfig, cacheConfig);
        } catch(Exception e) {
            throw new RuntimeException("could not connect to the redis container", e);
        }
    }
}
