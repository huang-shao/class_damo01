package models

type User struct {
	Id int`json:"id"`
	Nick string`json:"nick"`
	Password string`json:"password"`
	Birthday string`json:"birthday"`
	Hobby string`json:"hobby"`

}
