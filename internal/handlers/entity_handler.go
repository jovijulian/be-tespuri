package handlers

import (
	"errors"
	"strconv"

	"api-puri/internal/dto"
	"api-puri/internal/services"
	"api-puri/pkg/i18n"
	"api-puri/pkg/response"
	"api-puri/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type EntityHandler struct {
	service services.EntityService
}

func NewEntityHandler(service services.EntityService) *EntityHandler {
	return &EntityHandler{service: service}
}

func (h *EntityHandler) Get(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	if page < 1 {
		page = 1
	}
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	if perPage < 1 || perPage > 100 {
		perPage = 20
	}
	search := c.Query("search", "")
	category := c.Query("category", "")
	status := c.Query("status", "")

	data, total, err := h.service.GetPaginated(page, perPage, search, category, status)
	if err != nil {
		return response.InternalServerError(c, i18n.MsgInternalServerError)
	}

	return response.Paginated(c, i18n.EntityListRetrieved, data, total, page, perPage)
}

func (h *EntityHandler) GetMapPoints(c *fiber.Ctx) error {
	category := c.Query("category", "")
	status := c.Query("status", "")

	data, err := h.service.GetMapPoints(category, status)
	if err != nil {
		return response.InternalServerError(c, i18n.MsgInternalServerError)
	}

	return response.Success(c, i18n.EntityListRetrieved, data)
}

func (h *EntityHandler) Detail(c *fiber.Ctx) error {
	idParam := c.Params("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid entity ID format", nil)
	}

	data, err := h.service.GetByID(uint(parsedID))
	if err != nil {
		if errors.Is(err, services.ErrEntityNotFound) {
			return response.NotFound(c, i18n.EntityNotFound)
		}
		return response.InternalServerError(c, i18n.MsgInternalServerError)
	}

	return response.Success(c, i18n.MsgSuccess, data)
}

func (h *EntityHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateEntityRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid JSON body payload", nil)
	}

	if errs := validator.ValidateStruct(c, &req); len(errs) > 0 {
		return response.BadRequest(c, i18n.MsgValidationError, errs)
	}

	data, err := h.service.Create(&req)
	if err != nil {
		if errors.Is(err, services.ErrEntityCodeExists) {
			return response.Conflict(c, i18n.EntityCodeExists)
		}
		return response.InternalServerError(c, i18n.MsgInternalServerError)
	}

	return response.Created(c, i18n.EntityCreatedSuccess, data)
}

func (h *EntityHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid entity ID format", nil)
	}

	var req dto.UpdateEntityRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid JSON body payload", nil)
	}

	if errs := validator.ValidateStruct(c, &req); len(errs) > 0 {
		return response.BadRequest(c, i18n.MsgValidationError, errs)
	}

	data, err := h.service.Update(uint(parsedID), &req)
	if err != nil {
		if errors.Is(err, services.ErrEntityNotFound) {
			return response.NotFound(c, i18n.EntityNotFound)
		}
		return response.InternalServerError(c, i18n.MsgInternalServerError)
	}

	return response.Success(c, i18n.EntityUpdatedSuccess, data)
}

func (h *EntityHandler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return response.BadRequest(c, "Invalid entity ID format", nil)
	}

	if err := h.service.Delete(uint(parsedID)); err != nil {
		if errors.Is(err, services.ErrEntityNotFound) {
			return response.NotFound(c, i18n.EntityNotFound)
		}
		return response.InternalServerError(c, i18n.MsgInternalServerError)
	}

	return response.Success(c, i18n.EntityDeletedSuccess, nil)
}
