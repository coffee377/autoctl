# autoctl

自动化命令行工具

# 特性

- 多包或模块（monorepo）管理支持
- 语言支持 node,java,golang
- 包管理器识别（npm|yarn|pnpm）
- 自动版本号升级
- 升级日志生成

# 主要命令

- [ ] autoctl init 初始化
- [ ] autoctl changed 检查自上次发布以来哪些软件包被修改过
- [ ] autoctl release 创建一个新版本
- [ ] autoctl diff [package?]  列出所有或某个软件包自上次发布以来的修改情况

# 前端版本管理

# 后端版本管理

## Gradle

## Maven

# Credits

autoctl is powered by these awesome libraries

- [cobra](https://github.com/spf13/cobra)
- [git2go](https://github.com/libgit2/git2go)
- [logrus](https://github.com/sirupsen/logrus)