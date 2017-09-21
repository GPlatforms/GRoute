package com.gplatforms.groute;

import android.support.annotation.Keep;

@Keep
public interface GRouteCallBack {
    void onError(int code, String message);
    void onSuccess();
}