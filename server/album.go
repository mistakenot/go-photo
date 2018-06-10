package server

// Album represents a collection of photos.
type Album struct {
	Name  string `json:name`
	Size  int    `json:size`
	Count int    `json:count`
}

// AlbumOverview gives you an overview of all accounts.
type AlbumOverview struct {
	Albums []Album `json:albums`
}
