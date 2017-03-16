# GRoute
Server config route platform

## 愿景
- 服务端能够灵活的配置各个参数
- 服务端能根据客户端请求属性返回对应的 JSON、支持条件的配置
- 服务器能够定时检测服务可用性，做出预警（可选）
- SDK 可配置常规基本参数：版本，平台...
- SDK 提供给 APP 方便的查询 JSON 数据的接口

## 服务器返回结果
```json
{
    "code": 200,
        "msg": "success",
        "data": {
            "base_url": [
            {
                "reg": "fa",
                "url": "http://www.baidu.com"
            },
            {
                "reg": "fa",
                "url": "http://www.baidu.com"
            },
            {
                "reg": "*",
                "url": "http://www.baidu.com"
            }
            ]
        }
}
```

## SDK 参数配置

