package com.vaptcha.offline.domain;

public class Offline {
    private String challenge;
    private String callback;
    private String offline_action;
    private String vid;
    private String v;

    @Override
    public String toString() {
        return "Offline{" +
                "challenge='" + challenge + '\'' +
                ", callback='" + callback + '\'' +
                ", offline_action='" + offline_action + '\'' +
                ", vid='" + vid + '\'' +
                ", v='" + v + '\'' +
                '}';
    }

    public String getV() {
        return v;
    }

    public void setV(String v) {
        this.v = v;
    }

    public String getVid() {
        return vid;
    }

    public void setVid(String vid) {
        this.vid = vid;
    }

    public String getChallenge() {
        return challenge;
    }

    public void setChallenge(String challenge) {
        this.challenge = challenge;
    }

    public String getCallback() {
        return callback;
    }

    public void setCallback(String callback) {
        this.callback = callback;
    }

    public String getOffline_action() {
        return offline_action;
    }

    public void setOffline_action(String offline_action) {
        this.offline_action = offline_action;
    }
}
