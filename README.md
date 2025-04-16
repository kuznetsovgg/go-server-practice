# Просто небольшая практика по работе с Go, Docker и прочим

Упор на структуру проекта (насколько это можно, не доводя до абсурда https://github.com/golang-standards/project-layout) и использование технологий.

## Обзор:
Приложение сейчас это простой сайт со счётчиком дней до 2026 года, прикрутил туда middleware, которая собирает user-agent пользователя при обращении к / endpoint'у. Эти данные оно также и логирует.

![image](https://github.com/user-attachments/assets/8dd63548-48c3-4c96-b3ff-de7ade5f9754)

Сделал docker-compose, поднял Promtail+Loki для логов, смотрю в Grafan'е. Пока что просто вот так:

![image](https://github.com/user-attachments/assets/6ef9f1e0-a826-4f56-a806-04c43adf16cd)

Ещё сделаю отдельный эндпоинт для сбора метрик, но попозже. Там использую Prometheus. Впрочем, я его уже настроил.
