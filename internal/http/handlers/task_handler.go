package handlers

import (
	"LongTaskAPI/internal/services"
	"LongTaskAPI/internal/services/factory"
	"LongTaskAPI/internal/utils"
	"LongTaskAPI/pkg/dto"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	logger  *logrus.Logger
	service *services.TaskService
}

func CreateHandler(logger *logrus.Logger, service *services.TaskService) *Handler {
	return &Handler{
		logger:  logger,
		service: service}
}

func (h *Handler) CreateNewTasksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var taskReq dto.TaskRequestDto

		err := utils.DecodeJSONBody(r, &taskReq)
		if err != nil {
			h.logger.WithError(err).Error("Error in reading request body")
			if err = utils.RespondWithErrors(w, http.StatusBadRequest, "Invalid JSON"); err != nil {
				h.logger.WithError(err).Error("Failed to respond with 400")
			}
			return
		}
		task := factory.ToTask(taskReq)

		task, err = h.service.Create(task)
		if err != nil {
			if err = utils.RespondWithErrors(w, http.StatusInternalServerError, "Server error"); err != nil {
				h.logger.WithError(err).Error("Failed to respond with 500")
			}
			return
		}

		duration := 0.0
		if !task.EndAt.IsZero() {
			duration = task.EndAt.Sub(task.CreatedAt).Seconds()
		}

		response := dto.TaskResponse{
			Id:        task.ID,
			Status:    task.Status,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			Duration:  duration,
		}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			h.logger.WithError(err).Error("Failed to respond with 200")
		}
	}
}

func (h *Handler) GetAllTasksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		AllTasks, err := h.service.GetAll()
		if err != nil {
			if err = utils.RespondWithErrors(w, http.StatusInternalServerError, "Server error"); err != nil {
				h.logger.WithError(err).Error("Failed to respond with 500")
			}
			return
		}

		var response []dto.TaskResponse

		for _, task := range AllTasks {
			duration := 0.0
			if !task.EndAt.IsZero() {
				duration = task.EndAt.Sub(task.CreatedAt).Seconds()
			}

			response = append(response, dto.TaskResponse{
				Id:        task.ID,
				Status:    task.Status,
				Title:     task.Title,
				CreatedAt: task.CreatedAt,
				Duration:  duration,
			})
		}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			h.logger.WithError(err).Error("Failed to respond with 200")
		}
	}
}

func (h *Handler) GetTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, ok := h.getIDFromVars(w, r)
		if !ok {
			return
		}

		task, err := h.service.GetById(int64(id))
		if err != nil {
			if errors.Is(err, utils.ErrorNotFound) {
				h.logger.WithFields(logrus.Fields{
					"error": err,
				}).Error("not found task by id")

				if err = utils.RespondWithErrors(w, http.StatusNotFound, "not found"); err != nil {
					h.logger.WithError(err).Error("failed to respond with 404")
				}
			} else {
				h.logger.WithFields(logrus.Fields{
					"error": err,
				}).Error("server error")
				if err = utils.RespondWithErrors(w, http.StatusInternalServerError, "server error"); err != nil {
					h.logger.WithError(err).Error("failed to respond with 500")
				}
			}
			return
		}
		duration := 0.0
		if !task.EndAt.IsZero() {
			duration = task.EndAt.Sub(task.CreatedAt).Seconds()
		}

		response := dto.TaskResponse{
			Id:        task.ID,
			Status:    task.Status,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			Duration:  duration,
		}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			h.logger.WithError(err).Error("Failed to respond with 200")
		}
	}
}

func (h *Handler) DeleteTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, ok := h.getIDFromVars(w, r)
		if !ok {
			return
		}

		err := h.service.DeleteTask(int64(id))
		if err != nil {
			if err = utils.RespondWithErrors(w, http.StatusInternalServerError, "Server error"); err != nil {
				h.logger.WithError(err).Error("Failed to respond with 500")
			}
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *Handler) StartTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, ok := h.getIDFromVars(w, r)
		if !ok {
			return
		}

		task, err := h.service.GetById(int64(id))
		if err != nil {
			if errors.Is(err, utils.ErrorNotFound) {
				h.logger.WithFields(logrus.Fields{
					"error": err,
				}).Error("not found task by id")

				if err = utils.RespondWithErrors(w, http.StatusNotFound, "not found"); err != nil {
					h.logger.WithError(err).Error("failed to respond with 404")
				}
			} else {
				h.logger.WithFields(logrus.Fields{
					"error": err,
				}).Error("server error")
				if err = utils.RespondWithErrors(w, http.StatusInternalServerError, "server error"); err != nil {
					h.logger.WithError(err).Error("failed to respond with 500")
				}
			}
			return
		}

		task, err = h.service.StartTask(task)
		if err != nil {
			h.logger.WithError(err).Error("Failed to start task")
			if err = utils.RespondWithErrors(w, http.StatusInternalServerError, "Server error"); err != nil {
				h.logger.WithError(err).Error("Failed to respond with 500")
			}
			return
		}

		duration := 0.0
		if !task.EndAt.IsZero() {
			duration = task.EndAt.Sub(task.CreatedAt).Seconds()
		}

		response := dto.TaskResponse{
			Id:        task.ID,
			Status:    task.Status,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			Duration:  duration,
		}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			h.logger.WithError(err).Error("Failed to respond with 200")
		}
	}
}

func (h *Handler) getIDFromVars(w http.ResponseWriter, r *http.Request) (int64, bool) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		if err := utils.RespondWithErrors(w, http.StatusBadRequest, "Missing Task ID"); err != nil {
			h.logger.WithError(err).Error("Failed to respond with 400")
		}
		return 0, false
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		if err = utils.RespondWithErrors(w, http.StatusBadRequest, "Invalid Task ID"); err != nil {
			h.logger.WithError(err).Error("Failed to respond with 400")
		}
		return 0, false
	}
	return id, true
}
