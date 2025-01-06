package xerr

var message map[ErrCode]string

func init() {
	message = make(map[ErrCode]string)

	message[StatusCodeError] = "system error"
	message[StatusCodeSuccess] = "Success"
	message[StatusCode400] = "bad request"
	message[StatusCode401] = "Invalid Token"
	message[StatusCode401_2] = "TOKEN has expired"
	message[StatusCode402] = "you have been automatically logged out due to inactivity"
	message[StatusCode403] = "forbidden"
	message[StatusCode403_6] = "IP address rejected"
	message[StatusCode404] = "Not Found"
	message[StatusCode405] = "Resources are forbidden"
	message[StatusCode408] = "timed out"
	message[StatusCode500] = "Internal server error, please contact the administrator"

	message[ParametersInvalid] = "Parameters Invalid"
	message[HeaderLangError] = "Invalid language"
	message[HeaderDeviceError] = "Invalid device source"
	message[ConfigInvalid] = "Configuration error[%s]"
	message[HttpRequestInvalid] = "Http request invalid"

	message[AccountExists] = "Username already exists"
	message[AccountNotExists] = "Username not exists"
	message[PhoneNotExists] = "phone not exists"
	message[HideRegister] = "Sorry, the system is being upgraded and maintained, please contact online customer service to get more information and details."
	message[AgentAccountNotExists] = "Agent account not exists[%s]"
	message[AgentAccountInvalid] = "Agent account invalid[%s]"
	message[AgentLoginUrlInvalid] = "Please login the agent url"
	message[AgentTransferMemberToSuperAgentInvalid] = "Can't transfer member agent to super agent"
	message[AccountStatusInvalid] = "Account status is abnormal, please contact the online customer"
	message[VerificationCodeInvalid] = "Verification Code not exists"
	message[UserWalletLocked] = "User wallet locked, close game then try after %s second"
	message[UserWalletNotExists] = "User wallet does not exist"
	message[UserKycUnPass] = "Cash can be withdrawn only after KYC certification"
	message[AgencyShipExists] = "Agency relationship already exists"
	message[PasswordInvalid] = "Invalid password. Please try again!"
	message[GoogleCodeInvalid] = "INVALID GOOGLE CODE"
	message[MobileExists] = "Phone number already exists"
	message[EmailExists] = "email already exists"
	message[MobilePrefixError] = "Mobile no.needs to start with 09"
	message[MobileLengthError] = "mobile repeat"
	message[MobileRepeat] = "enter the correct mobile"
	message[PromoCodeInvalid] = "Share Code Invalid"
	message[CoinTransferInvalid] = "Transaction Error"
	message[CoinNotEnough] = "Insufficient balance"
	message[CoinTransferOverLimit] = "At least 1%s"
	message[LoginInvalidOverLimit] = "Too many login errors, please contact the administrator or retrieve your password"
	message[RegisterInvalid] = "Registration error"
	message[GoogleLoginCodeExpired] = "google login code expired"
	message[GoogleLoginEmailExists] = "google login email exists"
	message[FacebookLoginCodeExpired] = "facebook login code expired"
	message[FacebookLoginEmailExists] = "facebook login user id exists"
	message[UserBoundenGoogle] = "user bounden google"
	message[UserBoundenFacebook] = "user bounden facebook"
	message[GoogleLoginFailed] = "google login failed"
	message[FacebookLoginFailed] = "facebook login failed"
	message[GoogleLoginBindingFailed] = "google login binding failed"
	message[FacebookLoginBindingFailed] = "facebook login binding failed"
	message[GoogleHasBounden] = "google has bounden"
	message[FacebookHasBounden] = "facebook has bounden"
	message[UpdateError] = "modification failed"
	message[EmailInvalid] = "Email address invalid"
	message[PhoneInvalid] = "Phone number invalid"
	message[BindMobileFirst] = "bind mobile first"
	message[AlreadySetPassword] = "already set password"
	message[PasswordNotSet] = "password not set"
	message[UidNotSet] = "uid not set"

	message[SmsCodeInvalid] = "Sms Code Invalid"
	message[SmsCodeExpired] = "Sms Code expired"
	message[SmsCodeAlreadyUsed] = "Sms Code already used"
	message[RegLimitMax5] = "Reg limit over ip[%s]"
	message[RegIpLimit] = "Reg IP limit over ip[%s]"
	message[MaybeUnableToObtainSms] = "Please contact online customer service"
	message[AccountAlreadyBoundedMobile] = "account already bounded mobile"
	message[SmsSentHighFrequencyError] = "sms send too high frequency"

	message[ActiveBetCoinNoTEnough] = "Insufficient valid bet amount!"
	message[ActiveCodeRepeat] = "Active code can not be repeated!"
	message[ActiveCodeNull] = "Active code can not be null!"
	message[ActiveApplyIdInvalid] = "Active apply id invalid"
	message[ActiveFirstDepositAlreadyDeposit] = "Already deposit can't join this active"
	message[ActiveRepeatApply] = "Active repeat apply"
	message[ActiveDepositNumOutOfLimit] = "Active deposit num out of limit"
	message[ActiveOneTimeOnlyApply1] = "Only can join 1 activity at a time"
	message[ActiveFirstOrderException] = "The first order has been completed"
	message[ActiveLadderNameNotNull] = "Ladder Name is not null"
	message[CheckInRuleNull] = "check in rule is null"
	message[TodayAlreadyCheckedIn] = "today already checked in"
	message[ActivitySettingError] = "activity setting error"
	message[CheckInActivityExistsInPeriod] = "check in activity exists in this period"
	message[DepositNotMeetStandard] = "deposit not meet the standard"
	message[CheckInActivityNotExists] = "check in activity not exists"
	message[CheckInActivityExists] = "check in activity exists"

	message[ApiCoinCommissionInsufficient] = "Commission coin Insufficient"

	message[OK] = "SUCCESS"
	message[ServerCommonError] = "The server is missing, try again later"
	message[RequestParamError] = "Parameter error"
	message[TokenExpireError] = "token is invalid, please log in again"
	message[TokenGenerateError] = "Failed to generate token"
	message[DbError] = "The database is busy, please try again later"
	message[DbUpdateAffectedZeroError] = "The number of rows affected by updating data is 0"

	message[UpdateException] = "The database is busy, please try again later"
	message[CreateException] = "The database is busy, please try again later"
	message[DeleteException] = "The database is busy, please try again later"
	message[QueryException] = "The database is busy, please try again later"
	message[RedisException] = "The database is busy, please try again later"

	message[GameNotExists] = "not exit"
	message[GameNotOpen] = "The game is close for maintenance, please wait patiently"
	message[GameUnderMaintenance] = "The game is under maintenance"
	message[GameSlotFavoriteAlready] = "Already in favorite list"
	message[GameSlotFavoriteNotYet] = "Please collect to the favorite list first"
	message[GameRecordsEmpty] = "No such game record!"
	message[GameLinkGeneratorInvalid] = "Game link generator invalid, please try again later"
	message[PlatFactoryNotExists] = "Platform factory does not exist"
	message[PlatInvalidParam] = "Invalid Parameter"
	message[PlatRequestFrequent] = "Request too frequent"
	message[PlatInvalidCurrency] = "Invalid currency type"
	message[PlatInvalidLanguage] = "Invalid language type"

	message[UPLOAD_IMAGE_FAIL] = "Failed to upload file"
	message[EsSynchronousError] = "ES user coinLog synchronization exception"
	message[EsCurrentSumBetError] = "ES Statistics The current bet flow is abnormal"
	message[SaveCoinLogError] = "Failed to save coin Log"
	message[WithdrawalOrderError] = "withdrawal order  error"
	message[BetslipsRefundError] = "betslips refund error"
	message[WithdrawalAmountError] = "the limit per transaction is between %s-%s"
	message[WithdrawalOrdersNoProcess] = "withdrawal order is being processed"
	message[CoinRateException] = "coin rate does not exist"
	message[WithdrawalAddressCountException] = "withdrawal address article 3 the most"
	message[WithdrawalAddressNotExists] = "Withdrawal address not exists"
	message[BetslipsRefundOrderNothingnessError] = "betslips refund order nothingness error"
	message[NoBettingWasRecorded] = "No betting was recorded"
	message[RepeatedBets] = "Repeated bets"
	message[NoticeUidException] = "the uid of the in-station message cannot be empty"
	message[SaveStreamFailure] = "Save coin stream failure"

	message[SignatureVerificationError] = "sports signature verification error"
	message[PartnerError] = "sports PartnerId error"
	message[CurrencyError] = "sports currency id error"
	message[UserNotExist] = "sports user not exist"
	message[UserBlack] = "sports user black"
	message[BetTransactionIdExist] = "sports bet transaction id exist"
	message[BetOrderNumberExist] = "sports bet order number exist"
	message[BetAmountMax] = "sports bet amount max"
	message[RecordDoesNotExist] = "The record does not exist"
	message[WithdrawalAmountAbnormal] = "The withdrawal amount is abnormal"
	message[PayChannelIsClose] = "pay channel is close!"
	message[WithdrawalHasBeenInitiated] = "withdrawal order payment has been initiated"
	message[PayPlatIsClose] = "pay plat is close!"
	message[CoinInvalid] = "amount is invalid"
	message[ActivityIsContinue] = "If you participate in the current activity, the previous unfinished activity will automatically canceled. Do you want to participate in the new activity immediately?"
	message[WithdrawalIsDeposit] = "Users need to top up to initiate a withdrawal"
	message[PictureCodeException] = "Verification failure"

	message[MemberOnLine] = "Member On Line"                                                       //会员在线
	message[InvalidSecretKey] = "Invalid SecretKey"                                                //密钥无效
	message[MemberNotExist] = "Member Not Exist"                                                   //会员不存在
	message[OffLine] = "OffLine"                                                                   //会员已下线
	message[MemberNotExist] = "Member Not Exist"                                                   //会员不存在
	message[AccountNotExist] = "Account Not Exist"                                                 //帐户不存在
	message[InvalidAuthenticationKey] = "Invalid Authentication Key"                               //身份验证密钥无效
	message[InsufficientBalance] = "Insufficient player balance"                                   //玩家余额不足
	message[IpIsRestricted] = "IP address is restricted"                                           //IP地址被限制
	message[DuplicateTransaction] = "Repeated transactions"                                        //重复的交易
	message[BetNotFound] = "No such order found"                                                   //查无该注单
	message[BetNotFound] = "No information found"                                                  //查无资料
	message[InfoNotFound] = "No information found"                                                 //查无资料
	message[CannotExecute] = "Unable to execute, please try again later"                           //无法执行，请稍后再试一次
	message[BettingNotEligibleForBonus] = "The order does not meet the bonus eligibility criteria" //注单不符合 Bonus 资格
	message[PlacebetSeriesCannotBeRetried] = "Placebet series cannot be retried"                   //无法重试Placebet系列
	message[NetworkError] = "network error"                                                        //网络错误
	message[SystemMaintenance] = "System maintenance in progress"                                  //系统维修中
	message[RequestTimeout] = "Request overdue"                                                    //请求逾时
	message[SystemBusy] = "The system is busy"                                                     //系统忙碌中
	message[SystemError] = "system error"                                                          //系统错误
	message[ActiveBetCoinNoEnough] = "Insufficient Balance"                                        //余额不足
	message[BetslipsRepetitionError] = "Betslips Repetition Error"                                 //订单重复错误
	message[InvalidAmount] = "Invalid Amount"                                                      //无效金额
	message[RecordDoesNotExist] = "Record does not exist"                                          //记录不存在
	message[InsufficientPoints] = "Insufficient points"                                            //积分不足

}

func MapErrMsg(error ErrCode) string {
	if msg, ok := message[error]; ok {
		return msg
	} else {
		return "The server is missing, try again later"
	}
}

func IsCodeErr(error ErrCode) bool {
	if _, ok := message[error]; ok {
		return true
	} else {
		return false
	}
}
