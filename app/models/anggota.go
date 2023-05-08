package models

import (
	"ta_microservice_auth/db"
	"time"

	"gorm.io/gorm"
)

type Anggota struct {
	Id         int            `json:"id" gorm:"primarykey"`
	NRA        string         `json:"nra" gorm:"unique"`
	Password   string         `json:"password"`
	Role       *Roles         `json:"role" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Role_id    int            `json:"role_id" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;foreignKey:role_id"`
	Detail     *AnggotaDetail `json:"detail" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;foreignKey:Anggota_id"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
}

type AnggotaDetail struct {
	Id           int       `json:"id" gorm:"primarykey"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"unique"`
	Phone_Number string    `json:"phone_number"`
	Address      string    `json:"address" gorm:"text"`
	Divisi       string    `json:"divisi"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Anggota_id   int       `json:"anggota_id" gorm:"index"`
}

type Roles struct {
	Id   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}

type ReqAuth struct {
	AccesToken   string `json:"acces_token"`
	RefreshToken string `json:"refresh_token"`
}

// func CreateAnggota(a *Anggota) (*Anggota, error) {
// 	result := db.Db.Create(a)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	// Load the Role data for the created Anggota
// 	err := db.Db.Preload("Role").Find(&a.Role).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return a, nil
// }

func FindUserById(db *gorm.DB, Id int) (*Anggota, error) {
	var user Anggota
	err := db.Where("Id = ?", Id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func FindUserByNRA(db *gorm.DB, nra string) (*Anggota, error) {
	var user Anggota
	err := db.Where("nra = ?", nra).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByPassword(db *gorm.DB, Password string) (*Anggota, error) {
	var user Anggota
	err := db.Where("password = ?", Password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Delete(id int) (err error) {
	d := Anggota{}
	db.Db.Delete(d, "id = ?", id)
	if d.Id == 0 {
		return err
	}

	return nil
}

// anggota detail

func CreateAnggotaDetail(a *AnggotaDetail) (*AnggotaDetail, error) {
	result := db.Db.Create(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

func GetAnggotaDetailById(id int) (*AnggotaDetail, error) {
	var anggotaDetail AnggotaDetail
	result := db.Db.Where("anggota_id = ?", id).First(&anggotaDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	return &anggotaDetail, nil
}
