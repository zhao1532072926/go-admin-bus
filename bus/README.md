# Bus 服务模块

Bus 是一个基于配置的多维度服务生成框架，通过简单的配置即可生成完整的多维度服务架构。

## 目录结构

```
bus/
├── config/ # 配置相关
│ └── service.go # 配置加载和解析
├── controller/ # 控制器
├── middleware/ # 中间件
├── models/ # 数据模型
├── router/ # 路由
├── utils/ # 工具函数
└── init.go # 初始化入口
```


## 配置化生成服务维度
服务支持在配置文件中定义二维的服务名称，并自动生成数据表和 API 路由。
对应数据库表会在服务启动时自动检查并创建。

### 配置文件说明
配置文件路径: `config/service.yml`
配置文件格式:
```yaml
services: # 服务列表
  - name: "first" # 第一维度名称
    sub_services: # 子服务列表
      - "dd" # 子服务名称
      - "aa"
      - "bb"

jwt: # JWT 配置 
  secret: "your-secret-key" # 密钥
  timeout_second: 86400 # 24 hours
```

### 自动生成的数据表
服务启动时，会根据配置文件自动检查表是否存在，不存在则创建。

- 用户表: `{维度名称}_users`
  - 示例: first_users

- 详情表: `{维度名称}_{子服务名称}_details`
  - 示例: first_dd_details, first_aa_details

### 界面配置时，路径为 /info/{表名}
示例: /info/first_users

### API 路由

- 登录: `/api/{维度名称}/login`
- 详情: `/api/{维度名称}/{子服务名称}/details` 


