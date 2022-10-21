package handlers

import (
	"manager/models"
	"strings"

	"gorm.io/gorm"
)

type omit bool

type Flagset struct {
	Sdkkeys map[string]string     `json:"sdkKeys"`
	Flags   map[string]Flagrule `json:"flags"`
}

type Flagrule map[string]interface{}

type Audset struct {
	Combine    string     `json:"combine"`
	Conditions []CondInst `json:"conditions"`
}

type CondInst struct {
	*models.ConditionEmbedded
	AttributeID omit     `json:"attributeID,omitempty"`
	Vals        []string `json:"vals"`
}

func BuildFlagset(db *gorm.DB) (fs *Flagset) {
	sdks := buildSdkkeys(db)
	flrules := buildFlagrules(db, []uint{}) // empty array == flag ids

	fs = &Flagset{
		Sdkkeys: *sdks,
		Flags:   flrules,
	}

	return fs
}

func FlagUpdateForPublisher(db *gorm.DB, fls []models.Flag) map[string]Flagrule {
	ids := []uint{}
	for _, fl := range fls {
		ids = append(ids, fl.ID)
	}

	pub := buildFlagrules(db, ids)
	return pub
}

func buildSdkkeys(db *gorm.DB) *map[string]string {
	var sdks []models.Sdkkey
	db.Model(models.Sdkkey{}).Select("key", "type").Find(&sdks)

	hash := map[string]string{}
	for i := range sdks {
		hash[sdks[i].Key] = sdks[i].Type
	}
	return &hash
}

func buildFlagrules(db *gorm.DB, flIds []uint) (frs map[string]Flagrule) {
	var flags []models.Flag
	frs = map[string]Flagrule{}
	db.Model(models.Flag{}).Select("id", "key", "status").Find(&flags, flIds)

	for ind := range flags {
		flag := models.Flag{}
		flagrule := Flagrule{}
		db.Preload("Audiences").First(&flag, flags[ind].ID)
		flagrule["status"] = flag.Status

		for i := range flag.Audiences {
			flagrule[flag.Audiences[i].Key] = *buildAudrule(flag.Audiences[i], db)
		}

		frs[flags[ind].Key] = flagrule
	}

	return frs
}

func buildAudrule(aud models.Audience, db *gorm.DB) (ar *Audset) {
	db.Preload("Conditions").First(&aud)
	conds := getEmbeddedConds(aud, db)
	ar = &Audset{
		Combine:    aud.Combine,
		Conditions: conds,
	}
	return ar
}

func getEmbeddedConds(aud models.Audience, db *gorm.DB) []CondInst {
	conds := []CondInst{}
	for ind := range aud.Conditions {
		cond := aud.Conditions[ind]
		var attr models.Attribute
		db.Find(&attr, cond.AttributeID)
		db.Find(&cond)
		cond.Attribute = attr
		conds = append(conds, CondInst{
			ConditionEmbedded: &models.ConditionEmbedded{
				Condition:    &cond,
				AttributeKey: attr.Key,
			},
			Vals: strings.Split(cond.Vals, ", "),
		})
	}
	return conds
}
