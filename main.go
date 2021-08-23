package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"buddhism/api"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	r := gin.Default() //Have Logger Recovery
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "HEAD", "DELETE"}
	r.Use(cors.New(config))
	/*
		r.Use(cors.New(cors.Config{
					AllowOrigins:     []string{"*"},
					AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			})) //open  allows all origins and allowheaders
	*/

	r.GET("/members", api.Getmember)
	r.GET("/members/mun/:tmnum", api.Getmember)
	r.GET("/members/id/:tid", api.Getmember)
	r.POST("/members", api.Postmember)
	r.PUT("/members/:tid", api.Putmember)
	r.DELETE("/members/:tid", api.Deletemember)

	//r.GET("/zhangqis", api.Getzhangqi)
	r.GET("/zhangqisinfoall", api.Getzhangqiinfoall)
	r.GET("/zhangqis/mun/:tmnum", api.Getzhangqi)
	r.GET("/zhangqis/id/:tid", api.Getzhangqi)
	r.POST("/zhangqis", api.Postzhangqi)
	r.PUT("/zhangqis/:tid", api.Putzhangqi)
	r.DELETE("/zhangqis/:tid", api.Deletezhangqi)

	r.GET("/chaojiangerensinfoall", api.Getchaojiangereninfoall)
	r.GET("/chaojiangerens/mun/:tmnum", api.Getchaojiangeren)
	r.GET("/chaojiangerens/id/:tid", api.Getchaojiangeren)
	r.POST("/chaojiangerens", api.Postchaojiangeren)
	r.PUT("/chaojiangerens/:tid", api.Putchaojiangeren)
	r.DELETE("/chaojiangerens/:tid", api.Deletechaojiangeren)

	r.GET("/chaojianzuxiansinfoall", api.Getchaojianzuxianinfoall)
	r.GET("/chaojianzuxians/mun/:tmnum", api.Getchaojianzuxian)
	r.GET("/chaojianzuxians/id/:tid", api.Getchaojianzuxian)
	r.POST("/chaojianzuxians", api.Postchaojianzuxian)
	r.PUT("/chaojianzuxians/:tid", api.Putchaojianzuxian)
	r.DELETE("/chaojianzuxians/:tid", api.Deletechaojianzuxian)

	r.GET("/yinglingsinfoall", api.Getyinglinginfoall)
	r.GET("/yinglings/mun/:tmnum", api.Getyingling)
	r.GET("/yinglings/id/:tid", api.Getyingling)
	r.POST("/yinglings", api.Postyingling)
	r.PUT("/yinglings/:tid", api.Putyingling)
	r.DELETE("/yinglings/:tid", api.Deleteyingling)

	r.GET("/dijizhusinfoall", api.Getdijizhuinfoall)
	r.GET("/dijizhus/mun/:tmnum", api.Getdijizhu)
	r.GET("/dijizhus/id/:tid", api.Getdijizhu)
	r.POST("/dijizhus", api.Postdijizhu)
	r.PUT("/dijizhus/:tid", api.Putdijizhu)
	r.DELETE("/dijizhus/:tid", api.Deletedijizhu)

	r.GET("/dongwulingsinfoall", api.Getdongwulinginfoall)
	r.GET("/dongwulings/mun/:tmnum", api.Getdongwuling)
	r.GET("/dongwulings/id/:tid", api.Getdongwuling)
	r.POST("/dongwulings", api.Postdongwuling)
	r.PUT("/dongwulings/:tid", api.Putdongwuling)
	r.DELETE("/dongwulings/:tid", api.Deletedongwuling)

	r.GET("/membersearchall/:tkeyword", api.Getmembersearchall)

	r.Use(static.Serve("/", static.LocalFile("./web", true))) //static web
	r.Run(":8081")
}
