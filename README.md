# ЭФМО-01-25 Буров М.А. ПР11

# Описание проекта
Проектирование REST API (CRUD для заметок). Разработка структуры

# Требования к проекту
* Go 1.25+
* Git

# Версия Go
<img width="317" height="55" alt="image" src="https://github.com/user-attachments/assets/43f9087b-95b9-4c7d-86e9-746258c45c63" />

# Цели:
1.	Освоить принципы проектирования REST API.
2.	Научиться разрабатывать структуру проекта backend-приложения на Go.
3.	Спроектировать и реализовать CRUD-интерфейс (Create, Read, Update, Delete) для сущности «Заметка».
4.	Освоить применение слоистой архитектуры (handler → service → repository).
5.	Подготовить основу для интеграции с базой данных и JWT-аутентификацией в следующих занятиях.

# Структура проекта
Дерево структуры проекта: 
```
pz11-notes-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── core/
│   │   ├── note.go
│   │   └── service/
│   │       └── note_service.go
│   ├── http/
│   │   ├── router.go
│   │   └── handlers/
│   │       └── notes.go
│   └── repo/
│       └── note_mem.go
├── api/
│   └── openapi.yaml
├── go.mod
└── go.sum
```

# Теоретические положения REST API и CRUD

2.1 REST API - Понятие и принципы
REST (Representational State Transfer) - это архитектурный стиль для проектирования распределённых систем, основанный на стандартном протоколе HTTP.

Основные принципы REST:
| Принцип | Описание | Пример |
|---------|---------|--------|
| **Ресурсность** | Все данные представлены как ресурсы с уникальными URI | `/api/v1/notes/{id}` |
| **Методы HTTP** | Использование стандартных HTTP-методов для операций | GET, POST, PATCH, DELETE |
| **Stateless** | Каждый запрос содержит полную информацию для обработки | Не требует сессий |
| **JSON формат** | Обмен данными в формате JSON | `{"id": 1, "title": "Note"}` |
| **Единообразие** | Одинаковая структура запросов и ответов | Предсказуемое API |


CRUD - Create, Read, Update, Delete - четыре базовые операции для работы с данными.
| CRUD | HTTP | Эндпоинт | Статус | Описание |
|------|------|----------|--------|---------|
| **Create** | POST | `/api/v1/notes` | 201 | Создание новой заметки |
| **Read (одна)** | GET | `/api/v1/notes/{id}` | 200 | Получение одной заметки |
| **Read (все)** | GET | `/api/v1/notes` | 200 | Получение всех заметок |
| **Update** | PATCH | `/api/v1/notes/{id}` | 200 | Обновление заметки |
| **Delete** | DELETE | `/api/v1/notes/{id}` | 204 | Удаление заметки |

# Скриншоты

Создание заметки:
<img width="1379" height="617" alt="image" src="https://github.com/user-attachments/assets/99e50838-dde9-42f8-9494-194746af9946" />

# Примеры кода

cmd/api/main.go - Точка входа
```
package main

import (
  "log"
  "net/http"
  "example.com/pz11-notes-api/internal/http"
  "example.com/pz11-notes-api/internal/http/handlers"
  "example.com/pz11-notes-api/internal/repo"
)

func main() {
  repo := repo.NewNoteRepoMem()
  h := &handlers.Handler{Repo: repo}
  r := httpx.NewRouter(h)

  log.Println("Server started at :8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}
```

internal/core/note.go - Модель данных
```
package core

import "time"

type Note struct {
  ID        int64
  Title     string
  Content   string
  CreatedAt time.Time
  UpdatedAt *time.Time
}
```

internal/http/handlers/notes.go - HTTP-обработчики
```
package handlers

import (
  "encoding/json"
  "net/http"
  "example.com/pz11-notes-api/internal/core"
  "example.com/pz11-notes-api/internal/repo"
)

type Handler struct {
  Repo *repo.NoteRepoMem
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
  var n core.Note
  if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
    http.Error(w, "Invalid input", http.StatusBadRequest)
    return
  }
  id, _ := h.Repo.Create(n)
  n.ID = id
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(n)
}
```

