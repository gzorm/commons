package payment

type ResultCode string

// resultCodeId resultStatus resultCodeRemarks
// 00000000 S SUCCESS It means Gcash has received the transaction successfully
// 00000002 F PARAM_MISSING One or more mandatory parameters is/are missing
// 00000004 F PARAM_ILLEGAL Illegal parameters.For example,non-numeric input,invalid date,etc.
// 00000007 F INVALID_SIGNATURE Signature is invalid
// 00000013 F NO_INTERFACE_DEF API is not defined
// 00000017 F FUNCTION_NOT_MATCH Function parameter does not match API
// 00000023 F ACCESSTOKEN_VALIDATE_FAILED Access Token Validate failed
// 00000024 F CLIENT_VALIDATE_FAILED Client Authentication failed
// 00000900 U SYSTEM_ERROR System Error.Merchant can initiate the transaction again with the same parameters
// 12003001 F MERCHANT_NOT_EXIST merchant account does not exist
// 12003040 F SUBMERCHANT_NOT_REGISTERED sub merchant not register ed under the merchant
// 12005002 F ORDER_NOT_EXISTS Order does not exist
// 12005003 F ORDER_STATUS_INVALID Need to verify payment status(should be SUCCESS)
// 12005004 F ORDER_IS_FROZEN order is frozen(but currently not implemented)
// 12005005 F NOAUTH No auth
// 12005007 F ACCOUNT_STATUS_ABNORMAL Account status abnormal GCashOne-TimePayment Web pay v1.2.0 IntegrationGuideCopyright©2021GlobeFintechInnovations,Inc.AllRightsReserved.50
// 12005008 F ORDER_IS_CANCELED Order is canceled
// 12005011 F PAYMENT_REQUEST_HAS_RISK Payment request has risk
// 12005014 F WITHOUT_AVAILABLE_PAY_METHOD Without available pay method
// 12005017 F CANCEL_NOT_ALLOWED cancel not allowed in merchant contract
// 12005018 F CANCEL_EXPIRED Allowed duration for cancellation has expired.
// 12005019 F REFUND_TRANSACTION_EXIST cancel service cannot push through due to are fund transaction exist
// 12005020 F DISPUTE_TRANSACTION_EXIST No usage
// 12005021 F ERCHANT_ACCOUNT_BALANCE_NOT_ENOUGH Merchant account does not have enough balance
// 12005022 F MERCHANT_ACCOUNT_ABNORMAL Merchant account
// 12005023 F USER_BOUND_ASSET_NOT_EXIST User bound asset not exist
// 12005024 F CHANNEL_STATUS_NOT_ENABLED Channel status not enabled<Note:Not in use currently>
// 12005025 F CHANNEL_OVER_LIMIT Channel over limit<Note:Not in use currently>
// 12005027 F BALANCE_EXCEED_LIMIT user limit exceeded
// 12005100 F ORDER_IS_CLOSED Order is closed
// 12005102 F CAPTURE_NOT_EXISTS no code/usage
// 12005103 F CURRENCY_NOT_CORRECT Check the set currency value.It should be in PHP
// 12005104 F AMOUNT_EXCEEDS_LIMIT Transaction amount exceeds limit
// 12005105 F AGREEMENT_REFUND_NOT_ALLOWED refund not allowed in contract
// 12005106 F AGREEMENT_MULTI_REFUND_NOT_ALLOWED multi refund not allowed in contract
// 12005107 F AGREEMENT_SPECIFIED_ACCOUNT_NOT_EXIST request source account is not in contract refund source account list GCash One-Time Payment Web pay v1.2.0IntegrationGuideCopyright©2021GlobeFintechInnovations,Inc.AllRightsReserved.51
// 12005108 F AGREEMENT_REFUND_TIME_EXCEEDS_LIMIT refund service used after refund deadline in contract
// 12005110 F USER_STATUS_ABNORMAL User status abnormal
// 12005112 F REPEAT_REQ_INCONSISTENT Repeated submission and requests are in consistent
// 12005115 F BALANCE_NOT_ENOUGH Balance is not enough to finish the transaction
// 12005124 F USER_TYPE_ILLEGAL user type not found in cust center
// 12005200 F MERCHANT_STATUS_ABNORMAL Merchant status abnormal
// 12005436 U REFUND_IN_PROCESS Need to wait,refund is in process.Recommended to confirm with GCash support teambyfilingaticketfortheconfirmation of refund.
// 12005437 F CANCEL_IN_PROCESS Need to wait,cancel is in process.No cancel query.RecommendedtoconfirmwithGCashsupportteambyfilingaticketfortheconfirmationoftransactioncancellation.
// 12125003 F TIME_INVALID Time is invalid
const (
	SUCCESS           ResultCode = "00000000"
	PARAM_MISSING     ResultCode = "00000002"
	PARAM_ILLEGAL     ResultCode = "00000004"
	INVALID_SIGNATURE ResultCode = "00000007"
	SYSTEM_ERROR      ResultCode = "00000900"
	ORDER_NOT_EXISTS  ResultCode = "12005002"
	ORDER_IS_CLOSED   ResultCode = "12005100"
	REFUND_IN_PROCESS ResultCode = "12005436"
)

//12005100 F ORDER_IS_CLOSED Order is closed

func (r *ResultCode) String() string {
	switch *r {
	case SUCCESS:
		return "SUCCESS"
	case PARAM_ILLEGAL:
		return "PARAM_ILLEGAL"
	case INVALID_SIGNATURE:
		return "INVALID_SIGNATURE"
	case SYSTEM_ERROR:
		return "SYSTEM_ERROR"
	case ORDER_NOT_EXISTS:
		return "ORDER_NOT_EXISTS"
	case ORDER_IS_CLOSED:
		return "ORDER_IS_CLOSED"
	default:
		return ""
	}
}

func (r *ResultCode) Status() string {
	switch *r {
	case SUCCESS:
		return "S"
	case SYSTEM_ERROR, REFUND_IN_PROCESS:
		return "U"
	default:
		return "F"
	}
}

func (r *ResultCode) Msg() string {
	switch *r {
	case SUCCESS:
		return "It means Gcash has received the transaction successfully"
	case SYSTEM_ERROR:
		return "System Error.Merchant can initiate the transaction again with the same parameters"
	case PARAM_ILLEGAL:
		return "Illegal parameters.For example,non-numeric input,invalid date,etc."
	case ORDER_NOT_EXISTS:
		return "Order does not exist"
	case INVALID_SIGNATURE:
		return "Signature is invalid"
	case ORDER_IS_CLOSED:
		return "Order is closed"
	default:
		return ""
	}
}
