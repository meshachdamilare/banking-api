package database

import (
	"crypto/rand"
	"fmt"
	"github.com/meshachdamilare/banking-api/config"
	"github.com/meshachdamilare/banking-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var databaseURI string = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s",
	config.GetEnv("POSTGRES_HOSTNAME"),
	config.GetEnv("POSTGRES_USERNAME"),
	config.GetEnv("POSTGRES_PASSWORD"),
	config.GetEnv("POSTGRES_DATABASE"),
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func verifyPassword(userPassword string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil {
		return err
	}
	return nil
}

// MigratePostgres Create tables using structs.
func MigratePostgres() error {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Account{}, &models.User{})
	return nil
}

// CreateUser Add a new user to the database. The user password is hashed
func CreateUser(user *models.User) (*string, error) {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	passwordHasH, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = passwordHasH
	accNumber, _ := rand.Prime(rand.Reader, 32)
	account := &models.Account{
		Email:         user.Email,
		User:          user,
		Balance:       0,
		AccountNumber: accNumber.String(),
	}
	if query := db.Create(account); query.Error != nil {
		return nil, query.Error
	}
	return &account.AccountNumber, nil
}

// AuthUser Validate user password. The hash of password is compared
// in order to validate the password.
func AuthUser(user *models.UserAuth) error {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return err
	}

	result := models.User{}

	query := db.Where("email = ?", user.Email).First(&result)
	if query.Error != nil {
		return query.Error
	}
	err = verifyPassword(result.Password, user.Password)

	if err != nil {
		return err
	}
	return nil
}

// GetUser Get user detail of the email provided
func GetUser(email string) (*models.User, error) {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	result := models.User{}

	if query := db.Where("email = ?", email).First(&result); query.Error != nil {
		return nil, query.Error
	}
	return &result, nil
}

// GetAccount Get account detail of the provided mail
func GetAccount(email string) (*models.Account, error) {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	result := models.Account{}
	query := db.Preload(clause.Associations).Where("email =?", email).First(&result)
	if query.Error != nil {
		return nil, query.Error
	}
	return &result, nil
}

// UpdateAccountBalance Update users' account balance
func UpdateAccountBalance(email string, balance uint64) error {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return err
	}

	query := db.Model(&models.Account{}).Where("email = ?", email).Update("balance", balance)
	if query.Error != nil {
		return err
	}
	return nil
}

// UpdateAccountEmail Update account email and user email
func UpdateAccountEmail(email string, newEmail string) error {
	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return err
	}
	query := db.Where("email = ?", email).Update("email", newEmail)
	if query.Error != nil {
		return err
	}
	return nil
}
