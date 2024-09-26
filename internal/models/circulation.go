package models

// Website PriamrySection is the primary section for the circulation

type WebsitePrimarySection struct {
	Type     string `json:"type"`
	Referent struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Provider string `json:"provider"`
		Website  string `json:"website"`
	} `json:"referent"`
}

// WebsiteSection is a circulation section i.e., /news/local /sports/baseball, etc
type WebsiteSection struct {
	Type     string `json:"type"`
	Referent struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Provider string `json:"provider"`
		Website  string `json:"website"`
	} `json:"referent"`
}

// Circulation is an individual document circulation which makes up a list of circulations for an ANS document
type Circulation struct {
	DocumentID            string                `json:"document_id"`
	ID                    string                `json:"id"`
	WebsiteID             string                `json:"website_id"`
	WebsitePrimarySection WebsitePrimarySection `json:"website_primary_section"`
	WebsiteSections       []WebsiteSection      `json:"website_sections"`
	WebsiteURL            string                `json:"website_url"`
}

// A list of circulations for a document which is fetched by ansID
type Circulations struct {
	Circulations []Circulation `json:"circulations"`
}
