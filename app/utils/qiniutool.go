package utils

import (
	"fmt"
	"github.com/opesun/goquery"
	. "github.com/qiniu/api.v6/conf"
	qio "github.com/qiniu/api.v6/io"
	qrs "github.com/qiniu/api.v6/rs"
	"github.com/revel/revel"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func UploadString(content string, name string) error {
	var err error
	var ret qio.PutRet
	ACCESS_KEY = "NPfcHtb0e2EH7lCmmJot21MRr0lCel81S-QlUaJF"
	SECRET_KEY = "6DzF_oVRYhBkq0mqb4txThza_IfQEUey107VXaPq"
	var policy = qrs.PutPolicy{
		Scope: "isaac",
	}
	f := strings.NewReader(content)
	err = qio.Put(nil, &ret, policy.Token(nil), name, f, nil)
	return err
}
func UploadReader(reader io.Reader, name string) error {
	var err error
	var ret qio.PutRet
	ACCESS_KEY = "NPfcHtb0e2EH7lCmmJot21MRr0lCel81S-QlUaJF"
	SECRET_KEY = "6DzF_oVRYhBkq0mqb4txThza_IfQEUey107VXaPq"
	var policy = qrs.PutPolicy{
		Scope: "isaac",
	}
	err = qio.Put(nil, &ret, policy.Token(nil), name, reader, nil)
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
func HttpPostForm(sel, mac, key, tab, tp string) string {
	v := url.Values{}
	v.Set("select_act", sel)
	v.Set("match_act", mac)
	v.Set("key", key)
	v.Set("table", tab)
	url := "https://www.findmima.com/ajax.php?act=select"
	if tp == "2" {
		url = "http://kf.findmima.com/ajax.php?act=select"
	} else if tp == "3" {
		url = "https://qq.findmima.com/ajax.php?act=select"
	} else if tp == "4" {
		url = "http://qun.findmima.com/ajax.php?act=select"
	}
	fmt.Println(url)
	param := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, _ := http.NewRequest("POST", url, param)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data)

}
func GetHTML(url string) (string, error) {
	var err error
	//向服务端发送get请求
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, err := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		return bodystr, err
	}
	return "", err
}
func GetHTMLContentWithURL(url string) (string, error) {
	//
	var err error
	//向服务端发送get请求
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, err := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		return bodystr, err
	}
	return "", err
}
func GetTwitterHTML(url string) (string, error) {
	//
	var err error
	//向服务端发送get请求
	request, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, err := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		//f := bytes.NewReader(str)

		node, er := goquery.ParseString(bodystr)
		if er == nil {
			ns := node.Find(".cards-media-container div")
			if ns != nil && ns.Length() > 0 {
				for i := 0; i < ns.Length(); i++ {
					no := ns.Eq(i)
					img := no.Find("").Attrs("data-url")
					fmt.Println(img[0])
					if img[0] != "" {
						la := strings.LastIndex(img[0], "/")
						na := img[0][(la + 1):]
						la = strings.LastIndex(na, ":")
						na = na[:la]
						if na != "" {
							stt, _ := GetHTMLContentWithURL(img[0])
							go UploadString(stt, na)
						}
					}
				}
			}
		}
		return bodystr, err
	}
	return "", err
}

func GetCodeContent(name string) (string, error) {
	var err error
	//向服务端发送get请求
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s_code.html", QiNiuDomain, name), nil)
	response, err := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		str, err := ioutil.ReadAll(response.Body)
		bodystr := string(str)

		return bodystr, err
	}
	return "", err
}
func getReader(url string) io.Reader {

	//向服务端发送get请求
	request, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		return response.Body
	}
	return nil
}

func DeleteFile(fileName string) bool {
	entryPathes := []qrs.EntryPath{
		qrs.EntryPath{
			Bucket: BulketName,
			Key:    fmt.Sprintf("%s.html", fileName),
		},
		qrs.EntryPath{
			Bucket: BulketName,
			Key:    fmt.Sprintf("%s_code.html", fileName),
		},
	}
	rs := qrs.New(nil)
	_, err := rs.BatchDelete(nil, entryPathes)
	if err != nil {
		//产生错误
		revel.ERROR.Printf("rs.BatchMove failed:%v", err)
		return false
	}
	return true
}
