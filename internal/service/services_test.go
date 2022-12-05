package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bugslayer_332/crudAPI/internal/model"
	//"github.com/bugslayer_332/crudAPI/internal/service"
)

func TestCreateUser(t *testing.T) {
	var user model.User

	user.Name = "Yashwanth"
	user.Age = 20
	user.Location = "Banglore"

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		t.Fatalf("couldnt encode the test struct into the io.reader")
	}

	req, err := http.NewRequest("POST", "/create", &buf)
	if err != nil {
		t.Fatalf("couldn't create the new request")
	}

	rec := httptest.NewRecorder()

	//r := mux.NewRouter()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rec, req)

	//http.HandleFunc("/create", CreateUser)

	//log.Fatal(http.ListenAndServe(":8080", nil))

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expectd status ok , recived status %v", res.StatusCode)
	}

	/*defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("couldnt read the reponse body")
	}
	//fmt.Println(b)

	// required_response := json.NewDecoder().Decode(&)
	*/

}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/read", nil)

	if err != nil {
		t.Fatalf("couldn't create the new request")
	}

	rec := httptest.NewRecorder()

	//r := mux.NewRouter()
	handler := http.HandlerFunc(GetAllUser)
	handler.ServeHTTP(rec, req)

	//http.HandleFunc("/create", CreateUser)

	//log.Fatal(http.ListenAndServe(":8080", nil))

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expectd status ok , recived status %v", res.StatusCode)

	}

}
