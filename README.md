# GRoute
Server config route platform

## Feature List
- 服务端能够灵活的配置各个参数
- 服务端能根据客户端请求属性返回对应的 JSON、支持条件的配置
- 服务器能够定时检测服务可用性，做出预警（可选）
- SDK 可配置常规基本参数：版本，平台...
- SDK 提供给 APP 方便的查询 JSON 数据的接口
- SDK 支持缓存设置
- SDK 支持多源配置，保证配置服务器的高可用性

## Config Json
```json
{
    "code": 200,
     "msg": "success",
     "data": {
         "app_id": "aaaa",
         "share": {
             "title": "分享",
             "url": "http://www.taobao.com"
         },
         "base_url": [
         {
             "reg": "fa",
             "url": "http://www.baidu.com"
         },
         {
             "reg": "fa",
             "url": "http://www.163.com"
         },
         {
             "reg": "*",
             "url": "http://www.sina.com"
         }
         ],
         "arr": [
             "arr1",
         "arr2"
         ],
         "arr2": [
             1,
         2
         ],
         "is_check": true,
         "count": 10
     }
}
```

## Sign
sign = sha1({app_id}data{secret})
app_id+数据+密钥

## Usage

Android: [详情](Android/README.md)

iOS: [详情](iOS/README.md)

## License

```
Copyright (C) GRoute Open Source Project

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
