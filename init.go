package main

import (
	"os"
	"runtime"
	"time"
	"vicore_hrd/config"
	"vicore_hrd/exception"
	"vicore_hrd/pkg/logs"

	hrdHandler "vicore_hrd/modules/hrd/handler"
	hrdMapper "vicore_hrd/modules/hrd/mapper"
	hrdRepository "vicore_hrd/modules/hrd/repository"
	hrdUseCase "vicore_hrd/modules/hrd/usecase"

	libHandler "vicore_hrd/modules/lib/handler"
	libMapper "vicore_hrd/modules/lib/mapper"
	libRepository "vicore_hrd/modules/lib/repository"

	antreanHandler "vicore_hrd/modules/antrean/handler"
	antreanMapper "vicore_hrd/modules/antrean/mapper"
	antreanRepo "vicore_hrd/modules/antrean/repository"
	antreanUsecase "vicore_hrd/modules/antrean/usecase"

	// ASESMEN
	asesmenHandler "vicore_hrd/modules/asesmen/handler"
	asesmenMapper "vicore_hrd/modules/asesmen/mapper"
	asesmenRepo "vicore_hrd/modules/asesmen/repository"

	asesmenUsecase "vicore_hrd/modules/asesmen/usecase"

	eduksaiHandler "vicore_hrd/modules/edukasi_terintegrasi/handler"
	mapperEdukasi "vicore_hrd/modules/edukasi_terintegrasi/mapper"
	repoEdukasi "vicore_hrd/modules/edukasi_terintegrasi/repository"
	edukasiUseCase "vicore_hrd/modules/edukasi_terintegrasi/usecase"

	generalHandler "vicore_hrd/modules/general_consent/handler"
	generalMapper "vicore_hrd/modules/general_consent/mapper"
	generalRepo "vicore_hrd/modules/general_consent/repository"
	generalUseCase "vicore_hrd/modules/general_consent/usecase"

	// RESUME MEDIS
	resumeMedisHandler "vicore_hrd/modules/resume_medis/handler"
	resumeMedisMapper "vicore_hrd/modules/resume_medis/mapper"
	resumeMedisRepo "vicore_hrd/modules/resume_medis/repository"
	resumeMedisUseCase "vicore_hrd/modules/resume_medis/usecase"

	// TRIASE
	triaseHandler "vicore_hrd/modules/triase/handler"
	triaseMapper "vicore_hrd/modules/triase/mapper"
	triaseRepository "vicore_hrd/modules/triase/repository"
	triaseUseCase "vicore_hrd/modules/triase/usecase"

	lembarKonsulHandler "vicore_hrd/modules/lembar_konsul/handler"
	lembarKonsulRepo "vicore_hrd/modules/lembar_konsul/repository"

	lembarMapper "vicore_hrd/modules/lembar_konsul/mapper"
	lembarUseCase "vicore_hrd/modules/lembar_konsul/usecase"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	HRDHandler          *hrdHandler.HRDHandler
	LibHandler          *libHandler.LibHandler
	AntreanHandler      *antreanHandler.AntreanHandler
	AsemsenHandler      *asesmenHandler.AsesmenHandler
	EdukasiHandler      *eduksaiHandler.EdukasiHandler
	GeneralHandler      *generalHandler.GeneralHandler
	Resumehandler       *resumeMedisHandler.ResumeHandler
	TriaseHandler       *triaseHandler.TriaseHandler
	LembarKonsulHandler *lembarKonsulHandler.LembarKonsulhandler
}

