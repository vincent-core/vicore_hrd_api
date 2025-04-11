package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"vicore_hrd/app/rest"
	"vicore_hrd/modules/hrd/dto"
	"vicore_hrd/modules/hrd/entity"

	"github.com/sirupsen/logrus"
)

type hrdUsecase struct {
	logging       *logrus.Logger
	hrdMapper     entity.VicoreHRDMapper
	hrdRepository entity.VicoreHRDRepository
}

func NewHRDUseCase(hrdRepository entity.VicoreHRDRepository, logging *logrus.Logger, hrdMapper entity.VicoreHRDMapper) entity.VicoreHRDUseCase {
	return &hrdUsecase{
		logging:       logging,
		hrdRepository: hrdRepository,
		hrdMapper:     hrdMapper,
	}
}

func (su *hrdUsecase) OnLoginUserByEmailAndPasswordUseCase(Email string, Password string, BaseURL string) (res dto.ResponseDataUser, message string, err error) {
	kemploye, er11 := su.hrdRepository.FindHRDByEmailRepository(Email)

	if er11 != nil || kemploye.Email == "" {
		return res, "Email tidak ditemukan", er11
	}

	// CEKK APAKAH KARYAWAN DOKTER ATAU PERAWAT
	// if kemploye.KeteranganPerson != "DOKTER" && kemploye.KeteranganPerson != "PERAWAT" {
	// 	return res, "Aplikasi hanya dapat diakses oleh Dokter dan Perawat", errors.New("Error App")
	// }

	// CEK PASSWORD APAKAH SESUAI DENGAN DATA
	var sha = sha1.New()
	sha.Write([]byte(Password))
	var encrypted = sha.Sum(nil)
	var passwordStr = fmt.Sprintf("%x", encrypted)

	if passwordStr == kemploye.Password {
		token, _ := rest.GenerateTokenUser(kemploye.Email, kemploye.KeteranganPerson, kemploye.KodePelayanan, kemploye.IDK)
		pelayanan, _ := su.hrdRepository.OnFindPelayananRepository(kemploye.KodePelayanan)

		user := dto.ResponseDataUser{
			Nama:         kemploye.Nama,
			KdBagian:     kemploye.KodePelayanan,
			Bagian:       pelayanan.Bagian,
			Token:        token["token"],
			RefreshToken: token["refresh_token"],
			KetPerson:    kemploye.KeteranganPerson,
			Photo:        toUserPhotoProfile(kemploye.JenisKelamin, kemploye.KeteranganPerson, kemploye.Photo, BaseURL),
		}

		return user, "Login berhasil", nil
	}

	if passwordStr != Password {
		return res, "Password salah", errors.New("Error")
	}

	return res, "Login gagal", errors.New("Error App")
}

func toUserPhotoProfile(JenisKelamin string, UserPerson string, UserPicture string, BaseURL string) (UserProfile string) {
	if UserProfile == "" {
		if JenisKelamin == "Laki-Laki" && UserPerson == "DOKTER" {
			return BaseURL + os.Getenv("USER_IMAGES") + "dokter_pria.png"
		}

		if JenisKelamin == "Perempuan" && UserPerson == "DOKTER" {
			return BaseURL + os.Getenv("USER_IMAGES") + "dokter_wanita.png"
		}
	}

	if UserProfile != "" {
		return BaseURL + os.Getenv("USER_IMAGES") + UserPicture

	}

	return BaseURL + os.Getenv("USER_IMAGES") + "dokter_pria.png"
}
