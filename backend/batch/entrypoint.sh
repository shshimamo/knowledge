#!/bin/sh

# 実行例: docker run -it knowledge-batch "/go/bin/echo hoge"
# The first argument is the command
CMD=$1

# Check if CMD is empty
if [ -z "$CMD" ]; then
  echo "No command specified"
  exit 1
fi

# Check if CMD_NAME starts with /go/bin/
if [ "${CMD#"/go/bin/"}" = "$CMD" ]; then
  # プレフィックス /go/bin/ を取り除いた結果が同じなら、/go/bin/ で始まっていない
  echo "Command must start with /go/bin/"
  exit 1
fi

echo "Executing command: $CMD"

# Execute the command
$CMD