# Краткие выводы
Успешно разработан REST API для управления заметками на языке Go с реализацией CRUD-операций и применением слоистой архитектуры (обработчики → сервис → репозиторий).

# Ответы на контрольные вопросы

1.	Что означает аббревиатура REST и в чём её суть?

REST (Representational State Transfer) - это архитектурный стиль для построения распределённых веб-приложений.

Суть REST:

- Использование стандартных HTTP-методов (GET, POST, PATCH, DELETE)

- Представление данных как ресурсов с уникальными URI

- Отсутствие состояния (stateless) - каждый запрос независим

- Единообразный интерфейс для всех операций

- Использование стандартных HTTP-кодов ответов

Пример: /api/v1/notes/1 - это ресурс (заметка с ID=1), с которым работают через стандартные HTTP-методы.

2.	Как связаны CRUD-операции и методы HTTP?

| CRUD | HTTP | Описание |
|------|------|---------|
| **Create** | POST | Создание нового ресурса |
| **Read** | GET | Получение ресурса |
| **Update** | PATCH/PUT | Изменение ресурса |
| **Delete** | DELETE | Удаление ресурса |

3.	Для чего нужна слоистая архитектура (handler → service → repository)?

Слоистая архитектура разделяет ответственность:

- HTTP Handlers - парсинг запросов, валидация, форматирование ответов

- Service (Core) - бизнес-логика, обработка данных

- Repository - работа с данными (в памяти, БД, файлы)

Преимущества:

- Разделение ответственности - каждый слой делает одно

- Тестируемость - можно тестировать каждый слой отдельно

- Реиспользуемость - service можно использовать в разных контекстах

- Масштабируемость - легко добавлять новые функции

- Гибкость - можно менять реализацию (например, in-memory → БД)

4.	Что означает принцип «stateless» в REST API?

Stateless означает, что сервер НЕ хранит состояние клиента между запросами.

Преимущества:

- Масштабируемость - можно распределить на несколько серверов

- Надёжность - отказ одного сервера не теряет состояние

- Производительность - не нужно искать сессию

5.	Почему важно использовать стандартные коды ответов HTTP?

Стандартные HTTP-коды - это язык, на котором общаются клиент и сервер.

| Код | Статус | Значение |
|-----|--------|----------|
| 200 | OK | Успешное выполнение |
| 201 | Created | Ресурс создан |
| 204 | No Content | Операция выполнена, контент не возвращается |
| 400 | Bad Request | Ошибка в запросе |
| 404 | Not Found | Ресурс не найден |
| 500 | Internal | Ошибка сервера |

Почему это важно:

- Предсказуемость - клиент знает, как обработать ответ

- Совместимость - работает с любыми HTTP-библиотеками

- Автоматизация - фреймворки и инструменты понимают коды

- Debugging - быстро понять, в чём проблема

6.	Как можно добавить аутентификацию в REST API?

Основные методы:

1. JWT (JSON Web Token) - наиболее распространённый:
```
curl -H "Authorization: Bearer eyJhbGc..." http://localhost:8080/api/v1/notes
```

2. API Key:
```
curl -H "X-API-Key: abc123xyz" http://localhost:8080/api/v1/notes
```

3. OAuth 2.0 - для социальной аутентификации

В коде:
```
func (h *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	// Проверка токена в заголовке
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// Верификация токена...
	// Если валиден, обработать запрос
}
```

7.	В чём преимущество версионирования API (например, /api/v1/)?

Версионирование позволяет развивать API без нарушения совместимости.

```
/api/v1/notes  - Старая версия (для старых клиентов)
/api/v2/notes  - Новая версия (с улучшениями)
```
