package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPLN struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	UserId         uint64             `json:"user_id" bson:"user_id"`
	MeterId        uint64             `json:"meter_id" bson:"meter_id"`
	MeterNo        uint64             `json:"meter_no" bson:"meter_no"`
	Number         uint64             `json:"number" bson:"number"`
	Name           string             `json:"name" bson:"name"`
	AliasName      string             `json:"alias_name" bson:"alias_name"`
	NoKtp          uint64             `json:"no_ktp" bson:"no_ktp"`
	Email          string             `json:"email" bson:"email"`
	Type           string             `json:"type" bson:"type"`
	Unitup         int32              `json:"unitup" bson:"unitup"`
	Unitupi        int32              `json:"unitupi" bson:"unitupi"`
	Namaup         string             `json:"namaup" bson:"namaup"`
	EnergyType     string             `json:"energy_type" bson:"energy_type"`
	Energy         uint32             `json:"energy" bson:"energy"`
	Fasa           int32              `json:"fasa" bson:"fasa"`
	IsSplu         bool               `json:"is_splu" bson:"is_splu"`
	MeterKwhNumber int64              `json:"meter_kwh_number" bson:"meter_kwh_number"`
	PeruntukanId   string             `json:"peruntukan_id" bson:"peruntukan_id"`
	Rpujl          int32              `json:"rpujl" bson:"rpujl"`
	Address        string             `json:"address" bson:"address"`
	Latitude       float64            `json:"latitude" bson:"latitude"`
	Longitude      float64            `json:"longitude" bson:"longitude"`
	KodeProv       int32              `json:"kode_prov" bson:"kode_prov"`
	KodeKab        int32              `json:"kode_kab" bson:"kode_kab"`
	NamaProv       string             `json:"nama_prov" bson:"nama_prov"`
	NamaKab        string             `json:"nama_kab" bson:"nama_kab"`
	TahunProduksi  int32              `json:"tahun_produksi" bson:"tahun_produksi"`
	HavePlts       bool               `json:"have_plts" bson:"have_plts"`
	HaveEstove     bool               `json:"have_estove" bson:"have_estove"`
}
