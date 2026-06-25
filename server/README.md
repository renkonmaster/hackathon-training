# バックエンドの練習

## 起動

```sh
go -C server run ./cmd/api
```

デフォルトでは `:8080` で起動します。ポートを変えたい場合は `PORT` を指定してください。

## コード生成

OpenAPI 定義を変更したら、リポジトリルートで以下を実行します。

```sh
mise run codegen
```

## 確認

```sh
go -C server test ./...
```
