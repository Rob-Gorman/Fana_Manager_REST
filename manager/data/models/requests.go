package models

import (
	"encoding/json"
	"manager/utils"
)

type FlagSubmit struct {
	Key         string   `json:"key"`
	DisplayName string   `json:"displayName"`
	SdkKey      string   `json:"sdkKey"`
	Audiences   []string `json:"audiences"`
}

// type AttrSubmit struct {
// 	Key         string `json:"key"`
// 	DisplayName string `json:"displayName"`
// 	Type        string `json:"attrType"`
// }

type AudSubmit struct {
	Key         string      `json:"key"`
	DisplayName string      `json:"displayName"`
	Combine     string      `json:"combine"`
	Conditions  []Condition `json:"conditions"`
}

type CondSubmit struct {
	AttributeID uint   `json:"attributeID"`
	Operator    string `json:"operator"`
	Vals        string `json:"vals"`
	Negate      bool   `json:"negate"`
}

func (f *FlagSubmit) FromJSON(req *[]byte) error {
	err := json.Unmarshal(*req, f)
	if err != nil {
		return utils.UnmarshalError(err)
	}
	return err
}

func (a *Attribute) FromJSON(req *[]byte) error {
	err := json.Unmarshal(*req, a)
	if err != nil {
		return utils.UnmarshalError(err)
	}
	return err
}

func (a *Audience) FromJSON(req *[]byte) error {
	err := json.Unmarshal(*req, a)
	if err != nil {
		return utils.UnmarshalError(err)
	}
	return err
}
