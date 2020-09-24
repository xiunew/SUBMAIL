package main

import (
	"SUBMAIL/submail"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	mailconfig := make(map[string]string)
	mailconfig["appid"] = "your_mail_app_id"
	mailconfig["appkey"] = "your_mail_app_key"
	mailconfig["signtype"] = "md5"
	messageconfig := make(map[string]string)
	messageconfig["appid"] = "your_message_app_id"
	messageconfig["appkey"] = "your_message_app_key"
	messageconfig["signtype"] = "md5"

	//mail send
	mailsend := submail.CreateMailSend()
	submail.MailSendAddTo(mailsend, "leo@submail.cn", "leo")
	submail.MailSendAddCc(mailsend, "mailer@submail.cn", "")
	submail.MailSendAddBcc(mailsend, "leo@drinkfans.com", "")
	submail.MailSendSetSender(mailsend, "no-reply@submail.cn", "SUBMAIL")
	submail.MailSendSetReply(mailsend, "service@submail.cn")
	submail.MailSendSetText(mailsend, "test SDK text")
	submail.MailSendSetHtml(mailsend, "test sdk html")
	submail.MailSendSetSubject(mailsend, "test SDK")
	fmt.Println("MailSend %s", submail.MailSendRun(submail.MailSendBuildRequest(mailsend), mailconfig))

	//mail xsend
	mailxsend := submail.CreateMailXSend()
	submail.MailXSendAddTo(mailxsend, "leo@submail.cn", "leo")
	submail.MailXSendSetProject(mailxsend, "uigGk1")
	submail.MailXSendAddVar(mailxsend, "name", "leo")
	submail.MailXSendAddVar(mailxsend, "age", "32")
	submail.MailXSendAddLink(mailxsend, "developer", "http://submail.cn/chs/developer")
	submail.MailXSendAddLink(mailxsend, "store", "http://submail.cn/chs/store")
	submail.MailXSendAddHeader(mailxsend, "X-Accept", "zh-cn")
	submail.MailXSendAddHeader(mailxsend, "X-Mailer", "leo App")
	fmt.Println("MailXSend ", submail.MailXSendRun(submail.MailXSendBuildRequest(mailxsend), mailconfig))

	//messagexsend
	messagexsend := submail.CreateMessageXSend()
	submail.MessageXSendAddTo(messagexsend, "138xxxxxxxx")
	submail.MessageXSendSetProject(messagexsend, "kZ9Ky3")
	submail.MessageXSendAddVar(messagexsend, "code", "198276")
	fmt.Println("MessageXSend ", submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig))

	//AddressBookMail Subscribe and UnSubscribe
	addressbookmail := submail.CreateAddressBookMail()
	submail.AddressBookMailSetAddress(addressbookmail, "leo@apple.cn", "leo")
	fmt.Println("AddressBookMail Subscribe ", submail.MailSubscribeRun(submail.AddressBookMailBuildRequest(addressbookmail), mailconfig))
	fmt.Println("AddressBookMail UnSubscribe ", submail.MailUnSubscribeRun(submail.AddressBookMailBuildRequest(addressbookmail), mailconfig))

	//AddressBookMessage Subscribe and UnSubscribe
	addressbookmessage := submail.CreateAddressBookMessage()
	submail.AddressBookMessageSetAddress(addressbookmessage, "138xxxxxxxx", "")
	fmt.Println("AddressBookMessage Subscribe ", submail.MessageSubscribeRun(submail.AddressBookMessageBuildRequest(addressbookmessage), messageconfig))
	fmt.Println("AddressBookMessage UnSubscribe ", submail.MessageUnSubscribeRun(submail.AddressBookMessageBuildRequest(addressbookmessage), messageconfig))
}
