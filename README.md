# API
APIのテンプレート

## 使い方

### Makefile版
```
make up
```

### docker-compose版
```
go get github.com/cespare/reflex
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.3.8
mkdir -p internal/http/echo/gen/
oapi-codegen -package gen -o internal/http/echo/gen/gen.go -generate "types,server" resource/openapi/openapi.yaml
docker-compose up
```

## 構成

### domain
インターフェイスを格納する場所。  
関数群を定義してある

### 機能ディレクトリ
テンプレートでは`user`という機能として作りましたが、このディレクトリ配下に
- delivery
- usecase
- reposiroty

が配置されてます。

#### delivery
ユーザからアクセスされる場所。  
例えばhttpサーバだったりのルーティングを書く

#### usecase
処理の中身を書く場所。  
実処理はここに書いて、deliveryには書かないようにすると疎結合になる

#### repository
データベース処理などを書く場所。  
データベース処理などもdomainに書かれたインターフェイスを基に関数の実部を作りあげ、返却するようにする。

### 使い方
ユーザからアクセスされる`delivery`を作り、それに必要な`usecase`を指定してあげる。`usecase`を作成するときも必要な`repository`を指定することで、お互いが疎な関係になるので付け替えが容易になる。

## テスト方法

### Makefile版
```
make test
```

### docker-compose版
```
go get github.com/cespare/reflex
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.3.8
mkdir -p internal/http/echo/gen/
oapi-codegen -package gen -o internal/http/echo/gen/gen.go -generate "types,server" resource/openapi/openapi.yaml
go test -v ./...
```

