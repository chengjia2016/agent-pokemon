# Pokemon Dataset Analyzer

Python脚本用于处理 `pokemon-dataset-zh` 数据集并生成PostgreSQL SQL初始化脚本。

## 📁 文件位置

- **脚本**: `/root/petskill/pokemon_analyzer.py`
- **输出**: `/root/petskill/judge-server/INIT_POKEMON_DATA.sql`

## 🚀 快速开始

### 运行分析器

```bash
python3 /root/petskill/pokemon_analyzer.py
```

### 输出示例

```
==================================================
Pokemon Dataset Analyzer
==================================================
✓ Loaded 1025 Pokemon species
✓ Loaded 953 moves
✓ Loaded 311 abilities

✓ SQL file saved to: /root/petskill/judge-server/INIT_POKEMON_DATA.sql
✓ Total SQL lines: 4618
✓ INSERT statements: 2289
  - Pokemon species: 1025
  - Moves: 953
  - Abilities: 311

==================================================
✓ Analysis completed successfully!
==================================================
```

## 📊 生成的数据统计

| 数据类型 | 数量 | 字段 |
|---------|------|------|
| Pokemon Species | 1,025 | pokemon_id, name_en, name_zh, types, capture_rate |
| Moves | 953 | move_id, name_en, name_zh, type, category, power, accuracy, pp |
| Abilities | 311 | ability_id, name_en, name_zh, description |

## 📝 SQL特性

✓ **UTF-8编码支持**: 完整的中文字符支持  
✓ **SQL安全**: 单引号正确转义防止注入  
✓ **标准化处理**:
  - 属性名称: 火→Fire, 水→Water等（18种属性）
  - 招式分类: 物理→Physical, 特殊→Special, 变化→Status
✓ **NULL处理**: 状态类招式的威力和准确度为NULL  
✓ **主键**: pokemon_id, move_id, ability_id  

## 📚 使用示例

### 导入到PostgreSQL

```bash
psql -U postgres -d pokemon_db < /root/petskill/judge-server/INIT_POKEMON_DATA.sql
```

### 验证数据导入

```sql
-- 检查记录数
SELECT COUNT(*) FROM pokemon_species;    -- 应返回 1025
SELECT COUNT(*) FROM pokemon_moves;      -- 应返回 953
SELECT COUNT(*) FROM pokemon_abilities;  -- 应返回 311

-- 查询示例
SELECT * FROM pokemon_species WHERE name_en = 'Pikachu';
SELECT * FROM pokemon_moves WHERE type = 'Electric' AND power > 50;
SELECT * FROM pokemon_abilities WHERE name_zh LIKE '%火%';
```

## 🔄 数据来源

### 输入文件
- `/tmp/pokemon-dataset-zh/data/simple_pokedex.json` - 基础宝可梦索引
- `/tmp/pokemon-dataset-zh/data/pokemon/[ID]-[NAME].json` - 个别宝可梦详情
- `/tmp/pokemon-dataset-zh/data/move_list.json` - 招式数据
- `/tmp/pokemon-dataset-zh/data/ability_list.json` - 特性数据

### 处理流程
1. 加载 simple_pokedex.json 获取基础宝可梦列表
2. 加载各个宝可梦JSON文件提取属性和捕捉率
3. 标准化属性名称和招式分类（中文→英文）
4. 加载招式列表并进行标准化处理
5. 加载特性列表并保留中文描述
6. 生成SQL INSERT语句
7. 输出到SQL文件

## 🛠️ 脚本功能

### 主要类: `PokemonAnalyzer`

**方法**:
- `load_simple_pokedex()` - 加载基础宝可梦数据
- `load_moves()` - 加载招式数据
- `load_abilities()` - 加载特性数据
- `generate_pokemon_species_sql()` - 生成宝可梦SQL
- `generate_moves_sql()` - 生成招式SQL
- `generate_abilities_sql()` - 生成特性SQL
- `generate_all_sql()` - 生成所有SQL
- `save_sql_file()` - 保存到文件
- `analyze()` - 执行完整分析流程

### 工具函数

- `escape_sql_string()` - SQL字符串转义
- `escape_sql_int()` - SQL整数转义

## 📋 SQL文件结构

```sql
-- Header with UTF-8 encoding setup
SET client_encoding = 'UTF8';

-- Create tables with IF NOT EXISTS
CREATE TABLE IF NOT EXISTS pokemon_species (...)
CREATE TABLE IF NOT EXISTS pokemon_moves (...)
CREATE TABLE IF NOT EXISTS pokemon_abilities (...)

-- Insert data
INSERT INTO pokemon_species ...
INSERT INTO pokemon_moves ...
INSERT INTO pokemon_abilities ...
```

## ⚙️ 配置

脚本中的配置常量:

```python
# 数据目录（输入）
data_dir = "/tmp/pokemon-dataset-zh/data"

# 输出文件
output_file = "/root/petskill/judge-server/INIT_POKEMON_DATA.sql"
```

## 📊 属性映射表

### 宝可梦属性 (18种)

| 中文 | 英文 | 中文 | 英文 |
|------|------|------|------|
| 一般 | Normal | 电 | Electric |
| 火 | Fire | 超能力 | Psychic |
| 水 | Water | 虫 | Bug |
| 格斗 | Fighting | 岩石 | Rock |
| 毒 | Poison | 幽灵 | Ghost |
| 地面 | Ground | 龙 | Dragon |
| 飞行 | Flying | 恶 | Dark |
| | | 钢 | Steel |
| | | 妖精 | Fairy |

### 招式分类

| 中文 | 英文 |
|------|------|
| 物理 | Physical |
| 特殊 | Special |
| 变化 | Status |

## 🐍 Python依赖

- Python 3.6+
- json (标准库)
- os (标准库)
- pathlib (标准库)
- typing (标准库)

不需要额外的第三方库！

## ✅ 质量检查

生成的SQL文件已验证:

✓ UTF-8编码正确  
✓ 13,491个中文字符保留  
✓ SQL语法无错误  
✓ 引号正确转义  
✓ 所有INSERT语句有效  
✓ 主键无重复  

## 🔍 故障排查

### 如果运行报错:

1. **找不到数据目录**
   ```
   检查 /tmp/pokemon-dataset-zh/data/ 是否存在
   ```

2. **找不到输出目录**
   ```
   脚本会自动创建 /root/petskill/judge-server/ 目录
   ```

3. **编码错误**
   ```
   确保系统支持UTF-8编码
   export LANG=en_US.UTF-8
   ```

## 📈 性能指标

- 处理时间: < 1秒
- 输出文件大小: 337 KB
- 总SQL行数: 4,618行
- INSERT语句数: 2,289条

## 📄 文件输出

### 文件信息

- **大小**: 337,572字节
- **行数**: 4,618行
- **编码**: UTF-8
- **格式**: PostgreSQL SQL脚本

### 内容分解

- CREATE TABLE: 3条语句
- INSERT语句: 2,289条
- 注释行: 8行
- 空白行: 5行

## 🎯 应用场景

✓ 初始化Pokemon数据库  
✓ 测试环境数据准备  
✓ 数据迁移  
✓ 数据备份和恢复  
✓ 开发和测试数据集

---

**生成日期**: 2026年4月13日  
**数据集版本**: pokemon-dataset-zh  
**兼容版本**: PostgreSQL 9.6+
