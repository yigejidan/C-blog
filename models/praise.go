package models

import (
	"liteblog/syserror"
	"github.com/jinzhu/gorm"
)


type PraiseLog struct {
	Model
	Key    string 
	UserID int    
	Table   string 
	Flag   bool
}

type TempPraise struct {
	Praise int
}

func UpdatePraise(table,key string,userid int)(pcnt int,err error){
	//需要判断后台是否已经点赞过，如果点赞过就返回点赞过得错误
	//查询点赞表，看是否有记录
	d := db.Begin()
	defer func(){
		if err != nil {
			d.Rollback()
		}else {
			d.Commit()
		}
	}()
	var p PraiseLog
	err = d.Model(&PraiseLog{}).Where("`key` = ? and `table` = ? and user_id = ?",key,table,userid).Take(&p).Error
	if err == gorm.ErrRecordNotFound{
		p = PraiseLog{
			Key : key,
			Table : table,
			UserID : userid,
			Flag : false,
		}
	}else if err != nil {
		return 0,err
	}
	if p.Flag {
		return 0,syserror.HasPraiseError{}
	}
	p.Flag = true
	if err = d.Save(&p).Error;err != nil {
		return 0,err
	}
	var ppp TempPraise
	err = d.Table(table).Where("key = ?", key).Select("praise").Scan(&ppp).Error
	if err != nil {
		return 0,err
	}
	pcnt = ppp.Praise + 1
	if err = d.Table(table).Where("key = ?", key).Update("praise",pcnt).Error;err != nil{
		return 0,err	
	}

	return pcnt,nil
}