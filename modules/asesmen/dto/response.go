package dto

import (
	"vicore_hrd/modules/asesmen"
	resumemedis "vicore_hrd/modules/resume_medis"
)

type (
	ReportAsesmenDokterIGD struct {
		ProfilePasien    DataProfilePasien                        `json:"profil_pasien"`
		Tanggal          string                                   `json:"tanggal"`
		KeluhanUtama     string                                   `json:"keluhan_utama"`
		PenyakitSekarang string                                   `json:"penyakit_sekarang"`
		PenyakitDahulu   string                                   `json:"penyakit_dahulu"`
		PenyakitKeluarga string                                   `json:"penyakit_keluarga"`
		Prognosis        string                                   `json:"prognosis"`
		Dokter           string                                   `json:"dokter"`
		ImageLokalis     string                                   `json:"image_lokalis"`
		Diagnosa         []asesmen.DiagnosaResponse               `json:"diagnosa"`
		Labor            []resumemedis.ResHasilLaborTableLama     `json:"labor"`
		Radiologi        []resumemedis.RegHasilRadiologiTabelLama `json:"radiologi"`
		Fiso             []resumemedis.RegHasilRadiologiTabelLama `json:"fisioterapi"`
		Gizi             []resumemedis.RegHasilRadiologiTabelLama `json:"gizi"`
		PemeriksaanFisik ResponsePemfisik                         `json:"pemeriksaan_fisik"`
		Planning         []InstruksiObat                          `json:"planning"`
		VitalSign        ResponseTandaVital                       `json:"vital_sign"`
		KonsulKe         string                                   `json:"konsul_ke"`
		Terapi           string                                   `json:"terapi"`
		CaraKeluar       string                                   `json:"cara_keluar"`
		CaraKeluarDetail string                                   `json:"cara_keluar_detail"`
	}

	ResponsePemfisik struct {
		Kepala              string `json:"kepala"`
		Mata                string `json:"mata"`
		THT                 string `json:"tht"`
		Mulut               string `json:"mulut"`
		Leher               string `json:"leher"`
		Dada                string `json:"dada"`
		Jantung             string `json:"jantung"`
		Paru                string `json:"paru"`
		Perut               string `json:"perut"`
		Hati                string `json:"hati"`
		Limpa               string `json:"limpa"`
		Ginjal              string `json:"ginjal"`
		ALatKelamin         string `json:"alat_kelamin"`
		AnggotaGerak        string `json:"anggota_gerak"`
		Refleks             string `json:"relfeks"`
		KekuatanOtot        string `json:"kekuatan_otot"`
		Kulit               string `json:"kulit"`
		KelenjarGetahBening string `json:"kelenjar_getah_bening"`
		RtVt                string `json:"rt_vt"`
	}

	PemeriksanFisikAwalMedis struct {
		E                   string
		V                   string
		M                   string
		Kesadaran           string
		Kepala              string `json:"kepala"`
		Mata                string `json:"mata"`
		THT                 string `json:"tht"`
		Mulut               string `json:"mulut"`
		Leher               string `json:"leher"`
		Dada                string `json:"dada"`
		Jantung             string `json:"jantung"`
		Paru                string `json:"paru"`
		Perut               string `json:"perut"`
		Hati                string `json:"hati"`
		Limpa               string `json:"limpa"`
		Ginjal              string `json:"ginjal"`
		ALatKelamin         string `json:"alat_kelamin"`
		AnggotaGerak        string `json:"anggota_gerak"`
		Refleks             string `json:"relfeks"`
		KekuatanOtot        string `json:"kekuatan_otot"`
		Kulit               string `json:"kulit"`
		KelenjarGetahBening string `json:"kelenjar_getah_bening"`
		RtVt                string `json:"rt_vt"`
	}

	ResponseTandaVital struct {
		GCS        string `json:"gcs"`
		TD         string `json:"td"`
		Nadi       string `json:"nadi"`
		Suhu       string `json:"suhu"`
		Kesadaran  string `json:"kesadaran"`
		Pernafasan string `json:"pernafasan"`
		SPO2       string `json:"spo2"`
	}

	ReportPengantarRawatInap struct {
		Tanggal              string                     `json:"tanggal"`
		Bagian               string                     `json:"bagian"`
		Mohon                string                     `json:"mohon"`
		NamaPasien           string                     `json:"nama_pasien"`
		JenisKelamin         string                     `json:"jenis_kelamin"`
		KeluhanUtama         string                     `json:"keluhan_utama"`
		TanggalLahir         string                     `json:"tgl_lahir"`
		Alamat               string                     `json:"alamat"`
		NomorRM              string                     `json:"no_rm"`
		NamaDPJP             string                     `json:"nama_dpjp"`
		Diagnosa             []asesmen.DiagnosaResponse `json:"diagnosa"`
		PemeriksaanFisik     PemeriksaanFisik           `json:"pemeriksaan_fisik"`
		DokterPenangungJawab string                     `json:"dokter_penangung_jawab"`
		InstruksiObat        []InstruksiObat            `json:"instruksi"`
		InstruksiNarasi      string                     `json:"instruksi_narasi"`
	}

	PemeriksaanFisik struct {
		Sens         string `json:"sens"`
		E            string `json:"e"`
		M            string `json:"m"`
		V            string `json:"v"`
		TekananDarah string `json:"tekanan_darah"`
		RR           string `json:"rr"`
		Temp         string `json:"temp"`
		Hr           string `json:"hr"`
	}
)
