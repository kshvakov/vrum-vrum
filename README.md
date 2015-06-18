# vrum-vrum


@see: vrum-vrum/example

```go
package main

import (
    "fmt"
    "github.com/kshvakov/vrum-vrum/server"
)

type User struct {
    Name string
}

func main() {
    app := server.New()
    app.Use(func(c *server.Context) {
      //  if you.Auth {
            c.Set("user", User{Name: "UserName"})
      //  }
    })
    app.Get("/", func(c *server.Context) {
            if user, ok := c.Get("user").(User); ok {
                fmt.Fprint(c, user.Name)
            } else {
                fmt.Fprint(c, "Undefined")
            }
        },
    )
}
```

**App metrics**

`http://node_addr:port/metrics/`

```json
{
   "Version":"go1.4.1",
   "NumGoroutine":7,
   "RequestsTotalCount":2,
   "RequestsTotalTime":374890,
   "RequestCountByLocation":{
      "/metrics/":2
   },
   "RequestTimeByLocation":{
      "/metrics/":374890
   },
   "MemStats":{
      "Alloc":337376,
      "TotalAlloc":438072,
      "Sys":2885880,
      "Lookups":15,
      "Mallocs":1408,
      "Frees":304,
      "HeapAlloc":304,
      "HeapSys":835584,
      "HeapIdle":172032,
      "HeapInuse":663552,
      "HeapReleased":0,
      "HeapObjects":1104,
      "StackInuse":212992,
      "StackSys":212992,
      "MSpanInuse":5824,
      "MSpanSys":16384,
      "MCacheInuse":1200,
      "MCacheSys":16384,
      "BuckHashSys":1440192,
      "GCSys":72048,
      "OtherSys":292296,
      "NextGC":458672,
      "LastGC":1434572370519910434,
      "PauseTotalNs":1025631,
      "NumGC":6
   }
}
```