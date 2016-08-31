package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8000/user/new"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"usuario":"test","password":"$2y$10$XNgjZfD.SyopQG446F8j2eZ5biew3Feg\/GCWgAoPWgkzpewph.kju","nombres":"prueba","apellidos":"go","email":"rodrigo_2392@hotmail.com","tipoUsuario":"2","activo":"0","created_at":"2015-10-09 11:36:00","updated_at":"2015-10-09 16:36:00","remember_token":"PuwAZOhUkNfJ8vMp1djUV43p0pfkCn4RYoRR6aN866ySimkt8uG1B8MUVkxe"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}