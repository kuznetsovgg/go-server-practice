# Просто небольшая практика по работе с Go, Docker и прочим

Упор на структуру проекта (насколько это можно, не доводя до абсурда https://github.com/golang-standards/project-layout) и использование технологий.

## Обзор
Приложение - это простой сайт со счётчиком дней до 2026 года, прикрутил туда middleware, которая собирает user-agent пользователя при обращении к / endpoint'у. Эти данные оно также и логирует.

![image](https://github.com/user-attachments/assets/8dd63548-48c3-4c96-b3ff-de7ade5f9754)

Сделал docker-compose, поднял Promtail+Loki для логов и Prometheus (promhttp https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/promhttp), смотрю в Grafan'е. 

Сделал очень простой дашборд, настроил под логи и пару метрик (приложения):

![image](https://github.com/user-attachments/assets/3d93d467-84d2-4119-89f3-1b3c26365757)

Также написал CI-workflow, состоящий из 3 шагов: проверки go.mod (практика работы с зависимостями приложения), линтинга кода, а также сборки и пуша образа в docker hub.

![image](https://github.com/user-attachments/assets/d381ba86-f1ab-4d4e-ae88-a65a780473ff)

![image](https://github.com/user-attachments/assets/9ab5c3f6-1802-44b4-a295-929a76c15dcd)

## Итоги:
Поработал со всем, чем хотел: Prometheus, Loki, Grafana, Github Actions; также впервые поработал с promhttp и Promtail.
