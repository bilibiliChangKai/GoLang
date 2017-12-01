package entities

type userInfoDao DaoSource

var userInfoCreateTable = "CREATE TABLE `userinfo` ( `uid` INT(10) NOT NULL AUTO_INCREMENT, `username` VARCHAR(64) NULL DEFAULT NULL, `departname` VARCHAR(64) NULL DEFAULT NULL, `created` DATE NULL DEFAULT NULL, PRIMARY KEY (`uid`));"

// Create
func (dao *userInfoDao) Create() error {
	stmt, err := dao.Prepare(userInfoCreateTable)
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	checkErr(err)
	if err != nil {
		return err
	}
	return nil
}

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	stmt, err := dao.Prepare(userInfoInsertStmt)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(u.UserName, u.DepartName, u.CreateAt)
	checkErr(err)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.UID = int(id)
	return nil
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"

// FindAll .
func (dao *userInfoDao) FindAll() []UserInfo {
	rows, err := dao.Query(userInfoQueryAll)
	checkErr(err)
	defer rows.Close()

	ulist := make([]UserInfo, 0, 0)
	for rows.Next() {
		u := UserInfo{}
		err := rows.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
		checkErr(err)
		ulist = append(ulist, u)
	}
	return ulist
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) *UserInfo {
	stmt, err := dao.Prepare(userInfoQueryByID)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	u := UserInfo{}
	err = row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
	checkErr(err)

	return &u
}
