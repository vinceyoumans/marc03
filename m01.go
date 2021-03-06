package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url01 := "https://spreadsheets.google.com/feeds/list/168xdxefP3gGnrTGg2hwJoeRVfmbEuTB5plFTyd6I5Qo/1/public/values?alt=json"
	//url02 := "https://spreadsheets.google.com/feeds/list/168xdxefP3gGnrTGg2hwJoeRVfmbEuTB5plFTyd6I5Qo/2/public/values?alt=json"
	//url03 := "https://spreadsheets.google.com/feeds/list/168xdxefP3gGnrTGg2hwJoeRVfmbEuTB5plFTyd6I5Qo/3/public/values?alt=json"
	//url04 := "https://spreadsheets.google.com/feeds/list/168xdxefP3gGnrTGg2hwJoeRVfmbEuTB5plFTyd6I5Qo/4/public/values?alt=json"
	//url05 := "https://spreadsheets.google.com/feeds/list/168xdxefP3gGnrTGg2hwJoeRVfmbEuTB5plFTyd6I5Qo/5/public/values?alt=json"
	println("=============  starting main =============")

	// res, err := http.Get("https://www.citibikenyc.com/stations/json")
	res, err := http.Get(url01)
	if err != nil {
		panic(err.Error())
	}

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }

	var m AutoGenerated
	// err := json.Unmarshal(body, &m)
	//json.NewDecoder(res.Body).Decode(&m)
	// json.NewDecoder([]byte(body)).Decode(&m)

	// json.NewDecoder([]byte(res.Body)).Decode(&m)
	json.NewDecoder(res.Body).Decode(&m)

	//err := json.Unmarshal([]byte(body), &m)
	if err != nil {
		fmt.Println("Whoops...:", err)
	}

	fmt.Println("============  about to print m ============")
	fmt.Println(m.Feed.Title.T)
	fmt.Println("============  about to print m2 ============")
	// fmt.Println(m.Feed)
	fmt.Println("============  about to print m3 ============")
	// fmt.Println(m)
	fmt.Println("============  about to print m4 ============")
}
