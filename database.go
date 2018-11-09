package weserver

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //init myqsl
)

// IC table
type IC struct {
	Num         int    `gorm:"primary_key" json:"num"`
	PN          string `gorm:"index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Datasheet   string
	Vendor1     string `gorm:"not null" json:"vendor1"`
	Vendor1PN   string `gorm:"not null" json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

//RES table
type RES struct {
	Num         int    `gorm:"primary_key" json:"num"`
	PN          string `gorm:"index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Datasheet   string
	Vendor1     string `gorm:"not null" json:"vendor1"`
	Vendor1PN   string `gorm:"not null" json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

//CAP table
type CAP struct {
	Num         int    `gorm:"primary_key" json:"num"`
	PN          string `gorm:"index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Datasheet   string
	Vendor1     string `gorm:"not null" json:"vendor1"`
	Vendor1PN   string `gorm:"not null" json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

//Inductor table
type Inductor struct {
	Num         int    `gorm:"primary_key" json:"num"`
	PN          string `gorm:"index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Datasheet   string
	Vendor1     string `gorm:"not null" json:"vendor1"`
	Vendor1PN   string `gorm:"not null" json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

//TransistorDiode table
type TransistorDiode struct {
	Num         int    `gorm:"primary_key" json:"num"`
	PN          string `gorm:"index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Datasheet   string
	Vendor1     string `gorm:"not null" json:"vendor1"`
	Vendor1PN   string `gorm:"not null" json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

//SwitchConnector table
type SwitchConnector struct {
	Num         int    `gorm:"primary_key" json:"num"`
	PN          string `gorm:"index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Datasheet   string
	Vendor1     string `gorm:"not null" json:"vendor1"`
	Vendor1PN   string `gorm:"not null" json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

//Other table
type Other struct {
	Num         int    `gorm:"primary_key" json:"num"`
	PN          string `gorm:"index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Datasheet   string
	Vendor1     string `gorm:"not null" json:"vendor1"`
	Vendor1PN   string `gorm:"not null" json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

//DbConnect connect to the db server and return the object
func DbConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "frankie:71451085a@tcp(www.whyengineer.com:3306)/hwdb?charset=utf8&parseTime=True")
	if err != nil {
		log.Panic(err)
		return nil
	}
	db.AutoMigrate(&IC{})
	db.AutoMigrate(&CAP{})
	db.AutoMigrate(&RES{})
	db.AutoMigrate(&Inductor{})
	db.AutoMigrate(&TransistorDiode{})
	db.AutoMigrate(&SwitchConnector{})
	db.AutoMigrate(&Other{})
	return db
}
