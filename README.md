# TasksAPI

RESTful API на Go для управления задачами.  
Позволяет создавать, получать, запускать и удалять задачи.

---

## Возможности

- Создание новой задачи  
- Получение всех задач  
- Получение задачи по ID  
- Запуск задачи  
- Удаление задачи  

---

## Технологии

- Go  
- Gorilla Mux  
- Logrus   

---

## Запуск проекта

1. Клонировать репозиторий:  
   ```bash
   git clone https://github.com/folada2007/TasksAPI
   ```
2. Перейти в папку с сервером:
   ```bash
   cd TasksAPI/cmd/server
   ```
3. Установить зависимости:
   ```bash
   go mod tidy
   ```
4. Запустить сервер:
   ```bash
   go run main.go
   ```
5. По умолчанию сервер будет доступен на http://localhost:8080

### API эндпоинты

- `POST /tasks/create` — создать новую задачу  
- `GET /tasks` — получить все задачи  
- `GET /tasks/{id}` — получить задачу по ID  
- `POST /tasks/{id}/start` — запустить задачу  
- `DELETE /tasks/{id}` — удалить задачу


## Примеры запросов

### Создание задачи (windows):
```bash
curl -X POST http://localhost:8080/tasks/create -H "Content-Type: application/json" -d "{\"title\":\"Заголовок\"}"
```
### Создание задачи (linux):
```bash
curl -X POST http://localhost:8080/tasks/create -H "Content-Type: application/json" -d '{"title":"Заголовок"}'
```
### Получение всех задач:
```bash
curl http://localhost:8080/tasks
```
### Получение задачи по ID:
```bash
curl http://localhost:8080/tasks/{id}
```
### Запуск задачи:
```bash
curl -X POST http://localhost:8080/tasks/1/start
```
### Удаление задачи:
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## Формат ответа

- Все ответы приходят в формате JSON.
- В случае ошибки также возвращается JSON с полем `error`.
- После запуска задачи в поле `duration` будет указано время выполнения в секундах.

### Пример ответа (GET /tasks):
```json
[
  {
    "id": "1",
    "status": "created",
    "title": "Заголовок",
    "createdAt": "2025-06-23T12:00:00Z",
    "duration": 0
  }
]
```
