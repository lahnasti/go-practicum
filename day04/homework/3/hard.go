/*
Расширьте API, добавив аутентификацию с использованием токена JWT (JSON Web Token). Создайте маршруты:
POST /login: Вход пользователя, где можно отправить имя пользователя и пароль для получения токена.
GET /profile: Получить информацию о текущем пользователе, требующую аутентификации с использованием токена.
Защитите маршруты для управления задачами и пользователями с использованием Middleware для проверки наличия и валидности токена JWT.
*/