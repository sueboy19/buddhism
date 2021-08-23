package tablestruct

type Member struct {
	Tid             string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum            string `db:"m_num" form:"m_num" json:"m_num"`
	Mname           string `db:"m_name" form:"m_name" json:"m_name"`
	Mtel            string `db:"m_tel" form:"m_tel" json:"m_tel"`
	Maddr           string `db:"m_addr" form:"m_addr" json:"m_addr"`
	Mdiandeng       string `db:"m_diandeng" form:"m_diandeng" json:"m_diandeng"`
	Mzhangqi        string `db:"m_zhangqi" form:"m_zhangqi" json:"m_zhangqi"`
	Mchaojiangeren  string `db:"m_chaojiangeren" form:"m_chaojiangeren" json:"m_chaojiangeren"`
	Mchaojianzuxian string `db:"m_chaojianzuxian" form:"m_chaojianzuxian" json:"m_chaojianzuxian"`
	Myingling       string `db:"m_yingling" form:"m_yingling" json:"m_yingling"`
	Mdijizhu        string `db:"m_dijizhu" form:"m_dijizhu" json:"m_dijizhu"`
	Mdongwuling     string `db:"m_dongwuling" form:"m_dongwuling" json:"m_dongwuling"`
	Created         string `form:"created" json:"created"` //db:"created" 避免Sqlcols產生出現，db自動產生
}

/*type Diandeng struct { //點燈
	Tid  string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum string `db:"m_num" form:"m_num" json:"m_num"`
	Name string `db:"name" form:"name" json:"name"`
	Addr string `db:"addr" form:"addr" json:"addr"`
}*/

type Zhangqi struct { //長祈
	Tid  string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum string `db:"m_num" form:"m_num" json:"m_num"`
	Name string `db:"name" form:"name" json:"name"`
}

type Chaojiangeren struct { //超薦個人
	Tid  string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum string `db:"m_num" form:"m_num" json:"m_num"`
	Name string `db:"name" form:"name" json:"name"`
	Addr string `db:"addr" form:"addr" json:"addr"`
}

type Chaojianzuxain struct { //超薦祖先
	Tid     string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum    string `db:"m_num" form:"m_num" json:"m_num"`
	Zuxian  string `db:"zuxian" form:"zuxian" json:"zuxian"`
	Addr    string `db:"addr" form:"addr" json:"addr"`
	Newdead string `db:"newdead" form:"newdead" json:"newdead"`
	Name    string `db:"name" form:"name" json:"name"`
}

type Yingling struct { //嬰靈
	Tid  string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum string `db:"m_num" form:"m_num" json:"m_num"`
	Name string `db:"name" form:"name" json:"name"`
	Addr string `db:"addr" form:"addr" json:"addr"`
}

type Dijizhu struct { //地基主
	Tid  string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum string `db:"m_num" form:"m_num" json:"m_num"`
	Name string `db:"name" form:"name" json:"name"`
	Addr string `db:"addr" form:"addr" json:"addr"`
}

type Dongwuling struct { //動物靈
	Tid  string `form:"tid" json:"tid"` //db:"tid" 避免Sqlcols產生出現，db自動產生
	Mnum string `db:"m_num" form:"m_num" json:"m_num"`
	Name string `db:"name" form:"name" json:"name"`
	Addr string `db:"addr" form:"addr" json:"addr"`
}
