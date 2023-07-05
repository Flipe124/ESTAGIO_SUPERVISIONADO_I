FROM php:fpm-alpine

RUN apk update && apk upgrade; apk add gcc
RUN docker-php-ext-install mysqli && docker-php-ext-enable mysqli
