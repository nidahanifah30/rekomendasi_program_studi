package main

import (
	"html/template"
	"rekomendasi_program_studi/config"
	"rekomendasi_program_studi/controllers"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()

	// Inisialisasi session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"js":  template.JSEscapeString,
		"currentYear": func() int {
			return time.Now().Year()
		},
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// endpoint untuk dashboard
	r.GET("/dashboard", controllers.ShowDashboard)

	// endpoint untuk siswa
	r.GET("/siswa", controllers.GetAllSiswaHTML)
	r.POST("/siswa/tambah", controllers.SimpanSiswa)
	r.POST("/siswa/edit/:nisn", controllers.EditSiswa)
	r.POST("/siswa/hapus/:nisn", controllers.HapusSiswa)

	// endpoint untuk program studi
	r.GET("/prodi", controllers.GetAllProgramStudiHTML)
	r.POST("/prodi/tambah", controllers.SimpanProgramStudi)
	r.POST("/prodi/edit/:kode", controllers.EditProgramStudi)
	r.POST("/prodi/hapus/:kode", controllers.HapusProgramStudi)

	// Route untuk perhitungan dan hasil
	r.GET("/perhitungan/gap/:nisn", controllers.ShowGapCalculation)
	r.GET("/perhitungan/cf-sf/:nisn", controllers.ShowCFSFCalculation)
	r.GET("/perhitungan/hasil/:nisn", controllers.ShowHasilRekomendasi)

	// endpoint untuk rekomendasi
	r.GET("/rekomendasi-html/:nisn", controllers.ShowHasilRekomendasi)

	// endpoint untuk data rekomendasi
	r.GET("/data-rekomendasi", controllers.ShowDataRekomendasi)

	r.Run(":8080")
}
