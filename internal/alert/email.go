package alert

/*
负责 email方式的通知
*/

type EmailAlertInfo struct {
	Subject string
	Content string
	SendTo  string
}
