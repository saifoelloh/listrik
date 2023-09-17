package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/listrik/pkg/entities"
)

func UsersPLNSuccessResponse(data *[]entities.UserPLN) *fiber.Map {
	return &fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
		"error":  nil,
	}
}

func UserPLNSuccessResponse(data *entities.UserPLN) *fiber.Map {
	user := entities.UserPLN{
		ID:             data.ID,
		UserId:         data.UserId,
		MeterId:        data.MeterId,
		MeterNo:        data.MeterNo,
		Number:         data.Number,
		Name:           data.Name,
		AliasName:      data.AliasName,
		NoKtp:          data.NoKtp,
		Email:          data.Email,
		Type:           data.Type,
		Unitup:         data.Unitup,
		Unitupi:        data.Unitupi,
		Namaup:         data.Namaup,
		EnergyType:     data.EnergyType,
		Energy:         data.Energy,
		Fasa:           data.Fasa,
		IsSplu:         data.IsSplu,
		MeterKwhNumber: data.MeterKwhNumber,
		PeruntukanId:   data.PeruntukanId,
		Rpujl:          data.Rpujl,
		Address:        data.Address,
		Latitude:       data.Latitude,
		Longitude:      data.Longitude,
		KodeProv:       data.KodeProv,
		KodeKab:        data.KodeKab,
		NamaProv:       data.NamaProv,
		NamaKab:        data.NamaKab,
		TahunProduksi:  data.TahunProduksi,
		HavePlts:       data.HavePlts,
		HaveEstove:     data.HaveEstove,
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
