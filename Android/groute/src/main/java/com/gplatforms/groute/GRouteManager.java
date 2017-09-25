package com.gplatforms.groute;

import android.content.Context;
import android.text.TextUtils;
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

    private Context mContext;
    private String mAppId;
    private String mSecret;
    private static HashSet<String> mConfigUrls = new HashSet<>();

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

    public GRouteManager setContext(Context context) {
        this.mContext = context;
        return this;
    }

    public GRouteManager setAppId(String appId) {
        this.mAppId = appId;
        return this;
    }

    public GRouteManager setSecret(String secret) {
        this.mSecret = secret;
        return this;
    }

    public GRouteManager setConfigUrl(List<String> configUrls) {
        if (configUrls != null) {
            for (String configUrl : configUrls) {
                long currentTime = System.currentTimeMillis() / 1000;
                String sign = GRouteUtil.sha1(mSecret + mAppId + currentTime);
                String result = configUrl + "?app_id=" + mAppId + "&timestamp=" + currentTime + "&sign=" + sign;

                Log.d(TAG, "config url: " + result);

                mConfigUrls.add(result);
            }
        }

        return this;
    }

    public void build() {

        if (mContext == null) {
            throw new RuntimeException("please init groute context first.");
        }

        if (TextUtils.isEmpty(mAppId)) {
            throw new RuntimeException("please init groute appid first.");
        }

        if (TextUtils.isEmpty(mSecret)) {
            throw new RuntimeException("please init groute secret first.");
        }

        if (mConfigUrls.size() == 0) {
            throw new RuntimeException("please init groute config url first.");
        }

    }

    public void update() {
        update(null);
    }

    public void update(final GRouteCallBack callback) {
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
                        if (callback != null) {
                            callback.onError(CODE_ERROR_HTTP, e.getMessage());
                        }
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
                                GRouteUtil.writeCache(mContext, gRouteJson);
                                if (callback != null) {
                                    callback.onSuccess();
                                }
                            } else {
                                Log.d(TAG, "discard response after request successed.");
                            }
                        } else {
                            if (!isSuccessed && mCallCount == 0) {
                                if (callback != null) {
                                    callback.onError(CODE_ERROR_PARSE, "model为空");
                                }
                            }
                        }
                    } catch (Exception e) {
                        mCallCount--;
                        if (!isSuccessed && mCallCount == 0) {
                            if (callback != null) {
                                callback.onError(CODE_ERROR_PARSE, e.getMessage());
                            }
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
        ensureGrouteData();
        return mGRouteData.getCode();
    }

    public String getMsg() {
        ensureGrouteData();
        return mGRouteData.getMsg();
    }

    public String getBaseUrl() {
        ensureGrouteData();
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
        ensureGrouteData();
        T result = (T) mGRouteData.get(key);
        return result;
    }

    public boolean isAvaliable() {
        ensureGrouteData();
        return mGRouteData != null;
    }

    private void ensureGrouteData() {
        if (mGRouteData == null) {
            String cache = GRouteUtil.readCache(mContext);
            if (!TextUtils.isEmpty(cache)) {
                mGRouteData = new Gson().fromJson(cache, GRouteData.class);
            }
        }
    }

}
