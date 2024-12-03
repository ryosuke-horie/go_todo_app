# opensslコマンドのインストール

brewを利用しない環境のためメモ

## 手順(Ubuntu)

```bash
# バイナリを取得
wget https://github.com/openssl/openssl/releases/download/openssl-3.4.0/openssl-3.4.0.tar.gz

# 解凍
tar -xzvf openssl-3.4.0.tar.gz

# 解凍したディレクトリへ移動
cd openssl-3.4.0

# 必要な依存関係をインストール
sudo apt install build-essential

# ソースコードからインストール
sudo ./config
sudo make
sudo make install
```