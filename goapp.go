package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	//"log"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := "./data/" + p.Title + ".txt"
	err := ioutil.WriteFile(filename, p.Body, 0600)
	return err
}

func loadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	body, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	page := &Page{Title: title, Body: body}
	return page, error
}

func viewHandler(response http.ResponseWriter, request *http.Request)  {
	title := request.URL.Path[len("/view/"):]
	page, error := loadPage(title)
	if error != nil {
		//log.Fatal("Error on load page")
		fmt.Fprintf(response, "<h1>%s</h1>", error)
	}
	fmt.Fprintf(response, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func main()  {
	/*page := &Page{Title: "primer", Body: []byte("Nuestra primer página")}
	page.save()
	page := loadPage("primer")
	fmt.Println(page.Title, string(page.Body))*/
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}