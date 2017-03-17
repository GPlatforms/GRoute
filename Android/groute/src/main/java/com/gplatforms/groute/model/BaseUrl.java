package com.gplatforms.groute.model;

import android.support.annotation.Keep;

@Keep
public class BaseUrl {
    private String reg;
    private String url;

    public String getReg() {
        return reg;
    }

    public void setReg(String reg) {
        this.reg = reg;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }
}
