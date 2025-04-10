package dto

import (
	"vicore_hrd/modules/asesmen"
	asesmenDTO "vicore_hrd/modules/asesmen/dto"
)

type (
	ReponseLembarKonsul struct {
		NamaPasien            string                     `json:"nama_pasien"`
		TanggalLahir          string                     `json:"tanggal_lahir"`
		NomorRekamMedis       string                     `json:"no_rm"`
		NoReg                 string                     `json:"noreg"`
		Ruangan               string                     `json:"ruangan"`
		Dokter                string                     `json:"dokter"`
		JenisKonsultasi       string                     `json:"jenis_konsultasi"`
		DokterMemintaKonsul   string                     `json:"dokter_meminta_konsul"`
		Tanggal               string                     `json:"tanggal_konsul"`
		MohonKonulstasiPasien string                     `json:"mohon_konsul_pasien"`
		Umur                  string                     `json:"umur_pasien"`
		IktisarKlinik         string                     `json:"iktisar_klinik"`
		DiagnosaKerja         []asesmen.DiagnosaResponse `json:"diagnosa_kerja"`
	}

	ReponseLembarKonsulV2 struct {
		DokterDpjp            string                     `json:"dokter_dpjp"`
		NamaPasien            string                     `json:"nama_pasien"`
		TanggalLahir          string                     `json:"tanggal_lahir"`
		NomorRekamMedis       string                     `json:"no_rm"`
		NoReg                 string                     `json:"noreg"`
		Ruangan               string                     `json:"ruangan"`
		MohonKonulstasiPasien string                     `json:"mohon_konsul_pasien"`
		Umur                  string                     `json:"umur_pasien"`
		DiagnosaKerja         []asesmen.DiagnosaResponse `json:"diagnosa_kerja"`
		KonsulanDokter        []KonsulanDokter           `json:"konsulan_dokter"`
	}

	KonsulanDokter struct {
		DokterKonsul        string `json:"dokter_konsul"`
		DokterMemintaKonsul string `json:"dokter_meminta_konsul"`
		Spesialisasi        string `json:"spesialisasi"`
		KonsuleKe           int    `json:"konsule_ke"`
		JenisKonsultasi     string `json:"jenis_konsultasi"`
		Tanggal             string `json:"tanggal_konsul"`
		IktisarKlinik       string `json:"iktisar_klinik"`
		Ruangan             string `json:"ruangan"`
	}

	DokterKonsule struct {
		KonsuleKe    int    `json:"konsule_ke"`
		NamaDokter   string `json:"nama_dokter"`
		Spesialisasi string `json:"spesialisasi"`
	}

	JenisKonsule struct {
		KonsulKe    int    `json:"konsul_ke"`
		JenisKonsul string `json:"jenis_konsul"`
	}

	JawabanKonsul struct {
		KonsulKe int    `json:"konsul_ke"`
		Tanggal  string `json:"tanggal"`
		Penemuan string `json:"penemuan"`
		// Diagnosa   []asesmen.DiagnosaResponse `json:"diagnosa"`
		Terapi     string `json:"terapi"`
		Anjuran    string `json:"anjuran"`
		NamaDokter string `json:"nama_dokter"`
	}

	DataReponseLembarKonsule struct {
		ProfilePasien   asesmenDTO.DataProfilePasien `json:"profil_pasien"`
		LembarKonsul    ReponseLembarKonsul          `json:"lembar_konsul"`
		JawabanKonsulen JawabanKonsul                `json:"jawaban_konsul"`
	}

	DataReponseLembarKonsuleV2 struct {
		ProfilePasien   asesmenDTO.DataProfilePasien `json:"profil_pasien"`
		LembarKonsul    ReponseLembarKonsulV2        `json:"lembar_konsul"`
		JawabanKonsulen []JawabanKonsul              `json:"jawaban_konsul"`
	}
)
