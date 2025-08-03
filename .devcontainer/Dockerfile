# Dev Container用のDockerfile
FROM mcr.microsoft.com/devcontainers/base:ubuntu

# 必要なパッケージをインストール
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    git \
    curl \
    wget \
    ca-certificates \
    build-essential \
    postgresql-client \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# Go のインストール (Go 1.21)
RUN wget -O go.tar.gz https://go.dev/dl/go1.21.8.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz

# Node.js のインストール (Node.js 18)
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
    && apt-get install -y nodejs

# 環境変数の設定
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GOROOT="/usr/local/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Go のディレクトリを作成
RUN mkdir -p /go/bin /go/src /go/pkg

# vscode ユーザーでの作業用にディレクトリ作成
RUN mkdir -p /workspaces && chown -R vscode:vscode /workspaces
RUN mkdir -p /go && chown -R vscode:vscode /go

# Claude Code のインストール（NPM版）
RUN npm install -g @anthropic-ai/claude-code

USER vscode
WORKDIR /workspaces/knowledge