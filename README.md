# Leetcode每日一题
获取Leetcode每日一题，可部署在腾讯云函数上，通过企业微信API每日定时推送

## 配置
**cron表达式：**
`0 0 10,14,16 * * * *`

**触发器附加信息：**
```json
{
    "notify": {
        "serverChan": {
            "secretKey": ""
        }
    }
}
```

## 参考项目
[wangwangit/python_sign](https://github.com/wangwangit/python_sign/blob/master/com/ww/leetcode/leetcode.py)