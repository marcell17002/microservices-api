package informations

import "time"

type Informations struct {
	ID             int
	IDUser         int
	Npm            string
	Prodi          string
	Ktp            string
	Pekerjaan      string
	JenisPekerjaan string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
