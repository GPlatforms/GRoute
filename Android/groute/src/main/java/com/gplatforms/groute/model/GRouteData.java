package com.gplatforms.groute.model;

import android.support.annotation.Keep;

import java.util.List;
import java.util.Map;

@Keep
public class GRouteData {

    private List<BaseUrl> base_url;
    private Map<String, String> ext_kv;

    public List<BaseUrl> getBase_url() {
        return base_url;
    }

    public void setBase_url(List<BaseUrl> base_url) {
        this.base_url = base_url;
    }

    public Map<String, String> getExt_kv() {
        return ext_kv;
    }

    public void setExt_kv(Map<String, String> ext_kv) {
        this.ext_kv = ext_kv;
    }
}
