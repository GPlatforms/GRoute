package com.gplatforms.groute.sample;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import com.google.gson.reflect.TypeToken;
import com.gplatforms.groute.GRouteManager;
import com.gplatforms.groute.callback.GRouteCallBack;
import com.gplatforms.groute.model.BaseUrl;
import com.gplatforms.groute.model.GRouteData;

import java.util.List;

public class MainActivity extends AppCompatActivity {

    private Button mRequestButton;
    private TextView mResultView;

    private StringBuilder mResult = new StringBuilder();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);


        GRouteManager.addConfigUrl("http://api.dianchibbs.com/config/definition");

        mRequestButton = (Button) findViewById(R.id.request);
        mResultView = (TextView) findViewById(R.id.result);
        mRequestButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                request();
            }
        });
    }

    public void request() {
        mResultView.setText("正在请求中....");
        mResult = new StringBuilder();
        GRouteManager.getInstance().request(new GRouteCallBack() {
            @Override
            public void onError(int code, String message) {
                mResult.append("发生错误：\n\n");
                mResult.append("code: " + code + "\n\n");
                mResult.append("message:" + message);

                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        mResultView.setText(mResult);
                    }
                });
            }

            @Override
            public void onSuccess() {
                GRouteManager routeManager = GRouteManager.getInstance();

                mResult.append("JSON: " + routeManager.getJson() + "\n\n\n");

                mResult.append("BaseUrl：" + routeManager.getBaseUrl() + "\n");
                mResult.append("BaseUrl.fa：" + routeManager.getBaseUrl("fa") + "\n");

                Number count = routeManager.get("count");
                String app_id = routeManager.get("app_id");
                boolean is_check = routeManager.get("is_check");
                List<String> arr = routeManager.getList("arr");
                List<Number> arr2 = routeManager.getList("arr2");
                List<BaseUrl> baseUrls = routeManager.getList("base_url", new TypeToken<List<BaseUrl>>(){}.getType());
                Share share = routeManager.get("share", Share.class);

                mResult.append("count: " + count.intValue() + "\n");
                mResult.append("app_id: " + app_id + "\n");
                mResult.append("is_check: " + is_check + "\n");
                for (BaseUrl baseUrl : baseUrls) {
                    mResult.append("reg: " + baseUrl.getReg() + ", url:" + baseUrl.getUrl() +"\n");
                }
                mResult.append("share title: " + share.getTitle() + ", url:" + share.getUrl());
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        mResultView.setText(mResult.toString());
                    }
                });
            }
        });
    }
}
