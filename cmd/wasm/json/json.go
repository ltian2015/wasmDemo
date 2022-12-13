package jsonhelper

import "encoding/json"

//将一个json字符串格式化为缩进美观的json字符串
func PrettyJson(input string) (string, error) {
	var raw interface{}
	//将给定的JSON字符串解码为二进制形式。
	err := json.Unmarshal([]byte(input), &raw)
	if err != nil {
		return "", err
	}
	//将二进制的JSON数据以缩进的方式进行编码
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
