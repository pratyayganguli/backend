package com.github.pratyayganguli.controller;

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
    public Password createPassword(@RequestBody Password password) {
        // generate the random id
        // do not store the data
        // send the response
        return passwordService.createPassword(password);
    }

    @GetMapping
    public Password getPassword(@PathVariable String id) {
        return passwordService.getPassword();
    }
}
