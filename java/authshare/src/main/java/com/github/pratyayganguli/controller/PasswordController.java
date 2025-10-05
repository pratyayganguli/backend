package com.github.pratyayganguli.controller;

import com.github.pratyayganguli.model.HttpResponse;
import com.github.pratyayganguli.model.Password;
import com.github.pratyayganguli.service.PasswordService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;


@RestController
@RequestMapping("/api/passwords")
@RequiredArgsConstructor

public class PasswordController {
    private final PasswordService passwordService;

    @PostMapping
    public HttpResponse<Password> createPassword(@RequestBody Password password) {
        return passwordService.createPassword(password);
    }

    @GetMapping
    public HttpResponse<Password> getPassword(@PathVariable String id) {
        return passwordService.getPassword(id);
    }
}
