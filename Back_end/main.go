package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", "root:@tcp(localhost)/pierre")
var query, _ = db.Query("SELECT * FROM pierre.user")

type Users struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Confirmation string `json:"confirmation"`
	Email        string `json:"email"`
	Adresse      string
	PP           string `json:"PP"`
}

type Articles struct {
	ID          int     `json:"ID"`
	Name        string  `json:"Name"`
	Description string  `json:"description"`
	Prix        float64 `json:"prix"`
	Picture     string  `json:"picture"`
}

type Avis struct {
	ID         int    `json:"ID"`
	Articlesid int    `json:"ArticlesId"`
	Txt        string `json:"Txt"`
}

type Panier struct {
	Id        int    `json:"ID"`
	ArticleId string `json:"article-id"`
}

type Session struct {
	Id         int    `json:"ID"`
	User_id    int    `json:"User_id"`
	Token      string `json:"token"`
	Expiration string `json:"expiration"`
}

func api(w http.ResponseWriter, r *http.Request) {

}

func login_handler(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user Users
	var session Session
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	// fmt.Println(user)
	emailVar := `SELECT ID, NAME, PASSWORD, EMAIL FROM bobato.usr WHERE EMAIL="` + user.Email + `" AND PASSWORD="` + user.Password + `"`
	var getRaw = db.QueryRow(emailVar)
	getRaw.Scan(&user.ID, &user.Password, &user.Email)
	//fmt.Println(user.Password, session.UserName)
	// fmt.Println(user)
	if user.ID != 0 {
		session.Token = uuid.New().String()
		session.User_id = user.ID
		// fmt.Println(session.Token)
		insert := `INSERT INTO bobato.session (USR_ID, TOKEN) VALUES (` + strconv.Itoa(session.User_id) + `,"` + session.Token + `");`
		// fmt.Println(insert)
		_, err := db.Query(insert)
		a, _ := json.Marshal(session)
		w.Write(a)
		if err != nil {
			// fmt.Println(err)
		}

	}
	*/
}

func register_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var register Users
	// var session Session
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&register)
	// fmt.Println(register)

	if register.Password == register.Confirmation {
		insert := `INSERT INTO pierre.user (NAME, PASSWORD, EMAIL) VALUES ("` + register.Name + `","` + register.Password + `","` + register.Email + `");`
		_, err := db.Query(insert)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// fmt.Println("password != confirm")
	}
}

func articles_handler(w http.ResponseWriter, r *http.Request) {

}

func artcile_handler(w http.ResponseWriter, r *http.Request) {

}

func avis_handler(w http.ResponseWriter, r *http.Request) {

}

func users_handler(w http.ResponseWriter, r *http.Request) {

}

func user_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// var user []Users
	// for query.Next() {
	// 	var usr Users
	// 	query.Scan(&usr.ID, &usr.Name, &usr.Password, &usr.Email, &usr.Adresse)
	// 	user = append(user, usr)
	// }
	// // fmt.Println(user)
	var querys , _ = db.Query("SELECT * FROM pierre.user")
	var usr Users
	querys.Scan(&usr.ID, &usr.Name, &usr.Password, &usr.Email, &usr.Adresse, &usr.PP)

	a, _ := json.Marshal(usr)
	w.Write(a)
}

func panier_handler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/api/", api)
	http.HandleFunc("/api/login", login_handler)
	http.HandleFunc("/api/register", register_handler)
	http.HandleFunc("/api/articles", articles_handler)
	http.HandleFunc("/api/articles/", artcile_handler)
	http.HandleFunc("/api/articles/avis/", avis_handler)
	http.HandleFunc("/api/user", user_handler)
	http.HandleFunc("/api/users/", users_handler)
	http.HandleFunc("/api/panier", panier_handler)
	log.Fatal(http.ListenAndServe(":55", nil))
}
