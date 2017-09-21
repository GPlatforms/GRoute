package com.gplatforms.groute;

import android.util.Log;

import com.google.gson.Gson;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

import okhttp3.Call;
import okhttp3.Callback;
import okhttp3.Headers;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class GRouteManager {

    public static final String TAG = "GRoute";

    /**
     * ====================================
     * Groute Singleton
     * ====================================
     */
    private volatile static GRouteManager instance;

    private GRouteManager() {
    }

    public static GRouteManager getInstance() {
        if (instance == null) {
            synchronized (GRouteManager.class) {
                if (instance == null) {
                    instance = new GRouteManager();
                }
            }
        }
        return instance;
    }

    /**
     * ====================================
     * Config Url
     * ====================================
     */
    private static HashSet<String> mConfigUrls = new HashSet<>();

    public GRouteManager addConfigUrl(String url, String appId, String secret) {

        long currentTime = System.currentTimeMillis() / 1000;
        String sign = GRouteUtil.sha1(secret + appId + currentTime);
        String result = url + "?app_id=" + appId + "&timestamp=" + currentTime + "&sign=" + sign;

        Log.d(TAG, "-------------------result: " + result);

        mConfigUrls.add(result);
        return this;
    }

    /**
     * ====================================
     * Groute Error Code
     * ====================================
     */
    public static final int CODE_OK = 200;
    public static final int CODE_ERROR_HTTP = -1;
    public static final int CODE_ERROR_PARSE = -2;
    public static final int CODE_ERROR_SIGN = -3;

    /**
     * ====================================
     * Groute Support
     * ====================================
     */
    private GRouteData mGRouteData;
    private OkHttpClient mHttpClient = new OkHttpClient();
    private boolean isSuccessed = false;
    private int mCallCount = 0;

    public void request(final GRouteCallBack callBack) {
        if (mHttpClient.dispatcher().runningCallsCount() > 0) {
            Log.d(TAG, "request is running, please wait it finished...");
            return;
        }
        final List<Call> calls = new ArrayList<>();
        for (String url : mConfigUrls) {
            Request request = new Request.Builder()
                    .url(url)
                    .build();
            Call call = mHttpClient.newCall(request);
            calls.add(call);
        }

        isSuccessed = false;
        mCallCount = calls.size();
        for (Call call : calls) {
            call.enqueue(new Callback() {
                @Override
                public void onFailure(Call call, IOException e) {
                    mCallCount--;
                    if (!isSuccessed && mCallCount == 0) {
                        callBack.onError(CODE_ERROR_HTTP, e.getMessage());
                    }
                }

                @Override
                public void onResponse(Call call, Response response) throws IOException {
                    // 处理响应结果
                    Headers headers = response.headers();
                    for (String name : headers.names()) {
                        Log.d(TAG, name + ": " + headers.get(name));
                    }

                    // json
                    String gRouteJson = response.body().string();
                    Log.d(TAG, "groute json: " + gRouteJson);

                    // model
                    Gson gson = new Gson();
                    try {
                        GRouteData gRouteData = gson.fromJson(gRouteJson, GRouteData.class);
                        if (gRouteData != null && gRouteData.getCode() == CODE_OK) {
                            mGRouteData = gRouteData;
                            if (!isSuccessed) {
                                isSuccessed = true;
                                callBack.onSuccess();
                            } else {
                                Log.d(TAG, "discard response after request successed.");
                            }
                        } else {
                            if (!isSuccessed && mCallCount == 0) {
                                callBack.onError(CODE_ERROR_PARSE, "model为空");
                            }
                        }
                    } catch (Exception e) {
                        if (!isSuccessed && mCallCount == 0) {
                            callBack.onError(CODE_ERROR_PARSE, e.getMessage());
                        }
                    }
                }
            });
        }
    }

    /**
     * 内置默认的code, msg, base_url
     *
     * @return
     */
    public int getCode() {
        return mGRouteData.getCode();
    }

    public String getMsg() {
        return mGRouteData.getMsg();
    }

    public String getBaseUrl() {
        List<String> baseUrls = mGRouteData.getBase_url();
        if (baseUrls != null && baseUrls.size() > 0) {
            return baseUrls.get(0);
        }
        return null;
    }


    /**
     * 获取基本类型： Number, Boolean, String
     *
     * @param key
     * @param <T>
     * @return
     */
    public <T> T get(String key) {
        T result = (T) mGRouteData.get(key);
        return result;
    }

}
