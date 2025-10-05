package com.github.pratyayganguli.db;

import com.github.pratyayganguli.model.Password;

import java.util.List;
import java.util.concurrent.CompletableFuture;

public class RedisPasswordRepository implements PasswordRepository<Password>{
    @Override
    public CompletableFuture<Void> create(Password password) {
        return CompletableFuture.runAsync(()->{

        });
    }

    @Override
    public CompletableFuture<Password> getPassword(String id) {
        return CompletableFuture.supplyAsync(()->{
            return null;
        });
    }

    @Override
    public CompletableFuture<List<Password>> getPasswords(int limit, int offset) {
        return CompletableFuture.supplyAsync(()->{
            return null;
        });
    }
}
