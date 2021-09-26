package repository

import (
	"github.com/afz204/golang-second-api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// contract what user can do to db
type UserResponsitory interface {
	Insertuser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (trx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

// is create a new instance of UserReponsitory
func NewUserReponsitory(db *gorm.DB) UserResponsitory {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Insertuser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}
func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
func (db *userConnection) IsDuplicateEmail(email string) (trx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}
func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}
func (db *userConnection) ProfileUser(UserId string) entity.User {
	var user entity.User
	db.connection.Find(&user, UserId)
	return user
}


func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("Failed to has password!")
	}

	return string(hash)
}
