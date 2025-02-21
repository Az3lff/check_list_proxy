package http

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"

	"github.com/Az3lff/check_list_proxy/internal/config"
	"github.com/Az3lff/check_list_proxy/internal/models"
	"github.com/Az3lff/check_list_proxy/internal/service"
)

type Handler struct {
	cfg   config.HTTPServer
	valid *validator.Validate
	svs   *service.Service
}

func NewHandler(cfg config.HTTPServer, svs *service.Service) *Handler {
	return &Handler{
		cfg:   cfg,
		valid: validator.New(),
		svs:   svs,
	}
}

func (h *Handler) CreateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), h.cfg.ResponseTimeout)
	defer cancel()

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Errorf("failed to parse userID: %w", err).Error()})
	}

	var reqBody models.CreateTaskRequest
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.valid.Struct(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := h.svs.CreateTask(ctx, userID, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *Handler) GetList(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), h.cfg.ResponseTimeout)
	defer cancel()

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Errorf("failed to parse userID: %w", err).Error()})
	}

	resp, err := h.svs.GetList(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *Handler) DeleteTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), h.cfg.ResponseTimeout)
	defer cancel()

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Errorf("failed to parse userID: %w", err).Error()})
	}

	var reqBody models.DeleteTaskRequest
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.valid.Struct(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := h.svs.DeleteTask(ctx, userID, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *Handler) DoneTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), h.cfg.ResponseTimeout)
	defer cancel()

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Errorf("failed to parse userID: %w", err).Error()})
	}

	var reqBody models.DoneTaskRequest
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.valid.Struct(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := h.svs.DoneTask(ctx, userID, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
