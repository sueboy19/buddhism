package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"buddhism/functions"
	"buddhism/tablestruct"
)

//var db, err = sql.Open("sqlite3", "buddhism.db")
var db, err = sqlx.Connect("sqlite3", "buddhism.db")

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
func Checklogin(userId string, password string) []tablestruct.Worker {
	tworker := []tablestruct.Worker{}
	ssql := `select * from workers where w_loginid = $1 and w_loginpasswd = $2 `
	err = db.Select(&tworker, ssql, userId, password)
	checkErr(err)
	return tworker
}

func Logininfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"rid":  claims["id"],
		"text": "Hello World.",
	})
}
*/

func Getmember(c *gin.Context) {
	ssql := `select * FROM members `
	mnum := c.Param("tmnum")
	tid := c.Param("tid")
	members := []tablestruct.Member{}

	if mnum == "" && tid == "" {
		err = db.Select(&members, ssql)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		}
	}

	if mnum != "" {
		ssql += " where m_num = $1" //多筆
		err = db.Select(&members, ssql, mnum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		}
	}

	if tid != "" {
		tmember := tablestruct.Member{} //單筆
		ssql += " where tid = $1"
		err = db.Get(&tmember, ssql, tid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		} else {
			members = append(members, tmember) //確定無錯，才會進行填值
		}
	}

	c.JSON(http.StatusOK, members)
}

func Postmember(c *gin.Context) {
	var member tablestruct.Member

	err = c.ShouldBind(&member)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Post member Error! Form key value have Error!",
			"err": err.Error()})
		return
	}

	code, msg := functions.Sqlpost(db, "members", member)
	c.JSON(code, msg)
}

func Putmember(c *gin.Context) {
	var member tablestruct.Member
	tid := c.Param("tid") // c.PostForm("rid")
	err = c.ShouldBind(&member)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Put member Error! Form key value have Error!",
			"err": err.Error()})
		return
	}
	code, msg := functions.Sqlput(db, "members", tid, member)
	c.JSON(code, msg)
}

func Deletemember(c *gin.Context) {
	tid := c.Param("tid") // c.PostForm("rid")
	code, msg := functions.Sqldelete(db, "members", tid)
	c.JSON(code, msg)
}

//長祈    zhangqi
func Getzhangqi(c *gin.Context) {
	ssql := `select * FROM zhangqi `
	mnum := c.Param("tmnum")
	tid := c.Param("tid")
	zhangqis := []tablestruct.Zhangqi{}

	if mnum == "" && tid == "" {
		err = db.Select(&zhangqis, ssql)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found! Getzhangqi",
				"err": err.Error()})
			return
		}
	}

	if mnum != "" {
		ssql += " where m_num = $1" //多筆
		err = db.Select(&zhangqis, ssql, mnum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found! m_num: " + mnum,
				"err": err.Error()})
			return
		}
	}

	if tid != "" {
		tzhangqi := tablestruct.Zhangqi{} //單筆
		ssql += " where tid = $1"
		err = db.Get(&tzhangqi, ssql, tid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found! tid: " + tid,
				"err": err.Error()})
			return
		} else {
			zhangqis = append(zhangqis, tzhangqi) //確定無錯，才會進行填值
		}
	}

	c.JSON(http.StatusOK, zhangqis)
}

func Getzhangqiinfoall(c *gin.Context) {
	// select members.m_num, members.m_addr, members.m_name as name
	// 	   from members
	//       where m_zhangqi = '1' and members.m_name is not null
	//     union
	ssql :=
		`select members.m_num, members.m_addr, zhangqi.name
		   from members left join zhangqi on members.m_num = zhangqi.m_num
		  where m_zhangqi = '1' and zhangqi.name is not null
		  order by members.m_num, zhangqi.tid `

	type t struct {
		Mnum  string `db:"m_num" json:"m_num"`
		Maddr string `db:"m_addr" json:"m_addr"`
		Name  string `db:"name" json:"name"`
	}

	zhangqis := []t{}

	err = db.Select(&zhangqis, ssql)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not Found! Getzhangqiinfoall",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, zhangqis)
}

func Postzhangqi(c *gin.Context) {
	var zhangqi tablestruct.Zhangqi

	err = c.ShouldBind(&zhangqi)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Post zhangqi Error! Form key value have Error!",
			"err": err.Error()})
		return
	}

	code, msg := functions.Sqlpost(db, "zhangqi", zhangqi)
	c.JSON(code, msg)
}

func Putzhangqi(c *gin.Context) {
	var zhangqi tablestruct.Zhangqi
	tid := c.Param("tid") // c.PostForm("rid")
	err = c.ShouldBind(&zhangqi)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Put zhangqi Error! Form key value have Error!",
			"err": err.Error()})
		return
	}
	code, msg := functions.Sqlput(db, "zhangqi", tid, zhangqi)
	c.JSON(code, msg)
}

