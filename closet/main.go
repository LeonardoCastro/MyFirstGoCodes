package main

import(
      "github.com/jmoiron/sqlx"
        _ "github.com/lib/pq"
        "log"
        "fmt"
        "database/sql"
)

type Closet struct {
  Id int
  Type string
  Branch sql.NullString
  Color string
}

func RemoveDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func AskColor(colores []string, color string) bool {
  var answer string
  var election, res bool

  for _, c := range(colores) {
    if c == color {
      election = true
    }
  }

  switch election{
  case true:
    res = true

  default:
    var seen bool
    for seen == false {
      fmt.Println("No such color, do you want to see available colors? [Y/N]")
      fmt.Scanln(&answer)
      switch answer{
      case "Y":
        fmt.Println(colores)
        seen = true
      case "N":
        seen = true
        res =  false
      default:
        fmt.Println("Wrong answer, please try again.")
      }
    }
  }
  return res
}

func main() {

  // Conexión con la base de datos
  db, err := sqlx.Open("postgres", "user=Leo dbname=testdb sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  colores := []string{}

  err = db.Select(&colores, "SELECT color FROM closet")
  if err != nil {
    log.Fatal(err)
  }

  RemoveDuplicates(&colores)

  // Interacción con el usuario. Elección de dos colores
  var color1, color2 string
  var election bool

  for election == false {

    fmt.Println("Which colors do you want to dress?")
    fmt.Scanln(&color1)
    election = AskColor(colores, color1)
  }

    fmt.Println("and...")
  election = false
  for election == false {
    fmt.Scanln(&color2)
    election = AskColor(colores, color2)
  }

  closet := []Closet{}

  err = db.Select(&closet, "SELECT * FROM closet WHERE color LIKE $1 OR color LIKE $2", "%"+color1+"%", "%"+color2+"%")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Tus opciones son:")
  for _, c := range closet {
    fmt.Println(c.Color, c.Type," from ",c.Branch.String)
  }
}
