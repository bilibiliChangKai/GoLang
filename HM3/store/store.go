package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// UserItem 用户结构
type UserItem struct {
	// 注册用学号
	ID string
	// 用户名字
	Name string
	// hash过的密码
	Password string
	// 注册用的邮箱
	Email string
	// 注册用的电话号码
	PhoneNumber string
}

// NewUser 同New
func NewUser(ID string, Name string, Password string, Email string, PhoneNumber string) *UserItem {
	return &UserItem{ID, Name, Password, Email, PhoneNumber}
}

// AddUser 新加入一个User,若存在返回error
func AddUser(user UserItem) error {
	if _, ok := userItems[user.Name]; ok {
		return errors.New("ERROR: Couldn't add existed user")
	}

	userItems[user.Name] = user
	writeJSON()
	return nil
}

// IsExistedUser 判断该用户是否已注册
func IsExistedUser(userName string) bool {
	_, ok := userItems[userName]
	return ok
}

// 文件名
var userItemsFilePath = "store/Json/UserItems.json"

// 用于储存用户
var userItems map[string](UserItem)

func init() {
	userItems = make(map[string](UserItem))
	readJSON()
}

func readJSON() {
	// 解析userItems
	b1, err1 := ioutil.ReadFile(userItemsFilePath)
	if err1 == nil {
		json.Unmarshal(b1, &userItems)
	}
}

func writeJSON() {
	// 写入userItems
	b1, err := json.Marshal(userItems)

	if err == nil {
		if _, err := os.Open(userItemsFilePath); err != nil {
			os.Create(userItemsFilePath)
		}
		ioutil.WriteFile(userItemsFilePath, b1, 0755)
	}
}
