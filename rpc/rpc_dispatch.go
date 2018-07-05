package rpc

import (
	"fmt"
	"os"
	"eosgo-client/common"
	"eosgo-client/model"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"strings"
)

func TChainConfig() {

	uri := "E:/goworkspace/helloword/src/main/eos.conf"
	fmt.Println("loading config file: " + uri)
	file, _ := os.Open(uri)
	common.ConfigInit(file)
}

func CreateAccount() {

	// 1、初始化配置文件
	if common.Config.API_URL == "" {
		TChainConfig()
	}

	// 2、create account name
	account := common.ToolsAccountGenerateName("eosgoeosgo")
	fmt.Println("account: "+account)

	// 3、create account
	trx, err := ContractNewAccount(common.Config.NODE_PRODUCER_NAME, account, common.Config.NODE_PUB_KEY, "", "")
	if err != nil {
		fmt.Println("err: ", err)
	}
	if err != nil {
		fmt.Println("err: ", err)
	}
	if trx != nil {
		fmt.Println("transaction id: ", trx.ID)
	}
}

func RpcPushTransaction() {

	// 1、初始化配置文件
	if common.Config.API_URL == "" {
		TChainConfig()
	}

	auth := model.Authorization{
		common.Config.NODE_PRODUCER_NAME,
		"active",
	}

	// 2、发型货币
	action := model.Action{
		common.Config.NODE_PRODUCER_NAME,
		common.Config.NODE_PRODUCER_NAME,
		"issue",
		[]string{common.Config.NODE_PRODUCER_NAME, "account"},
		[]model.Authorization{auth},
		"",
		map[string]interface{}{"to": "account", "quantity": "2.0001 EOS", "memo": "just 2 coins"},
	}

	trx := model.Transaction{
		0,
		0,
		0,
		"",
		[]string{},
		[]string{},
		[]model.Action{action},
		[]string{},
		[]model.Authorization{auth},
		"",
		0,
		0,
		0,
		"",
		"none",
		[]map[string]interface{}{},
		0,
		0,
		[]model.Action{},
	}

	trxPushed, err := ChainPushTransaction(trx, []string{common.Config.NODE_PUB_KEY}, "")
	if err != nil {
		fmt.Println("err: ", err)
	}
	if trxPushed != nil {
		fmt.Println("transaction id: ", trxPushed.ID)
	}
}


func BuyRam() {

	// 1、初始化配置文件
	if common.Config.API_URL == "" {
		TChainConfig()
	}

	// 2、查询最新价格
	resp, err_http := http.Get("https://tbeospre.mytokenpocket.vip/v1/ram_price")
	if err_http != nil {
		fmt.Println("err: ", err_http)
	}
	defer resp.Body.Close()
	body, err_ReadAll := ioutil.ReadAll(resp.Body)
	if err_ReadAll != nil {
		fmt.Println("err: ", err_ReadAll)
	}
	fmt.Println(string(body))

	type RamPrice struct {
		result  string
		message string
		data float64
	}
	var ramPrice RamPrice
	err_Unmarshal := json.Unmarshal(body, ramPrice)
	if err_Unmarshal != nil {
		fmt. Println ( "error:" , err_Unmarshal )
	}
	fmt.Printf ( "%+v" , ramPrice)
	i, _ := strconv.ParseFloat("1024", 64)
	floatStr := fmt.Sprintf("%.5", i / ramPrice.data )
	real_time_price, _ := strconv.ParseFloat(floatStr, 64)


	/*	maxRam = br[0].rows[0].max_ram_size;
		var ramBaseBalance = ar[0].rows[0].base.balance; // Amount of RAM bytes in use

		var ramUsed = 1 - (ramBaseBalance - maxRam);
		target = document.getElementById("maxRam");
		target.innerHTML = (maxRam / 1024 / 1024 / 1024).toFixed(2) + " GB";

		var ramUtilization = (ramUsed / maxRam) * 100;*/
	reqbody_a := `
        {json:"true", code:"eosio", scope:"eosio", table:"rammarket", limit:"10"}
        `
	body_a, err_a :=httpPostForm("https://api.eosnewyork.io/v1/chain/get_table_rows",reqbody_a)
	var dat_a map[string]interface{}
	json.Unmarshal([]byte(body_a), &dat_a)
	fmt. Println ( "error:" , err_a )

	reqbody_b := `
        {json:"true", code:"eosio", scope:"eosio", table:"global"}
        `
	body_b, err_b :=httpPostForm("https://api.eosnewyork.io/v1/chain/get_table_rows",reqbody_b)
	var dat_b map[string]interface{}
	json.Unmarshal([]byte(body_b), &dat_b)
	fmt. Println ( "error:" , err_b )

	if real_time_price>0.8 {

		wallet := "wallet name"
		err := WalletLock(wallet)
		fmt.Println("err: ", err)

		err_WalletOpen := WalletOpen(wallet)
		fmt.Println("err: ", err_WalletOpen)

		err_WalletUnlock := WalletUnlock(wallet, "")
		fmt.Println("err: ", err_WalletUnlock)

		auth := model.Authorization{
			common.Config.NODE_PRODUCER_NAME,
			"active",
		}

		// 2、内存操作
		action := model.Action{
			common.Config.NODE_PRODUCER_NAME,
			common.Config.NODE_PRODUCER_NAME,
			"issue",
			[]string{common.Config.NODE_PRODUCER_NAME, "account"},
			[]model.Authorization{auth},
			"",
			map[string]interface{}{"to": "account", "quantity": "2.0001 EOS", "memo": "just 2 coins"},
		}

		trx := model.Transaction{
			0,
			0,
			0,
			"",
			[]string{},
			[]string{},
			[]model.Action{action},
			[]string{},
			[]model.Authorization{auth},
			"",
			0,
			0,
			0,
			"",
			"none",
			[]map[string]interface{}{},
			0,
			0,
			[]model.Action{},
		}
		trxPushed, err := ChainPushTransaction(trx, []string{common.Config.NODE_PUB_KEY}, "")
		if err != nil {
			fmt.Println("err: ", err)
		}
		if trxPushed != nil {
			fmt.Println("transaction id: ", trxPushed.ID)
		}
	}
}


func httpPostForm(url string,reqbody string) ([]byte, error){
	//创建请求
	postReq, err := http.NewRequest("POST",
		"http://baidu.com", //post链接
		strings.NewReader(reqbody)) //post内容

	if err != nil {
		fmt.Println("POST请求:创建请求失败", err)
		return []byte(""), err
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	//执行请求
	client := &http.Client{}
	resp, err := client.Do(postReq)
	body := []byte("")
	if err != nil {
		fmt.Println("POST请求:创建请求失败", err)
		return []byte(""), err
	} else {
		//读取响应
		body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
		if err != nil {
			fmt.Println("POST请求:读取body失败", err)
			return []byte(""), err
		}

		fmt.Println("POST请求:创建成功", string(body))
	}
	defer resp.Body.Close()
	return body, err
}
