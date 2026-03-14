package handlers

import (
	"assignment4/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type PaginationHandler struct {
	usecase *usecase.UserUsecase
}

func NewPaginationHandler(uc *usecase.UserUsecase) *PaginationHandler {
	return &PaginationHandler{usecase: uc}
}

// GetPaginatedUsers - GET /users/paginated?page=1&pageSize=10&name=alice&orderBy=name
func (h *PaginationHandler) GetPaginatedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Получаем параметры пагинации
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize < 1 {
		pageSize = 10
	}
	
	// Фильтры
	filters := make(map[string]interface{})
	if name := r.URL.Query().Get("name"); name != "" {
		filters["name"] = name
	}
	if email := r.URL.Query().Get("email"); email != "" {
		filters["email"] = email
	}
	if gender := r.URL.Query().Get("gender"); gender != "" {
		filters["gender"] = gender
	}
	
	// Сортировка
	orderBy := r.URL.Query().Get("orderBy")
	
	// Получаем результат
	result, err := h.usecase.GetPaginatedUsers(page, pageSize, filters, orderBy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetCommonFriends - GET /users/common-friends?user1=2&user2=3
func (h *PaginationHandler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	user1Str := r.URL.Query().Get("user1")
	user2Str := r.URL.Query().Get("user2")
	
	if user1Str == "" || user2Str == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user1 and user2 parameters required"})
		return
	}
	
	user1, err1 := strconv.Atoi(user1Str)
	user2, err2 := strconv.Atoi(user2Str)
	
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid user IDs"})
		return
	}
	
	friends, err := h.usecase.GetCommonFriends(user1, user2)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friends)
}

// AddFriend - POST /users/add-friend {"user_id": 2, "friend_id": 3}
func (h *PaginationHandler) AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var req struct {
		UserID   int `json:"user_id"`
		FriendID int `json:"friend_id"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}
	
	if req.UserID == 0 || req.FriendID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user_id and friend_id are required"})
		return
	}
	
	err := h.usecase.AddFriend(req.UserID, req.FriendID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "friend added successfully"})
}