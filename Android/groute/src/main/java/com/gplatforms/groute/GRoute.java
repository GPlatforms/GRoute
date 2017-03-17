package com.gplatforms.groute;

import com.google.gson.Gson;
import com.gplatforms.groute.callback.GRouteCallBack;
import com.gplatforms.groute.config.GRouteConfig;
import com.gplatforms.groute.model.BaseUrl;
import com.gplatforms.groute.model.GRouteData;
import com.gplatforms.groute.model.GRouteModel;

import java.io.IOException;
import java.util.List;
import java.util.Map;

import okhttp3.Call;
import okhttp3.Callback;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class GRoute {

    /**
     * ====================================
     * Groute Singleton
     * ====================================
     */
    private volatile static GRoute instance;

    private GRoute() {
    }

    public static GRoute getInstance() {
        if (instance == null) {
            synchronized (GRoute.class) {
                if (instance == null) {
                    instance = new GRoute();
                }
            }
        }
        return instance;
    }

    /**
     * ====================================
     * Groute Error Code
     * ====================================
     */
    public static final int CODE_OK = 200;
    public static final int CODE_ERROR_HTTP = -1;

    /**
     * ====================================
     * Groute Support
     * ====================================
     */

    private String mGRouteJson;
    private GRouteModel mGRouteModel;
    private GRouteData mGRouteData;

    public void request(final GRouteCallBack gRouteCallBack) {
        OkHttpClient client = new OkHttpClient();
        Request request = new Request.Builder()
                .url(GRouteConfig.getUrl())
                .build();
        client.newCall(request).enqueue(new Callback() {
            @Override
            public void onFailure(Call call, IOException e) {
                gRouteCallBack.onError(CODE_ERROR_HTTP, e.getMessage());
            }

            @Override
            public void onResponse(Call call, Response response) throws IOException {
                String gRouteJson = response.body().string();
                Gson gson = new Gson();
                GRouteModel gRouteModel = gson.fromJson(gRouteJson, GRouteModel.class);
                GRouteData gRouteData = gRouteModel.getData();

                if (gRouteModel.getCode() == CODE_OK) {
                    mGRouteJson = gRouteJson;
                    mGRouteModel = gRouteModel;
                    mGRouteData = gRouteData;
                    gRouteCallBack.onSuccess(gRouteJson, gRouteModel, gRouteData);
                } else {
                    gRouteCallBack.onError(gRouteModel.getCode(), gRouteModel.getMsg());
                }
            }
        });
    }

    // get default key *
    public String getBaseUrl() {
        List<BaseUrl> list = mGRouteData.getBase_url();
        for (BaseUrl baseUrl : list) {
            if ("*".equals(baseUrl.getReg())) {
                return baseUrl.getUrl();
            }
        }
        return null;
    }

    public String getBaseUrl(String module) {
        List<BaseUrl> list = mGRouteData.getBase_url();
        for (BaseUrl baseUrl : list) {
            if (module.matches(baseUrl.getReg())) {
                return baseUrl.getUrl();
            }
        }
        return null;
    }

    public String getExtValue(String key) {
        Map<String, String> extKV = mGRouteData.getExt_kv();
        String result = extKV.get(key);
        return result;
    }
}
