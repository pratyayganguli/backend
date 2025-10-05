package com.github.pratyayganguli.util;

import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import java.util.Base64;

public class PasswordGenerator {
    // todo: add extra methods to add a size to the password
    public static String GeneratePassword() {
        try {
            SecureRandom secureRandom = new SecureRandom();
            byte[] secureBytes = secureRandom.generateSeed(1024);
            return Base64.getEncoder().encodeToString(secureBytes).substring(10, 30);
        } catch(Exception e) {
            // todo: use custom exceptions for a better message.
            throw new RuntimeException(e);
        }
    }

    public static byte[] GeneratePasswordBytes() {
        try {
            SecureRandom secureRandom = new SecureRandom();
            byte[] secureBytes = secureRandom.generateSeed(1024);
            String b64Trimmed = Base64.getEncoder().encodeToString(secureBytes).substring(10, 30);
            return b64Trimmed.getBytes(StandardCharsets.UTF_8);
        } catch(Exception e) {
            // todo: use custom exceptions for a better message.
            throw new RuntimeException(e);
        }
    }
}
