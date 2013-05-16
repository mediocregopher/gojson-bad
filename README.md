# Encoding

```go
package main

import (
    "fmt"
    "gojson"
)

func main() {
    //TODO this needs a lot more abstraction
    j := gojson.JsonMap(map[string]gojson.Json{
        "a": gojson.JsonInt(4),
        "b": gojson.JsonFloat(4.5),
        "c": gojson.JsonBool(true),
        "d": gojson.JsonString("wut"),
        "e": gojson.JsonList([]gojson.Json{
            gojson.JsonString("ohai"),
        }),
        "f": gojson.JsonMap(map[string]gojson.Json{
            "a": gojson.JsonInt(4),
            "b": gojson.JsonFloat(4.5),
        }),
    })

    //gojson.Encode returns []byte, EncodeStr returns string
    f := gojson.EncodeStr(j)
    fmt.Printf("%s\n",f)
}
```

# Decoding

```go
package main

import (
    "fmt"
    "gojson"
)

func main() {
    j,err := gojson.DecodeStr(`{"a":1,"b":2.2,"c":["foo",false,null]}`)
    if err != nil { panic(err) }

    if gojson.IsMap(j) {
        m := gojson.ToMap(j)

        //Going through all keys
        for k,v := range m {
            fmt.Println(k,v)
        }

        //Going through each key individually:

        //For some reason go's json always turns ints into floats
        if gojson.IsFloat(m["a"]) {
            fmt.Println(gojson.ToFloat(m["a"]))
        }

        if gojson.IsFloat(m["b"]) {
            fmt.Println(gojson.ToFloat(m["b"]))
        }

        if gojson.IsList(m["c"]) {
            l := gojson.ToList(m["c"])

            if gojson.IsString(l[0]) {
                fmt.Println(gojson.ToString(l[0]))
            }

            if gojson.IsBool(l[1]) {
                fmt.Println(gojson.ToBool(l[1]))
            }

            if l[2] == nil {
                fmt.Println(l[2])
            }
        }

    }
}
```
