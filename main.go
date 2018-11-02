// file: main.go

package main

import (
	"github.com/cooladdr/xiuxiubeidanci/datasource"
	"github.com/cooladdr/xiuxiubeidanci/repositories"
	"github.com/cooladdr/xiuxiubeidanci/services"
	"github.com/cooladdr/xiuxiubeidanci/web/controllers"
	_ "github.com/cooladdr/xiuxiubeidanci/web/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	_ "github.com/kataras/iris/sessions"
)

func main() {
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	// Load the template files.
	tmpl := iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)

	app.StaticWeb("/public", "./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})


	db, err := datasource.NewMysql()
	if err != nil {
		app.Logger().Fatalf("error while loading datasource: %v", err)
		return
	}
	defer db.Close()


	repo := repositories.NewWordRepository(db)
	wordService := services.NewWordService(repo)

	words := mvc.New(app.Party("/words"))

	words.Register(wordService)

	words.Handle(new(controllers.WordController))

	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// Disables the updater.
		iris.WithoutVersionChecker,
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
