package seed

import (
	"ta_microservice_auth/app/models"
	"ta_microservice_auth/db"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		// {
		// 	Name: "Delete",
		// 	Run: func(db *gorm.DB) error {
		// 		err := DeleteSeed()
		// 		return err
		// 	},
		// },
		{
			Name: "Create Roles",
			Run: func(d *gorm.DB) error {
				err := CreateRole(d, "admin")
				return err
			},
		},
		{
			Name: "Create Roles",
			Run: func(d *gorm.DB) error {
				err := CreateRole(d, "anggota")
				return err
			},
		},
		// {
		// 	Name: "Create Akun Admin 1",
		// 	Run: func(db *gorm.DB) error {
		// 		err := CreateAdmin(db, "I.288IP", "nanda123", 1)
		// 		return err
		// 	},
		// },
	}
}

// func CreateAdmin(db *gorm.DB, username string, password string, role int) error {
// 	create := &models.Anggota{
// 		NRA:      username,
// 		Password: password,
// 		Role_id:  role,
// 	}

// 	err, _ := models.CreateAnggota(create)
// 	if err == nil {
// 		return fmt.Errorf("Failed to create admin: %v", err)
// 	}

// 	return nil
// }

func CreateRole(db *gorm.DB, name string) error {

	return db.Create(&models.Roles{
		Name: name,
	}).Error
}

func DeleteSeed() error {
	// 	return app.Db.Delete(models.Packages{}).Error
	err := db.Db.Exec("DELETE FROM roles").Error
	return err
}
