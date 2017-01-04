package admin

import (
	"fox/model"
	"fox/util/db"
	"fmt"
)

type Attachment struct {
	*model.Attachment
}

func NewAttachmentSercice() *Attachment {
	return new(Attachment)
}
func (c *Attachment)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewAttachment()
	data, err := mode.GetAll(q, fields, orderBy, page, 20)
	if err != nil {
		return nil, err
	}
	return data, nil
}
//更新添加的数据
//@type_id 模块
//@id      id
//@aid     管理员ID
func (c *Attachment)UpdateByTypeIdId(type_id, id, aid int) (bool, error) {
	maps := make(map[string]interface{})
	maps["type_id"] = type_id
	maps["id"] = id
	maps["aid"] = aid
	num, err := db.Filter(maps).Update(c)
	if err != nil {
		return false, err
	}
	fmt.Println("更新条数", num)
	return true, nil
}