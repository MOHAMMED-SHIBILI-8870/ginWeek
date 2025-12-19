package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Post() {
	postData := map[string]interface{}{
		"userId": 111,
		"id":     111,
		"title":  "helooo",
		"body":   "hai",
	}

	jsonData,err:= json.Marshal(postData)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}

	resp,err := http.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)

	if err != nil{
		fmt.Println("Error:",err)
		return
	}

	fmt.Printf("Response :%s \n",string(body))
	fmt.Printf("status code :%d \n",resp.StatusCode)

}