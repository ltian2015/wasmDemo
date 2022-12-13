package qrcode

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

//GenerateQr函数根据给定的字符串内容生成二维码图片，并将
//二维码图片转换为base64的字符串编码，以便通过HTTP协议进行传输。
func GenerateQr(content string) (string, error) {
	//调用qr库的函数根据内容生成二维码对象。
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return "", err
	}
	//将二维码大小设置为200*200
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return "", err
	}
	//将二维码图片的二进制字节写到字节缓存区中。
	var buf bytes.Buffer
	if err := png.Encode(&buf, qrCode); err != nil {
		return "", err
	}
	data := buf.Bytes()
	//将字节缓冲区中的数据转换为HTTP HTML <image>元素所要求的base64格式的编码。
	imgBase64Str := base64.StdEncoding.EncodeToString(data)
	return imgBase64Str, nil
}
