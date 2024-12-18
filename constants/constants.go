package constants

import "time"

const (
	TestMode  = "test"
	DebugMode = "debug"

	JWTRefreshTokenExpireDuration = time.Hour * 72
	JWTAccessTokenExpireDuration  = time.Minute * 60
	ContextTimeoutDuration        = time.Second * 7

	CustomerRoleInSignup = "customer_in_signup"
	CustomerRole         = "customer"
	UserRole             = "user"

	UzLang = "uz"
	RuLang = "ru"

	VerifyCodeLength = 6

	CasbinConfigPath    = "configs/rbac_model.conf"
	MiddlewareRolesPath = "configs/models.csv"

	FirebaseReturnURL = "https://firebasestorage.googleapis.com/v0/b/phleybo.appspot.com/o/"

	Success           = "success"
	InternelServError = "Sizni so'rovingizni bajarishda kutilmagan xatolik, Iltimos keyinroq urunib ko'ring"
	BadRequest        = "Yuborgan so'rovingizda xatolik"
	Active            = 1
	InActive          = 0
	ColorType         = 5
)

const (
	IncomeTransactionID = iota
	ExpenseTransactionID
	TransferTransactionID
)
