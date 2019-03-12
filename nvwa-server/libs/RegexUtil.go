package libs

import (
	"regexp"
	"errors"
)

func ReplaceTemplateContent(fileContent []byte, configMap map[string]string) (changeContent []byte, err error) {
	changeContent = fileContent
	for k, v := range configMap {
		// 转换为普通字符
		tk := regexp.QuoteMeta(k)
		if isMatch, _ := regexp.Match(tk, changeContent); isMatch {
			reg, err := regexp.Compile(tk)
			if err != nil {
				return nil, err
			}
			changeContent = reg.ReplaceAll(changeContent, []byte(v))
		} else {
			err = errors.New("无法在文件中找到要匹配的pattern, pattern: " + k)
			return
		}
	}
	return
}
