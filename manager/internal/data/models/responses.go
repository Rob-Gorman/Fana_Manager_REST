package models

import (
	"bytes"
	"encoding/json"
	"manager/utils"
)

type omit bool

type FlagResponse struct {
	*Flag
	Audiences []AudienceNoCondsResponse `json:"audiences"`
	SdkKey    omit                      `json:"sdkKey,omitempty"`
	DeletedAt omit                      `json:"deleted_at,omitempty"`
}

type FlagNoAudsResponse struct {
	*Flag
	SdkKey    omit `json:"sdkKey,omitempty"`
	DeletedAt omit `json:"deleted_at,omitempty"`
	Audiences omit `json:"audiences,omitempty"`
}

type AudienceResponse struct {
	*Audience
	Conditions []ConditionEmbedded  `json:"conditions"`
	Flags      []FlagNoAudsResponse `json:"flags"`
	DeletedAt  omit                 `json:"deleted_at,omitempty"`
}

type AudienceNoCondsResponse struct {
	*Audience
	Flags      omit `json:"flags,omitempty"`
	DeletedAt  omit `json:"deleted_at,omitempty"`
	Conditions omit `json:"conditions,omitempty"`
	Combine    omit `json:"combine,omitempty"`
}

type ConditionEmbedded struct {
	*Condition
	ID           omit   `json:"id,omitempty"`
	AudienceID   omit   `json:"audienceID,omitempty"`
	Attribute    omit   `json:"Attribute,omitempty"`
	AttributeKey string `json:"attribute"`
}

type AttributeResponse struct {
	*Attribute
	Conditions omit                      `json:",omitempty"`
	Audiences  []AudienceNoCondsResponse `json:"audiences"`
}

type AuditResponse struct {
	FlagLogs      []FlagLog      `json:"flagLogs" gorm:"embedded"`
	AudienceLogs  []AudienceLog  `json:"audienceLogs" gorm:"embedded"`
	AttributeLogs []AttributeLog `json:"attributeLogs" gorm:"embedded"`
}

func (fr *FlagResponse) ToJSON() (res *[]byte, err error) {
	w := bytes.Buffer{}
	err = json.NewEncoder(&w).Encode(fr)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	body := w.Bytes()
	res = &body
	return res, err
}

func ToJSON(obj interface{}) (res *[]byte, err error) {
	w := bytes.Buffer{}
	err = json.NewEncoder(&w).Encode(&obj)
	if err != nil {
		return nil, utils.MarshalError(err)
	}

	body := w.Bytes()
	res = &body
	return res, err
}

func (a *Audience) ToResponse(conds []ConditionEmbedded) *AudienceResponse {
	return &AudienceResponse{
		Audience:   a,
		Conditions: conds,
		Flags:      []FlagNoAudsResponse{},
	}
}

func AllFlagsRes(fs *[]Flag) (res []FlagNoAudsResponse) {
	for ind := range *fs {
		res = append(res, FlagNoAudsResponse{Flag: &(*fs)[ind]})
	}
	return
}

func AllAudsRes(auds *[]Audience) (res []AudienceNoCondsResponse) {
	for ind := range *auds {
		res = append(res, AudienceNoCondsResponse{Audience: &(*auds)[ind]})
	}
	return
}
