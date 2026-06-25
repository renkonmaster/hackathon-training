# ハッカソン前準備

## はじめに
このリポジトリは、[Webなろう講習会](https://traptitech.github.io/naro-text/) 第一部を終えた人がハッカソンを始め共同開発に入る前に練習するための簡易テンプレートリポジトリです。

基本的な使い方は、少人数のチームを組みここにあるIssueを解決しながら進めていくのがよいと思います。

## リポジトリの構成


## 環境構築

- miseをインストールする
- `mise run setup` を実行



## Docker Compose での開発

MySQL と Adminer だけを起動する場合:

```sh
docker compose up
```

バックエンドを単体で開発する場合:

```sh
docker compose --profile server up --build
```

実フロントエンドも使いたい場合は、別途 `client` を起動するか fullstack profile を使ってください。

フロントエンドを OpenAPI の仮サーバー付きで開発する場合:

```sh
docker compose --profile client up --build
```

フロントエンドとバックエンドを両方起動する場合:

```sh
docker compose --profile fullstack up --build
```

- API: http://localhost:8080
- フロントエンド: http://localhost:5173
- 仮 API サーバー: http://localhost:4010
- Adminer: http://localhost:8081

## Let's Start
それでは早速Issue #1 を見てそこにある手順に従って進めていきましょう！
