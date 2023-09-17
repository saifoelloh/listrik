package presenter

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/listrik/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPLN struct {
	ID            primitive.ObjectID
	Name          string
	Email         string
	NoKTP         uint64
	Address       string
	Energy        uint16
	NamaKab       string
	NamaProv      string
	TahunProduksi uint16
}

func UsersPLNSuccessResponse(data *[]UserPLN) *fiber.Map {
	fmt.Println("konjay", data)
	return &fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
		"error":  nil,
	}
}

func UserPLNSuccessResponse(data *entities.UserPLN) *fiber.Map {
	user := entities.UserPLN{
		ID:            data.ID,
		Name:          data.Name,
		Email:         data.Email,
		Energy:        data.Energy,
		NamaKab:       data.NamaKab,
		NamaProv:      data.NamaProv,
		TahunProduksi: data.TahunProduksi,
		NoKTP:         data.NoKTP,
		Address:       data.Address,
	}

	return &fiber.Map{
		"status": fiber.StatusOK,
		"data":   user,
		"error":  nil,
	}
}

func UserPLNErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": fiber.StatusBadRequest,
		"data":   nil,
		"error":  err.Error(),
	}
}
