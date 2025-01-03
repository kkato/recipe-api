FROM golang:1.23.0

WORKDIR /app

# go.mod と go.sum を先にコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# アプリのソースコードをコピー
COPY . .

# バイナリをビルド
RUN go build -o main

# アプリケーションを実行するコマンド
CMD ["./main"]
