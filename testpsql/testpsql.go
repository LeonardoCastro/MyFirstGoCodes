package main

import( "github.com/jmoiron/sqlx"
        _ "github.com/lib/pq"
        "log"
        "fmt"
      )

 var schema = `
 CREATE TABLE closet2(
   id int,
   kind text,
   branch text,
   color text
   )`

type Closet struct {
  Id int
  Kind string
  Branch string
  Color string
}

func main() {

  db, err := sqlx.Open("postgres", "user=Leo dbname=testdb sslmode=disable")

  if err != nil {
    log.Fatal(err)
  }

  // en testdb ya tengo una TABLA llamada closet.
  // Primero coloquemos algo en closet 2.

  db.MustExec("INSERT INTO closet2 (id, kind, branch, color) VALUES ($1, $2, $3, $4)", 1, "sweater", "DEFAULT", "blue")

  // Mostrémoslo
  closet2 := []Closet{}
  db.Select(&closet2, "SELECT * FROM closet2")
  fmt.Println(closet2[0])


  // Ahora leamos la TABLA ya existente
  closet := []Closet{}
  db.Select(&closet, "SELECT * FROM closet")
   for _, c := range closet {
     fmt.Println(c)
   }
}
