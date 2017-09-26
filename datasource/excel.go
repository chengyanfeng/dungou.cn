package datasource

import (
	. "dungou.cn/def"
	. "dungou.cn/util"
	"github.com/tealeg/xlsx"
	//"fmt"
	"fmt"
	"strings"
)

type Excel struct {
}

func (this *Excel) Xsl2Csv(xsl string, csv string) (r []string, e error) {
	cmd := fmt.Sprintf("ssconvert -S %v %v", xsl, csv)
	_, e = Exec(cmd)
	r = []string{}
	if e == nil {
		for i := 0; i < 50; i++ {
			fn := JoinStr(csv, ".", i)
			if FileExists(fn) {
				r = append(r, fn)
			} else {
				break
			}
		}
	} else {
		r = append(r, xsl)
	}
	return
}

func (this *Excel) List(xsl string) (r []string, e error) {
	cmd := fmt.Sprintf("xlsx -l %v", xsl)
	out, e := Exec(cmd)
	r = []string{}
	if e == nil {
		for _, v := range strings.Split(out, "\n") {
			v = Trim(v)
			if !IsEmpty(v) {
				r = append(r, v)
			}
		}
	}
	return
}

func (this *Excel) Bytes2Json(b []byte) (head []P, r []P) {
	r = []P{}
	xlFile, err := xlsx.OpenBinary(b)
	if err != nil {
		Error("Bytes2Json", err)
	} else {
		sheet := xlFile.Sheets[0]
		if len(sheet.Rows) > 1 {
			head = []P{}
			first := sheet.Rows[0]
			for _, cell := range first.Cells {
				col, _ := cell.String()
				col = Replace(col, PUNCTUATION, "")
				if !IsEmpty(col) {
					p := P{"o": col}
					head = append(head, p)
				}
			}
			for _, row := range sheet.Rows[1:] {
				p := P{}
				for i, cell := range row.Cells {
					if i < len(head) {
						tp := cell.Type()
						str, _ := cell.String()
						str = Trim(str)
						v := head[i]
						k := ToString(v["o"])
						if tp == xlsx.CellTypeNumeric {
							if IsInt(str) {
								v["type"] = "long"
								p[k] = ToInt64(str)
							} else if IsFloat(str) {
								v["type"] = "float"
								p[k] = ToFloat(str)
							} else {
								v["type"] = "string"
								p[k], _ = ToDate(str)
							}
							if IsEmpty(str) {
								p[k] = 0
							}
						} else if tp == xlsx.CellTypeDate {
							Debug("CellTypeDate", str)
						} else {
							v["type"] = "string"
							p[k] = Trim(str)
						}
					}
				}
				r = append(r, p)
			}
		}
	}
	return
}
