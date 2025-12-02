# ЭФМО-01-25 Буров М.А. ПР11

# Описание проекта
Проектирование REST API (CRUD для заметок). Разработка структуры

# Требования к проекту
* Go 1.25+
* Git

# Версия Go
<img width="317" height="55" alt="image" src="https://github.com/user-attachments/assets/43f9087b-95b9-4c7d-86e9-746258c45c63" />

# Команды запуска и переменные окружения

Linux/macOS:
```
export JWT_SECRET="dev-secret"
export ACCESS_TTL="15m"
export REFRESH_TTL="168h"
export APP_PORT="8080"
go run ./cmd/server
```

Windows PowerShell:
```
$env:JWT_SECRET="dev-secret"
$env:ACCESS_TTL="15m"
$env:REFRESH_TTL="168h"
$env:APP_PORT="8080"
go run ./cmd/server
```

# Цели:
1.	Освоить принципы проектирования REST API.
2.	Научиться разрабатывать структуру проекта backend-приложения на Go.
3.	Спроектировать и реализовать CRUD-интерфейс (Create, Read, Update, Delete) для сущности «Заметка».
4.	Освоить применение слоистой архитектуры (handler → service → repository).
5.	Подготовить основу для интеграции с базой данных и JWT-аутентификацией в следующих занятиях.

# Структура проекта
Дерево структуры проекта: 
```
pz10-auth/
├── go.mod
├── go.sum
├── cmd/server/main.go
├── internal/
│   ├── core/
│   │   ├── user.go
│   │   └── service.go
│   ├── http/
│   │   ├── router.go
│   │   └── middleware/
│   │       ├── authn.go
│   │       └── authz.go
│   ├── repo/
│   │   └── user_mem.go
│   └── platform/
│       ├── config/config.go
│       └── jwt/jwt.go
└── README.md
```

# Скриншоты

Успешный /login администратора:

<img width="1380" height="638" alt="image" src="https://github.com/user-attachments/assets/8345ceae-8e0d-4cdf-9e26-4d3894701bf7" />

Успешный /login пользователя:

<img width="1369" height="647" alt="image" src="https://github.com/user-attachments/assets/144c77c6-282a-450c-8036-cfdd67d354bd" />

/me администратора:

<img width="1371" height="557" alt="image" src="https://github.com/user-attachments/assets/0380c76a-c4e3-4c28-bcef-609b6a4b7d4e" />

/admin/stats администратора:

<img width="1363" height="569" alt="image" src="https://github.com/user-attachments/assets/20e228ef-93c5-4c32-ba5c-5e5bcb7d9e75" />

403 для user на /admin/stats:

<img width="1368" height="513" alt="image" src="https://github.com/user-attachments/assets/7d944520-b2a4-43cd-840c-70bb29b9efad" />

refresh-флоу (старый/новый access):

<img width="1359" height="626" alt="image" src="https://github.com/user-attachments/assets/81343d14-83f3-4577-a291-c6a69454bfa3" />

<img width="1384" height="559" alt="image" src="https://github.com/user-attachments/assets/15c41d49-3a30-4422-9774-8c9b5ef2cad4" />

# Краткие выводы

Реализована полнофункциональная JWT-аутентификация с поддержкой двух типов токенов: short-lived access (15 мин) и long-lived refresh (7 дней). При логине система выдаёт обе токена; access используется для доступа к защищённым ресурсам, refresh позволяет обновить пару без повторного логина. Система хранит отозванные refresh-токены в in-memory blacklist, что позволяет реализовать корректный logout.

Реализована ABAC-авторизация (Attribute-Based Access Control) на примере эндпоинта /api/v1/users/{id}: пользователи с ролью user могут получить только собственный профиль (id == sub из токена), в то время как администраторы имеют полный доступ. Архитектура использует middleware-цепочку: AuthN (аутентификация через JWT) → AuthZ (авторизация через RBAC/ABAC), что разделяет ответственность и облегчает тестирование.

# Ответы на контрольные вопросы


