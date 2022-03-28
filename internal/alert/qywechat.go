package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golo3/global"
	"golo3/model"
	"golo3/pkg/setting"
	"log"
	"net/http"
)

/*
负责 企业微信方式的通知
*/

type QywechatInfo struct {
	CorpID  string
	AgentID int32
	Secret  string
	MsgType string
	ToUser  string
}

type QywechatAlertInfo struct {
	Title   string
	Content string
}

//bindQywechatInfo 绑定企业微信通知各项参数
func bindQywechatInfo(q *setting.QywechatSettings) *QywechatInfo {
	return &QywechatInfo{
		CorpID:  q.CorpID,
		Secret:  q.Secret,
		ToUser:  q.ToUser,
		MsgType: q.MsgType,
		AgentID: q.AgentID,
	}
}

//getAccessTokenUrl 拼接出请求access_token的url
func getAccessTokenUrl(info *QywechatInfo) string {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", info.CorpID, info.Secret)
	return url
}

//getAccessToken  获取access_token用来发送消息
func getAccessToken() (interface{}, error) {
	qywechatinfo := bindQywechatInfo(global.QywechatSetting)
	url := getAccessTokenUrl(qywechatinfo)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	if data["errcode"].(float64) != 0 {
		log.Println(data["errcode"])
		return "", err
	}
	return data["access_token"], nil
}

//getQywechatAlertInfo  企业微信通知消息具体内容
func getQywechatAlertInfo(l *model.LogInfo) *QywechatAlertInfo {
	return &QywechatAlertInfo{
		Title:   fmt.Sprintf("%s 环境 %s 的日志报警", l.Env, l.ServiceName),
		Content: fmt.Sprintf("Time: %s \n ERROR INFO:\n %s\n", l.Timestamp, l.LogInfo),
	}
}

//getPostUrl 获取Post请求的url
func getPostUrl() (string, error) {
	accessToken, err := getAccessToken()
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", accessToken)
	return url, nil

}

//getPostBody 绑定Post请求消息体
func getPostBody(info *QywechatInfo, alertinfo *QywechatAlertInfo) *map[string]interface{} {
	switch info.MsgType {
	case "textcard":
		return &map[string]interface{}{
			"touser":  info.ToUser,
			"msgtype": info.MsgType,
			"agentid": info.AgentID,
			"textcard": map[string]interface{}{
				"title":       alertinfo.Title,
				"description": alertinfo.Content,
				"url":         "www.baidu.com",
				"btntext":     "跳转到百度",
			},
		}
	case "markdown":
		return &map[string]interface{}{
			"touser":  info.ToUser,
			"msgtype": info.MsgType,
			"agentid": info.AgentID,
			"markdown": map[string]interface{}{
				"content": "您的会议室已经预定，稍后会同步到`邮箱`\n\t\t\t\t>**事项详情**\n\t\t\t\t>事　项：<font color=\\\"info\\\">开会</font> \n\t\t\t\t>组织者：@miglioguan\n\t\t\t\t>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang\n\t\t\t\t>\n\t\t\t\t>会议室：<font color=\\\"info\\\">广州TIT 1楼 301</font> \n\t\t\t\t>日　期：<font color=\\\"warning\\\">2018年5月18日</font> \n\t\t\t\t>时　间：<font color=\\\"comment\\\">上午9:00-11:00</font> \n\t\t\t\t>\n\t\t\t\t>请准时参加会议。\n\t\t\t\t>\n\t\t\t\t>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com),",
			},
		}
	}
	return nil
}

func QywechatAlert(l *model.LogInfo) (map[string]interface{}, error) {

	//绑定企业微信参数设置
	qywechatInfo := bindQywechatInfo(global.QywechatSetting)
	// 绑定报警信息
	qywechatAlertInfo := getQywechatAlertInfo(l)
	// 获取Post请求url
	url, err := getPostUrl()
	if err != nil {
		return nil, err
	}
	// 绑定Post请求的消息体
	body := getPostBody(qywechatInfo, qywechatAlertInfo)
	res, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(res))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data2 map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data2)
	if err != nil {
		return nil, err
	}
	return data2, nil
}
