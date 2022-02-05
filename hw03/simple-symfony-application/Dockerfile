FROM php:8.1-cli

COPY . /usr/src/app
WORKDIR /usr/src/app

# Install Postgre PDO
RUN apt-get update
RUN apt-get install -y libpq-dev \
    && docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql \
    && docker-php-ext-install pdo pdo_pgsql pgsql

EXPOSE 80

CMD [ "php", "-S", "0.0.0.0:80", "-t", "/usr/src/app", "/usr/src/app/public/index.php"]