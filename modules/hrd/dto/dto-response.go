package dto

type (
	ResponseDataUser struct {
		Nama         string `json:"nama"`
		Bagian       string `json:"bagian"`
		KdBagian     string `json:"kd_bagian"`
		KetPerson    string `json:"ket_person"`
		Photo        string `json:"photo"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
