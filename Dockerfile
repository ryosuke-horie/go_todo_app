# デプロイ用コンテナに含めるバイナリを作成するコンテナ
# 注airのバージョンのため、golang:1.23.2-bullseyeを利用（書籍では1.18.2）
FROM golang:1.23.2-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o app

# ------------------------------------------------------------------------------

# デプロイ用のコンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ------------------------------------------------------------------------------

# ローカル環境で利用するホットリロード環境
# 注airのバージョンのため、golang:1.23.2-bullseyeを利用（書籍では1.18.2）
FROM golang:1.23.2-bullseye AS dev
WORKDIR /app

# 書籍の時と異なりair-verseに変更
RUN go install github.com/air-verse/air@latest
CMD ["air"]
