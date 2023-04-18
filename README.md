***POST to SMTP***

Service is listenign POST requests and transfer their body to emails via SMTP. 

**Переменные окружения**

Переменные автоматом затягиваются из файла `.env` или берутся текущие переменные окружения.

#GIN 
* GIN_MODE=release # Http server mode, default: debug
* PORT # Application listening port, default 8080

#SMTP
* SMTP_HOST=smtp.example.com
* SMTP_PORT=587
* SMTP_USER=user@example.com
* SMTP_PASSWORD=your_password
* SMTP_FROM=from@example.com
* SMTP_TO=user1@example.com,user2@example.com[,...]
* SMTP_DEFAULT_SUBJECT=Test subject
