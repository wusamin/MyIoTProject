package handler

import (
	"log"

	conf "maid/pkg/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

// ViewController sets parameter for view.
func ViewController(r *gin.Engine) {
	r.LoadHTMLGlob("template/html/*.html")
	r.Static("js", "./template/js")
	r.Static("css", "./template/css")
	r.Static("images", "./template/images")
	r.Static("favicon", "./template/favicon")
	r.Static("manifest", "./template/manifest")

	r.GET("/maid/view/chart", chartView)

	r.GET("/maid/view/test", test)

	r.GET("/maid/view/dashboard", dashboardView)
	r.GET("/maid/view/install-dashboard", installDashboard)

	r.GET("/maid/view/local/install-dashboard", installLocalDashboard)
	r.GET("/maid/view/local/dashboard", localDashboard)

	for _, v := range dashboardURL {
		r.GET("/maid/view/local/dashboard/"+v, redirectLocalDashboard)
	}

	for _, v := range dashboardURL {
		r.GET("/maid/view/dashboard/"+v, redirectDashboard)
	}

	{
		m := conf.WebSocket

		c := r.Group("/maid/view/telegram")

		c.GET("/connect", func(c *gin.Context) {
			m.HandleRequest(c.Writer, c.Request)
		})

		// c.GET("/ws/cast", func(c *gin.Context) {
		// 	// m.Broadcast([]byte(fmt.Sprintf("sended:%v", fmt.Sprint(time.Now()))))
		// 	m.Broadcast([]byte("日付が変わりました。本日はわたしが、お側で時刻をお知らせしますね。うふふっ"))
		// })

		c.GET("/broadcast", func(c *gin.Context) {
			// m.Broadcast([]byte(fmt.Sprintf("sended:%v", fmt.Sprint(time.Now()))))
			text := c.Query("message")
			m.Broadcast([]byte(text))
		})

		m.HandleMessage(func(s *melody.Session, msg []byte) {
			log.Printf("websocket connection open. [session: %v]\n", string(msg))
		})

		m.HandleConnect(func(s *melody.Session) {
			log.Printf("websocket connection open. [session: %#v]\n", s)
			m.Broadcast([]byte("接続されました"))
		})

		m.HandleDisconnect(func(s *melody.Session) {
			log.Printf("websocket connection closed. [session: %#v]\n", s)
		})
	}

	{
		m := conf.WebSocketVoice

		c := r.Group("/maid/view/telegram-voice")

		c.GET("/connect", func(c *gin.Context) {
			m.HandleRequest(c.Writer, c.Request)
		})

		// c.GET("/ws/cast", func(c *gin.Context) {
		// 	// m.Broadcast([]byte(fmt.Sprintf("sended:%v", fmt.Sprint(time.Now()))))
		// 	m.Broadcast([]byte("日付が変わりました。本日はわたしが、お側で時刻をお知らせしますね。うふふっ"))
		// })

		c.GET("/broadcast", func(c *gin.Context) {
			// m.Broadcast([]byte(fmt.Sprintf("sended:%v", fmt.Sprint(time.Now()))))
			text := c.Query("message")
			m.Broadcast([]byte(text))
		})

		m.HandleMessage(func(s *melody.Session, msg []byte) {
			log.Printf("websocket connection open. [session: %v]\n", string(msg))
		})

		m.HandleConnect(func(s *melody.Session) {
			log.Printf("websocket connection open. [session: %#v]\n", s)
			m.Broadcast([]byte("接続されました"))
		})

		m.HandleDisconnect(func(s *melody.Session) {
			log.Printf("websocket connection closed. [session: %#v]\n", s)
		})
	}

}

var dashboardURL = []string{"roomstatus", "weather", "system", "schedule", "apiwindow"}

var chartView = func(cxt *gin.Context) {
	cxt.HTML(200, "index.html", gin.H{})
}

var dashboardView = func(cxt *gin.Context) {
	cxt.HTML(200, "dashboard.html", gin.H{})
}

var test = func(cxt *gin.Context) {
	cxt.HTML(200, "develop.html", gin.H{})
}

var installDashboard = func(cxt *gin.Context) {
	cxt.HTML(200, "wpa-entrance.html", gin.H{})
}

var installLocalDashboard = func(cxt *gin.Context) {
	cxt.HTML(200, "install-local-dashboard.html", gin.H{})
}

var localDashboard = func(cxt *gin.Context) {
	cxt.HTML(200, "local-dashboard.html", gin.H{})
}

var redirectLocalDashboard = func(cxt *gin.Context) {
	cxt.Redirect(301, "/maid/view/local/dashboard")
}

var redirectDashboard = func(cxt *gin.Context) {
	cxt.Redirect(301, "/maid/view/dashboard")
}

var setRedirect = func(cxt *gin.Context) {

	baseURL := "/maid/view/local/dashboard"
	URLSet := []string{"roomstatus", "weather", "system", "apiwindow", "schedule"}

	for _, URL := range URLSet {
		cxt.Redirect(301, baseURL+URL)
	}
}
