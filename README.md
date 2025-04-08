# Telegram Bot
Здесь представлены файлы разработки из основного репозитория

## Разработан телеграмм бот - текстовый помошник в беседе
- Внутри него заложена нейростеть Gemini для генерации ответов пользователей.
- Реализована поддержка конекста беседы с участниками беседы.
- Реализована функция очистки контекста бота.
- Бот помнит контекс в течении определенного времени.
- Ассинхронная работа бота.

## Используемые технологии для разроботки проекта:
- Разработка приложения выполнена на Go.
- Использован подход чистой архитектуры для построения приложения, также использована техника внедрения зависимостей.
- Для связи тг бота используется библиотека [telebot.v4](https://gopkg.in/telebot.v4),
- Для связи с gemini используется [github.com/cohesion-org/deepseek-go](https://github.com/cohesion-org/deepseek-go).
- Для работы с контекстом беседы используется Redis.

## Дальнейшие планы:
- Внедрение очередей сообщений (обработка в gemini занимает разное количество времени, поэтому ответы могут приходить в неправильнос порядке).
- Использование Kubernetus.
- Внедрения нейросети для генерации картинок.
- Использования multi-stage сборки.