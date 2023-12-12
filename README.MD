# Devops Super

## 功能

本项目的定位目标是成为一个轻量、易使用、易部署、功能丰富的运维管理平台，已实现以下功能：

- [x] 基础系统功能：用户管理、部门管理、角色管理、权限管理；
    - [x] 一套权限管理逻辑控制前后端路由、按钮级别权限；
- [x] 堡垒机功能：主机管理、主机组管理；
    - [x] 支持密码、秘钥方式建立 SSH Web 终端（基于 WebSocket）；
    - [x] 支持 Web 文件管理器（上传、下载、删除）管理主机文件系统（基于 WebSocket + SFTP）；
    - [x] 支持终端会话记录持久化到硬盘、回放会话；
    - [x] 支持主机访问授权（主机所属主机组授权到角色或用户）；
- [x] 基于 Kubernetes 原生 Pod 的持续集成（CI）；
    - [x] 秘钥管理；
    - [x] 构建环境管理；
    - [x] 流水线管理（编排）；
    - [x] 运行流水线 & 构建流水线客户端镜像；
    - [x] 流水线运行历史展示；
    - [x] 获取流水线运行日志；
    - [x] 支持终止执行中的构建任务；
    - [x] 流水线参数化支持；
    - [x] 支持克隆流水线；
    - [ ] 支持挂载多种类型的卷以持久化构建数据；
    - [ ] 支持 kaniko 构建上传镜像

## 预览地址

[点我预览](http://ds-demo.zze.xyz)

- 管理员：`admin`，密码：`devops.zze`；
- 测试账号：`test`，密码：`devops.zze`；

> 数据库每小时自动重置。

效果图：

![用户管理](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/user-manage.png)
![角色管理](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/role-manage.png)
![权限管理](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/permission-manage.png)
![部门管理](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/dept-manage.png)
![主机管理](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/host-manage.png)
![主机组授权](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/host-group-authorize.png)
![主机终端记录](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/host-terminal-session-history.png)
![全屏终端](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/terminal-single.png)
![终端文件管理器](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/terminal-file-manager.png)
![终端主机树](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/terminal-tree.png)
![流水线管理](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/pipeline-manage.png)
![流水线编排](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/pipeline-arrange.png)
![流水线日志](https://raw.githubusercontent.com/zze326/devops-super/main/resource/imgs/pipeline-log.png)

## 技术栈

- 语言：Golang、Typescript；
- 后端：Go Frame、Casbin、Pongo2、Kubernetes Client Go、Go Git；
- 前端：Vue3、Vite、Element-Plus、TypeScript、Pinia、Xterm 等；

## 项目运行

1、在 MySQL 中执行 `manifest/db/devops_super.sql` 创建好数据库以及初始化数据；

2、然后修改 `manifest/config/config.yaml` 中的数据库连接地址（`database.default.link`），格式如下：

```sql
mysql
:<用户名>:<密码>@tcp(<数据库地址>)/<库名>?loc=Local&parseTime=true
-- 例：mysql:zze:zze.admin@tcp(127.0.0.1:3306)/devops_super?loc=Local&parseTime=true
```

3、直接运行项目根目录的 `main.go` 就可以跑起来啦~

## 前端项目

本仓库是后端项目，对应前端项目地址为：<https://github.com/zze326/devops-super-fe>。

## 快速部署

1、在 MySQL 中执行 `manifest/db/devops_super.sql` 创建好数据库以及初始化数据；

2、从 [releases](https://github.com/zze326/devops-super/releases)
获取到要部署的镜像版本，以 [v1.0](https://github.com/zze326/devops-super/releases/tag/v1.0) 为例：主程序 Docker
镜像为 `registry.cn-shenzhen.aliyuncs.com/zze/devops-super:202310271518`；

3、运行容器，看如下实例：

```bash
docker run -d --name devops-super \
        -v/opt/devops-super/host-sessions:/app/host-sessions \
        -eSERVER_PORT=8001 \
        -eDB_HOST=127.0.0.1 \
        -eDB_PORT=3306 \
        -eDB_NAME=devops_super \
        -eDB_USER=zze \
        -eDB_PWD=zze.admin \
        -eJWT_SECRET=vIIEngfamdsaGZasdsasdasadkseadgF9fe \
        --network host registry.cn-shenzhen.aliyuncs.com/zze/devops-super:202310271518
```

> 容器的 `/app/host-sessions` 目录保存了主机终端会话记录信息，可以将它挂载出来以持久化保存。
> 
> 环境变量说明：
> - `DB_HOST`：MySQL 主机地址；
> - `DB_PORT`：MySQL 端口；
> - `DB_NAME`：数据库名称；
> - `DB_USER`：连接数据库使用的用户；
> - `DB_PWD`：连接数据库的用户密码；
> - `JWT_SECRET`：JWT 秘钥，随机生成即可；
> - `SERVER_PORT`：服务监听端口；
>
> 参数说明（按需使用）：
> - `-d`：后台运行容器；
> - `--name`：指定容器名称；
> - `--network host`：容器共用宿主机网络空间；

4、容器运行成功后浏览器访问 8001 端口即可；