package main

import (
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// 节点 ,机房 ,解析线路 ,解析地区 ,Ingress 网关地址 ,是否可用 ,测试 ,动态加速网关vip ,规避,,,,,,,,,,
type Tutorial struct {
	Cluster        string
	Jifang         string
	JiexiXianlu    string
	JiexiDiqu      string
	IngressGateway string
	Enable         string
	Test           string
	DynamicGateway string
	Guibi          string
}

var Url = "/search/?device_id=58168804978&os_version=13.5.1&action_type=history_keyword_search&ab_feature=794527,1662483,1730957,1538699&search_position=search_bar&net_type=0&keyword_type=hist&app_name=news_article&is_native_req=0&ab_version=668779,660830,662176,1859936,1843682,662099,1861936,668774,1804363,1855226,1459146,1654131,1788408,1469498,1493166,1419598,1157750,1853622,1673135,1593455,1847464,1257520,1809700,1789127,1426442,1529252,1549048,1484967,1789132,1419036,668775,1851965&is_top_searchbar=1&iid=3737993910687872&ac=WIFI&type=hist&pos=5pe9vb%2F88Pzt3vTp5L%2B9p72%2FeBEKeScxeCUfv7GXvb2%2F%2FvTp5L%2B9p72%2FeBEKeScxeCUfv7GXvb2%2F8fLz%2BvTp6Pn4v72nvaysq7Our66qqKqlq6uqqK6pqrGXvb2%2F8fzp9Ono%2Bfi%2Fvae9rqSzpKutqKusr6ivrKqtrK6qsZe9vb%2F88Pzt0fzp9Ono%2Bfi%2Fvae9rqSzpKutqKusr6ivrKqtrK6qsZe9vb%2F88Pzt0fLz%2BvTp6Pn4v72nvaysq7Our66qqKqlq6uqqK6pqrGXvb2%2F8fL%2B%2FPHC8fzp%2BO7pwu3y7r%2B9p73ml729vb2%2F6fTw%2BO7p%2FPDtv72nvayopKmlraqlr6qzra2trKqsqbGXvb29vb%2Ft7%2FLr9PP%2B%2BL%2B9p72%2FeBEKeScxeCUfv7GXvb29vb%2F%2B9Onkv72nvb94EQp5JzF4JR%2B%2FsZe9vb29v%2F7y8u%2F59PP86fjL%2FPHo%2BO6%2Fvae95pe9vb29vb2%2F8fLz%2BvTp6Pn4v72nvaysq7Our66qqKqlq6uqqK6pqrGXvb29vb29v%2FH86fTp6Pn4v72nva6ks6SrrairrK%2Bor6yqrayuqpe9vb294LGXvb29vb%2F8%2Bfnv%2BO7uv72nvb94EQp5JzF4JR97KCp7Kh14ESd4EQp5JRZ4GC51PAp0HA55JTB4GC57AAx4EAp4OTp1PAqspaV4Eip4OTR5IAF4BiB0BBh5JTB4Ih6%2Fl7294Jfg&followbtn_template=%7B%22color_style%22%3A%22red%22%7D&forum=1&fetch_by_ttnet=0&ssmix=a&version_code=7.8.0&vid=385AB3A9-BB2D-44A8-AEEA-2CB1F4E59F2C&channel=local_test&source=search_subtab_switch&tma_jssdk_version=1.70.0.2&from=search_tab&longitude=116.32376&search_sug=1&ab_group=794527,1662483,1730957,1538699&update_version_code=78003&idfa=7E877572-6B14-451A-907E-71DE8C37D2D9&idfv=385AB3A9-BB2D-44A8-AEEA-2CB1F4E59F2C&device_platform=iphone&device_type=iPhone%208%20Plus&keyword=%E7%8E%89%E7%B1%B3%E7%85%AE%E5%A4%9A%E4%B9%85%E6%89%8D%E7%86%9F&openudid=b355ce1358a3744a0e923733b9f5a69b84a2816a&is_incognito=0&ab_client=a1,f2,f7,e1&pd=synthesis&aid=13&search_start_time=1594807850611&cdid=A8100BF5-8481-4B17-8EBC-4A5661C20010&resolution=1242%2A2208&latitude=39.96056&from_pd=synthesis&index_resource=&min_time=&max_time="
var Host = "api3-trial-esr-c.snssdk.com"
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var Client = &http.Client{Transport: tr}

func main() {

	// 打开这个 csv 文件
	file, err := os.Open("esr.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 初始化一个 csv reader，并通过这个 reader 从 csv 文件读取数据
	reader := csv.NewReader(file)
	// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// 通过 readAll 方法返回 csv 文件中的所有内容
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// 遍历从 csv 文件中读取的所有内容，并将其追加到 tutorials2 切片中
	var allClusters []Tutorial
	for _, item := range record {
		//fmt.Println(item)
		tutorial := Tutorial{
			Cluster:        item[0],
			Jifang:         item[1],
			JiexiXianlu:    item[2],
			JiexiDiqu:      item[3],
			IngressGateway: item[4],
			Enable:         item[5],
			Test:           item[6],
			DynamicGateway: item[7],
			Guibi:          item[8],
		}
		allClusters = append(allClusters, tutorial)
	}

	// test

	for i := 1; i < len(allClusters); i++ {
		isOk, _ := DoRequest(Client, Host, Url, allClusters[i].IngressGateway)
		if !isOk {
			fmt.Println("集群：", allClusters[i].Cluster, "检查失败!!!!!!!!")
		} else {
			fmt.Println("集群：", allClusters[i].Cluster, "成功")
		}
	}

}

func DoRequest(client *http.Client, Host, url, addr string) (bool, error) {

	Host = strings.Replace(Host, " ", "", -1)
	url = strings.Replace(url, " ", "", -1)
	addr = strings.Replace(addr, " ", "", -1)

	requst, err := http.NewRequest("GET",
		fmt.Sprintf("https://%s%s", addr, url),
		nil)
	if err != nil {
		return false, err
	}
	requst.Host = Host
	//requst.URL.Host = Host
	requst.Header = http.Header{"Host": []string{Host}}
	response, err := client.Do(requst)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		//fmt.Println(response.Header)
		return true, nil
	}
	return false, err
}
