package main

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v2"
)

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
	s := session()
	defer s.Close()

	userToModify, err := Read(id)
	if err != nil {
		return err
	}

	u := userToModify.(UsuarioHotel)
	d := data.(UsuarioHotel)

	u.Apellido = d.Apellido
	u.Nombre = d.Nombre
	u.Numero = d.Numero
	u.NumeroDeHabitacion = d.NumeroDeHabitacion
	u.PaisDeOrigen = d.PaisDeOrigen

	err = r.DB("hotel").Table("usuarios").Get(id).Replace(&u).Exec(s)
	if err != nil {
		return err
	}

	return nil
}

func Delete(id uint) error {
	s := session()
	defer s.Close()

	err := r.DB("hotel").Table("usuarios").Get(id).Delete().Exec(s)

	if err != nil {
		return err
	}
	return nil
}
