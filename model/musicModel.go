/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  musicModel
 * @Date: 2021/8/20 19:31
 */

package model

type Data struct {
	Name string `json:"name"`
	Auther string `json:"auther"`
	Mp3url	string	`json:"mp3url"`
	Picurl string `json:"imgurl"`
}

type MusicData struct {
	Info Data `json:"info"`
}