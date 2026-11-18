# open-img-sync

把 Docker 镜像同步到无法直连外网（或网络受限）的服务器，终结手动 `docker pull` → `docker save` → `scp` → `docker load` 的循环。

## 痛点

| 现在的手动流程 | 用 open-img-sync |
| --- | --- |
| 本地 `docker pull`，占用本机磁盘与带宽 | 平台侧拉取，不依赖本机 Docker |
| `docker save` 导出大文件 | 流式传输，无中间文件 |
| `scp` 到目标服务器，手动输入账号密码 | 服务器与账号集中管理，一次配置 |
| 登录目标机 `docker load` | 自动远端 load，进度与失败重试可见 |
| 阿里云企业版等临时密码每次去控制台复制 | 自动获取并刷新临时凭据 |

## 规划能力

本仓库当前处于**需求收集与拆解阶段**，实现以 [issues](https://github.com/LinuxSuRen/open-img-sync/issues) 为准，欢迎认领：

- 需求总览与架构决策：[#1](https://github.com/LinuxSuRen/open-img-sync/issues/1)
- Linux 服务器与 SSH 账号管理：[#2](https://github.com/LinuxSuRen/open-img-sync/issues/2)
- 镜像源与 Registry 凭据管理：[#3](https://github.com/LinuxSuRen/open-img-sync/issues/3)
- 阿里云 ACR 企业版临时凭据自动获取：[#4](https://github.com/LinuxSuRen/open-img-sync/issues/4)
- 同步引擎与任务编排：[#5](https://github.com/LinuxSuRen/open-img-sync/issues/5)
- Web 管理界面（Vue 3）：[#6](https://github.com/LinuxSuRen/open-img-sync/issues/6)
- 单二进制分发与 CLI 规范：[#7](https://github.com/LinuxSuRen/open-img-sync/issues/7)
- 容器化运行与镜像发布：[#8](https://github.com/LinuxSuRen/open-img-sync/issues/8)
- CI 门面与仓库自动化：[#9](https://github.com/LinuxSuRen/open-img-sync/issues/9)
- 安全基线：凭据加密与错误码规范：[#10](https://github.com/LinuxSuRen/open-img-sync/issues/10)

## 技术形态

- 后端 Go，前端 Vue 3 + Vite，构建产物嵌入二进制
- 单二进制即可运行，同时提供多架构容器镜像（GHCR）
- CLI 提供 `serve` 与 `version` 子命令

## 协作约定

- 功能开发一律 `feat/`、`fix/` 分支 + Pull Request，主分支受保护（禁直推、禁强推），合并由维护者执行
- PR 描述须附验证证据（测试命令与输出）；涉及 UI 的改动须附页面截图
- 提交信息使用 Conventional Commits（中文描述）

## 许可证

[Apache-2.0](LICENSE)
