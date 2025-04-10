package dto

import (
	"vicore_hrd/modules/asesmen/dto"
	resumemedis "vicore_hrd/modules/resume_medis"
)

type (
	ResponseResumeMedis struct {
		NamaPasien         string                                   `json:"nama_pasien"`
		NoRegister         string                                   `json:"no_register"`
		TanggalLahir       string                                   `json:"tanggal_lahir"`
		TanggalMasuk       string                                   `json:"tgl_masuk"`
		Alamat             string                                   `json:"alamat"`
		JamMasuk           string                                   `json:"jam_masuk"`
		TanggalKeluar      string                                   `json:"tgl_keluar"`
		JamKeluar          string                                   `json:"jam_keluar"`
		NomorRekamMedis    string                                   `json:"no_rm"`
		RiwayatPenyakit    string                                   `json:"riwayat_penyakit"`
		JenisKelamin       string                                   `json:"jenis_kelamin"`
		Ruang              string                                   `json:"ruang"`
		Pelayanan          string                                   `json:"pelayanan"`
		Kelas              string                                   `json:"kelas"`
		KeadaanWaktuKeluar string                                   `json:"keadaan"`
		DokterMerawat      string                                   `json:"dokter_merawat"`
		DataObat           []DataKeluarObat                         `json:"obat_waktu_pulang"`
		Labor              []resumemedis.ResHasilLaborTableLama     `json:"labor"`
		Radiologi          []resumemedis.RegHasilRadiologiTabelLama `json:"radiologi"`
		Fiso               []resumemedis.RegHasilRadiologiTabelLama `json:"fisioterapi"`
		Gizi               []resumemedis.RegHasilRadiologiTabelLama `json:"gizi"`
		PemeriksaanFisik   ResponsePemerikssanFisik                 `json:"pemeriksaan_fisik"`
		Diagnosa           []resumemedis.DiagnosaResponse           `json:"diagnosa"`
		Tindakan           []resumemedis.TindakanResponse           `json:"tindakan"`
	}

	DataKeluarObat struct {
		TglKeluar string `json:"tgl_keluar"`
		NamaObat  string `json:"nama_obat"`
		Jumlah    int    `json:"jumlah"`
	}

	ResponsePemerikssanFisik struct {
		Tb         string `json:"tb"`
		Td         string `json:"td"`
		Bb         string `json:"bb"`
		Nadi       string `json:"nadi"`
		Suhu       string `json:"suhu"`
		Spo2       string `json:"spo2"`
		Pernafasan string `json:"penafasan"`
	}

	DataResponseCPPT struct {
		DataCPPT []dto.ResponseCPPT `json:"cppt"`
		Pasien   ProfilePasien      `json:"profile_pasien"`
	}

	ProfilePasien struct {
		NoReg     string
		Nama      string
		KdBagian  string
		Pelayanan string
	}
)
