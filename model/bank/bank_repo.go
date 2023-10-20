package bank

import "github.com/jinzhu/gorm"

type BankRepo struct {
	db *gorm.DB
}

func NewBankRepo(db *gorm.DB) *BankRepo {
	return &BankRepo{db: db}
}

func (b BankRepo) GetDataBank() ([]BankList, error) {
	list := make([]BankList, 0)

	rows, e := b.db.Model(&BankList{}).Where("is_deleted = 'N'").Rows()
	defer rows.Close()

	for rows.Next() {
		var bank BankList
		b.db.ScanRows(rows, &bank)
		list = append(list, bank)
	}

	return list, e
}
