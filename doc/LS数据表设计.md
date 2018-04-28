# 数据表设计

## 用户信息表

| 字段名    | 类型        | 备注 |
| --------- | ----------- | ---- |
| qq_number | BIGINT(20)  | qq号 |
| name      | VARCHAR(64) | 昵称 |
| pic       | VARCHAR(64) | 头像 |

## 心情表

| 字段名      | 类型        | 备注                       |
| ----------- | ----------- | -------------------------- |
| id          | BIGINT(20)  |                            |
| content     | VARCHAR(64) | 心情内容                   |
| label_id    | TINYINT(4)  | 心情标签ID，需存在标签表中 |
| strong      | TINYINT(4)  | 强度                       |
| create_time | DATE        | 创建时间                   |
| visable     | TINYINT(4)  | 1. 个人可见；2. 社区可见   |
| poster      | BIGINT(20)  | 发布人qq号                 |





## 评论表 

| 字段名      | 类型        | 备注       |
| ----------- | ----------- | ---------- |
| id          | BIGINT(20)  |            |
| emotion_id  | BIGINT(20)  | 心情ID     |
| comment     | VARCHAR(64) | 评论内容   |
| poster      | BIGINT(20)  | 发布人qq号 |
| create_time | DATE        | 时间       |
| rspto       | BIGINT(20)  | 回复人qq号 |

## 点赞表

| 字段名      | 类型       | 备注       |
| ----------- | ---------- | ---------- |
| id          | BIGINT(20) |            |
| emotion_id  | BIGINT(20) | 心情id     |
| poster      | BIGINT(20) | 发布人qq号 |
| create_time | DATE       | 时间       |

## 标签表

| 字段名 | 类型        | 备注                 |
| ------ | ----------- | -------------------- |
| id     | TINYINT(4)  |                      |
| name   | VARCHAR(64) | 心情标签名，不可重复 |