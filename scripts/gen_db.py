import os
import json
from pathlib import Path
import yaml
from sqlalchemy import Column, Integer, String, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

Base = declarative_base()

# 定义数据库模型
class Command(Base):
    __tablename__ = 'commands'
    id = Column(Integer, primary_key=True)
    title = Column(String, nullable=False)  # title不能为空
    category = Column(String, nullable=False)  # category不能为空
    command = Column(String, default="")  # command不能为null，默认值为空字符串
    description = Column(String, default="")  # description不能为null，默认值为空字符串
    content = Column(String)

def generate_sql_file(db_path, sql_path):
    """生成SQL文件"""
    # 使用sqlite3命令行工具导出SQL
    import subprocess
    try:
        subprocess.run(f"sqlite3 {db_path} .dump > {sql_path}", shell=True, check=True)
        print(f"成功生成SQL文件: {sql_path}")
    except subprocess.CalledProcessError as e:
        print(f"生成SQL文件失败: {e}")

def main():
    # 使用Pathlib处理数据库路径
    db_path = Path(__file__).parent.parent / "resources" / "commands.db"
    sql_path = Path(__file__).parent.parent / "resources" / "commands.sql"
    
    # 创建数据库
    engine = create_engine(f'sqlite:///{db_path}')
    Base.metadata.create_all(engine)
    Session = sessionmaker(bind=engine)
    session = Session()

    # 获取所有目录
    docs_path = Path("../docs/zh")

    # 遍历每个目录
    for category_dir in docs_path.iterdir():
        if category_dir.is_dir():
            category_name = category_dir.name

            # 遍历目录下的markdown文件
            for md_file in category_dir.glob("*.md"):
                # 查找对应的meta文件
                meta_file = md_file.with_suffix(".meta")
                if meta_file.exists():
                    with open(meta_file, "r", encoding="utf-8") as f:
                        meta_data = yaml.safe_load(f)
                    
                    with open(md_file, "r", encoding="utf-8") as f:
                        md_content = f.read()
                    
                    command = Command(
                        category=category_name,
                        title=meta_data.get("title", ""),
                        command=meta_data.get("command", ""),
                        description=meta_data.get("description", ""),
                        content=md_content
                    )
                    session.add(command)

    # 提交事务
    session.commit()
    session.close()

    # 生成SQL文件
    generate_sql_file(db_path, sql_path)

if __name__ == "__main__":
    main()