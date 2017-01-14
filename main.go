package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"strconv"
)

var port = os.Getenv("PORT")

func methodHandler(method string, get, post func()) {
	switch method {
	case "GET":
		get()
	case "POST":
		post()
	}
}

func create(w http.ResponseWriter, r *http.Request) {

	get := func() {
		tpl, err := template.New("create").Parse(`
		<html>
			<head>
			<title>Pagina de prubas CRUD</title>
			</head>

			<body>
				<button><a href="/">Back</a></button>
				<div>
					<form>
					<input type="text" name="id" placeholder="ID Card Number">
					<input type="text" name="name" placeholder="Name">
					<input type="text" name="surname" placeholder="Surname">
					<input type="tel" name="telnum" placeholder="Telephone Number">
					<input type="text" name="country" placeholder="Country Of Origin">
					<input type="number" name="room_number" placeholder="Room Number">

					<button formmethod="post" formaction="/create">Create</button>
					</form>
				</div>
			</body>

		</html>
		`)
		if err != nil {
			log.Fatalln(err)
		}
		tpl.Execute(w, nil)
	}
	post := func() {
		stringToInt := func(s string) int {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			return i
		}

		nuevoUsuario := UsuarioHotel{
			ID:                 uint(stringToInt(r.FormValue("id"))),
			Nombre:             r.FormValue("name"),
			Apellido:           r.FormValue("surname"),
			Numero:             stringToInt(r.FormValue("telnum")),
			PaisDeOrigen:       r.FormValue("country"),
			NumeroDeHabitacion: stringToInt(r.FormValue("room_number")),
		}

		err := Create(nuevoUsuario)
		if err != nil {
			log.Fatal(err)
		}
	}

	methodHandler(r.Method, get, post)
}

func index(w http.ResponseWriter, r *http.Request) {

	get := func() {
		tpl, err := template.New("index").Parse(`
		<html>
			<head>
			<title>Pagina de prubas CRUD</title>
			</head>

			<body>
				<ul>
					<li><a href="/create">Create</a></li>
					<li><a href="/read">Read</a></li>
					<li><a href="/update">Update</a></li>
					<li><a href="/delete">Delete</a></li>
				</ul>
			</body>

		</html>
		`)
		if err != nil {
			log.Fatalln(err)
		}
		tpl.Execute(w, nil)
	}
	post := func() {

	}

	methodHandler(r.Method, get, post)
}

func main() {
	// Servidor web
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)

	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
