# golib

### 安装
```shell
go get github.com/shrimps80/golib
```

### 使用方法
+ 获取当前时间戳 `Time()`
+ 时间戳转时间格式 `Date("Y-m-d H:i:s", time)`
+ 时间格式转时间戳 `Strtotime("2012/12/12 12:12:12")`
+ 截取字符串 `Substr(str, 3, -1)`
+ 翻转字符串 `Strrev(str)`
+ 数组中是否存在指定的值 `InArray(3, number)`
+ 查找字符串在另一字符串中第一次出现的位置 `Strpos("You love php, I love php too!", "p1hp", 10)`
+ 根据map值进行排序 `SortMapByValue(age, "DESC")`
+ 用于小程序解码 `AesCBCDecrypt(session_key, iv)`
+ 字符转数字 `ParseInt("10", 0)`
+ 获取随机字符串 `GetRandomString(6)`
+ sha1加密 `Sha1Str("123456")`
+ md5加密 `Md5Str("123456")`
+ 四舍五入保留n位小数 `Round(3.164, 1)`
+ 四舍保留n位小数 `TruncRound(3.164, 1)`