package com.vaptcha.offline.domain;

public class ValidateResp {
    private int code;
    private String msg;

    @Override
    public String toString() {
        return "ValidateResp{" +
                "code=" + code +
                ", msg='" + msg + '\'' +
                '}';
    }

    public int getCode() {
        return code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public ValidateResp() {
    }

    public ValidateResp(int code, String msg) {
        this.code = code;
        this.msg = msg;
    }
}
