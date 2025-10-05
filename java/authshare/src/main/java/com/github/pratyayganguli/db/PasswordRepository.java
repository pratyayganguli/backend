package com.github.pratyayganguli.db;

import java.util.List;
import java.util.concurrent.CompletableFuture;

public interface PasswordRepository<T> {
    CompletableFuture<Void> create(T data);
    CompletableFuture<T> getPassword(String id);
    CompletableFuture<List<T>> getPasswords(int limit, int offset);
}
