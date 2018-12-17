package weserver

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

const (
	OrcadPrefix = "D:\\whyengineer\\WE_HW\\WE_ORCAD\\"
)

type Status struct {
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

type QueryDb struct {
	Value     string `json:"value" form:"value" query:"value"`
	TableName string `json:"table" form:"table" query:"table"`
	Num       int    `json:"num" form:"num" query:"num"`
}

type QueryDb1 struct {
	TableName string `json:"table" form:"table" query:"table"`
	Offset    int    `json:"offset" form:"offset" query:"offset"`
	Num       int    `json:"num" form:"num" query:"num"`
}

func hwdbRouter() {
	hwdb.GET("/queryHw", queryHwDb)
	hwdb.GET("/deleteHw", deleteHwDb)
	hwdb.GET("/queryHw1", queryHwDb1)
	hwdb.POST("/addHw", addHwDb)
	hwdb.POST("/updateHw", updateHwDb)

	hwdb.GET("/queryFp", queryFp)
	hwdb.POST("/addFp", addFp)
	hwdb.GET("/listFp", listFp)
	hwdb.GET("/deleteFp", deleteFp)
	hwdb.GET("/findFp", findFp)

	hwdb.POST("/addSymbol", addSymbol)
	hwdb.GET("/querySymbol", querySymbol)
	hwdb.GET("/listSymbol", listSymbol)
	hwdb.GET("/deleteSymbol", deleteSymbol)
	hwdb.GET("/findSymbol", findSymbol)

	hwdb.GET("/getUpToken", upToken)

}

type symbolHandle struct {
	Symbol string `json:"symbol"`
	Url    string `json:"url"`
	Offset int    `json:"offset" form:"offset" query:"offset"`
	Num    int    `json:"num" form:"num" query:"num"`
}

func listFp(c echo.Context) (err error) {
	fp := new(fpHandle)
	ret := new(Status)
	if err = c.Bind(fp); err != nil {
		return
	}
	var tmp []Footprint
	db.Order("id desc").Limit(fp.Num).Offset(fp.Offset).Find(&tmp)

	ret.Error = 0
	ret.Msg = "successful"
	ret.Data = tmp
	return c.JSON(http.StatusOK, ret)
}

func listSymbol(c echo.Context) (err error) {
	symbol := new(symbolHandle)
	ret := new(Status)
	if err = c.Bind(symbol); err != nil {
		return
	}
	var tmp []Symbol
	db.Order("id desc").Limit(symbol.Num).Offset(symbol.Offset).Find(&tmp)

	ret.Error = 0
	ret.Msg = "successful"
	ret.Data = tmp
	return c.JSON(http.StatusOK, ret)
}
func addFp(c echo.Context) (err error) {
	fp := new(fpHandle)
	ret := new(Status)
	if access, msg := accessCheck(c); !access {
		ret.Error = -1
		ret.Msg = msg
		return c.JSON(http.StatusOK, ret)
	}
	if err = c.Bind(fp); err != nil {
		return
	}
	s := new(Footprint)
	s.Footprint = fp.Footprint
	s.Url = fp.Url

	ret.Error = 0
	if dbErr := db.Create(s).Error; dbErr != nil {
		ret.Error = -1
		ret.Msg = dbErr.Error()
	}
	return c.JSON(http.StatusOK, ret)

}
func addSymbol(c echo.Context) (err error) {
	symbol := new(symbolHandle)
	ret := new(Status)
	if access, msg := accessCheck(c); !access {
		ret.Error = -1
		ret.Msg = msg
		return c.JSON(http.StatusOK, ret)
	}
	if err = c.Bind(symbol); err != nil {
		return
	}
	s := new(Symbol)
	s.Symbol = symbol.Symbol
	s.Url = symbol.Url

	ret.Error = 0
	if dbErr := db.Create(s).Error; dbErr != nil {
		ret.Error = -1
		ret.Msg = dbErr.Error()
	}
	return c.JSON(http.StatusOK, ret)

}
func findFp(c echo.Context) (err error) {
	fp := new(fpHandle)
	ret := new(Status)
	if err = c.Bind(fp); err != nil {
		return
	}
	var result []Footprint
	db.Where("footprint = ?", fp.Footprint).Find(&result)
	ret.Data = result
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)
}
func findSymbol(c echo.Context) (err error) {
	symbol := new(symbolHandle)
	ret := new(Status)
	if err = c.Bind(symbol); err != nil {
		return
	}
	var result []Symbol
	db.Where("symbol = ?", symbol.Symbol).Find(&result)
	ret.Data = result
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)
}
func querySymbol(c echo.Context) (err error) {
	symbol := new(symbolHandle)
	ret := new(Status)
	if err = c.Bind(symbol); err != nil {
		return
	}
	var result []Symbol
	db.Where("symbol LIKE ?", "%"+symbol.Symbol+"%").Find(&result)
	ret.Data = result
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)
}

