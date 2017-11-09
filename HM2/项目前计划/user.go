// currentUser是当前User，如果没有登录为nil
var CurrentUser userItem

// 判断当前有没有用户登录，并不是很必要
func IsLogin() bool {}

// 注册用户，如果用户名一样，则返回err
func RegisterUser(name string, password string, email string, phoneNumber string) error {}

// 登录用户
// 如果用户名不存在，则返回err = errors.New("name")
// 或者用户名密码不正确，则返回err = errors.New("password")
func LoginUser(name string, password string) error {}

// 登出用户，如果当前没有用户登录，则返回err
func LogoutUser() error {}

// 列出当前所有用户名，邮箱，密码并组合成字符串返回
// 如果当前没有用户登录，返回err
func ListUsers() (string, error) {}

// 删除当前登录用户，删除后当前登录用户置为nil
// 如果当前没有用户登录，返回err
func DeleteUser() error {}

// 判断当前姓名的用户是否注册
func IsRegisteredUser(name string) bool {}

// 得到当前已登录用户的姓名，如果没有登录，返回""
func GetLogonUsername() string {}