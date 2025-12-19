package get

import (
	"fmt"
	"io"
	"net/http"
)

func Get() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err !=nil{
		fmt.Println("Error :",err)
		return
	}

	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error : " ,err)
		return
	}

	fmt.Printf("Response:%s\n",string(body))
	fmt.Printf("status Code:%d\n",resp.StatusCode)
}
