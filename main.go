package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	r "gopkg.in/gorethink/gorethink.v2"
)

var port = os.Getenv("PORT")

type CRUD interface {
	Create(data interface{}) error
	Read(id uint) (interface{}, error)
	Update(id uint, data interface{}) error
	Delete(id uint) error
}

type UsuarioHotel struct {
	ID                 uint   `json:"id"`
	Nombre             string `json:"nombre"`
	Apellido           string `json:"apellido"`
	Numero             int    `json:"numero"`
	PaisDeOrigen       string `json:"pais_de_origen"`
	NumeroDeHabitacion int    `json:"numero_de_habitacion"`
}

func session() (*r.Session) {
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

func index(response http.ResponseWriter, request *http.Request) {
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
	tpl.Execute(response, nil)
}

func main() {
	// Servidor web
	http.HandleFunc("/", index)

	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
