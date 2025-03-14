// Code generated by ent, DO NOT EDIT.

package wincoinlog

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gzorm/commons/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldID, id))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldUID, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldUsername, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCategory, v))
}

// ReferID applies equality check predicate on the "refer_id" field. It's identical to ReferIDEQ.
func ReferID(v int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldReferID, v))
}

// Coin applies equality check predicate on the "coin" field. It's identical to CoinEQ.
func Coin(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoin, v))
}

// CoinReal applies equality check predicate on the "coin_real" field. It's identical to CoinRealEQ.
func CoinReal(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoinReal, v))
}

// PlatID applies equality check predicate on the "plat_id" field. It's identical to PlatIDEQ.
func PlatID(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldPlatID, v))
}

// OutIn applies equality check predicate on the "out_in" field. It's identical to OutInEQ.
func OutIn(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldOutIn, v))
}

// GameID applies equality check predicate on the "game_id" field. It's identical to GameIDEQ.
func GameID(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldGameID, v))
}

// CoinBefore applies equality check predicate on the "coin_before" field. It's identical to CoinBeforeEQ.
func CoinBefore(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoinBefore, v))
}

// CoinAfter applies equality check predicate on the "coin_after" field. It's identical to CoinAfterEQ.
func CoinAfter(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoinAfter, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldUID, vs...))
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldUID, v))
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldUID, v))
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldUID, v))
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldUID, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldContainsFold(FieldUsername, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldCategory, v))
}

// ReferIDEQ applies the EQ predicate on the "refer_id" field.
func ReferIDEQ(v int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldReferID, v))
}

// ReferIDNEQ applies the NEQ predicate on the "refer_id" field.
func ReferIDNEQ(v int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldReferID, v))
}

// ReferIDIn applies the In predicate on the "refer_id" field.
func ReferIDIn(vs ...int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldReferID, vs...))
}

// ReferIDNotIn applies the NotIn predicate on the "refer_id" field.
func ReferIDNotIn(vs ...int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldReferID, vs...))
}

// ReferIDGT applies the GT predicate on the "refer_id" field.
func ReferIDGT(v int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldReferID, v))
}

// ReferIDGTE applies the GTE predicate on the "refer_id" field.
func ReferIDGTE(v int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldReferID, v))
}

// ReferIDLT applies the LT predicate on the "refer_id" field.
func ReferIDLT(v int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldReferID, v))
}

// ReferIDLTE applies the LTE predicate on the "refer_id" field.
func ReferIDLTE(v int) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldReferID, v))
}

// ReferIDIsNil applies the IsNil predicate on the "refer_id" field.
func ReferIDIsNil() predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIsNull(FieldReferID))
}

// ReferIDNotNil applies the NotNil predicate on the "refer_id" field.
func ReferIDNotNil() predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotNull(FieldReferID))
}

// CoinEQ applies the EQ predicate on the "coin" field.
func CoinEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoin, v))
}

// CoinNEQ applies the NEQ predicate on the "coin" field.
func CoinNEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldCoin, v))
}

// CoinIn applies the In predicate on the "coin" field.
func CoinIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldCoin, vs...))
}

// CoinNotIn applies the NotIn predicate on the "coin" field.
func CoinNotIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldCoin, vs...))
}

// CoinGT applies the GT predicate on the "coin" field.
func CoinGT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldCoin, v))
}

// CoinGTE applies the GTE predicate on the "coin" field.
func CoinGTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldCoin, v))
}

// CoinLT applies the LT predicate on the "coin" field.
func CoinLT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldCoin, v))
}

// CoinLTE applies the LTE predicate on the "coin" field.
func CoinLTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldCoin, v))
}

// CoinRealEQ applies the EQ predicate on the "coin_real" field.
func CoinRealEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoinReal, v))
}

// CoinRealNEQ applies the NEQ predicate on the "coin_real" field.
func CoinRealNEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldCoinReal, v))
}

// CoinRealIn applies the In predicate on the "coin_real" field.
func CoinRealIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldCoinReal, vs...))
}

// CoinRealNotIn applies the NotIn predicate on the "coin_real" field.
func CoinRealNotIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldCoinReal, vs...))
}

// CoinRealGT applies the GT predicate on the "coin_real" field.
func CoinRealGT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldCoinReal, v))
}

// CoinRealGTE applies the GTE predicate on the "coin_real" field.
func CoinRealGTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldCoinReal, v))
}

// CoinRealLT applies the LT predicate on the "coin_real" field.
func CoinRealLT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldCoinReal, v))
}

// CoinRealLTE applies the LTE predicate on the "coin_real" field.
func CoinRealLTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldCoinReal, v))
}

// PlatIDEQ applies the EQ predicate on the "plat_id" field.
func PlatIDEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldPlatID, v))
}

// PlatIDNEQ applies the NEQ predicate on the "plat_id" field.
func PlatIDNEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldPlatID, v))
}

// PlatIDIn applies the In predicate on the "plat_id" field.
func PlatIDIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldPlatID, vs...))
}

