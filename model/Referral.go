package model

import (
	"database/sql"
	"mq/service"
)

type Referral struct {
	Type sql.NullInt64 `orm:"type"`
	Follow sql.NullInt64 `orm:"follow"`
	Net_follow_num sql.NullInt64 `orm:"net_follow_num"`
	Money sql.NullFloat64 `orm:"money"`
	Createtime sql.NullInt64 `orm:"createtime"`
	Short_id sql.NullInt64 `orm:"short_id"`
	Utime sql.NullString `orm:"utime"`
	Manage_title_id sql.NullInt64 `orm:"manage_title_id"`
	Manage_preview_id sql.NullString `orm:"manage_preview_id"`
	Guide_sex sql.NullString `orm:"guide_sex"`
	Cost sql.NullFloat64 `orm:"cost"`
	Short_url_qq sql.NullString `orm:"short_url_qq"`
	Incr_num sql.NullInt64 `orm:"incr_num"`
	Guide_idx sql.NullInt64 `orm:"guide_idx"`
	Id sql.NullInt64 `orm:"id"`
	Book_id sql.NullInt64 `orm:"book_id"`
	Name sql.NullString `orm:"name"`
	Short_url_weibo sql.NullString `orm:"short_url_weibo"`
	Push sql.NullString `orm:"push"`
	Chapter_id sql.NullInt64 `orm:"chapter_id"`
	Guide_chapter_idx sql.NullInt64 `orm:"guide_chapter_idx"`
	Wx_type sql.NullString `orm:"wx_type"`
	Guide_follow_num sql.NullInt64 `orm:"guide_follow_num"`
	State sql.NullString `orm:"state"`
	Guide_title sql.NullString `orm:"guide_title"`
	Chapter_name sql.NullString `orm:"chapter_name"`
	Admin_id sql.NullInt64 `orm:"admin_id"`
	Source_url sql.NullString `orm:"source_url"`
	Unfollow_num sql.NullInt64 `orm:"unfollow_num"`
	Manage_template_id sql.NullString `orm:"manage_template_id"`
	Updatetime sql.NullInt64 `orm:"updatetime"`
	Guide_image sql.NullString `orm:"guide_image"`
	Manage_cover_id sql.NullInt64 `orm:"manage_cover_id"`
	Chapter_idx sql.NullInt64 `orm:"chapter_idx"`
	Uv sql.NullInt64 `orm:"uv"`
}

var L = service.LogService{}

func GetReferralIds() []sql.NullInt64 {
	session := service.GetDb()
	var referral []Referral
	var list []sql.NullInt64
	err1 := session.Table(&referral).Offset(0).Limit(100).Order("id desc").Select()
	L.FailOnError(err1, "查询失败" + session.LastSql)
	for _,v := range referral {
		list = append(list, v.Id)
	}

	return list
}