func Deletezhangqi(c *gin.Context) {
	tid := c.Param("tid") // c.PostForm("rid")
	code, msg := functions.Sqldelete(db, "zhangqi", tid)
	c.JSON(code, msg)
}

//超薦個人  chaojiangeren
func Getchaojiangeren(c *gin.Context) {
	ssql := `select * FROM chaojiangeren `
	mnum := c.Param("tmnum")
	tid := c.Param("tid")
	chaojiangerens := []tablestruct.Chaojiangeren{}

	if mnum != "" {
		ssql += " where m_num = $1" //多筆
		err = db.Select(&chaojiangerens, ssql, mnum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		}
	}

	if tid != "" {
		tchaojiangeren := tablestruct.Chaojiangeren{} //單筆
		ssql += " where tid = $1"
		err = db.Get(&tchaojiangeren, ssql, tid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		} else {
			chaojiangerens = append(chaojiangerens, tchaojiangeren) //確定無錯，才會進行填值
		}
	}

	c.JSON(http.StatusOK, chaojiangerens)
}

func Getchaojiangereninfoall(c *gin.Context) {
	// select members.m_num, members.m_name as name, members.m_addr as addr
	// 	   from members
	//       where m_chaojiangeren = '1' and members.m_name is not null
	//       union
	ssql :=
		`select members.m_num, chaojiangeren.name, chaojiangeren.addr
		   from members left join chaojiangeren on members.m_num = chaojiangeren.m_num
		  where m_chaojiangeren = '1' and chaojiangeren.name is not null
		  order by members.m_num, chaojiangeren.tid `

	type t struct {
		Mnum string `db:"m_num" json:"m_num"`
		Addr string `db:"addr" json:"addr"`
		Name string `db:"name" json:"name"`
	}

	chaojiangerens := []t{}

	err = db.Select(&chaojiangerens, ssql)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not Found!",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chaojiangerens)
}

func Postchaojiangeren(c *gin.Context) {
	var chaojiangeren tablestruct.Chaojiangeren

	err = c.ShouldBind(&chaojiangeren)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Post chaojiangeren Error! Form key value have Error!",
			"err": err.Error()})
		return
	}

	code, msg := functions.Sqlpost(db, "chaojiangeren", chaojiangeren)
	c.JSON(code, msg)
}

func Putchaojiangeren(c *gin.Context) {
	var chaojiangeren tablestruct.Chaojiangeren
	tid := c.Param("tid") // c.PostForm("rid")
	err = c.ShouldBind(&chaojiangeren)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Put chaojiangeren Error! Form key value have Error!",
			"err": err.Error()})
		return
	}
	code, msg := functions.Sqlput(db, "chaojiangeren", tid, chaojiangeren)
	c.JSON(code, msg)
}

func Deletechaojiangeren(c *gin.Context) {
	tid := c.Param("tid") // c.PostForm("rid")
	code, msg := functions.Sqldelete(db, "chaojiangeren", tid)
	c.JSON(code, msg)
}

//超薦袓先  chaojianzuxian
func Getchaojianzuxian(c *gin.Context) {
	ssql := `select * FROM chaojianzuxian `
	mnum := c.Param("tmnum")
	tid := c.Param("tid")
	chaojianzuxians := []tablestruct.Chaojianzuxain{}

	if mnum != "" {
		ssql += " where m_num = $1" //多筆
		err = db.Select(&chaojianzuxians, ssql, mnum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		}
	}

	if tid != "" {
		tchaojianzuxian := tablestruct.Chaojianzuxain{} //單筆
		ssql += " where tid = $1"
		err = db.Get(&tchaojianzuxian, ssql, tid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		} else {
			chaojianzuxians = append(chaojianzuxians, tchaojianzuxian) //確定無錯，才會進行填值
		}
	}

	c.JSON(http.StatusOK, chaojianzuxians)
}

func Getchaojianzuxianinfoall(c *gin.Context) {
	ssql :=
		`select members.m_num, chaojianzuxian.zuxian, chaojianzuxian.name, chaojianzuxian.addr, chaojianzuxian.newdead
		   from members left join chaojianzuxian on members.m_num = chaojianzuxian.m_num
		  where m_chaojianzuxian = '1' 
		    and chaojianzuxian.name is not null
		    and chaojianzuxian.zuxian is not null
		  order by members.m_num, chaojianzuxian.tid `

	type t struct {
		Mnum    string `db:"m_num" json:"m_num"`
		Zuxian  string `db:"zuxian" json:"zuxian"`
		Addr    string `db:"addr" json:"addr"`
		Name    string `db:"name" json:"name"`
		Newdead string `db:"newdead" json:"newdead"`
	}

	chaojianzuxians := []t{}

	err = db.Select(&chaojianzuxians, ssql)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not Found!",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chaojianzuxians)
}

