package response

import (
	"reflect"
	"time"

	"api-puri/pkg/i18n"
	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Total      int64  `json:"total"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	TotalPages int    `json:"total_pages"`
	ServerTime string `json:"server_time"`
}

func resolveMessage(c *fiber.Ctx, message string, fallbackKey string) string {
	if message == "" {
		return i18n.T(c, fallbackKey)
	}
	return i18n.T(c, message)
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Success: true,
		Message: resolveMessage(c, message, i18n.MsgSuccess),
		Data:    data,
	})
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(ApiResponse{
		Success: true,
		Message: resolveMessage(c, message, i18n.MsgCreated),
		Data:    data,
	})
}

func Paginated(c *fiber.Ctx, message string, data interface{}, total int64, page int, perPage int) error {
	if data == nil {
		data = []interface{}{}
	} else {
		val := reflect.ValueOf(data)
		if val.Kind() == reflect.Slice && val.IsNil() {
			data = []interface{}{}
		}
	}
	totalPages := 0
	if perPage > 0 {
		totalPages = int(total) / perPage
		if int(total)%perPage != 0 {
			totalPages++
		}
	}

	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Success: true,
		Message: resolveMessage(c, message, i18n.MsgSuccess),
		Data:    data,
		Meta: &Meta{
			Total:      total,
			Page:       page,
			PerPage:    perPage,
			TotalPages: totalPages,
			ServerTime: time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func BadRequest(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(ApiResponse{
		Success: false,
		Message: resolveMessage(c, message, i18n.MsgBadRequest),
		Data:    nil,
		Errors:  errors,
	})
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(ApiResponse{
		Success: false,
		Message: resolveMessage(c, message, i18n.MsgUnauthorized),
		Data:    nil,
	})
}

func Forbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(ApiResponse{
		Success: false,
		Message: resolveMessage(c, message, i18n.MsgForbidden),
		Data:    nil,
	})
}

func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(ApiResponse{
		Success: false,
		Message: resolveMessage(c, message, i18n.MsgNotFound),
		Data:    nil,
	})
}

func Conflict(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusConflict).JSON(ApiResponse{
		Success: false,
		Message: resolveMessage(c, message, i18n.MsgConflict),
		Data:    nil,
	})
}

func InternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(ApiResponse{
		Success: false,
		Message: resolveMessage(c, message, i18n.MsgInternalServerError),
		Data:    nil,
	})
}
