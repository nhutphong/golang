package main

import "github.com/kataras/iris/v12"



func main() {
    app := iris.Default()
    app.Any("/", index)

    // {tail:path}; slice params /name/age/.../.../.../
    app.Get("/{name}/{age:int}/{tail:path}", func(ctx iris.Context) {
        var p myParams
        if err := ctx.ReadParams(&p); err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }

        ctx.Writef("myParams: %#v", p)
    })


    app.Get("/header", func(ctx iris.Context) {
        var hs myHeaders
        if err := ctx.ReadHeaders(&hs); err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }

        ctx.JSON(hs)
    })


    //form
    app.RegisterView(iris.HTML("./templates", ".html"))

    app.Get("/form", showForm)
    app.Post("/form", handleForm)


    app.Listen(":8080")

}



//
func showForm(ctx iris.Context) {
    ctx.View("form.html")
}

type formExample struct {
    Colors []string `form:"colors[]"` // or just "colors".
}

func handleForm(ctx iris.Context) {
    var form formExample
    err := ctx.ReadForm(&form)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.JSON(iris.Map{"Colors": form.Colors})
}



type myParams struct {
    Name string   `param:"name"`
    Age  int      `param:"age"`
    Tail []string `param:"tail"`
}
// All parameters are required, as we already know,
// the router will fire 404 if name or int or tail are missing.



type Person struct {
    Name    string `url:"name,required"`
    Address string `url:"address"`
}

func index(ctx iris.Context) {
    var person Person
    if err := ctx.ReadQuery(&person); err!=nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.Application().Logger().Infof("Person: %#+v", person)
    // ctx.WriteString(person)
    ctx.Writef("person %#v", person)
}


type myHeaders struct {
    RequestID      string `header:"X-Request-Id,required"`
    Authentication string `header:"Authentication,required"`
}


