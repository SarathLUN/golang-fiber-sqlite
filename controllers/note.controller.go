package controllers

import (
	"github.com/SarathLUN/golang-fiber-sqlite/initializers"
	"github.com/SarathLUN/golang-fiber-sqlite/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

// CreateNote accept *fiber.Ctx and return error if any
func CreateNote(c *fiber.Ctx) error {
	var payload *models.CreateNoteSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "fail parse body: " + err.Error(),
		})
	}

	// run ValidateStruct of payload
	if err := models.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	// create new note
	note := models.Note{
		Title:     payload.Title,
		Content:   payload.Content,
		Category:  payload.Category,
		Published: payload.Published,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// save note to database
	result := initializers.DB.Create(&note)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Note created successfully",
		"data":    fiber.Map{"note": note},
	})
}

// FindNotes accept *fiber.Ctx and return error if any
func FindNotes(c *fiber.Ctx) error {
	// implement pagination
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	// find all notes
	var notes []models.Note
	result := initializers.DB.Limit(intLimit).Offset(offset).Find(&notes)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Notes fetched successfully",
		"data":    fiber.Map{"notes": notes},
	})
}

// UpdateNote accept `*fiber.Ctx` to update existing record base on given param: `noteId` and return error if any
func UpdateNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")
	var payload *models.UpdateNoteSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "fail parse body: " + err.Error(),
		})
	}

	// run ValidateStruct of payload
	if err := models.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// find note by id
	var note models.Note
	result := initializers.DB.First(&note, "id = ?", noteId)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "Note not found",
		})
	}

	// update note
	updatedResult := initializers.DB.Model(&note).Updates(models.UpdateNoteSchema{
		Title:     payload.Title,
		Content:   payload.Content,
		Category:  payload.Category,
		Published: payload.Published,
		UpdatedAt: time.Now(),
	})
	if updatedResult.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": updatedResult.Error.Error(),
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Note updated successfully",
		"data":    fiber.Map{"note": note},
	})
}

// FindNoteById accept `*fiber.Ctx` to find note by given param: `noteId` and return error if any
func FindNoteById(c *fiber.Ctx) error {
	noteId := c.Params("noteId")
	var note models.Note
	result := initializers.DB.First(&note, "id = ?", noteId)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "Note not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Note fetched successfully",
		"data":    fiber.Map{"note": note},
	})
}

// DeleteNote accept `*fiber.Ctx` to delete note by given param: `noteId` and return error if any
func DeleteNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")
	var note models.Note
	result := initializers.DB.First(&note, "id = ?", noteId)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "Note not found",
		})
	}
	result = initializers.DB.Delete(&note)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Note deleted successfully",
	})
}
