package seeds

import (
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/pkg/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateUser1",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Dmytro", "dmytroyakimuk@gmail.com", "123456", "admin")

			},
		},
		seed.Seed{
			Name: "CreateUser2",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "vova", "vova@gmail.com", "123456", "owner")
			},
		},
		seed.Seed{
			Name: "CreateUser3",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "roma", "roma@gmail.com", "123456", "client")
			},
		},

		seed.Seed{
			Name: "CreateCategory1",
			Run: func(db *gorm.DB) error {
				return CreateCategory(db, "Піші прогулянки", "pishi_prohulianky", "icon-hiking")
			},
		},

		seed.Seed{
			Name: "CreateCategory2",
			Run: func(db *gorm.DB) error {
				return CreateCategory(db, "Катання на гірських велосипедах", "katannia_na_hirskykh_velosypedakh", "icon-mountain-biking")
			},
		},

		seed.Seed{
			Name: "CreateCategory3",
			Run: func(db *gorm.DB) error {
				return CreateCategory(db, "Підводне плавання", "pidvodne_plavannia", "icon-scuba-diving")
			},
		},

		seed.Seed{
			Name: "CreateCategory4",
			Run: func(db *gorm.DB) error {
				return CreateCategory(db, "Риболовля", "rybolovlia", "icon-hunting")
			},
		},

		seed.Seed{
			Name: "CreateRegion1",
			Run: func(db *gorm.DB) error {
				return CreateRegion(db, "Волиньська", "volynska")
			},
		},

		seed.Seed{
			Name: "CreateRegion2",
			Run: func(db *gorm.DB) error {
				return CreateRegion(db, "Рівненська", "rivnenska")
			},
		},

		seed.Seed{
			Name: "CreateRegion3",
			Run: func(db *gorm.DB) error {
				return CreateRegion(db, "Львівська", "lvivska")
			},
		},

		seed.Seed{
			Name: "CreateTour1",
			Run: func(db *gorm.DB) error {
				return CreateTour(db,
					"Тур на вихідні в стилі активного відпочинку 1",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_1",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_1",
					"unpublic",
					1,
					2,
					1,
				)
			},
		},

		seed.Seed{
			Name: "CreateTour2",
			Run: func(db *gorm.DB) error {
				return CreateTour(db,
					"Тур на вихідні в стилі активного відпочинку 2",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					"unpublic",
					4,
					1,
					2,
				)
			},
		},

		seed.Seed{
			Name: "CreateImage1",
			Run: func(db *gorm.DB) error {
				return CreateImage(db, "test1.jpg", 1)
			},
		},

		seed.Seed{
			Name: "CreateImage2",
			Run: func(db *gorm.DB) error {
				return CreateImage(db, "test2.jpg", 1)
			},
		},
		seed.Seed{
			Name: "CreateTrip1",
			Run: func(db *gorm.DB) error {
				return CreateTrip(db, time.Date(2021, time.November, 10, 0, 0, 0, 0, time.UTC), time.Date(2021, time.November, 12, 0, 0, 0, 0, time.UTC), 1500, 1)
			},
		},
		seed.Seed{
			Name: "CreateTrip2",
			Run: func(db *gorm.DB) error {
				return CreateTrip(db, time.Date(2021, time.November, 12, 0, 0, 0, 0, time.UTC), time.Date(2021, time.November, 14, 0, 0, 0, 0, time.UTC), 1500, 1)
			},
		},
		seed.Seed{
			Name: "CreateTrip3",
			Run: func(db *gorm.DB) error {
				return CreateTrip(db, time.Date(2021, time.November, 14, 0, 0, 0, 0, time.UTC), time.Date(2021, time.November, 16, 0, 0, 0, 0, time.UTC), 1500, 1)
			},
		},
	}
}
