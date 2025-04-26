Tencent TTS for Legado Reader
===============================

基于腾讯语音合成API的[Legado(开源阅读)](https://github.com/gedoor/legado)TTS服务. 本项目实现了如下功能

1. 基于Legado的协议和腾讯云的语音合成API实现TTS功能
2. 输入的文字超长时(大于150字), 自动分割文字并且并发请求腾讯云接口, 并在内存中合并最终的音频文件
3. 支持调整朗读速度


基本配置
-----------------

### 服务配置

在当前目录创建`config.json`文件, 输入如下的JSON, 并填入腾讯云的`SecretId`和`SecretKey`.


```
{
    "secretId": "",
    "secretKey": "",
    "region": ""
}
```

可以在[腾讯云API秘钥管理](https://console.cloud.tencent.com/cam/capi)页面创建或查看秘钥信息. 

> 注意: 腾讯云推荐通过创建子账号的模式使用API秘钥. 即创建一个和当前账号关联的子账号, 并授予子账号最少的必要权限. 从而避免API秘钥泄露产生重大安全风险.



### Legado配置

在Legado的朗读功能的设置页面上新增一个朗读引擎, 在URL部分输入如下内容

```
http://192.168.1.8:8000/tts,
{
    "method": "POST",
    "body": {
        "chat_name": "501000",
        "text": "{{speakText}}",
        "speed": "{{speakSpeed}}"
    }
}
```

其中`chat_name`是一个可调整的参数, 用于决定朗读的音色, 详细取值可参考腾讯云的[音色列表](https://cloud.tencent.com/document/product/1073/92668)

> 注意: 以上配置中的IP地址(192.168.1.8)需要替换为部署本服务的机器的IP地址. 


服务启动
-------------

本项目纯Go实现, 因此仅需要编译并启动项目即可

```
go build && ./legado-tts-tencent
```

或者可以直接从Release页面下载已经编译好的可执行文件.



腾讯云相关入口
-----------------

以下是一些腾讯云相关的入口, 以便于快速体验不同音色效果和查看使用费用

- [控制台体验入口](https://console.cloud.tencent.com/tts/complexaudio)
- [用量查询](https://console.cloud.tencent.com/tts)
- [费用列表](https://cloud.tencent.com/document/product/1073/34112)

参考项目
----------

感谢[freefrank/tts](https://github.com/freefrank/tts)项目提供的思路, 虽然实现上完全不同, 但这个项目让我确定我的这个项目在理论上是可行的.
