配置分发SDK - Android
=====

## 接口
GRoute提供以下方法，方便开发者：
```java
void request(fGRouteCallBack gRouteCallBack) // 请求配置
String getBaseUrl()                          // 获取可用URL
String getBaseUrl(String module)             // 获取特定模块可用URL
String getExtValue(String key)               // 获取Key-Value属性值
```

## 使用步骤
#### 初始化
```java
GRouteConfig.setUrl("http://api.dianchibbs.com/config/definition");
```
#### 请求配置
```java
GRoute.getInstance().request(new GRouteCallBack() {
        @Override
        public void onError(int code, String message) {
        }

        @Override
        public void onSuccess(String gRouteJson, GRouteModel gRouteModel, GRouteData gRouteData) {
        }
    });
```