配置分发SDK - Android
=====

下载：[groute-0.2.jar](release/groute-0.2.jar)

## 接口
GRoute提供以下方法请求配置并获取KV配置：

```java

void request(GRouteCallBack gRouteCallBack)  // 请求配置
T get(String key)                            // 获取基本类型： Number, Boolean, String, List<String>等

```

每个应用都需要配置BaseUrl(支持分模块)，为了方便开发者，我们额外提供了两个方法：

```java

String getBaseUrl()                          // 获取可用URL
String getBaseUrl(String module)             // 获取特定模块可用URL

```

## 使用步骤
#### 初始化

```java

GRouteManager.getInstance()
    .addConfigUrl("http://api.dianchibbs.com/config/definition")
    .addConfigUrl("http://api.dianchibbs.com/config/definition2")
    .addConfigUrl("http://api.dianchibbs.com/config/definition3");

```

#### 请求配置

```java

GRouteManager.getInstance().request(new GRouteCallBack() {
        @Override
        public void onError(int code, String message) {
            mResult.append("发生错误：\n\n");
            mResult.append("code: " + code + "\n\n");
            mResult.append("message:" + message);
        }

        @Override
        public void onSuccess() {
            GRouteManager routeManager = GRouteManager.getInstance();

            Number count = routeManager.get("count");
            String app_id = routeManager.get("app_id");
            boolean is_check = routeManager.get("is_check");
            List<String> arr = routeManager.getList("arr");
            List<Number> arr2 = routeManager.getList("arr2");
            List<BaseUrl> baseUrls = routeManager.getList("base_url",
                                        new TypeToken<List<BaseUrl>>(){}.getType());
            Share share = routeManager.get("share", Share.class);
        }
});

```
