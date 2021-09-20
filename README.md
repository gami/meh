# MEH
Go Web application like Twitter

# How to start development

## Setup DB

ローカルDBを構築する
```
docker compose -f docker/docker-compose.local.yaml up -d db
docker compose -f docker/docker-compose.local.yaml run meh make reset_local_db
docker compose -f docker/docker-compose.local.yaml run meh make migrate
```

## How to test
```
docker compose -f docker/docker-compose.local.yaml run meh make test
```

## How to lint

```
docker compose -f docker/docker-compose.local.yaml run meh make lint
```

## Run local app

`air`を使ってHot Reloadをおこなっています。
```
docker compose -f docker/docker-compose.local.yaml run --service-ports meh make run
```

```
# curl -X POST -H "Content-Type: application/json" -d '{"screen_name":"gami"}' localhost:8080/users | jq
{
  "ID": 1,
  "ScreenName": "gami"
}

# curl -X POST -H "Content-Type: application/json"  -d '{"screen_name":"taro"}' localhost:8080/users | jq
{
  "ID": 2,
  "ScreenName": "taro"
}

# curl -X POST -H "Content-Type: application/json" -H"Authorization: 1" -d '{"user_id":1,"followee_id":2}' localhost:8080/follows/create
# curl -X POST -H "Content-Type: application/json" -H"Authorization: 1" -d '{"user_id":1,"text":"hello"}' localhost:8080/mehs
# curl -X POST -H "Content-Type: application/json" -H"Authorization: 2" -d '{"user_id":2,"text":"hello2"}' localhost:8080/mehs

# curl -X GET -H "Content-Type: application/json" -H"Authorization: 1" localhost:8080/me/timeline | jq 
{
  "mehs": [
    {
      "id": 2,
      "text": "hello2",
      "user": {
        "id": 2,
        "screen_name": "taro"
      }
    },
    {
      "id": 1,
      "text": "hello",
      "user": {
        "id": 1,
        "screen_name": "gami"
      }
    }
  ],
  "pagination": {
    "count": 0,
    "last_id": 1
  }
}
```



## Generate codes

### Oapi-codegen

OpenAPIスキーマからServerInterfaceとレスポンス型を生成します。

```
docker compose -f docker/docker-compose.local.yaml run meh make gen-api
```

### ent

entのORMモデルを生成します。

ORMモデルの生成
```
docker compose -f docker/docker-compose.local.yaml run meh make gen-rom
```

### Project Layout
cmd/server -> controller -> usecase -> core/service <- repository

```
├── Makefile
├── README.md
├── api
│   ├── openapi ... OpenAPI生成ファイル
│   └── openapi.yaml ... OpenAPIスキーマ
├── cmd
│   └── server ... エントリポイント
├── config ... 設定
├── controller ... コントローラー/router
├── core ... ビジネスロジック/ドメインモデル
│   ├── follow
│   ├── meh
│   └── user
├── di ... DI
├── docker
├── ent ... ent（ORM）の生成ファイル
│   └── schema ... DBスキーマ
├── go.mod
├── go.sum
├── mysql ... DB接続
├── pkg ... utilility関数群
├── repository ... ORMを使って、DBアクセスする
├── tools ... マイグレーションなど
└── usecase ... coreを使って、ユースケースを実現する
```