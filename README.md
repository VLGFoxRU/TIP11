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
<img width="1364" height="580" alt="image" src="https://github.com/user-attachments/assets/3bc0b686-aa57-45e2-ae4a-9de6e05c54c6" />

Получение всех заметок:
<img width="1356" height="614" alt="image" src="https://github.com/user-attachments/assets/3870ecfb-91b5-4067-9e7b-3823c1aff687" />

Получение заметки по ID:
<img width="1357" height="594" alt="image" src="https://github.com/user-attachments/assets/128ea21a-2dd3-4b96-8b18-242827026ea5" />

Обновление заметки:
<img width="1361" height="608" alt="image" src="https://github.com/user-attachments/assets/f4b860da-6473-47f8-849a-8482c8164400" />

Удаление заметки:
<img width="1371" height="475" alt="image" src="https://github.com/user-attachments/assets/5faa66c2-4ece-45ec-8b63-70d8b8097b78" />

# Примеры кода

cmd/api/main.go - Точка входа
```
package main

import (
	"log"
	"net/http"

	"example.com/pz11-notes-api/internal/http"
	"example.com/pz11-notes-api/internal/repo"
)

func main() {
	// Инициализация репозитория
	noteRepo := repo.NewNoteRepoMem()

	// Создание маршрутизатора
	r := router.NewRouter(noteRepo)

	// Запуск сервера
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
```

internal/core/note.go - Модель данных
```
package core

import "time"

type Note struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
```

internal/http/handlers/notes.go - HTTP-обработчики
```
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"example.com/pz11-notes-api/internal/core"
	"example.com/pz11-notes-api/internal/repo"
)

type NoteHandler struct {
	repo *repo.NoteRepoMem
}

func NewNoteHandler(repo *repo.NoteRepoMem) *NoteHandler {
	return &NoteHandler{repo: repo}
}

// CreateNote - POST /api/v1/notes
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req core.CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	note := core.Note{
		Title:   req.Title,
		Content: req.Content,
	}

	id, err := h.repo.Create(note)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	note.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// GetNote - GET /api/v1/notes/{id}
func (h *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// GetAllNotes - GET /api/v1/notes
func (h *NoteHandler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes := h.repo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// UpdateNote - PATCH /api/v1/notes/{id}
func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req core.UpdateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	note, err := h.repo.Update(id, req)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// DeleteNote - DELETE /api/v1/notes/{id}
func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
