package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	r "gopkg.in/gorethink/gorethink.v2"
	"strconv"
)

var port = os.Getenv("PORT")

type CRUD interface {
	Create(data interface{}) error
	Read(id uint) (interface{}, error)
	Update(id uint, data interface{}) error
	Delete(id uint) error
}

type UsuarioHotel struct {
	ID                 uint   `gorethink:"id"`
	Nombre             string `gorethink:"nombre"`
	Apellido           string `gorethink:"apellido"`
	Numero             int    `gorethink:"numero"`
	PaisDeOrigen       string `gorethink:"pais_de_origen"`
	NumeroDeHabitacion int    `gorethink:"numero_de_habitacion"`
}

func session() *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func Create(data interface{}) error {
	s := session()
	defer s.Close()

	d := data.(UsuarioHotel)

	err := r.DB("hotel").Table("usuarios").Insert(&d).Exec(s)
	if err != nil {
		return err
	}
	return nil
}

func Read(id uint) (interface{}, error) {
	s := session()
	defer s.Close()

	nUser := UsuarioHotel{}

	cursor, err := r.DB("hotel").Table("usuarios").Get(id).Run(s)
	if err != nil {
		return nil, err
	}

	cursor.One(&nUser)

	return nUser, nil
}

func Update(id uint, data interface{}) error {
	return nil
}

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
