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

# opensslコマンドが実行可能か確認 → エラー発生
openssl version
#openssl: /lib/x86_64-linux-gnu/libssl.so.3:version `OPENSSL_3.4.0' not found (required by #openssl)
#openssl: /lib/x86_64-linux-gnu/libssl.so.3: #version `OPENSSL_3.2.0' not found (required by #openssl)
#openssl: /lib/x86_64-linux-gnu/libcrypto.so.3: #version `OPENSSL_3.0.9' not found (required by #openssl)
#openssl: /lib/x86_64-linux-gnu/libcrypto.so.3: #version `OPENSSL_3.3.0' not found (required by #openssl)
#openssl: /lib/x86_64-linux-gnu/libcrypto.so.3: #version `OPENSSL_3.4.0' not found (required by #openssl)
#openssl: /lib/x86_64-linux-gnu/libcrypto.so.3: #version `OPENSSL_3.2.0' not found (required by #openssl)

# libssl.so.3の場所を確認
sudo find /usr/local/ -name libssl.so.3
# /usr/local/lib64/libssl.so.3

# libcrypto.so.3
sudo find /usr/local/ -name libcrypto.so.3
# /usr/local/lib64/libcrypto.so.3

# pathを通す
vi ~/.bashrc
# 以下を記載
# export LD_LIBRARY_PATH=/usr/local/lib64
source ~/.bashrc

# インストールできたか確認
openssl version
# OpenSSL 3.4.0 22 Oct 2024 (Library: OpenSSL 3.4.0 22 Oct 2024)

# ディレクトリを削除
sudo rm -rf openssl-3.4.0/
```
