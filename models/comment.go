package models

import "github.com/astaxie/beego/orm"

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Stamp       int
	Status      int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId, offset, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	//求总数
	num, _ := o.Raw("SELECT id FROM comment WHERE status=1 AND episodes_id=?", episodesId).QueryRows(&comments)
	_, err := o.Raw("SELECT id,content,add_time,user_id,stamp,praise_count FROM comment WHERE status=1 AND episodes_id=? ORDER BY add_time DESC LIMIT ?,?", episodesId, offset, limit).QueryRows(&comments)
	return num, comments, err
}
