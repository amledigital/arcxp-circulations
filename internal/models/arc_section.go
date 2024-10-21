package models

type QResult struct {
	QResults []ArcSection `json:"q_results,omitempty"`
}

type ArcSection struct {
	ID         string               `json:"_id,omitempty"`
	Site       ArcSectionSite       `json:"site,omitempty"`
	SiteTopper ArcSectionSiteTopper `json:"site_topper,omitempty"`
	Social     ArcSectionSocial     `json:"social,omitempty"`
	Navigation ArcSectionNavigation `json:"navigation,omitempty"`
	Admin      ArcSectionAdmin      `json:"_admin,omitempty"`
	Website    string               `json:"_website,omitempty"`
	Name       string               `json:"name,omitempty"`
	Order      ArcSectionOrder      `json:"order,omitempty"`
	Parent     ArcSectionParent     `json:"pareant,omitempty"`
	Ancestors  ArcSectionAncestors  `json:"ancestors,omitempty"`
	Inactive   bool                 `json:"inactive,omitempty"`
	NodeType   string               `json:"node_type,omitempty"`
}

type ArcSectionSite struct {
	SiteTitle                    string `json:"site_title,omitempty"`
	SiteKeywords                 string `json:"site_keywords,omitempty"`
	SiteURL                      string `json:"string,omitempty"`
	SiteAbout                    string `json:"site_about,omitempty"`
	PageBuilderPathForNativeApps string `json:"pagebuilder_path_for_native_apps,omitempty"`
	SiteTagline                  string `json:"site_tagline,omitempty"`
	SiteDescription              string `json:"site_description"`
}

type ArcSectionSiteTopper struct {
	SiteLogoImage string `json:"site_logo_image"`
}

type ArcSectionSocial struct {
	RSS       string `json:"rss,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Facebook  string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
}

type ArcSectionNavigation struct {
	NavTitle string `json:"nav_title,omitempty"`
}

type ArcSectionAdmin struct {
	AliasIDs []string `json:"alias_ids,omitempty"`
}

type ArcSectionOrder struct {
	Footer   int `json:"footer,omitempty"`
	Default  int `json:"default,omitempty"`
	MegaMenu int `json:"mega-menu,omitempty"`
	Testing  int `json:"testing,omitempty"`
}

type ArcSectionParent struct {
	Default  string `json:"default,omitempty"`
	Footer   string `json:"footer,omitempty"`
	MegaMenu string `json:"mega-menu,omitempty"`
	Testing  string `json:"testing,omitempty"`
}

type ArcSectionAncestors struct {
	Default  []string `json:"default,omitempty"`
	Footer   []string `json:"footer,omitempty"`
	MegaMenu []string `json:"mega-menu,omitempty"`
	Testing  []string `json:"testing,omitempty"`
}
