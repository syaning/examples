"""Minimal deepagents demo using Alibaba Cloud Qwen via OpenAI-compatible API."""

import os
import sys
from pathlib import Path

from deepagents import create_deep_agent
from dotenv import load_dotenv
from langchain_openai import ChatOpenAI

# Project root: .../deep-agents (src/deep_agents/__init__.py -> parents[2])
_ENV_FILE = Path(__file__).resolve().parents[2] / ".env"

def _require_env(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise SystemExit(
            f"缺少 {name}。\n"
            "请在项目根目录的 .env 中配置，可参考 .env.example。"
        )
    return value


def _build_model() -> ChatOpenAI:
    return ChatOpenAI(
        model=_require_env("QWEN_MODEL"),
        api_key=_require_env("DASHSCOPE_API_KEY"),
        base_url=_require_env("DASHSCOPE_BASE_URL"),
        temperature=0.3,
    )


def main() -> None:
    load_dotenv(_ENV_FILE)
    prompt = (
        " ".join(sys.argv[1:]).strip() or "用一句话介绍你自己，并说明你基于哪个模型。"
    )

    agent = create_deep_agent(
        model=_build_model(),
        system_prompt="你是一个简洁、友好的中文助手。回答尽量短。",
    )

    result = agent.invoke({"messages": [{"role": "user", "content": prompt}]})
    messages = result.get("messages", [])
    if not messages:
        print(result)
        return

    last = messages[-1]
    content = getattr(last, "content", None) or str(last)
    print(content)


if __name__ == "__main__":
    main()
