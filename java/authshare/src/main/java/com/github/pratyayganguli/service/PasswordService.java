package com.github.pratyayganguli.service;

import com.github.pratyayganguli.model.HttpResponse;
import com.github.pratyayganguli.model.Password;
import com.github.pratyayganguli.util.PasswordGenerator;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.UUID;

@Service
public class PasswordService {
    public HttpResponse<Password> createPassword(@RequestBody Password password) {
        // todo: validate the request,
        //       set the values
        //       propagate it
        byte[] passBytes = PasswordGenerator.GeneratePasswordBytes();
        password.setId(UUID.randomUUID().toString());
        password.setPassword(passBytes);
        password.setCreatedAt(System.currentTimeMillis());
        HttpResponse<Password> response = new HttpResponse<>();
        response.setData(password);
        response.setCode(200);
        response.setMessage("Password created successfully!");
        return response;
    }


    public HttpResponse<Password> getPassword(@PathVariable String id) {
        return null;
    }
}


