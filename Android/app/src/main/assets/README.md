配置分发SDK - Android
=====

下载：[groute-0.3.jar](release/groute-0.3.jar)

## 接口
GRoute提供以下方法请求配置并获取KV配置：

```

    boolean isAvaliable()                        // 本地是否有配置数据
    void update()                                // 请求配置
    void update(GRouteCallBack callback)         // 请求配置
    T get(String key)                            // 获取基本类型： Number, Boolean, String, List<String>等

```

为了方便开发者，特别提供了获取 Server Url 方法：

```

    String getBaseUrl()                          // 获取可用 Server URL

```

## 使用步骤
#### 初始化

```

    String appId = "11";
    String secretKey = "8e";
    List<String> configUrls = new ArrayList<>();
    configUrls.add("http://111.111.111.111/groute/v1/config");
    configUrls.add("http://222.222.222.2222/groute/v1/config");
    GRouteManager.getInstance()
            .setContext(context)
            .setAppId(appId)
            .setSecret(secretKey)
            .setConfigUrl(configUrls)
            .build();

```

#### 使用配置

```

    GRouteManager routeManager = GRouteManager.getInstance();
    if (routeManager.isAvaliable()) {
        int code = routeManager.getCode();
        String msg = routeManager.getMsg();
        String baseUrl = routeManager.getBaseUrl();
        boolean is_vip = routeManager.get("is_vip");
    }

```

#### 请求配置

```

    GRouteManager.getInstance().update();

```

也可以设置回调：


```

    GRouteManager.getInstance().update(new GRouteCallBack() {
        @Override
        public void onError(int code, String message) {
            mResult.append("发生错误：\n\n");
            mResult.append("code: " + code + "\n\n");
            mResult.append("message:" + message);
        }

        @Override
        public void onSuccess() {
            GRouteManager routeManager = GRouteManager.getInstance();

            int code = routeManager.getCode();
            String msg = routeManager.getMsg();
            String baseUrl = routeManager.getBaseUrl();
            boolean is_vip = routeManager.get("is_vip");
        }
    });

```
