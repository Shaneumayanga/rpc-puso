package main

type API struct {
	db *Database
}

func NewAPI(db *Database) *API {
	return &API{
		db: db,
	}
}

func (api *API) Save(pusa Pusa, reply *Pusa) error {
	api.db.Puso = append(api.db.Puso, pusa)
	*reply = pusa
	return nil
}

func (api *API) DumpDB(empty string, reply *[]Pusa) error {
	*reply = api.db.Puso
	return nil
}
