[TOC]

# 接口设计

## 登录

### 请求

POST /api/login

| 字段名       | 类型   | 备注   |
| ------------ | ------ | ------ |
| phone_number | string | 手机号 |
| password     | string | 密码   |


### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     |        |                                      |

data:

| 字段名 | 类型   | 备注 |
| ------ | ------ | ---- |
| token  | string |      |

## 登出

### 请求

POST /api/logout?:token

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     |        |                                      |

data:


## 注册

### 请求

POST /api/user

| 字段名       | 类型   | 备注           |
| ------------ | ------ | -------------- |
| phone_number | string | 手机号         |
| nickname     | string | 昵称，可以重复 |
| password     | string | 密码           |

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     | []     |                                      |

## 查看用户自身信息

### 请求

GET /api/user?token={token}

### 响应

| 字段名   | 类型   | 备注                                  |
| -------- | ------ | ------------------------------------- |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                              |
| data     |        |                                       |

data:

| 字段名    | 类型   | 备注                                  |
| --------- | ------ | ------------------------------------- |
| id           | INT         |           |
| phone_number | VARCHAR(64) | 手机号    |
| nickname     | VARCHAR(64) | 昵称      |
| avatar       | VARCHAR(64) | 头像      |



## 发布心情

### 请求

POST  /api/emotion? token={token}

| 字段名   | 类型   | 备注                     |
| -------- | ------ | ------------------------ |
| content  | string | 心情内容                 |
| label_id | int    | 心情标签id               |
| strong   | int    | 强度                     |
| visiable | int    | 1. 个人可见；2. 社区可见 |
| poster   | int    | 发布人id                 |

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     | []     |                                      |



## 查看用户自身心情（“我的”界面）

### 请求

GET  /api/emotion/self?token={token} & pageno= {pageno}

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     |      |                                      |

data:

| 字段名      | 类型   | 备注       |
| ----------- | ------ | ---------- |
| emotion_id  | int    | 心情id     |
| content     | string | 心情内容   |
| label_id    | int    | 心情标签id |
| label_name  | string | 心情标签名 |
| strong      | int    | 强度       |
| create_time | date   | 创建时间   |

 ## 查询所有指定标签心情（“广场”界面）

### 请求

GET  /api/emotion?label_id={label_id}& token={token} & pageno= {pageno}

> 无label_id时查询全部

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string |                                      |
| data     | []     |                                      |

data[i]

| 字段名      | 类型   | 备注                         |
| ----------- | ------ | ---------------------------- |
| emotion_id  | int    | 心情id                       |
| content     | string | 心情内容                     |
| label_id    | int    | 心情标签id                   |
| label_name  | string | 心情标签名                   |
| strong      | int    | 强度                         |
| create_time | date   | 创建时间                     |
| poster      | int64  | 发布人id                     |
| nickname    | string | 发布人昵称                   |
| avatar      | string | url                          |
| like_cnt    | int    | 点赞数                       |
| comment_cnt | int    | 评论数                       |
| is_like     | int    | 用户是否点过赞, 1点过, 0没有 |

## 发布评论

### 请求

POST /api/comment? token={token}

| 字段名     | 类型   | 备注                            |
| ---------- | ------ | ------------------------------- |
| emotion_id | int64  | 心情ID                          |
| comment    | string | 评论内容                        |
| rspto      | int64  | 回复人id，为0时表示不回复指定人 |

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     | []     |                                      |

## 查看心情下的评论

### 请求

GET /api/comment?emotion_id={emotion_id} & token={token} & pageno= {pageno}

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     | []     |                                      |

data[i]:

| 字段名      | 类型   | 备注                            |
| ----------- | ------ | ------------------------------- |
| emotion_id  | int64  | 心情ID                          |
| comment     | string | 评论内容                        |
| poster      | int64  | 发布人id                        |
| create_time | date   | 时间                            |
| rspto       | int64  | 回复人id，为0时表示不回复指定人 |



## 发布点赞

### 请求

POST /api/like?token={token}

| 字段名     | 类型  | 备注   |
| ---------- | ----- | ------ |
| emotion_id | int64 | 心情id |

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     | []     |                                      |



## 查看动态

>  动态：别人对自己心情的评论和点赞，以及心情下的回复

### 请求

GET /api/message? token={token} & pageno= {pageno}

### 响应

| 字段名   | 类型   | 备注                                 |
| -------- | ------ | ------------------------------------ |
| ret_code | int    | 0. 成功； -1. 失败； -2 token验证失败 |
| message  | string | 错误消息                             |
| data     | []     |                                      |

data[i]:

| 字段名     | 类型   | 备注                                     |
| ---------- | ------ | ---------------------------------------- |
| emotion_id | int64  | 心情id                                   |
| comment    | string | 评论或回复内容                           |
| poster     | int64  | 评论、点赞或回复的发布人id               |
| nickname | string | 昵称 |
| avatar | string | 头像 |
| type       | int    | 1. 点赞； 2. 心情评论；3. 回复 |



## 表情识别

通过优图接口实现