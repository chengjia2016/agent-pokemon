-- =========================================
-- Pokemon Database Initialization Script
-- Auto-generated from pokemon-dataset-zh
-- =========================================
-- Set encoding
SET client_encoding = 'UTF8';

-- Pokemon Species Data
CREATE TABLE IF NOT EXISTS pokemon_species (
    pokemon_id VARCHAR(10) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    types TEXT,
    capture_rate INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0001', 'Bulbasaur', '妙蛙种子', 'Grass,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0002', 'Ivysaur', '妙蛙草', 'Grass,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0003', 'Venusaur', '妙蛙花', 'Grass,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0004', 'Charmander', '小火龙', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0005', 'Charmeleon', '火恐龙', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0006', 'Charizard', '喷火龙', 'Fire,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0007', 'Squirtle', '杰尼龟', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0008', 'Wartortle', '卡咪龟', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0009', 'Blastoise', '水箭龟', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0010', 'Caterpie', '绿毛虫', 'Bug', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0011', 'Metapod', '铁甲蛹', 'Bug', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0012', 'Butterfree', '巴大蝶', 'Bug,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0013', 'Weedle', '独角虫', 'Bug,Poison', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0014', 'Kakuna', '铁壳蛹', 'Bug,Poison', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0015', 'Beedrill', '大针蜂', 'Bug,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0016', 'Pidgey', '波波', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0017', 'Pidgeotto', '比比鸟', 'Normal,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0018', 'Pidgeot', '大比鸟', 'Normal,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0019', 'Rattata', '小拉达', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0020', 'Raticate', '拉达', 'Normal', 127);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0021', 'Spearow', '烈雀', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0022', 'Fearow', '大嘴雀', 'Normal,Flying', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0023', 'Ekans', '阿柏蛇', 'Poison', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0024', 'Arbok', '阿柏怪', 'Poison', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0025', 'Pikachu', '皮卡丘', 'Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0026', 'Raichu', '雷丘', 'Electric', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0027', 'Sandshrew', '穿山鼠', 'Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0028', 'Sandslash', '穿山王', 'Ground', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0029', 'Nidoran♀', '尼多兰', 'Poison', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0030', 'Nidorina', '尼多娜', 'Poison', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0031', 'Nidoqueen', '尼多后', 'Poison,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0032', 'Nidoran♂', '尼多朗', 'Poison', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0033', 'Nidorino', '尼多力诺', 'Poison', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0034', 'Nidoking', '尼多王', 'Poison,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0035', 'Clefairy', '皮皮', 'Fairy', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0036', 'Clefable', '皮可西', 'Fairy', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0037', 'Vulpix', '六尾', 'Fire', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0038', 'Ninetales', '九尾', 'Fire', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0039', 'Jigglypuff', '胖丁', 'Normal,Fairy', 170);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0040', 'Wigglytuff', '胖可丁', 'Normal,Fairy', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0041', 'Zubat', '超音蝠', 'Poison,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0042', 'Golbat', '大嘴蝠', 'Poison,Flying', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0043', 'Oddish', '走路草', 'Grass,Poison', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0044', 'Gloom', '臭臭花', 'Grass,Poison', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0045', 'Vileplume', '霸王花', 'Grass,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0046', 'Paras', '派拉斯', 'Bug,Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0047', 'Parasect', '派拉斯特', 'Bug,Grass', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0048', 'Venonat', '毛球', 'Bug,Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0049', 'Venomoth', '摩鲁蛾', 'Bug,Poison', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0050', 'Diglett', '地鼠', 'Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0051', 'Dugtrio', '三地鼠', 'Ground', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0052', 'Meowth', '喵喵', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0053', 'Persian', '猫老大', 'Normal', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0054', 'Psyduck', '可达鸭', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0055', 'Golduck', '哥达鸭', 'Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0056', 'Mankey', '猴怪', 'Fighting', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0057', 'Primeape', '火暴猴', 'Fighting', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0058', 'Growlithe', '卡蒂狗', 'Fire', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0059', 'Arcanine', '风速狗', 'Fire', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0060', 'Poliwag', '蚊香蝌蚪', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0061', 'Poliwhirl', '蚊香君', 'Water', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0062', 'Poliwrath', '蚊香泳士', 'Water,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0063', 'Abra', '凯西', 'Psychic', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0064', 'Kadabra', '勇基拉', 'Psychic', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0065', 'Alakazam', '胡地', 'Psychic', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0066', 'Machop', '腕力', 'Fighting', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0067', 'Machoke', '豪力', 'Fighting', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0068', 'Machamp', '怪力', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0069', 'Bellsprout', '喇叭芽', 'Grass,Poison', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0070', 'Weepinbell', '口呆花', 'Grass,Poison', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0071', 'Victreebel', '大食花', 'Grass,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0072', 'Tentacool', '玛瑙水母', 'Water,Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0073', 'Tentacruel', '毒刺水母', 'Water,Poison', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0074', 'Geodude', '小拳石', 'Rock,Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0075', 'Graveler', '隆隆石', 'Rock,Ground', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0076', 'Golem', '隆隆岩', 'Rock,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0077', 'Ponyta', '小火马', 'Fire', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0078', 'Rapidash', '烈焰马', 'Fire', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0079', 'Slowpoke', '呆呆兽', 'Water,Psychic', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0080', 'Slowbro', '呆壳兽', 'Water,Psychic', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0081', 'Magnemite', '小磁怪', 'Electric,Steel', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0082', 'Magneton', '三合一磁怪', 'Electric,Steel', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0083', 'Farfetch''d', '大葱鸭', 'Normal,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0084', 'Doduo', '嘟嘟', 'Normal,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0085', 'Dodrio', '嘟嘟利', 'Normal,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0086', 'Seel', '小海狮', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0087', 'Dewgong', '白海狮', 'Water,Ice', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0088', 'Grimer', '臭泥', 'Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0089', 'Muk', '臭臭泥', 'Poison', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0090', 'Shellder', '大舌贝', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0091', 'Cloyster', '刺甲贝', 'Water,Ice', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0092', 'Gastly', '鬼斯', 'Ghost,Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0093', 'Haunter', '鬼斯通', 'Ghost,Poison', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0094', 'Gengar', '耿鬼', 'Ghost,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0095', 'Onix', '大岩蛇', 'Rock,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0096', 'Drowzee', '催眠貘', 'Psychic', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0097', 'Hypno', '引梦貘人', 'Psychic', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0098', 'Krabby', '大钳蟹', 'Water', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0099', 'Kingler', '巨钳蟹', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0100', 'Voltorb', '霹雳电球', 'Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0101', 'Electrode', '顽皮雷弹', 'Electric', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0102', 'Exeggcute', '蛋蛋', 'Grass,Psychic', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0103', 'Exeggutor', '椰蛋树', 'Grass,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0104', 'Cubone', '卡拉卡拉', 'Ground', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0105', 'Marowak', '嘎啦嘎啦', 'Ground', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0106', 'Hitmonlee', '飞腿郎', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0107', 'Hitmonchan', '快拳郎', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0108', 'Lickitung', '大舌头', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0109', 'Koffing', '瓦斯弹', 'Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0110', 'Weezing', '双弹瓦斯', 'Poison', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0111', 'Rhyhorn', '独角犀牛', 'Ground,Rock', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0112', 'Rhydon', '钻角犀兽', 'Ground,Rock', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0113', 'Chansey', '吉利蛋', 'Normal', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0114', 'Tangela', '蔓藤怪', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0115', 'Kangaskhan', '袋兽', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0116', 'Horsea', '墨海马', 'Water', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0117', 'Seadra', '海刺龙', 'Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0118', 'Goldeen', '角金鱼', 'Water', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0119', 'Seaking', '金鱼王', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0120', 'Staryu', '海星星', 'Water', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0121', 'Starmie', '宝石海星', 'Water,Psychic', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0122', 'Mr. Mime', '魔墙人偶', 'Psychic,Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0123', 'Scyther', '飞天螳螂', 'Bug,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0124', 'Jynx', '迷唇姐', 'Ice,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0125', 'Electabuzz', '电击兽', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0126', 'Magmar', '鸭嘴火兽', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0127', 'Pinsir', '凯罗斯', 'Bug', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0128', 'Tauros', '肯泰罗', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0129', 'Magikarp', '鲤鱼王', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0130', 'Gyarados', '暴鲤龙', 'Water,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0131', 'Lapras', '拉普拉斯', 'Water,Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0132', 'Ditto', '百变怪', 'Normal', 35);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0133', 'Eevee', '伊布', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0134', 'Vaporeon', '水伊布', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0135', 'Jolteon', '雷伊布', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0136', 'Flareon', '火伊布', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0137', 'Porygon', '多边兽', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0138', 'Omanyte', '菊石兽', 'Rock,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0139', 'Omastar', '多刺菊石兽', 'Rock,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0140', 'Kabuto', '化石盔', 'Rock,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0141', 'Kabutops', '镰刀盔', 'Rock,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0142', 'Aerodactyl', '化石翼龙', 'Rock,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0143', 'Snorlax', '卡比兽', 'Normal', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0144', 'Articuno', '急冻鸟', 'Ice,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0145', 'Zapdos', '闪电鸟', 'Electric,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0146', 'Moltres', '火焰鸟', 'Fire,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0147', 'Dratini', '迷你龙', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0148', 'Dragonair', '哈克龙', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0149', 'Dragonite', '快龙', 'Dragon,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0150', 'Mewtwo', '超梦', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0151', 'Mew', '梦幻', 'Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0152', 'Chikorita', '菊草叶', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0153', 'Bayleef', '月桂叶', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0154', 'Meganium', '大竺葵', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0155', 'Cyndaquil', '火球鼠', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0156', 'Quilava', '火岩鼠', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0157', 'Typhlosion', '火暴兽', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0158', 'Totodile', '小锯鳄', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0159', 'Croconaw', '蓝鳄', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0160', 'Feraligatr', '大力鳄', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0161', 'Sentret', '尾立', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0162', 'Furret', '大尾立', 'Normal', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0163', 'Hoothoot', '咕咕', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0164', 'Noctowl', '猫头夜鹰', 'Normal,Flying', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0165', 'Ledyba', '芭瓢虫', 'Bug,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0166', 'Ledian', '安瓢虫', 'Bug,Flying', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0167', 'Spinarak', '圆丝蛛', 'Bug,Poison', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0168', 'Ariados', '阿利多斯', 'Bug,Poison', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0169', 'Crobat', '叉字蝠', 'Poison,Flying', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0170', 'Chinchou', '灯笼鱼', 'Water,Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0171', 'Lanturn', '电灯怪', 'Water,Electric', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0172', 'Pichu', '皮丘', 'Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0173', 'Cleffa', '皮宝宝', 'Fairy', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0174', 'Igglybuff', '宝宝丁', 'Normal,Fairy', 170);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0175', 'Togepi', '波克比', 'Fairy', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0176', 'Togetic', '波克基古', 'Fairy,Flying', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0177', 'Natu', '天然雀', 'Psychic,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0178', 'Xatu', '天然鸟', 'Psychic,Flying', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0179', 'Mareep', '咩利羊', 'Electric', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0180', 'Flaaffy', '茸茸羊', 'Electric', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0181', 'Ampharos', '电龙', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0182', 'Bellossom', '美丽花', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0183', 'Marill', '玛力露', 'Water,Fairy', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0184', 'Azumarill', '玛力露丽', 'Water,Fairy', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0185', 'Sudowoodo', '树才怪', 'Rock', 65);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0186', 'Politoed', '蚊香蛙皇', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0187', 'Hoppip', '毽子草', 'Grass,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0188', 'Skiploom', '毽子花', 'Grass,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0189', 'Jumpluff', '毽子棉', 'Grass,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0190', 'Aipom', '长尾怪手', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0191', 'Sunkern', '向日种子', 'Grass', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0192', 'Sunflora', '向日花怪', 'Grass', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0193', 'Yanma', '蜻蜻蜓', 'Bug,Flying', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0194', 'Wooper', '乌波', 'Water,Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0195', 'Quagsire', '沼王', 'Water,Ground', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0196', 'Espeon', '太阳伊布', 'Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0197', 'Umbreon', '月亮伊布', 'Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0198', 'Murkrow', '黑暗鸦', 'Dark,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0199', 'Slowking', '呆呆王', 'Water,Psychic', 70);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0200', 'Misdreavus', '梦妖', 'Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0201', 'Unown', '未知图腾', 'Psychic', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0202', 'Wobbuffet', '果然翁', 'Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0203', 'Girafarig', '麒麟奇', 'Normal,Psychic', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0204', 'Pineco', '榛果球', 'Bug', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0205', 'Forretress', '佛烈托斯', 'Bug,Steel', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0206', 'Dunsparce', '土龙弟弟', 'Normal', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0207', 'Gligar', '天蝎', 'Ground,Flying', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0208', 'Steelix', '大钢蛇', 'Steel,Ground', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0209', 'Snubbull', '布鲁', 'Fairy', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0210', 'Granbull', '布鲁皇', 'Fairy', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0211', 'Qwilfish', '千针鱼', 'Water,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0212', 'Scizor', '巨钳螳螂', 'Bug,Steel', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0213', 'Shuckle', '壶壶', 'Bug,Rock', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0214', 'Heracross', '赫拉克罗斯', 'Bug,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0215', 'Sneasel', '狃拉', 'Dark,Ice', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0216', 'Teddiursa', '熊宝宝', 'Normal', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0217', 'Ursaring', '圈圈熊', 'Normal', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0218', 'Slugma', '熔岩虫', 'Fire', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0219', 'Magcargo', '熔岩蜗牛', 'Fire,Rock', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0220', 'Swinub', '小山猪', 'Ice,Ground', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0221', 'Piloswine', '长毛猪', 'Ice,Ground', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0222', 'Corsola', '太阳珊瑚', 'Water,Rock', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0223', 'Remoraid', '铁炮鱼', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0224', 'Octillery', '章鱼桶', 'Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0225', 'Delibird', '信使鸟', 'Ice,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0226', 'Mantine', '巨翅飞鱼', 'Water,Flying', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0227', 'Skarmory', '盔甲鸟', 'Steel,Flying', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0228', 'Houndour', '戴鲁比', 'Dark,Fire', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0229', 'Houndoom', '黑鲁加', 'Dark,Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0230', 'Kingdra', '刺龙王', 'Water,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0231', 'Phanpy', '小小象', 'Ground', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0232', 'Donphan', '顿甲', 'Ground', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0233', 'Porygon2', '多边兽Ⅱ', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0234', 'Stantler', '惊角鹿', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0235', 'Smeargle', '图图犬', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0236', 'Tyrogue', '无畏小子', 'Fighting', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0237', 'Hitmontop', '战舞郎', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0238', 'Smoochum', '迷唇娃', 'Ice,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0239', 'Elekid', '电击怪', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0240', 'Magby', '鸭嘴宝宝', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0241', 'Miltank', '大奶罐', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0242', 'Blissey', '幸福蛋', 'Normal', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0243', 'Raikou', '雷公', 'Electric', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0244', 'Entei', '炎帝', 'Fire', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0245', 'Suicune', '水君', 'Water', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0246', 'Larvitar', '幼基拉斯', 'Rock,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0247', 'Pupitar', '沙基拉斯', 'Rock,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0248', 'Tyranitar', '班基拉斯', 'Rock,Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0249', 'Lugia', '洛奇亚', 'Psychic,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0250', 'Ho-Oh', '凤王', 'Fire,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0251', 'Celebi', '时拉比', 'Psychic,Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0252', 'Treecko', '木守宫', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0253', 'Grovyle', '森林蜥蜴', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0254', 'Sceptile', '蜥蜴王', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0255', 'Torchic', '火稚鸡', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0256', 'Combusken', '力壮鸡', 'Fire,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0257', 'Blaziken', '火焰鸡', 'Fire,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0258', 'Mudkip', '水跃鱼', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0259', 'Marshtomp', '沼跃鱼', 'Water,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0260', 'Swampert', '巨沼怪', 'Water,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0261', 'Poochyena', '土狼犬', 'Dark', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0262', 'Mightyena', '大狼犬', 'Dark', 127);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0263', 'Zigzagoon', '蛇纹熊', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0264', 'Linoone', '直冲熊', 'Normal', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0265', 'Wurmple', '刺尾虫', 'Bug', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0266', 'Silcoon', '甲壳茧', 'Bug', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0267', 'Beautifly', '狩猎凤蝶', 'Bug,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0268', 'Cascoon', '盾甲茧', 'Bug', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0269', 'Dustox', '毒粉蛾', 'Bug,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0270', 'Lotad', '莲叶童子', 'Water,Grass', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0271', 'Lombre', '莲帽小童', 'Water,Grass', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0272', 'Ludicolo', '乐天河童', 'Water,Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0273', 'Seedot', '橡实果', 'Grass', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0274', 'Nuzleaf', '长鼻叶', 'Grass,Dark', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0275', 'Shiftry', '狡猾天狗', 'Grass,Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0276', 'Taillow', '傲骨燕', 'Normal,Flying', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0277', 'Swellow', '大王燕', 'Normal,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0278', 'Wingull', '长翅鸥', 'Water,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0279', 'Pelipper', '大嘴鸥', 'Water,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0280', 'Ralts', '拉鲁拉丝', 'Psychic,Fairy', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0281', 'Kirlia', '奇鲁莉安', 'Psychic,Fairy', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0282', 'Gardevoir', '沙奈朵', 'Psychic,Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0283', 'Surskit', '溜溜糖球', 'Bug,Water', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0284', 'Masquerain', '雨翅蛾', 'Bug,Flying', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0285', 'Shroomish', '蘑蘑菇', 'Grass', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0286', 'Breloom', '斗笠菇', 'Grass,Fighting', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0287', 'Slakoth', '懒人獭', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0288', 'Vigoroth', '过动猿', 'Normal', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0289', 'Slaking', '请假王', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0290', 'Nincada', '土居忍士', 'Bug,Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0291', 'Ninjask', '铁面忍者', 'Bug,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0292', 'Shedinja', '脱壳忍者', 'Bug,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0293', 'Whismur', '咕妞妞', 'Normal', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0294', 'Loudred', '吼爆弹', 'Normal', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0295', 'Exploud', '爆音怪', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0296', 'Makuhita', '幕下力士', 'Fighting', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0297', 'Hariyama', '铁掌力士', 'Fighting', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0298', 'Azurill', '露力丽', 'Normal,Fairy', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0299', 'Nosepass', '朝北鼻', 'Rock', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0300', 'Skitty', '向尾喵', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0301', 'Delcatty', '优雅猫', 'Normal', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0302', 'Sableye', '勾魂眼', 'Dark,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0303', 'Mawile', '大嘴娃', 'Steel,Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0304', 'Aron', '可可多拉', 'Steel,Rock', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0305', 'Lairon', '可多拉', 'Steel,Rock', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0306', 'Aggron', '波士可多拉', 'Steel,Rock', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0307', 'Meditite', '玛沙那', 'Fighting,Psychic', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0308', 'Medicham', '恰雷姆', 'Fighting,Psychic', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0309', 'Electrike', '落雷兽', 'Electric', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0310', 'Manectric', '雷电兽', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0311', 'Plusle', '正电拍拍', 'Electric', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0312', 'Minun', '负电拍拍', 'Electric', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0313', 'Volbeat', '电萤虫', 'Bug', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0314', 'Illumise', '甜甜萤', 'Bug', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0315', 'Roselia', '毒蔷薇', 'Grass,Poison', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0316', 'Gulpin', '溶食兽', 'Poison', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0317', 'Swalot', '吞食兽', 'Poison', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0318', 'Carvanha', '利牙鱼', 'Water,Dark', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0319', 'Sharpedo', '巨牙鲨', 'Water,Dark', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0320', 'Wailmer', '吼吼鲸', 'Water', 125);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0321', 'Wailord', '吼鲸王', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0322', 'Numel', '呆火驼', 'Fire,Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0323', 'Camerupt', '喷火驼', 'Fire,Ground', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0324', 'Torkoal', '煤炭龟', 'Fire', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0325', 'Spoink', '跳跳猪', 'Psychic', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0326', 'Grumpig', '噗噗猪', 'Psychic', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0327', 'Spinda', '晃晃斑', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0328', 'Trapinch', '大颚蚁', 'Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0329', 'Vibrava', '超音波幼虫', 'Ground,Dragon', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0330', 'Flygon', '沙漠蜻蜓', 'Ground,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0331', 'Cacnea', '刺球仙人掌', 'Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0332', 'Cacturne', '梦歌仙人掌', 'Grass,Dark', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0333', 'Swablu', '青绵鸟', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0334', 'Altaria', '七夕青鸟', 'Dragon,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0335', 'Zangoose', '猫鼬斩', 'Normal', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0336', 'Seviper', '饭匙蛇', 'Poison', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0337', 'Lunatone', '月石', 'Rock,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0338', 'Solrock', '太阳岩', 'Rock,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0339', 'Barboach', '泥泥鳅', 'Water,Ground', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0340', 'Whiscash', '鲶鱼王', 'Water,Ground', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0341', 'Corphish', '龙虾小兵', 'Water', 205);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0342', 'Crawdaunt', '铁螯龙虾', 'Water,Dark', 155);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0343', 'Baltoy', '天秤偶', 'Ground,Psychic', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0344', 'Claydol', '念力土偶', 'Ground,Psychic', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0345', 'Lileep', '触手百合', 'Rock,Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0346', 'Cradily', '摇篮百合', 'Rock,Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0347', 'Anorith', '太古羽虫', 'Rock,Bug', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0348', 'Armaldo', '太古盔甲', 'Rock,Bug', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0349', 'Feebas', '丑丑鱼', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0350', 'Milotic', '美纳斯', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0351', 'Castform', '飘浮泡泡', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0352', 'Kecleon', '变隐龙', 'Normal', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0353', 'Shuppet', '怨影娃娃', 'Ghost', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0354', 'Banette', '诅咒娃娃', 'Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0355', 'Duskull', '夜巡灵', 'Ghost', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0356', 'Dusclops', '彷徨夜灵', 'Ghost', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0357', 'Tropius', '热带龙', 'Grass,Flying', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0358', 'Chimecho', '风铃铃', 'Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0359', 'Absol', '阿勃梭鲁', 'Dark', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0360', 'Wynaut', '小果然', 'Psychic', 125);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0361', 'Snorunt', '雪童子', 'Ice', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0362', 'Glalie', '冰鬼护', 'Ice', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0363', 'Spheal', '海豹球', 'Ice,Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0364', 'Sealeo', '海魔狮', 'Ice,Water', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0365', 'Walrein', '帝牙海狮', 'Ice,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0366', 'Clamperl', '珍珠贝', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0367', 'Huntail', '猎斑鱼', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0368', 'Gorebyss', '樱花鱼', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0369', 'Relicanth', '古空棘鱼', 'Water,Rock', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0370', 'Luvdisc', '爱心鱼', 'Water', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0371', 'Bagon', '宝贝龙', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0372', 'Shelgon', '甲壳龙', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0373', 'Salamence', '暴飞龙', 'Dragon,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0374', 'Beldum', '铁哑铃', 'Steel,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0375', 'Metang', '金属怪', 'Steel,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0376', 'Metagross', '巨金怪', 'Steel,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0377', 'Regirock', '雷吉洛克', 'Rock', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0378', 'Regice', '雷吉艾斯', 'Ice', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0379', 'Registeel', '雷吉斯奇鲁', 'Steel', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0380', 'Latias', '拉帝亚斯', 'Dragon,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0381', 'Latios', '拉帝欧斯', 'Dragon,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0382', 'Kyogre', '盖欧卡', 'Water', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0383', 'Groudon', '固拉多', 'Ground', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0384', 'Rayquaza', '烈空坐', 'Dragon,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0385', 'Jirachi', '基拉祈', 'Steel,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0386', 'Deoxys', '代欧奇希斯', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0387', 'Turtwig', '草苗龟', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0388', 'Grotle', '树林龟', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0389', 'Torterra', '土台龟', 'Grass,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0390', 'Chimchar', '小火焰猴', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0391', 'Monferno', '猛火猴', 'Fire,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0392', 'Infernape', '烈焰猴', 'Fire,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0393', 'Piplup', '波加曼', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0394', 'Prinplup', '波皇子', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0395', 'Empoleon', '帝王拿波', 'Water,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0396', 'Starly', '姆克儿', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0397', 'Staravia', '姆克鸟', 'Normal,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0398', 'Staraptor', '姆克鹰', 'Normal,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0399', 'Bidoof', '大牙狸', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0400', 'Bibarel', '大尾狸', 'Normal,Water', 127);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0401', 'Kricketot', '圆法师', 'Bug', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0402', 'Kricketune', '音箱蟀', 'Bug', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0403', 'Shinx', '小猫怪', 'Electric', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0404', 'Luxio', '勒克猫', 'Electric', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0405', 'Luxray', '伦琴猫', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0406', 'Budew', '含羞苞', 'Grass,Poison', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0407', 'Roserade', '罗丝雷朵', 'Grass,Poison', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0408', 'Cranidos', '头盖龙', 'Rock', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0409', 'Rampardos', '战槌龙', 'Rock', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0410', 'Shieldon', '盾甲龙', 'Rock,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0411', 'Bastiodon', '护城龙', 'Rock,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0412', 'Burmy', '结草儿', 'Bug', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0413', 'Wormadam', '结草贵妇', 'Bug,Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0414', 'Mothim', '绅士蛾', 'Bug,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0415', 'Combee', '三蜜蜂', 'Bug,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0416', 'Vespiquen', '蜂女王', 'Bug,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0417', 'Pachirisu', '帕奇利兹', 'Electric', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0418', 'Buizel', '泳圈鼬', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0419', 'Floatzel', '浮潜鼬', 'Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0420', 'Cherubi', '樱花宝', 'Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0421', 'Cherrim', '樱花儿', 'Grass', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0422', 'Shellos', '无壳海兔', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0423', 'Gastrodon', '海兔兽', 'Water,Ground', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0424', 'Ambipom', '双尾怪手', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0425', 'Drifloon', '飘飘球', 'Ghost,Flying', 125);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0426', 'Drifblim', '随风球', 'Ghost,Flying', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0427', 'Buneary', '卷卷耳', 'Normal', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0428', 'Lopunny', '长耳兔', 'Normal', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0429', 'Mismagius', '梦妖魔', 'Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0430', 'Honchkrow', '乌鸦头头', 'Dark,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0431', 'Glameow', '魅力喵', 'Normal', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0432', 'Purugly', '东施喵', 'Normal', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0433', 'Chingling', '铃铛响', 'Psychic', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0434', 'Stunky', '臭鼬噗', 'Poison,Dark', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0435', 'Skuntank', '坦克臭鼬', 'Poison,Dark', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0436', 'Bronzor', '铜镜怪', 'Steel,Psychic', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0437', 'Bronzong', '青铜钟', 'Steel,Psychic', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0438', 'Bonsly', '盆才怪', 'Rock', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0439', 'Mime Jr.', '魔尼尼', 'Psychic,Fairy', 145);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0440', 'Happiny', '小福蛋', 'Normal', 130);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0441', 'Chatot', '聒噪鸟', 'Normal,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0442', 'Spiritomb', '花岩怪', 'Ghost,Dark', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0443', 'Gible', '圆陆鲨', 'Dragon,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0444', 'Gabite', '尖牙陆鲨', 'Dragon,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0445', 'Garchomp', '烈咬陆鲨', 'Dragon,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0446', 'Munchlax', '小卡比兽', 'Normal', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0447', 'Riolu', '利欧路', 'Fighting', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0448', 'Lucario', '路卡利欧', 'Fighting,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0449', 'Hippopotas', '沙河马', 'Ground', 140);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0450', 'Hippowdon', '河马兽', 'Ground', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0451', 'Skorupi', '钳尾蝎', 'Poison,Bug', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0452', 'Drapion', '龙王蝎', 'Poison,Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0453', 'Croagunk', '不良蛙', 'Poison,Fighting', 140);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0454', 'Toxicroak', '毒骷蛙', 'Poison,Fighting', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0455', 'Carnivine', '尖牙笼', 'Grass', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0456', 'Finneon', '荧光鱼', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0457', 'Lumineon', '霓虹鱼', 'Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0458', 'Mantyke', '小球飞鱼', 'Water,Flying', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0459', 'Snover', '雪笠怪', 'Grass,Ice', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0460', 'Abomasnow', '暴雪王', 'Grass,Ice', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0461', 'Weavile', '玛狃拉', 'Dark,Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0462', 'Magnezone', '自爆磁怪', 'Electric,Steel', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0463', 'Lickilicky', '大舌舔', 'Normal', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0464', 'Rhyperior', '超甲狂犀', 'Ground,Rock', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0465', 'Tangrowth', '巨蔓藤', 'Grass', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0466', 'Electivire', '电击魔兽', 'Electric', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0467', 'Magmortar', '鸭嘴炎兽', 'Fire', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0468', 'Togekiss', '波克基斯', 'Fairy,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0469', 'Yanmega', '远古巨蜓', 'Bug,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0470', 'Leafeon', '叶伊布', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0471', 'Glaceon', '冰伊布', 'Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0472', 'Gliscor', '天蝎王', 'Ground,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0473', 'Mamoswine', '象牙猪', 'Ice,Ground', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0474', 'Porygon-Z', '多边兽Ｚ', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0475', 'Gallade', '艾路雷朵', 'Psychic,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0476', 'Probopass', '大朝北鼻', 'Rock,Steel', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0477', 'Dusknoir', '黑夜魔灵', 'Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0478', 'Froslass', '雪妖女', 'Ice,Ghost', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0479', 'Rotom', '洛托姆', 'Electric,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0480', 'Uxie', '由克希', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0481', 'Mesprit', '艾姆利多', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0482', 'Azelf', '亚克诺姆', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0483', 'Dialga', '帝牙卢卡', 'Steel,Dragon', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0484', 'Palkia', '帕路奇亚', 'Water,Dragon', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0485', 'Heatran', '席多蓝恩', 'Fire,Steel', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0486', 'Regigigas', '雷吉奇卡斯', 'Normal', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0487', 'Giratina', '骑拉帝纳', 'Ghost,Dragon', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0488', 'Cresselia', '克雷色利亚', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0489', 'Phione', '霏欧纳', 'Water', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0490', 'Manaphy', '玛纳霏', 'Water', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0491', 'Darkrai', '达克莱伊', 'Dark', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0492', 'Shaymin', '谢米', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0493', 'Arceus', '阿尔宙斯', 'Normal', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0494', 'Victini', '比克提尼', 'Psychic,Fire', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0495', 'Snivy', '藤藤蛇', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0496', 'Servine', '青藤蛇', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0497', 'Serperior', '君主蛇', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0498', 'Tepig', '暖暖猪', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0499', 'Pignite', '炒炒猪', 'Fire,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0500', 'Emboar', '炎武王', 'Fire,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0501', 'Oshawott', '水水獭', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0502', 'Dewott', '双刃丸', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0503', 'Samurott', '大剑鬼', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0504', 'Patrat', '探探鼠', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0505', 'Watchog', '步哨鼠', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0506', 'Lillipup', '小约克', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0507', 'Herdier', '哈约克', 'Normal', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0508', 'Stoutland', '长毛狗', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0509', 'Purrloin', '扒手猫', 'Dark', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0510', 'Liepard', '酷豹', 'Dark', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0511', 'Pansage', '花椰猴', 'Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0512', 'Simisage', '花椰猿', 'Grass', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0513', 'Pansear', '爆香猴', 'Fire', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0514', 'Simisear', '爆香猿', 'Fire', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0515', 'Panpour', '冷水猴', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0516', 'Simipour', '冷水猿', 'Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0517', 'Munna', '食梦梦', 'Psychic', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0518', 'Musharna', '梦梦蚀', 'Psychic', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0519', 'Pidove', '豆豆鸽', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0520', 'Tranquill', '咕咕鸽', 'Normal,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0521', 'Unfezant', '高傲雉鸡', 'Normal,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0522', 'Blitzle', '斑斑马', 'Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0523', 'Zebstrika', '雷电斑马', 'Electric', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0524', 'Roggenrola', '石丸子', 'Rock', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0525', 'Boldore', '地幔岩', 'Rock', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0526', 'Gigalith', '庞岩怪', 'Rock', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0527', 'Woobat', '滚滚蝙蝠', 'Psychic,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0528', 'Swoobat', '心蝙蝠', 'Psychic,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0529', 'Drilbur', '螺钉地鼠', 'Ground', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0530', 'Excadrill', '龙头地鼠', 'Ground,Steel', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0531', 'Audino', '差不多娃娃', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0532', 'Timburr', '搬运小匠', 'Fighting', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0533', 'Gurdurr', '铁骨土人', 'Fighting', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0534', 'Conkeldurr', '修建老匠', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0535', 'Tympole', '圆蝌蚪', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0536', 'Palpitoad', '蓝蟾蜍', 'Water,Ground', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0537', 'Seismitoad', '蟾蜍王', 'Water,Ground', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0538', 'Throh', '投摔鬼', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0539', 'Sawk', '打击鬼', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0540', 'Sewaddle', '虫宝包', 'Bug,Grass', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0541', 'Swadloon', '宝包茧', 'Bug,Grass', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0542', 'Leavanny', '保姆虫', 'Bug,Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0543', 'Venipede', '百足蜈蚣', 'Bug,Poison', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0544', 'Whirlipede', '车轮球', 'Bug,Poison', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0545', 'Scolipede', '蜈蚣王', 'Bug,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0546', 'Cottonee', '木棉球', 'Grass,Fairy', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0547', 'Whimsicott', '风妖精', 'Grass,Fairy', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0548', 'Petilil', '百合根娃娃', 'Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0549', 'Lilligant', '裙儿小姐', 'Grass', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0550', 'Basculin', '野蛮鲈鱼', 'Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0551', 'Sandile', '黑眼鳄', 'Ground,Dark', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0552', 'Krokorok', '混混鳄', 'Ground,Dark', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0553', 'Krookodile', '流氓鳄', 'Ground,Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0554', 'Darumaka', '火红不倒翁', 'Fire', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0555', 'Darmanitan', '达摩狒狒', 'Fire', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0556', 'Maractus', '沙铃仙人掌', 'Grass', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0557', 'Dwebble', '石居蟹', 'Bug,Rock', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0558', 'Crustle', '岩殿居蟹', 'Bug,Rock', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0559', 'Scraggy', '滑滑小子', 'Dark,Fighting', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0560', 'Scrafty', '头巾混混', 'Dark,Fighting', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0561', 'Sigilyph', '象征鸟', 'Psychic,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0562', 'Yamask', '哭哭面具', 'Ghost', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0563', 'Cofagrigus', '死神棺', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0564', 'Tirtouga', '原盖海龟', 'Water,Rock', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0565', 'Carracosta', '肋骨海龟', 'Water,Rock', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0566', 'Archen', '始祖小鸟', 'Rock,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0567', 'Archeops', '始祖大鸟', 'Rock,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0568', 'Trubbish', '破破袋', 'Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0569', 'Garbodor', '灰尘山', 'Poison', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0570', 'Zorua', '索罗亚', 'Dark', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0571', 'Zoroark', '索罗亚克', 'Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0572', 'Minccino', '泡沫栗鼠', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0573', 'Cinccino', '奇诺栗鼠', 'Normal', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0574', 'Gothita', '哥德宝宝', 'Psychic', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0575', 'Gothorita', '哥德小童', 'Psychic', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0576', 'Gothitelle', '哥德小姐', 'Psychic', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0577', 'Solosis', '单卵细胞球', 'Psychic', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0578', 'Duosion', '双卵细胞球', 'Psychic', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0579', 'Reuniclus', '人造细胞卵', 'Psychic', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0580', 'Ducklett', '鸭宝宝', 'Water,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0581', 'Swanna', '舞天鹅', 'Water,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0582', 'Vanillite', '迷你冰', 'Ice', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0583', 'Vanillish', '多多冰', 'Ice', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0584', 'Vanilluxe', '双倍多多冰', 'Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0585', 'Deerling', '四季鹿', 'Normal,Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0586', 'Sawsbuck', '萌芽鹿', 'Normal,Grass', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0587', 'Emolga', '电飞鼠', 'Electric,Flying', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0588', 'Karrablast', '盖盖虫', 'Bug', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0589', 'Escavalier', '骑士蜗牛', 'Bug,Steel', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0590', 'Foongus', '哎呀球菇', 'Grass,Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0591', 'Amoonguss', '败露球菇', 'Grass,Poison', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0592', 'Frillish', '轻飘飘', 'Water,Ghost', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0593', 'Jellicent', '胖嘟嘟', 'Water,Ghost', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0594', 'Alomomola', '保姆曼波', 'Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0595', 'Joltik', '电电虫', 'Bug,Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0596', 'Galvantula', '电蜘蛛', 'Bug,Electric', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0597', 'Ferroseed', '种子铁球', 'Grass,Steel', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0598', 'Ferrothorn', '坚果哑铃', 'Grass,Steel', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0599', 'Klink', '齿轮儿', 'Steel', 130);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0600', 'Klang', '齿轮组', 'Steel', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0601', 'Klinklang', '齿轮怪', 'Steel', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0602', 'Tynamo', '麻麻小鱼', 'Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0603', 'Eelektrik', '麻麻鳗', 'Electric', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0604', 'Eelektross', '麻麻鳗鱼王', 'Electric', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0605', 'Elgyem', '小灰怪', 'Psychic', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0606', 'Beheeyem', '大宇怪', 'Psychic', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0607', 'Litwick', '烛光灵', 'Ghost,Fire', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0608', 'Lampent', '灯火幽灵', 'Ghost,Fire', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0609', 'Chandelure', '水晶灯火灵', 'Ghost,Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0610', 'Axew', '牙牙', 'Dragon', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0611', 'Fraxure', '斧牙龙', 'Dragon', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0612', 'Haxorus', '双斧战龙', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0613', 'Cubchoo', '喷嚏熊', 'Ice', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0614', 'Beartic', '冻原熊', 'Ice', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0615', 'Cryogonal', '几何雪花', 'Ice', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0616', 'Shelmet', '小嘴蜗', 'Bug', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0617', 'Accelgor', '敏捷虫', 'Bug', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0618', 'Stunfisk', '泥巴鱼', 'Ground,Electric', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0619', 'Mienfoo', '功夫鼬', 'Fighting', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0620', 'Mienshao', '师父鼬', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0621', 'Druddigon', '赤面龙', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0622', 'Golett', '泥偶小人', 'Ground,Ghost', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0623', 'Golurk', '泥偶巨人', 'Ground,Ghost', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0624', 'Pawniard', '驹刀小兵', 'Dark,Steel', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0625', 'Bisharp', '劈斩司令', 'Dark,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0626', 'Bouffalant', '爆炸头水牛', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0627', 'Rufflet', '毛头小鹰', 'Normal,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0628', 'Braviary', '勇士雄鹰', 'Normal,Flying', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0629', 'Vullaby', '秃鹰丫头', 'Dark,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0630', 'Mandibuzz', '秃鹰娜', 'Dark,Flying', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0631', 'Heatmor', '熔蚁兽', 'Fire', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0632', 'Durant', '铁蚁', 'Bug,Steel', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0633', 'Deino', '单首龙', 'Dark,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0634', 'Zweilous', '双首暴龙', 'Dark,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0635', 'Hydreigon', '三首恶龙', 'Dark,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0636', 'Larvesta', '燃烧虫', 'Bug,Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0637', 'Volcarona', '火神蛾', 'Bug,Fire', 15);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0638', 'Cobalion', '勾帕路翁', 'Steel,Fighting', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0639', 'Terrakion', '代拉基翁', 'Rock,Fighting', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0640', 'Virizion', '毕力吉翁', 'Grass,Fighting', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0641', 'Tornadus', '龙卷云', 'Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0642', 'Thundurus', '雷电云', 'Electric,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0643', 'Reshiram', '莱希拉姆', 'Dragon,Fire', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0644', 'Zekrom', '捷克罗姆', 'Dragon,Electric', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0645', 'Landorus', '土地云', 'Ground,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0646', 'Kyurem', '酋雷姆', 'Dragon,Ice', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0647', 'Keldeo', '凯路迪欧', 'Water,Fighting', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0648', 'Meloetta', '美洛耶塔', 'Normal,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0649', 'Genesect', '盖诺赛克特', 'Bug,Steel', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0650', 'Chespin', '哈力栗', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0651', 'Quilladin', '胖胖哈力', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0652', 'Chesnaught', '布里卡隆', 'Grass,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0653', 'Fennekin', '火狐狸', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0654', 'Braixen', '长尾火狐', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0655', 'Delphox', '妖火红狐', 'Fire,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0656', 'Froakie', '呱呱泡蛙', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0657', 'Frogadier', '呱头蛙', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0658', 'Greninja', '甲贺忍蛙', 'Water,Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0659', 'Bunnelby', '掘掘兔', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0660', 'Diggersby', '掘地兔', 'Normal,Ground', 127);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0661', 'Fletchling', '小箭雀', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0662', 'Fletchinder', '火箭雀', 'Fire,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0663', 'Talonflame', '烈箭鹰', 'Fire,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0664', 'Scatterbug', '粉蝶虫', 'Bug', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0665', 'Spewpa', '粉蝶蛹', 'Bug', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0666', 'Vivillon', '彩粉蝶', 'Bug,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0667', 'Litleo', '小狮狮', 'Fire,Normal', 220);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0668', 'Pyroar', '火炎狮', 'Fire,Normal', 65);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0669', 'Flabébé', '花蓓蓓', 'Fairy', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0670', 'Floette', '花叶蒂', 'Fairy', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0671', 'Florges', '花洁夫人', 'Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0672', 'Skiddo', '坐骑小羊', 'Grass', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0673', 'Gogoat', '坐骑山羊', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0674', 'Pancham', '顽皮熊猫', 'Fighting', 220);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0675', 'Pangoro', '流氓熊猫', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0676', 'Furfrou', '多丽米亚', 'Normal', 160);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0677', 'Espurr', '妙喵', 'Psychic', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0678', 'Meowstic', '超能妙喵', 'Psychic', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0679', 'Honedge', '独剑鞘', 'Steel,Ghost', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0680', 'Doublade', '双剑鞘', 'Steel,Ghost', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0681', 'Aegislash', '坚盾剑怪', 'Steel,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0682', 'Spritzee', '粉香香', 'Fairy', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0683', 'Aromatisse', '芳香精', 'Fairy', 140);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0684', 'Swirlix', '绵绵泡芙', 'Fairy', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0685', 'Slurpuff', '胖甜妮', 'Fairy', 140);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0686', 'Inkay', '好啦鱿', 'Dark,Psychic', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0687', 'Malamar', '乌贼王', 'Dark,Psychic', 80);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0688', 'Binacle', '龟脚脚', 'Rock,Water', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0689', 'Barbaracle', '龟足巨铠', 'Rock,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0690', 'Skrelp', '垃垃藻', 'Poison,Water', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0691', 'Dragalge', '毒藻龙', 'Poison,Dragon', 55);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0692', 'Clauncher', '铁臂枪虾', 'Water', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0693', 'Clawitzer', '钢炮臂虾', 'Water', 55);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0694', 'Helioptile', '伞电蜥', 'Electric,Normal', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0695', 'Heliolisk', '光电伞蜥', 'Electric,Normal', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0696', 'Tyrunt', '宝宝暴龙', 'Rock,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0697', 'Tyrantrum', '怪颚龙', 'Rock,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0698', 'Amaura', '冰雪龙', 'Rock,Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0699', 'Aurorus', '冰雪巨龙', 'Rock,Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0700', 'Sylveon', '仙子伊布', 'Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0701', 'Hawlucha', '摔角鹰人', 'Fighting,Flying', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0702', 'Dedenne', '咚咚鼠', 'Electric,Fairy', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0703', 'Carbink', '小碎钻', 'Rock,Fairy', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0704', 'Goomy', '黏黏宝', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0705', 'Sliggoo', '黏美儿', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0706', 'Goodra', '黏美龙', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0707', 'Klefki', '钥圈儿', 'Steel,Fairy', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0708', 'Phantump', '小木灵', 'Ghost,Grass', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0709', 'Trevenant', '朽木妖', 'Ghost,Grass', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0710', 'Pumpkaboo', '南瓜精', 'Ghost,Grass', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0711', 'Gourgeist', '南瓜怪人', 'Ghost,Grass', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0712', 'Bergmite', '冰宝', 'Ice', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0713', 'Avalugg', '冰岩怪', 'Ice', 55);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0714', 'Noibat', '嗡蝠', 'Flying,Dragon', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0715', 'Noivern', '音波龙', 'Flying,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0716', 'Xerneas', '哲尔尼亚斯', 'Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0717', 'Yveltal', '伊裴尔塔尔', 'Dark,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0718', 'Zygarde', '基格尔德', 'Dragon,Ground', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0719', 'Diancie', '蒂安希', 'Rock,Fairy', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0720', 'Hoopa', '胡帕', 'Psychic,Ghost', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0721', 'Volcanion', '波尔凯尼恩', 'Fire,Water', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0722', 'Rowlet', '木木枭', 'Grass,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0723', 'Dartrix', '投羽枭', 'Grass,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0724', 'Decidueye', '狙射树枭', 'Grass,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0725', 'Litten', '火斑喵', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0726', 'Torracat', '炎热喵', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0727', 'Incineroar', '炽焰咆哮虎', 'Fire,Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0728', 'Popplio', '球球海狮', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0729', 'Brionne', '花漾海狮', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0730', 'Primarina', '西狮海壬', 'Water,Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0731', 'Pikipek', '小笃儿', 'Normal,Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0732', 'Trumbeak', '喇叭啄鸟', 'Normal,Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0733', 'Toucannon', '铳嘴大鸟', 'Normal,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0734', 'Yungoos', '猫鼬少', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0735', 'Gumshoos', '猫鼬探长', 'Normal', 127);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0736', 'Grubbin', '强颚鸡母虫', 'Bug', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0737', 'Charjabug', '虫电宝', 'Bug,Electric', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0738', 'Vikavolt', '锹农炮虫', 'Bug,Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0739', 'Crabrawler', '好胜蟹', 'Fighting', 225);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0740', 'Crabominable', '好胜毛蟹', 'Fighting,Ice', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0741', 'Oricorio', '花舞鸟', 'Fire,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0742', 'Cutiefly', '萌虻', 'Bug,Fairy', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0743', 'Ribombee', '蝶结萌虻', 'Bug,Fairy', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0744', 'Rockruff', '岩狗狗', 'Rock', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0745', 'Lycanroc', '鬃岩狼人', 'Rock', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0746', 'Wishiwashi', '弱丁鱼', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0747', 'Mareanie', '好坏星', 'Poison,Water', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0748', 'Toxapex', '超坏星', 'Poison,Water', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0749', 'Mudbray', '泥驴仔', 'Ground', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0750', 'Mudsdale', '重泥挽马', 'Ground', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0751', 'Dewpider', '滴蛛', 'Water,Bug', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0752', 'Araquanid', '滴蛛霸', 'Water,Bug', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0753', 'Fomantis', '伪螳草', 'Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0754', 'Lurantis', '兰螳花', 'Grass', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0755', 'Morelull', '睡睡菇', 'Grass,Fairy', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0756', 'Shiinotic', '灯罩夜菇', 'Grass,Fairy', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0757', 'Salandit', '夜盗火蜥', 'Poison,Fire', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0758', 'Salazzle', '焰后蜥', 'Poison,Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0759', 'Stufful', '童偶熊', 'Normal,Fighting', 140);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0760', 'Bewear', '穿着熊', 'Normal,Fighting', 70);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0761', 'Bounsweet', '甜竹竹', 'Grass', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0762', 'Steenee', '甜舞妮', 'Grass', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0763', 'Tsareena', '甜冷美后', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0764', 'Comfey', '花疗环环', 'Fairy', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0765', 'Oranguru', '智挥猩', 'Normal,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0766', 'Passimian', '投掷猴', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0767', 'Wimpod', '胆小虫', 'Bug,Water', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0768', 'Golisopod', '具甲武者', 'Bug,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0769', 'Sandygast', '沙丘娃', 'Ghost,Ground', 140);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0770', 'Palossand', '噬沙堡爷', 'Ghost,Ground', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0771', 'Pyukumuku', '拳海参', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0772', 'Type: Null', '属性：空', 'Normal', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0773', 'Silvally', '银伴战兽', 'Normal', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0774', 'Minior', '小陨星', 'Rock,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0775', 'Komala', '树枕尾熊', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0776', 'Turtonator', '爆焰龟兽', 'Fire,Dragon', 70);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0777', 'Togedemaru', '托戈德玛尔', 'Electric,Steel', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0778', 'Mimikyu', '谜拟Ｑ', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0779', 'Bruxish', '磨牙彩皮鱼', 'Water,Psychic', 80);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0780', 'Drampa', '老翁龙', 'Normal,Dragon', 70);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0781', 'Dhelmise', '破破舵轮', 'Ghost,Grass', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0782', 'Jangmo-o', '心鳞宝', 'Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0783', 'Hakamo-o', '鳞甲龙', 'Dragon,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0784', 'Kommo-o', '杖尾鳞甲龙', 'Dragon,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0785', 'Tapu Koko', '卡璞・鸣鸣', 'Electric,Fairy', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0786', 'Tapu Lele', '卡璞・蝶蝶', 'Psychic,Fairy', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0787', 'Tapu Bulu', '卡璞・哞哞', 'Grass,Fairy', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0788', 'Tapu Fini', '卡璞・鳍鳍', 'Water,Fairy', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0789', 'Cosmog', '科斯莫古', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0790', 'Cosmoem', '科斯莫姆', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0791', 'Solgaleo', '索尔迦雷欧', 'Psychic,Steel', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0792', 'Lunala', '露奈雅拉', 'Psychic,Ghost', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0793', 'Nihilego', '虚吾伊德', 'Rock,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0794', 'Buzzwole', '爆肌蚊', 'Bug,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0795', 'Pheromosa', '费洛美螂', 'Bug,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0796', 'Xurkitree', '电束木', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0797', 'Celesteela', '铁火辉夜', 'Steel,Flying', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0798', 'Kartana', '纸御剑', 'Grass,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0799', 'Guzzlord', '恶食大王', 'Dark,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0800', 'Necrozma', '奈克洛兹玛', 'Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0801', 'Magearna', '玛机雅娜', 'Steel,Fairy', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0802', 'Marshadow', '玛夏多', 'Fighting,Ghost', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0803', 'Poipole', '毒贝比', 'Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0804', 'Naganadel', '四颚针龙', 'Poison,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0805', 'Stakataka', '垒磊石', 'Rock,Steel', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0806', 'Blacephalon', '砰头小丑', 'Fire,Ghost', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0807', 'Zeraora', '捷拉奥拉', 'Electric', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0808', 'Meltan', '美录坦', 'Steel', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0809', 'Melmetal', '美录梅塔', 'Steel', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0810', 'Grookey', '敲音猴', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0811', 'Thwackey', '啪咚猴', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0812', 'Rillaboom', '轰擂金刚猩', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0813', 'Scorbunny', '炎兔儿', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0814', 'Raboot', '腾蹴小将', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0815', 'Cinderace', '闪焰王牌', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0816', 'Sobble', '泪眼蜥', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0817', 'Drizzile', '变涩蜥', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0818', 'Inteleon', '千面避役', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0819', 'Skwovet', '贪心栗鼠', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0820', 'Greedent', '藏饱栗鼠', 'Normal', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0821', 'Rookidee', '稚山雀', 'Flying', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0822', 'Corvisquire', '蓝鸦', 'Flying', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0823', 'Corviknight', '钢铠鸦', 'Flying,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0824', 'Blipbug', '索侦虫', 'Bug', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0825', 'Dottler', '天罩虫', 'Bug,Psychic', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0826', 'Orbeetle', '以欧路普', 'Bug,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0827', 'Nickit', '偷儿狐', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0828', 'Thievul', '狐大盗', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0829', 'Gossifleur', '幼棉棉', 'Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0830', 'Eldegoss', '白蓬蓬', 'Grass', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0831', 'Wooloo', '毛辫羊', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0832', 'Dubwool', '毛毛角羊', 'Normal', 127);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0833', 'Chewtle', '咬咬龟', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0834', 'Drednaw', '暴噬龟', 'Water,Rock', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0835', 'Yamper', '来电汪', 'Electric', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0836', 'Boltund', '逐电犬', 'Electric', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0837', 'Rolycoly', '小炭仔', 'Rock', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0838', 'Carkol', '大炭车', 'Rock,Fire', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0839', 'Coalossal', '巨炭山', 'Rock,Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0840', 'Applin', '啃果虫', 'Grass,Dragon', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0841', 'Flapple', '苹裹龙', 'Grass,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0842', 'Appletun', '丰蜜龙', 'Grass,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0843', 'Silicobra', '沙包蛇', 'Ground', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0844', 'Sandaconda', '沙螺蟒', 'Ground', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0845', 'Cramorant', '古月鸟', 'Flying,Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0846', 'Arrokuda', '刺梭鱼', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0847', 'Barraskewda', '戽斗尖梭', 'Water', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0848', 'Toxel', '毒电婴', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0849', 'Toxtricity', '颤弦蝾螈', 'Electric,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0850', 'Sizzlipede', '烧火蚣', 'Fire,Bug', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0851', 'Centiskorch', '焚焰蚣', 'Fire,Bug', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0852', 'Clobbopus', '拳拳蛸', 'Fighting', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0853', 'Grapploct', '八爪武师', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0854', 'Sinistea', '来悲茶', 'Ghost', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0855', 'Polteageist', '怖思壶', 'Ghost', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0856', 'Hatenna', '迷布莉姆', 'Psychic', 235);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0857', 'Hattrem', '提布莉姆', 'Psychic', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0858', 'Hatterene', '布莉姆温', 'Psychic,Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0859', 'Impidimp', '捣蛋小妖', 'Dark,Fairy', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0860', 'Morgrem', '诈唬魔', 'Dark,Fairy', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0861', 'Grimmsnarl', '长毛巨魔', 'Dark,Fairy', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0862', 'Obstagoon', '堵拦熊', 'Dark,Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0863', 'Perrserker', '喵头目', 'Steel', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0864', 'Cursola', '魔灵珊瑚', 'Ghost', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0865', 'Sirfetch''d', '葱游兵', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0866', 'Mr. Rime', '踏冰人偶', 'Ice,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0867', 'Runerigus', '死神板', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0868', 'Milcery', '小仙奶', 'Fairy', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0869', 'Alcremie', '霜奶仙', 'Fairy', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0870', 'Falinks', '列阵兵', 'Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0871', 'Pincurchin', '啪嚓海胆', 'Electric', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0872', 'Snom', '雪吞虫', 'Ice,Bug', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0873', 'Frosmoth', '雪绒蛾', 'Ice,Bug', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0874', 'Stonjourner', '巨石丁', 'Rock', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0875', 'Eiscue', '冰砌鹅', 'Ice', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0876', 'Indeedee', '爱管侍', 'Psychic,Normal', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0877', 'Morpeko', '莫鲁贝可', 'Electric,Dark', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0878', 'Cufant', '铜象', 'Steel', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0879', 'Copperajah', '大王铜象', 'Steel', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0880', 'Dracozolt', '雷鸟龙', 'Electric,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0881', 'Arctozolt', '雷鸟海兽', 'Electric,Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0882', 'Dracovish', '鳃鱼龙', 'Water,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0883', 'Arctovish', '鳃鱼海兽', 'Water,Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0884', 'Duraludon', '铝钢龙', 'Steel,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0885', 'Dreepy', '多龙梅西亚', 'Dragon,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0886', 'Drakloak', '多龙奇', 'Dragon,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0887', 'Dragapult', '多龙巴鲁托', 'Dragon,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0888', 'Zacian', '苍响', 'Fairy', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0889', 'Zamazenta', '藏玛然特', 'Fighting', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0890', 'Eternatus', '无极汰那', 'Poison,Dragon', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0891', 'Kubfu', '熊徒弟', 'Fighting', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0892', 'Urshifu', '武道熊师', 'Fighting,Dark', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0893', 'Zarude', '萨戮德', 'Dark,Grass', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0894', 'Regieleki', '雷吉艾勒奇', 'Electric', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0895', 'Regidrago', '雷吉铎拉戈', 'Dragon', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0896', 'Glastrier', '雪暴马', 'Ice', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0897', 'Spectrier', '灵幽马', 'Ghost', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0898', 'Calyrex', '蕾冠王', 'Psychic,Grass', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0899', 'Wyrdeer', '诡角鹿', 'Normal,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0900', 'Kleavor', '劈斧螳螂', 'Bug,Rock', 15);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0901', 'Ursaluna', '月月熊', 'Ground,Normal', 20);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0902', 'Basculegion', '幽尾玄鱼', 'Water,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0903', 'Sneasler', '大狃拉', 'Fighting,Poison', 20);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0904', 'Overqwil', '万针鱼', 'Dark,Poison', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0905', 'Enamorus', '眷恋云', 'Fairy,Flying', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0906', 'Sprigatito', '新叶喵', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0907', 'Floragato', '蒂蕾喵', 'Grass', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0908', 'Meowscarada', '魔幻假面喵', 'Grass,Dark', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0909', 'Fuecoco', '呆火鳄', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0910', 'Crocalor', '炙烫鳄', 'Fire', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0911', 'Skeledirge', '骨纹巨声鳄', 'Fire,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0912', 'Quaxly', '润水鸭', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0913', 'Quaxwell', '涌跃鸭', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0914', 'Quaquaval', '狂欢浪舞鸭', 'Water,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0915', 'Lechonk', '爱吃豚', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0916', 'Oinkologne', '飘香豚', 'Normal', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0917', 'Tarountula', '团珠蛛', 'Bug', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0918', 'Spidops', '操陷蛛', 'Bug', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0919', 'Nymble', '豆蟋蟀', 'Bug', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0920', 'Lokix', '烈腿蝗', 'Bug,Dark', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0921', 'Pawmi', '布拨', 'Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0922', 'Pawmo', '布土拨', 'Electric,Fighting', 80);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0923', 'Pawmot', '巴布土拨', 'Electric,Fighting', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0924', 'Tandemaus', '一对鼠', 'Normal', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0925', 'Maushold', '一家鼠', 'Normal', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0926', 'Fidough', '狗仔包', 'Fairy', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0927', 'Dachsbun', '麻花犬', 'Fairy', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0928', 'Smoliv', '迷你芙', 'Grass,Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0929', 'Dolliv', '奥利纽', 'Grass,Normal', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0930', 'Arboliva', '奥利瓦', 'Grass,Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0931', 'Squawkabilly', '怒鹦哥', 'Normal,Flying', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0932', 'Nacli', '盐石宝', 'Rock', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0933', 'Naclstack', '盐石垒', 'Rock', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0934', 'Garganacl', '盐石巨灵', 'Rock', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0935', 'Charcadet', '炭小侍', 'Fire', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0936', 'Armarouge', '红莲铠骑', 'Fire,Psychic', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0937', 'Ceruledge', '苍炎刃鬼', 'Fire,Ghost', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0938', 'Tadbulb', '光蚪仔', 'Electric', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0939', 'Bellibolt', '电肚蛙', 'Electric', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0940', 'Wattrel', '电海燕', 'Electric,Flying', 180);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0941', 'Kilowattrel', '大电海燕', 'Electric,Flying', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0942', 'Maschiff', '偶叫獒', 'Dark', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0943', 'Mabosstiff', '獒教父', 'Dark', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0944', 'Shroodle', '滋汁鼹', 'Poison,Normal', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0945', 'Grafaiai', '涂标客', 'Poison,Normal', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0946', 'Bramblin', '纳噬草', 'Grass,Ghost', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0947', 'Brambleghast', '怖纳噬草', 'Grass,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0948', 'Toedscool', '原野水母', 'Ground,Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0949', 'Toedscruel', '陆地水母', 'Ground,Grass', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0950', 'Klawf', '毛崖蟹', 'Rock', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0951', 'Capsakid', '热辣娃', 'Grass', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0952', 'Scovillain', '狠辣椒', 'Grass,Fire', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0953', 'Rellor', '虫滚泥', 'Bug', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0954', 'Rabsca', '虫甲圣', 'Bug,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0955', 'Flittle', '飘飘雏', 'Psychic', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0956', 'Espathra', '超能艳鸵', 'Psychic', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0957', 'Tinkatink', '小锻匠', 'Fairy,Steel', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0958', 'Tinkatuff', '巧锻匠', 'Fairy,Steel', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0959', 'Tinkaton', '巨锻匠', 'Fairy,Steel', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0960', 'Wiglett', '海地鼠', 'Water', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0961', 'Wugtrio', '三海地鼠', 'Water', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0962', 'Bombirdier', '下石鸟', 'Flying,Dark', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0963', 'Finizen', '波普海豚', 'Water', 200);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0964', 'Palafin', '海豚侠', 'Water', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0965', 'Varoom', '噗隆隆', 'Steel,Poison', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0966', 'Revavroom', '普隆隆姆', 'Steel,Poison', 75);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0967', 'Cyclizar', '摩托蜥', 'Dragon,Normal', 190);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0968', 'Orthworm', '拖拖蚓', 'Steel', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0969', 'Glimmet', '晶光芽', 'Rock,Poison', 70);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0970', 'Glimmora', '晶光花', 'Rock,Poison', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0971', 'Greavard', '墓仔狗', 'Ghost', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0972', 'Houndstone', '墓扬犬', 'Ghost', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0973', 'Flamigo', '缠红鹤', 'Flying,Fighting', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0974', 'Cetoddle', '走鲸', 'Ice', 150);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0975', 'Cetitan', '浩大鲸', 'Ice', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0976', 'Veluza', '轻身鳕', 'Water,Psychic', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0977', 'Dondozo', '吃吼霸', 'Water', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0978', 'Tatsugiri', '米立龙', 'Dragon,Water', 100);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0979', 'Annihilape', '弃世猴', 'Fighting,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0980', 'Clodsire', '土王', 'Poison,Ground', 90);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0981', 'Farigiraf', '奇麒麟', 'Normal,Psychic', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0982', 'Dudunsparce', '土龙节节', 'Normal', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0983', 'Kingambit', '仆斩将军', '', NULL);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0984', 'Great Tusk', '雄伟牙', 'Ground,Fighting', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0985', 'Scream Tail', '吼叫尾', 'Fairy,Psychic', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0986', 'Brute Bonnet', '猛恶菇', 'Grass,Dark', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0987', 'Flutter Mane', '振翼发', 'Ghost,Fairy', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0988', 'Slither Wing', '爬地翅', 'Bug,Fighting', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0989', 'Sandy Shocks', '沙铁皮', 'Electric,Ground', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0990', 'Iron Treads', '铁辙迹', 'Ground,Steel', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0991', 'Iron Bundle', '铁包袱', 'Ice,Water', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0992', 'Iron Hands', '铁臂膀', 'Fighting,Electric', 50);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0993', 'Iron Jugulis', '铁脖颈', 'Dark,Flying', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0994', 'Iron Moth', '铁毒蛾', 'Fire,Poison', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0995', 'Iron Thorns', '铁荆棘', 'Rock,Electric', 30);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0996', 'Frigibax', '凉脊龙', 'Dragon,Ice', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0997', 'Arctibax', '冻脊龙', 'Dragon,Ice', 25);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0998', 'Baxcalibur', '戟脊龙', 'Dragon,Ice', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('0999', 'Gimmighoul', '索财灵', 'Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1000', 'Gholdengo', '赛富豪', 'Steel,Ghost', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1001', 'Wo-Chien', '古简蜗', 'Dark,Grass', 6);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1002', 'Chien-Pao', '古剑豹', 'Dark,Ice', 6);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1003', 'Ting-Lu', '古鼎鹿', 'Dark,Ground', 6);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1004', 'Chi-Yu', '古玉鱼', 'Dark,Fire', 6);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1005', 'Roaring Moon', '轰鸣月', 'Dragon,Dark', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1006', 'Iron Valiant', '铁武者', 'Fairy,Fighting', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1007', 'Koraidon', '故勒顿', 'Fighting,Dragon', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1008', 'Miraidon', '密勒顿', 'Electric,Dragon', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1009', 'Walking Wake', '波荡水', 'Water,Dragon', 5);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1010', 'Iron Leaves', '铁斑叶', 'Grass,Psychic', 5);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1011', 'Dipplin', '裹蜜虫', 'Grass,Dragon', 45);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1012', 'Poltchageist', '斯魔茶', 'Grass,Ghost', 120);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1013', 'Sinistcha', '来悲粗茶', 'Grass,Ghost', 60);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1014', 'Okidogi', '够赞狗', 'Poison,Fighting', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1015', 'Munkidori', '愿增猿', 'Poison,Psychic', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1016', 'Fezandipiti', '吉雉鸡', 'Poison,Fairy', 3);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1017', 'Ogerpon', '厄诡椪', 'Grass', 5);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1018', 'Archaludon', '铝钢桥龙', 'Steel,Dragon', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1019', 'Hydrapple', '蜜集大蛇', 'Grass,Dragon', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1020', 'Gouging Fire', '破空焰', 'Fire,Dragon', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1021', 'Raging Bolt', '猛雷鼓', 'Electric,Dragon', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1022', 'Iron Boulder', '铁磐岩', 'Rock,Psychic', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1023', 'Iron Crown', '铁头壳', 'Steel,Psychic', 10);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1024', 'Terapagos', '太乐巴戈斯', 'Normal', 255);
INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ('1025', 'Pecharunt', '桃歹郎', 'Poison,Ghost', 3);
-- Pokemon Moves Data
CREATE TABLE IF NOT EXISTS pokemon_moves (
    move_id VARCHAR(10) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    type VARCHAR(50),
    category VARCHAR(50),
    power INT,
    accuracy INT,
    pp INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('1', 'Pound', '拍击', 'Normal', 'Physical', 40, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('2', 'Karate Chop', '空手劈', 'Fighting', 'Physical', 50, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('3', 'Double Slap', '连环巴掌', 'Normal', 'Physical', 15, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('4', 'Comet Punch', '连续拳', 'Normal', 'Physical', 18, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('5', 'Mega Punch', '百万吨重拳', 'Normal', 'Physical', 80, 85, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('6', 'Pay Day', '聚宝功', 'Normal', 'Physical', 40, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('7', 'Fire Punch', '火焰拳', 'Fire', 'Physical', 75, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('8', 'Ice Punch', '冰冻拳', 'Ice', 'Physical', 75, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('9', 'Thunder Punch', '雷电拳', 'Electric', 'Physical', 75, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('10', 'Scratch', '抓', 'Normal', 'Physical', 40, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('11', 'Vise Grip', '夹住', 'Normal', 'Physical', 55, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('12', 'Guillotine', '极落钳', 'Normal', 'Physical', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('13', 'Razor Wind', '旋风刀', 'Normal', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('14', 'Swords Dance', '剑舞', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('15', 'Cut', '居合劈', 'Normal', 'Physical', 50, 95, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('16', 'Gust', '起风', 'Flying', 'Special', 40, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('17', 'Wing Attack', '翅膀攻击', 'Flying', 'Physical', 60, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('18', 'Whirlwind', '吹飞', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('19', 'Fly', '飞翔', 'Flying', 'Physical', 90, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('20', 'Bind', '绑紧', 'Normal', 'Physical', 15, 85, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('21', 'Slam', '摔打', 'Normal', 'Physical', 80, 75, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('22', 'Vine Whip', '藤鞭', 'Grass', 'Physical', 45, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('23', 'Stomp', '踩踏', 'Normal', 'Physical', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('24', 'Double Kick', '二连踢', 'Fighting', 'Physical', 30, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('25', 'Mega Kick', '百万吨重踢', 'Normal', 'Physical', 120, 75, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('26', 'Jump Kick', '飞踢', 'Fighting', 'Physical', 100, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('27', 'Rolling Kick', '回旋踢', 'Fighting', 'Physical', 60, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('28', 'Sand Attack', '泼沙', 'Ground', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('29', 'Headbutt', '头锤', 'Normal', 'Physical', 70, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('30', 'Horn Attack', '角撞', 'Normal', 'Physical', 65, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('31', 'Fury Attack', '乱击', 'Normal', 'Physical', 15, 85, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('32', 'Horn Drill', '角钻', 'Normal', 'Physical', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('33', 'Tackle', '撞击', 'Normal', 'Physical', 40, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('34', 'Body Slam', '泰山压顶', 'Normal', 'Physical', 85, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('35', 'Wrap', '紧束', 'Normal', 'Physical', 15, 90, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('36', 'Take Down', '猛撞', 'Normal', 'Physical', 90, 85, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('37', 'Thrash', '大闹一番', 'Normal', 'Physical', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('38', 'Double-Edge', '舍身冲撞', 'Normal', 'Physical', 120, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('39', 'Tail Whip', '摇尾巴', 'Normal', 'Status', NULL, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('40', 'Poison Sting', '毒针', 'Poison', 'Physical', 15, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('41', 'Twineedle', '双针', 'Bug', 'Physical', 25, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('42', 'Pin Missile', '飞弹针', 'Bug', 'Physical', 25, 95, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('43', 'Leer', '瞪眼', 'Normal', 'Status', NULL, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('44', 'Bite', '咬住', 'Dark', 'Physical', 60, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('45', 'Growl', '叫声', 'Normal', 'Status', NULL, 100, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('46', 'Roar', '吼叫', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('47', 'Sing', '唱歌', 'Normal', 'Status', NULL, 55, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('48', 'Supersonic', '超音波', 'Normal', 'Status', NULL, 55, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('49', 'Sonic Boom', '音爆', 'Normal', 'Special', NULL, 90, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('50', 'Disable', '定身法', 'Normal', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('51', 'Acid', '溶解液', 'Poison', 'Special', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('52', 'Ember', '火花', 'Fire', 'Special', 40, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('53', 'Flamethrower', '喷射火焰', 'Fire', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('54', 'Mist', '白雾', 'Ice', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('55', 'Water Gun', '水枪', 'Water', 'Special', 40, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('56', 'Hydro Pump', '水炮', 'Water', 'Special', 110, 80, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('57', 'Surf', '冲浪', 'Water', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('58', 'Ice Beam', '冰冻光束', 'Ice', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('59', 'Blizzard', '暴风雪', 'Ice', 'Special', 110, 70, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('60', 'Psybeam', '幻象光线', 'Psychic', 'Special', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('61', 'Bubble Beam', '泡沫光线', 'Water', 'Special', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('62', 'Aurora Beam', '极光束', 'Ice', 'Special', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('63', 'Hyper Beam', '破坏光线', 'Normal', 'Special', 150, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('64', 'Peck', '啄', 'Flying', 'Physical', 35, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('65', 'Drill Peck', '啄钻', 'Flying', 'Physical', 80, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('66', 'Submission', '深渊翻滚', 'Fighting', 'Physical', 80, 80, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('67', 'Low Kick', '踢倒', 'Fighting', 'Physical', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('68', 'Counter', '双倍奉还', 'Fighting', 'Physical', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('69', 'Seismic Toss', '地球上投', 'Fighting', 'Physical', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('70', 'Strength', '怪力', 'Normal', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('71', 'Absorb', '吸取', 'Grass', 'Special', 20, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('72', 'Mega Drain', '超级吸取', 'Grass', 'Special', 40, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('73', 'Leech Seed', '寄生种子', 'Grass', 'Status', NULL, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('74', 'Growth', '生长', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('75', 'Razor Leaf', '飞叶快刀', 'Grass', 'Physical', 55, 95, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('76', 'Solar Beam', '日光束', 'Grass', 'Special', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('77', 'Poison Powder', '毒粉', 'Poison', 'Status', NULL, 75, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('78', 'Stun Spore', '麻痹粉', 'Grass', 'Status', NULL, 75, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('79', 'Sleep Powder', '催眠粉', 'Grass', 'Status', NULL, 75, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('80', 'Petal Dance', '花瓣舞', 'Grass', 'Special', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('81', 'String Shot', '吐丝', 'Bug', 'Status', NULL, 95, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('82', 'Dragon Rage', '龙之怒', 'Dragon', 'Special', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('83', 'Fire Spin', '火焰旋涡', 'Fire', 'Special', 35, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('84', 'Thunder Shock', '电击', 'Electric', 'Special', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('85', 'Thunderbolt', '十万伏特', 'Electric', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('86', 'Thunder Wave', '电磁波', 'Electric', 'Status', NULL, 90, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('87', 'Thunder', '打雷', 'Electric', 'Special', 110, 70, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('88', 'Rock Throw', '落石', 'Rock', 'Physical', 50, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('89', 'Earthquake', '地震', 'Ground', 'Physical', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('90', 'Fissure', '地裂', 'Ground', 'Physical', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('91', 'Dig', '挖洞', 'Ground', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('92', 'Toxic', '剧毒', 'Poison', 'Status', NULL, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('93', 'Confusion', '念力', 'Psychic', 'Special', 50, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('94', 'Psychic', '精神强念', 'Psychic', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('95', 'Hypnosis', '催眠术', 'Psychic', 'Status', NULL, 60, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('96', 'Meditate', '瑜伽姿势', 'Psychic', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('97', 'Agility', '高速移动', 'Psychic', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('98', 'Quick Attack', '电光一闪', 'Normal', 'Physical', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('99', 'Rage', '愤怒', 'Normal', 'Physical', 20, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('100', 'Teleport', '瞬间移动', 'Psychic', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('101', 'Night Shade', '黑夜魔影', 'Ghost', 'Special', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('102', 'Mimic', '模仿', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('103', 'Screech', '刺耳声', 'Normal', 'Status', NULL, 85, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('104', 'Double Team', '影子分身', 'Normal', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('105', 'Recover', '自我再生', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('106', 'Harden', '变硬', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('107', 'Minimize', '变小', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('108', 'Smokescreen', '烟幕', 'Normal', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('109', 'Confuse Ray', '奇异之光', 'Ghost', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('110', 'Withdraw', '缩入壳中', 'Water', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('111', 'Defense Curl', '变圆', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('112', 'Barrier', '屏障', 'Psychic', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('113', 'Light Screen', '光墙', 'Psychic', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('114', 'Haze', '黑雾', 'Ice', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('115', 'Reflect', '反射壁', 'Psychic', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('116', 'Focus Energy', '聚气', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('117', 'Bide', '忍耐', 'Normal', 'Physical', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('118', 'Metronome', '挥指', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('119', 'Mirror Move', '鹦鹉学舌', 'Flying', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('120', 'Self-Destruct', '玉石俱碎', 'Normal', 'Physical', 200, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('121', 'Egg Bomb', '炸蛋', 'Normal', 'Physical', 100, 75, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('122', 'Lick', '舌舔', 'Ghost', 'Physical', 30, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('123', 'Smog', '浊雾', 'Poison', 'Special', 30, 70, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('124', 'Sludge', '污泥攻击', 'Poison', 'Special', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('125', 'Bone Club', '骨棒', 'Ground', 'Physical', 65, 85, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('126', 'Fire Blast', '大字爆炎', 'Fire', 'Special', 110, 85, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('127', 'Waterfall', '攀瀑', 'Water', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('128', 'Clamp', '贝壳夹击', 'Water', 'Physical', 35, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('129', 'Swift', '高速星星', 'Normal', 'Special', 60, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('130', 'Skull Bash', '火箭头锤', 'Normal', 'Physical', 130, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('131', 'Spike Cannon', '尖刺加农炮', 'Normal', 'Physical', 20, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('132', 'Constrict', '缠绕', 'Normal', 'Physical', 10, 100, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('133', 'Amnesia', '瞬间失忆', 'Psychic', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('134', 'Kinesis', '折弯汤匙', 'Psychic', 'Status', NULL, 80, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('135', 'Soft-Boiled', '生蛋', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('136', 'High Jump Kick', '飞膝踢', 'Fighting', 'Physical', 130, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('137', 'Glare', '大蛇瞪眼', 'Normal', 'Status', NULL, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('138', 'Dream Eater', '食梦', 'Psychic', 'Special', 100, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('139', 'Poison Gas', '毒瓦斯', 'Poison', 'Status', NULL, 90, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('140', 'Barrage', '投球', 'Normal', 'Physical', 15, 85, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('141', 'Leech Life', '汲取', 'Bug', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('142', 'Lovely Kiss', '恶魔之吻', 'Normal', 'Status', NULL, 75, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('143', 'Sky Attack', '神鸟猛击', 'Flying', 'Physical', 140, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('144', 'Transform', '变身', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('145', 'Bubble', '泡沫', 'Water', 'Special', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('146', 'Dizzy Punch', '迷昏拳', 'Normal', 'Physical', 70, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('147', 'Spore', '蘑菇孢子', 'Grass', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('148', 'Flash', '闪光', 'Normal', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('149', 'Psywave', '精神波', 'Psychic', 'Special', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('150', 'Splash', '跃起', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('151', 'Acid Armor', '溶化', 'Poison', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('152', 'Crabhammer', '蟹钳锤', 'Water', 'Physical', 100, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('153', 'Explosion', '大爆炸', 'Normal', 'Physical', 250, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('154', 'Fury Swipes', '乱抓', 'Normal', 'Physical', 18, 80, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('155', 'Bonemerang', '骨头回力镖', 'Ground', 'Physical', 50, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('156', 'Rest', '睡觉', 'Psychic', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('157', 'Rock Slide', '岩崩', 'Rock', 'Physical', 75, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('158', 'Hyper Fang', '终结门牙', 'Normal', 'Physical', 80, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('159', 'Sharpen', '棱角化', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('160', 'Conversion', '纹理', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('161', 'Tri Attack', '三重攻击', 'Normal', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('162', 'Super Fang', '愤怒门牙', 'Normal', 'Physical', NULL, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('163', 'Slash', '劈开', 'Normal', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('164', 'Substitute', '替身', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('165', 'Struggle', '挣扎', 'Normal', 'Physical', 50, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('166', 'Sketch', '写生', 'Normal', 'Status', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('167', 'Triple Kick', '三连踢', 'Fighting', 'Physical', 10, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('168', 'Thief', '小偷', 'Dark', 'Physical', 60, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('169', 'Spider Web', '蛛网', 'Bug', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('170', 'Mind Reader', '心之眼', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('171', 'Nightmare', '恶梦', 'Ghost', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('172', 'Flame Wheel', '火焰轮', 'Fire', 'Physical', 60, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('173', 'Snore', '打鼾', 'Normal', 'Special', 50, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('174', 'Curse', '咒术', 'Ghost', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('175', 'Flail', '抓狂', 'Normal', 'Physical', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('176', 'Conversion 2', '纹理２', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('177', 'Aeroblast', '气旋攻击', 'Flying', 'Special', 100, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('178', 'Cotton Spore', '棉孢子', 'Grass', 'Status', NULL, 100, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('179', 'Reversal', '绝处逢生', 'Fighting', 'Physical', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('180', 'Spite', '怨恨', 'Ghost', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('181', 'Powder Snow', '细雪', 'Ice', 'Special', 40, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('182', 'Protect', '守住', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('183', 'Mach Punch', '音速拳', 'Fighting', 'Physical', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('184', 'Scary Face', '可怕面孔', 'Normal', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('185', 'Feint Attack', '出奇一击', 'Dark', 'Physical', 60, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('186', 'Sweet Kiss', '天使之吻', 'Fairy', 'Status', NULL, 75, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('187', 'Belly Drum', '腹鼓', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('188', 'Sludge Bomb', '污泥炸弹', 'Poison', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('189', 'Mud-Slap', '掷泥', 'Ground', 'Special', 20, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('190', 'Octazooka', '章鱼桶炮', 'Water', 'Special', 65, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('191', 'Spikes', '撒菱', 'Ground', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('192', 'Zap Cannon', '电磁炮', 'Electric', 'Special', 120, 50, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('193', 'Foresight', '识破', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('194', 'Destiny Bond', '同命', 'Ghost', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('195', 'Perish Song', '终焉之歌', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('196', 'Icy Wind', '冰冻之风', 'Ice', 'Special', 55, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('197', 'Detect', '看穿', 'Fighting', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('198', 'Bone Rush', '骨棒乱打', 'Ground', 'Physical', 25, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('199', 'Lock-On', '锁定', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('200', 'Outrage', '逆鳞', 'Dragon', 'Physical', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('201', 'Sandstorm', '沙暴', 'Rock', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('202', 'Giga Drain', '终极吸取', 'Grass', 'Special', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('203', 'Endure', '挺住', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('204', 'Charm', '撒娇', 'Fairy', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('205', 'Rollout', '滚动', 'Rock', 'Physical', 30, 90, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('206', 'False Swipe', '点到为止', 'Normal', 'Physical', 40, 100, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('207', 'Swagger', '虚张声势', 'Normal', 'Status', NULL, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('208', 'Milk Drink', '喝牛奶', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('209', 'Spark', '电光', 'Electric', 'Physical', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('210', 'Fury Cutter', '连斩', 'Bug', 'Physical', 40, 95, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('211', 'Steel Wing', '钢翼', 'Steel', 'Physical', 70, 90, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('212', 'Mean Look', '黑色目光', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('213', 'Attract', '迷人', 'Normal', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('214', 'Sleep Talk', '梦话', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('215', 'Heal Bell', '治愈铃声', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('216', 'Return', '报恩', 'Normal', 'Physical', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('217', 'Present', '礼物', 'Normal', 'Physical', NULL, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('218', 'Frustration', '迁怒', 'Normal', 'Physical', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('219', 'Safeguard', '神秘守护', 'Normal', 'Status', NULL, NULL, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('220', 'Pain Split', '分担痛楚', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('221', 'Sacred Fire', '神圣之火', 'Fire', 'Physical', 100, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('222', 'Magnitude', '震级', 'Ground', 'Physical', NULL, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('223', 'Dynamic Punch', '爆裂拳', 'Fighting', 'Physical', 100, 50, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('224', 'Megahorn', '超级角击', 'Bug', 'Physical', 120, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('225', 'Dragon Breath', '龙息', 'Dragon', 'Special', 60, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('226', 'Baton Pass', '接棒', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('227', 'Encore', '再来一次', 'Normal', 'Status', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('228', 'Pursuit', '追打', 'Dark', 'Physical', 40, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('229', 'Rapid Spin', '高速旋转', 'Normal', 'Physical', 50, 100, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('230', 'Sweet Scent', '甜甜香气', 'Normal', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('231', 'Iron Tail', '铁尾', 'Steel', 'Physical', 100, 75, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('232', 'Metal Claw', '金属爪', 'Steel', 'Physical', 50, 95, 35);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('233', 'Vital Throw', '借力摔', 'Fighting', 'Physical', 70, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('234', 'Morning Sun', '晨光', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('235', 'Synthesis', '光合作用', 'Grass', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('236', 'Moonlight', '月光', 'Fairy', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('237', 'Hidden Power', '觉醒力量', 'Normal', 'Special', 60, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('238', 'Cross Chop', '十字劈', 'Fighting', 'Physical', 100, 80, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('239', 'Twister', '龙卷风', 'Dragon', 'Special', 40, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('240', 'Rain Dance', '求雨', 'Water', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('241', 'Sunny Day', '大晴天', 'Fire', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('242', 'Crunch', '咬碎', 'Dark', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('243', 'Mirror Coat', '镜面反射', 'Psychic', 'Special', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('244', 'Psych Up', '自我暗示', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('245', 'Extreme Speed', '神速', 'Normal', 'Physical', 80, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('246', 'Ancient Power', '原始之力', 'Rock', 'Special', 60, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('247', 'Shadow Ball', '暗影球', 'Ghost', 'Special', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('248', 'Future Sight', '预知未来', 'Psychic', 'Special', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('249', 'Rock Smash', '碎岩', 'Fighting', 'Physical', 40, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('250', 'Whirlpool', '潮旋', 'Water', 'Special', 35, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('251', 'Beat Up', '围攻', 'Dark', 'Physical', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('252', 'Fake Out', '击掌奇袭', 'Normal', 'Physical', 40, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('253', 'Uproar', '吵闹', 'Normal', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('254', 'Stockpile', '蓄力', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('255', 'Spit Up', '喷出', 'Normal', 'Special', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('256', 'Swallow', '吞下', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('257', 'Heat Wave', '热风', 'Fire', 'Special', 95, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('258', 'Hail', '冰雹', 'Ice', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('259', 'Torment', '无理取闹', 'Dark', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('260', 'Flatter', '吹捧', 'Dark', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('261', 'Will-O-Wisp', '磷火', 'Fire', 'Status', NULL, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('262', 'Memento', '临别礼物', 'Dark', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('263', 'Facade', '硬撑', 'Normal', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('264', 'Focus Punch', '真气拳', 'Fighting', 'Physical', 150, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('265', 'Smelling Salts', '清醒', 'Normal', 'Physical', 70, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('266', 'Follow Me', '看我嘛', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('267', 'Nature Power', '自然之力', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('268', 'Charge', '充电', 'Electric', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('269', 'Taunt', '挑衅', 'Dark', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('270', 'Helping Hand', '帮助', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('271', 'Trick', '戏法', 'Psychic', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('272', 'Role Play', '扮演', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('273', 'Wish', '祈愿', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('274', 'Assist', '借助', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('275', 'Ingrain', '扎根', 'Grass', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('276', 'Superpower', '蛮力', 'Fighting', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('277', 'Magic Coat', '魔法反射', 'Psychic', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('278', 'Recycle', '回收利用', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('279', 'Revenge', '报复', 'Fighting', 'Physical', 60, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('280', 'Brick Break', '劈瓦', 'Fighting', 'Physical', 75, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('281', 'Yawn', '哈欠', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('282', 'Knock Off', '拍落', 'Dark', 'Physical', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('283', 'Endeavor', '蛮干', 'Normal', 'Physical', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('284', 'Eruption', '喷火', 'Fire', 'Special', 150, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('285', 'Skill Swap', '特性互换', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('286', 'Imprison', '封印', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('287', 'Refresh', '焕然一新', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('288', 'Grudge', '怨念', 'Ghost', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('289', 'Snatch', '化为己用', 'Dark', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('290', 'Secret Power', '秘密之力', 'Normal', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('291', 'Dive', '潜水', 'Water', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('292', 'Arm Thrust', '猛推', 'Fighting', 'Physical', 15, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('293', 'Camouflage', '保护色', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('294', 'Tail Glow', '萤火', 'Bug', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('295', 'Luster Purge', '洁净光芒', 'Psychic', 'Special', 95, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('296', 'Mist Ball', '薄雾球', 'Psychic', 'Special', 95, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('297', 'Feather Dance', '羽毛舞', 'Flying', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('298', 'Teeter Dance', '摇晃舞', 'Normal', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('299', 'Blaze Kick', '火焰踢', 'Fire', 'Physical', 85, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('300', 'Mud Sport', '玩泥巴', 'Ground', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('301', 'Ice Ball', '冰球', 'Ice', 'Physical', 30, 90, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('302', 'Needle Arm', '尖刺臂', 'Grass', 'Physical', 60, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('303', 'Slack Off', '偷懒', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('304', 'Hyper Voice', '巨声', 'Normal', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('305', 'Poison Fang', '剧毒牙', 'Poison', 'Physical', 50, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('306', 'Crush Claw', '撕裂爪', 'Normal', 'Physical', 75, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('307', 'Blast Burn', '爆炸烈焰', 'Fire', 'Special', 150, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('308', 'Hydro Cannon', '加农水炮', 'Water', 'Special', 150, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('309', 'Meteor Mash', '彗星拳', 'Steel', 'Physical', 90, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('310', 'Astonish', '惊吓', 'Ghost', 'Physical', 30, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('311', 'Weather Ball', '气象球', 'Normal', 'Special', 50, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('312', 'Aromatherapy', '芳香治疗', 'Grass', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('313', 'Fake Tears', '假哭', 'Dark', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('314', 'Air Cutter', '空气利刃', 'Flying', 'Special', 60, 95, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('315', 'Overheat', '过热', 'Fire', 'Special', 130, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('316', 'Odor Sleuth', '气味侦测', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('317', 'Rock Tomb', '岩石封锁', 'Rock', 'Physical', 60, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('318', 'Silver Wind', '银色旋风', 'Bug', 'Special', 60, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('319', 'Metal Sound', '金属音', 'Steel', 'Status', NULL, 85, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('320', 'Grass Whistle', '草笛', 'Grass', 'Status', NULL, 55, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('321', 'Tickle', '挠痒', 'Normal', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('322', 'Cosmic Power', '宇宙力量', 'Psychic', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('323', 'Water Spout', '喷水', 'Water', 'Special', 150, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('324', 'Signal Beam', '信号光束', 'Bug', 'Special', 75, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('325', 'Shadow Punch', '暗影拳', 'Ghost', 'Physical', 60, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('326', 'Extrasensory', '神通力', 'Psychic', 'Special', 80, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('327', 'Sky Uppercut', '冲天拳', 'Fighting', 'Physical', 85, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('328', 'Sand Tomb', '流沙深渊', 'Ground', 'Physical', 35, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('329', 'Sheer Cold', '绝对零度', 'Ice', 'Special', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('330', 'Muddy Water', '浊流', 'Water', 'Special', 90, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('331', 'Bullet Seed', '种子机关枪', 'Grass', 'Physical', 25, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('332', 'Aerial Ace', '燕返', 'Flying', 'Physical', 60, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('333', 'Icicle Spear', '冰锥', 'Ice', 'Physical', 25, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('334', 'Iron Defense', '铁壁', 'Steel', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('335', 'Block', '挡路', 'Normal', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('336', 'Howl', '长嚎', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('337', 'Dragon Claw', '龙爪', 'Dragon', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('338', 'Frenzy Plant', '疯狂植物', 'Grass', 'Special', 150, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('339', 'Bulk Up', '健美', 'Fighting', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('340', 'Bounce', '弹跳', 'Flying', 'Physical', 85, 85, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('341', 'Mud Shot', '泥巴射击', 'Ground', 'Special', 55, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('342', 'Poison Tail', '毒尾', 'Poison', 'Physical', 50, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('343', 'Covet', '渴望', 'Normal', 'Physical', 60, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('344', 'Volt Tackle', '伏特攻击', 'Electric', 'Physical', 120, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('345', 'Magical Leaf', '魔法叶', 'Grass', 'Special', 60, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('346', 'Water Sport', '玩水', 'Water', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('347', 'Calm Mind', '冥想', 'Psychic', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('348', 'Leaf Blade', '叶刃', 'Grass', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('349', 'Dragon Dance', '龙之舞', 'Dragon', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('350', 'Rock Blast', '岩石爆击', 'Rock', 'Physical', 25, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('351', 'Shock Wave', '电击波', 'Electric', 'Special', 60, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('352', 'Water Pulse', '水之波动', 'Water', 'Special', 60, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('353', 'Doom Desire', '破灭之愿', 'Steel', 'Special', 140, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('354', 'Psycho Boost', '精神突进', 'Psychic', 'Special', 140, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('355', 'Roost', '羽栖', 'Flying', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('356', 'Gravity', '重力', 'Psychic', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('357', 'Miracle Eye', '奇迹之眼', 'Psychic', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('358', 'Wake-Up Slap', '唤醒巴掌', 'Fighting', 'Physical', 70, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('359', 'Hammer Arm', '臂锤', 'Fighting', 'Physical', 100, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('360', 'Gyro Ball', '陀螺球', 'Steel', 'Physical', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('361', 'Healing Wish', '治愈之愿', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('362', 'Brine', '盐水', 'Water', 'Special', 65, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('363', 'Natural Gift', '自然之恩', 'Normal', 'Physical', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('364', 'Feint', '佯攻', 'Normal', 'Physical', 30, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('365', 'Pluck', '啄食', 'Flying', 'Physical', 60, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('366', 'Tailwind', '顺风', 'Flying', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('367', 'Acupressure', '点穴', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('368', 'Metal Burst', '金属爆炸', 'Steel', 'Physical', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('369', 'U-turn', '急速折返', 'Bug', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('370', 'Close Combat', '近身战', 'Fighting', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('371', 'Payback', '以牙还牙', 'Dark', 'Physical', 50, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('372', 'Assurance', '恶意追击', 'Dark', 'Physical', 60, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('373', 'Embargo', '查封', 'Dark', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('374', 'Fling', '投掷', 'Dark', 'Physical', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('375', 'Psycho Shift', '精神转移', 'Psychic', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('376', 'Trump Card', '王牌', 'Normal', 'Special', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('377', 'Heal Block', '回复封锁', 'Psychic', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('378', 'Wring Out', '绞紧', 'Normal', 'Special', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('379', 'Power Trick', '力量戏法', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('380', 'Gastro Acid', '胃液', 'Poison', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('381', 'Lucky Chant', '幸运咒语', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('382', 'Me First', '抢先一步', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('383', 'Copycat', '仿效', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('384', 'Power Swap', '力量互换', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('385', 'Guard Swap', '防守互换', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('386', 'Punishment', '惩罚', 'Dark', 'Physical', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('387', 'Last Resort', '珍藏', 'Normal', 'Physical', 140, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('388', 'Worry Seed', '烦恼种子', 'Grass', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('389', 'Sucker Punch', '突袭', 'Dark', 'Physical', 70, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('390', 'Toxic Spikes', '毒菱', 'Poison', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('391', 'Heart Swap', '心灵互换', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('392', 'Aqua Ring', '水流环', 'Water', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('393', 'Magnet Rise', '电磁飘浮', 'Electric', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('394', 'Flare Blitz', '闪焰冲锋', 'Fire', 'Physical', 120, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('395', 'Force Palm', '发劲', 'Fighting', 'Physical', 60, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('396', 'Aura Sphere', '波导弹', 'Fighting', 'Special', 80, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('397', 'Rock Polish', '岩石打磨', 'Rock', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('398', 'Poison Jab', '毒击', 'Poison', 'Physical', 80, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('399', 'Dark Pulse', '恶之波动', 'Dark', 'Special', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('400', 'Night Slash', '暗袭要害', 'Dark', 'Physical', 70, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('401', 'Aqua Tail', '水流尾', 'Water', 'Physical', 90, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('402', 'Seed Bomb', '种子炸弹', 'Grass', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('403', 'Air Slash', '空气之刃', 'Flying', 'Special', 75, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('404', 'X-Scissor', '十字剪', 'Bug', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('405', 'Bug Buzz', '虫鸣', 'Bug', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('406', 'Dragon Pulse', '龙之波动', 'Dragon', 'Special', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('407', 'Dragon Rush', '龙之俯冲', 'Dragon', 'Physical', 100, 75, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('408', 'Power Gem', '力量宝石', 'Rock', 'Special', 80, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('409', 'Drain Punch', '吸取拳', 'Fighting', 'Physical', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('410', 'Vacuum Wave', '真空波', 'Fighting', 'Special', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('411', 'Focus Blast', '真气弹', 'Fighting', 'Special', 120, 70, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('412', 'Energy Ball', '能量球', 'Grass', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('413', 'Brave Bird', '勇鸟猛攻', 'Flying', 'Physical', 120, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('414', 'Earth Power', '大地之力', 'Ground', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('415', 'Switcheroo', '掉包', 'Dark', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('416', 'Giga Impact', '终极冲击', 'Normal', 'Physical', 150, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('417', 'Nasty Plot', '诡计', 'Dark', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('418', 'Bullet Punch', '子弹拳', 'Steel', 'Physical', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('419', 'Avalanche', '雪崩', 'Ice', 'Physical', 60, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('420', 'Ice Shard', '冰砾', 'Ice', 'Physical', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('421', 'Shadow Claw', '暗影爪', 'Ghost', 'Physical', 70, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('422', 'Thunder Fang', '雷电牙', 'Electric', 'Physical', 65, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('423', 'Ice Fang', '冰冻牙', 'Ice', 'Physical', 65, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('424', 'Fire Fang', '火焰牙', 'Fire', 'Physical', 65, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('425', 'Shadow Sneak', '影子偷袭', 'Ghost', 'Physical', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('426', 'Mud Bomb', '泥巴炸弹', 'Ground', 'Special', 65, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('427', 'Psycho Cut', '精神利刃', 'Psychic', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('428', 'Zen Headbutt', '意念头锤', 'Psychic', 'Physical', 80, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('429', 'Mirror Shot', '镜光射击', 'Steel', 'Special', 65, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('430', 'Flash Cannon', '加农光炮', 'Steel', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('431', 'Rock Climb', '攀岩', 'Normal', 'Physical', 90, 85, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('432', 'Defog', '清除浓雾', 'Flying', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('433', 'Trick Room', '戏法空间', 'Psychic', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('434', 'Draco Meteor', '流星群', 'Dragon', 'Special', 130, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('435', 'Discharge', '放电', 'Electric', 'Special', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('436', 'Lava Plume', '喷烟', 'Fire', 'Special', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('437', 'Leaf Storm', '飞叶风暴', 'Grass', 'Special', 130, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('438', 'Power Whip', '强力鞭打', 'Grass', 'Physical', 120, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('439', 'Rock Wrecker', '岩石炮', 'Rock', 'Physical', 150, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('440', 'Cross Poison', '十字毒刃', 'Poison', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('441', 'Gunk Shot', '垃圾射击', 'Poison', 'Physical', 120, 80, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('442', 'Iron Head', '铁头', 'Steel', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('443', 'Magnet Bomb', '磁铁炸弹', 'Steel', 'Physical', 60, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('444', 'Stone Edge', '尖石攻击', 'Rock', 'Physical', 100, 80, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('445', 'Captivate', '诱惑', 'Normal', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('446', 'Stealth Rock', '隐形岩', 'Rock', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('447', 'Grass Knot', '打草结', 'Grass', 'Special', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('448', 'Chatter', '喋喋不休', 'Flying', 'Special', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('449', 'Judgment', '制裁光砾', 'Normal', 'Special', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('450', 'Bug Bite', '虫咬', 'Bug', 'Physical', 60, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('451', 'Charge Beam', '充电光束', 'Electric', 'Special', 50, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('452', 'Wood Hammer', '木槌', 'Grass', 'Physical', 120, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('453', 'Aqua Jet', '水流喷射', 'Water', 'Physical', 40, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('454', 'Attack Order', '攻击指令', 'Bug', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('455', 'Defend Order', '防御指令', 'Bug', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('456', 'Heal Order', '回复指令', 'Bug', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('457', 'Head Smash', '双刃头锤', 'Rock', 'Physical', 150, 80, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('458', 'Double Hit', '二连击', 'Normal', 'Physical', 35, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('459', 'Roar of Time', '时光咆哮', 'Dragon', 'Special', 150, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('460', 'Spacial Rend', '亚空裂斩', 'Dragon', 'Special', 100, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('461', 'Lunar Dance', '新月舞', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('462', 'Crush Grip', '捏碎', 'Normal', 'Physical', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('463', 'Magma Storm', '熔岩风暴', 'Fire', 'Special', 100, 75, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('464', 'Dark Void', '暗黑洞', 'Dark', 'Status', NULL, 50, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('465', 'Seed Flare', '种子闪光', 'Grass', 'Special', 120, 85, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('466', 'Ominous Wind', '奇异之风', 'Ghost', 'Special', 60, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('467', 'Shadow Force', '暗影潜袭', 'Ghost', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('468', 'Hone Claws', '磨爪', 'Dark', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('469', 'Wide Guard', '广域防守', 'Rock', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('470', 'Guard Split', '防守平分', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('471', 'Power Split', '力量平分', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('472', 'Wonder Room', '奇妙空间', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('473', 'Psyshock', '精神冲击', 'Psychic', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('474', 'Venoshock', '毒液冲击', 'Poison', 'Special', 65, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('475', 'Autotomize', '身体轻量化', 'Steel', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('476', 'Rage Powder', '愤怒粉', 'Bug', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('477', 'Telekinesis', '意念移物', 'Psychic', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('478', 'Magic Room', '魔法空间', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('479', 'Smack Down', '击落', 'Rock', 'Physical', 50, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('480', 'Storm Throw', '山岚摔', 'Fighting', 'Physical', 60, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('481', 'Flame Burst', '烈焰溅射', 'Fire', 'Special', 70, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('482', 'Sludge Wave', '污泥波', 'Poison', 'Special', 95, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('483', 'Quiver Dance', '蝶舞', 'Bug', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('484', 'Heavy Slam', '重磅冲撞', 'Steel', 'Physical', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('485', 'Synchronoise', '同步干扰', 'Psychic', 'Special', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('486', 'Electro Ball', '电球', 'Electric', 'Special', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('487', 'Soak', '浸水', 'Water', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('488', 'Flame Charge', '蓄能焰袭', 'Fire', 'Physical', 50, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('489', 'Coil', '盘蜷', 'Poison', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('490', 'Low Sweep', '下盘踢', 'Fighting', 'Physical', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('491', 'Acid Spray', '酸液炸弹', 'Poison', 'Special', 40, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('492', 'Foul Play', '移花接木', 'Dark', 'Physical', 95, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('493', 'Simple Beam', '单纯光束', 'Normal', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('494', 'Entrainment', '找伙伴', 'Normal', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('495', 'After You', '您先请', 'Normal', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('496', 'Round', '轮唱', 'Normal', 'Special', 60, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('497', 'Echoed Voice', '回声', 'Normal', 'Special', 40, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('498', 'Chip Away', '逐步击破', 'Normal', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('499', 'Clear Smog', '清除之烟', 'Poison', 'Special', 50, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('500', 'Stored Power', '辅助力量', 'Psychic', 'Special', 20, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('501', 'Quick Guard', '快速防守', 'Fighting', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('502', 'Ally Switch', '交换场地', 'Psychic', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('503', 'Scald', '热水', 'Water', 'Special', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('504', 'Shell Smash', '破壳', 'Normal', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('505', 'Heal Pulse', '治愈波动', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('506', 'Hex', '祸不单行', 'Ghost', 'Special', 65, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('507', 'Sky Drop', '自由落体', 'Flying', 'Physical', 60, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('508', 'Shift Gear', '换档', 'Steel', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('509', 'Circle Throw', '巴投', 'Fighting', 'Physical', 60, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('510', 'Incinerate', '烧净', 'Fire', 'Special', 60, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('511', 'Quash', '延后', 'Dark', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('512', 'Acrobatics', '杂技', 'Flying', 'Physical', 55, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('513', 'Reflect Type', '镜面属性', 'Normal', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('514', 'Retaliate', '报仇', 'Normal', 'Physical', 70, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('515', 'Final Gambit', '搏命', 'Fighting', 'Special', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('516', 'Bestow', '传递礼物', 'Normal', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('517', 'Inferno', '烈火深渊', 'Fire', 'Special', 100, 50, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('518', 'Water Pledge', '水之誓约', 'Water', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('519', 'Fire Pledge', '火之誓约', 'Fire', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('520', 'Grass Pledge', '草之誓约', 'Grass', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('521', 'Volt Switch', '伏特替换', 'Electric', 'Special', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('522', 'Struggle Bug', '虫之抵抗', 'Bug', 'Special', 50, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('523', 'Bulldoze', '重踏', 'Ground', 'Physical', 60, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('524', 'Frost Breath', '冰息', 'Ice', 'Special', 60, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('525', 'Dragon Tail', '龙尾', 'Dragon', 'Physical', 60, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('526', 'Work Up', '自我激励', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('527', 'Electroweb', '电网', 'Electric', 'Special', 55, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('528', 'Wild Charge', '疯狂伏特', 'Electric', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('529', 'Drill Run', '直冲钻', 'Ground', 'Physical', 80, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('530', 'Dual Chop', '二连劈', 'Dragon', 'Physical', 40, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('531', 'Heart Stamp', '爱心印章', 'Psychic', 'Physical', 60, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('532', 'Horn Leech', '木角', 'Grass', 'Physical', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('533', 'Sacred Sword', '圣剑', 'Fighting', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('534', 'Razor Shell', '贝壳刃', 'Water', 'Physical', 75, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('535', 'Heat Crash', '高温重压', 'Fire', 'Physical', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('536', 'Leaf Tornado', '青草搅拌器', 'Grass', 'Special', 65, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('537', 'Steamroller', '疯狂滚压', 'Bug', 'Physical', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('538', 'Cotton Guard', '棉花防守', 'Grass', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('539', 'Night Daze', '暗黑爆破', 'Dark', 'Special', 85, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('540', 'Psystrike', '精神击破', 'Psychic', 'Special', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('541', 'Tail Slap', '扫尾拍打', 'Normal', 'Physical', 25, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('542', 'Hurricane', '暴风', 'Flying', 'Special', 110, 70, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('543', 'Head Charge', '爆炸头突击', 'Normal', 'Physical', 120, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('544', 'Gear Grind', '齿轮飞盘', 'Steel', 'Physical', 50, 85, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('545', 'Searing Shot', '火焰弹', 'Fire', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('546', 'Techno Blast', '高科技光炮', 'Normal', 'Special', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('547', 'Relic Song', '古老之歌', 'Normal', 'Special', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('548', 'Secret Sword', '神秘之剑', 'Fighting', 'Special', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('549', 'Glaciate', '冰封世界', 'Ice', 'Special', 65, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('550', 'Bolt Strike', '雷击', 'Electric', 'Physical', 130, 85, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('551', 'Blue Flare', '青焰', 'Fire', 'Special', 130, 85, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('552', 'Fiery Dance', '火之舞', 'Fire', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('553', 'Freeze Shock', '冰冻伏特', 'Ice', 'Physical', 140, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('554', 'Ice Burn', '极寒冷焰', 'Ice', 'Special', 140, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('555', 'Snarl', '大声咆哮', 'Dark', 'Special', 55, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('556', 'Icicle Crash', '冰柱坠击', 'Ice', 'Physical', 85, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('557', 'V-create', 'Ｖ热焰', 'Fire', 'Physical', 180, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('558', 'Fusion Flare', '交错火焰', 'Fire', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('559', 'Fusion Bolt', '交错闪电', 'Electric', 'Physical', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('560', 'Flying Press', '飞身重压', 'Fighting', 'Physical', 100, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('561', 'Mat Block', '掀榻榻米', 'Fighting', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('562', 'Belch', '打嗝', 'Poison', 'Special', 120, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('563', 'Rototiller', '耕地', 'Ground', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('564', 'Sticky Web', '黏黏网', 'Bug', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('565', 'Fell Stinger', '致命针刺', 'Bug', 'Physical', 50, 100, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('566', 'Phantom Force', '潜灵奇袭', 'Ghost', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('567', 'Trick-or-Treat', '万圣夜', 'Ghost', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('568', 'Noble Roar', '战吼', 'Normal', 'Status', NULL, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('569', 'Ion Deluge', '等离子浴', 'Electric', 'Status', NULL, NULL, 25);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('570', 'Parabolic Charge', '抛物面充电', 'Electric', 'Special', 65, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('571', 'Forest''s Curse', '森林咒术', 'Grass', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('572', 'Petal Blizzard', '落英缤纷', 'Grass', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('573', 'Freeze-Dry', '冷冻干燥', 'Ice', 'Special', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('574', 'Disarming Voice', '魅惑之声', 'Fairy', 'Special', 40, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('575', 'Parting Shot', '抛下狠话', 'Dark', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('576', 'Topsy-Turvy', '颠倒', 'Dark', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('577', 'Draining Kiss', '吸取之吻', 'Fairy', 'Special', 50, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('578', 'Crafty Shield', '戏法防守', 'Fairy', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('579', 'Flower Shield', '鲜花防守', 'Fairy', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('580', 'Grassy Terrain', '青草场地', 'Grass', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('581', 'Misty Terrain', '薄雾场地', 'Fairy', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('582', 'Electrify', '输电', 'Electric', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('583', 'Play Rough', '嬉闹', 'Fairy', 'Physical', 90, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('584', 'Fairy Wind', '妖精之风', 'Fairy', 'Special', 40, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('585', 'Moonblast', '月亮之力', 'Fairy', 'Special', 95, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('586', 'Boomburst', '爆音波', 'Normal', 'Special', 140, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('587', 'Fairy Lock', '妖精之锁', 'Fairy', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('588', 'King''s Shield', '王者盾牌', 'Steel', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('589', 'Play Nice', '和睦相处', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('590', 'Confide', '密语', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('591', 'Diamond Storm', '钻石风暴', 'Rock', 'Physical', 100, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('592', 'Steam Eruption', '蒸汽爆炸', 'Water', 'Special', 110, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('593', 'Hyperspace Hole', '异次元洞', 'Psychic', 'Special', 80, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('594', 'Water Shuriken', '飞水手里剑', 'Water', 'Special', 15, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('595', 'Mystical Fire', '魔法火焰', 'Fire', 'Special', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('596', 'Spiky Shield', '尖刺防守', 'Grass', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('597', 'Aromatic Mist', '芳香薄雾', 'Fairy', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('598', 'Eerie Impulse', '怪异电波', 'Electric', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('599', 'Venom Drench', '毒液陷阱', 'Poison', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('600', 'Powder', '粉尘', 'Bug', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('601', 'Geomancy', '大地掌控', 'Fairy', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('602', 'Magnetic Flux', '磁场操控', 'Electric', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('603', 'Happy Hour', '欢乐时光', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('604', 'Electric Terrain', '电气场地', 'Electric', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('605', 'Dazzling Gleam', '魔法闪耀', 'Fairy', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('606', 'Celebrate', '庆祝', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('607', 'Hold Hands', '牵手', 'Normal', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('608', 'Baby-Doll Eyes', '圆瞳', 'Fairy', 'Status', NULL, 100, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('609', 'Nuzzle', '蹭蹭脸颊', 'Electric', 'Physical', 20, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('610', 'Hold Back', '手下留情', 'Normal', 'Physical', 40, 100, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('611', 'Infestation', '纠缠不休', 'Bug', 'Special', 20, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('612', 'Power-Up Punch', '增强拳', 'Fighting', 'Physical', 40, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('613', 'Oblivion Wing', '归天之翼', 'Flying', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('614', 'Thousand Arrows', '千箭齐发', 'Ground', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('615', 'Thousand Waves', '千波激荡', 'Ground', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('616', 'Land''s Wrath', '大地神力', 'Ground', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('617', 'Light of Ruin', '破灭之光', 'Fairy', 'Special', 140, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('618', 'Origin Pulse', '根源波动', 'Water', 'Special', 110, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('619', 'Precipice Blades', '断崖之剑', 'Ground', 'Physical', 120, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('620', 'Dragon Ascent', '画龙点睛', 'Flying', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('621', 'Hyperspace Fury', '异次元猛攻', 'Dark', 'Physical', 100, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('622', 'Breakneck Blitz', '究极无敌大冲撞', 'Normal', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('623', 'Breakneck Blitz', '究极无敌大冲撞', 'Normal', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('624', 'All-Out Pummeling', '全力无双激烈拳', 'Fighting', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('625', 'All-Out Pummeling', '全力无双激烈拳', 'Fighting', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('626', 'Supersonic Skystrike', '极速俯冲轰烈撞', 'Flying', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('627', 'Supersonic Skystrike', '极速俯冲轰烈撞', 'Flying', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('628', 'Acid Downpour', '强酸剧毒灭绝雨', 'Poison', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('629', 'Acid Downpour', '强酸剧毒灭绝雨', 'Poison', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('630', 'Tectonic Rage', '地隆啸天大终结', 'Ground', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('631', 'Tectonic Rage', '地隆啸天大终结', 'Ground', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('632', 'Continental Crush', '毁天灭地巨岩坠', 'Rock', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('633', 'Continental Crush', '毁天灭地巨岩坠', 'Rock', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('634', 'Savage Spin-Out', '绝对捕食回旋斩', 'Bug', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('635', 'Savage Spin-Out', '绝对捕食回旋斩', 'Bug', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('636', 'Never-Ending Nightmare', '无尽暗夜之诱惑', 'Ghost', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('637', 'Never-Ending Nightmare', '无尽暗夜之诱惑', 'Ghost', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('638', 'Corkscrew Crash', '超绝螺旋连击', 'Steel', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('639', 'Corkscrew Crash', '超绝螺旋连击', 'Steel', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('640', 'Inferno Overdrive', '超强极限爆焰弹', 'Fire', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('641', 'Inferno Overdrive', '超强极限爆焰弹', 'Fire', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('642', 'Hydro Vortex', '超级水流大漩涡', 'Water', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('643', 'Hydro Vortex', '超级水流大漩涡', 'Water', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('644', 'Bloom Doom', '绚烂缤纷花怒放', 'Grass', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('645', 'Bloom Doom', '绚烂缤纷花怒放', 'Grass', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('646', 'Gigavolt Havoc', '终极伏特狂雷闪', 'Electric', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('647', 'Gigavolt Havoc', '终极伏特狂雷闪', 'Electric', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('648', 'Shattered Psyche', '至高精神破坏波', 'Psychic', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('649', 'Shattered Psyche', '至高精神破坏波', 'Psychic', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('650', 'Subzero Slammer', '激狂大地万里冰', 'Ice', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('651', 'Subzero Slammer', '激狂大地万里冰', 'Ice', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('652', 'Devastating Drake', '究极巨龙震天地', 'Dragon', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('653', 'Devastating Drake', '究极巨龙震天地', 'Dragon', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('654', 'Black Hole Eclipse', '黑洞吞噬万物灭', 'Dark', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('655', 'Black Hole Eclipse', '黑洞吞噬万物灭', 'Dark', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('656', 'Twinkle Tackle', '可爱星星飞天撞', 'Fairy', 'Physical', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('657', 'Twinkle Tackle', '可爱星星飞天撞', 'Fairy', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('658', 'Catastropika', '皮卡皮卡必杀击', 'Electric', 'Physical', 210, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('659', 'Shore Up', '集沙', 'Ground', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('660', 'First Impression', '迎头一击', 'Bug', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('661', 'Baneful Bunker', '碉堡', 'Poison', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('662', 'Spirit Shackle', '缝影', 'Ghost', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('663', 'Darkest Lariat', 'ＤＤ金勾臂', 'Dark', 'Physical', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('664', 'Sparkling Aria', '泡影的咏叹调', 'Water', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('665', 'Ice Hammer', '冰锤', 'Ice', 'Physical', 100, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('666', 'Floral Healing', '花疗', 'Fairy', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('667', 'High Horsepower', '十万马力', 'Ground', 'Physical', 95, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('668', 'Strength Sap', '吸取力量', 'Grass', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('669', 'Solar Blade', '日光刃', 'Grass', 'Physical', 125, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('670', 'Leafage', '树叶', 'Grass', 'Physical', 40, 100, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('671', 'Spotlight', '聚光灯', 'Normal', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('672', 'Toxic Thread', '毒丝', 'Poison', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('673', 'Laser Focus', '磨砺', 'Normal', 'Status', NULL, NULL, 30);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('674', 'Gear Up', '辅助齿轮', 'Steel', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('675', 'Throat Chop', '深渊突刺', 'Dark', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('676', 'Pollen Puff', '花粉团', 'Bug', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('677', 'Anchor Shot', '掷锚', 'Steel', 'Physical', 80, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('678', 'Psychic Terrain', '精神场地', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('679', 'Lunge', '猛扑', 'Bug', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('680', 'Fire Lash', '火焰鞭', 'Fire', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('681', 'Power Trip', '嚣张', 'Dark', 'Physical', 20, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('682', 'Burn Up', '燃尽', 'Fire', 'Special', 130, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('683', 'Speed Swap', '速度互换', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('684', 'Smart Strike', '修长之角', 'Steel', 'Physical', 70, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('685', 'Purify', '净化', 'Poison', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('686', 'Revelation Dance', '觉醒之舞', 'Normal', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('687', 'Core Enforcer', '核心惩罚者', 'Dragon', 'Special', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('688', 'Trop Kick', '热带踢', 'Grass', 'Physical', 70, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('689', 'Instruct', '号令', 'Psychic', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('690', 'Beak Blast', '鸟嘴加农炮', 'Flying', 'Physical', 100, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('691', 'Clanging Scales', '鳞片噪音', 'Dragon', 'Special', 110, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('692', 'Dragon Hammer', '龙锤', 'Dragon', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('693', 'Brutal Swing', '狂舞挥打', 'Dark', 'Physical', 60, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('694', 'Aurora Veil', '极光幕', 'Ice', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('695', 'Sinister Arrow Raid', '遮天蔽日暗影箭', 'Ghost', 'Physical', 180, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('696', 'Malicious Moonsault', '极恶飞跃粉碎击', 'Dark', 'Physical', 180, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('697', 'Oceanic Operetta', '海神庄严交响乐', 'Water', 'Special', 195, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('698', 'Guardian of Alola', '巨人卫士・阿罗拉', 'Fairy', 'Special', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('699', 'Soul-Stealing 7-Star Strike', '七星夺魂腿', 'Ghost', 'Physical', 195, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('700', 'Stoked Sparksurfer', '驾雷驭电戏冲浪', 'Electric', 'Special', 175, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('701', 'Pulverizing Pancake', '认真起来大爆击', 'Normal', 'Physical', 210, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('702', 'Extreme Evoboost', '九彩昇华齐聚顶', 'Normal', 'Status', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('703', 'Genesis Supernova', '起源超新星大爆炸', 'Psychic', 'Special', 185, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('704', 'Shell Trap', '陷阱甲壳', 'Fire', 'Special', 150, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('705', 'Fleur Cannon', '花朵加农炮', 'Fairy', 'Special', 130, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('706', 'Psychic Fangs', '精神之牙', 'Psychic', 'Physical', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('707', 'Stomping Tantrum', '跺脚', 'Ground', 'Physical', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('708', 'Shadow Bone', '暗影之骨', 'Ghost', 'Physical', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('709', 'Accelerock', '冲岩', 'Rock', 'Physical', 40, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('710', 'Liquidation', '水流裂破', 'Water', 'Physical', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('711', 'Prismatic Laser', '棱镜镭射', 'Psychic', 'Special', 160, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('712', 'Spectral Thief', '暗影偷盗', 'Ghost', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('713', 'Sunsteel Strike', '流星闪冲', 'Steel', 'Physical', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('714', 'Moongeist Beam', '暗影之光', 'Ghost', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('715', 'Tearful Look', '泪眼汪汪', 'Normal', 'Status', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('716', 'Zing Zap', '麻麻刺刺', 'Electric', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('717', 'Nature''s Madness', '自然之怒', 'Fairy', 'Special', NULL, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('718', 'Multi-Attack', '多属性攻击', 'Normal', 'Physical', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('719', '10,000,000 Volt Thunderbolt', '千万伏特', 'Electric', 'Special', 195, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('720', 'Mind Blown', '惊爆大头', 'Fire', 'Special', 150, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('721', 'Plasma Fists', '等离子闪电拳', 'Electric', 'Physical', 100, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('722', 'Photon Geyser', '光子喷涌', 'Psychic', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('723', 'Clangorous Soulblaze', '炽魂热舞烈音爆', 'Dragon', 'Special', 185, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('724', 'Splintered Stormshards', '狼啸石牙飓风暴', 'Rock', 'Physical', 190, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('725', 'Let''s Snuggle Forever', '亲密无间大乱揍', 'Fairy', 'Physical', 190, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('726', 'Searing Sunraze Smash', '日光回旋下苍穹', 'Steel', 'Physical', 200, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('727', 'Menacing Moonraze Maelstrom', '月华飞溅落灵霄', 'Ghost', 'Special', 200, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('728', 'Light That Burns the Sky', '焚天灭世炽光爆', 'Psychic', 'Special', 200, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('729', 'Zippy Zap', '电电加速', 'Electric', 'Physical', 50, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('730', 'Splishy Splash', '滔滔冲浪', 'Water', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('731', 'Floaty Fall', '飘飘坠落', 'Flying', 'Physical', 90, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('732', 'Pika Papow', '闪闪雷光', 'Electric', 'Special', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('733', 'Bouncy Bubble', '活活气泡', 'Water', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('734', 'Buzzy Buzz', '麻麻电击', 'Electric', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('735', 'Sizzly Slide', '熊熊火爆', 'Fire', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('736', 'Glitzy Glow', '哗哗气场', 'Psychic', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('737', 'Baddy Bad', '坏坏领域', 'Dark', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('738', 'Sappy Seed', '茁茁炸弹', 'Grass', 'Physical', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('739', 'Freezy Frost', '冰冰霜冻', 'Ice', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('740', 'Sparkly Swirl', '亮亮风暴', 'Fairy', 'Special', 90, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('741', 'Veevee Volley', '砰砰击破', 'Normal', 'Physical', NULL, NULL, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('742', 'Double Iron Bash', '钢拳双击', 'Steel', 'Physical', 60, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('743', 'Max Guard', '极巨防壁', 'Normal', 'Status', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('744', 'Dynamax Cannon', '极巨炮', 'Dragon', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('745', 'Snipe Shot', '狙击', 'Water', 'Special', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('746', 'Jaw Lock', '紧咬不放', 'Dark', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('747', 'Stuff Cheeks', '大快朵颐', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('748', 'No Retreat', '背水一战', 'Fighting', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('749', 'Tar Shot', '沥青射击', 'Rock', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('750', 'Magic Powder', '魔法粉', 'Psychic', 'Status', NULL, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('751', 'Dragon Darts', '龙箭', 'Dragon', 'Physical', 50, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('752', 'Teatime', '茶会', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('753', 'Octolock', '蛸固', 'Fighting', 'Status', NULL, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('754', 'Bolt Beak', '电喙', 'Electric', 'Physical', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('755', 'Fishious Rend', '鳃咬', 'Water', 'Physical', 85, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('756', 'Court Change', '换场', 'Normal', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('757', 'Max Flare', '极巨火爆', 'Fire', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('758', 'Max Flutterby', '极巨虫蛊', 'Bug', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('759', 'Max Lightning', '极巨闪电', 'Electric', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('760', 'Max Strike', '极巨攻击', 'Normal', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('761', 'Max Knuckle', '极巨拳斗', 'Fighting', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('762', 'Max Phantasm', '极巨幽魂', 'Ghost', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('763', 'Max Hailstorm', '极巨寒冰', 'Ice', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('764', 'Max Ooze', '极巨酸毒', 'Poison', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('765', 'Max Geyser', '极巨水流', 'Water', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('766', 'Max Airstream', '极巨飞冲', 'Flying', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('767', 'Max Starfall', '极巨妖精', 'Fairy', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('768', 'Max Wyrmwind', '极巨龙骑', 'Dragon', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('769', 'Max Mindstorm', '极巨超能', 'Psychic', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('770', 'Max Rockfall', '极巨岩石', 'Rock', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('771', 'Max Quake', '极巨大地', 'Ground', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('772', 'Max Darkness', '极巨恶霸', 'Dark', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('773', 'Max Overgrowth', '极巨草原', 'Grass', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('774', 'Max Steelspike', '极巨钢铁', 'Steel', '极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('775', 'Clangorous Soul', '魂舞烈音爆', 'Dragon', 'Status', NULL, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('776', 'Body Press', '扑击', 'Fighting', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('777', 'Decorate', '装饰', 'Fairy', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('778', 'Drum Beating', '鼓击', 'Grass', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('779', 'Snap Trap', '捕兽夹', 'Grass', 'Physical', 35, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('780', 'Pyro Ball', '火焰球', 'Fire', 'Physical', 120, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('781', 'Behemoth Blade', '巨兽斩', 'Steel', 'Physical', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('782', 'Behemoth Bash', '巨兽弹', 'Steel', 'Physical', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('783', 'Aura Wheel', '气场轮', 'Electric', 'Physical', 110, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('784', 'Breaking Swipe', '广域破坏', 'Dragon', 'Physical', 60, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('785', 'Branch Poke', '木枝突刺', 'Grass', 'Physical', 40, 100, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('786', 'Overdrive', '破音', 'Electric', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('787', 'Apple Acid', '苹果酸', 'Grass', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('788', 'Grav Apple', '万有引力', 'Grass', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('789', 'Spirit Break', '灵魂冲击', 'Fairy', 'Physical', 75, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('790', 'Strange Steam', '神奇蒸汽', 'Fairy', 'Special', 90, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('791', 'Life Dew', '生命水滴', 'Water', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('792', 'Obstruct', '拦堵', 'Dark', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('793', 'False Surrender', '假跪真撞', 'Dark', 'Physical', 80, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('794', 'Meteor Assault', '流星突击', 'Fighting', 'Physical', 150, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('795', 'Eternabeam', '无极光束', 'Dragon', 'Special', 160, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('796', 'Steel Beam', '铁蹄光线', 'Steel', 'Special', 140, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Wildfire', '超极巨深渊灭焰', 'Fire', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Befuddle', '超极巨蝶影蛊惑', 'Bug', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Volt Crash', '超极巨万雷轰顶', 'Electric', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Gold Rush', '超极巨特大金币', 'Normal', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Chi Strike', '超极巨会心一击', 'Fighting', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Terror', '超极巨幻影幽魂', 'Ghost', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Resonance', '超极巨极光旋律', 'Ice', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Cuddle', '超极巨热情拥抱', 'Normal', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Replenish', '超极巨资源再生', 'Normal', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Malodor', '超极巨臭气冲天', 'Poison', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Stonesurge', '超极巨岩阵以待', 'Water', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Wind Rage', '超极巨旋风袭卷', 'Flying', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Stun Shock', '超极巨异毒电场', 'Electric', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Finale', '超极巨幸福圆满', 'Fairy', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Depletion', '超极巨劣化衰变', 'Dragon', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Gravitas', '超极巨天道七星', 'Psychic', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Volcalith', '超极巨炎石喷发', 'Rock', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Sandblast', '超极巨沙尘漫天', 'Ground', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Snooze', '超极巨睡魔降临', 'Dark', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Tartness', '超极巨酸不溜丢', 'Grass', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Sweetness', '超极巨琼浆玉液', 'Grass', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Smite', '超极巨天谴雷诛', 'Fairy', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Steelsurge', '超极巨钢铁阵法', 'Steel', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Meltdown', '超极巨液金熔击', 'Steel', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Foam Burst', '超极巨激漩泡涡', 'Water', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Centiferno', '超极巨百火焚野', 'Fire', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('797', 'Expanding Force', '广域战力', 'Psychic', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('798', 'Steel Roller', '铁滚轮', 'Steel', 'Physical', 130, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('799', 'Scale Shot', '鳞射', 'Dragon', 'Physical', 25, 90, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('800', 'Meteor Beam', '流星光束', 'Rock', 'Special', 120, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('801', 'Shell Side Arm', '臂贝武器', 'Poison', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('802', 'Misty Explosion', '薄雾炸裂', 'Fairy', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('803', 'Grassy Glide', '青草滑梯', 'Grass', 'Physical', 55, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('804', 'Rising Voltage', '电力上升', 'Electric', 'Special', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('805', 'Terrain Pulse', '大地波动', 'Normal', 'Special', 50, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('806', 'Skitter Smack', '爬击', 'Bug', 'Physical', 70, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('807', 'Burning Jealousy', '妒火', 'Fire', 'Special', 70, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('808', 'Lash Out', '泄愤', 'Dark', 'Physical', 75, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('809', 'Poltergeist', '灵骚', 'Ghost', 'Physical', 110, 90, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('810', 'Corrosive Gas', '腐蚀气体', 'Poison', 'Status', NULL, NULL, 40);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('811', 'Coaching', '指导', 'Fighting', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('812', 'Flip Turn', '快速折返', 'Water', 'Physical', 60, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('813', 'Triple Axel', '三旋击', 'Ice', 'Physical', 20, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('814', 'Dual Wingbeat', '双翼', 'Flying', 'Physical', 40, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('815', 'Scorching Sands', '热沙大地', 'Ground', 'Special', 70, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('816', 'Jungle Healing', '丛林治疗', 'Grass', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('817', 'Wicked Blow', '暗冥强击', 'Dark', 'Physical', 75, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('818', 'Surging Strikes', '水流连打', 'Water', 'Physical', 25, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Drum Solo', '超极巨狂擂乱打', 'Grass', '超极巨', 160, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Fireball', '超极巨破阵火球', 'Fire', '超极巨', 160, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Hydrosnipe', '超极巨狙击神射', 'Water', '超极巨', 160, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Vine Lash', '超极巨灰飞鞭灭', 'Grass', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Cannonade', '超极巨水炮轰灭', 'Water', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max One Blow', '超极巨夺命一击', 'Dark', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('—', 'G-Max Rapid Flow', '超极巨流水连击', 'Water', '超极巨', NULL, NULL, NULL);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('819', 'Thunder Cage', '雷电囚笼', 'Electric', 'Special', 80, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('820', 'Dragon Energy', '巨龙威能', 'Dragon', 'Special', 150, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('821', 'Freezing Glare', '冰冷视线', 'Psychic', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('822', 'Fiery Wrath', '怒火中烧', 'Dark', 'Special', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('823', 'Thunderous Kick', '雷鸣蹴击', 'Fighting', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('824', 'Glacial Lance', '雪矛', 'Ice', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('825', 'Astral Barrage', '星碎', 'Ghost', 'Special', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('826', 'Eerie Spell', '诡异咒语', 'Psychic', 'Special', 80, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('827', 'Dire Claw', '克命爪', 'Poison', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('828', 'Psyshield Bash', '屏障猛攻', 'Psychic', 'Physical', 70, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('829', 'Power Shift', '力量转换', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('830', 'Stone Axe', '岩斧', 'Rock', 'Physical', 65, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('831', 'Springtide Storm', '阳春风暴', 'Fairy', 'Special', 100, 80, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('832', 'Mystical Power', '神秘之力', 'Psychic', 'Special', 70, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('833', 'Raging Fury', '大愤慨', 'Fire', 'Physical', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('834', 'Wave Crash', '波动冲', 'Water', 'Physical', 120, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('835', 'Chloroblast', '叶绿爆震', 'Grass', 'Special', 150, 95, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('836', 'Mountain Gale', '冰山风', 'Ice', 'Physical', 100, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('837', 'Victory Dance', '胜利之舞', 'Fighting', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('838', 'Headlong Rush', '突飞猛扑', 'Ground', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('839', 'Barb Barrage', '毒千针', 'Poison', 'Physical', 60, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('840', 'Esper Wing', '气场之翼', 'Psychic', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('841', 'Bitter Malice', '冤冤相报', 'Ghost', 'Special', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('842', 'Shelter', '闭关', 'Steel', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('843', 'Triple Arrows', '三连箭', 'Fighting', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('844', 'Infernal Parade', '群魔乱舞', 'Ghost', 'Special', 60, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('845', 'Ceaseless Edge', '秘剑・千重涛', 'Dark', 'Physical', 65, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('846', 'Bleakwind Storm', '枯叶风暴', 'Flying', 'Special', 100, 80, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('847', 'Wildbolt Storm', '鸣雷风暴', 'Electric', 'Special', 100, 80, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('848', 'Sandsear Storm', '热沙风暴', 'Ground', 'Special', 100, 80, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('849', 'Lunar Blessing', '新月祈祷', 'Psychic', 'Status', NULL, NULL, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('850', 'Take Heart', '勇气填充', 'Psychic', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('851', 'Tera Blast', '太晶爆发', 'Normal', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('852', 'Silk Trap', '线阱', 'Bug', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('853', 'Axe Kick', '下压踢', 'Fighting', 'Physical', 120, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('854', 'Last Respects', '扫墓', 'Ghost', 'Physical', 50, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('855', 'Lumina Crash', '琉光冲激', 'Psychic', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('856', 'Order Up', '上菜', 'Dragon', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('857', 'Jet Punch', '喷射拳', 'Water', 'Physical', 60, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('858', 'Spicy Extract', '辣椒精华', 'Grass', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('859', 'Spin Out', '疾速转轮', 'Steel', 'Physical', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('860', 'Population Bomb', '鼠数儿', 'Normal', 'Physical', 20, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('861', 'Ice Spinner', '冰旋', 'Ice', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('862', 'Glaive Rush', '巨剑突击', 'Dragon', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('863', 'Revival Blessing', '复生祈祷', 'Normal', 'Status', NULL, NULL, 1);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('864', 'Salt Cure', '盐腌', 'Rock', 'Physical', 40, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('865', 'Triple Dive', '三连钻', 'Water', 'Physical', 30, 95, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('866', 'Mortal Spin', '晶光转转', 'Poison', 'Physical', 30, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('867', 'Doodle', '描绘', 'Normal', 'Status', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('868', 'Fillet Away', '甩肉', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('869', 'Kowtow Cleave', '仆刀', 'Dark', 'Physical', 85, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('870', 'Flower Trick', '千变万花', 'Grass', 'Physical', 70, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('871', 'Torch Song', '闪焰高歌', 'Fire', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('872', 'Aqua Step', '流水旋舞', 'Water', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('873', 'Raging Bull', '怒牛', 'Normal', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('874', 'Make It Rain', '淘金潮', 'Steel', 'Special', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('875', 'Psyblade', '精神剑', 'Psychic', 'Physical', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('876', 'Hydro Steam', '水蒸气', 'Water', 'Special', 80, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('877', 'Ruination', '大灾难', 'Dark', 'Special', NULL, 90, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('878', 'Collision Course', '全开猛撞', 'Fighting', 'Physical', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('879', 'Electro Drift', '闪电猛冲', 'Electric', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('880', 'Shed Tail', '断尾', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('881', 'Chilly Reception', '冷笑话', 'Ice', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('882', 'Tidy Up', '大扫除', 'Normal', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('883', 'Snowscape', '雪景', 'Ice', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('884', 'Pounce', '虫扑', 'Bug', 'Physical', 50, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('885', 'Trailblaze', '起草', 'Grass', 'Physical', 50, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('886', 'Chilling Water', '泼冷水', 'Water', 'Special', 50, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('887', 'Hyper Drill', '强力钻', 'Normal', 'Physical', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('888', 'Twin Beam', '双光束', 'Psychic', 'Special', 40, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('889', 'Rage Fist', '愤怒之拳', 'Ghost', 'Physical', 50, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('890', 'Armor Cannon', '铠农炮', 'Fire', 'Special', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('891', 'Bitter Blade', '悔念剑', 'Fire', 'Physical', 90, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('892', 'Double Shock', '电光双击', 'Electric', 'Physical', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('893', 'Gigaton Hammer', '巨力锤', 'Steel', 'Physical', 160, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('894', 'Comeuppance', '复仇', 'Dark', 'Physical', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('895', 'Aqua Cutter', '水波刀', 'Water', 'Physical', 70, 100, 20);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('896', 'Blazing Torque', '灼热暴冲', 'Fire', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('897', 'Wicked Torque', '黑暗暴冲', 'Dark', 'Physical', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('898', 'Noxious Torque', '剧毒暴冲', 'Poison', 'Physical', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('899', 'Combat Torque', '格斗暴冲', 'Fighting', 'Physical', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('900', 'Magical Torque', '魔法暴冲', 'Fairy', 'Physical', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('901', 'Blood Moon', '血月', 'Normal', 'Special', 140, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('902', 'Matcha Gotcha', '刷刷茶炮', 'Grass', 'Special', 80, 90, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('903', 'Syrup Bomb‎', '糖浆炸弹', 'Grass', 'Special', 60, 85, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('904', 'Ivy Cudgel', '棘藤棒', 'Grass', 'Physical', 100, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('905', 'Electro Shot', '电光束', 'Electric', 'Special', 130, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('906', 'Tera Starstorm', '晶光星群', 'Normal', 'Special', 120, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('907', 'Fickle Beam', '随机光', 'Dragon', 'Special', 80, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('908', 'Burning Bulwark', '火焰守护', 'Fire', 'Status', NULL, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('909', 'Thunderclap', '迅雷', 'Electric', 'Special', 70, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('910', 'Mighty Cleave', '强刃攻击', 'Rock', 'Physical', 95, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('911', 'Tachyon Cutter', '迅子利刃', 'Steel', 'Special', 50, NULL, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('912', 'Hard Press', '硬压', 'Steel', 'Physical', NULL, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('913', 'Dragon Cheer', '龙声鼓舞', 'Dragon', 'Status', NULL, NULL, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('914', 'Alluring Voice', '魅诱之声', 'Fairy', 'Special', 80, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('915', 'Temper Flare', '豁出去', 'Fire', 'Physical', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('916', 'Supercell Slam', '闪电强袭', 'Electric', 'Physical', 100, 95, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('917', 'Psychic Noise', '精神噪音', 'Psychic', 'Special', 75, 100, 10);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('918', 'Upper Hand', '快手还击', 'Fighting', 'Physical', 65, 100, 15);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('919', 'Malignant Chain', '邪毒锁链', 'Poison', 'Special', 100, 100, 5);
INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ('920', 'Nihil Light', '归无之光', 'Dragon', 'Special', 200, NULL, NULL);
-- Pokemon Abilities Data
CREATE TABLE IF NOT EXISTS pokemon_abilities (
    ability_id VARCHAR(10) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('001', 'Stench', '恶臭', '通过释放臭臭的气味，在攻击的时候，有时会使对手畏缩。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('002', 'Drizzle', '降雨', '出场时，会将天气变为下雨。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('003', 'Speed Boost', '加速', '每一回合速度会变快。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('004', 'Battle Armor', '战斗盔甲', '被坚硬的甲壳守护着，不会被对手的攻击击中要害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('005', 'Sturdy', '结实', '在ＨＰ全满时，即使受到招式攻击，也不会被一击打倒。一击必杀的招式也没有效果。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('006', 'Damp', '湿气', '通过把周围都弄湿，使谁都无法使用自爆等爆炸类的招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('007', 'Limber', '柔软', '因为身体柔软，不会变为麻痹状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('008', 'Sand Veil', '沙隐', '在沙暴的时候，闪避率会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('009', 'Static', '静电', '身上带有静电，有时会让接触到的对手麻痹。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('010', 'Volt Absorb', '蓄电', '受到电属性的招式攻击时，不会受到伤害，而是会回复。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('011', 'Water Absorb', '储水', '受到水属性的招式攻击时，不会受到伤害，而是会回复。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('012', 'Oblivious', '迟钝', '因为感觉迟钝，不会变为着迷和被挑衅状态。对威吓也毫不动摇。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('013', 'Cloud Nine', '无关天气', '任何天气的影响都会消失。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('014', 'Compound Eyes', '复眼', '因为拥有复眼，招式的命中率会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('015', 'Insomnia', '不眠', '因为有着睡不着的体质，所以不会陷入睡眠状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('016', 'Color Change', '变色', '自己的属性会变为从对手处所受招式的属性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('017', 'Immunity', '免疫', '因为体内拥有免疫能力，不会变为中毒状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('018', 'Flash Fire', '引火', '受到火属性的招式攻击时，吸收火焰，自己使出的火属性招式会变强。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('019', 'Shield Dust', '鳞粉', '被鳞粉守护着，不会受到招式的追加效果影响。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('020', 'Own Tempo', '我行我素', '因为我行我素，不会变为混乱状态。对威吓也毫不动摇。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('021', 'Suction Cups', '吸盘', '用吸盘牢牢贴在地面上，让替换宝可梦的招式和道具无效。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('022', 'Intimidate', '威吓', '出场时威吓对手，让其退缩，降低对手的攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('023', 'Shadow Tag', '踩影', '踩住对手的影子使其无法逃走或替换。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('024', 'Rough Skin', '粗糙皮肤', '受到攻击时，用粗糙的皮肤弄伤接触到自己的对手。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('025', 'Wonder Guard', '神奇守护', '不可思议的力量，只有效果绝佳的招式才能击中。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('026', 'Levitate', '飘浮', '从地面浮起，从而不会受到地面属性招式的攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('027', 'Effect Spore', '孢子', '受到攻击时，有时会把接触到自己的对手变为中毒、麻痹或睡眠状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('028', 'Synchronize', '同步', '将自己的中毒、麻痹或灼伤状态传染给对手。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('029', 'Clear Body', '恒净之躯', '不会因为对手的招式或特性而被降低能力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('030', 'Natural Cure', '自然回复', '回到同行队伍后，异常状态就会被治愈。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('031', 'Lightning Rod', '避雷针', '将电属性的招式吸引到自己身上，不会受到伤害，而是会提高特攻。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('032', 'Serene Grace', '天恩', '托天恩的福，招式的追加效果容易出现。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('033', 'Swift Swim', '悠游自如', '下雨天气时，速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('034', 'Chlorophyll', '叶绿素', '晴朗天气时，速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('035', 'Illuminate', '发光', '通过让周围变亮来保持命中率不会被降低。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('036', 'Trace', '复制', '出场时，复制对手的特性，变为与之相同的特性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('037', 'Huge Power', '大力士', '物理攻击的威力会变为２倍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('038', 'Poison Point', '毒刺', '有时会让接触到自己的对手变为中毒状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('039', 'Inner Focus', '精神力', '拥有经过锻炼的精神，而不会因对手的攻击而畏缩。对威吓也毫不动摇。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('040', 'Magma Armor', '熔岩铠甲', '将炽热的熔岩覆盖在身上，不会变为冰冻状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('041', 'Water Veil', '水幕', '将水幕裹在身上，不会变为灼伤状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('042', 'Magnet Pull', '磁力', '用磁力吸住钢属性的宝可梦，使其无法逃走。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('043', 'Soundproof', '隔音', '通过屏蔽声音，不受到声音招式的影响。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('044', 'Rain Dish', '雨盘', '下雨天气时，会缓缓回复ＨＰ。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('045', 'Sand Stream', '扬沙', '出场时，会把天气变为沙暴。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('046', 'Pressure', '压迫感', '给予对手压迫感，大量减少其使用招式的ＰＰ。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('047', 'Thick Fat', '厚脂肪', '因为被厚厚的脂肪保护着，会让火属性和冰属性的招式伤害减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('048', 'Early Bird', '早起', '即使变为睡眠状态，也能以２倍的速度提早醒来。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('049', 'Flame Body', '火焰之躯', '有时会让接触到自己的对手变为灼伤状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('050', 'Run Away', '逃跑', '一定能从野生宝可梦那儿逃走。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('051', 'Keen Eye', '锐利目光', '多亏了锐利的目光，命中率不会被降低。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('052', 'Hyper Cutter', '怪力钳', '因为拥有以力量自豪的钳子，不会被对手降低攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('053', 'Pickup', '捡拾', '有时会捡来对手用过的道具，冒险过程中也会捡到。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('054', 'Truant', '懒惰', '如果使出招式，下一回合就会休息。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('055', 'Hustle', '活力', '自己的攻击变高，但命中率会降低。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('056', 'Cute Charm', '迷人之躯', '有时会让接触到自己的对手着迷。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('057', 'Plus', '正电', '出场的伙伴之间如果有正电或负电特性的宝可梦，自己的特攻会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('058', 'Minus', '负电', '出场的伙伴之间如果有正电或负电特性的宝可梦，自己的特攻会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('059', 'Forecast', '阴晴不定', '受天气的影响，会变为水属性、火属性或冰属性中的某一个。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('060', 'Sticky Hold', '黏着', '因为道具是粘在黏性身体上的，所以不会被对手夺走。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('061', 'Shed Skin', '蜕皮', '通过蜕去身上的皮，有时会治愈异常状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('062', 'Guts', '毅力', '如果变为异常状态，会拿出毅力，攻击会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('063', 'Marvel Scale', '神奇鳞片', '如果变为异常状态，神奇鳞片会发生反应，防御会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('064', 'Liquid Ooze', '污泥浆', '吸收了污泥浆的对手会因强烈的恶臭而受到伤害，减少ＨＰ。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('065', 'Overgrow', '茂盛', 'ＨＰ减少的时候，草属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('066', 'Blaze', '猛火', 'ＨＰ减少的时候，火属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('067', 'Torrent', '激流', 'ＨＰ减少的时候，水属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('068', 'Swarm', '虫之预感', 'ＨＰ减少的时候，虫属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('069', 'Rock Head', '坚硬脑袋', '即使使出会受反作用力伤害的招式，ＨＰ也不会减少。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('070', 'Drought', '日照', '出场时，会将天气变为晴朗。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('071', 'Arena Trap', '沙穴', '在战斗中让对手无法逃走。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('072', 'Vital Spirit', '干劲', '通过激发出干劲，不会变为睡眠状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('073', 'White Smoke', '白色烟雾', '被白色烟雾保护着，不会被对手降低能力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('074', 'Pure Power', '瑜伽之力', '因瑜伽的力量，物理攻击的威力会变为２倍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('075', 'Shell Armor', '硬壳盔甲', '被坚硬的壳保护着，对手的攻击不会击中要害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('076', 'Cacophony', '杂音', '不受到声音招式的影响。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('076', 'Air Lock', '气闸', '所有天气的影响都会消失。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('077', 'Tangled Feet', '蹒跚', '在混乱状态时，闪避率会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('078', 'Motor Drive', '电气引擎', '受到电属性的招式攻击时，不会受到伤害，而是速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('079', 'Rivalry', '斗争心', '面对性别相同的对手，会燃起斗争心，变得更强。而面对性别不同的，则会变弱。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('080', 'Steadfast', '不屈之心', '每次畏缩时，不屈之心就会燃起，速度也会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('081', 'Snow Cloak', '雪隐', '下雪天气时，闪避率会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('082', 'Gluttony', '贪吃鬼', '原本ＨＰ变得很少时才会吃树果，在ＨＰ还有一半时就会把它吃掉。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('083', 'Anger Point', '愤怒穴位', '要害被击中时，会大发雷霆，攻击力变为最大。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('084', 'Unburden', '轻装', '失去所持有的道具时，速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('085', 'Heatproof', '耐热', '耐热的体质会让火属性的招式伤害减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('086', 'Simple', '单纯', '能力变化会变为平时的２倍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('087', 'Dry Skin', '干燥皮肤', '下雨天气时和受到水属性的招式时，ＨＰ会回复。晴朗天气时和受到火属性的招式时，ＨＰ会减少。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('088', 'Download', '下载', '比较对手的防御和特防，根据较低的那项能力相应地提高自己的攻击或特攻。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('089', 'Iron Fist', '铁拳', '使用拳类招式的威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('090', 'Poison Heal', '毒疗', '变为中毒状态时，ＨＰ不会减少，反而会增加起来。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('091', 'Adaptability', '适应力', '与自身同属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('092', 'Skill Link', '连续攻击', '如果使用连续招式，总是能使出最高次数。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('093', 'Hydration', '湿润之躯', '下雨天气时，异常状态会治愈。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('094', 'Solar Power', '太阳之力', '晴朗天气时，特攻会提高，而每回合ＨＰ会减少。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('095', 'Quick Feet', '飞毛腿', '变为异常状态时，速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('096', 'Normalize', '一般皮肤', '无论是什么属性的招式，全部会变为一般属性。威力会少量提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('097', 'Sniper', '狙击手', '击中要害时，威力会变得更强。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('098', 'Magic Guard', '魔法防守', '不会受到攻击以外的伤害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('099', 'No Guard', '无防守', '由于无防守战术，双方使出的招式都必定会击中。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('100', 'Stall', '慢出', '使出招式的顺序必定会变为最后。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('101', 'Technician', '技术高手', '攻击时可以将低威力招式的威力提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('102', 'Leaf Guard', '叶子防守', '晴朗天气时，不会变为异常状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('103', 'Klutz', '笨拙', '无法使用持有的道具。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('104', 'Mold Breaker', '破格', '可以不受对手特性的干扰，向对手使出招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('105', 'Super Luck', '超幸运', '因为拥有超幸运，攻击容易击中对手的要害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('106', 'Aftermath', '引爆', '变为濒死时，会对接触到自己的对手造成伤害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('107', 'Anticipation', '危险预知', '可以察觉到对手拥有的危险招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('108', 'Forewarn', '预知梦', '出场时，只读取１个对手拥有的招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('109', 'Unaware', '纯朴', '可以无视对手能力的变化，进行攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('110', 'Tinted Lens', '有色眼镜', '可以将效果不好的招式以通常的威力使出。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('111', 'Filter', '过滤', '受到效果绝佳的攻击时，可以减弱其威力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('112', 'Slow Start', '慢启动', '在５回合内，攻击和速度减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('113', 'Scrappy', '胆量', '一般属性和格斗属性的招式可以击中幽灵属性的宝可梦。对威吓也毫不动摇。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('114', 'Storm Drain', '引水', '将水属性的招式引到自己身上，不会受到伤害，而是会提高特攻。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('115', 'Ice Body', '冰冻之躯', '下雪天气时，会缓缓回复ＨＰ。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('116', 'Solid Rock', '坚硬岩石', '受到效果绝佳的攻击时，可以减弱其威力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('117', 'Snow Warning', '降雪', '出场时，会将天气变为下雪。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('118', 'Honey Gather', '采蜜', '战斗结束时，有时候会捡来甜甜蜜。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('119', 'Frisk', '察觉', '出场时，可以察觉对手的持有物。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('120', 'Reckless', '舍身', '自己会因反作用力受伤的招式，其威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('121', 'Multitype', '多属性', '自己的属性会根据持有的石板而改变。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('122', 'Flower Gift', '花之礼', '晴朗天气时，自己与同伴的攻击和特防能力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('123', 'Bad Dreams', '梦魇', '给予睡眠状态的对手伤害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('124', 'Pickpocket', '顺手牵羊', '盗取接触到自己的对手的道具。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('125', 'Sheer Force', '强行', '招式的追加效果消失，但因此能以更高的威力使出招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('126', 'Contrary', '唱反调', '能力的变化发生逆转，原本提高时会降低，而原本降低时会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('127', 'Unnerve', '紧张感', '让对手紧张，使其无法食用树果。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('128', 'Defiant', '不服输', '被对手降低能力时，攻击会大幅提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('129', 'Defeatist', '软弱', 'ＨＰ减半时，会变得软弱，攻击和特攻会减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('130', 'Cursed Body', '诅咒之躯', '受到攻击时，有时会把对手的招式变为定身法状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('131', 'Healer', '治愈之心', '有时会治愈异常状态的同伴。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('132', 'Friend Guard', '友情防守', '可以减少我方的伤害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('133', 'Weak Armor', '碎裂铠甲', '受到物理招式的伤害时，防御会降低，速度会大幅提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('134', 'Heavy Metal', '重金属', '自身的重量会变为２倍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('135', 'Light Metal', '轻金属', '自身的重量会减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('136', 'Multiscale', '多重鳞片', 'ＨＰ全满时，受到的伤害会变少。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('137', 'Toxic Boost', '中毒激升', '变为中毒状态时，物理招式的威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('138', 'Flare Boost', '受热激升', '变为灼伤状态时，特殊招式的威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('139', 'Harvest', '收获', '可以多次制作出已被使用掉的树果。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('140', 'Telepathy', '心灵感应', '读取我方的攻击，并闪避其招式伤害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('141', 'Moody', '心情不定', '每一回合，能力中的某项会大幅提高，而某项会降低。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('142', 'Overcoat', '防尘', '不会受到沙暴的伤害。也不会受到粉末类和孢子类招式的影响。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('143', 'Poison Touch', '毒手', '只通过接触就有可能让对手变为中毒状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('144', 'Regenerator', '再生力', '退回同行队伍后，ＨＰ会少量回复。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('145', 'Big Pecks', '健壮胸肌', '不会受到防御降低的效果。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('146', 'Sand Rush', '拨沙', '沙暴天气时，速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('147', 'Wonder Skin', '奇迹皮肤', '成为不易受到变化招式攻击的身体。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('148', 'Analytic', '分析', '如果在最后使出招式，招式的威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('149', 'Illusion', '幻觉', '假扮成同行队伍中的最后一只宝可梦出场，迷惑对手。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('150', 'Imposter', '变身者', '变身为当前面对的宝可梦。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('151', 'Infiltrator', '穿透', '可以穿透对手的壁障或替身进行攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('152', 'Mummy', '木乃伊', '被对手接触到后，会将对手变为木乃伊。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('153', 'Moxie', '自信过度', '如果打倒对手，就会充满自信，攻击会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('154', 'Justified', '正义之心', '受到恶属性的招式攻击时，因为正义感，攻击会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('155', 'Rattled', '胆怯', '受到恶属性、幽灵属性和虫属性的攻击或威吓时，会因胆怯而速度提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('156', 'Magic Bounce', '魔法镜', '可以不受到由对手使出的变化招式影响，并将其反弹。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('157', 'Sap Sipper', '食草', '受到草属性的招式攻击时，不会受到伤害，而是攻击会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('158', 'Prankster', '恶作剧之心', '可以率先使出变化招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('159', 'Sand Force', '沙之力', '沙暴天气时，岩石属性、地面属性和钢属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('160', 'Iron Barbs', '铁刺', '用铁刺给予接触到自己的对手伤害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('161', 'Zen Mode', '达摩模式', 'ＨＰ变为一半以下时，样子会改变。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('162', 'Victory Star', '胜利之星', '自己和同伴的命中率会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('163', 'Turboblaze', '涡轮火焰', '可以不受对手特性的干扰，向对手使出招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('164', 'Teravolt', '兆级电压', '可以不受对手特性的干扰，向对手使出招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('165', 'Aroma Veil', '芳香幕', '可以防住向自己和同伴发出的心灵攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('166', 'Flower Veil', '花幕', '我方的草属性宝可梦能力不会降低，也不会变为异常状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('167', 'Cheek Pouch', '颊囊', '无论是哪种树果，食用后，ＨＰ都会回复。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('168', 'Protean', '变幻自如', '变为与自己使出的招式相同的属性。每次出场战斗仅生效一次。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('169', 'Fur Coat', '毛皮大衣', '对手给予的物理招式的伤害会减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('170', 'Magician', '魔术师', '夺走被自己的招式击中的对手的道具。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('171', 'Bulletproof', '防弹', '可以防住对手的球和弹类招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('172', 'Competitive', '好胜', '如果被对手降低能力，特攻会大幅提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('173', 'Strong Jaw', '强壮之颚', '因为颚部强壮，啃咬类招式的威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('174', 'Refrigerate', '冰冻皮肤', '一般属性的招式会变为冰属性。威力会少量提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('175', 'Sweet Veil', '甜幕', '自己和同伴的宝可梦不会变为睡眠状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('176', 'Stance Change', '战斗切换', '如果使出攻击招式，会变为刀剑形态，如果使出招式“王者盾牌”，会变为盾牌形态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('177', 'Gale Wings', '疾风之翼', 'ＨＰ全满时，飞行属性的招式可以率先使出。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('178', 'Mega Launcher', '超级发射器', '波动和波导类招式的威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('179', 'Grass Pelt', '草之毛皮', '在青草场地时，防御会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('180', 'Symbiosis', '共生', '同伴使用道具时，会把自己持有的道具传递给同伴。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('181', 'Tough Claws', '硬爪', '接触到对手的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('182', 'Pixilate', '妖精皮肤', '一般属性的招式会变为妖精属性。威力会少量提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('183', 'Gooey', '黏滑', '对于用攻击接触到自己的对手，会降低其速度。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('184', 'Aerilate', '飞行皮肤', '一般属性的招式会变为飞行属性。威力会少量提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('185', 'Parental Bond', '亲子爱', '亲子俩可以合计攻击２次。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('186', 'Dark Aura', '暗黑气场', '全体的恶属性招式变强。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('187', 'Fairy Aura', '妖精气场', '全体的妖精属性招式变强。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('188', 'Aura Break', '气场破坏', '让气场的效果发生逆转，降低威力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('189', 'Primordial Sea', '始源之海', '变为不会受到火属性攻击的天气。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('190', 'Desolate Land', '终结之地', '变为不会受到水属性攻击的天气。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('191', 'Delta Stream', '德尔塔气流', '变为令飞行属性的弱点消失的天气。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('192', 'Stamina', '持久力', '受到攻击时，防御会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('193', 'Wimp Out', '跃跃欲逃', 'ＨＰ变为一半时，会慌慌张张逃走，退回同行队伍中。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('194', 'Emergency Exit', '危险回避', 'ＨＰ变为一半时，为了回避危险，会退回到同行队伍中。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('195', 'Water Compaction', '遇水凝固', '受到水属性的招式攻击时，防御会大幅提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('196', 'Merciless', '不仁不义', '攻击中毒状态的对手时，必定会击中要害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('197', 'Shields Down', '界限盾壳', 'ＨＰ变为一半时，壳会坏掉，变得有攻击性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('198', 'Stakeout', '蹲守', '可以对替换出场的对手以２倍的伤害进行攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('199', 'Water Bubble', '水泡', '降低自己受到的火属性招式的威力，不会灼伤。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('200', 'Steelworker', '钢能力者', '钢属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('201', 'Berserk', '怒火冲天', '因对手的攻击ＨＰ变为一半时，特攻会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('202', 'Slush Rush', '拨雪', '下雪天气时，速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('203', 'Long Reach', '远隔', '可以不接触对手就使出所有的招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('204', 'Liquid Voice', '湿润之声', '所有的声音招式都变为水属性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('205', 'Triage', '先行治疗', '可以率先使出回复招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('206', 'Galvanize', '电气皮肤', '一般属性的招式会变为电属性。威力会少量提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('207', 'Surge Surfer', '冲浪之尾', '电气场地时，速度会变为２倍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('208', 'Schooling', '鱼群', 'ＨＰ多的时候会聚起来变强。ＨＰ剩余量变少时，群体会分崩离析。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('209', 'Disguise', '画皮', '通过画皮覆盖住身体，可以防住１次攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('210', 'Battle Bond', '牵绊变身', '打倒对手时，与训练家的牵绊会增强，自己的攻击、特攻、速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('211', 'Power Construct', '群聚变形', 'ＨＰ变为一半时，细胞们会赶来支援，变为完全体形态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('212', 'Corrosion', '腐蚀', '可以使钢属性和毒属性的宝可梦也陷入中毒状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('213', 'Comatose', '绝对睡眠', '总是半梦半醒的状态，绝对不会醒来。可以就这么睡着进行攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('214', 'Queenly Majesty', '女王的威严', '向对手施加威慑力，使其无法对我方使出先制招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('215', 'Innards Out', '飞出的内在物', '被对手打倒的时候，会给予对手相当于ＨＰ剩余量的伤害。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('216', 'Dancer', '舞者', '有谁使出跳舞招式时，自己也能就这么接着使出跳舞招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('217', 'Battery', '蓄电池', '会提高我方的特殊招式的威力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('218', 'Fluffy', '毛茸茸', '会将对手所给予的接触类招式的伤害减半，但火属性招式的伤害会变为２倍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('219', 'Dazzling', '鲜艳之躯', '让对手吓一跳，使其无法对我方使出先制招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('220', 'Soul-Heart', '魂心', '宝可梦每次变为濒死状态时，特攻会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('221', 'Tangling Hair', '卷发', '对于用攻击接触到自己的对手，会降低其速度。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('222', 'Receiver', '接球手', '继承被打倒的同伴的特性，变为相同的特性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('223', 'Power of Alchemy', '化学之力', '继承被打倒的同伴的特性，变为相同的特性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('224', 'Beast Boost', '异兽提升', '打倒对手的时候，自己最高的那项能力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('225', 'RKS System', 'ＡＲ系统', '根据持有的存储碟，自己的属性会改变。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('226', 'Electric Surge', '电气制造者', '出场时，会布下电气场地。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('227', 'Psychic Surge', '精神制造者', '出场时，会布下精神场地。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('228', 'Misty Surge', '薄雾制造者', '出场时，会布下薄雾场地。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('229', 'Grassy Surge', '青草制造者', '出场时，会布下青草场地。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('230', 'Full Metal Body', '金属防护', '不会因为对手的招式或特性而被降低能力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('231', 'Shadow Shield', '幻影防守', 'ＨＰ全满时，受到的伤害会变少。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('232', 'Prism Armor', '棱镜装甲', '受到效果绝佳的攻击时，可以减弱其威力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('233', 'Neuroforce', '脑核之力', '效果绝佳的攻击，威力会变得更强。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('234', 'Intrepid Sword', '不挠之剑', '首次出场时，攻击会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('235', 'Dauntless Shield', '不屈之盾', '首次出场时，防御会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('236', 'Libero', '自由者', '变为与自己使出的招式相同的属性。每次出场战斗仅生效一次。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('237', 'Ball Fetch', '捡球', '没有携带道具时，会拾取第１个投出后捕捉失败的精灵球。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('238', 'Cotton Down', '棉絮', '受到攻击后撒下棉絮，降低除自己以外的所有宝可梦的速度。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('239', 'Propeller Tail', '螺旋尾鳍', '能无视具有吸引对手招式效果的特性或招式的影响。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('240', 'Mirror Armor', '镜甲', '只反弹自己受到的能力降低效果。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('241', 'Gulp Missile', '一口导弹', '冲浪或潜水时会叼来猎物。受到伤害时，会吐出猎物进行攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('242', 'Stalwart', '坚毅', '能无视具有吸引对手招式效果的特性或招式的影响。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('243', 'Steam Engine', '蒸汽机', '受到水属性或火属性的招式攻击时，速度会巨幅提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('244', 'Punk Rock', '庞克摇滚', '声音招式的威力会提高。受到的声音招式伤害会减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('245', 'Sand Spit', '吐沙', '受到攻击时，会刮起沙暴。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('246', 'Ice Scales', '冰鳞粉', '由于有冰鳞粉的守护，受到的特殊攻击伤害会减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('247', 'Ripen', '熟成', '使树果成熟，效果变为２倍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('248', 'Ice Face', '结冻头', '头部的冰会代替自己承受物理攻击，但是样子会改变。下雪时，冰会恢复原状。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('249', 'Power Spot', '能量点', '只要处在相邻位置，招式的威力就会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('250', 'Mimicry', '拟态', '宝可梦的属性会根据场地的状态而变化。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('251', 'Screen Cleaner', '除障', '出场时，敌方和我方的光墙、反射壁和极光幕的效果会消失。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('252', 'Steely Spirit', '钢之意志', '我方的钢属性攻击威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('253', 'Perish Body', '灭亡之躯', '受到接触类招式攻击时，双方都会在３回合后变为濒死状态。替换后效果消失。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('254', 'Wandering Spirit', '游魂', '与使用接触类招式攻击自己的宝可梦互换特性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('255', 'Gorilla Tactics', '一猩一意', '虽然攻击会提高，但是只能使出一开始所选的招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('256', 'Neutralizing Gas', '化学变化气体', '特性为化学变化气体的宝可梦在场时，场上所有宝可梦的特性效果都会消失或者无法生效。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('257', 'Pastel Veil', '粉彩护幕', '自己和同伴都不会陷入中毒的异常状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('258', 'Hunger Switch', '饱了又饿', '每回合结束时会在满腹花纹与空腹花纹之间交替改变样子。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('259', 'Quick Draw', '速击', '有时能比对手先一步行动。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('260', 'Unseen Fist', '无形拳', '如果使出的是接触到对手的招式，就可以无视守护效果进行攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('261', 'Curious Medicine', '怪药', '出场时会从贝壳撒药，将我方的能力变化复原。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('262', 'Transistor', '电晶体', '电属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('263', 'Dragon''s Maw', '龙颚', '龙属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('264', 'Chilling Neigh', '苍白嘶鸣', '打倒对手时会用冰冷的声音嘶鸣并提高攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('265', 'Grim Neigh', '漆黑嘶鸣', '打倒对手时会用恐怖的声音嘶鸣并提高特攻。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('266', 'As One', '人马一体', '兼备蕾冠王的紧张感和雪暴马的苍白嘶鸣这两种特性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('267', 'As One', '人马一体', '兼备蕾冠王的紧张感和灵幽马的漆黑嘶鸣这两种特性。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('268', 'Lingering Aroma', '甩不掉的气味', '被对手接触到后，甩不掉的气味会沾上对手。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('269', 'Seed Sower', '掉出种子', '受到攻击时，会将脚下变成青草场地。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('270', 'Thermal Exchange', '热交换', '受到火属性的招式攻击时，攻击会提高，且不会陷入灼伤状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('271', 'Anger Shell', '愤怒甲壳', '因被对手攻击而ＨＰ变为一半时，会因愤怒降低防御和特防。但攻击、特攻、速度会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('272', 'Purifying Salt', '洁净之盐', '因洁净的盐而不会陷入异常状态。会让幽灵属性的招式伤害减半。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('273', 'Well-Baked Body', '焦香之躯', '受到火属性的招式攻击时，不会受到伤害，而是会大幅提高防御。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('274', 'Wind Rider', '乘风', '吹起了顺风或受到风的招式攻击时，不会受到伤害，而是会提高攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('275', 'Guard Dog', '看门犬', '受到威吓时，攻击会提高。让替换宝可梦的招式和道具无效。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('276', 'Rocky Payload', '搬岩', '岩石属性的招式威力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('277', 'Wind Power', '风力发电', '受到风的招式攻击时，会变为充电状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('278', 'Zero to Hero', '全能变身', '回到同行队伍后，会变为全能形态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('279', 'Commander', '发号施令', '出场时，若我方当中有吃吼霸，就会进入其口中，并从其口中发出指令。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('280', 'Electromorphosis', '电力转换', '受到伤害时，会变为充电状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('281', 'Protosynthesis', '古代活性', '携带着驱劲能量或天气为晴朗时，数值最高的能力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('282', 'Quark Drive', '夸克充能', '携带着驱劲能量或在电气场地上时，数值最高的能力会提高。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('283', 'Good as Gold', '黄金之躯', '不会氧化的坚固黄金身躯不会受到对手的变化招式的影响。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('284', 'Vessel of Ruin', '灾祸之鼎', '以能呼唤灾厄的鼎的力量降低除自己以外的宝可梦的特攻。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('285', 'Sword of Ruin', '灾祸之剑', '以能呼唤灾厄的剑的力量降低除自己以外的宝可梦的防御。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('286', 'Tablets of Ruin', '灾祸之简', '以能呼唤灾厄的简的力量降低除自己以外的宝可梦的攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('287', 'Beads of Ruin', '灾祸之玉', '以能呼唤灾厄的勾玉的力量降低除自己以外的宝可梦的特防。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('288', 'Orichalcum Pulse', '绯红脉动', '出场时，会将天气变为晴朗。日照强烈时，会通过古代的脉动升高攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('289', 'Hadron Engine', '强子引擎', '出场时，会布下电气场地。处于电气场地时，会通过未来的机关升高特攻。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('290', 'Opportunist', '跟风', '对手的能力提高时，自己也会趁机同样地提高能力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('291', 'Cud Chew', '反刍', '吃了树果后，会在下一回合结束时从胃反刍出来再吃１次。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('292', 'Sharpness', '锋锐', '提高切割对手的招式的威力。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('293', 'Supreme Overlord', '大将', '出场时，攻击和特攻会按照目前被打倒的同伴数量逐渐提升，被打倒越多，提升越多。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('294', 'Costar', '同台共演', '出场时，复制同伴的能力变化。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('295', 'Toxic Debris', '毒满地', '受到物理招式的伤害时，会在对手脚下散布毒菱。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('296', 'Armor Tail', '尾甲', '包裹头部的神秘尾巴使对手无法对我方使出先制招式。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('297', 'Earth Eater', '食土', '受到地面属性的招式攻击时，不会受到伤害，而是会得到回复。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('298', 'Mycelium Might', '菌丝之力', '使出变化招式时，虽然行动必定会变慢，但能不受对手的特性妨碍。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('299', 'Hospitality', '款待', '出场时款待同伴，回复其少量ＨＰ。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('300', 'Mind''s Eye', '心眼', '一般属性和格斗属性的招式可以击中幽灵属性的宝可梦。无视对手的闪避率的变化，且命中率不会被降低。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('301', 'Embody Aspect', '面影辉映', '将回忆映于心中，让碧草面具发出光辉，提高自己的速度。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('302', 'Embody Aspect', '面影辉映', '将回忆映于心中，让火灶面具发出光辉，提高自己的攻击。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('303', 'Embody Aspect', '面影辉映', '将回忆映于心中，让水井面具发出光辉，提高自己的特防。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('304', 'Embody Aspect', '面影辉映', '将回忆映于心中，让础石面具发出光辉，提高自己的防御。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('305', 'Toxic Chain', '毒锁链', '凭借含有毒素的锁链的力量，有时能让被招式击中的对手陷入剧毒状态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('306', 'Supersweet Syrup', '甘露之蜜', '首次出场时，会散发出甜腻的蜜的香味来降低对手的闪避率。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('307', 'Tera Shift', '太晶变形', '出场时，会吸收周围的能量，变为太晶形态。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('308', 'Tera Shell', '太晶甲壳', '甲壳蕴藏着全部属性的力量，会将自己ＨＰ全满时受到的伤害全都变为效果不好。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('309', 'Teraform Zero', '归零化境', '太乐巴戈斯变为星晶形态时，蕴藏在它身上的力量会将天气和场地的影响全部归零。');
INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ('310', 'Poison Puppeteer', '毒傀儡', '因桃歹郎的招式而陷入中毒状态的对手同时也会陷入混乱状态。');
