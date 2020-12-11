package main

import (
	"fmt"
)

type Iuser struct {
	Name string
	Age int
	WeaponHeight int
}

// 选择最终的用户
type IuserProvide interface {
	Finduser(iuserspecific Iuserspecific) []Iuser
}

// 区别用户
type Iuserspecific interface {
	IsSatisfiedBy(iuser Iuser) bool
}

//
type Icommanstatic interface {
	And(iuserspecific Iuserspecific) Iandshixian
}

// 实现区别用户这个接口 姓名
type Iuserspecificbyname struct {
	Name string
}

func (i *Iuserspecificbyname) IsSatisfiedBy(iuser Iuser) bool {
	return i.Name == iuser.Name
}

// 实现区别用户这个接口 年龄
type Iuserspecificbyage struct {
	Age int
}

func (i *Iuserspecificbyage) IsSatisfiedBy(iuser Iuser) bool {
	return i.Age < iuser.Age
}

// 实现区别用户这个接口 wepon
type Iuserspecificbywh struct {
	WeaponHeight int
}

func (i *Iuserspecificbywh) IsSatisfiedBy(iuser Iuser) bool {
	return i.WeaponHeight < iuser.WeaponHeight
}

// 实现用户筛选
type UserProvider struct {
	userlist []Iuser
}

func (u *UserProvider)Finduser(iuserspecific Iuserspecific) []Iuser{
	var result []Iuser
	for _, v := range u.userlist {
		if iuserspecific.IsSatisfiedBy(v) {
			result = append(result, v)
		}

	}
	return result
}

// 实现and
type Iandshixian struct {
	left Iuserspecific
	right []Iuserspecific
}

func (i Iandshixian) And(iuserspecific Iuserspecific)  Iandshixian{
	i.right = append(i.right, iuserspecific)
	return i
}

func(i Iandshixian) End()  Iuserspecific{
	// i.right = append(i.right, iuserspecific)
	return &i
}

func (i *Iandshixian) IsSatisfiedBy(iuser Iuser) bool{
	var result bool = true
	for _, v := range i.right {
		result = v.IsSatisfiedBy(iuser) && result
	}
	return result && i.left.IsSatisfiedBy(iuser)
}









func main() {
	var userlist = []Iuser{
		Iuser{Name: "zhangfei", Age: 38, WeaponHeight: 15},
		Iuser{Name: "zhangfei", Age: 17, WeaponHeight: 15},
		Iuser{Name: "guanyu", Age: 39, WeaponHeight: 20},
		Iuser{Name: "zhaoyun", Age: 20, WeaponHeight: 18},
		Iuser{Name: "liubei", Age: 40, WeaponHeight: 10},
		Iuser{Name: "machao", Age: 30, WeaponHeight: 21},
		Iuser{Name: "huangzhong", Age: 45, WeaponHeight: 13},
	}
	var isuserp IuserProvide= &UserProvider{userlist: userlist}
	// var result []Iuser
	var ilogicname = Iuserspecificbyname{Name: "zhangfei"}
	var ilogicage = Iuserspecificbyage{Age: 20}
	var ilogicwh = Iuserspecificbywh{WeaponHeight: 13}
	userlist = isuserp.Finduser(&ilogicage)
	userlist = isuserp.Finduser(&ilogicname)
	var iandshixan = Iandshixian{left: &ilogicage}
	userlist = isuserp.Finduser(iandshixan.And(&ilogicname).And(&ilogicwh).End())

	//var andorother = Andorother{left: &ilogicname ,right: &ilogicage}
	//for _, v := range userlist {

	//}
	fmt.Println(userlist)
	//fmt.Println(ilogicname)
	//fmt.Println(ilogicage)



}