func Postchaojianzuxian(c *gin.Context) {
	var chaojianzuxian tablestruct.Chaojianzuxain

	err = c.ShouldBind(&chaojianzuxian)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Post chaojianzuxian Error! Form key value have Error!",
			"err": err.Error()})
		return
	}

	code, msg := functions.Sqlpost(db, "chaojianzuxian", chaojianzuxian)
	c.JSON(code, msg)
}

func Putchaojianzuxian(c *gin.Context) {
	var chaojianzuxian tablestruct.Chaojianzuxain
	tid := c.Param("tid") // c.PostForm("rid")
	err = c.ShouldBind(&chaojianzuxian)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Put chaojianzuxian Error! Form key value have Error!",
			"err": err.Error()})
		return
	}
	code, msg := functions.Sqlput(db, "chaojianzuxian", tid, chaojianzuxian)
	c.JSON(code, msg)
}

func Deletechaojianzuxian(c *gin.Context) {
	tid := c.Param("tid") // c.PostForm("rid")
	code, msg := functions.Sqldelete(db, "chaojianzuxian", tid)
	c.JSON(code, msg)
}

//嬰靈  yingling
func Getyingling(c *gin.Context) {
	ssql := `select * FROM yingling `
	mnum := c.Param("tmnum")
	tid := c.Param("tid")
	yinglings := []tablestruct.Yingling{}

	if mnum != "" {
		ssql += " where m_num = $1" //多筆
		err = db.Select(&yinglings, ssql, mnum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		}
	}

	if tid != "" {
		tyingling := tablestruct.Yingling{} //單筆
		ssql += " where tid = $1"
		err = db.Get(&tyingling, ssql, tid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		} else {
			yinglings = append(yinglings, tyingling) //確定無錯，才會進行填值
		}
	}

	c.JSON(http.StatusOK, yinglings)
}

func Getyinglinginfoall(c *gin.Context) {
	ssql :=
		`select members.m_num, yingling.name, yingling.addr
		   from members left join yingling on members.m_num = yingling.m_num
		  where m_yingling = '1' 
		    and yingling.name is not null
		  order by members.m_num, yingling.tid `

	type t struct {
		Mnum string `db:"m_num" json:"m_num"`
		Name string `db:"name" json:"name"`
		Addr string `db:"addr" json:"addr"`
	}

	yinglings := []t{}

	err = db.Select(&yinglings, ssql)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not Found!",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, yinglings)
}

func Postyingling(c *gin.Context) {
	var yingling tablestruct.Yingling

	err = c.ShouldBind(&yingling)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Post yingling Error! Form key value have Error!",
			"err": err.Error()})
		return
	}

	code, msg := functions.Sqlpost(db, "yingling", yingling)
	c.JSON(code, msg)
}

func Putyingling(c *gin.Context) {
	var yingling tablestruct.Yingling
	tid := c.Param("tid") // c.PostForm("rid")
	err = c.ShouldBind(&yingling)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Put yingling Error! Form key value have Error!",
			"err": err.Error()})
		return
	}
	code, msg := functions.Sqlput(db, "yingling", tid, yingling)
	c.JSON(code, msg)
}

func Deleteyingling(c *gin.Context) {
	tid := c.Param("tid") // c.PostForm("rid")
	code, msg := functions.Sqldelete(db, "yingling", tid)
	c.JSON(code, msg)
}

//地基主  dijizhu
func Getdijizhu(c *gin.Context) {
	ssql := `select * FROM dijizhu `
	mnum := c.Param("tmnum")
	tid := c.Param("tid")
	dijizhus := []tablestruct.Dijizhu{}

	if mnum != "" {
		ssql += " where m_num = $1" //多筆
		err = db.Select(&dijizhus, ssql, mnum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		}
	}

	if tid != "" {
		tdijizhu := tablestruct.Dijizhu{} //單筆
		ssql += " where tid = $1"
		err = db.Get(&tdijizhu, ssql, tid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		} else {
			dijizhus = append(dijizhus, tdijizhu) //確定無錯，才會進行填值
		}
	}

	c.JSON(http.StatusOK, dijizhus)
}

func Getdijizhuinfoall(c *gin.Context) {
	ssql :=
		`select members.m_num, dijizhu.name, dijizhu.addr
		   from members left join dijizhu on members.m_num = dijizhu.m_num
		  where m_dijizhu = '1' 
		    and dijizhu.name is not null
		  order by members.m_num, dijizhu.tid `

	type t struct {
		Mnum string `db:"m_num" json:"m_num"`
		Name string `db:"name" json:"name"`
		Addr string `db:"addr" json:"addr"`
	}

	dijizhus := []t{}

	err = db.Select(&dijizhus, ssql)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not Found!",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dijizhus)
}

