*POST to SMTP*
Сервис слушает POST запросы и отправляет содиржимое через http

**Переменные окружения**
Переменные автоматом затягиваются из файла `.env` или берутся текущие переменные окружения.

//gin
GIN_MODE=release //debug
PORT // default 8080

//smtp
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=user@example.com
SMTP_PASSWORD=your_password
SMTP_FROM=from@example.com
SMTP_TO=user1@example.com,user2@example.com[,...]
SMTP_DEFAULT_SUBJECT=Тестовая тема письма
