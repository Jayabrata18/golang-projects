package main

import(
	"github.com/"
	"github.com/gofiber/fiber"
)
func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func intDatabase(){

}

func main(){
	app := fiber.New()
	setupRoutes(app)
	app.listen(3000)

}