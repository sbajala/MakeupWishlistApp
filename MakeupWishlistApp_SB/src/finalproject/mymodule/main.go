package main

import (
	beauty "finalproject/mymodule/static"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Message struct {
	Title  string
	Status string
}

var message Message

func main() {
	http.HandleFunc("/emptyWishlist", emptyWishlist)
	http.HandleFunc("/removeItem", removeItem)
	http.HandleFunc("/addItem", addItem)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/wishlist", wishList)
	http.HandleFunc("/", homepage)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}

func wishList(response http.ResponseWriter, request *http.Request) {
	wishlist := beauty.GetWishlist()
	renderPage(response, "wishlist.gohtml", wishlist)
}

func removeItem(response http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	beauty.RemoveFromWishlist(item)
	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(item))
}

func addItem(response http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	beauty.AddToWishlist(item)

	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(item))
}

func emptyWishlist(response http.ResponseWriter, request *http.Request) {
	beauty.EmptyWishlist()
	response.Header().Set("Content-Type", "application/json")
}

func homepage(response http.ResponseWriter, request *http.Request) {
	//GET COOKIE TO CUSTOMIZE TITLE
	cookie := GetCookieByName("title", response, request)
	beauty.SetWishlistName(cookie)
	message.Title = cookie

	renderPage(response, "homepage.gohtml", message)
}

func renderPage(response http.ResponseWriter, page string, object interface{}) {

	template, err := template.ParseFiles(page)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = template.ExecuteTemplate(response, page, object)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func setCookie(response http.ResponseWriter, request *http.Request) {
	value := request.FormValue("title")
	value = strings.ToUpper(value)
	value += "'S BEAUTY WISHLIST"
	cookie := &http.Cookie{
		Name:   "title",
		Value:  value,
		MaxAge: 0,
	}
	http.SetCookie(response, cookie)
	http.Redirect(response, request, "/", http.StatusSeeOther)
}

func GetCookieByName(cookieName string, response http.ResponseWriter, request *http.Request) string {
	cookie, err := request.Cookie(cookieName)
	if err != nil {
		cookie = &http.Cookie{
			Name:   "title",
			Value:  "BEAUTY WISHLIST",
			MaxAge: 0,
		}
		http.SetCookie(response, cookie)
		response.WriteHeader(200)
	}
	return cookie.Value
}
