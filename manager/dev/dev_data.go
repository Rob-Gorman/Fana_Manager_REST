package dev

// import (
// 	"fmt"
// 	"manager/handlers"
// 	"manager/models"

// 	"gorm.io/gorm"
// )

// func SeedDB(db *gorm.DB) {
// 	seedFlags(db)
// 	seedAttributes(db)
// 	seedAudiences(db)
// 	seedFlagAuds(db) // see this function for tricker query implementation
// 	// seedLogs(db)
// 	seedSdks(db)
// 	handlers.BuildFlagset(db)
// }

// func seedSdks(db *gorm.DB) {
// 	key1 := handlers.NewSDKKey("***-*****-**")
// 	key2 := handlers.NewSDKKey("***-*****-**")
// 	var sdkkeys = []models.Sdkkey{
// 		{Key: key1},
// 		{Key: key2, Type: "server"},
// 	}
// 	db.Create(&sdkkeys)
// }

// func seedFlags(db *gorm.DB) {
// 	var flags = []models.Flag{
// 		{Key: "fake-flag-1", DisplayName: "FAKE FLAG ONE"},
// 		{Key: "experimental-flag-1", DisplayName: "Exp Fl 1"},
// 		{Key: "development-flag-1", DisplayName: "Dev Flag One"},
// 	}
// 	db.Create(&flags)
// }

// func seedAttributes(db *gorm.DB) {
// 	var attrs = []models.Attribute{
// 		{Key: "state", Type: "STR", DisplayName: "State"},
// 		{Key: "student", Type: "BOOL", DisplayName: "Student"},
// 		{Key: "beta", Type: "BOOL", DisplayName: "Beta"},
// 	}
// 	db.Create(&attrs)
// }

// func seedAudiences(db *gorm.DB) {
// 	ca_stu := models.Audience{
// 		Key:         "california_students",
// 		DisplayName: "California Students",
// 		Conditions: []models.Condition{
// 			{
// 				AttributeID: 1, // this references the actual attribute! WOOT
// 				Operator:    "EQ",
// 				Vals:        "california",
// 			},
// 			{
// 				AttributeID: 2,
// 				Operator:    "EQ",
// 				Vals:        "true",
// 			},
// 		},
// 	}
// 	beta_test := models.Audience{
// 		Key:         "beta_testers",
// 		DisplayName: "Beta Testers",
// 		Conditions: []models.Condition{
// 			{
// 				AttributeID: 3,
// 				Operator:    "EQ",
// 				Vals:        "true",
// 			},
// 		},
// 	}

// 	// db.Create(&ca_stu) // could just do this for a single audience
// 	auds := []models.Audience{ca_stu, beta_test}
// 	db.Create(&auds)
// }

// func seedFlagAuds(db *gorm.DB) {
// 	// we're trying to add audiences to a flag (see ERD)
// 	// This function uses 3 DB queries to add audiences to a flag
// 	// (or rather, to add rows to the flag_audiences join table)
// 	// 1 fetch flag being updated (this is the perspective of GORM)
// 	// 1 fetch audience(s) being applied to that flag
// 	// 1 insert to apply the association between those two objects (i.e., add a row) in the flag_audiences join
// 	// there might be better ways to do it, but gotta build the API now

// 	// actually, this particular function updates two flags with multiple audiences
// 	// but the flow remains
// 	var firstFlag, lastFlag models.Flag // initialize targets for query results
// 	// SELECT * FROM flags ORDER BY id LIMIT 1
// 	// i.e., get the fist flag (by id)
// 	db.First(&firstFlag) // results _MARSHALLED_ into a Flag object firstFlag
// 	// SELECT * FROM flags ORDER BY id DESC LIMIT 1
// 	// i.e., get the last flag (by id)
// 	db.Last(&lastFlag)

// 	var auds []models.Audience // initialize slice of Audiences for query result
// 	// SELECT * FROM audiences LIMIT 2
// 	db.Limit(2).Find(&auds) // auds now holds a slice of 2 Audience objects

// 	// this is just for logging, show what our results are
// 	// for i, aud := range auds {
// 	// 	fmt.Println("Item num", i, aud.Key) // printing out the _key_
// 	// }

// 	firstFlag.Audiences = []models.Audience{auds[0]}
// 	lastFlag.Audiences = []models.Audience{auds[1]}

// 	// this line was the needle in the haystack of their docs to make this work:
// 	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&[]models.Flag{firstFlag, lastFlag})
// 	fmt.Println(lastFlag.Audiences[0].Key)
// }

// func seedLogs(db *gorm.DB) {
// 	var flagLogs = []models.FlagLog{
// 		{FlagID: 1, FlagKey: "key1", EventDesc: "We plant'n seeds"},
// 		{FlagID: 1, FlagKey: "key1", EventDesc: "Fr a bountiful harvest"},
// 		{FlagID: 2, FlagKey: "key2", EventDesc: ":blobsweat"},
// 		{FlagID: 2, FlagKey: "key2", EventDesc: "ahhhhhhhhhhh"},
// 	}
// 	db.Create(&flagLogs)

// 	var audLogs = []models.AudienceLog{
// 		{AudienceID: 1, AudienceKey: "akey1", EventDesc: "changed some stuff"},
// 		{AudienceID: 1, AudienceKey: "akey1", EventDesc: "reverted it"},
// 		{AudienceID: 2, AudienceKey: "akey2", EventDesc: "reverted it"},
// 		{AudienceID: 2, AudienceKey: "akey2", EventDesc: "git rebase"},
// 	}

// 	db.Create(&audLogs)

// 	var attrLogs = []models.AttributeLog{
// 		{AttributeID: 2, AttributeKey: "student", EventDesc: "created"},
// 		{AttributeID: 1, AttributeKey: "state", EventDesc: "created"},
// 	}

// 	db.Create(&attrLogs)
// }
