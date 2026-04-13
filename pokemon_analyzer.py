#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Pokemon Dataset Analyzer - Extract data from pokemon-dataset-zh and generate SQL
This script processes JSON files and generates SQL INSERT statements for Pokemon database.
"""

import json
import os
import sys
from pathlib import Path
from typing import Dict, List, Optional, Tuple, Any

# ==================== Type Mapping ====================
# Map Chinese type names to English
TYPE_MAPPING = {
    "一般": "Normal",
    "火": "Fire",
    "水": "Water",
    "电": "Electric",
    "草": "Grass",
    "冰": "Ice",
    "格斗": "Fighting",
    "毒": "Poison",
    "地面": "Ground",
    "飞行": "Flying",
    "超能力": "Psychic",
    "虫": "Bug",
    "岩石": "Rock",
    "幽灵": "Ghost",
    "龙": "Dragon",
    "恶": "Dark",
    "钢": "Steel",
    "妖精": "Fairy",
}

# Map Chinese category names to English
CATEGORY_MAPPING = {
    "物理": "Physical",
    "特殊": "Special",
    "变化": "Status",
}


class PokemonAnalyzer:
    """Main analyzer class for processing Pokemon dataset"""

    def __init__(self, data_dir: str):
        """Initialize the analyzer with dataset directory"""
        self.data_dir = Path(data_dir)
        self.pokemon_species: List[Dict] = []
        self.pokemon_moves: List[Dict] = []
        self.pokemon_abilities: List[Dict] = []
        self.type_map: Dict[str, int] = {}  # Map type names to IDs
        self.ability_map: Dict[str, int] = {}  # Map ability names to IDs

    def escape_sql_string(self, s: Optional[str]) -> str:
        """Escape string for SQL"""
        if s is None:
            return "NULL"
        # Escape single quotes
        escaped = str(s).replace("'", "''")
        return f"'{escaped}'"

    def escape_sql_int(self, val: Optional[Any]) -> str:
        """Escape integer for SQL"""
        if val is None or val == "" or val == "NULL":
            return "NULL"
        try:
            return str(int(val))
        except (ValueError, TypeError):
            return "NULL"

    def load_simple_pokedex(self) -> bool:
        """Load basic Pokemon data from simple_pokedex.json"""
        pokedex_file = self.data_dir / "simple_pokedex.json"

        if not pokedex_file.exists():
            print(f"Error: {pokedex_file} not found")
            return False

        try:
            with open(pokedex_file, "r", encoding="utf-8") as f:
                simple_pokedex = json.load(f)

            # Load individual Pokemon files to get full details
            pokemon_dir = self.data_dir / "pokemon"
            if not pokemon_dir.exists():
                print(f"Error: {pokemon_dir} not found")
                return False

            for entry in simple_pokedex:
                pokemon_id = entry.get("index", "0")
                name_en = entry.get("name_en", "Unknown")
                name_zh = entry.get("name_zh", "未知")

                # Try to load individual Pokemon file for more details
                pokemon_file = pokemon_dir / f"{pokemon_id}-{name_zh}.json"

                types = []
                capture_rate = None
                base_stats = {}

                if pokemon_file.exists():
                    try:
                        with open(pokemon_file, "r", encoding="utf-8") as f:
                            pokemon_data = json.load(f)

                            # Extract types from forms
                            if (
                                "forms" in pokemon_data
                                and len(pokemon_data["forms"]) > 0
                            ):
                                form = pokemon_data["forms"][0]
                                types = form.get("types", [])

                                # Extract catch rate
                                catch_rate_str = form.get("catch_rate", "")
                                if catch_rate_str:
                                    # Parse "45（5.9%）" format
                                    parts = catch_rate_str.split("（")
                                    if parts:
                                        try:
                                            capture_rate = int(parts[0].strip())
                                        except ValueError:
                                            pass
                    except Exception as e:
                        print(f"Warning: Could not parse {pokemon_file}: {e}")

                # Normalize types to English
                types_en = [TYPE_MAPPING.get(t, t) for t in types]

                pokemon_entry = {
                    "id": pokemon_id,
                    "name_en": name_en,
                    "name_zh": name_zh,
                    "types": types_en,
                    "capture_rate": capture_rate,
                }

                self.pokemon_species.append(pokemon_entry)

            print(f"✓ Loaded {len(self.pokemon_species)} Pokemon species")
            return True

        except Exception as e:
            print(f"Error loading simple_pokedex.json: {e}")
            return False

    def load_moves(self) -> bool:
        """Load move data from move_list.json"""
        moves_file = self.data_dir / "move_list.json"

        if not moves_file.exists():
            print(f"Error: {moves_file} not found")
            return False

        try:
            with open(moves_file, "r", encoding="utf-8") as f:
                moves_data = json.load(f)

            for move in moves_data:
                move_id = move.get("id", "0")
                name_en = move.get("name_en", "Unknown")
                name_zh = move.get("name_zh", "未知")
                move_type = move.get("type", "")
                category = move.get("category", "")
                power = move.get("power", None)
                accuracy = move.get("accuracy", None)
                pp = move.get("pp", None)

                # Normalize type and category
                type_en = TYPE_MAPPING.get(move_type, move_type)
                category_en = CATEGORY_MAPPING.get(category, category)

                move_entry = {
                    "id": move_id,
                    "name_en": name_en,
                    "name_zh": name_zh,
                    "type": type_en,
                    "category": category_en,
                    "power": self.escape_sql_int(power),
                    "accuracy": self.escape_sql_int(accuracy),
                    "pp": self.escape_sql_int(pp),
                }

                self.pokemon_moves.append(move_entry)

            print(f"✓ Loaded {len(self.pokemon_moves)} moves")
            return True

        except Exception as e:
            print(f"Error loading move_list.json: {e}")
            return False

    def load_abilities(self) -> bool:
        """Load ability data from ability_list.json"""
        abilities_file = self.data_dir / "ability_list.json"

        if not abilities_file.exists():
            print(f"Error: {abilities_file} not found")
            return False

        try:
            with open(abilities_file, "r", encoding="utf-8") as f:
                abilities_data = json.load(f)

            for ability in abilities_data:
                ability_id = ability.get("id", "0")
                name_en = ability.get("name_en", "Unknown")
                name_zh = ability.get("name_zh", "未知")
                description = ability.get("description", "")

                ability_entry = {
                    "id": ability_id,
                    "name_en": name_en,
                    "name_zh": name_zh,
                    "description": description,
                }

                self.pokemon_abilities.append(ability_entry)

            print(f"✓ Loaded {len(self.pokemon_abilities)} abilities")
            return True

        except Exception as e:
            print(f"Error loading ability_list.json: {e}")
            return False

    def generate_pokemon_species_sql(self) -> List[str]:
        """Generate SQL INSERT statements for pokemon_species table"""
        sql_statements = []

        sql_statements.append("""-- Pokemon Species Data
