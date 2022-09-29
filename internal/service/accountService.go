package service

import (
	"gorm.io/gorm"
	models "simple_bank/db"
	"time"
)

type AccountStore interface {
	CreateAccount(request models.Account) (int64, error)
	GetAccount(reqID uint) (*models.Account, error)
	ListAccounts() ([]*models.Account, error)
}

const (
	tableName = "simple_bank.accounts"
)

type accountService struct {
	db *gorm.DB
}

type Account struct {
	ID        int64     `gorm:"column:id"`
	Owner     string    `gorm:"column:owner"`
	Balance   int64     `gorm:"column:balance"`
	Currency  string    `gorm:"column:currency"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func NewAccountService(db *gorm.DB) AccountStore {
	return &accountService{db: db}
}

func (acc Account) TableName() string {
	return tableName
}

func (s *accountService) CreateAccount(request models.Account) (int64, error) {
	account := Account{
		ID:        0,
		Owner:     request.Owner,
		Balance:   request.Balance,
		Currency:  request.Currency,
		CreatedAt: time.Now(),
	}
	retDB := s.db.Create(&account)
	if retDB.Error != nil {
		return 0, retDB.Error
	}
	return account.ID, nil
}

func (s *accountService) GetAccount(reqID uint) (*models.Account, error) {
	var account *models.Account
	retDB := s.db.Table(tableName).Where("id = ?", reqID).Find(&account)
	if retDB.Error != nil {
		return nil, retDB.Error
	}
	return account, nil
}

func (s *accountService) ListAccounts() ([]*models.Account, error) {
	var accounts []*models.Account
	retDB := s.db.Table(tableName).Find(&accounts)
	if retDB.Error != nil {
		return nil, retDB.Error
	}
	return accounts, nil
}
