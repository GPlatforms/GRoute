package com.gplatforms.groute.sample;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import com.gplatforms.groute.GRoute;
import com.gplatforms.groute.callback.GRouteCallBack;
import com.gplatforms.groute.config.GRouteConfig;
import com.gplatforms.groute.model.GRouteData;
import com.gplatforms.groute.model.GRouteModel;

public class MainActivity extends AppCompatActivity {

    private Button mRequestButton;
    private TextView mResultView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);


        GRouteConfig.setUrl("http://api.dianchibbs.com/config/definition");

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
        GRoute.getInstance().request(new GRouteCallBack() {
            @Override
            public void onError(int code, String message) {
                final StringBuilder result = new StringBuilder("请求失败：\n\n");
                result.append("code: " + code + "\n\n");
                result.append("message:" + message);

                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        mResultView.setText(result);
                    }
                });
            }

            @Override
            public void onSuccess(String gRouteJson, GRouteModel gRouteModel, GRouteData gRouteData) {
                final StringBuilder result = new StringBuilder("请求成功：\n\n");
                result.append("GRouteModel 解析成功：" + gRouteModel.toString() + "\n\n");
                result.append("GRouteData 解析成功：" + gRouteData.toString() + "\n\n");
                result.append("JSON 数据结果：" + gRouteJson + "\n\n");

                result.append("BaseUrl：" + GRoute.getInstance().getBaseUrl() + "\n");
                result.append("BaseUrl.fa：" + GRoute.getInstance().getBaseUrl("fa") + "\n");

                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        mResultView.setText(result.toString());
                    }
                });
            }
        });
    }
}