CREATE TABLE IF NOT EXISTS pokemon_species (
    pokemon_id VARCHAR(10) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    types TEXT,
    capture_rate INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

""")

        for pokemon in self.pokemon_species:
            pokemon_id = pokemon["id"]
            name_en = self.escape_sql_string(pokemon["name_en"])
            name_zh = self.escape_sql_string(pokemon["name_zh"])
            types = self.escape_sql_string(",".join(pokemon["types"]))
            capture_rate = self.escape_sql_int(pokemon["capture_rate"])

            sql = f"""INSERT INTO pokemon_species (pokemon_id, name_en, name_zh, types, capture_rate) 
VALUES ({self.escape_sql_string(pokemon_id)}, {name_en}, {name_zh}, {types}, {capture_rate});"""
            sql_statements.append(sql)

        return sql_statements

    def generate_moves_sql(self) -> List[str]:
        """Generate SQL INSERT statements for pokemon_moves table"""
        sql_statements = []

        sql_statements.append("""-- Pokemon Moves Data
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

""")

        for move in self.pokemon_moves:
            move_id = move["id"]
            name_en = self.escape_sql_string(move["name_en"])
            name_zh = self.escape_sql_string(move["name_zh"])
            move_type = self.escape_sql_string(move["type"])
            category = self.escape_sql_string(move["category"])
            power = move["power"]
            accuracy = move["accuracy"]
            pp = move["pp"]

            sql = f"""INSERT INTO pokemon_moves (move_id, name_en, name_zh, type, category, power, accuracy, pp) 
VALUES ({self.escape_sql_string(move_id)}, {name_en}, {name_zh}, {move_type}, {category}, {power}, {accuracy}, {pp});"""
            sql_statements.append(sql)

        return sql_statements

    def generate_abilities_sql(self) -> List[str]:
        """Generate SQL INSERT statements for pokemon_abilities table"""
        sql_statements = []

        sql_statements.append("""-- Pokemon Abilities Data
CREATE TABLE IF NOT EXISTS pokemon_abilities (
    ability_id VARCHAR(10) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

""")

        for ability in self.pokemon_abilities:
            ability_id = ability["id"]
            name_en = self.escape_sql_string(ability["name_en"])
            name_zh = self.escape_sql_string(ability["name_zh"])
            description = self.escape_sql_string(ability["description"])

            sql = f"""INSERT INTO pokemon_abilities (ability_id, name_en, name_zh, description) 
VALUES ({self.escape_sql_string(ability_id)}, {name_en}, {name_zh}, {description});"""
            sql_statements.append(sql)

        return sql_statements

    def generate_all_sql(self) -> List[str]:
        """Generate all SQL statements"""
        all_sql = []

        # Add header
        all_sql.append("""-- =========================================
-- Pokemon Database Initialization Script
-- Auto-generated from pokemon-dataset-zh
-- =========================================
-- Set encoding
SET client_encoding = 'UTF8';

""")

        # Add all SQL statements
        all_sql.extend(self.generate_pokemon_species_sql())
        all_sql.append("")
        all_sql.extend(self.generate_moves_sql())
        all_sql.append("")
        all_sql.extend(self.generate_abilities_sql())

        return all_sql

    def save_sql_file(self, output_file: str) -> Tuple[bool, int]:
        """Save generated SQL to file"""
        try:
            all_sql = self.generate_all_sql()

            with open(output_file, "w", encoding="utf-8") as f:
                for statement in all_sql:
                    f.write(statement)
                    if statement and not statement.endswith("\n"):
                        f.write("\n")

            # Count INSERT statements
            insert_count = sum(
                1 for line in all_sql if line.strip().startswith("INSERT INTO")
            )

            print(f"\n✓ SQL file saved to: {output_file}")
            print(f"✓ Total SQL lines: {len(all_sql)}")
            print(f"✓ INSERT statements: {insert_count}")
            print(f"  - Pokemon species: {len(self.pokemon_species)}")
            print(f"  - Moves: {len(self.pokemon_moves)}")
            print(f"  - Abilities: {len(self.pokemon_abilities)}")

            return True, insert_count

        except Exception as e:
            print(f"Error saving SQL file: {e}")
            return False, 0

    def analyze(self, output_file: str) -> bool:
        """Run full analysis"""
        print("=" * 50)
        print("Pokemon Dataset Analyzer")
        print("=" * 50)

        # Load data
        if not self.load_simple_pokedex():
            return False

        if not self.load_moves():
            return False

        if not self.load_abilities():
            return False

        # Generate and save SQL
        success, insert_count = self.save_sql_file(output_file)

        if success:
            print("\n" + "=" * 50)
            print("✓ Analysis completed successfully!")
            print("=" * 50)
            return True

        return False


def main():
    """Main entry point"""
    # Configuration
    data_dir = "/tmp/pokemon-dataset-zh/data"
    output_file = "/root/petskill/judge-server/INIT_POKEMON_DATA.sql"

    # Verify input directory exists
    if not Path(data_dir).exists():
        print(f"Error: Data directory not found: {data_dir}")
        sys.exit(1)

    # Verify output directory exists
    output_path = Path(output_file)
    output_path.parent.mkdir(parents=True, exist_ok=True)

    # Run analyzer
    analyzer = PokemonAnalyzer(data_dir)
    if analyzer.analyze(output_file):
        sys.exit(0)
    else:
        sys.exit(1)


if __name__ == "__main__":
    main()
