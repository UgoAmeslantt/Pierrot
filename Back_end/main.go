package main

import(
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/api/",api)
	http.HandleFunc("/api/login",login_handler)
	http.HandleFunc("/api/register", register_handler)
	http.HandleFunc("/api/articles", articles_handler)
	http.HandleFunc("/api/articles/", artcile_handler)
	http.HandleFunc("/api/articles/avis/", avis_handler)
	http.HandleFunc("/api/users", users_handler)
	http.HandleFunc("/api/users/", user_handler)
	http.HandleFunc("/api/panier", panier_handler)
	log.Fatal(http.ListenAndServe(":55", nil))
}

func api(w http.ResponseWriter, r *http.Request){

}



func login_handler(w http.ResponseWriter, r *http.Request){

}



func register_handler(w http.ResponseWriter, r *http.Request){

}



func articles_handler(w http.ResponseWriter, r *http.Request){

}



func artcile_handler(w http.ResponseWriter, r *http.Request){

}



func avis_handler(w http.ResponseWriter, r *http.Request){

}



func users_handler(w http.ResponseWriter, r *http.Request){

}



func user_handler(w http.ResponseWriter, r *http.Request){

}


func panier_handler(w http.ResponseWriter, r *http.Request){

}


