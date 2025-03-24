from openai import OpenAI
import os
from pathlib import Path
import time


with open('template.md', 'r', encoding='utf-8') as f:
    template_data = f.read()

system_promot = "你是一个专业的命令行工具的文档生成器，你需要根据用户的命令行工具的描述，生成对应的命令行工具的文档。现在要生成git命令行工具的文档。严格参考以下模板,不要自行添加总结性文字,例如通过以上命令，可以更方便地管理和查看 Git 仓库中的标签。\n" + template_data

client = OpenAI(api_key="xxx", base_url="https://api.deepseek.com")

def get_content(user_promot):
    response = client.chat.completions.create(
        model="deepseek-chat",
        messages=[
            {"role": "system", "content": system_promot},
            {"role": "user", "content": user_promot},
        ],
        stream=True
    )
    content_list = []
    for chunk in response:
        print(chunk.choices[0].delta.content, end="", flush=True)
        content_list.append(chunk.choices[0].delta.content)
    return "".join(content_list)

def generate_content(md_file):
    client = get_client()
    with open(md_file, 'r', encoding='utf-8') as f:
        md_data = f.read()
    if "命令介绍" in md_data:
        return
    user_promot = "请完善该文档:\n"+"file_name:"+md_file.name+"\n"+"conten:\n"+md_data
    content = get_content(user_promot)
    with open(md_file, 'w', encoding='utf-8') as f:
        f.write(content)

def generate(dir_path):
    md_list = Path(dir_path).glob("*.md")
    for inddex,md_file in enumerate(md_list):
        print("开始生成第"+str(inddex)+"个文档："+md_file.name)
        generate_content(md_file)
        print("生成第"+str(inddex)+"个文档完成")

if __name__ == "__main__":
    generate("../docs/zh")