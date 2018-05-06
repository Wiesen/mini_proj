# 数据表设计

## 用户信息表(user)

| 字段名       | 类型        | 备注      |
| ------------ | ----------- | --------- |
| id           | INT         |           |
| qq_number    | VARCHAR(64) | qq号      |
| phone_number | VARCHAR(64) | 手机号    |
| password     | VARCHAR(64) | 密码      |
| nickname     | VARCHAR(64) | 昵称      |
| token        | VARCHAR(64) | 登录token |
| avatar       | VARCHAR(64) | 头像      |

## 心情表(emotion)

| 字段名      | 类型         | 备注                       |
| ----------- | ------------ | -------------------------- |
| id          | INT          |                            |
| content     | VARCHAR(256) | 心情内容                   |
| label_id    | TINYINT(4)   | 心情标签ID，需存在标签表中 |
| strong      | TINYINT(4)   | 强度                       |
| create_time | DATE         | 创建时间                   |
| visiable    | TINYINT(4)   | 1. 个人可见；2. 社区可见   |
| poster      | INT          | 发布人id                   |
| comment_cnt | int          | 评论数量                   |
| like_cnt    | int          | 点赞数量                   |

## 评论表(comment) 

| 字段名      | 类型         | 备注       |
| ----------- | ------------ | ---------- |
| id          | INT          |            |
| emotion_id  | INT          | 心情ID     |
| content     | VARCHAR(256) | 评论内容   |
| poster      | INT          | 发布人id   |
| create_time | DATE         | 时间       |
| rspto       | INT          | 被回复人id |

## 点赞表(like)

| 字段名      | 类型 | 备注     |
| ----------- | ---- | -------- |
| id          | INT  |          |
| emotion_id  | INT  | 心情id   |
| poster      | INT  | 发布人id |
| create_time | DATE | 时间     |

## 标签表(label)

| 字段名 | 类型        | 备注                 |
| ------ | ----------- | -------------------- |
| id     | TINYINT(4)  |                      |
| name   | VARCHAR(64) | 心情标签名，不可重复 |

## 消息表(message)

| 字段名      | 类型 | 备注     |
| ----------- | ---- | -------- |
| id          | INT  |          |
| emotion_id  | INT  | 心情id   |
| poster      | INT  | 发布人id |
| owner      | INT  | 被通知人id |
| create_time | DATE | 时间     |
| content     | VARCHAR(256) | 评论内容   |
| type_id     | INT   | 消息类型，1. 点赞； 2. 心情评论；3. 回复 |
