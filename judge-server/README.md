# Judge Server - Agent Monster

> 🏛️ 裁判服务器 - 验证宠物、战斗和排行榜的合法性

## 功能

### 1. 宠物验证
- 验证宠物精灵数据是否合规
- 检查属性值是否在合法范围内
- 验证基因权重总和
- 验证签名

### 2. 战斗验证
- 验证战斗结果是否合法
- 检查攻击/防御技能是否有效
- 验证战斗时长
- 防止作弊行为

### 3. 排行榜系统
- 每日自动生成排行榜
- ELO 评级系统
- 胜负记录追踪

### 4. 食物和生长记录
- 记录食物消费
- 追踪宠物成长
- 验证升级合法性

### 5. GitHub 同步
- 同步排行榜到 Issues
- 同步战斗报告到 Issues
- 更新 Wiki 页面
- 触发 GitHub Actions

---

## 快速开始

### 1. 安装依赖

```bash
cd judge-server
go mod tidy
```

### 2. 配置 PostgreSQL

```sql
CREATE DATABASE agent_monster;
```

### 3. 配置文件

编辑 `.config/config.yaml`:

```yaml
server:
  host: 0.0.0.0
  port: 8080

database:
  host: localhost
  port: 5432
  user: postgres
  password: xiaodudu
  dbname: agent_monster
  sslmode: disable

github:
  token: your_github_token
  owner: chengjia2016
  repo: agent-monster

settlement:
  daily_time: "00:00"
  sync_to_github: true
```

### 4. 运行服务器

```bash
go run cmd/main.go
```

---

## API 端点

| 端点 | 方法 | 描述 |
|------|------|------|
| `/health` | GET | 健康检查 |
| `/api/pet/validate` | POST | 验证宠物数据 |
| `/api/battle/validate` | POST | 验证战斗结果 |
| `/api/leaderboard` | GET | 获取排行榜 |
| `/api/leaderboard/daily` | POST | 生成每日排行榜 |
| `/api/food/record` | POST | 记录食物消费 |
| `/api/growth/record` | POST | 记录宠物成长 |

---

## 数据库表结构

### pets
存储所有宠物数据

### battles
存储战斗记录

### leaderboard
玩家排行榜

### food_transactions
食物消费记录

### growth_records
宠物成长记录

### daily_leaderboards
每日排行榜快照

---

## 部署

### Systemd 服务

```bash
sudo cp judge-server.service /etc/systemd/system/
sudo systemctl enable judge-server
sudo systemctl start judge-server
```

### Docker

```bash
docker build -t judge-server .
docker run -p 8080:8080 judge-server
```

---

## 安全注意事项

⚠️ **此代码包含敏感配置，不应提交到 GitHub**

- 数据库密码
- GitHub Token
- 服务器配置

确保 `judge-server/` 目录在 `.gitignore` 中。

---

## 开发

```bash
# 运行测试
go test ./...

# 格式化代码
go fmt ./...

# 构建
go build -o judge-server cmd/main.go
```
