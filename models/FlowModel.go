package models

import "time"

type FlowWeebhookDetail struct {
	Url string `json:"Url"`
}
type FlowActionBody struct {
	ObjectId   string `json:"ObjectId"`
	Name       string `json:"Name"`
	ActionType string `json:"ActionType"`
	StepName   string `json:"StepName"`
	StepId     int64  `json:"StepId"`
	FLowName   string `json:"FlowName"`
	FlowId     int64  `json:"FlowId"`
}
type FlowActionDetails struct {
	Body FlowActionBody `json:"Body"`
}
type FlowAction struct {
	Type       string            `json:"Type"`
	ActionArgs FlowActionDetails `jsom:"ActionArgs"`
}
type FlowAdditionalCondition struct {
	CondittionName string `json:"CondittionName"`
	Value          string `json:"Value"`
}
type FlowCondition struct {
	FiledType                    string                  `json:"FiledType"`
	FieldName                    string                  `json:"FieldName"`
	FiledValue                   string                  `json:"FiledValue"`
	ValidateAdditionalConditions bool                    `json:"ValidateAdditionalConditions"`
	AdditionalConditions         FlowAdditionalCondition `json:"AdditionalConditions"`
}

type FlowValidation struct {
	ConditionOperator string          `json:"ConditionOperator"`
	ConditionList     []FlowCondition `json:"ConditionList"`
}
type FlowEntity struct {
	Id                 int64        `json:"Id"`
	Priority           float64      `json:"Priority"`
	ConnectionId       int64        `json:"ConnectionId"`
	NextFlowId         int64        `json:"NextFlowId"`
	Name               string       `json:"Name"`
	EntityArgs         string       `json:"EntityArgs"`
	ETA                int64        `json:"ETA"`
	ReminderFrequency  int64        `json:"ReminderFrequency"`
	Reminders          []FlowAction `json:"Reminders"`
	IsConnectingEntity bool         `json:"IsConnectingEntity"`
	IsWaitEntity       bool         `json:"IsWaitEntity"`
}
type FlowConnection struct {
	Id             int64          `json:"Id"`
	Name           string         `json:"Name"`
	ParentEntityId int64          `json:"ParentEntityId"`
	Action         []FlowAction   `json:"Action"`
	SuccessId      int64          `json:"SuccessId"`
	FailId         int64          `json:"FailId"`
	Validation     FlowValidation `json:"Validation"`
}

type Flow struct {
	Id            int64                     `json:"Id"`
	Version       string                    `json:"Version"`
	FlowName      string                    `json:"flowName" binding:"required"`
	CompanyId     int64                     `json:"CompanyId"`
	Webhook       FlowWeebhookDetail        `json:"Webhook"`
	DefaultArgs   []string                  `json:"DefaultArgs" binding:"required"`
	StartEntityId int64                     `json:"StartEntityId"`
	Entities      map[string]FlowEntity     `json:"Entities" binding:"required"`
	Connections   map[string]FlowConnection `json:"Connections" binding:"required"`
}

type FlowTaskFields struct {
	FiledType  string `json:"FiledType"`
	FieldName  string `json:"FieldName"`
	FiledValue string `json:"FiledValue"`
}

type FlowTaskMetadata struct {
	CustomArgs        string           `json:"CustomArgs"`
	Fields            []FlowTaskFields `json:"Fields"`
	PauseTime         int64            `json:"PauseTime"`
	PauseDurationUnit string           `json:"PauseDurationUnit"`
}

// Task details to be send by 3rd part application to move the objectId in the workflow
type FlowTask struct {
	ObjectId string           `json:"ObjectId"`
	FlowId   int64            `json:"FlowId"`
	MetaData FlowTaskMetadata `json:"MetaData"`
}

// will be used for out internal item flow and maintain this in db for scheduling
type FlowItem struct {
	ObjectId        string           `json:"ObjectId"`
	FlowId          int64            `json:"FlowId"`
	EntityId        int64            `json:"EntityId"`
	MetaData        FlowTaskMetadata `json:"MetaData"`
	CreatedAt       time.Time        `json:"CreatedAt"`
	IsCurrentEntity bool             `json:"IsCurrentEntity"`
}

//task description to be displaied to the user
type FlowItemDescription struct {
	ItemDetails FlowItem `json:"ItemDetails"`
	FlowName    string   `json:"FlowName"`
	EntityName  string   `json:"EntityName"`
}
