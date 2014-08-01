package main

import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type", 
        "text/html",
    )
    io.WriteString(
        res, 
        `<doctype html>
<html>
    <head>
        <title>Hello Gopher</title>
    </head>
    <body>
        Hello Gopher!
    </body>
</html>`,
    )
}
func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8000", nil)   
}