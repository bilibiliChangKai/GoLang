package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Create .
func (*UserInfoAtomicService) Create() error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userInfoDao{tx}
	err = dao.Create()

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return nil
}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userInfoDao{tx}
	err = dao.Save(u)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{mydb}
	return dao.FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	dao := userInfoDao{mydb}
	return dao.FindByID(id)
}
