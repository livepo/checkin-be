## 1. 签到用户列表
GET /api/v1/checkin-users

response
```json
{
    "code": 0,
    "data": [
        {
            "avatar": "https://www.baidu1.com",
            "id": 1,
            "poster": "https://www.jd.com111",
            "username": "1212"
        },
        {
            "avatar": "https://www.baidu1.com",
            "id": 2,
            "poster": "https://www.jd.com111",
            "username": "李四"
        },
        {
            "avatar": "https://www.baidu1.com",
            "id": 3,
            "poster": "https://www.jd.com111",
            "username": "AAA"
        }
    ],
    "msg": "success"
}
```

## 2. 签到
POST /api/v1/checkin

json:
```json
{
    "username": "大人物3",
    "avatar": "https://www.baidu1.com",
    "poster": "https://www.jd.com111"
}
```

response
```json
{
    "code": 0,
    "msg": "success"
}
```

## 3. 删除有问题的签到数据
DELETE /api/v1//checkin-users/:userid

response
```json
{
    "code": 0,
    "msg": "success"
}
```

## 4. 生成交换贺卡列表
POST /api/v1//generate-exchange-card-pairs

response
```json
{
    "code": 0,
    "msg": "success"
}
```

## 5. 列出交换贺卡列表
GET /api/v1/exchange-card-pairs

response
```json
{
    "code": 0,
    "data": [
        {
            "recv_id": 5,
            "send_id": 8
        },
        {
            "recv_id": 8,
            "send_id": 7
        },
        {
            "recv_id": 7,
            "send_id": 3
        }
    ],
    "msg": "success"
}
```

## 6. 抽奖设置
POST /api/v1/lottery-setup

JSON:
```json
{
    "label": "三等奖",
    "description": "iphone, macbookair...",
    "amount": 5
}
```
response
```json
{
    "code": 0,
    "msg": "success"
}
```

## 7. 点击生成抽奖
POST /api/v1/generate-lottery

JSON:
```json
{
    "label": "三等奖"
}
```

response
```json
{
    "code": 0,
    "msg": "success"
}
```

## 8. 抽奖结果
GET /api/v1/lotteries?label=一等奖

response
```json
{
    "code": 0,
    "data": [
        {
            "label": "一等奖",
            "user_id": 2
        },
        {
            "label": "一等奖",
            "user_id": 4
        },
        {
            "label": "一等奖",
            "user_id": 9
        }
    ],
    "msg": "success"
}
```