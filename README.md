# Company-Official-App-Backend

## サービス起動方法

```
$ make up
```

## MySQL初期設定(ローカル環境以外)

以下コマンドでCloud SQL Proxyに接続しプロセスを立ち上げた後([参考](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy?hl=ja#macos-64-bit))

```
$ ./cloud_sql_proxy -instances=INSTANCE_CONNECTION_NAME=tcp:0.0.0.0:PORT
```

MySQLに接続して以下を実行

```
REVOKE 'cloudsqlsuperuser' FROM 'app'@'';
REVOKE 'cloudsqlsuperuser' FROM 'migration_ro'@'';
GRANT INSERT, SELECT, UPDATE, DELETE, INDEX ON `app-mysql-database`.* TO 'app'@'';
GRANT SELECT ON `app-mysql-database`.* TO 'migration_ro'@'';
```