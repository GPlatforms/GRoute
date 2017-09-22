package com.gplatforms.groute.sample;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import com.gplatforms.groute.GRouteManager;
import com.gplatforms.groute.GRouteCallBack;

import java.util.ArrayList;
import java.util.List;

import us.feras.mdv.MarkdownView;

public class MainActivity extends AppCompatActivity {

    private TextView mCacheView;
    private Button mRequestButton;
    private TextView mResultView;
    private MarkdownView markdownView;

    private StringBuilder mResult = new StringBuilder();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);


        String appId = "11";
        String secretKey = "8e";
        List<String> configUrls = new ArrayList<>();
        configUrls.add("http://111.111.111.111/groute/v1/config");
        configUrls.add("http://222.222.222.2222/groute/v1/config");
        GRouteManager.getInstance()
                .setContext(this)
                .setAppId(appId)
                .setSecret(secretKey)
                .setConfigUrl(configUrls)
                .build();

        mCacheView = (TextView) findViewById(R.id.cache);
        mRequestButton = (Button) findViewById(R.id.request);
        mResultView = (TextView) findViewById(R.id.result);
        markdownView = (MarkdownView) findViewById(R.id.markdown);

        mRequestButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                request();
            }
        });
        showCache();
        setMarkdownText();
    }

    private void showCache() {
        GRouteManager routeManager = GRouteManager.getInstance();
        StringBuilder mCache = new StringBuilder();
        mCache.append("isAvaliable：" + routeManager.isAvaliable() + "\n");
        if (routeManager.isAvaliable()) {
            mCache.append("- - - - - - - - - - - - - - - -\n");
            mCache.append("code：" + routeManager.getCode() + "\n");
            mCache.append("msg：" + routeManager.getMsg() + "\n");
            mCache.append("base_url：" + routeManager.getBaseUrl() + "\n");
            boolean is_vip = routeManager.get("is_vip");
            mCache.append("is_vip: " + is_vip + "\n");
        }
        mCacheView.setText(mCache);
    }

    public void request() {
        mResultView.setText("正在请求中....");
        mResult = new StringBuilder();
        GRouteManager.getInstance().update(new GRouteCallBack() {
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
                mResult.append("请求成功：\n\n");
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        mResultView.setText(mResult.toString());
                    }
                });

                GRouteManager routeManager = GRouteManager.getInstance();
                mResult.append("code：" + routeManager.getCode() + "\n");
                mResult.append("msg：" + routeManager.getMsg() + "\n");
                mResult.append("base_url：" + routeManager.getBaseUrl() + "\n");
                boolean is_vip = routeManager.get("is_vip");
                mResult.append("is_vip: " + is_vip + "\n");
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        mResultView.setText(mResult.toString());
                        showCache();
                    }
                });
            }
        });
    }

    private void setMarkdownText() {
        markdownView.loadMarkdownFile("file:///android_asset/README.md");
    }
}
