package models

type ContentApiFetchResultAdditionalProperties struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
}

type ContentApiFetchResult struct {
	Type                 string                                    `json:"type"`
	Version              string                                    `json:"version"`
	ContentElements      []Article                                 `json:"content_elements"`
	AdditionalProperties ContentApiFetchResultAdditionalProperties `json:"additional_properties"`
	Count                int                                       `json:"count"`
	Next                 int                                       `json:"next"`
}