type fpHandle struct {
	Footprint string `json:"footprint"`
	Url       string `json:"url"`
	Offset    int    `json:"offset" form:"offset" query:"offset"`
	Num       int    `json:"num" form:"num" query:"num"`
}

func queryFp(c echo.Context) (err error) {
	fp := new(fpHandle)
	ret := new(Status)
	if err = c.Bind(fp); err != nil {
		return
	}
	var result []Footprint
	db.Where("footprint LIKE ?", "%"+fp.Footprint+"%").Find(&result)
	ret.Data = result
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)
}

type DeleteDb struct {
	Pn        string `json:"pn" form:"pn" query:"pn"`
	TableName string `json:"table" form:"table" query:"table"`
}

type AddDb struct {
	TableName   string `json:"table" form:"table" query:"table"`
	PN          string `json:"pn"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Footprint   string `json:"footprint"`
	Symbol      string `json:"symbol"`
	Orcad       string `json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func updateHwDb(c echo.Context) (err error) {
	update := new(AddDb)
	ret := new(Status)
	if access, msg := accessCheck(c); !access {
		ret.Error = -1
		ret.Msg = msg
		return c.JSON(http.StatusOK, ret)
	}
	if err = c.Bind(update); err != nil {
		return
	}
	var result interface{}

	switch update.TableName {
	case "IC":
		var tmp IC
		db.Where("pn = ?", update.PN).First(&tmp)
		result = &tmp
	case "RES":
		var tmp RES
		db.Where("pn = ?", update.PN).First(&tmp)
		result = &tmp
	case "CAP":
		var tmp CAP
		db.Where("pn = ?", update.PN).First(&tmp)
		result = &tmp
	case "Other":
		var tmp Other
		db.Where("pn = ?", update.PN).First(&tmp)
		result = tmp
	case "Inductor":
		var tmp Inductor
		db.Where("pn = ?", update.PN).First(&tmp)
		result = &tmp
	case "Switch/Connector":
		var tmp SwitchConnector
		db.Where("pn = ?", update.PN).First(&tmp)
		result = &tmp
	case "Transistor/Diode":
		var tmp TransistorDiode
		db.Where("pn = ?", update.PN).First(&tmp)
		result = &tmp
	default:
		ret.Error = -1
		ret.Msg = "not found the table"
		ret.Data = nil
		return c.JSON(http.StatusOK, ret)
	}
	output := reflect.ValueOf(result).Elem()
	input := reflect.ValueOf(update).Elem()

	for i := 0; i < output.NumField(); i++ {
		fieldInfo := output.Type().Field(i) // a reflect.StructField
		//fmt.Println(fieldInfo)

		tmp := input.FieldByName(fieldInfo.Name)

		// fmt.Println(tmp)
		if tmp.Kind() != reflect.Invalid {
			output.Field(i).Set(tmp)
		}

	}
	if dbErr := db.Save(result).Error; dbErr != nil {
		ret.Error = -1
		ret.Msg = dbErr.Error()
	}
	return c.JSON(http.StatusOK, ret)
}
func addHwDb(c echo.Context) (err error) {
	add := new(AddDb)
	ret := new(Status)
	if access, msg := accessCheck(c); !access {
		ret.Error = -1
		ret.Msg = msg
		return c.JSON(http.StatusOK, ret)
	}
	if err = c.Bind(add); err != nil {
		return
	}
	add.Orcad = OrcadPrefix + add.Value
	var result interface{}

	switch add.TableName {
	case "IC":
		result = new(IC)
	case "RES":
		result = new(RES)
	case "CAP":
		result = new(CAP)
	case "Other":
		result = new(Other)
	case "Inductor":
		result = new(Inductor)
	case "Switch/Connector":
		result = new(SwitchConnector)
	case "Transistor/Diode":
		result = new(TransistorDiode)
	default:
		ret.Error = -1
		ret.Msg = "not found the table"
		ret.Data = nil
		return c.JSON(http.StatusOK, ret)
	}

	output := reflect.ValueOf(result).Elem()
	input := reflect.ValueOf(add).Elem()

	for i := 0; i < output.NumField(); i++ {
		fieldInfo := output.Type().Field(i) // a reflect.StructField
		//fmt.Println(fieldInfo)

		tmp := input.FieldByName(fieldInfo.Name)

		// fmt.Println(tmp)
		if tmp.Kind() != reflect.Invalid {
			output.Field(i).Set(tmp)
		}

	}
	if dbErr := db.Create(result).Error; dbErr != nil {
		ret.Error = -1
		ret.Msg = dbErr.Error()
	}
	return c.JSON(http.StatusOK, ret)

}
func queryHwDb1(c echo.Context) (err error) {
	q := new(QueryDb1)
	ret := new(Status)
	if err = c.Bind(q); err != nil {
		return
	}
	var result interface{}
	switch q.TableName {
	case "IC":
		var tmp []IC
		db.Order("id desc").Limit(q.Num).Offset(q.Offset).Find(&tmp)
		result = tmp
	case "RES":
		var tmp []RES
		db.Order("id desc").Limit(q.Num).Offset(q.Offset).Find(&tmp)
		result = tmp
	case "CAP":
		var tmp []CAP
		db.Order("id desc").Limit(q.Num).Offset(q.Offset).Find(&tmp)
		result = tmp
	case "Other":
		var tmp []Other
		db.Order("id desc").Limit(q.Num).Offset(q.Offset).Find(&tmp)
		result = tmp
	case "Inductor":
		var tmp []Inductor
		db.Order("id desc").Limit(q.Num).Offset(q.Offset).Find(&tmp)
		result = tmp
	case "Switch/Connector":
		var tmp []SwitchConnector
		db.Order("id desc").Limit(q.Num).Offset(q.Offset).Find(&tmp)
		result = tmp
	case "Transistor/Diode":
		var tmp []TransistorDiode
		db.Order("id desc").Limit(q.Num).Offset(q.Offset).Find(&tmp)
		result = tmp
	default:
		ret.Error = -1
		ret.Msg = "not found the table"
		ret.Data = nil
		return c.JSON(http.StatusOK, ret)
	}

	ret.Error = 0
	ret.Msg = "successful"
	ret.Data = result
	return c.JSON(http.StatusOK, ret)

}
func queryHwDb(c echo.Context) (err error) {
	q := new(QueryDb)
	ret := new(Status)
	// if access, msg := accessCheck(c); !access {
	// 	ret.Error = -1
	// 	ret.Msg = msg
	// 	return c.JSON(http.StatusOK, ret)
	// }
	if err = c.Bind(q); err != nil {
		return
	}
	var result interface{}
	switch q.TableName {
	case "IC":
		var tmp []IC
		db.Order("id desc").Limit(q.Num).Where("value LIKE ?", "%"+q.Value+"%").Find(&tmp)
		result = tmp
	case "RES":
		var tmp []RES
		db.Order("id desc").Limit(q.Num).Where("value LIKE ?", "%"+q.Value+"%").Find(&tmp)
		result = tmp
	case "CAP":
		var tmp []CAP
		db.Order("id desc").Limit(q.Num).Where("value LIKE ?", "%"+q.Value+"%").Find(&tmp)
		result = tmp
	case "Other":
		var tmp []Other
		db.Order("id desc").Limit(q.Num).Where("value LIKE ?", "%"+q.Value+"%").Find(&tmp)
		result = tmp
	case "Inductor":
		var tmp []Inductor
		db.Order("id desc").Limit(q.Num).Where("value LIKE ?", "%"+q.Value+"%").Find(&tmp)
		result = tmp
	case "Switch/Connector":
		var tmp []SwitchConnector
		db.Order("id desc").Limit(q.Num).Where("value LIKE ?", "%"+q.Value+"%").Find(&tmp)
		result = tmp
	case "Transistor/Diode":
		var tmp []TransistorDiode
		db.Order("id desc").Limit(q.Num).Where("value LIKE ?", "%"+q.Value+"%").Find(&tmp)
		result = tmp
	default:
		ret.Error = -1
		ret.Msg = "not found the table"
		ret.Data = nil
		return c.JSON(http.StatusOK, ret)
	}

	ret.Error = 0
	ret.Msg = "successful"
	ret.Data = result
	return c.JSON(http.StatusOK, ret)

}
func accessCheck(c echo.Context) (bool, string) {
	sess, _ := session.Get("session", c)
	if sess.Values["groups"] == nil {
		return false, "Please login"
	}
	groups := sess.Values["groups"].(string)
	group := strings.Split(groups, ",")
	r := false
	for _, g := range group {
		fmt.Println(g)
		if g == "hwdb" {
			r = true
		}
	}
	return r, "You don't have right to access database"

}
func deleteFp(c echo.Context) (err error) {
	ret := new(Status)
	if access, msg := accessCheck(c); !access {
		ret.Error = -1
		ret.Msg = msg
		return c.JSON(http.StatusOK, ret)
	}
	d := new(fpHandle)
	ret.Data = nil
	if err = c.Bind(d); err != nil {
		return
	}

	db.Unscoped().Where("footprint = ?", d.Footprint).Delete(Footprint{})
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)

}
func deleteSymbol(c echo.Context) (err error) {
	ret := new(Status)
	if access, msg := accessCheck(c); !access {
		ret.Error = -1
		ret.Msg = msg
		return c.JSON(http.StatusOK, ret)
	}
	d := new(symbolHandle)
	ret.Data = nil
	if err = c.Bind(d); err != nil {
		return
	}

	db.Unscoped().Where("symbol = ?", d.Symbol).Delete(Symbol{})
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)

}
func deleteHwDb(c echo.Context) (err error) {
	ret := new(Status)
	if access, msg := accessCheck(c); !access {
		ret.Error = -1
		ret.Msg = msg
		return c.JSON(http.StatusOK, ret)
	}
	d := new(DeleteDb)

	ret.Data = nil
	if err = c.Bind(d); err != nil {
		return
	}

	switch d.TableName {
	case "IC":
		db.Unscoped().Where("pn = ?", d.Pn).Delete(IC{})
	case "RES":
		db.Unscoped().Where("pn = ?", d.Pn).Delete(RES{})
	case "CAP":
		db.Unscoped().Where("pn = ?", d.Pn).Delete(CAP{})
	case "Other":
		db.Unscoped().Where("pn = ?", d.Pn).Delete(Other{})
	case "Inductor":
		db.Unscoped().Where("pn = ?", d.Pn).Delete(Inductor{})
	case "Switch/Connector":
		db.Unscoped().Where("pn = ?", d.Pn).Delete(SwitchConnector{})
	case "Transistor/Diode":
		db.Unscoped().Where("pn = ?", d.Pn).Delete(TransistorDiode{})
	default:
		ret.Error = -1
		ret.Msg = "not found the table"
		return c.JSON(http.StatusOK, ret)
	}
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)

}

func upToken(c echo.Context) (err error) {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, getUpToken(name))
}
