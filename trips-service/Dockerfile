FROM composer:2 AS composer_builder
WORKDIR /app

FROM php:8.3-fpm-alpine as runtime

COPY composer.json composer.lock ./

RUN composer install --prefer-dist --no-progress --no-interaction

RUN apk add --no-cache \
    icu-dev libzip-dev oniguruma-dev \
    zlib-dev libpng-dev libjpeg-turbo-dev libxml2-dev \
    bash git curl tzdata

RUN docker-php-ext-configure gd --with-jpeg \
 && docker-php-ext-install -j$(nproc) \
      intl \
      pdo_mysql \
      opcache \
      zip \
      mbstring \
      gd \
      xml

WORKDIR /var/www/html

COPY --from=composer_builder /app/vendor ./vendor
COPY . .

EXPOSE 9000

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD php -v || exit 1

CMD ["php-fmp"]

