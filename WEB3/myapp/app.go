package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)                      // POST방식 데이터 마샬링
	w.Header().Add("content-type", "application/json") // 형식이 예쁘게 나오게 하기 위해
	w.WriteHeader((http.StatusOK))
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() // 라우터 클래스를 만들어서 등록하는 방식
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})
	return mux
}
