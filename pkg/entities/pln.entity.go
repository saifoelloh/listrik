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
	Unitup         uint16             `json:"unitup" bson:"unitup"`
	Unitupi        uint8              `json:"unitupi" bson:"unitupi"`
	Namaup         string             `json:"namaup" bson:"namaup"`
	EnergyType     string             `json:"energy_type" bson:"energy_type"`
	Energy         uint16             `json:"energy" bson:"energy"`
	Fasa           string             `json:"fasa" bson:"fasa"`
	IsSplu         string             `json:"is_splu" bson:"is_splu"`
	MeterKwhNumber uint64             `json:"meter_kwh_number" bson:"meter_kwh_number"`
	PeruntukanId   string             `json:"peruntukan_id" bson:"peruntukan_id"`
	Rpujl          uint16             `json:"rpujl" bson:"rpujl"`
	Address        string             `json:"address" bson:"address"`
	Latitude       int64              `json:"latitude" bson:"latitude"`
	Longitude      int64              `json:"longitude" bson:"longitude"`
	KodeProv       uint16             `json:"kode_prov" bson:"kode_prov"`
	KodeKab        uint16             `json:"kode_kab" bson:"kode_kab"`
	NamaProv       string             `json:"nama_prov" bson:"nama_prov"`
	NamaKab        string             `json:"nama_kab" bson:"nama_kab"`
	TahunProduksi  uint16             `json:"tahun_produksi" bson:"tahun_produksi"`
	HavePlts       string             `json:"have_plts" bson:"have_plts"`
	HaveEstove     string             `json:"have_estove" bson:"have_estove"`
}
