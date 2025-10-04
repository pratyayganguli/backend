package com.github.pratyayganguli.model;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor

public class Password {
    // UUID
    private String id;
    // save the byte value of the password
    private byte[] password;
    // generated on the flight
    private long createdAt;
    @JsonProperty("passwordMetadata")
    private PasswordMetadata passwordMetadata;
}
