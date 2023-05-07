# これはaws_lambdaを使用して、簡易的なAPIサーバーを作成するためのものになります

## 使用言語
Golang

## フレームワーク
gin

## インフラ
aws s3
aws aurora


## 開発環境

dockerイメージを作成する時のコマンド
`docker image build --build-arg MYSQL_DATABASE=$MYSQL_DATABASE MYSQL_USER=$MYSQL_USER　DB_PASSWORD=$DB_PASSWORD　DB_ENDPOINT=$DB_ENDPOINT　ENV=$ENV -t local-test:latest .`

docker composeのサービスを作成する時のコマンド
`docker compose build --build-arg　ENV=$ENV SERVER_PORT=$SERVER_PORT web `

`docker compose up -d`


## 環境変数
PORT -> APIサーバーのポート番号
ENV　-> 開発環境か本番環境かを判定する
DB_USERNAME
DB_PASSWORD
DB_ENDPOINT
DB_DATABASE
