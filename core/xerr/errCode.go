package xerr

type ErrCode int32

// 成功返回
const OK ErrCode = 200

// 默认状态码
const (
	StatusCodeError   ErrCode = -1
	StatusCodeSuccess ErrCode = 0
	StatusCode400     ErrCode = 398 + iota
	StatusCode401
	StatusCode401_2 ErrCode = 397 + iota
	StatusCode402
	StatusCode403
	StatusCode403_6 ErrCode = 396 + iota
	StatusCode404
	StatusCode405
	StatusCode408 ErrCode = 398 + iota
	StatusCode500 ErrCode = 500
)

// 请求统一异常
const (
	ParametersInvalid  ErrCode = 10000
	HeaderLangError    ErrCode = 10002
	HeaderDeviceError  ErrCode = 10003
	ConfigInvalid      ErrCode = 10007
	HttpRequestInvalid ErrCode = 10008
)

/***************** 用户模块 20000 - 29999 ******************/

const (
	// AccountExists [账户 20100 - 20199]
	AccountExists ErrCode = 20100 + iota
	_
	AccountNotExists,
	PhoneNotExists,
	AgentAccountNotExists,
	AgentAccountInvalid,
	AgentLoginUrlInvalid,
	AgentTransferMemberToSuperAgentInvalid,
	HideRegister ErrCode = 20100 + iota, 20100 + iota,
		20100 + iota, 20100 + iota, 20100 + iota,
		20100 + iota, 20100 + iota
	_ ErrCode = 20100 + iota
	AccountStatusInvalid
	VerificationCodeInvalid
	UserWalletLocked
	UserKycUnPass
	UserWalletNotExists
	// PasswordInvalid [密码 20400 - 20499]
	PasswordInvalid   ErrCode = 20401
	GoogleCodeInvalid ErrCode = 20403
	// MobileExists [手机 20500 - 20599]
	MobileExists ErrCode = 20502 - 11 + iota
	EmailExists
	MobileRepeat
	MobilePrefixError
	MobileLengthError
	// PromoCodeInvalid [推广码 20600 - 20699]
	PromoCodeInvalid ErrCode = 20601
	// CoinTransferInvalid [金额 20700 - 20799]
	CoinTransferInvalid ErrCode = 20701
	CoinNotEnough       ErrCode = 20701 - 13 + iota
	CoinTransferOverLimit
	// LoginInvalidOverLimit [登录/登出 20800 - 20899]
	LoginInvalidOverLimit ErrCode = 20802
	// RegisterInvalid [注册 20900 - 20999]
	RegisterInvalid ErrCode = 20902 - 21 + iota
	GoogleLoginCodeExpired
	GoogleLoginEmailExists
	FacebookLoginCodeExpired
	FacebookLoginEmailExists
	// UserBoundenGoogle [三方绑定 21000 - 21100]
	UserBoundenGoogle ErrCode = 21000 - 26 + iota
	UserBoundenFacebook
	GoogleLoginFailed
	FacebookLoginFailed
	GoogleLoginBindingFailed
	FacebookLoginBindingFailed
	GoogleHasBounden
	FacebookHasBounden

	UpdateError ErrCode = 21203 - 34 + iota
	EmailInvalid
	PhoneInvalid
	BindMobileFirst
	AlreadySetPassword
	PasswordNotSet
	UidNotSet

	AgencyShipExists ErrCode = 22000
)

/***************** SMS 30000 - 30099 ******************/
const (
	SmsCodeInvalid ErrCode = 30003 + iota
	SmsCodeExpired
	SmsCodeAlreadyUsed
	SmsSentHighFrequencyError
	RegLimitMax5
	RegIpLimit
	MaybeUnableToObtainSms
	AccountAlreadyBoundedMobile
)