func Postdijizhu(c *gin.Context) {
	var dijizhu tablestruct.Dijizhu

	err = c.ShouldBind(&dijizhu)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Post dijizhu Error! Form key value have Error!",
			"err": err.Error()})
		return
	}

	code, msg := functions.Sqlpost(db, "dijizhu", dijizhu)
	c.JSON(code, msg)
}

func Putdijizhu(c *gin.Context) {
	var dijizhu tablestruct.Dijizhu
	tid := c.Param("tid") // c.PostForm("rid")
	err = c.ShouldBind(&dijizhu)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Put dijizhu Error! Form key value have Error!",
			"err": err.Error()})
		return
	}
	code, msg := functions.Sqlput(db, "dijizhu", tid, dijizhu)
	c.JSON(code, msg)
}

func Deletedijizhu(c *gin.Context) {
	tid := c.Param("tid") // c.PostForm("rid")
	code, msg := functions.Sqldelete(db, "dijizhu", tid)
	c.JSON(code, msg)
}

//動物靈  dongwuling
func Getdongwuling(c *gin.Context) {
	ssql := `select * FROM dongwuling `
	mnum := c.Param("tmnum")
	tid := c.Param("tid")
	dongwulings := []tablestruct.Dongwuling{}

	if mnum != "" {
		ssql += " where m_num = $1" //多筆
		err = db.Select(&dongwulings, ssql, mnum)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		}
	}

	if tid != "" {
		tdongwuling := tablestruct.Dongwuling{} //單筆
		ssql += " where tid = $1"
		err = db.Get(&tdongwuling, ssql, tid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "Not Found!",
				"err": err.Error()})
			return
		} else {
			dongwulings = append(dongwulings, tdongwuling) //確定無錯，才會進行填值
		}
	}

	c.JSON(http.StatusOK, dongwulings)
}

func Getdongwulinginfoall(c *gin.Context) {
	ssql :=
		`select members.m_num, dongwuling.name, dongwuling.addr
		   from members left join dongwuling on members.m_num = dongwuling.m_num
		  where m_dongwuling = '1' 
		    and dongwuling.name is not null
		  order by members.m_num, dongwuling.tid `

	type t struct {
		Mnum string `db:"m_num" json:"m_num"`
		Name string `db:"name" json:"name"`
		Addr string `db:"addr" json:"addr"`
	}

	dongwulings := []t{}

	err = db.Select(&dongwulings, ssql)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not Found!",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dongwulings)
}

func Postdongwuling(c *gin.Context) {
	var dongwuling tablestruct.Dongwuling

	err = c.ShouldBind(&dongwuling)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Post dongwuling Error! Form key value have Error!",
			"err": err.Error()})
		return
	}

	code, msg := functions.Sqlpost(db, "dongwuling", dongwuling)
	c.JSON(code, msg)
}

func Putdongwuling(c *gin.Context) {
	var dongwuling tablestruct.Dongwuling
	tid := c.Param("tid") // c.PostForm("rid")
	err = c.ShouldBind(&dongwuling)
	if err != nil {
		c.JSON(404, gin.H{
			"tas": "Put dongwuling Error! Form key value have Error!",
			"err": err.Error()})
		return
	}
	code, msg := functions.Sqlput(db, "dongwuling", tid, dongwuling)
	c.JSON(code, msg)
}

func Deletedongwuling(c *gin.Context) {
	tid := c.Param("tid") // c.PostForm("rid")
	code, msg := functions.Sqldelete(db, "dongwuling", tid)
	c.JSON(code, msg)
}

func Getmembersearchall(c *gin.Context) {
	ssql :=
		`select distinct m_num, m_name, title from
		(
		select '0會員基本資料' as title, m_num, m_name from members where m_name like $1 or m_tel like $1 or m_addr like $1
		union
		select '1長生祈福' as title, m_num, '' as m_name from zhangqi where name like $1
		union
		select '2個人冤親債主' as title, m_num, '' as m_name from chaojiangeren where name like $1 or addr like $1
		union
		select '3超薦祖先' as title, m_num, '' as m_name from chaojianzuxian where zuxian like $1 or name like $1 or addr like $1
		union
		select '4嬰靈' as title, m_num, '' as m_name from yingling where name like $1 or addr like $1
		union
		select '5地基主' as title, m_num, '' as m_name from dijizhu where name like $1 or addr like $1
		union
		select '6動物靈' as title, m_num, '' as m_name from dongwuling where name like $1 or addr like $1
		) order by title, m_num `

	keyword := c.Param("tkeyword")

	type t struct {
		Title string `db:"title" json:"title"`
		Mnum  string `db:"m_num" json:"m_num"`
		Mname string `db:"m_name" json:"m_name"`
	}

	searchs := []t{}

	err = db.Select(&searchs, ssql, "%"+keyword+"%")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not Found!",
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, searchs)
}
