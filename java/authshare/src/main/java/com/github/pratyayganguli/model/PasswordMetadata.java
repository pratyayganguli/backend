package com.github.pratyayganguli.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.lang.NonNull;

@Data
@AllArgsConstructor
@NoArgsConstructor

public class PasswordMetadata {
    private String icon;
    @NonNull
    @JsonProperty("vendor")
    private String vendor;
}
