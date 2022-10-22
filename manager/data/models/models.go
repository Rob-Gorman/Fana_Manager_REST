package models

import (
	"time"

	"gorm.io/gorm"
)

type Flag struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Key         string         `json:"key" gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	DisplayName string         `json:"displayName" gorm:"type:varchar(30)"`
	Sdkkey      string         `json:"sdkKey" gorm:"type:varchar(30);default:'not_used_sdk_key'"`
	Status      bool           `json:"status" gorm:"default:false; NOT NULL"`
	Audiences   []Audience     `json:"audiences" gorm:"many2many:flag_audiences; joinForeignKey:FlagID;joinReferences:AudienceID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Audience struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	DisplayName string         `json:"displayName" gorm:"type:varchar(30)"`
	Key         string         `json:"key" gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	Combine     string         `json:"combine" gorm:"default:'ANY'"`
	Flags       []Flag         `json:"flags" gorm:"many2many:flag_audiences; foreignKey:ID"`
	Conditions  []Condition    `json:"conditions" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Attribute struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Key         string         `json:"key" gorm:"type:varchar(30); UNIQUE; NOT NULL"`
	Type        string         `json:"attrType" gorm:"type:varchar(10)"`
	DisplayName string         `json:"displayName" gorm:"type:varchar(30)"`
	Conditions  []Condition    `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Condition struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	AudienceID  uint      `json:"audienceID"`
	Negate      bool      `json:"negate" gorm:"default:false"`
	AttributeID uint      `json:"attributeID"`
	Attribute   Attribute `gorm:"foreignKey:AttributeID; references:ID"`
	Operator    string    `json:"operator" gorm:"default:'EQ'"`
	Vals        string    `json:"vals" gorm:"default:'';not null"`
}

type Sdkkey struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Key       string         `json:"key" gorm:"type:varchar(30); NOT NULL; UNIQUE"`
	Status    bool           `json:"status" gorm:"default:true"`
	Type      string         `json:"type" gorm:"default:'client'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
