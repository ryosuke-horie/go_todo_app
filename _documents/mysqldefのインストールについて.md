# mysqldefのインストールについて

書籍で紹介されている方法と異なるため記載する。

## 違う点

- GitHubのリポジトリ名が変わっている
  -  `k0kubun/sqldef`が`sqldef/sqldef`になっている
  -  go getで取得してもうまくいかなかった

## 対応

バイナリをGitHubから取得してアーカイブを解凍する方法を利用した。

```bash
# 実行時の最新版を指定してインストール
wget https://github.com/sqldef/sqldef/releases/download/v0.17.24/mysqldef_linux_amd64.tar.gz

# 解凍
tar -xzvf mysqldef_linux_amd64.tar.gz

# バイナリの配置と実行権限の設定
sudo mv mysqldef /usr/local/bin/
sudo chmod +x /usr/local/bin/mysqldef

# アーカイブを削除
rm mysqldef_linux_amd64.tar.gz
```