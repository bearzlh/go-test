package model

import (
	"database/sql"
	"mq/service"
)

type Referral struct {
	Type             sql.NullInt64   `orm:"type"`
	Follow           sql.NullInt64   `orm:"follow"`
	NetFollowNum     sql.NullInt64   `orm:"net_follow_num"`
	Money            sql.NullFloat64 `orm:"money"`
	CreateTime       sql.NullInt64   `orm:"createtime"`
	ShortId          sql.NullInt64   `orm:"short_id"`
	UTime            sql.NullString  `orm:"utime"`
	ManageTitleId    sql.NullInt64   `orm:"manage_title_id"`
	ManagePreviewId  sql.NullString  `orm:"manage_preview_id"`
	GuideSex         sql.NullString  `orm:"guide_sex"`
	Cost             sql.NullFloat64 `orm:"cost"`
	ShortUrlQq       sql.NullString  `orm:"short_url_qq"`
	IncrNum          sql.NullInt64   `orm:"incr_num"`
	GuideIdx         sql.NullInt64   `orm:"guide_idx"`
	Id               sql.NullInt64   `orm:"id"`
	BookId           sql.NullInt64   `orm:"book_id"`
	Name             sql.NullString  `orm:"name"`
	ShortUrlWeiBo    sql.NullString  `orm:"short_url_weibo"`
	Push             sql.NullString  `orm:"push"`
	ChapterId        sql.NullInt64   `orm:"chapter_id"`
	GuideChapterIdx  sql.NullInt64   `orm:"guide_chapter_idx"`
	WxType           sql.NullString  `orm:"wx_type"`
	GuideFollowNum   sql.NullInt64   `orm:"guide_follow_num"`
	State            sql.NullString  `orm:"state"`
	GuideTitle       sql.NullString  `orm:"guide_title"`
	ChapterName      sql.NullString  `orm:"chapter_name"`
	AdminId          sql.NullInt64   `orm:"admin_id"`
	SourceUrl        sql.NullString  `orm:"source_url"`
	UnFollowNum      sql.NullInt64   `orm:"unfollow_num"`
	ManageTemplateId sql.NullString  `orm:"manage_template_id"`
	UpdateTime       sql.NullInt64   `orm:"updatetime"`
	GuideImage       sql.NullString  `orm:"guide_image"`
	ManageCoverId    sql.NullInt64   `orm:"manage_cover_id"`
	ChapterIdx       sql.NullInt64   `orm:"chapter_idx"`
	Uv               sql.NullInt64   `orm:"uv"`
}

var L = service.LogService{}

func GetReferralIds() []sql.NullInt64 {
	session := service.GetDefaultDb()
	var referral []Referral
	var list []sql.NullInt64
	err1 := session.Table(&referral).Offset(0).Limit(100).Order("id desc").Select()
	L.FailOnError(err1, "查询失败"+session.LastSql)
	for _, v := range referral {
		list = append(list, v.Id)
	}

	return list
}
