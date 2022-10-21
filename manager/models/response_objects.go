package models

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
