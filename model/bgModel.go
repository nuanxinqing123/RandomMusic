/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  bgModel
 * @Date: 2021/8/20 17:38
 */

package model

// BgAPI 背景图片数据
type BgAPI struct {
	Status int    `json:"status"`
	Bing   BgData `json:"bing"`
}

type BgData struct {
	Url       string `json:"url"`
	Copyright string `json:"copyright"`
}
