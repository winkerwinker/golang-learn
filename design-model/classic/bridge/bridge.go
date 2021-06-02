package bridge

//桥接模式

// part 2
type ImsgSender interface {
	Send(msg string) error
}
type EmailMsgSender struct {
	emails []string
}

func (e EmailMsgSender) Send(msg string) error {
	return nil
	//panic("implement me")
}

// NewEmailMsgSender NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}


// part 1
type INotify interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
	sender ImsgSender
}

// NewErrorNotification NewErrorNotification
func NewErrorNotification(sender ImsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}


// Notify 发送通知
func (n *ErrorNotification) Notify(msg string) error {
	// 自己不做让别人来做
	//  todo 与适配器的区别是不会在过程中修改参数，
	//  todo 与代理模式的区别是不会在上下添加操作
	return n.sender.Send(msg)
}