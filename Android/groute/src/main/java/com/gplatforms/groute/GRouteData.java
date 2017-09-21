package com.gplatforms.groute;

import android.support.annotation.Keep;

import java.util.HashMap;
import java.util.List;

@Keep
public class GRouteData extends HashMap<String, Object> {
    public int getCode() {
        return ((Number) get("code")).intValue();
    }

    public String getMsg() {
        return (String) get("msg");
    }

    public List<String> getBase_url() {
        List<String> baseUrl = (List<String>) get("base_url");
        return baseUrl;
    }
}
