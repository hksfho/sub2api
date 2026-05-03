#!/usr/bin/env bash
# 本地构建镜像的快速脚本，避免在命令行反复输入构建参数。

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

docker buildx build -t hksfho/hkaiproxy:latest \
    --platform linux/amd64,linux/arm64 \
    --push \
    -f "${REPO_ROOT}/Dockerfile" \
    "${REPO_ROOT}"
