package models

type Siswa struct {
	NISN      uint    `gorm:"column:nisn;primaryKey"` // Sesuaikan nama kolom
	NamaSiswa string  `gorm:"column:nama_siswa"`
	Jurusan   string  `gorm:"column:jurusan"`
	N1        float64 `gorm:"column:N1"`
	N2        float64 `gorm:"column:N2"`
	N3        float64 `gorm:"column:N3"`
	N4        float64 `gorm:"column:N4"`
	N5        float64 `gorm:"column:N5"`
}

func (Siswa) TableName() string {
	return "data_siswa"
}
