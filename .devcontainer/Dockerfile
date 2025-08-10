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
    apt-transport-https \
    gnupg \
    lsb-release \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*


# Docker CLI をインストール（docker-outside-of-docker方式）
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null \
    && apt-get update \
    && apt-get -y install docker-ce-cli \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# GitHub CLI をインストール
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg \
    && chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
    && apt-get update \
    && apt-get install gh -y \
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

# dockerグループを作成し、vscode ユーザーをdockerグループに追加
# daemonグループ（GID=1）を作成し、vscode ユーザーをdaemonグループに追加
RUN groupadd docker || true \
    && usermod -aG docker vscode \
    && groupadd -g 1 daemon || true \
    && usermod -aG daemon vscode

# Claude Code のインストール（NPM版）
RUN npm install -g @anthropic-ai/claude-code

USER vscode
WORKDIR /workspaces/knowledge