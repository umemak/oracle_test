# oracle_test

## build and run

```sh
$ docker compose build
$ docker compose up -d
```

### wait db status running (healthy)

```sh
$ docker compose ps
NAME                COMMAND                  SERVICE             STATUS               PORTS
oracle_test-cli-1   "bash"                   cli                 running              
oracle_test-db-1    "/bin/sh -c 'exec $O…"   db                  running (starting)   0.0.0.0:1521->1521/tcp
```

```sh
$ docker compose ps
NAME                COMMAND                  SERVICE             STATUS              PORTS
oracle_test-cli-1   "bash"                   cli                 running             
oracle_test-db-1    "/bin/sh -c 'exec $O…"   db                  running (healthy)   0.0.0.0:1521->1521/tcp
```

## use go-ora

```sh
$ docker compose run cli go run go-ora/main.go -server "oracle://system:OraclePwd@db:1521/XE" "select 1 from dual"
1                        : 1
```

## use go-oci8

```sh
$ docker compose run cli go run go-oci8/main.go
1
```
