package models

// type Auth struct {
// 	Id         uint      `json:"id" gorm:"primarykey"`
// 	Username   string    `json:"username"`
// 	Password   string    `json:"password"`
// 	Role       string    `json:"role"`
// 	Created_at time.Time `json:"created_at"`
// 	Updated_at time.Time `json:"updated_at"`
// }

// func FindUserByd(db *gorm.DB, Id uint) (*Auth, error) {
// 	var user Auth
// 	err := db.Where("Id = ?", Id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func FindUserByUsername(db *gorm.DB, username string) (*Auth, error) {
// 	var user Auth
// 	err := db.Where("username = ?", username).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func FindUserByPasswrd(db *gorm.DB, Password string) (*Auth, error) {
// 	var user Auth
// 	err := db.Where("password = ?", Password).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func CreateUser(db *gorm.DB, a *Auth) (*Auth, error) {
// 	result := db.Create(a)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return a, nil
// }

// func GetAll(query *gorm.DB, Package *[]Alat) (err error) {
// 	err = query.Find(Package).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func GetAlatById(id uint) (*Alat, error) {
// 	p := Alat{}
// 	db.Db.First(&p, id)
// 	if p.Id == 0 {
// 		return nil, errors.New("Alat not found")
// 	}
// 	return &p, nil
// }

// func Delete(id uint) {

// }

// func (repo *Alat) Update() {
// 	db.Db.Save(&repo)
// }
