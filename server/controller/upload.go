package controller

import (
	"image"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
)

type uploadController struct {
	cfg *domain.ConfigEnv
}

func NewUploadController(cfg *domain.ConfigEnv) *uploadController {
	return &uploadController{
		cfg: cfg,
	}
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

// @Summary Get uploaded file
// @Description Get an uploaded file by filename
// @Tags Upload
// @Produce application/octet-stream
// @Param filename path string true "Filename"
// @Success 200 {file} file
// @Failure 404 {object} map[string]string
// @Router /api/upload/{filename} [get]
func (c *uploadController) GetFile(ctx *fiber.Ctx) error {
	filename := ctx.Params("filename")
	if filename == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Filename is required",
		})
	}

	filePath := filepath.Join("uploads", filename)
	return ctx.SendFile(filePath)
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
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: "Failed to get file",
			SUCCESS: false,
		})
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			MESSAGE: "Failed to open file",
			SUCCESS: false,
		})
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to read file",
		})
	}
	if _, err := src.Seek(0, 0); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to seek file",
		})
	}

	fileType := http.DetectContentType(buffer)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}

	if !allowedTypes[fileType] {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: "File type not allowed",
			SUCCESS: false,
		})
	}

	var img image.Image
	if fileType == "image/webp" {
		img, err = webp.Decode(src)
	} else {
		img, _, err = image.Decode(src)
	}
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			MESSAGE: "Failed to decode image",
			SUCCESS: false,
		})
	}

	const maxWidth = 1920
	const maxHeight = 1080
	resizedImg := imaging.Fit(img, maxWidth, maxHeight, imaging.Lanczos)
	timestamp := time.Now().Format("20060102150405")
	sanitizedFileName := sanitizeFileName(file.Filename)
	newFileName := timestamp + "_" + sanitizedFileName
	dst := filepath.Join("uploads", newFileName)

	if fileType == "image/webp" {
		outFile, err := os.Create(dst)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
				MESSAGE: "Failed to save image",
				SUCCESS: false,
			})
		}
		defer outFile.Close()
		if err := webp.Encode(outFile, resizedImg, &webp.Options{Lossless: false, Quality: 80}); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
				MESSAGE: "Failed to encode WebP",
				SUCCESS: false,
			})
		}
	} else {
		if err := imaging.Save(resizedImg, dst, imaging.JPEGQuality(85)); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
				MESSAGE: "Failed to save image",
				SUCCESS: false,
			})
		}
	}

	path := "/api/upload/" + newFileName
	url := c.cfg.BACKEND_URL + path
	return ctx.JSON(fiber.Map{
		"message": "Image uploaded and processed successfully",
		"file":    newFileName,
		"url":     url,
		"path":    path,
		"type":    fileType,
		"size":    file.Size,
	})

}
