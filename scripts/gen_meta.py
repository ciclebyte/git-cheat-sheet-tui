import os
import json
from pathlib import Path
import yaml
import re

def extract_command_info(content):
    # 提取命令介绍
    intro_pattern = r"## 一、命令介绍\n\n(.*?)\n\n## "
    intro_match = re.search(intro_pattern, content, re.DOTALL)
    
    # 提取命令格式
    format_pattern = r"### 命令格式\n\n```bash\n(.*?)\n```"
    format_match = re.search(format_pattern, content, re.DOTALL)
    
    return {
        "description": intro_match.group(1).strip() if intro_match else None,
        "command": format_match.group(1).strip() if format_match else None
    }

def main():
    md_list = Path("../docs/zh").glob("**/*.md")
    for md in md_list:
        meta_file_path = md.with_suffix(".meta")
        if not meta_file_path.exists():
            meta_file_path.unlink(missing_ok=True)      
        with open(md, 'r', encoding='utf-8') as f:
            content = f.read()
            info = extract_command_info(content)
            info["title"] = md.stem
        # 写入 YAML 元数据到.meta 文件
        with open(meta_file_path, 'w', encoding='utf-8') as meta_f:
            yaml.dump(info, meta_f, allow_unicode=True, default_flow_style=False, sort_keys=False)


if __name__ == "__main__":
    main()