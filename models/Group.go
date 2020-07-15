package models

type Group struct {
	GroupID int
	IsGroupBusy bool
}

type Groups []*Group
