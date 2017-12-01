package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// // Create .
// func (*UserInfoAtomicService) Create() error {
// 	tx, err := mydb.Begin()
// 	checkErr(err)
//
// 	dao := userInfoDao{tx}
// 	err = dao.Create()
//
// 	if err == nil {
// 		tx.Commit()
// 	} else {
// 		tx.Rollback()
// 	}
//
// 	return nil
// }

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := mydb.Insert(u)
	checkErr(err)
	return err
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	as := []UserInfo{}
	err := mydb.Desc("id").Find(&as)
	checkErr(err)
	return as
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	a := &UserInfo{}
	_, err := mydb.Id(id).Get(a)
	checkErr(err)
	return a
}
