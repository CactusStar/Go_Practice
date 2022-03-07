package main_1

import (
	"fmt"
	"net/http"
	"encoding/base64"
	"time"
	"html/template"
)

// type HelloHandler struct{}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

// type WorldHandler struct{}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie {
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie {
		Name: "second_cookie",
		Value: "Manning publication Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World")
	c := http.Cookie {
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie {
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func process_1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "")
}
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.HandleFunc("/process", process)
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	http.HandleFunc("/process_1", process_1)
	server.ListenAndServe()
}
