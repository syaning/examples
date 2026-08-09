# deep-agents × 通义千问（最小 Demo）

基于 [`deepagents`](https://github.com/langchain-ai/deepagents)，通过阿里云百炼的 **OpenAI 兼容接口** 调用通义千问。

## 准备

1. 在[阿里云百炼控制台](https://bailian.console.aliyun.com/)创建 API Key
2. 安装依赖：

```bash
uv sync
```

3. 配置 `.env`（必填项见 `.env.example`）：

```bash
cp .env.example .env
# 编辑 .env，至少填入 DASHSCOPE_API_KEY
```

三项均从环境读取，无代码内默认值：`DASHSCOPE_API_KEY`、`DASHSCOPE_BASE_URL`、`QWEN_MODEL`。
已导出的同名环境变量会覆盖 `.env`。

## 运行

```bash
uv run deep-agents
uv run deep-agents 今天北京天气怎么样？用一句话回答
```

核心配置就是三处：`api_key`、`base_url`、`model`。

```python
from langchain_openai import ChatOpenAI
from deepagents import create_deep_agent

model = ChatOpenAI(
    model="qwen-plus",
    api_key="sk-xxx",
    base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
)

agent = create_deep_agent(model=model, system_prompt="你是一个简洁的中文助手。")
result = agent.invoke({"messages": [{"role": "user", "content": "你好"}]})
```
