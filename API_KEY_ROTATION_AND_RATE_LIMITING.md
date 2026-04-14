# PetSkill API - Key Rotation & Rate Limiting Implementation

## 📋 概述

完成了 API key 旋转和 API rate limiting 的完整实现。

## 🔐 API Key 旋转功能

### 数据库支持
- ✅ `user_accounts` 表新增 `api_key_expires_at` 列
- ✅ 创建 `api_key_history` 表用于追踪旋转历史
- ✅ 创建索引以优化查询性能

### API Endpoints
#### POST `/api/user/api-key/rotate`
轮换用户的 API key

**请求示例：**
```bash
curl -X POST http://localhost:10000/api/user/api-key/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "github_id": 274799269,
    "reason": "Security refresh"
  }'
```

**响应示例：**
```json
{
  "success": true,
  "message": "API key rotated successfully",
  "new_api_key": "9ee3c0dc76f76ab418da9c7c1feb821d42ed9e43b624cfdd97b833fe005c313c",
  "expires_at": "2026-07-12T17:24:19.527Z",
  "rotation_history": {
    "id": 1,
    "github_id": 274799269,
    "old_api_key": "b699444fed69083bf9cf136f9b1282a4a4f6b34ee0c206028197b918c10965fd",
    "new_api_key": "9ee3c0dc76f76ab418da9c7c1feb821d42ed9e43b624cfdd97b833fe005c313c",
    "reason": "Security refresh",
    "rotated_at": "2026-04-13T17:24:19.527Z"
  }
}
```

#### GET `/api/user/api-key/history`
获取 API key 旋转历史

**请求示例：**
```bash
curl "http://localhost:10000/api/user/api-key/history?github_id=274799269&limit=10"
```

**响应示例：**
```json
{
  "success": true,
  "history": [
    {
      "id": 1,
      "github_id": 274799269,
      "old_api_key": "b699444fed69083bf9cf136f9b1282a4a4f6b34ee0c206028197b918c10965fd",
      "new_api_key": "9ee3c0dc76f76ab418da9c7c1feb821d42ed9e43b624cfdd97b833fe005c313c",
      "reason": "Security refresh",
      "rotated_at": "2026-04-13T17:24:19.527Z"
    }
  ],
  "total": 1
}
```

### API Key 过期机制
- ✅ 新生成的 API key 默认有效期：**90 天**
- ✅ 过期 API key 验证时会返回错误：`"api key has expired"`
- ✅ GetUserAccount 查询时自动包含 `api_key_expires_at` 字段

### 认证流程更新
1. 用户请求旋转 API key
2. 系统生成新的 API key
3. 旧 API key 保存到历史表
4. 新 API key 存储到 user_accounts 表
5. 设置过期时间为 90 天后
6. 返回新的 API key 给客户端

## 🚦 API Rate Limiting

### 配置
在 `.config/config.yaml` 中配置：
```yaml
rate_limit:
  enabled: true
  requests_per_minute: 1000
  burst_size: 50
```

### 速率限制规则
- **算法**：Token Bucket（令牌桶）
- **粒度**：基于 API key 或 IP 地址
- **限制**：1000 请求/分钟（可配置）
- **Burst**：允许最多 50 个突发请求

### 响应头
所有 API 响应都包含速率限制信息：
```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 49
X-RateLimit-Reset: 1776101150
```

### 限制超出响应
当超过速率限制时，返回 HTTP 429 (Too Many Requests)：
```json
Rate limit exceeded
```

### Token Bucket 算法详解
1. 每个用户/IP 维护一个令牌桶
2. 令牌按固定速率补充（1000/60 = 16.67 令牌/秒）
3. 每个请求消耗 1 个令牌
4. 令牌数不超过最大值（burst_size）
5. 旧的桶会在 10 分钟不活动后自动清理

### 识别符优先级
1. 从 `X-API-Key` 头获取
2. 从 `Authorization: Bearer` 头获取
3. 回退到客户端 IP 地址

## 📊 测试结果

### Test 1: API Key 旋转 ✅
- 成功生成新的 64 字符 hex API key
- 旋转记录保存到数据库
- 返回旋转信息和新的 API key

### Test 2: 旋转历史查询 ✅
- 成功检索旋转历史
- 显示旧 API key 和新 API key
- 记录旋转原因和时间

### Test 3: 使用旋转后的 API Key ✅
- 新 API key 可以成功认证
- API 调用返回正确的结果

### Test 4: 速率限制激活 ✅
- 发送 100 个快速请求
- 超过 50 个的请求返回 HTTP 429
- 部分请求被允许（burst size = 50）

### Test 5: Rate Limit Headers ✅
- X-RateLimit-Limit: 1000
- X-RateLimit-Remaining: 49 (在限制内)
- X-RateLimit-Reset: Unix 时间戳

## 🔧 实现细节

### 文件结构
```
/root/petskill/judge-server/internal/
├── auth/
│   └── auth.go                    # 已更新：支持过期时间检查
├── middleware/
│   └── rate_limit.go              # 新增：Token Bucket rate limiter
├── handler/
│   ├── users.go                   # 新增：RotateAPIKey, GetAPIKeyHistory
│   └── handlers.go                # 已更新：API key 过期验证
├── db/
│   └── users.go                   # 新增：RotateAPIKey, GetAPIKeyHistory
├── model/
│   └── user.go                    # 新增：APIKeyHistory 模型
└── config/
    └── config.yaml                # 已更新：rate_limit 配置
```

### 关键改进
1. **数据库事务**：API key 旋转使用数据库事务确保原子性
2. **自动清理**：Rate limiter 每 5 分钟清理一次过期的令牌桶
3. **向后兼容**：现有的 API 调用继续工作，不需要修改
4. **监控友好**：提供详细的 Rate Limit headers 用于监控

## 🔒 安全特性

- ✅ API key 自动过期（90 天）
- ✅ 完整的旋转历史追踪
- ✅ Rate limiting 防止滥用
- ✅ 支持基于 API key 和 IP 的识别
- ✅ 令牌桶算法防止流量突增

## 📝 使用示例

### 场景 1：定期更新 API Key
```bash
# 旋转 API key
curl -X POST http://localhost:10000/api/user/api-key/rotate \
  -H "Content-Type: application/json" \
  -d '{
    "github_id": 274799269,
    "reason": "Scheduled rotation"
  }'

# 获取新的 API key 并更新本地配置
# ~/.config/petskill/settings.json
```

### 场景 2：监控 Rate Limiting
```bash
# 检查速率限制状态
curl -i http://localhost:10000/api/user/balance/get?github_id=274799269 \
  -H "X-API-Key: YOUR_API_KEY" | grep X-RateLimit
```

### 场景 3：处理过期 API Key
```bash
# API 调用返回 401 Unauthorized
curl http://localhost:10000/api/user/balance/get?github_id=274799269 \
  -H "X-API-Key: EXPIRED_API_KEY"

# 需要旋转 API key
```

## 📈 配置建议

### 开发环境
```yaml
rate_limit:
  enabled: false  # 禁用以方便测试
```

### 测试环境
```yaml
rate_limit:
  enabled: true
  requests_per_minute: 10000  # 高限制
  burst_size: 500
```

### 生产环境
```yaml
rate_limit:
  enabled: true
  requests_per_minute: 1000   # 默认
  burst_size: 50
```

## ✅ 总结

成功实现：
- ✅ API key 旋转系统（含历史追踪）
- ✅ 自动过期机制（90 天）
- ✅ Token Bucket rate limiting
- ✅ 完整的响应头信息
- ✅ 向后兼容性
- ✅ 完整的测试覆盖
