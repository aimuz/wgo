package hook

// Message ...
type Message struct {
	// Encrypt
	Encrypt string `xml:"Encrypt"`

	// ToUserName 开发者微信号
	ToUserName string `xml:"ToUserName"`

	// FromUserName 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName"`

	// CreateTime 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime"`

	// MsgType 消息类型，文本为text
	MsgType string `xml:"MsgType"`

	// MsgID 消息id，64位整型
	MsgID int64 `xml:"MsgId"`

	// Content 文本消息内容
	Content string `xml:"Content"`

	// Event 事件类型，subscribe
	Event string `xml:"Event"`

	//Image *ImageMessage `xml:"Image"` // 消息类型为 image 可用
	//Voice *VoiceMessage `xml:"Voice"` // 消息为 voice 可用
	//Video *VideoMessage `xml:"Video"` // 消息为 video 可用
	//Music *MusicMessage `xml:"Music"` // 消息为 music 可用
	//ArticleCount int32 `xml:"article_count"` // 图文消息个数；当用户发送文本、图片、语音、视频、图文、地理位置这六种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息
	//Articles *CallArticles `xml:"Articles"` // 消息为 news 可用
}

// Messager ..
type Messager interface {
	_is()
}

type implement struct{}

func (implement) _is() {}

// TextMessage ...
type TextMessage struct {
	implement
	Content string
}

// ImageMessage ...
type ImageMessage struct {
	implement
}

// VoiceMessage ...
type VoiceMessage struct {
	implement
}

// VideoMessage ...
type VideoMessage struct {
	implement
	// MediaID 通过素材管理中的接口上传多媒体文件，得到的id。
	MediaID string `xml:"MediaId"`

	// Title 视频消息的标题
	Title string `xml:"Title"`

	// Description 视频消息的描述
	Description string `xml:"Description"`
}

// MusicMessage ...
type MusicMessage struct {
	implement
}

// ArticleMessage ...
type ArticleMessage struct {
	implement
}
