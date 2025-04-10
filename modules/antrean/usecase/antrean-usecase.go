package usecase

import (
	"fmt"
	"time"
	"vicore_hrd/modules/antrean/dto"
	"vicore_hrd/modules/antrean/entity"
	asesmenEntity "vicore_hrd/modules/asesmen/entity"
	hrdEnitty "vicore_hrd/modules/hrd/entity"
	libEntity "vicore_hrd/modules/lib/entity"

	"github.com/sirupsen/logrus"
)

type antreanUseCase struct {
	logging           *logrus.Logger
	antreanMapper     entity.AntreanMapper
	antreanRepository entity.AntreanRepository
	AsesmenRepository asesmenEntity.AsesmenRepository
	LibRepository     libEntity.LibRepository
	hrdRepository     hrdEnitty.VicoreHRDRepository
}

func NewAntreanUseCase(antreanRepository entity.AntreanRepository, logging *logrus.Logger, antreanMapper entity.AntreanMapper, asesemenRepo asesmenEntity.AsesmenRepository, libRepo libEntity.LibRepository, hrdRepo hrdEnitty.VicoreHRDRepository) entity.AntreanUseCase {
	return &antreanUseCase{
		logging:           logging,
		antreanMapper:     antreanMapper,
		antreanRepository: antreanRepository,
		AsesmenRepository: asesemenRepo,
		LibRepository:     libRepo,
		hrdRepository:     hrdRepo,
	}
}

