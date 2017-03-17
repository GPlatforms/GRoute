package com.gplatforms.groute.callback;

import android.support.annotation.Keep;

import com.gplatforms.groute.model.GRouteData;
import com.gplatforms.groute.model.GRouteModel;

@Keep
public interface GRouteCallBack {
    void onError(int code, String message);
    void onSuccess(String gRouteJson, GRouteModel gRouteModel, GRouteData gRouteData);
}
