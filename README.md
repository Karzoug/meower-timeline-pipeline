# Meower timeline pipeline

Cервис, отслеживающий события генерируемые сервисами [пользователей](https://github.com/Karzoug/meower-user-outbox), [сообщений](https://github.com/Karzoug/meower-post-outbox) и [связей](https://github.com/Karzoug/meower-relation-outbox), обрабатывающий их по сценариям работы ленты и отправляющий события-задачи для [сервиса ленты](https://github.com/Karzoug/meower-timeline-service).

### Стек
- Основной язык: go
- Брокер: kafka
- Наблюдаемость: opentelemetry, jaeger, prometheus
- Контейнеры: docker, docker compose