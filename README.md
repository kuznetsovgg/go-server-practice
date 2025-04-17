# Просто небольшая практика по работе с Go, Docker и прочим

Упор на структуру проекта (насколько это можно, не доводя до абсурда https://github.com/golang-standards/project-layout) и использование технологий.

## Обзор
Приложение - это простой сайт со счётчиком дней до 2026 года, прикрутил туда middleware, которая собирает user-agent пользователя при обращении к / endpoint'у. Эти данные оно также и логирует.

![image](https://github.com/user-attachments/assets/8dd63548-48c3-4c96-b3ff-de7ade5f9754)

Сделал docker-compose, поднял Promtail+Loki для логов и Prometheus (promhttp https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/promhttp), смотрю в Grafan'е. 

Сделал очень простой, не слишком информативный дашборд:

![image](https://github.com/user-attachments/assets/3d93d467-84d2-4119-89f3-1b3c26365757)

## Планы
CI ?
