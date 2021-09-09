/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  userModel
 * @Date: 2021/9/4 16:32
 */

package model

// User 用户注册数据
type User struct {
	Id           int
	Username     string `form:"username"`
	Password     string `form:"password"`
	Email        string `form:"email"`
	RegisterTime string
}

// HistoryRecord 历史歌曲数据
type HistoryRecord struct {
	Id        int
	Username  string
	MusicPic  string
	MusicInfo string
	MusicLink string
	EndTime   string
}

// UserEmailCode 用户验证码过期数据
type UserEmailCode struct {
	Id      int
	Email   string `form:"email"`
	Code    string
	EndTime int
}
