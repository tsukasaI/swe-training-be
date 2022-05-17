# auWellness_forBiz_api

## 概要

xxx

## 環境

- Go v1.17-bullseye
- gin v1.7.4
- gqlgen v0.16.0
- MySQL 8.0

## 開発環境

リモートコンテナを用いたローカル開発環境を構築する手順を記載する

### 初期設定

#### 事前準備

- Docker インストール
- VSCode インストール
- VSCode の拡張機能「Remote - Containers」インストール
  - <https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers>
- 本リポジトリの clone
  - Windows 環境の場合、WSL を使用する場合でも WSL 上には clone しない事。
- MySQLクライアント (任意)
  - DBeaver
  - A5:SQL Mk-2 (Windowsのみ)
- APIテスト用クライアント (任意)
  - Postman
  - curl

#### 環境変数設定

1. /.devcontaner 直下に.env を作成する
1. 下記の環境変数を記載

   ```bash
   TIME_ZONE=Asia/Tokyo
   LOCALE=ja_JP.UTF-8

   GO_VERSION=1.17-bullseye
   WORKING_DIR=/forBiz-cms-api

   DB_DATABASE=test
   DB_USERNAME=docker
   DB_PASSWORD=docker
   ```

#### SSH 設定

リモートコンテナからホストマシンの SSH キーにアクセスできるようにする

1. (Win, Linux) ssh-agent を起動する
1. 下記コマンドで ssh-agent に private key を追懐する

   ```bash
   # github_rsaにはgithubにアクセス可能なssh private keyファイル名を指定
   ssh-add $HOME/.ssh/github_rsa
   ```

### Romote-Container 作成

1. VSCode 起動
1. 左下のアイコンクリック
1. 「Remote-Containers - Open Folder in Container...」クリック
1. しばらく待つ
   - 初回の場合コンテナ image の取得や作成が行われるため、それなりに時間がかかります
1. `Cmd + j`でターミナルを起動語下記を入力

   ```bash
   go mod tidy
   ```

1. 開発可能

#### 補足

- DB 用コンテナの起動のためにはホストマシンのポート番号 3306 が未使用である必要がある

### API サーバ起動

#### 通常起動

- 以下をコンソール画面で実行
- `Cmd + d`で終了

```bash
npm run dev
```

- テーブルが作成されない場合は、`/config/config*.yml`の中に定義されている`isInitTable`がfalseになっていないか確認する。

  ```yml
  server:
    isInitTable: true # ここがfalseになっているとテーブルが作成されない。
  ```

#### デバッグ起動

ブレイクポイントやスタックトレースが利用可能

- F5 キー押下
- デバッグコンソール表示（ `Ctrl + Shift + Y` or `Command + Shift + Y`）
  - 最下部の入力欄でブレイクした箇所でのコード実行が可能です

### Unit Test 実行

- 下記コマンドを実行
  - `Cmd + ; -> a`

### テスト実行方法

1. 下記コマンドを入力して DB にサンプルデータを投入する

   ```bash
   mysql --host=host.docker.internal -P 3306 --user=docker --password=docker for_biz_test < resource/for_biz_test.sql
   ```
2. 下記コマンドでテストを実行する

   ```bash
   go test ./apps/cms/api/
   ```

### config.yml

```
/config/
   config.yml ... ビルド時に実際に読み込みされるファイル。ローカルでのビルド時はこちらを編集。
   config_dev.yml
   config_stg.yml ... 各環境用のconfigファイル。こちらに設定しておくことで環境へのデプロイ時にgithub actionsがconfig.ymlにリネームしてdocker buildされる。
   config_prd.yml
```

### Database Initialize

1. `Cmd + j`でターミナルを起動
1. (必要に応じて) 下記コマンドを入力し、DB にサンプルデータを投入する

   ```bash
   mysql --host=host.docker.internal -P 3306 --user=docker --password=docker test < resource/initdb-sample.sql
   ```