func (iu *antreanUseCase) OnGetAntrianIGDUseCase(modulID string, person string, userID string) (res []dto.AntrianPasien, message string, err error) {
	switch modulID {
	case "IGD001":
		if person == "DOKTER" {
			var antrianpasien = []dto.AntrianPasien{}

			antrianDokterUmum, er12 := iu.antreanRepository.GetAntrianIGDDokterUmumRepository(userID)

			if er12 != nil {
				return make([]dto.AntrianPasien, 0), "Data tidak ditemukan", er12
			}

			// JIKA ANTRIAN PASIEN KOSONG
			if len(antrianDokterUmum) == 0 {
				return make([]dto.AntrianPasien, 0), "Data kosong", nil
			}

			// LOOPING ANTRIAN DOKTER UMUM
			for i := 0; i <= len(antrianDokterUmum)-1; i++ {
				asesmenPerawat, _ := iu.AsesmenRepository.OnGetPengkajianKeperawatanRepository(modulID, "RAJAL", antrianDokterUmum[i].Noreg)
				asesmenDokter, _ := iu.AsesmenRepository.OnGeAsesmenDokterRepository(antrianDokterUmum[i].Noreg, modulID)

				antrianpasien = append(antrianpasien, dto.AntrianPasien{
					Tgllahir:       tanggalIndoFromISO(antrianDokterUmum[i].Tgllahir),
					NoAntrean:      antrianDokterUmum[i].NoAntrian,
					JenisKelamin:   antrianDokterUmum[i].Jeniskelamin,
					Debitur:        "-",
					KodeDebitur:    "-",
					Noreg:          antrianDokterUmum[i].Noreg,
					Mrn:            antrianDokterUmum[i].Id,
					Keterangan:     "-",
					NamaPasien:     antrianDokterUmum[i].Nama,
					KdBag:          "IGD001",
					Bagian:         "Instalasi Gawat Darurat",
					Pelayanan:      "RAJAL",
					NamaDokter:     antrianDokterUmum[i].Dokter,
					KdDokter:       antrianDokterUmum[i].Kodedr,
					Kamar:          "-",
					Kasur:          "-",
					AsesmenDokter:  asesmenDokter.Dokter.Namadokter,
					AsesmenPerawat: asesmenPerawat.Perawat.Namaperawat,
				})
			}

			return antrianpasien, "OK", nil
		}

		antarinUGD, er12 := iu.antreanRepository.GetAntrianUGD()

		if er12 != nil {
			return make([]dto.AntrianPasien, 0), er12.Error(), er12
		}

		var antrianpasien = []dto.AntrianPasien{}

		if len(antarinUGD) == 0 {
			return make([]dto.AntrianPasien, 0), "Data tidak ditemukan", nil
		}

		for i := 0; i <= len(antarinUGD)-1; i++ {
			asesmenPerawat, _ := iu.AsesmenRepository.OnGetPengkajianKeperawatanRepository(modulID, "RAJAL", antarinUGD[i].Noreg)
			asesmenDokter, _ := iu.AsesmenRepository.OnGeAsesmenDokterRepository(antarinUGD[i].Noreg, modulID)

			antrianpasien = append(antrianpasien, dto.AntrianPasien{
				Tgllahir:       tanggalIndoFromISO(antarinUGD[i].Tgllahir),
				NoAntrean:      antarinUGD[i].NoAntrian,
				JenisKelamin:   antarinUGD[i].Jeniskelamin,
				Debitur:        "-",
				KodeDebitur:    "-",
				Noreg:          antarinUGD[i].Noreg,
				Mrn:            antarinUGD[i].Id,
				Keterangan:     "-",
				NamaPasien:     antarinUGD[i].Nama,
				KdBag:          "IGD001",
				Bagian:         "Instalasi Gawat Darurat",
				Pelayanan:      "RAJAL",
				NamaDokter:     antarinUGD[i].Dokter,
				KdDokter:       antarinUGD[i].Kodedr,
				Kamar:          "-",
				Kasur:          "-",
				AsesmenDokter:  asesmenDokter.Dokter.Namadokter,
				AsesmenPerawat: asesmenPerawat.Perawat.Namaperawat,
			})
		}

		return antrianpasien, "OK", nil
	default:
		if person == "DOKTER" {
			antrians, err := iu.antreanRepository.GetPasienBangsalForDokter(modulID, userID)

			if err != nil {
				return make([]dto.AntrianPasien, 0), "Data tidak ditemukan", err
			}

			if len(antrians) == 0 {
				return make([]dto.AntrianPasien, 0), "Data kosong", nil
			}

			var antrianpasien = []dto.AntrianPasien{}
			bagian, _ := iu.hrdRepository.OnFindPelayananRepository(modulID)

			// LOOPING DATA JIKA DITEMUKAN
			for i := 0; i <= len(antrians)-1; i++ {
				asesmenDokter, _ := iu.AsesmenRepository.OnGeAsesmenDokterRepository(antrians[i].Noreg, modulID)
				asesmenPerawat, _ := iu.AsesmenRepository.OnGetPengkajianKeperawatanRepository(modulID, "RANAP", antrians[i].Noreg)
				antrianpasien = append(antrianpasien, dto.AntrianPasien{
					Tgllahir:       tanggalIndoFromISO(antrians[i].Tgllahir),
					NoAntrean:      "-",
					JenisKelamin:   antrians[i].Sex,
					Debitur:        "-",
					KodeDebitur:    "-",
					Noreg:          antrians[i].Noreg,
					Mrn:            antrians[i].Id,
					Keterangan:     "-",
					NamaPasien:     antrians[i].Nama,
					KdBag:          bagian.KdBag,
					Bagian:         bagian.Bagian,
					Pelayanan:      "RANAP",
					NamaDokter:     antrians[i].Dokter,
					KdDokter:       antrians[i].Kodedr,
					Kamar:          "-",
					Kasur:          "-",
					AsesmenDokter:  asesmenDokter.Dokter.Namadokter,
					AsesmenPerawat: asesmenPerawat.Perawat.Namaperawat,
				})
			}

			return antrianpasien, "OK", nil
		}

		// JIKA TIDAK DOKTER
		antrianBangsal, er122 := iu.antreanRepository.GetPasienBangsal(modulID)

		if er122 != nil {
			return make([]dto.AntrianPasien, 0), "Data tidak ditemukan", er122
		}

		if len(antrianBangsal) == 0 {
			return make([]dto.AntrianPasien, 0), "Data kosong", nil
		}

		var antrianpasien = []dto.AntrianPasien{}
		bagian, _ := iu.hrdRepository.OnFindPelayananRepository(modulID)

		for i := 0; i <= len(antrianBangsal)-1; i++ {
			asesmenDokter, _ := iu.AsesmenRepository.OnGetAsesmenDokterRANAPRepository(antrianBangsal[i].Noreg, modulID)
			asesmenPerawat, _ := iu.AsesmenRepository.OnGetPengkajianKeperawatanRepository(modulID, "RANAP", antrianBangsal[i].Noreg)

			antrianpasien = append(antrianpasien, dto.AntrianPasien{
				Tgllahir:       tanggalIndoFromISO(antrianBangsal[i].Tgllahir),
				NoAntrean:      "-",
				JenisKelamin:   antrianBangsal[i].Sex,
				Debitur:        "-",
				KodeDebitur:    "-",
				Noreg:          antrianBangsal[i].Noreg,
				Mrn:            antrianBangsal[i].Id,
				Keterangan:     "-",
				NamaPasien:     antrianBangsal[i].Nama,
				KdBag:          bagian.KdBag,
				Bagian:         bagian.Bagian,
				Pelayanan:      "RANAP",
				NamaDokter:     antrianBangsal[i].Dokter,
				KdDokter:       antrianBangsal[i].Kodedr,
				Kamar:          antrianBangsal[i].Kamar,
				Kasur:          antrianBangsal[i].Kasur,
				AsesmenDokter:  asesmenDokter.Dokter.Namadokter,
				AsesmenPerawat: asesmenPerawat.Perawat.Namaperawat,
			})
		}

		return antrianpasien, "OK", nil
	}
}

func (iu *antreanUseCase) OnDashboardUseCase(modulID string) (res dto.ResponseDashboard, err error) {

	if modulID == "IGD001" {
		data, _ := iu.antreanRepository.GetAntrianUGD()

		return dto.ResponseDashboard{
				Jumlah: len(data),
			},
			nil
	}

	if modulID != "IGD001" {
		bangsal, _ := iu.antreanRepository.GetPasienBangsal(modulID)

		return dto.ResponseDashboard{
				Jumlah: len(bangsal),
			},
			nil
	}

	return dto.ResponseDashboard{}, nil
}

func tanggalIndoFromISO(isoTime string) string {

	t, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return ""
	}

	// Definisi bulan dalam bahasa Indonesia
	bulan := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	// Format tanggal: 26 Juni 2020
	return fmt.Sprintf("%d %s %d", t.Day(), bulan[t.Month()-1], t.Year())
}
