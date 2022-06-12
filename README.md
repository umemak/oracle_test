# oracle_test

## build and run

```sh
$ docker compose build
$ docker compose up -d
```

### wait db status running (healthy)

```sh
$ docker compose ps
NAME                    COMMAND                  SERVICE             STATUS               PORTS
oracle_test-cli-1       "bash"                   cli                 running              
oracle_test-db-1        "/bin/sh -c 'exec $O…"   db                  running (starting)   0.0.0.0:1521->1521/tcp
oracle_test-go-oci8-1   "bash"                   go-oci8             running              
oracle_test-go-ora-1    "bash"                   go-ora              running              
```

```sh
NAME                    COMMAND                  SERVICE             STATUS              PORTS
oracle_test-cli-1       "bash"                   cli                 running             
oracle_test-db-1        "/bin/sh -c 'exec $O…"   db                  running (healthy)   0.0.0.0:1521->1521/tcp
oracle_test-go-oci8-1   "bash"                   go-oci8             running             
oracle_test-go-ora-1    "bash"                   go-ora              running             
```

## use go-ora

```sh
$ docker compose run cli go run go-ora/main.go -server "oracle://system:OraclePwd@db:1521/XE" "select 1 from dual"
1                        : 1
```

```sh
$ docker compose run go-ora bin/go-ora -server "oracle://system:OraclePwd@db:1521/XE" "select 1 from dual"
1                        : 1
```

## use go-oci8

```sh
$ docker compose run cli go run go-oci8/main.go
1
```

```sh
$ docker compose run go-oci8 bin/go-oci8
1
```

## docker image size
```sh
$ docker images
REPOSITORY                                       TAG          IMAGE ID       CREATED          SIZE
go-ora                                           local        f66e952d8ade   2 minutes ago    13.6MB
go-oci8                                          local        7b03c82b0a87   3 minutes ago    1.25GB
oraclegocli                                      latest       4ac9385215a8   16 hours ago     1.41GB
oraclelinux                                      8            3bbe8a2c4b82   10 days ago      226MB
oraclelinux                                      8-slim       1fcc1e6dda05   3 weeks ago      101MB
container-registry.oracle.com/database/express   latest       e986fd612413   2 months ago     11.2GB
ghcr.io/oracle/oraclelinux8-instantclient        21           823f111da487   2 months ago     504MB
```
