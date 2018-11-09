package weserver

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

type Status struct {
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

type QueryDb struct {
	Value     string `json:"value" form:"value" query:"value"`
	TableName string `json:"table" form:"table" query:"table"`
}

func hwdbRouter() {
	hwdb.GET("/query", queryHwDb)
	hwdb.GET("/delete", deleteHwDb)
	hwdb.POST("/add", addHwDb)
}

type DeleteDb struct {
	Pn        string `json:"pn" form:"pn" query:"pn"`
	TableName string `json:"table" form:"table" query:"table"`
}

type AddDb struct {
	TableName   string `json:"table" form:"table" query:"table"`
	Num         int    `json:"-"`
	PN          string `json:"pn"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Footprint   string `json:"footprint"`
	Symbol      string `json:"symbol"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func addHwDb(c echo.Context) (err error) {
	add := new(AddDb)
	ret := new(Status)
	if err = c.Bind(add); err != nil {
		return
	}
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
		tmp := input.FieldByName(fieldInfo.Name)

		output.Field(i).Set(tmp)
	}
	fmt.Println(result)

	return c.JSON(http.StatusOK, ret)

}
func queryHwDb(c echo.Context) (err error) {
	q := new(QueryDb)
	ret := new(Status)
	if err = c.Bind(q); err != nil {
		return
	}
	var result interface{}
	switch q.TableName {
	case "IC":
		var tmp []IC
		db.Where("value = ?", q.Value).Find(&tmp)
		result = tmp
	case "RES":
		var tmp []RES
		db.Where("value = ?", q.Value).Find(&tmp)
		result = tmp
	case "CAP":
		var tmp []CAP
		db.Where("value = ?", q.Value).Find(&tmp)
		result = tmp
	case "Other":
		var tmp []Other
		db.Where("value = ?", q.Value).Find(&tmp)
		result = tmp
	case "Inductor":
		var tmp []Inductor
		db.Where("value = ?", q.Value).Find(&tmp)
		result = tmp
	case "Switch/Connector":
		var tmp []SwitchConnector
		db.Where("value = ?", q.Value).Find(&tmp)
		result = tmp
	case "Transistor/Diode":
		var tmp []TransistorDiode
		db.Where("value = ?", q.Value).Find(&tmp)
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

func deleteHwDb(c echo.Context) (err error) {
	d := new(DeleteDb)
	ret := new(Status)
	ret.Data = nil
	if err = c.Bind(d); err != nil {
		return
	}
	var result interface{}
	switch d.TableName {
	case "IC":
		var tmp []IC
		db.Where("pn = ?", d.Pn).Find(&tmp)
		result = tmp
	case "RES":
		var tmp []RES
		db.Where("pn = ?", d.Pn).Find(&tmp)
		result = tmp
	case "CAP":
		var tmp []CAP
		db.Where("pn = ?", d.Pn).Find(&tmp)
		result = tmp
	case "Other":
		var tmp []Other
		db.Where("pn = ?", d.Pn).Find(&tmp)
		result = tmp
	case "Inductor":
		var tmp []Inductor
		db.Where("pn = ?", d.Pn).Find(&tmp)
		result = tmp
	case "Switch/Connector":
		var tmp []SwitchConnector
		db.Where("pn = ?", d.Pn).Find(&tmp)
		result = tmp
	case "Transistor/Diode":
		var tmp []TransistorDiode
		db.Where("pn = ?", d.Pn).Find(&tmp)
		result = tmp
	default:
		ret.Error = -1
		ret.Msg = "not found the table"
		return c.JSON(http.StatusOK, ret)
	}
	data := reflect.ValueOf(result)
	if data.Len() == 0 {
		ret.Error = -1
		ret.Msg = "not found the item"
		return c.JSON(http.StatusOK, ret)
	}
	db.Delete(result)
	ret.Error = 0
	ret.Msg = "successful"

	return c.JSON(http.StatusOK, ret)

}

// func deleteHwDb(c echo.Context) (err error) {
// 	d := new(DeleteDb)
// 	ret := new(Status)
// 	ret.Data = nil
// 	if err = c.Bind(d); err != nil {
// 		return
// 	}
// 	fmt.Println(d)

// 	if d.TableName == "IC" {
// 		var result []IC
// 		db.Where("pn = ?", d.Pn).Find(&result)
// 		if len(result) == 0 {
// 			ret.Error = StatusNotFound
// 		} else {
// 			for _, val := range result {
// 				db.Delete(&val)
// 			}
// 			ret.Error = StatusNo
// 		}
// 		return c.JSON(http.StatusOK, ret)
// 	}

// 	if d.TableName == "RES" {
// 		var result []RES
// 		db.Where("pn = ?", d.Pn).Find(&result)
// 		if len(result) == 0 {
// 			ret.Error = StatusNotFound
// 		} else {
// 			for _, val := range result {
// 				db.Delete(&val)
// 			}
// 			ret.Error = StatusNo
// 		}
// 		return c.JSON(http.StatusOK, ret)
// 	}

// 	if d.TableName == "CAP" {
// 		var result []CAP
// 		db.Where("pn = ?", d.Pn).Find(&result)
// 		if len(result) == 0 {
// 			ret.Error = StatusNotFound
// 		} else {
// 			for _, val := range result {
// 				db.Delete(&val)
// 			}
// 			ret.Error = StatusNo
// 		}
// 		return c.JSON(http.StatusOK, ret)
// 	}

// 	if d.TableName == "Inductor" {
// 		var result []Inductor
// 		db.Where("pn = ?", d.Pn).Find(&result)
// 		if len(result) == 0 {
// 			ret.Error = StatusNotFound
// 		} else {
// 			for _, val := range result {
// 				db.Delete(&val)
// 			}
// 			ret.Error = StatusNo
// 		}
// 		return c.JSON(http.StatusOK, ret)
// 	}

// 	if d.TableName == "Other" {
// 		var result []Other
// 		db.Where("pn = ?", d.Pn).Find(&result)
// 		if len(result) == 0 {
// 			ret.Error = StatusNotFound
// 		} else {
// 			for _, val := range result {
// 				db.Delete(&val)
// 			}
// 			ret.Error = StatusNo
// 		}
// 		return c.JSON(http.StatusOK, ret)
// 	}

// 	if d.TableName == "Switch/Connector" {
// 		var result []SwitchConnector
// 		db.Where("pn = ?", d.Pn).Find(&result)
// 		if len(result) == 0 {
// 			ret.Error = StatusNotFound
// 		} else {
// 			for _, val := range result {
// 				db.Delete(&val)
// 			}
// 			ret.Error = StatusNo
// 		}
// 		return c.JSON(http.StatusOK, ret)
// 	}

// 	if d.TableName == "Transistor/Diode" {
// 		var result []TransistorDiode
// 		db.Where("pn = ?", d.Pn).Find(&result)
// 		if len(result) == 0 {
// 			ret.Error = StatusNotFound
// 		} else {
// 			for _, val := range result {
// 				db.Delete(&val)
// 			}
// 			ret.Error = StatusNo
// 		}
// 		return c.JSON(http.StatusOK, ret)
// 	}

// 	ret.Error = StatusError
// 	return c.JSON(http.StatusOK, ret)
// }
