package dmv

import (
	"strings"
	"time"

	"github.com/hexcraft-biz/xtime"
)

// License suspensions
type LicenseSuspension struct {
	VehicleType string     `json:"vehicleType"`
	Status      string     `json:"status"`
	EndDate     xtime.Time `json:"endDate"`
}

type LicenseSuspensions []*LicenseSuspension

func (rs LicenseSuspensions) IsLoopThresholdPassed(now xtime.Time, hasDUI *bool) bool {
	if len(rs) <= 0 {
		return true
	}

	for _, r := range rs {
		if strings.Contains(r.Status, "酒駕") {
			*hasDUI = true
		}

		if r.VehicleType == "汽車" && r.EndDate.After(now) {
			return false
		}
	}

	return true
}

// Penalties
type Penalty struct {
	CitationNumber    string     `json:"citationNumber"`
	ViolationDate     xtime.Time `json:"violationDate"`
	ViolationLocation string     `json:"violationLocation"`
	Cause             string     `json:"cause"`
	HandlingOffice    string     `json:"handlingOffice"`
	Points            int        `json:"points"`
}

type Penalties []*Penalty

func (rs Penalties) IsLoopThresholdPassed(now xtime.Time, hasDUI *bool) bool {
	if len(rs) <= 0 {
		return true
	}

	halfYearAgo := now.Add(-180 * 24 * time.Hour)
	points := 0

	for _, r := range rs {
		if strings.Contains(r.Cause, "酒駕") {
			*hasDUI = true
		}

		if r.ViolationDate.After(halfYearAgo) {
			points += r.Points
		}
	}

	if points > 3 {
		return false
	}

	return true
}

// Tickets
type Ticket struct {
	CitationNumber    string     `json:"citationNumber"`
	ViolationDate     xtime.Time `json:"violationDate"`
	ViolationLocation string     `json:"violationLocation"`
	Cause             string     `json:"cause"`
	PenalAmount       int        `json:"penalAmount"`
	ResponsibleParty  string     `json:"responsibleParty"`
	PlateNumber       string     `json:"plateNumber"`
	PenalCode         string     `json:"penalCode"`
	HandlingOffice    string     `json:"handlingOffice"`
	AppearanceDate    xtime.Time `json:"appearanceDate"`
	CanPayOnline      bool       `json:"canPayOnline"`
}

type Tickets []*Ticket

func (rs Tickets) IsLoopThresholdPassed() bool {
	if len(rs) > 4 {
		return false
	}

	return true
}
