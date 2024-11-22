package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10" // バリデーションでよく利用されるらしい
	"github.com/ryosuke-horie/go_todo_app/entity"
	"github.com/ryosuke-horie/go_todo_app/store"
)

type AddTask struct {
	Store     *store.TaskStore
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}

	// リクエストボディを構造体にデコードする
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// バリデーションに引っかかった場合はエラーレスポンスを返す
	err := at.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t := &entity.Task{
		Title:     b.Title,
		Status:    entity.TaskStatusTodo,
		CreatedAt: time.Now(),
	}
	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		ID int `json:"id"`
	}{ID: int(id)}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
