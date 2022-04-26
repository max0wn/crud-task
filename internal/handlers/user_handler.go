package handlers

import (
	"crud-task/internal/entities"
	"crud-task/internal/requests"
	"crud-task/internal/responses"
	"crud-task/internal/services"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type UserHandler struct {
	http.ServeMux

	service *services.UserService
}

func decode(reader io.ReadCloser) (*entities.UserEntity, error) {
	var user entities.UserEntity

	if err := json.NewDecoder(reader).Decode(&user); err != nil {
		return nil, errors.New("failed to decode user")
	}

	return &user, nil
}

func encode(response *responses.UserResponse, writer http.ResponseWriter) error {
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		return errors.New("failed to encode user")
	}

	return nil
}

func (handler *UserHandler) create(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")

		user, err := decode(req.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := &responses.UserResponse{Status: "failed"}
			encode(response, w)
			return
		}

		request := &requests.UserRequest{User: user}
		err = handler.service.Create(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := &responses.UserResponse{Status: "failed"}
			encode(response, w)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := &responses.UserResponse{Status: "OK"}
		encode(response, w)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (handler *UserHandler) update(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPatch {
		w.Header().Set("Content-Type", "application/json")

		user, err := decode(req.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := &responses.UserResponse{Status: "failed"}
			encode(response, w)
			return
		}

		request := &requests.UserRequest{User: user}
		err = handler.service.Update(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := &responses.UserResponse{Status: "failed"}
			encode(response, w)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := &responses.UserResponse{Status: "OK"}
		encode(response, w)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (handler *UserHandler) delete(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")

		user, err := decode(req.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := &responses.UserResponse{Status: "failed"}
			encode(response, w)
			return
		}

		request := &requests.UserRequest{User: user}
		err = handler.service.Delete(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := &responses.UserResponse{Status: "failed"}
			encode(response, w)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := &responses.UserResponse{Status: "OK"}
		encode(response, w)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (handler *UserHandler) hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func NewUserHandler() *UserHandler {
	handler := &UserHandler{
		service: services.NewUserService(),
	}

	handler.HandleFunc("/api/v1/hello", handler.hello)
	handler.HandleFunc("/api/v1/create", handler.create)
	handler.HandleFunc("/api/v1/update", handler.update)
	handler.HandleFunc("/api/v1/delete", handler.delete)

	return handler
}
