package templates_struct

type SearchResult struct{
	Title string `json:Title`
	Author string `json:Author`
	Year string `json:Year`
	ID string `json:ID`
}
