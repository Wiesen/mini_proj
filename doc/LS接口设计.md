[TOC]

# 接口设计



## 创建用户

描述：第一次登录成功后创建，包括qq号，昵称以及头像。创建后信息不能修改

### 请求

POST /api/user

| 字段名    | 类型   | 备注                  |
| --------- | ------ | --------------------- |
| qq_number | int64  | qq号                  |
| name      | string | 昵称，可以重复        |
| pic       | string | 头像路径。考虑存在cos |

### 响应

| 字段名   | 类型 | 备注              |
| -------- | ---- | ----------------- |
| ret_code | Int  | 0. 成功； 1. 失败 |



## 发布心情

### 请求

POST  /api/emotion

| 字段名      | 类型   | 备注                     |
| ----------- | ------ | ------------------------ |
| content     | string | 心情内容，为空时不能评论 |
| label_id    | int    | 心情标签id               |
| strong      | int    | 强度                     |
| create_time | date   | 创建时间                 |
| visable     | int    | 1. 个人可见；2. 社区可见 |
| poster      | string | 发布人qq号               |

### 响应

| 字段名   | 类型 | 备注              |
| -------- | ---- | ----------------- |
| ret_code | int  | 0. 成功； 1. 失败 |



## 查看用户心情

### 请求

GET  /api/emotion?qq_number={qq_number}

### 响应

| 字段名       | 类型      | 备注              |
| ------------ | --------- | ----------------- |
| ret_code     | int       | 0. 成功； 1. 失败 |
| emotion_list | []emotion |                   |

emotion:

| 字段名      | 类型   | 备注                     |
| ----------- | ------ | ------------------------ |
| content     | string | 心情内容。为空时不能评论 |
| label_id    | int    | 心情标签id               |
| label_name  | string | 心情标签名               |
| strong      | int    | 强度                     |
| create_time | date   | 创建时间                 |
| poster      | int64  | 发布人qq号               |
| like_cnt    | int    | 点赞数                   |
| comment_cnt | int    | 评论数                   |

## 查看所有心情

### 请求

GET  /api/emotion

### 响应

| 字段名       | 类型      | 备注              |
| ------------ | --------- | ----------------- |
| ret_code     | int       | 0. 成功； 1. 失败 |
| emotion_list | []emotion |                   |

emotion:

| 字段名      | 类型   | 备注                     |
| ----------- | ------ | ------------------------ |
| content     | string | 心情内容。为空时不能评论 |
| label_id    | int    | 心情标签id               |
| label_name  | string | 心情标签名               |
| strong      | int    | 强度                     |
| create_time | date   | 创建时间                 |
| poster      | int64  | 发布人qq号               |
| like_cnt    | int    | 点赞数                   |
| comment_cnt | int    | 评论数                   |

 ## 查询所有指定标签心情

### 请求

GET  /api/emotion?label_id={label_id}

### 响应

| 字段名       | 类型      | 备注              |
| ------------ | --------- | ----------------- |
| ret_code     | int       | 0. 成功； 1. 失败 |
| emotion_list | []emotion |                   |

emotion:

| 字段名      | 类型   | 备注                     |
| ----------- | ------ | ------------------------ |
| content     | string | 心情内容。为空时不能评论 |
| label_id    | int    | 心情标签id               |
| label_name  | string | 心情标签名               |
| strong      | int    | 强度                     |
| create_time | date   | 创建时间                 |
| poster      | int64  | 发布人qq号               |
| like_cnt    | int    | 点赞数                   |
| comment_cnt | int    | 评论数                   |

## 发布评论

### 请求

POST /api/comment

| 字段名      | 类型   | 备注                              |
| ----------- | ------ | --------------------------------- |
| emotion_id  | int64  | 心情ID                            |
| comment     | string | 评论内容                          |
| poster      | int64  | 发布人qq号                        |
| create_time | date   | 时间                              |
| rspto       | int64  | 回复人qq号，为0时表示不回复指定人 |

### 响应

| 字段名   | 类型 | 备注              |
| -------- | ---- | ----------------- |
| ret_code | int  | 0. 成功； 1. 失败 |

## 查看心情下的评论

### 请求

GET /api/comment?emotion_id={emotion_id}

### 响应

| 字段名       | 类型      | 备注              |
| ------------ | --------- | ----------------- |
| ret_code     | int       | 0. 成功； 1. 失败 |
| comment_list | []comment |                   |

comment:

| 字段名      | 类型   | 备注                              |
| ----------- | ------ | --------------------------------- |
| emotion_id  | int64  | 心情ID                            |
| comment     | string | 评论内容                          |
| poster      | int64  | 发布人qq号                        |
| create_time | date   | 时间                              |
| rspto       | int64  | 回复人qq号，为0时表示不回复指定人 |



## 发布点赞



### 请求

POST /api/like

| 字段名      | 类型  | 备注       |
| ----------- | ----- | ---------- |
| emotion_id  | int64 | 心情id     |
| poster      | int64 | 发布人qq号 |
| create_time | Date  | 时间       |

### 响应

| 字段名   | 类型 | 备注              |
| -------- | ---- | ----------------- |
| ret_code | Int  | 0. 成功； 1. 失败 |

## 查看动态

>  动态：别人对自己心情的评论和点赞，以及心情下的回复

### 请求

GET /api/message?qq_number={qq_number}

### 响应

| 字段名     | 类型   | 备注                                 |
| ---------- | ------ | ------------------------------------ |
| ret_code   | int    | 0. 成功； 1. 失败                    |
| emotion_id | int64  | 心情id                               |
| comment    | string | 评论                                 |
| like       | int    | 0.非点赞，1.点赞。和评论不能同时为空 |
| poster     | int64  | 发布人                               |



## 表情识别

通过优图接口实现