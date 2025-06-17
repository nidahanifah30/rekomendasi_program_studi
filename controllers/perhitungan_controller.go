package controllers

import (
	"math"
	"rekomendasi_program_studi/config"
	"rekomendasi_program_studi/models"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Fungsi helper untuk konversi GAP ke bobot sesuai dokumen
func gapToWeight(gap float64) float64 {
	switch gap {
	case 0:
		return 6
	case 1, -1:
		return 5.5
	case 2, -2:
		return 5
	case 3, -3:
		return 4.5
	case 4, -4:
		return 4
	case 5, -5:
		return 3.5
	default:
		if gap > 5 || gap < -5 {
			return 3
		}
		return 3 // default value
	}
}

func ShowGapCalculation(c *gin.Context) {
	nisn := c.Param("nisn")
	var siswa models.Siswa
	var prodiList []models.ProgramStudi

	// Ambil data siswa dan prodi
	config.DB.Where("NISN = ?", nisn).First(&siswa)
	config.DB.Find(&prodiList)

	// Hitung GAP untuk semua prodi
	type GapResult struct {
		Prodi     models.ProgramStudi
		N1        models.ProgramStudi
		N2        models.ProgramStudi
		N3        models.ProgramStudi
		N4        models.ProgramStudi
		N5        models.ProgramStudi
		GapN1     float64
		GapN2     float64
		GapN3     float64
		GapN4     float64
		GapN5     float64
		RoundedN1 float64
		RoundedN2 float64
		RoundedN3 float64
		RoundedN4 float64
		RoundedN5 float64
		BobotN1   float64
		BobotN2   float64
		BobotN3   float64
		BobotN4   float64
		BobotN5   float64
	}

	var results []GapResult
	for _, prodi := range prodiList {
		// Hitung GAP sesuai rumus dari dokumen
		gapN1 := (siswa.N1 - prodi.N1) / 10
		gapN2 := (siswa.N2 - prodi.N2) / 10
		gapN3 := (siswa.N3 - prodi.N3) / 10
		gapN4 := (siswa.N4 - prodi.N4) / 10
		gapN5 := (siswa.N5 - prodi.N5) / 10

		// Pembulatan GAP ke bilangan bulat terdekat
		roundedN1 := math.Round(gapN1)
		roundedN2 := math.Round(gapN2)
		roundedN3 := math.Round(gapN3)
		roundedN4 := math.Round(gapN4)
		roundedN5 := math.Round(gapN5)

		// Konversi ke bobot sesuai tabel dari dokumen
		bobotN1 := gapToWeight(roundedN1)
		bobotN2 := gapToWeight(roundedN2)
		bobotN3 := gapToWeight(roundedN3)
		bobotN4 := gapToWeight(roundedN4)
		bobotN5 := gapToWeight(roundedN5)

		results = append(results, GapResult{
			Prodi:     prodi,
			GapN1:     gapN1,
			GapN2:     gapN2,
			GapN3:     gapN3,
			GapN4:     gapN4,
			GapN5:     gapN5,
			RoundedN1: roundedN1,
			RoundedN2: roundedN2,
			RoundedN3: roundedN3,
			RoundedN4: roundedN4,
			RoundedN5: roundedN5,
			BobotN1:   bobotN1,
			BobotN2:   bobotN2,
			BobotN3:   bobotN3,
			BobotN4:   bobotN4,
			BobotN5:   bobotN5,
		})
	}

	c.HTML(200, "perhitungan_gap.html", gin.H{
		"Page":    "gap",
		"Siswa":   siswa,
		"Results": results,
	})
}

func ShowCFSFCalculation(c *gin.Context) {
	nisn := c.Param("nisn")
	var siswa models.Siswa
	var prodiList []models.ProgramStudi

	config.DB.Where("NISN = ?", nisn).First(&siswa)
	config.DB.Find(&prodiList)

	type CFSFResult struct {
		Prodi      models.ProgramStudi
		CF         float64
		SF         float64
		NilaiAkhir float64
	}

	var results []CFSFResult
	for _, prodi := range prodiList {
		// Hitung GAP dan pembulatan
		gapN1 := (siswa.N1 - prodi.N1) / 10
		gapN2 := (siswa.N2 - prodi.N2) / 10
		gapN3 := (siswa.N3 - prodi.N3) / 10
		gapN4 := (siswa.N4 - prodi.N4) / 10
		gapN5 := (siswa.N5 - prodi.N5) / 10

		roundedN1 := math.Round(gapN1)
		roundedN2 := math.Round(gapN2)
		roundedN3 := math.Round(gapN3)
		roundedN4 := math.Round(gapN4)
		roundedN5 := math.Round(gapN5)

		// Konversi ke bobot
		bobotN1 := gapToWeight(roundedN1)
		bobotN2 := gapToWeight(roundedN2)
		bobotN3 := gapToWeight(roundedN3)
		bobotN4 := gapToWeight(roundedN4)
		bobotN5 := gapToWeight(roundedN5)

		// Hitung CF (Core Factor) dan SF (Secondary Factor) sesuai dokumen
		// CF = (N2 + N3 + N5) / 3
		// SF = (N1 + N4) / 2
		cf := (bobotN2 + bobotN3 + bobotN5) / 3
		sf := (bobotN1 + bobotN4) / 2

		// Hitung Nilai Akhir = 60% CF + 40% SF
		nilaiAkhir := 0.6*cf + 0.4*sf

		results = append(results, CFSFResult{
			Prodi:      prodi,
			CF:         cf,
			SF:         sf,
			NilaiAkhir: nilaiAkhir,
		})
	}

	c.HTML(200, "perhitungan_cfsf.html", gin.H{
		"Page":    "cf-sf",
		"Siswa":   siswa,
		"Results": results,
	})
}

type HasilRekomendasi struct {
	KodeProdi string
	NamaProdi string
	GapN1     float64
	GapN2     float64
	GapN3     float64
	GapN4     float64
	GapN5     float64
	BobotN1   float64
	BobotN2   float64
	BobotN3   float64
	BobotN4   float64
	BobotN5   float64
	CF        float64
	SF        float64
	Total     float64
}

func ShowHasilRekomendasi(c *gin.Context) {
	nisn := c.Param("nisn")

	// Get student data
	var siswa models.Siswa
	if err := config.DB.Where("NISN = ?", nisn).First(&siswa).Error; err != nil {
		c.String(404, "Data siswa tidak ditemukan")
		return
	}

	// Get all study programs
	var prodiList []models.ProgramStudi
	config.DB.Find(&prodiList)

	var hasil []HasilRekomendasi

	for _, prodi := range prodiList {
		// Calculate GAP
		gapN1 := (siswa.N1 - prodi.N1) / 10
		gapN2 := (siswa.N2 - prodi.N2) / 10
		gapN3 := (siswa.N3 - prodi.N3) / 10
		gapN4 := (siswa.N4 - prodi.N4) / 10
		gapN5 := (siswa.N5 - prodi.N5) / 10

		// Round GAP values
		roundedN1 := math.Round(gapN1)
		roundedN2 := math.Round(gapN2)
		roundedN3 := math.Round(gapN3)
		roundedN4 := math.Round(gapN4)
		roundedN5 := math.Round(gapN5)

		// Convert to weights
		bobotN1 := gapToWeight(roundedN1)
		bobotN2 := gapToWeight(roundedN2)
		bobotN3 := gapToWeight(roundedN3)
		bobotN4 := gapToWeight(roundedN4)
		bobotN5 := gapToWeight(roundedN5)

		// Calculate CF and SF
		cf := (bobotN2 + bobotN3 + bobotN5) / 3
		sf := (bobotN1 + bobotN4) / 2

		// Calculate total score
		total := 0.6*cf + 0.4*sf

		hasil = append(hasil, HasilRekomendasi{
			KodeProdi: prodi.KodeProdi,
			NamaProdi: prodi.NamaProdi,
			GapN1:     gapN1,
			GapN2:     gapN2,
			GapN3:     gapN3,
			GapN4:     gapN4,
			GapN5:     gapN5,
			BobotN1:   bobotN1,
			BobotN2:   bobotN2,
			BobotN3:   bobotN3,
			BobotN4:   bobotN4,
			BobotN5:   bobotN5,
			CF:        math.Round(cf*100) / 100,
			SF:        math.Round(sf*100) / 100,
			Total:     math.Round(total*1000) / 1000,
		})
	}

	// Sort by total score (descending)
	sort.SliceStable(hasil, func(i, j int) bool {
		return hasil[i].Total > hasil[j].Total
	})

	c.HTML(200, "rekomendasi.html", gin.H{
		"Page":  "rekomendasi",
		"Siswa": siswa,
		"Hasil": hasil,
	})
}

// ✅ Tambahan: ShowDataRekomendasi untuk menampilkan semua siswa dengan rekomendasi terbaiknya
func ShowDataRekomendasi(c *gin.Context) {
	var siswaList []models.Siswa
	config.DB.Find(&siswaList)

	var data []map[string]string

	for _, siswa := range siswaList {
		var prodiList []models.ProgramStudi
		config.DB.Find(&prodiList)

		var hasil []HasilRekomendasi
		for _, prodi := range prodiList {
			gapN1 := (siswa.N1 - prodi.N1) / 10
			gapN2 := (siswa.N2 - prodi.N2) / 10
			gapN3 := (siswa.N3 - prodi.N3) / 10
			gapN4 := (siswa.N4 - prodi.N4) / 10
			gapN5 := (siswa.N5 - prodi.N5) / 10

			bobotN1 := gapToWeight(math.Round(gapN1))
			bobotN2 := gapToWeight(math.Round(gapN2))
			bobotN3 := gapToWeight(math.Round(gapN3))
			bobotN4 := gapToWeight(math.Round(gapN4))
			bobotN5 := gapToWeight(math.Round(gapN5))

			cf := (bobotN2 + bobotN3 + bobotN5) / 3
			sf := (bobotN1 + bobotN4) / 2
			total := 0.6*cf + 0.4*sf

			hasil = append(hasil, HasilRekomendasi{
				KodeProdi: prodi.KodeProdi,
				NamaProdi: prodi.NamaProdi,
				Total:     math.Round(total*1000) / 1000,
			})
		}

		sort.SliceStable(hasil, func(i, j int) bool {
			return hasil[i].Total > hasil[j].Total
		})

		if len(hasil) > 0 {
			data = append(data, map[string]string{
				"NISN":             strconv.FormatUint(uint64(siswa.NISN), 10),
				"NamaSiswa":        siswa.NamaSiswa,
				"ProdiRekomendasi": hasil[0].NamaProdi,
			})
		}
	}

	c.HTML(200, "data_rekomendasi.html", gin.H{
		"DataRekomendasi": data,
	})
}
