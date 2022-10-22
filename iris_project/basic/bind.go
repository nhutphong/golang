package main

import "github.com/kataras/iris/v12"

var CONST = " 
ReadJSON(outPtr interface{}) error

ReadJSONProtobuf(ptr proto.Message, opts ...ProtoUnmarshalOptions) error
ReadProtobuf(ptr proto.Message) error
ReadMsgPack(ptr interface{}) error
ReadXML(outPtr interface{}) error
ReadYAML(outPtr interface{}) error

// dung nhieu
ReadForm(formObject interface{}) error
ReadQuery(ptr interface{}) error

ReadBody(ptr interface{}) error

ptr = &struct
code struct, struct tags cho phu hop voi cac func o tren,


"

func main() {
    app := iris.Default()

    // localhost:8080/?name=nguyen-chi-thong&address=gia-lai
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

    //localhost:8080/kataras/27/iris/web/framework
    app.Get("/header", func(ctx iris.Context) {
        var hs myHeaders
        if err := ctx.ReadHeaders(&hs); err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return //end func
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

type FormExample struct {
    Colors []string `form:"colors[]"` // or just "colors".
}

func handleForm(ctx iris.Context) {
    var form FormExample
    err := ctx.ReadForm(&form)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.JSON(iris.Map{"Colors": form.Colors})
}


// app.Get("/{name}/{age:int}/{tail:path}", func(ctx iris.Context) {}
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

// localhost:8080/?name=nguyen-chi-thong&address=gia-lai
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


