package com.vaptcha.offline.domain;

public class GetResp {
    private String api;
    private int state;
    private int dt;
    private String dtkey;

    @Override
    public String toString() {
        return "GetResp{" +
                "api='" + api + '\'' +
                ", state=" + state +
                ", dt=" + dt +
                ", dtkey='" + dtkey + '\'' +
                '}';
    }

    public String getApi() {
        return api;
    }

    public void setApi(String api) {
        this.api = api;
    }

    public int getState() {
        return state;
    }

    public void setState(int state) {
        this.state = state;
    }

    public int getDt() {
        return dt;
    }

    public void setDt(int dt) {
        this.dt = dt;
    }

    public String getDtkey() {
        return dtkey;
    }

    public void setDtkey(String dtkey) {
        this.dtkey = dtkey;
    }

    public GetResp(String api, int state, int dt, String dtkey) {
        this.api = api;
        this.state = state;
        this.dt = dt;
        this.dtkey = dtkey;
    }
}
