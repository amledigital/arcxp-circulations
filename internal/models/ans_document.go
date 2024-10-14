package models

type PrimarySection struct {
	ID                   string      `json:"_id,omitempty"`
	Website              string      `json:"_website,omitempty"`
	AdditionalProperties interface{} `json:"additional_properties"`
	Description          string      `json:"description,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Parent               interface{} `json:"parent,omitempty"`
	ParentID             string      `json:"parent_id,omitempty"`
	Path                 string      `json:"path,omitempty"`
	Type                 string      `json:"type,omitempty"`
	Version              string      `json:"version,omitempty"`
}

type PrimarySite struct {
	ID                   string      `json:"_id,omitempty"`
	AdditionalProperties interface{} `json:"additional_properties,omitempty"`
	Description          string      `json:"description,omitempty"`
	Name                 string      `json:"name,omitempty"`
	ParentID             string      `json:"parent_id,omitempty"`
	Path                 string      `json:"path,omitempty"`
	Type                 string      `json:"type,omitempty"`
	Version              string      `json:"version,omitempty"`
}

type Section struct {
	ID                   string      `json:"_id,omitempty"`
	Website              string      `json:"_website,omitempty"`
	WebsiteSectionID     string      `json:"_website_section_id,omitempty"` // concats website.WebsiteSectionID
	AdditionalProperties interface{} `json:"additional_properties,omitempty"`
	Description          string      `json:"description,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Parent               interface{} `json:"parent,omitempty"`
	ParentID             string      `json:"parent_id,omitempty"`
	Path                 string      `json:"path,omitempty"`
	Type                 string      `json:"type,omitempty"`
	Version              string      `json:"version,omitempty"`
}

type Taxonomy struct {
	PrimarySection `json:"primary_section,omitempty"`
	PrimarySite    `json:"primary_site,omitempty"`
	Sections       []Section `json:"sections,omitempty"`
}

type Source struct {
	System     string `json:"system,omitempty"`
	Name       string `json:"name,omitempty"`
	SourceType string `json:"source_type,omitempty"`
}

type Owner struct {
	Sponsored bool   `json:"sponsored,omitempty"`
	ID        string `json:"id,omitempty"`
}
type Distributor struct {
	Name        string `json:"name,omitempty"`
	Category    string `json:"category,omitempty"`
	SubCategory string `json:"subcategory,omitempty"`
}

type WebsitesWebSiteSection struct {
	Path   string `json:"path,omitempty"`
	Parent struct {
		Default string `json:"default,omitempty"`
	} `json:"parent,omitempty"`
	Website              string      `json:"_website,omitempty"`
	ParentID             string      `json:"parent_id,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Description          string      `json:"description,omitempty"`
	ID                   string      `json:"_id,omitempty"`
	AdditionalProperties interface{} `json:"additional_properties,omitempty"`
	WebsiteSectionID     string      `json:"_website_section_id,omitempty"`
	Type                 string      `json:"type,omitempty"`
	Version              string      `json:"version,omitempty"`
}
type Website struct {
	WebsiteURL string                 `json:"website_url,omitempty"`
	Section    WebsitesWebSiteSection `json:"website_section,omitempty"`
}

type Websites map[string]Website

type Article struct {
	ID          string      `json:"_id,omitempty"`
	Headlines   interface{} `json:"headlines,omitempty"`
	Source      `json:"source,omitempty"`
	Taxonomy    `json:"taxonomy,omitempty"`
	Owner       `json:"owner,omitempty"`
	Distributor `json:"distributor,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	PublishDate string `json:"publish_date,omitempty"`
	Website     string `json:"website,omitempty"`
	WebsiteURL  string `json:"website_url,omitempty"`
	Websites    `json:"websites,omitempty"`
}
