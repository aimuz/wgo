package xcrypto

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
)

// WXBizMsgCrypt 开放平台的消息加密解密技术方案基于 AES 加解密算法来实现，具体如下：
//
//	EncodingAESKey： 即消息加解密 Key，长度固定为 43 个字符，从 a-z,A-Z,0-9 共 62 个字符中选取。由开发者在创建公众号插件时填写，后也可申请修改。
//	AESKey： AESKey=Base64_Decode(EncodingAESKey + "=")，EncodingAESKey 尾部填充一个字符的 "=", 用 Base64_Decode 生成 32 个字节的 AESKey；
//	AES 采用 CBC 模式，秘钥长度为 32 个字节（256 位），数据采用 PKCS#7 填充； PKCS#7：K 为秘钥字节数（采用 32），
//	Buf 为待加密的内容，N 为其字节数。Buf 需要被填充为 K 的整数倍。在 Buf 的尾部填充(K - N%K)个字节，每个字节的内容 是(K - N%K)。
//
// see: https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/Technical_Plan.html
type WXBizMsgCrypt struct {
	aesCrypto Crypto
	appID     string
}

// NewWXBizMsgCrypt ...
func NewWXBizMsgCrypt(encodingAESKey, appID string) (*WXBizMsgCrypt, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return nil, err
	}

	aesCrypto := NewAESUseCBCWithPKCS7(aesKey, aesKey[:16], len(aesKey))
	wx := &WXBizMsgCrypt{
		aesCrypto: aesCrypto,
		appID:     appID,
	}
	return wx, err
}

var binaryEndian = binary.BigEndian

// Encrypt 对输入的明文进行加密
//
//	AESKey = Base64_Decode(EncodingAESKey + "=");
//	FullStr = random(16B) + msg_len(4B) + msg + appid;
//	msg_encrypt = Base64_Encode( AES_Encrypt( FullStr, AESKey ) );
func (w *WXBizMsgCrypt) Encrypt(plain []byte) ([]byte, error) {
	// random(16B) + msgLen(4B) + len(msg) + len(appID)
	fullLen := 20 + len(plain) + len(w.appID)
	full := make([]byte, fullLen)
	_, err := rand.Read(full[:16])
	if err != nil {
		return nil, err
	}
	binaryEndian.PutUint32(full[16:20], uint32(len(plain)))
	copy(full[20:], plain)
	copy(full[20+len(plain):], w.appID)
	return w.aesCrypto.Encrypt(full)
}

// Decrypt 对输入的加密内容进行解密
func (w *WXBizMsgCrypt) Decrypt(encrypt []byte) ([]byte, error) {
	plain, err := w.aesCrypto.Decrypt(encrypt)
	if err != nil {
		return nil, err
	}
	binaryLen := plain[16:20]
	msgLen := int(binaryEndian.Uint32(binaryLen))
	return plain[20 : msgLen+20], nil
}
