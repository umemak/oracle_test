FROM php:7.4-bullseye

RUN apt update -y && apt install wget unzip libaio1 -y
RUN wget https://download.oracle.com/otn_software/linux/instantclient/217000/instantclient-basic-linux.x64-21.7.0.0.0dbru.zip
RUN wget https://download.oracle.com/otn_software/linux/instantclient/217000/instantclient-sdk-linux.x64-21.7.0.0.0dbru.zip
RUN unzip instantclient-basic-linux.x64-21.7.0.0.0dbru.zip
RUN unzip instantclient-sdk-linux.x64-21.7.0.0.0dbru.zip
ENV ORACLE_HOME /instantclient_21_7
ENV LD_LIBRARY_PATH /instantclient_21_7

RUN	echo "upload_max_filesize = 128M" >> /usr/local/etc/php/conf.d/0-upload_large_dumps.ini \
&&	echo "post_max_size = 128M" >> /usr/local/etc/php/conf.d/0-upload_large_dumps.ini \
&&	echo "memory_limit = 1G" >> /usr/local/etc/php/conf.d/0-upload_large_dumps.ini \
&&	echo "max_execution_time = 600" >> /usr/local/etc/php/conf.d/0-upload_large_dumps.ini \
&&	echo "max_input_vars = 5000" >> /usr/local/etc/php/conf.d/0-upload_large_dumps.ini

STOPSIGNAL SIGINT

RUN	addgroup --system adminer \
&&	adduser --system --gecos adminer adminer \
&&	mkdir -p /var/www/html \
&&	mkdir /var/www/html/plugins-enabled \
&&	chown -R adminer:adminer /var/www/html

WORKDIR /var/www/html

RUN	set -x \
&&	apt install -y \
	libpq-dev \
	libsqlite3-dev \
	unixodbc-dev \
	freetds-dev \
    pax-utils
RUN	docker-php-ext-configure pdo_odbc --with-pdo-odbc=unixODBC,/usr
RUN	docker-php-ext-configure oci8 --with-oci8=instantclient,/instantclient_21_7
RUN	docker-php-ext-configure pdo_oci --with-pdo_oci=instantclient,/instantclient_21_7
RUN ln -s /usr/lib/x86_64-linux-gnu/libsybdb.a /usr/lib
RUN docker-php-ext-install \
	mysqli \
	pdo_pgsql \
	pdo_sqlite \
	pdo_odbc \
	pdo_dblib \
    pdo_oci \
    oci8

COPY	*.php /var/www/html/

ENV	ADMINER_VERSION 4.8.1
ENV	ADMINER_DOWNLOAD_SHA256 2fd7e6d8f987b243ab1839249551f62adce19704c47d3d0c8dd9e57ea5b9c6b3
ENV	ADMINER_COMMIT 1f173e18bdf0be29182e0d67989df56eadea4754

RUN	set -x \
&&	apt install -y git \
&&	curl -fsSL "https://github.com/vrana/adminer/releases/download/v$ADMINER_VERSION/adminer-$ADMINER_VERSION.php" -o adminer.php \
&&	echo "$ADMINER_DOWNLOAD_SHA256  adminer.php" |sha256sum -c - \
&&	git clone --recurse-submodules=designs --depth 1 --shallow-submodules --branch "v$ADMINER_VERSION" https://github.com/vrana/adminer.git /tmp/adminer \
&&	commit="$(git -C /tmp/adminer/ rev-parse HEAD)" \
&&	[ "$commit" = "$ADMINER_COMMIT" ] \
&&	cp -r /tmp/adminer/designs/ /tmp/adminer/plugins/ . \
&&	rm -rf /tmp/adminer/

COPY	entrypoint.sh /usr/local/bin/
ENTRYPOINT	[ "entrypoint.sh", "docker-php-entrypoint" ]

USER	adminer
CMD	[ "php", "-S", "0.0.0.0:8080", "-t", "/var/www/html" ]

EXPOSE 8080
