# Leetcode每日一题
获取Leetcode每日一题，可部署在腾讯云函数上，通过企业微信API每日定时推送

> 重要提示：2022年6月1日之后，腾讯云对函数服务进行了收费。前三个月可以每月领取一定的免费额度，之后便需要购买了。学生可以享受一元购买一年的优惠活动。具体请以官网活动为主。

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