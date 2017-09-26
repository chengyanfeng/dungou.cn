package datasource

type Dungouset struct {
	Id        string `gorm:"size:512;column:id"`
	Type      string `gorm:"size:512;column:type"`
	Status    int    `gorm:"column:status"`
	Nowsta    string `gorm:"size:512;column:nowsta"`
	Jacks     int    `gorm:"column:jacks"`
	Pressures int    `gorm:"column:pressures"`
}

type Daopan struct {
	Dungou      string `gorm:"size:512;column:dungou"`
	Jilutime    string `gorm:"size:512;column:Jilutime"`
	Shike       string `gorm:"size:512;column:shike"`
	No1         string `gorm:"size:512;column:no1"`
	No2         string `gorm:"size:512;column:no2"`
	No3         string `gorm:"size:512;column:no3"`
	No4         string `gorm:"size:512;column:no4"`
	Zongniuli   string `gorm:"size:512;column:zongniuli"`
	Waizhou     string `gorm:"size:512;column:waizhou"`
	Neizhou     string `gorm:"size:512;column:neizhou"`
	Zuozhuan    string `gorm:"size:512;column:zuozhuan"`
	Youzhuan    string `gorm:"size:512;column:youzhuan"`
	Chaowali    string `gorm:"size:512;column:chaowali"`
	Huixuansudu string `gorm:"size:512;column:huixuansudu"`
}

type Jingbao struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Dyfx     string `gorm:"size:512;column:dyfx"`
	Dpcfh    string `gorm:"size:512;column:dpcfh"`
	Pzjhz    string `gorm:"size:512;column:pzjhz"`
	No1lx    string `gorm:"size:512;column:no1lx"`
	No2lx    string `gorm:"size:512;column:no2lx"`
	PzjJPU   string `gorm:"size:512;column:pzjJPU"`
	Zy       string `gorm:"size:512;column:zy"`
	Sb       string `gorm:"size:512;column:sb"`
	Jn       string `gorm:"size:512;column:jn"`
	No1zz    string `gorm:"size:512;column:no1zz"`
	No2zz    string `gorm:"size:512;column:no2zz"`
	Yl       string `gorm:"size:512;column:yl"`
	Xh       string `gorm:"size:512;column:xh"`
	Pdcfh    string `gorm:"size:512;column:pdcfh"`
	Zzdy     string `gorm:"size:512;column:zzdy"`
	Pddy     string `gorm:"size:512;column:pddy`
	Dpdy     string `gorm:"size:512;column:dpdy"`
	Zzhl     string `gorm:"size:512;column:zzhl"`
	Yxhl     string `gorm:"size:512;column:yxhl"`
	Dpcnl    string `gorm:"size:512;column:dpcnl"`
	Hzj      string `gorm:"size:512;column:hzj"`
	Mf       string `gorm:"size:512;column:mf"`
	Dpss     string `gorm:"size:512;column:dpss"`
}

type Jiaojie struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Xc1      string `gorm:"size:512;column:xc1"`
	Xc2      string `gorm:"size:512;column:xc2"`
	Xc3      string `gorm:"size:512;column:xc3"`
	Xc4      string `gorm:"size:512;column:xc4"`
	Yl1      string `gorm:"size:512;column:yl1"`
	Yl2      string `gorm:"size:512;column:yl2"`
	Yl3      string `gorm:"size:512;column:yl3"`
	Yl4      string `gorm:"size:512;column:yl4"`
	Jdsx     string `gorm:"size:512;column:jdsx"`
	Jdzy     string `gorm:"size:512;column:jdzy"`
}

type Luoxuanji struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Hzm      string `gorm:"size:512;column:hzm"`
	Yl       string `gorm:"size:512;column:yl"`
	Zt       string `gorm:"size:512;column:zt"`
	Sd       string `gorm:"size:512;column:sd"`
}

type Juejin struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Fy       string `gorm:"size:512;column:fy"`
	Hz       string `gorm:"size:512;column:hz"`
	Spq      string `gorm:"size:512;column:spq"`
	CZq      string `gorm:"size:512;column:cZq"`
	Sph      string `gorm:"size:512;column:sph"`
	Czh      string `gorm:"size:512;column:czh"`
	Fw       string `gorm:"size:512;column:fw"`
	Zjdqh    string `gorm:"size:512;column:zjdqh"`
	Pmhhy    string `gorm:"size:512;column:pmhhy"`
	Jhw      string `gorm:"size:512;column:jhw"`
	Dp       string `gorm:"size:512;column:dp"`
	HBW      string `gorm:"size:512;column:hBW"`
	Ep2      string `gorm:"size:512;column:ep2"`
	Zjzs     string `gorm:"size:512;column:zjzs"`
	Zjys     string `gorm:"size:512;column:zjys"`
	Zjyx     string `gorm:"size:512;column:zjyx"`
	Zjzx     string `gorm:"size:512;column:zjzx"`
	Jjjxz    string `gorm:"size:512;column:jjjxz"`
	Jjms     string `gorm:"size:512;column:jjms"`
	Zzms     string `gorm:"size:512;column:zzms"`
}

type Tuya struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Qjdque   string `gorm:"size:512;column:qjdque"`
	Tuyaque  string `gorm:"size:512;column:tuyaque"`
	Tls      string `gorm:"size:512;column:tls"`
	Xcs      string `gorm:"size:512;column:xcs"`
	Sds      string `gorm:"size:512;column:sds"`
	Tlx      string `gorm:"size:512;column:tlx"`
	Xcx      string `gorm:"size:512;column:xcx"`
	Sdx      string `gorm:"size:512;column:sdx"`
	Tlz      string `gorm:"size:512;column:tlz"`
	Xcz      string `gorm:"size:512;column:xcz"`
	Sdz      string `gorm:"size:512;column:sdz"`
	Tly      string `gorm:"size:512;column:tly"`
	Xcy      string `gorm:"size:512;column:xcy"`
	Sdy      string `gorm:"size:512;column:sdy"`
	Tuya1    string `gorm:"size:512;column:tuya1"`
	Tuya2    string `gorm:"size:512;column:tuya2"`
	Tuya3    string `gorm:"size:512;column:tuya3"`
	Tuya4    string `gorm:"size:512;column:tuya4"`
	Tuya5    string `gorm:"size:512;column:tuya5"`
	Ztl      string `gorm:"size:512;column:ztl"`
}
