# Hertz-Vue-Admin 项目功能详解

## 项目概述

Hertz-Vue-Admin 是一个基于 **Go (Hertz) + Vue3 + MySQL + Redis** 的全栈管理后台系统，支持动态路由、权限管理、代码生成器等高级功能。项目采用前后端分离架构，前端使用 Vue3 + Element Plus，后端使用字节跳动 Hertz 框架。

## 技术栈

### 前端
- **Vue 3.4.x**: 主框架
- **Vite 5.4.x**: 构建工具
- **Element Plus 2.8.x**: UI 组件库
- **Pinia 2.0.x**: 状态管理
- **Vue Router 4.3.x**: 路由管理
- **Axios 1.7.x**: HTTP 客户端
- **ECharts 5.3.2**: 图表库

### 后端
- **Go 1.18+**: 主语言
- **Hertz**: 高性能 HTTP 框架
- **GORM**: ORM 框架
- **MySQL**: 主数据库
- **Redis**: 缓存/会话存储
- **JWT**: 用户认证
- **Casbin**: 权限控制

---

## 核心功能模块

### 1. 用户认证与授权系统

#### 1.1 用户登录认证
**功能描述**: 支持用户名密码登录，集成验证码验证
**实现方式**:
- 前端: `web/src/view/login/index.vue` - 登录表单
- 后端: `server/api/v1/system/user_api.go#Login` - 登录接口
- 流程: 验证码校验 → 用户密码验证 → JWT Token 生成 → 返回用户信息

#### 1.2 JWT 令牌管理
**功能描述**: 基于 JWT 的无状态认证，支持多点登录控制
**实现方式**:
- `server/utils/jwt.go` - JWT 工具类
- `server/middleware/jwt.go` - JWT 中间件验证
- `server/model/system/sys_jwt_blacklist.go` - 黑名单机制

#### 1.3 多点登录控制
**功能描述**: 支持单点登录或多点登录配置
**配置位置**: `server/config.yaml` - `system.single-point-login`
**实现**: Redis 缓存在线用户信息，自动清理过期会话

### 2. 权限管理系统

#### 2.1 角色权限管理
**功能描述**: 基于角色的访问控制 (RBAC)
**实现方式**:
- 后端: Casbin 权限引擎
- 前端: `web/src/view/superAdmin/authority/` - 角色管理页面
- API: `server/api/v1/system/authority_api.go`

#### 2.2 菜单权限控制
**功能描述**: 动态菜单权限，支持按钮级别权限控制
**实现方式**:
- `server/model/system/sys_base_menu.go` - 菜单模型
- `server/model/system/sys_authority_menu.go` - 角色菜单关联
- 前端: `web/src/directive/auth.js` - 权限指令

#### 2.3 API 权限控制
**功能描述**: 接口级别的权限验证
**实现方式**:
- `server/model/system/sys_api.go` - API 权限模型
- `server/middleware/casbin_rbac.go` - Casbin 中间件
- 前端: `web/src/api/authority.js` - 权限相关 API

### 3. 动态路由系统

#### 3.1 后端路由生成
**功能描述**: 根据用户权限动态生成前端路由配置
**实现方式**:
- `server/api/v1/system/menu_api.go` - 菜单数据接口
- `server/model/system/sys_base_menu.go` - 菜单数据结构

#### 3.2 前端路由加载
**功能描述**: 异步加载用户权限路由
**实现方式**:
- `web/src/pinia/modules/router.js` - 路由状态管理
- `web/src/utils/asyncRouter.js` - 异步路由处理
- `web/src/permission.js` - 路由权限守卫

### 4. 代码生成器 (核心特色功能)

#### 4.1 数据库表分析
**功能描述**: 自动分析数据库表结构生成 CRUD 代码
**实现方式**:
- `server/service/system/sys_auto_code_mysql.go` - MySQL 表结构解析
- `server/api/v1/system/auto_code_api.go` - 代码生成接口

#### 4.2 模板代码生成
**功能描述**: 基于模板生成前端 Vue 组件和后端 Go 代码
**实现方式**:
- `server/resource/autocode_template/` - 代码模板
- `server/resource/plug_template/` - 插件模板
- 前端: `web/src/view/systemTools/autoCode/index.vue` - 代码生成界面

