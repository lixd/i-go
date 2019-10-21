package com.vaptcha.offline.domain;

public class Validate {
    private String token;

    public Validate() {
    }

    @Override
    public String toString() {
        return "Validate{" +
                "token='" + token + '\'' +
                '}';
    }

    public Validate(String token) {
        this.token = token;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }
}
