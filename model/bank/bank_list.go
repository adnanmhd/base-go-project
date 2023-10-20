package bank

type BankList struct {
	Prefix    string `gorm:"column:prefix"`
	BankCode  string `gorm:"column:bank_code"`
	BankName  string `gorm:"column:bank_name"`
	IsDeleted string `gorm:"column:is_deleted"`
}

func (b BankList) TableName() string {
	return "bank_list"
}
