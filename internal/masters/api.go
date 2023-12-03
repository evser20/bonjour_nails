package masters

type APIConfig struct {
	MastersReader MastersReader
}

type API struct {
	mastersReader MastersReader
}

func NewAPI(c APIConfig) (*API, error) {
	api := API{
		mastersReader: c.MastersReader,
	}
	return &api, nil
}

func (api *API) GetMasters() ([]Master, error) {
	return api.mastersReader.GetMasters()
}
