package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.Default()
    // defer app.Run(iris.Addr(":8080"))

    // Load all templates from the "./templates" folder
    // where extension is ".html" and parse them
    // using the standard `html/template` package.
    app.RegisterView(iris.HTML("./templates", ".html"))
    app.Get("/", func(ctx iris.Context) {
        // Bind: {{.message}} with "Hello world!"
        ctx.ViewData("message", "use templates/hello.html!") // code ngoai html {{ .message }}
        // Render template file: ./templates/hello.html
        ctx.View("hello.html")
    })

    app.Get("/hello", func(ctx iris.Context) {
        ctx.HTML("<h1>Hello World!</h1>")
    })

   
    // app.Get("/user/{id:string regexp(^[0-9]+$)}")
    app.Get("/user/{id:uint64}", func(ctx iris.Context) {
        userID, _ := ctx.Params().GetUint64("id") //get "id" tu url
        ctx.Writef("User ID: %d", userID) // render to html
    })


    //  QUERY /welcome?firstname=Jane&lastname=Doe
    app.Get("/welcome", func(ctx iris.Context) {
        firstname := ctx.URLParamDefault("firstname", "Guest")
        lastname := ctx.URLParam("lastname") // shortcut for ctx.Request().URL.Query().Get("lastname")

        ctx.Writef("Hello %s %s", firstname, lastname)
    })


    app.Post("/post", func(ctx iris.Context) {

        ids := ctx.URLParamSlice("id")
        names, err := ctx.PostValues("name")
        if err != nil {
            ctx.StopWithError(iris.StatusBadRequest, err)
            return
        }

        ctx.Writef("ids: %v; names: %v", ids, names)
    })




    // h := func(ctx iris.Context) {
    //     ctx.HTML("<b>Hi</b1>")
    // }

    // // handler registration and naming
    // home := app.Get("/", h)
    // home.Name = "home" // {{ urlpath "home"}}
    // // or
    // app.Get("/about", h).Name = "about" // {{ urlpath "about"}}
    // app.Get("/page/{id}", h).Name = "page" // {{ urlpath "page" "17"}}



    app.Listen(":8080")


}