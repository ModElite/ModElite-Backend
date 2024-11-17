package controller

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type uploadController struct{}

func NewUploadController() *uploadController {
	return &uploadController{}
}

func sanitizeFileName(fileName string) string {
	// Remove spaces and special characters
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = strings.Map(func(r rune) rune {
		if r == '_' || r == '.' || r == '-' || ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9') {
			return r
		}
		return -1
	}, fileName)
	return fileName
}

// @Summary Upload file
// @Description Upload a file
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/upload [post]
func (c *uploadController) UploadFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get file",
		})
	}

	timestamp := time.Now().Format("20060102150405")
	sanitizedFileName := sanitizeFileName(file.Filename)
	newFileName := timestamp + "_" + sanitizedFileName
	dst := filepath.Join("uploads", newFileName)
	if err := ctx.SaveFile(file, dst); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save file",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "File uploaded successfully",
		"file":    dst,
	})
}
