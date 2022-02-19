FROM php:8.1-cli

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN apt-get update
RUN apt-get install -y libpq-dev

# Install PostgreSQL PDO
RUN docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql \
    && docker-php-ext-install pdo pdo_pgsql pgsql

# Install Apcu
RUN pecl install apcu && docker-php-ext-enable apcu

EXPOSE 80

CMD [ "php", "-S", "0.0.0.0:80", "-t", "/usr/src/app", "/usr/src/app/public/index.php"]