package constant

const LimitPages = 6
const PortWeb string = "8080"
const IpAddress = "localhost"

const AddressResetPassword = "http://" + IpAddress + ":" + PortWeb + "/resetpass"

const (
	DefaultRoleAdmin             = "Admin"
	DefaultRoleUser              = "User"
	DefaultTypeHyBirder   int    = 2
	DefaultTypeHyLearning int    = 3
	DefaultTypeHyNew      int    = 1
	ImageTypePost         int    = 1
	ImageTypeProfile      int    = 0
	ImageTypeSmallImage   int    = 2
	WaitTimeSendEmail     string = "Please wait 5 minute to try again"
)
const (
	ErrorEmailPasswordEmpty string = "Email and password is empty"
	ErrorWrongEmailPassword string = "Wrong password or email"
	ErrorWrongEmail         string = "Wrong email"
	ErrorCreateCategory     string = "Error create category "
	ErrorEditCategory       string = "Error Edit  category "
	ErrorDeleteCategory     string = "Error Delete  category "
	ErrorCreateSubCategory  string = "Error create sub-category "
	ErrorEditSubCategory    string = "Error Edit sub-category "
	ErrorDeleteSubCategory  string = "Error Delete  sub-category "
	ErrorCreateUser         string = "Error create user "
	ErrorEditUser           string = "Error Edit user"
	ErrorDeleteUser         string = "Error Delete  user"
	ErrorCreatePost         string = "Error create Post  "
	ErrorEditPost           string = "Error Edit Post "
	ErrorDeletePost         string = "Error Delete  Post "
	ErrorChangePassword     string = "Error change password"
	ImageNotFormat          string = "Image is not in the correct format"
	DuplicateEmail          string = "Duplicate email"
	InformationNotEmpty     string = "Information is not empty"
	IDNotFound              string = "ID not Found"
	NoImages                string = "Not image"
	TokenExpiration         string = "Token Expiration"
	NoPermission            string = "You do not have permission to access"
)

const (
	RequestCreateUserNotEmpty string = "Email or userName or password cannot be empty "
	RequestLengthEmail        string = "Email length is less than 225 characters"
	RequestLengthPassword     string = "Password length is less than 225 characters"
	RequestLengthName         string = "Name length is less than 225 characters"
	RequestOldPassword        string = "Old password is not correct"
	RequestPasswordConfirm    string = "Password and confirm password is empty"
)

const (
	NotificationCheckEmail               string = "Check email or email spam"
	NotificationCreateSuccessCategory    string = "Create success category"
	NotificationEditSuccessCategory      string = "Edit success category"
	NotificationDeleteSuccessCategory    string = "Delete success category"
	NotificationCreateSuccessSubCategory string = "Create success sub-category"
	NotificationEditSuccessSubCategory   string = "Edit success sub-category"
	NotificationDeleteSuccessSubCategory string = "Delete success sub-category"
	NotificationCreateUser               string = "Create success user"
	NotificationEditUser                 string = "Edit success user"
	NotificationDeleteUser               string = "Delete success user"
	NotificationCreatePost               string = "Create success Post "
	NotificationEditPost                 string = "Edit success Post "
	NotificationDeletePost               string = "Delete success Post "
	ChangePasswordSuccessfully           string = "Change password successfully"
)

const (
	URLHome          string = "/home"
	URLLogin         string = "/login"
	URLCategories    string = "/category/s"
	URLSubCategories string = "/subCategory/s"
	URLUsers         string = "/user/s"
)

const (
	JSONUrl     string = "Url"
	JSONError   string = "Error"
	JSONSuccess string = "Success"
)

const (
	SessionEmail    string = "SessionEmail"
	SessionUserID   string = "SessionUserID"
	SessionAvatar   string = "SessionAvatar"
	SessionName     string = "SessionName"
	SessionRole     string = "SessionRole"
	SessionIDClient string = "SessionIDClient"
)

const (
	FormPassword string = "password"
	FormEmail    string = "email"
)
