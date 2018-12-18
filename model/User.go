package model

import "database/sql"

type User struct {
	Id                 sql.NullInt64 `orm:"id"`
	OpenId             sql.NullString `orm:"openid"`
	UnionId            sql.NullString `orm:"unionid"`
	Visitor            sql.NullString `orm:"visitor"`
	Nickname           sql.NullString `orm:"nickname"`
	Sex                sql.NullString `orm:"sex"`
	Mobile             sql.NullString `orm:"mobile"`
	Avatar             sql.NullString `orm:"avatar"`
	IsSubscribe        sql.NullString `orm:"is_subscribe"`
	SubscriptionExtend sql.NullString `orm:"subscription_extend"`
	SubscribeTime      sql.NullInt64 `orm:"subscribe_time"`
	BookCategoryIds    sql.NullString `orm:"book_category_ids"`
	OperateTime        sql.NullInt64 `orm:"operate_time"`
	IsPay              sql.NullString `orm:"is_pay"`
	KanDian            sql.NullInt64 `orm:"kandian"`
	FreeKanDian        sql.NullInt64 `orm:"free_kandian"`
	VipEndTime         sql.NullInt64 `orm:"vip_endtime"`
	RegisterIp         sql.NullString `orm:"register_ip"`
	Country            sql.NullString `orm:"country"`
	Area               sql.NullString `orm:"area"`
	Province           sql.NullString `orm:"province"`
	City               sql.NullString `orm:"city"`
	Isp                sql.NullString `orm:"isp"`
	ChannelId          sql.NullInt64 `orm:"channel_id"`
	HasFirstUnPay      sql.NullString `orm:"has_first_unpay"`
	IsFirstUnFollow    sql.NullString `orm:"is_first_unfollow"`
	PushId             sql.NullInt64 `orm:"push_id"`
	PushIdx            sql.NullInt64 `orm:"push_idx"`
	Mark               sql.NullInt64 `orm:"mark"`
	CityCode           sql.NullString `orm:"city_code"`
	AgentId            sql.NullInt64 `orm:"agent_id"`
	ReferralId         sql.NullInt64 `orm:"referral_id"`
	State              sql.NullString `orm:"state"`
	CreateTime         sql.NullInt64 `orm:"createtime"`
	UpdateTime         sql.NullInt64 `orm:"updatetime"`
}