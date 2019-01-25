package response

//===百度IP普通查询模型
type LocationModel struct {
	Address string          `json:"address"`
	Content LocationContent `json:"content"`
	Status  int             `json:"status"`
}

type LocationContent struct {
	Address       string         `json:"address"`
	AddressDetail LocationDetail `json:"address_detail"`
	Point         LocationPoint  `json:"point"`
}

type LocationDetail struct {
	City         string `json:"city"`
	CityCode     int    `json:"city_code"`
	District     string `json:"district"`
	Province     string `json:"province"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
}

type LocationPoint struct {
	X string `json:"x"`
	Y string `json:"y"`
}
