package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPLN struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	Email         string             `bson:"email"`
	NoKTP         uint64             `bson:"no_ktp"`
	Address       string             `bson:"address"`
	Energy        uint16             `bson:"energy"`
	NamaKab       string             `bson:"nama_kab"`
	NamaProv      string             `bson:"nama_prov"`
	TahunProduksi uint16             `bson:"tahun_produksi"`
}
