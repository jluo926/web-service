package album

import (
	"log"

	"github.com/jluo926/GoProjects/web-service/internal/db"
)

func test2() {
	db.Pool.Query("select count(*) from album")
}

func init() {
	log.Println("initializing rep 2")
}
