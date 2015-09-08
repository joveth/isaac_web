package utils

import (
	"fmt"
	. "github.com/qiniu/api.v6/conf"
	qio "github.com/qiniu/api.v6/io"
	"github.com/qiniu/api.v6/rs"
	"io/ioutil"
	"net/http"
	"strings"
)

func UploadString(content string, name string) error {
	var err error
	var ret qio.PutRet
	ACCESS_KEY = "NPfcHtb0e2EH7lCmmJot21MRr0lCel81S-QlUaJF"
	SECRET_KEY = "6DzF_oVRYhBkq0mqb4txThza_IfQEUey107VXaPq"
	var policy = rs.PutPolicy{
		Scope: "isaac",
	}
	f := strings.NewReader(content)
	err = qio.Put(nil, &ret, policy.Token(nil), name, f, nil)
	return err
}

var client = &http.Client{}

func GetHTMLContent(name string) (string, error) {
	var err error
	//向服务端发送get请求
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s.html", QiNiuDomain, name), nil)
	response, err := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, err := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		return bodystr, err
	}
	return "", err
}
