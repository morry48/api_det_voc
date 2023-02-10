# set up

### docker set up

```
$ cp go/src/.env.expample go/src/.env
$ docker compose up -d
```

### db client set up

```
HOST: 127.0.0.1
PORT: 5432
USER: user
PASSWORD: password

```
### init data

```
det_voc_test_data.csv 
```

### api

```
http://localhost:3000/vocabularies/
```


# アーキテクチャ

[クリーンアーキテクチャ](https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60)っぽい何か

## ディレクトリ構成

- go
  - src
    - /config(プロダクト全般の設定ファイル)
    - /packages(各機能)
      - /〇〇(各機能名)
        - /handler(外部との接続・インターフェースの差を吸収/コントローラー)
        - /usecase(サービスを提供する機能の単位)
        - /dmain(ドメインルールの集積場)
          - /entity(何にも依存しない層)
          - /interface_repository(依存性逆転のためのリポジトリインターフェース)
        - /infra(DBや永続化層を隠蔽)
          - /postgres(ポスグレ関連)
            - /model(DBモデル・データベースのテーブルを1対1で表現・grom用)
            - /repository(データ永続化の貯蔵庫)
    - server(サーバーを起動するための設定・ルーティング)
    - tmp(フレームワークのファイル)
    - 以下ファイル(FW/ライブラリ関連)

## todo
DIを引数でインジェクションする形式にしたい