// PlatIDNotIn applies the NotIn predicate on the "plat_id" field.
func PlatIDNotIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldPlatID, vs...))
}

// PlatIDGT applies the GT predicate on the "plat_id" field.
func PlatIDGT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldPlatID, v))
}

// PlatIDGTE applies the GTE predicate on the "plat_id" field.
func PlatIDGTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldPlatID, v))
}

// PlatIDLT applies the LT predicate on the "plat_id" field.
func PlatIDLT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldPlatID, v))
}

// PlatIDLTE applies the LTE predicate on the "plat_id" field.
func PlatIDLTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldPlatID, v))
}

// OutInEQ applies the EQ predicate on the "out_in" field.
func OutInEQ(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldOutIn, v))
}

// OutInNEQ applies the NEQ predicate on the "out_in" field.
func OutInNEQ(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldOutIn, v))
}

// OutInIn applies the In predicate on the "out_in" field.
func OutInIn(vs ...int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldOutIn, vs...))
}

// OutInNotIn applies the NotIn predicate on the "out_in" field.
func OutInNotIn(vs ...int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldOutIn, vs...))
}

// OutInGT applies the GT predicate on the "out_in" field.
func OutInGT(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldOutIn, v))
}

// OutInGTE applies the GTE predicate on the "out_in" field.
func OutInGTE(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldOutIn, v))
}

// OutInLT applies the LT predicate on the "out_in" field.
func OutInLT(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldOutIn, v))
}

// OutInLTE applies the LTE predicate on the "out_in" field.
func OutInLTE(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldOutIn, v))
}

// GameIDEQ applies the EQ predicate on the "game_id" field.
func GameIDEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldGameID, v))
}

// GameIDNEQ applies the NEQ predicate on the "game_id" field.
func GameIDNEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldGameID, v))
}

// GameIDIn applies the In predicate on the "game_id" field.
func GameIDIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldGameID, vs...))
}

// GameIDNotIn applies the NotIn predicate on the "game_id" field.
func GameIDNotIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldGameID, vs...))
}

// GameIDGT applies the GT predicate on the "game_id" field.
func GameIDGT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldGameID, v))
}

// GameIDGTE applies the GTE predicate on the "game_id" field.
func GameIDGTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldGameID, v))
}

// GameIDLT applies the LT predicate on the "game_id" field.
func GameIDLT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldGameID, v))
}

// GameIDLTE applies the LTE predicate on the "game_id" field.
func GameIDLTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldGameID, v))
}

// CoinBeforeEQ applies the EQ predicate on the "coin_before" field.
func CoinBeforeEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoinBefore, v))
}

// CoinBeforeNEQ applies the NEQ predicate on the "coin_before" field.
func CoinBeforeNEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldCoinBefore, v))
}

// CoinBeforeIn applies the In predicate on the "coin_before" field.
func CoinBeforeIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldCoinBefore, vs...))
}

// CoinBeforeNotIn applies the NotIn predicate on the "coin_before" field.
func CoinBeforeNotIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldCoinBefore, vs...))
}

// CoinBeforeGT applies the GT predicate on the "coin_before" field.
func CoinBeforeGT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldCoinBefore, v))
}

// CoinBeforeGTE applies the GTE predicate on the "coin_before" field.
func CoinBeforeGTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldCoinBefore, v))
}

// CoinBeforeLT applies the LT predicate on the "coin_before" field.
func CoinBeforeLT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldCoinBefore, v))
}

// CoinBeforeLTE applies the LTE predicate on the "coin_before" field.
func CoinBeforeLTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldCoinBefore, v))
}

// CoinAfterEQ applies the EQ predicate on the "coin_after" field.
func CoinAfterEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCoinAfter, v))
}

// CoinAfterNEQ applies the NEQ predicate on the "coin_after" field.
func CoinAfterNEQ(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldCoinAfter, v))
}

// CoinAfterIn applies the In predicate on the "coin_after" field.
func CoinAfterIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldCoinAfter, vs...))
}

// CoinAfterNotIn applies the NotIn predicate on the "coin_after" field.
func CoinAfterNotIn(vs ...float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldCoinAfter, vs...))
}

// CoinAfterGT applies the GT predicate on the "coin_after" field.
func CoinAfterGT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldCoinAfter, v))
}

// CoinAfterGTE applies the GTE predicate on the "coin_after" field.
func CoinAfterGTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldCoinAfter, v))
}

// CoinAfterLT applies the LT predicate on the "coin_after" field.
func CoinAfterLT(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldCoinAfter, v))
}

// CoinAfterLTE applies the LTE predicate on the "coin_after" field.
func CoinAfterLTE(v float64) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldCoinAfter, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int32) predicate.WinCoinLog {
	return predicate.WinCoinLog(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WinCoinLog) predicate.WinCoinLog {
	return predicate.WinCoinLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WinCoinLog) predicate.WinCoinLog {
	return predicate.WinCoinLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WinCoinLog) predicate.WinCoinLog {
	return predicate.WinCoinLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
