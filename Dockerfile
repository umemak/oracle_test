FROM ghcr.io/oracle/oraclelinux8-instantclient:21

RUN yum install -y wget tar pkgconfig gcc

RUN wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz

RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz

ENV PATH "${PATH}:/usr/local/go/bin"
ENV PKG_CONFIG_PATH "/usr/share/pkgconfig/"

COPY ./go-oci8/oci8.pc /usr/share/pkgconfig/
WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY go-oci8 ./go-oci8/
COPY go-ora ./go-ora/

CMD [ "go", "version" ]
