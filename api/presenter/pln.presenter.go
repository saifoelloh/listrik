package presenter

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/listrik/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPLN struct {
	ID             primitive.ObjectID
	UserId         uint64
	MeterId        uint64
	MeterNo        uint64
	Number         uint64
	Name           string
	AliasName      string
	NoKtp          uint64
	Email          string
	Type           string
	Unitup         uint16
	Unitupi        uint8
	Namaup         string
	EnergyType     string
	Energy         uint16
	Fasa           string
	IsSplu         string
	MeterKwhNumber uint64
	PeruntukanId   string
	Rpujl          uint16
	Address        string
	Latitude       int64
	Longitude      int64
	KodeProv       uint16
	KodeKab        uint16
	NamaProv       string
	NamaKab        string
	TahunProduksi  uint16
	HavePlts       string
	HaveEstove     string
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
