package main

import (
	"fmt"
	foo "net/http" //named import
)

func main() {
	fmt.Println("Hello, Go std library")
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	fmt.Println("Response Body:", resp.Body)
	fmt.Println("Response Content Length:", resp.ContentLength)
	fmt.Println("Response Transfer Encoding:", resp.TransferEncoding)
	fmt.Println("Response Uncompressed:", resp.Uncompressed)
	fmt.Println("Response Trailer:", resp.Trailer)
	fmt.Println("Response Request:", resp.Request)
	fmt.Println("Response TLS:", resp.TLS)

}
