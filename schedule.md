# schedule


原则：

- Restful 风格的API不使用第三方库、GraphQL 风格的API 使用第三方库
- 源码梳理


## 1. 基本使用

- 启动的三种方式
    - http.Handle 、http.HandleFunc
    - struct、ServerHttp
    - ServerMux、ServerHttp
- 响应写入的四种方式
    - Fmt.Fprintf
    - io.WriteString
    - W.Write
        - bytes.Buffer
        - []bytes
        - strings.Builder
    - template.Execute
- 请求
    - URL
    - Method
    - Params
    - Body
        - strings.NewReader
        - bytes.NewBuffer


为什么可以这么操作？

- 接口
- 实现
- 源码分享


## 2. 模版

- 使用
- 继承/复用
- 管道
- 循环
- 判断
- 函数

## 3. 调整项目结构

- mvc
- project-layout

## 4. database

模型建立：
    - 微信朋友圈后端数据模型设计
        - 朋友
        - 聊天组
        - 个人信息等
        - 朋友圈等

- 原生
- gorm
- xorm

## 5. 前端模版

## 6.




## 第二部

> gin, go-restful, echo , iris, beego

