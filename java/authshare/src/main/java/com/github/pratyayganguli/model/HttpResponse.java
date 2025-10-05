package com.github.pratyayganguli.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.lang.NonNull;

@Data
@AllArgsConstructor
@NoArgsConstructor

public class HttpResponse<T> {
    private T data;
    private int code;
    @NonNull
    private String message;
}
