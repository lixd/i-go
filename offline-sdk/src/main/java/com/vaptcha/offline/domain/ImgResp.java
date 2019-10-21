package com.vaptcha.offline.domain;

public class ImgResp {
    private String code;
    private String imgid;
    private String challenge;
    private String msg;
    private String key;

    public ImgResp(String code, String imgid, String challenge, String msg, String key) {
        this.code = code;
        this.imgid = imgid;
        this.challenge = challenge;
        this.msg = msg;
        this.key = key;
    }

    public ImgResp() {
    }

    @Override
    public String toString() {
        return "ImgResp{" +
                "code='" + code + '\'' +
                ", imgid='" + imgid + '\'' +
                ", challenge='" + challenge + '\'' +
                ", msg='" + msg + '\'' +
                ", key='" + key + '\'' +
                '}';
    }

    public String getKey() {
        return key;
    }

    public void setKey(String key) {
        this.key = key;
    }

    public String getCode() {
        return code;
    }

    public void setCode(String code) {
        this.code = code;
    }

    public String getImgid() {
        return imgid;
    }

    public void setImgid(String imgid) {
        this.imgid = imgid;
    }

    public String getChallenge() {
        return challenge;
    }

    public void setChallenge(String challenge) {
        this.challenge = challenge;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }
}
