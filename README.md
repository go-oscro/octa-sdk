# README #

这是基于微服务设计的 OSCRO 平台中使用的SDK


### 为什么要使用此sdk

- 数据库表结构统一设计管理
- 公共插件统一使用及处理

### 数据库表结构

平台中所有的微服务均使用src/models/xx. 数据表结构是在此设计，数据库迁移时由另外一些程序或者开发者来提供

### 公共插件

所有插件均由当前项目即oscro平台中使用的，共用的组件，目的就是为了方便统一进行配置使用

### 如何使用

在其他对应服务中进行引用即可