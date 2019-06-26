package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	cli "crud-go-redis/redis"
	"io/ioutil"
	"net/http"
	"strconv"
	"fmt"
)

var conn *cli.PoolConn

func Connect(port int) {
	conn = cli.New()
	router := httprouter.New()
	router.GET("/students", Retrieve)
	router.POST("/students", Insert)
	router.DELETE("/students/:id", Delete)
	fmt.Println("Server Connected and can be accessed at localhost:"+strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), router)
}

func Retrieve(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()
	course := cli.Course{
		Department: q.Get("department"),
		Code:       q.Get("code"),
		Section:    q.Get("section"),
	}
	receivers, err := conn.Get(course)
	if err != nil {
		w.Write([]byte("No content"))
		return
	}
	out, _ := json.Marshal(receivers)
	w.WriteHeader(200)
	w.Write(out)
}

func Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()
	var (
		department = q.Get("department")
		code       = q.Get("code")
		section    = q.Get("section")
	)
	course := cli.Course{
		Department: department,
		Code:       code,
		Section:    section,
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	receiver := cli.Student{}
	if err = json.Unmarshal(b, &receiver); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = conn.Add(course, receiver.Id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(201)
}



func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	q := r.URL.Query()
	var (
		id         = ps.ByName("id")
		department = q.Get("department")
		code       = q.Get("code")
		section    = q.Get("section")
	)
	course := &cli.Course{
		Department: department,
		Code:       code,
		Section:    section,
	}
	if err := conn.Remove(course, id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(204)
}
