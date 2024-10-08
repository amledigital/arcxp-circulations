package models

type PrimarySection struct {
	ID                   string      `json:"_id"`
	Website              string      `json:"_website"`
	AdditionalProperties interface{} `json:"additional_properties"`
	Description          string      `json:"description,omitempty"`
	Name                 string      `json:"name"`
	Parent               interface{} `json:"parent,omitempty"`
	ParentID             string      `json:"parent_id,omitempty"`
	Path                 string      `json:"path"`
	Type                 string      `json:"type"`
	Version              string      `json:"version"`
}

type PrimarySite struct {
	ID                   string      `json:"_id"`
	AdditionalProperties interface{} `json:"additional_properties,omitempty"`
	Description          string      `json:"description,omitempty"`
	Name                 string      `json:"name"`
	ParentID             string      `json:"parent_id,omitempty"`
	Path                 string      `json:"path"`
	Type                 string      `json:"type"`
	Version              string      `json:"version"`
}

type Section struct {
	ID                   string      `json:"_id"`
	Website              string      `json:"_website"`
	WebsiteSectionID     string      `json:"_website_section_id"` // concats website.WebsiteSectionID
	AdditionalProperties interface{} `json:"additional_properties,omitempty"`
	Description          string      `json:"description,omitempty"`
	Name                 string      `json:"name"`
	Parent               interface{} `json:"parent"`
	ParentID             string      `json:"parent_id"`
	Path                 string      `json:"path"`
	Type                 string      `json:"type"`
	Version              string      `json:"version"`
}

type Taxonomy struct {
	PrimarySection `json:"primary_section"`
	PrimarySite    `json:"primary_site"`
	Sections       []Section `json:"sections"`
}

type Source struct {
	System     string `json:"system"`
	Name       string `json:"name"`
	SourceType string `json:"source_type"`
}

type Owner struct {
	Sponsored bool   `json:"sponsored"`
	ID        string `json:"id"`
}
type Distributor struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory,omitempty"`
}

type WebsitesWebSiteSection struct {
	Path   string `json:"path"`
	Parent struct {
		Default string `json:"default,omitempty"`
	} `json:"parent"`
	Website              string      `json:"_website"`
	ParentID             string      `json:"parent_id"`
	Name                 string      `json:"name"`
	Description          string      `json:"description,omitempty"`
	ID                   string      `json:"_id"`
	AdditionalProperties interface{} `json:"additional_properties"`
	WebsiteSectionID     string      `json:"_website_section_id"`
	Type                 string      `json:"type"`
	Version              string      `json:"version"`
}
type Website struct {
	WebsiteURL string                 `json:"website_url"`
	Section    WebsitesWebSiteSection `json:"website_section"`
}

type Websites map[string]Website

type Article struct {
	ID          string      `json:"_id"`
	Headlines   interface{} `json:"headlines"`
	Source      `json:"source"`
	Taxonomy    `json:"taxonomy"`
	Owner       `json:"owner"`
	Distributor `json:"distributor"`
	CreatedDate string `json:"created_date"`
	PublishDate string `json:"publish_date"`
	Website     string `json:"website"`
	WebsiteURL  string `json:"website_url"`
	Websites    `json:"websites"`
}
