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