/***************** 优惠活动 30100 - 30199 ******************/
const (
	ActiveBetCoinNoTEnough ErrCode = 30119
	ActiveCodeRepeat       ErrCode = 30126
	ActiveCodeNull         ErrCode = 30127
	ActiveApplyIdInvalid   ErrCode = 30131
	ActiveFirstDepositAlreadyDeposit
	ActiveRepeatApply
	ActiveDepositNumOutOfLimit
	ActiveOneTimeOnlyApply1 ErrCode = 30132 - 8 + iota
	ActiveFirstOrderException
	ActiveLadderNameNotNull
	CheckInRuleNull
	TodayAlreadyCheckedIn
	ActivitySettingError
	CheckInActivityExistsInPeriod
	DepositNotMeetStandard
	CheckInActivityNotExists
	CheckInActivityExists
)

/**(前3位代表业务,后三位代表具体功能)**/
// 全局错误码
const (
	ServerCommonError ErrCode = iota + 100001
	RequestParamError
	TokenExpireError
	TokenGenerateError
	DbError
	DbUpdateAffectedZeroError
	UpdateException
	CreateException
	DeleteException
	QueryException
	RedisException
)

// ActiveBetCoinNoEnough gameCallback
const (
	ActiveBetCoinNoEnough ErrCode = 200001 + iota // 余额不足
	BetslipsRepetitionError
)

/***************** 游戏平台相关 39000 - 39999 ******************/
const (
	// GameNotExists 本平台异常
	GameNotExists ErrCode = 39001 + iota
	GameNotOpen
	GameUnderMaintenance
	GameSlotFavoriteAlready
	GameSlotFavoriteNotYet
	GameRecordsEmpty
	GameLinkGeneratorInvalid
	PlatFactoryNotExists

	// PlatInvalidParam 三方平台映射异常
	PlatInvalidParam    ErrCode = 39901
	PlatRequestFrequent ErrCode = 399899 + iota
	PlatInvalidCurrency
	PlatInvalidLanguage

	// 三方平台映射异常
)

// ApiCoinCommissionInsufficient /***************** APIEND 50000 - 59999 ******************/
const (
	ApiCoinCommissionInsufficient ErrCode = 50006
)

/***************** 图片上传 70000 - 70100 ******************/
const (
	UPLOAD_IMAGE_FAIL ErrCode = 70003
)

/*****************注单异常  70100 - 70200 ******************/
const (
	EsSynchronousError ErrCode = 70100 + iota
	EsCurrentSumBetError
	_
	SaveCoinLogError
	_
	WithdrawalOrderError
	BetslipsRefundError
	WithdrawalAmountError
	WithdrawalOrdersNoProcess
	CoinRateException
	WithdrawalAddressCountException
	BetslipsRefundOrderNothingnessError
	WithdrawalAddressNotExists
	_
	NoBettingWasRecorded
	RepeatedBets
	NoticeUidException
	SaveStreamFailure
)

/*****************Digi 体育异常  70300 - 70400 ******************/
const (
	SignatureVerificationError ErrCode = 70300 + iota
	PartnerError
	CurrencyError
	_
	UserNotExist
	BetTransactionIdExist
	BetOrderNumberExist
	BetAmountMax
	RecordDoesNotExist
	WithdrawalAmountAbnormal
	PayChannelIsClose
	PayPlatIsClose
	WithdrawalHasBeenInitiated
	CoinInvalid
	ActivityIsContinue
	WithdrawalIsDeposit
	PictureCodeException
	UserBlack
	InsufficientPoints
)

/******新加plat 结算code*****/
const (
	MemberOnLine                  = 0
	InvalidSecretKey              = 1
	MemberNotExist                = 2
	OffLine                       = 3
	AccountNotExist               = 203
	InvalidAuthenticationKey      = 311
	InsufficientBalance           = 502
	IpIsRestricted                = 503
	BetNotFound                   = 504
	InfoNotFound                  = 505
	CannotExecute                 = 506
	BettingNotEligibleForBonus    = 507
	PlacebetSeriesCannotBeRetried = 508
	DuplicateTransaction          = 509
	InvalidAmount                 = 510
	NetworkError                  = 902
	SystemMaintenance             = 903
	RequestTimeout                = 904
	SystemBusy                    = 905
	SystemError                   = 999
)
