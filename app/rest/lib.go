package rest

import (
	"fmt"
	"time"
)

func FormatTanggalWaktu(isoTime string) string {
	t, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat parsing waktu:", err)
		return ""
	}

	lokasi, _ := time.LoadLocation("Asia/Jakarta") // Ganti dengan lokasi Anda
	return t.In(lokasi).Format("2 Januari 2006 15:04 WIB")
}

func FormatTanggal(isoTime string) string {
	t, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat parsing waktu:", err)
		return ""
	}

	lokasi, _ := time.LoadLocation("Asia/Jakarta") // Ganti dengan lokasi Anda
	return t.In(lokasi).Format("2 Januari 2006")
}

func UbahTanggalIndo(tanggal string) (string, error) {
	// Parse tanggal dari string ke objek time.Time
	t, err := time.Parse(time.RFC3339, tanggal)
	if err != nil {
		return "", err
	}

	// Buat map nama bulan dalam bahasa Indonesia
	namaBulan := map[int]string{
		1:  "Januari",
		2:  "Februari",
		3:  "Maret",
		4:  "April",
		5:  "Mei",
		6:  "Juni",
		7:  "Juli",
		8:  "Agustus",
		9:  "September",
		10: "Oktober",
		11: "November",
		12: "Desember",
	}

	// Format tanggal dalam bahasa Indonesia
	tanggalIndo := fmt.Sprintf("%d %s %d", t.Day(), namaBulan[int(t.Month())], t.Year())

	return tanggalIndo, nil
}

func UbahTanggalIndoAndTime(tanggal string) (string, error) {
	// Parse tanggal dari string ke objek time.Time
	t, err := time.Parse(time.RFC3339, tanggal)
	if err != nil {
		return "", err
	}

	// Buat map nama bulan dalam bahasa Indonesia
	namaBulan := map[int]string{
		1:  "Januari",
		2:  "Februari",
		3:  "Maret",
		4:  "April",
		5:  "Mei",
		6:  "Juni",
		7:  "Juli",
		8:  "Agustus",
		9:  "September",
		10: "Oktober",
		11: "November",
		12: "Desember",
	}

	// Format tanggal dalam bahasa Indonesia
	tanggalIndo := fmt.Sprintf("%d %s %d %d:%d:%d WIB", t.Day(), namaBulan[int(t.Month())], t.Year(), t.Hour(), t.Minute(), t.Second())

	return tanggalIndo, nil
}

func UbahJamIndo(tanggal string) (string, error) {
	// Parse tanggal dari string ke objek time.Time
	t, err := time.Parse(time.RFC3339, tanggal)
	if err != nil {
		return "", err
	}

	// Format tanggal dalam bahasa Indonesia
	tanggalIndo := fmt.Sprintf("%d:%d:%d WIB", t.Hour(), t.Minute(), t.Second())

	return tanggalIndo, nil
}

func FormatDateIndonesian(dateString string) string {
	// Parse the date string into a time.Time object
	t, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	// Format the date into Indonesian format
	formattedDate := t.Format("2 Januari 2006")

	return formattedDate
}