#### 4.3 多数据库支持 (已删除)
**原功能**: 支持 MySQL/PgSQL/Oracle 多数据库代码生成
**删除原因**: 为简化项目，仅保留 MySQL 支持

### 5. 文件上传与存储

#### 5.1 对象存储集成
**功能描述**: 集成阿里云 OSS 文件存储
**实现方式**:
- `server/utils/upload/aliyun_oss.go` - OSS 上传实现
- `server/utils/upload/upload.go` - 统一上传接口

#### 5.2 本地文件存储 (已删除)
**原功能**: 支持本地文件系统存储
**删除原因**: 为简化项目，仅保留阿里云 OSS

#### 5.3 分片上传 (断点续传)
**功能描述**: 支持大文件分片上传，解决网络不稳定时文件上传中断的问题

**工作原理**:
1. **文件分片**: 将大文件分割成多个小片段（例如1MB一片）
2. **分片上传**: 逐个上传每个分片到服务器
3. **断点续传**: 如果网络中断，可以从上次上传的位置继续上传
4. **文件合并**: 所有分片上传完成后，在服务器端合并成完整文件

**举例说明**:
假设上传一个100MB的视频文件：
- 前端将文件分成100个1MB的分片
- 每个分片都有编号（如第1片、第2片...）
- 每个分片有自己的MD5校验码
- 上传过程中记录已上传的分片信息
- 如果第50片上传失败，重新连接后从第51片继续上传
- 所有分片上传完成后，服务器按顺序合并成完整视频文件

**实现方式**:
- 后端:
  - `server/utils/breakpoint_continue.go` - 核心分片逻辑
  - `server/api/v1/example/exa_breakpoint_continue.go` - 分片上传API
  - `server/model/example/exa_file_chunk.go` - 分片数据记录
- 前端: `web/src/view/example/upload/` - 分片上传组件

### 6. 数据字典管理

#### 6.1 字典数据维护
**功能描述**: 系统级数据字典配置
**实现方式**:
- `server/model/system/sys_dictionary.go` - 字典主表
- `server/model/system/sys_dictionary_detail.go` - 字典详情
- 前端: `web/src/view/superAdmin/dictionary/` - 字典管理界面

### 7. 系统监控与日志

#### 7.1 操作日志记录
**功能描述**: 记录用户操作行为
**实现方式**:
- `server/model/system/sys_operation_record.go` - 操作记录模型
- `server/middleware/operation.go` - 操作记录中间件
- 前端: `web/src/view/superAdmin/operation/` - 日志查看界面

#### 7.2 系统信息监控
**功能描述**: 显示服务器状态、内存使用等
**实现方式**:
- `server/api/v1/system/system_api.go` - 系统信息接口
- 前端: `web/src/view/system/state.vue` - 系统状态页面

### 8. 用户管理

#### 8.1 用户 CRUD
**功能描述**: 用户的增删改查操作
**实现方式**:
- `server/api/v1/system/user_api.go` - 用户管理接口
- `server/model/system/sys_user.go` - 用户模型
- 前端: `web/src/view/superAdmin/user/` - 用户管理界面

#### 8.2 用户角色分配
**功能描述**: 为用户分配角色权限
**实现方式**:
- `server/model/system/sys_user_authority.go` - 用户角色关联
- 前端权限分配界面集成在用户管理页面

### 9. 菜单管理系统

#### 9.1 动态菜单配置
**功能描述**: 可视化配置系统菜单结构
**实现方式**:
- `server/model/system/sys_base_menu.go` - 菜单数据模型
- 前端: `web/src/view/superAdmin/menu/` - 菜单管理界面

### 10. API 管理系统

#### 10.1 API 权限控制
**功能描述**: 对接口进行权限分组和控制
**实现方式**:
- `server/model/system/sys_api.go` - API 模型
- `server/model/system/sys_api_group.go` - API 分组
- 前端: `web/src/view/superAdmin/api/` - API 管理界面

### 11. Excel 导入导出

#### 11.1 Excel 数据处理
**功能描述**: 支持 Excel 文件的导入导出功能
**实现方式**:
- `server/utils/excel.go` - Excel 处理工具
- `server/api/v1/example/excel_api.go` - Excel 接口
- 前端: `web/src/view/example/excel/` - Excel 功能界面

