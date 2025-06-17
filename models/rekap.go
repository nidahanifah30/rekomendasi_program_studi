package models

type RekapSiswa struct {
	No          int    `gorm:"column:No;primaryKey"`
	Kelas       string `gorm:"column:Kelas"`
	JumlahSiswa int    `gorm:"column:Jumlah Siswa"`
	IPA         int    `gorm:"column:IPA"`
	IPS         int    `gorm:"column:IPS"`
}

func (RekapSiswa) TableName() string {
	return "data_seluruh_siswa"
}
