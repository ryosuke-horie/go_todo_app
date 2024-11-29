# [google/wireを使ったGoらしいアーキテクチャへの取り組み](https://speakerdeck.com/budougumi0617/gocon-fukuoka-2019-summer)

- google/wireはDIを行うためのCLIツールと実装用のAPIライブラリ
- package構成
  - entity >>> データ定義
  - domain >>>  ドメインロジック実装
  - repository >>> データの永続化
    - RDBMSに対するCRUD
    - UTにsqlmockなどを利用する
  - usecase >>>  ドメインモデルとリポジトリからシナリオを作る
    - リポジトリ層に対する抽象化
    - ドメインとリポジトリを組み合わせる層
    - UTではgo-mockを利用
  - http >>> Httpとプリミティブなパラメータを仲介
  - app >>> Injectorを配置しwire自動生成コードを置く