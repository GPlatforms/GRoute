package com.gplatforms.groute.callback;

import android.support.annotation.Keep;

@Keep
public interface BaseCallBack {
    void onError(int code, String message);
    void onSuccess();
}
