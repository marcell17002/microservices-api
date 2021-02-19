package informations

type GetInformationsDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateInformationsInput struct {
	IDUser         int    `json:"IDUser" binding:"required"`
	Npm            string `json:"npm" binding:"required"`
	Prodi          string `json:"prodi" binding:"required"`
	Ktp            string `json:"ktp" binding:"required"`
	Pekerjaan      string `json:"pekerjaan" binding:"required"`
	JenisPekerjaan string `json:"jenisPekerjaan" binding:"required"`
	Status         string `json:"status" binding:"required"`
}
