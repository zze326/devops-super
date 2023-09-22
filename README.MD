# zze Admin go

## 功能

本项目主要用来作为新项目启动时的模板，如果你想要使用 Golang 来新开发一个后台管理系统，那么选我就对了。

- 已实现后台管理系统最基础的用户管理、角色管理、权限管理，基于它来开发可以让你可以更聚焦业务的实现。
- 选用了相对成熟、稳定、功能齐全的的框架，后端基于 **[GoFrame](https://goframe.org/pages/viewpage.action?pageId=1114119)**
  ，前端基于 **[PureAdmin](https://yiming_chang.gitee.io/pure-admin-doc/pages/introduction/)**；
- 支持一套权限管理逻辑控制前后端路由、按钮级别权限；

## 预览地址

[点我预览](http://admin.zze.xyz)

- 管理员：`admin`，密码：`devops.zze`；
- 测试账号：`test`，密码：`devops.zze`；

效果图：

![用户管理](https://raw.githubusercontent.com/zze326/zze-admin-go/main/resource/imgs/user-manage.png)
![角色管理](https://raw.githubusercontent.com/zze326/zze-admin-go/main/resource/imgs/role-manage.png)
![权限管理](https://raw.githubusercontent.com/zze326/zze-admin-go/main/resource/imgs/permission-manage.png)
![部门管理](https://raw.githubusercontent.com/zze326/zze-admin-go/main/resource/imgs/dept-manage.png)

## 技术栈

- 语言：Golang、Typescript；
- 后端：GoFrame、Casbin；
- 前端：Vue3、Vite、Element-Plus、TypeScript、Pinia 等；

## 项目运行
1、在 MySQL 中执行 `manifest/db/zze_admin_go.sql` 创建好数据库以及初始化数据；

2、然后修改 `manifest/config/config.yaml` 中的数据库连接地址（`database.default.link`），格式如下：
```sql
mysql:<用户名>:<密码>@tcp(<数据库地址>)/<库名>?loc=Local&parseTime=true
-- 例：mysql:zze:zze.admin@tcp(127.0.0.1:3306)/zze_admin_go?loc=Local&parseTime=true
```
3、直接运行项目根目录的 `main.go` 就可以跑起来啦~
## 前端项目
本仓库是后端项目，对应前端项目地址为：<https://github.com/zze326/zze-admin-fe>。