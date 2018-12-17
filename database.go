package weserver

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //init myqsl
)

// Symbol table
type Symbol struct {
	gorm.Model
	Symbol string `json:"symbol" gorm:"index;not null;unique"`
	Url    string `json:"url"`
}

//Footprint table
type Footprint struct {
	gorm.Model
	Footprint string `json:"footprint" gorm:"index;not null;unique"`
	Url       string `json:"url"`
}

// IC table
type IC struct {
	gorm.Model
	PN          string `gorm:"unique;index;not null;" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null;" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null;" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Orcad       string `gorm:"not null;" json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func (IC) TableName() string {
	return "IC"
}

//RES table
type RES struct {
	gorm.Model
	PN          string `gorm:"unique;index;not null;" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null;" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null;" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Orcad       string `gorm:"not null;" json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func (RES) TableName() string {
	return "Resistor"
}

//CAP table
type CAP struct {
	gorm.Model
	PN          string `gorm:"unique;index;not null;" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null;" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null;" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Orcad       string `gorm:"not null;" json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func (CAP) TableName() string {
	return "Capacitor"
}

//Inductor table
type Inductor struct {
	gorm.Model
	PN          string `gorm:"unique;index;not null;" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null;" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null;" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Orcad       string `gorm:"not null;" json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func (Inductor) TableName() string {
	return "Inductor"
}

//TransistorDiode table
type TransistorDiode struct {
	gorm.Model
	PN          string `gorm:"unique;index;not null" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Orcad       string `gorm:"not null" json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func (TransistorDiode) TableName() string {
	return "Transistor_Diode"
}

//SwitchConnector table
type SwitchConnector struct {
	gorm.Model
	PN          string `gorm:"unique;index;not null;" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null;" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null;" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Orcad       string `gorm:"not null;" json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func (SwitchConnector) TableName() string {
	return "Switch_Connector"
}

//Other table
type Other struct {
	gorm.Model
	PN          string `gorm:"unique;index;not null;" json:"pn"`
	Value       string `gorm:"index;not null" json:"value"`
	Type        string `gorm:"not null;" json:"type"`
	Description string `gorm:"not null" json:"description"`
	Footprint   string `gorm:"not null;" json:"footprint"`
	Symbol      string `gorm:"not null" json:"symbol"`
	Orcad       string `gorm:"not null;" json:"orcad"`
	Datasheet   string `json:"datasheet"`
	Vendor1     string `json:"vendor1"`
	Vendor1PN   string `json:"vendor1pn"`
	Vendor2     string `json:"vendor2"`
	Vendor2PN   string `json:"vendor2pn"`
}

func (Other) TableName() string {
	return "Others"
}

type Weuser struct {
	gorm.Model
	Username string `gorm:"unique;index;not null" json:"username"`
	Email    string `gorm:"unique;index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Level    int    `gorm:"not null" json:"level"`
}

//DbConnect connect to the db server and return the object
func DbConnect(conn string) *gorm.DB {
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Panic(err)
		return nil
	}
	// db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&IC{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&CAP{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&RES{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&Inductor{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&TransistorDiode{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&SwitchConnector{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&Other{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&Symbol{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii").AutoMigrate(&Footprint{})
	// db.AutoMigrate(&Weuser{})
	return db
}
