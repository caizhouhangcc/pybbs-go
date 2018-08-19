package models

import (
	"pybbs-go/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// tags
type Tag struct {
	Id        int      `orm:"pk;auto"`
	Name      string   `orm:"unique"`
	alphabeta string   `orm:"index"` // 为了未来的排序
	Topics    []*Topic `orm:"reverse(many)"`
}

// SaveTag save tag
func SaveTag(tag *Tag) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(tag)
	return id
}

func FindTagById(id int) *Tag {
	o := orm.NewOrm()
	var tag Tag
	o.QueryTable(tag).RelatedSel().Filter("Id", id).One(&tag)
	return &tag
}

// FindTagByName to find tag by its name
func FindTagByName(name string) *Tag {
	o := orm.NewOrm()
	var tag Tag
	o.QueryTable(tag).RelatedSel().Filter("Name", name).One(&tag)
	return &tag
}

// QueryTagsByName to query tag by its name
func QueryTagsByName(q string) utils.Page {
	p, size := 0, 30
	o := orm.NewOrm()
	var tag Tag
	var list []Tag
	qs := o.QueryTable(tag)
	qs.RelatedSel().OrderBy("name").Limit(size).Offset((p - 1) * size).All(&list)
	return utils.PageUtil(-1, p, size, list)
}

func PageTag(p int, size int) utils.Page {
	o := orm.NewOrm()
	var tag Tag
	var list []Tag
	qs := o.QueryTable(tag)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("name").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}
