# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

这是一个 Go 语言 monorepo 自动化命令行工具项目，支持多模块管理、语言版本升级（node/java/golang）、包管理器识别（npm/yarn/pnpm）、自动版本号升级和升级日志生成。

## 常用命令

### 构建和测试

```bash
# 运行所有测试
go test ./...

# 运行单个测试
go test -v ./internal/bid/test/... -run TestName

# 静态构建测试（需要先构建 libgit2）
go test --tags "static" ./...

# 安装静态版本
go install --tags "static" ./...

# 初始化 go work（首次拉取代码后需要运行）
make init
```

### 数据库迁移 (internal/bid 模块)

迁移使用 [Atlas](https://atlasgo.io/) 管理，位于 `internal/bid` 目录：

```bash
cd internal/bid

# 生成迁移差异
make diff

# 应用迁移到本地开发数据库
make local-apply

# 查看迁移状态
make local-status

# 应用到测试环境
make test-apply

# 应用到生产环境
make prod-apply
```

### ent 代码生成

修改 ent schema 后需要重新生成代码：

```bash
cd internal/bid
go generate ./...
# 或直接运行
go run ./cmd/entc.go
```

## 项目架构

```
autoctl/
├── cmd/                    # CLI 入口
│   └── autoctl/
│       ├── main.go         # 主程序入口
│       └── root/           # Cobra 根命令
├── internal/               # 内部模块（不导出）
│   ├── bid/                # 投标业务模块（主要业务）
│   │   ├── ent/            # ent ORM 生成的代码（勿手动修改）
│   │   ├── ent/schema/     # ent 数据模型定义
│   │   ├── migrations/     # Atlas 数据库迁移文件
│   │   ├── data/           # 数据初始化脚本
│   │   ├── ds/             # 数据源配置
│   │   └── cmd/            # ent 代码生成入口
│   ├── dingtalk/           # 钉钉集成模块
│   │   ├── agent/          # 钉钉 Agent 服务
│   │   ├── app/            # 钉钉应用配置
│   │   ├── oa/             # OA 审批相关
│   │   ├── es/             # 事件订阅处理
│   │   └── cicd/           # CI/CD 集成
│   └── cmd/                # 命令实现
├── pkg/                    # 公共包（可导出）
│   ├── semver/             # 语义化版本管理
│   ├── log/                # 日志封装
│   ├── security/           # 安全工具
│   ├── git/                # Git 操作封装
│   └── api/                # API 通用组件
└── configs/                # 配置文件
```

## 技术栈

- **CLI 框架**: Cobra + Viper
- **ORM**: Ent (基于 Go generate)
- **数据库**: MySQL + Atlas 迁移
- **缓存**: Redis
- **钉钉 SDK**: open-dingtalk/dingtalk-stream-sdk-go
- **其他**: JWT, MinIO, Docker

## 关键配置

- 配置文件默认查找路径：`$HOME/auto.yml`、`.`、`./conf`
- 配置文件格式支持：YAML、TOML、JSON
- 数据库配置通过 Atlas HCL 文件管理（`internal/bid/atlas.hcl`）

## 注意事项

- `internal/bid/ent/` 目录下的文件由 ent 自动生成，勿手动修改
- 数据库凭证和敏感配置不要提交到版本控制
- 运行测试时注意环境变量和数据库连接配置