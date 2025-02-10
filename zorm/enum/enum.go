package enum

type BetEnum int

const (
	ROUND_ID BetEnum = iota + 1
	TRANSACTION_ID
)

type OutInEnum uint8

const (
	OUT OutInEnum = iota
	IN
)

type FlowCategoryTypeEnum int16

const (
	BANK_SAVINGS FlowCategoryTypeEnum = iota + 1
	WITHDRAWAL
	BET
	DRAW
	REBATE
	BROKERAGE
	AWARD
	SYSTEM_RECONCILIATION
	REFUND
	COMM_COIN_TRANSFER
	WITHDRAWAL_REFUND FlowCategoryTypeEnum = iota + 2
	TRANSFER_OUT
	TRANSFER_IN
	BALANCE_GAME
	GAME_ADJUSTMENT_BALANCE
)

// Xb_STATUS状态
const (
	STATUS_WAITING              = iota + 1 // 待开奖
	STATUS_FINISH                          // 已开奖
	STATUS_REFUND                          // 退款
	STATUS_WAITING_UNDETERMINED            // 已下注待开彩(未确定下注)
)

type WalletCategory int8

const (
	WalletCategoryPay WalletCategory = iota + 1
	WalletCategoryGame
	WalletCategoryActivity
	WalletCategoryCommission
)

type WalletCurrency int8

const (
	WalletCurrencyPhp WalletCurrency = iota + 1
)

type COIN_STATUS_ENUM int64

const (
	PROCESSING COIN_STATUS_ENUM = iota
	SUCCESS
	FAILED
)

type PLAT int8

const (
	BNG PLAT = iota + 1
	ONE_GAME
	PG
	EZUGI
	HABA
	JDB
	FB
	CQ9
	JiLi
	DS88
	BTI
	WM
	ELBET
	HOT_SLOTS
	EVO
	TPG
	TBS
	Booming
	ENDING
	AVIATRIX
	PLAY_STAR
	CALETA
	SABA
	PB
	RTG
	FUNKY
	AG
	PP
	TG
	ELECTRON
	TABLES

	//HABANERO1
)

func (p PLAT) String() string {
	switch p {
	case BNG:
		return "BNG"
	case ONE_GAME:
		return "ONE_GAME"
	case PG:
		return "PG"
	case EZUGI:
		return "EZUGI"
	case HABA:
		return "HABA"
	case JDB:
		return "JDB"
	case FB:
		return "FB"
	case CQ9:
		return "CQ9"
	case JiLi:
		return "JiLi"
	case DS88:
		return "DS88"
	case BTI:
		return "BTI"
	case WM:
		return "WM"
	case ELBET:
		return "ELBET"
	case HOT_SLOTS:
		return "HOT_SLOTS"
	case EVO:
		return "EVO"
	case TPG:
		return "TPG"
	case TBS:
		return "TBS"
	case Booming:
		return "Booming"
	case ENDING:
		return "ENDING"
	case AVIATRIX:
		return "AVIATRIX"
	case PLAY_STAR:
		return "PLAY_STAR"
	case CALETA:
		return "CALETA"
	case SABA:
		return "Filbet_Sports"
	case PB:
		return "PB"
	case RTG:
		return "RTG"
		//case HABANERO:
		// return "HABANERO"
	case FUNKY:
		return "FUNKY"
	case AG:
		return "AG"
	case PP:
		return "PP"
	case TG:
		return "v99_100001"
	case TABLES:
		return "v99_100003"
	case ELECTRON:
		return "v99_100002"
	}
	return ""
}

type TICKESTATUS string

type STATUS int8

const (
	STATUSRUNNING STATUS = iota + 1
	STATUSWON
	STATUSLOSE
	STATUSDRAW
	STATUSVOID
	STATUSREFUND
	STATUSREJECT
)

func (s STATUS) String() string {
	switch s {
	case STATUSRUNNING:
		return "running"
	case STATUSWON:
		return "won"
	case STATUSLOSE:
		return "lose"
	case STATUSDRAW:
		return "draw"
	case STATUSVOID:
		return "void"
	case STATUSREFUND:
		return "refund"
	case STATUSREJECT:
		return "reject"
	}
	return ""
}

type GAME int8

const (
	JDB_SLOT GAME = iota + 1
	BNG_SLOT
	PG_SLOT
	ONE_GAME_SLOT
	CQ9_SLOT
	EZUGI_LIVE
	JDB_FISH
	DS88_ANIMAL
	BOOMING_SLOTS
	AVIATRIX_SLOT
)

func (g GAME) String() string {
	switch g {
	case JDB_SLOT:
		return "JDB_SLOT"
	case BNG_SLOT:
		return "BNG_SLOT"
	case PG_SLOT:
		return "PG_SLOT"
	case ONE_GAME_SLOT:
		return "ONE_GAME_SLOT"
	case CQ9_SLOT:
		return "CQ9_SLOT"
	case EZUGI_LIVE:
		return "CQ9_SLOT"
	case JDB_FISH:
		return "JDB_FISH"
	case DS88_ANIMAL:
		return "DS88_ANIMAL"
	case BOOMING_SLOTS:
		return "BOOMING_SLOTS"
	case AVIATRIX_SLOT:
		return "AVIATRIX_SLOT"
	}

	return ""
}

type Group int8

const (
	SPORTS Group = iota + 1
	SLOT
	CASINO
	FINISH
	CARDS
	E_SPORT
	LOTT
	ANIMAL
	ARCADE
)

func (g Group) String() string {
	switch g {
	case SPORTS:
		return "体育"
	case SLOT:
		return "电子"
	case CASINO:
		return "真人"
	case FINISH:
		return "捕鱼"
	case CARDS:
		return "棋牌"
	case E_SPORT:
		return "电竞"
	case LOTT:
		return "彩票"
	case ANIMAL:
		return "动物"
	case ARCADE:
		return "街机"
	}
	return ""
}
