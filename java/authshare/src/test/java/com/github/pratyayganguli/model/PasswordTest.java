package com.github.pratyayganguli.model;

import org.junit.jupiter.api.Test;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import java.util.Base64;
import java.util.UUID;

public class PasswordTest {

    @Test
    void testPasswordObject() {
        SecureRandom secureRandom = new SecureRandom();
        byte[] randomBytes = secureRandom.generateSeed(1024);
        String b64Val = Base64.getEncoder().encodeToString(randomBytes).substring(10, 30);
        Password password = new Password(UUID.randomUUID().toString(), b64Val.getBytes(StandardCharsets.UTF_8));
    }
}
