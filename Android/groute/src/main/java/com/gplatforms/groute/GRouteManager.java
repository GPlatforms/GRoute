package com.gplatforms.groute;

import android.util.Log;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.gplatforms.groute.callback.BaseCallBack;
import com.gplatforms.groute.model.BaseUrl;
import com.gplatforms.groute.model.GRouteData;
import com.gplatforms.groute.model.GRouteModel;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.IOException;
import java.lang.reflect.Type;
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

    public GRouteManager addConfigUrl(String url) {
        mConfigUrls.add(url);
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
    private String mGRouteJson;
    private JSONObject mGRouteJsonObject;
    private GRouteData mGRouteData;
    private OkHttpClient mHttpClient = new OkHttpClient();
    private boolean isSuccessed = false;
    private int mCallCount = 0;

    public void request(final BaseCallBack baseCallBack) {
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
                        baseCallBack.onError(CODE_ERROR_HTTP, e.getMessage());
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
                    try {
                        mGRouteJsonObject = new JSONObject(gRouteJson);
                    } catch (JSONException e) {
                        e.printStackTrace();
                    }
                    Log.d(TAG, "groute json: " + gRouteJson);

                    // model
                    Gson gson = new Gson();
                    try {
                        GRouteModel gRouteModel = gson.fromJson(gRouteJson, GRouteModel.class);
                        if (gRouteModel != null) {
                            // data
                            GRouteData gRouteData = gRouteModel.getData();
                            mGRouteData = gRouteData;

                            if (gRouteModel.getCode() == CODE_OK) {
                                if (!isSuccessed) {
                                    isSuccessed = true;
                                    mGRouteJson = gRouteJson;
                                    baseCallBack.onSuccess();
                                } else {
                                    Log.d(TAG, "discard response after request successed.");
                                }
                            } else {
                                if (!isSuccessed && mCallCount == 0) {
                                    baseCallBack.onError(gRouteModel.getCode(), gRouteModel.getMsg());
                                }
                            }
                        } else {
                            if (!isSuccessed && mCallCount == 0) {
                                baseCallBack.onError(CODE_ERROR_PARSE, "model为空");
                            }
                        }
                    } catch (Exception e) {
                        if (!isSuccessed && mCallCount == 0) {
                            baseCallBack.onError(CODE_ERROR_PARSE, e.getMessage());
                        }
                    }
                }
            });
        }
    }

    /**
     * 内置默认的BaseUrl
     *
     * @return
     */
    public String getBaseUrl() {
        List<BaseUrl> baseUrls = getList("base_url", new TypeToken<List<BaseUrl>>(){}.getType());
        for (BaseUrl baseUrl : baseUrls) {
            if ("*".equals(baseUrl.getReg())) {
                return baseUrl.getUrl();
            }
        }
        return null;
    }

    /**特定模块的BaseUrl
     *
     * @return
     */
    public String getBaseUrl(String module) {
        List<BaseUrl> baseUrls = getList("base_url", new TypeToken<List<BaseUrl>>(){}.getType());
        for (BaseUrl baseUrl : baseUrls) {
            if (module.matches(baseUrl.getReg())) {
                return baseUrl.getUrl();
            }
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

    /**
     * 获取自定义类型： Clazz
     *
     * @param key
     * @param clazz
     * @param <T>
     * @return
     */
    public <T> T get(String key, Class clazz) {
        try {
            JSONObject jsonObject = mGRouteJsonObject.getJSONObject("data").getJSONObject(key);
            T result = (T) new Gson().fromJson(jsonObject.toString(), clazz);

            return result;
        } catch (JSONException e) {
            e.printStackTrace();
        }
        return null;
    }

    /**
     * 获取基本类型列表： List<Number, Boolean, String>
     *
     * @param key
     * @param <T>
     * @return
     */
    public <T> List<T> getList(String key) {
        try {
            JSONArray jsonArray = mGRouteJsonObject.getJSONObject("data").getJSONArray(key);
            List<T> list = new Gson().fromJson(jsonArray.toString(), new TypeToken<List<T>>(){}.getType());

            return list;
        } catch (JSONException e) {
            e.printStackTrace();
        }
        return null;
    }

    /**
     * 获取自定义类型列表： List<Type>
     *
     * @param key
     * @param type
     * @param <T>
     * @return
     */
    public <T> List<T> getList(String key, Type type) {
        try {
            JSONArray jsonArray = mGRouteJsonObject.getJSONObject("data").getJSONArray(key);
            List<T> list = new Gson().fromJson(jsonArray.toString(), type);

            return list;
        } catch (JSONException e) {
            e.printStackTrace();
        }
        return null;
    }

    public String getJson() {
        return mGRouteJson;
    }

}
