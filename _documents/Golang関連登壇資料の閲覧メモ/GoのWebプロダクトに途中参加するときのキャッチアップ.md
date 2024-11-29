# [GoのWebプロダクトに途中参加するときのキャッチアップ](https://speakerdeck.com/budougumi0617/how-to-catch-up-go-web-product)

1. 初期設定に関して
     -  README.mdで環境変数などの情報がないか確認
     -  makefileを確認しておく
2.  依存関係の確認
       -   go modを見て依存pkgの確認を行う
3. 起動してローカルで操作してみる
   - APIの実行
   - UIを操作
   - アクセスログやエラーログを確認する
4. ディレクトリ構成を理解する
     - GolangにはデファクトFWがないため組織に合わせてキャッチアップする必要がある。
     - よく見るディレクトリ名
       -  `handler`, `api`
          -  `middleware`
          -  `interceptor`
       -  `usecase`
       -  `service`
       -  `domain`
       -  `repository`
       -  `gateway`
       -  `infrastructure`
  5. ソースコードを読む
      - GitHub Copilot Chatなども活用する  
  6. 次のNew Joinerのためにキャッチアップ結果をアウトプットしておく
   - ドキュメントの古くなっているところは修正する
   - 人に聞かないとわからなかったことはドキュメントにする