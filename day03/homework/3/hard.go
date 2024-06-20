// realization on framework 'chi'

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type Task struct {
	UID       uuid.UUID `json:"uid"`
	NewTask   string    `json:"newtask"`
	CreatedAt time.Time `json:"createdAt"`
}

// структура для хранения задач
type TaskMap struct {
	List map[string]Task `json:"list"`
}

// инициализация хранилища задач
var taskMap = TaskMap{List: make(map[string]Task)}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/tasks", getTasks)
	r.Post("/tasks", createTask)

	r.Put("/tasks/{id}", updateTask)
	r.Delete("/tasks/{id}", deleteTask)
	r.Get("/tasks/{id}", getTask)

	r.Get("/tasks/last", getLastTasks)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(taskMap.List); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.UID = uuid.New()

	// Проверка корректности генерации UUID
	if task.UID == uuid.Nil {
		http.Error(w, "Failed to generate UUID", http.StatusInternalServerError)
		return
	}

	idString := task.UID.String()

	taskMap.List[idString] = task

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uid := chi.URLParam(r, "id")

	task.UID, _ = uuid.Parse(uid)
	taskMap.List[uid] = task

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")

	if _, exists := taskMap.List[uid]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	delete(taskMap.List, uid)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	task := taskMap.List[uid]

	if _, exists := taskMap.List[uid]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// функция вывода последних задач
func getLastTasks(w http.ResponseWriter, r *http.Request) {

	// Определяем количество задач для вывода (по умолчанию 5)
	limitStr := r.URL.Query().Get("limit")
	limit := 5 // значение по умолчанию
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}
	// Создаем слайс для сортировки задач
	tasks := make([]Task, 0, len(taskMap.List))
	for _, task := range taskMap.List {
		tasks = append(tasks, task)
	}

	//сортировка в порядке убывания, чтобы вывести последние добавленные задачи в числе первых
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CreatedAt.After(tasks[j].CreatedAt)
	})

	// Ограничиваем список задач нужным количеством
	if len(tasks) > limit {
		tasks = tasks[:limit]
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
