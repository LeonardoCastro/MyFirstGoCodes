package main

import( "github.com/jmoiron/sqlx"
        _ "github.com/lib/pq"
        "log"
        "fmt"
        "database/sql"
      )

 var schema = `
 CREATE TABLE closet2(
   id int,
   type text,
   branch text,
   color text
   )`

type Closet struct {
  Id int
  Type string
  Branch sql.NullString
  Color string
}

func main() {

  db, err := sqlx.Open("postgres", "user=Leo dbname=testdb sslmode=disable")

  if err != nil {
    log.Fatal(err)
  }

  // en testdb ya tengo una TABLA llamada closet.
  // Primero coloquemos algo en closet 2.
  db.MustExec(schema)

  db.MustExec("INSERT INTO closet2 (id, type, branch, color) VALUES ($1, $2, $3, $4)", 1, "sweater", "DEFAULT", "blue")

  // Mostrémoslo
  closet2 := []Closet{}
  err = db.Select(&closet2, "SELECT * FROM closet2")
  if err != nil {
      log.Println(err)
    }
  fmt.Println(closet2[0])


  // Ahora leamos la TABLA ya existente
  closet := []Closet{}
  err = db.Select(&closet, "SELECT * FROM closet")
  if err != nil {
    log.Println(err)
  }

  // Para poder observar los valores en la tabla:
   for _, c := range closet {
     fmt.Println(c)
   }

   // La columna branch puede tener valores NULL, por lo que hay que hacer entonces:
   for _, c := range closet {
     fmt.Println(c.Branch.String)
   }
}
