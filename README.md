# Auth API for gRPC course project

[![build-and-test](https://github.com/satanaroom/auth/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/satanaroom/auth/actions/workflows/build-and-test.yml)
[![golangci-lint](https://github.com/satanaroom/auth/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/satanaroom/auth/actions/workflows/golangci-lint.yml)

Auth API - это микросервис, который предоставляет API для создания, аутентификации и авторизации пользователей.

## Quick start
Для работы необходимо установить [Docker](https://docs.docker.com/engine/install/) и [Docker Compose](https://docs.docker.com/compose/install/).

1. Склонировать репозиторий
``` bash
git clone https://github.com/satanaroom/auth.git
```

2. Запустить контейнеры PostgreSQL и Prometheus
``` bash
docker-compose up -d
```

3. Установить зависимости и сбилдить проект
``` bash
make build
```

Описание API можно найти в вики проекта - https://github.com/satanaroom/auth/wiki/%D0%9E%D0%BF%D0%B8%D1%81%D0%B0%D0%BD%D0%B8%D0%B5-%D1%80%D1%83%D1%87%D0%B5%D0%BA-%D0%BC%D0%B8%D0%BA%D1%80%D0%BE%D1%81%D0%B5%D1%80%D0%B2%D0%B8%D1%81%D0%B0
