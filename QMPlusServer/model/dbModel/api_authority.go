package dbModel

import "main/init/qmsql"

type ApiAuthority struct {
	AuthorityId string `json:"-"`
	Path  string       `json:"_"`
}


//创建角色api关联关系
func (a *ApiAuthority)SetAuthAndPath(authId string,apis []Api)(err error){
	err = qmsql.DEFAULTDB.Where("authority_id = ?",authId).Delete(&ApiAuthority{}).Error
	for _,v := range apis{
		err = qmsql.DEFAULTDB.Create(&ApiAuthority{AuthorityId:authId,Path:v.Path}).Error
		if (err!=nil){
			return err
		}
	}
	return nil
}