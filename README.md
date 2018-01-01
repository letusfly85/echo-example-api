# Echo Example API written in Golang

## Docker Run

```sh
docker run --name example \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=example \
  -v $HOME/work/docker/mysql/example:/var/lib/mysql \
  -p 3306:3306 \
  -d mysql \
  mysqld
  --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

### Sign Up Example

```sh
curl -vv -H 'Content-Type: application/json' -c ses.txt -XPOST localhost:3030/signIn -d '{"email": "hoge@com", "password": "hogehoge"}'
```

### Sign In Example

```sh
curl -vv -H 'Content-Type: application/json' -c ses.txt -XGET localhost:3030/accounts
```

### You can Access Resource by Session stored txt file

```sh
curl -vv -H 'Content-Type: application/json' -b ses.txt -XPOST localhost:3030/signIn -d '{"email": "fuga@com", "password": "fugafuga"}'
```
