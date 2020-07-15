package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/esuwu/SilentChatBot/delivery"
	"github.com/esuwu/SilentChatBot/repository"
	"github.com/esuwu/SilentChatBot/usecase"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)


func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}


func initDatabase() *pgx.ConnPool{
	db, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			User:     "docker",
			Password: "docker",
			Port:     5432,
			Database: "docker",
		},
		MaxConnections: 50,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func InitRouter(api *delivery.Handlers) *fasthttprouter.Router {
	r := fasthttprouter.New()
	r.POST("/", api.MainHandler)
	return r
}


func initLevels(token string)  *delivery.Handlers {
	db := initDatabase()
	useCase := usecase.NewUseCase(repository.NewDBStore(db), token)
	api := delivery.NewHandlers(useCase)
	return api
}

func main(){
	token, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		log.Fatal("Failed to get token")
	}
	api := initLevels(token)
	router := InitRouter(api)
	log.Println("http server started on 3000 port: ")
	err := fasthttp.ListenAndServe(":3000", router.Handler)
	if err != nil {
		log.Fatal(err)
		return
	}

}