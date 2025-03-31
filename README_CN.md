
# data-space

[![Go Report Card](https://goreportcard.com/badge/github.com/hexiaopi/data-space)](https://goreportcard.com/report/github.com/hexiaopi/data-space)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/hexiaopi/data-space)](https://github.com/Hexiaopi/data-space)

通用的后台管理系统

## 功能

- 用户登录和日志
- 部门管理
- 用户管理
- 角色管理
- 菜单管理

## 🛠 技术栈

### 前端

- vue3
- vite
- pina
- element-plus

### 后端

- go
- mysql
- gin
- opentelemetry
- prometheus metric

## 安装

### 前端

```bash
  cd web
  npm install
```

### 后端

```bash
  go mod tidy
```
    
## 本地运行

### 前端

```bash
  cd web
  npm run dev
```

### 后端

```bash
  go run cmd/root.go
```

## 截图

![login](./docs/images/login.png)
![home](./docs/images/home.png)
![depart](./docs/images/depart.png)
![user](./docs/images/user.png)
![role](./docs/images/role.png)
![role-menu](./docs/images/role_menu.png)
![menu](./docs/images/menu.png)
![log](./docs/images/log.png)
![jaeger-all](./docs/images/jaeger_all.png)
![jaeger-detail](./docs/images/jaeger_detail.png)
![swagger](./docs/images/swagger.png)
![metrics](./docs/images/metrics.png)

## 收藏历史

[![Star History Chart](https://api.star-history.com/svg?repos=hexiaopi/data-space&type=Date)](https://star-history.com/#hexiaopi/data-space&Date)


## 许可证

[MIT](https://choosealicense.com/licenses/mit/)

