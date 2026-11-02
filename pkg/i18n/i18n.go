package i18n

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Language string

const (
	LangEN Language = "en"
	LangID Language = "id"
)

const (
	MsgSuccess             = "common.success"
	MsgCreated             = "common.created"
	MsgBadRequest          = "common.bad_request"
	MsgValidationError     = "common.validation_error"
	MsgUnauthorized        = "common.unauthorized"
	MsgForbidden           = "common.forbidden"
	MsgNotFound            = "common.not_found"
	MsgConflict            = "common.conflict"
	MsgInternalServerError = "common.internal_server_error"

	EntityCreatedSuccess   = "entity.created_success"
	EntityUpdatedSuccess   = "entity.updated_success"
	EntityDeletedSuccess   = "entity.deleted_success"
	EntityNotFound         = "entity.not_found"
	EntityCodeExists       = "entity.code_already_exists"
	EntityListRetrieved    = "entity.list_retrieved"
)

var dictionary = map[Language]map[string]string{
	LangEN: {
		MsgSuccess:             "Success",
		MsgCreated:             "Resource created successfully",
		MsgBadRequest:          "Bad request format or parameter",
		MsgValidationError:     "Validation failed for one or more fields",
		MsgUnauthorized:        "Authentication credentials missing or invalid",
		MsgForbidden:           "You do not have permission to perform this action",
		MsgNotFound:            "Resource not found",
		MsgConflict:            "Data conflict or resource already exists",
		MsgInternalServerError: "An unexpected internal server error occurred",

		EntityCreatedSuccess:   "Entity created successfully",
		EntityUpdatedSuccess:   "Entity updated successfully",
		EntityDeletedSuccess:   "Entity deleted successfully",
		EntityNotFound:         "Entity not found",
		EntityCodeExists:       "Entity code is already in use",
		EntityListRetrieved:    "Entities retrieved successfully",
	},
	LangID: {
		MsgSuccess:             "Berhasil",
		MsgCreated:             "Data berhasil dibuat",
		MsgBadRequest:          "Format atau parameter request tidak valid",
		MsgValidationError:     "Validasi input gagal",
		MsgUnauthorized:        "Autentikasi gagal atau sesi berakhir",
		MsgForbidden:           "Anda tidak memiliki izin akses",
		MsgNotFound:            "Data tidak ditemukan",
		MsgConflict:            "Data sudah ada atau terjadi konflik",
		MsgInternalServerError: "Terjadi kesalahan internal pada server",

		EntityCreatedSuccess:   "Entitas berhasil ditambahkan",
		EntityUpdatedSuccess:   "Entitas berhasil diperbarui",
		EntityDeletedSuccess:   "Entitas berhasil dihapus",
		EntityNotFound:         "Entitas tidak ditemukan",
		EntityCodeExists:       "Kode entitas sudah digunakan",
		EntityListRetrieved:    "Daftar entitas berhasil diambil",
	},
}

func ResolveLanguage(c *fiber.Ctx) Language {
	if c == nil {
		return LangEN
	}
	lang := c.Get("Accept-Language")
	if strings.HasPrefix(strings.ToLower(lang), "id") {
		return LangID
	}
	return LangEN
}

func T(c *fiber.Ctx, key string) string {
	lang := ResolveLanguage(c)
	if translations, ok := dictionary[lang]; ok {
		if msg, found := translations[key]; found {
			return msg
		}
	}
	if translations, ok := dictionary[LangEN]; ok {
		if msg, found := translations[key]; found {
			return msg
		}
	}
	return key
}