func RunApplication() {
	runtime.GOMAXPROCS(3)
	store := session.New()

	logging := logs.NewLogger()
	err := godotenv.Load(".env")

	if err != nil {
		logging.Error(err)
		log.Println(".env is not loaded properly")
		os.Exit(1)
	}

	db := config.InitMysqlDB()

	sqlDB, err := db.DB()

	if err != nil {
		logging.Info(err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour * 2)

	hrdMapper := hrdMapper.NewHRDMapperImple()
	antreanMapper := antreanMapper.NewAntreanMapperImple(logging)
	libMapper := libMapper.NewLibMapperImple()
	asesmenMapper := asesmenMapper.NewAsesmenMapperImple()
	mapperEdukasi := mapperEdukasi.NewEdukasiTerintegrasiMapper()
	resumeMapper := resumeMedisMapper.NewResumeMedisMapper()
	lembarMapper := lembarMapper.NewLembarKonsulMapperImpl()
	generalMapper := generalMapper.NewGeneralConsentMapper()
	triaseMapper := triaseMapper.NewTriaseMapper(logging)

	resumeRepo := resumeMedisRepo.NewResumeMedisRespository(db, logging)
	generalRepo := generalRepo.NewGeneralConsentRepository(db, logging)
	libRepo := libRepository.NewLibRepository(db, logging)
	antrianRepo := antreanRepo.NewAntreanRepository(db, logging)
	repoTriase := triaseRepository.NewTriaseRepository(db, logging)
	repoAsesmen := asesmenRepo.NewAsesmenRepository(db, logging)
	repoEdukasi := repoEdukasi.NewEdukasiTerintegrasi(db, logging)
	repoLembar := lembarKonsulRepo.NewLembarKonsulRepository(db, logging)

	generalUseCase := generalUseCase.NewGeneralUseCase(logging, generalRepo, generalMapper)
	resumeUsecase := resumeMedisUseCase.NewResumeUseCase(logging, resumeMapper, resumeRepo)
	repoHRD := hrdRepository.NewHisRepository(db, logging)
	hrdUseCase := hrdUseCase.NewHRDUseCase(repoHRD, logging, hrdMapper)

	antreanUseCase := antreanUsecase.NewAntreanUseCase(antrianRepo, logging, antreanMapper, repoAsesmen, libRepo, repoHRD)
	asesmenUseCase := asesmenUsecase.NewAsesmenUseCase(repoAsesmen, logging, asesmenMapper, resumeUsecase, resumeRepo, antrianRepo, repoHRD)
	edukasiUseCase := edukasiUseCase.NewEdukasiTerintegrasiUseCase(logging, mapperEdukasi, repoEdukasi)
	lembarUseCase := lembarUseCase.NewLembarKonsulUseCase(logging, repoLembar, lembarMapper, resumeRepo, repoHRD, repoAsesmen)
	triaseUseCase := triaseUseCase.NewTriaseUsecase(logging, triaseMapper, resumeRepo, repoTriase)

	lembarHandler := lembarKonsulHandler.LembarKonsulhandler{
		Logging:          logging,
		LembarKonsulRepo: repoLembar,
		LembarUseCase:    lembarUseCase,
	}
	eduksaiHandler := eduksaiHandler.EdukasiHandler{
		EdukasiRepository: repoEdukasi,
		EdukasiUseCase:    edukasiUseCase,
		EdukasiMapper:     mapperEdukasi,
		Logging:           logging,
		HRDRepository:     repoHRD,
	}

	hrdHandler := hrdHandler.HRDHandler{
		HRDRepository: repoHRD,
		HRDUseCase:    hrdUseCase,
		HRDMapper:     hrdMapper,
		Logging:       logging,
	}

	// ========= RESUME HANDLER
	resumeHandler := resumeMedisHandler.ResumeHandler{
		Logging:           logging,
		ResumerMapper:     resumeMapper,
		ResumeUseCase:     resumeUsecase,
		ResumeRepository:  resumeRepo,
		AntreanRepository: antrianRepo,
	}

	// ============== LIB HANDLER
	libHandler := libHandler.LibHandler{
		LibRepository: libRepo,
		LibMapper:     libMapper,
		Logging:       logging,
	}

	// ============== ANTREAN HANDLER
	antreanHandler := antreanHandler.AntreanHandler{
		AntreanRepository: antrianRepo,
		Logging:           logging,
		AntreanUseCase:    antreanUseCase,
		AntreanMapper:     antreanMapper,
	}

	asesmenHandler := asesmenHandler.AsesmenHandler{
		Logging:           logging,
		AsesmenRepository: repoAsesmen,
		AsesmenUseCase:    asesmenUseCase,
		AsesmenMapper:     asesmenMapper,
		ResumeRepoitory:   resumeRepo,
		GeneralRepository: generalRepo,
	}

	generalHandler := generalHandler.GeneralHandler{
		Logging:        logging,
		GeneralUsecase: generalUseCase,
		GeneralRepo:    generalRepo,
		GeneralMapper:  generalMapper,
	}

	triaseHandler := triaseHandler.TriaseHandler{
		Logging:       logging,
		TriaseMapper:  triaseMapper,
		TriaseUseCase: triaseUseCase,
		TriaseRepo:    repoTriase,
	}

	service := &Service{
		HRDHandler:          &hrdHandler,
		LibHandler:          &libHandler,
		AntreanHandler:      &antreanHandler,
		AsemsenHandler:      &asesmenHandler,
		EdukasiHandler:      &eduksaiHandler,
		GeneralHandler:      &generalHandler,
		Resumehandler:       &resumeHandler,
		TriaseHandler:       &triaseHandler,
		LembarKonsulHandler: &lembarHandler,
	}

	app := fiber.New(config.NewFiberConfig())

	app.Get("/dashboard", monitor.New())

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
		URL:  "/favicon.ico",
	}))

	app.Static("/app/images/users/", "./images/users/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Static("/app/images/public/", "./images/public/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Static("/app/images/nyeri/", "./images/nyeri/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Static("/app/images/general_consent/", "./images/general_consent/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Static("/app/images/general_consent", "./images/general_consent/", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	service.FiberRoutingResumeListen(logging, app, store)
	service.FiberLembarKonsulListen(logging, app, store)
	service.FiberTriaseRoutingListen(logging, app, store)
	service.FiberRoutingAsesmenListen(logging, app, store)
	service.FiberRoutingUserAndListen(logging, app, store)
	service.FiberRoutingPelayananListen(logging, app, store)
	service.FiberAntreanRouterListen(logging, app, store)
	service.FiberRuterCPPTV1(logging, app, store)
	service.FiberEdukasiTeringrasi(logging, app, store)
	service.FiberRoutingGeneralConsent(logging, app, store)

	err = app.Listen(os.Getenv("DEPLOY_PORT"))
	exception.PanicIfNeeded(err)
}
