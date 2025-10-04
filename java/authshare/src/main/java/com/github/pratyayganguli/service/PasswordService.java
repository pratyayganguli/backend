package com.github.pratyayganguli.service;

import com.github.pratyayganguli.model.Password;
import com.github.pratyayganguli.utils.PasswordGenerator;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.UUID;

@Service
public class PasswordService {
    public Password createPassword(@RequestBody Password password) {
        // todo: validate the request,
        //       set the values
        //       propagate it
        byte[] passBytes = PasswordGenerator.GeneratePasswordBytes();
        password.setId(UUID.randomUUID().toString());
        password.setPassword(passBytes);
        password.setCreatedAt(System.currentTimeMillis());
        return password;
    }

    public Password getPassword(@PathVariable String id) {
        return null;
    }
}


