package models

type ProgramStudi struct {
	KodeProdi string  `gorm:"column:kode;primaryKey" form:"kode"`
	NamaProdi string  `gorm:"column:nama_prodi" form:"nama_prodi"`
	N1        float64 `gorm:"column:N1" form:"N1"`
	N2        float64 `gorm:"column:N2" form:"N2"`
	N3        float64 `gorm:"column:N3" form:"N3"`
	N4        float64 `gorm:"column:N4" form:"N4"`
	N5        float64 `gorm:"column:N5" form:"N5"`
}

func (ProgramStudi) TableName() string {
	return "program_studi"
}
