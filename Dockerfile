FROM ghcr.io/oracle/oraclelinux8-instantclient:21

RUN yum install -y wget tar

RUN wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz

RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz

ENV PATH="${PATH}:/usr/local/go/bin"

CMD [ "go", "version" ]