### 12. 系统工具集成

#### 12.1 插件系统
**功能描述**: 支持插件的安装和管理
**实现方式**:
- `server/utils/plugin/` - 插件管理工具
- 前端: `web/src/view/systemTools/installPlugin/` - 插件安装界面

#### 12.2 自动化打包
**功能描述**: 项目自动化打包部署
**实现方式**:
- `Makefile` - 构建脚本 (已简化，移除 Docker 相关)
- 前端: `web/src/view/systemTools/autoPkg/` - 打包工具界面

### 13. 定时任务系统

#### 13.1 Cron 任务调度
**功能描述**: 支持定时任务的配置和管理
**实现方式**:
- `server/utils/timer/` - 定时器工具
- `server/config/timer.go` - 定时任务配置
- 前端: 系统配置中的 Timer 设置

---

## 架构设计特点

### 1. 分层架构
- **API 层**: 接口定义和参数验证
- **Service 层**: 业务逻辑处理
- **Model 层**: 数据模型和数据库操作
- **Utils 层**: 工具函数封装

### 2. 中间件架构
- JWT 认证中间件
- Casbin 权限中间件
- 操作记录中间件
- CORS 跨域中间件

### 3. 配置驱动
- YAML 配置文件驱动
- 环境变量支持
- 动态配置热更新

### 4. 插件化设计
- 代码生成器插件
- 文件上传插件
- 定时任务插件

---

## 已删除功能说明

为简化项目，以下功能已被移除：

### 1. Docker 部署支持
- **删除内容**: 所有 Dockerfile、docker-compose.yml、deploy/docker/ 目录
- **原因**: 用户明确表示不使用 Docker
- **替代方案**: 直接使用 `go run main.go` 和 `npm run serve`

### 2. 多数据库支持
- **删除内容**: PgSQL、Oracle 相关代码和服务
- **保留**: 仅 MySQL 支持
- **影响**: 代码生成器现在只支持 MySQL 数据库

### 3. 多 OSS 存储支持
- **删除内容**: 本地存储、七牛云、腾讯 COS、华为 OBS、AWS S3
- **保留**: 仅阿里云 OSS
- **影响**: 文件上传现在只支持阿里云 OSS

### 4. ESLint 代码检查
- **删除内容**: ESLint 配置和依赖
- **原因**: 用户不需要代码规范检查

### 5. VSCode 配置
- **删除内容**: `.code-workspace` 文件
- **原因**: 用户不使用 VSCode 开发

---

## 数据库设计

### 核心数据表
1. `sys_users` - 用户表
2. `sys_authorities` - 角色表
3. `sys_base_menus` - 菜单表
4. `sys_apis` - API 权限表
5. `sys_operation_records` - 操作日志表
6. `sys_dictionaries` - 数据字典表
7. `jwt_blacklists` - JWT 黑名单表

### 关系设计
- 用户 ↔ 角色 (多对多)
- 角色 ↔ 菜单 (多对多)
- 角色 ↔ API (多对多)

---

## 部署说明

### 环境要求
- Go 1.18+
- Node.js 18+
- MySQL 5.7+
- Redis (可选，用于会话管理)

### 启动步骤
1. 配置数据库连接 (`server/config.yaml`)
2. 配置阿里云 OSS (`server/config.yaml`)
3. 后端启动: `cd server && go run main.go`
4. 前端启动: `cd web && npm install && npm run serve`

### 生产部署
- 后端: `go build -o server main.go`
- 前端: `npm run build`

---

## 扩展建议

1. **监控告警**: 可集成 Prometheus + Grafana
2. **日志收集**: 可接入 ELK 栈
3. **缓存优化**: 增加 Redis 缓存层
4. **API 网关**: 可在前端增加 Nginx 反向代理
5. **容器化**: 如需要可重新添加 Docker 支持

---

## 注意事项

1. 代码生成器仅支持 MySQL 数据库
2. 文件上传仅支持阿里云 OSS
3. JWT Token 默认 7 天过期
4. 默认管理员账号: admin / 123456
5. 生产环境请修改默认密码和 JWT 密钥
