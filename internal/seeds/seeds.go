package seeds

import (
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/pkg/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateRole1",
			Run: func(db *gorm.DB) error {
				return CreateRole(db, "owner")
			},
		},
		seed.Seed{
			Name: "CreateRole2",
			Run: func(db *gorm.DB) error {
				return CreateRole(db, "client")
			},
		},
		seed.Seed{
			Name: "CreateUser1",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "vova", "saenko", "volodumurovuch", "vova@gmail.com", "123456", "0953417774", 1)
			},
		},
		seed.Seed{
			Name: "CreateUser2",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "roma", "boresenko", "batckovuch", "roma@gmail.com", "123456", "0953417774", 2)
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
					"opus opus opus opus opus opus opus opus opus tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_1",
					false,
					true,
					1,
					2,
					1,
					50.385197334367426,
					30.474394711883974,
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
					"opus opus opus opus opus opus opus opus opus tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					false,
					true,
					4,
					1,
					1,
					50.37928638275505,
					30.516623409457985,
				)
			},
		},

		seed.Seed{
			Name: "CreateTour3",
			Run: func(db *gorm.DB) error {
				return CreateTour(db,
					"Тур на вихідні в стилі активного відпочинку 3",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_3",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					"opus opus opus opus opus opus opus opus opus tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					false,
					true,
					4,
					1,
					1,
					50.37928638275505,
					30.516623409457985,
				)
			},
		},

		seed.Seed{
			Name: "CreateTour4",
			Run: func(db *gorm.DB) error {
				return CreateTour(db,
					"Тур на вихідні в стилі активного відпочинку 4",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_4",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					"opus opus opus opus opus opus opus opus opus tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					false,
					true,
					4,
					1,
					1,
					50.37928638275505,
					30.516623409457985,
				)
			},
		},

		seed.Seed{
			Name: "CreateTour5",
			Run: func(db *gorm.DB) error {
				return CreateTour(db,
					"Тур на вихідні в стилі активного відпочинку 5",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_5",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					"opus opus opus opus opus opus opus opus opus tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					false,
					true,
					4,
					1,
					1,
					50.37928638275505,
					30.516623409457985,
				)
			},
		},

		seed.Seed{
			Name: "CreateTour6",
			Run: func(db *gorm.DB) error {
				return CreateTour(db,
					"Тур на вихідні в стилі активного відпочинку 6",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_6",
					"tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					"opus opus opus opus opus opus opus opus opus tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku tur_na_vykhidni_v_styli_aktyvnoho_vidpochynku_2",
					false,
					true,
					4,
					1,
					1,
					50.37928638275505,
					30.516623409457985,
				)
			},
		},

		seed.Seed{
			Name: "CreateImage1",
			Run: func(db *gorm.DB) error {
				return CreateImage(db, "vopros-otvet-pryzhki-s-bandzhi-naskolko-eto-opasno_15370996521560642868.jpg", 1)
			},
		},
		seed.Seed{
			Name: "CreateImage1a",
			Run: func(db *gorm.DB) error {
				return CreateImage(db, "vopros-otvet-pryzhki-s-bandzhi-naskolko-eto-opasno_15370996521560642868.jpg", 1)
			},
		},

		seed.Seed{
			Name: "CreateImage2",
			Run: func(db *gorm.DB) error {
				return CreateImage(db, "1614867483_img-01.jpg", 2)
			},
		},
		seed.Seed{
			Name: "CreateTrip1",
			Run: func(db *gorm.DB) error {
				return CreateTrip(db, time.Date(2021, 5, 25, 0, 0, 0, 0, time.UTC), time.Date(2021, time.November, 27, 0, 0, 0, 0, time.UTC), 1500, 5, 1)
			},
		},
		seed.Seed{
			Name: "CreateTrip2",
			Run: func(db *gorm.DB) error {
				return CreateTrip(db, time.Date(2021, 5, 27, 0, 0, 0, 0, time.UTC), time.Date(2021, time.November, 29, 0, 0, 0, 0, time.UTC), 1500, 5, 1)
			},
		},
		seed.Seed{
			Name: "CreateTrip3",
			Run: func(db *gorm.DB) error {
				return CreateTrip(db, time.Date(2021, 5, 29, 0, 0, 0, 0, time.UTC), time.Date(2021, time.November, 30, 0, 0, 0, 0, time.UTC), 1500, 5, 1)
			},
		},

		// Events
		seed.Seed{
			Name: "CreateEvent1",
			Run: func(db *gorm.DB) error {
				return CreateEvent(db, "Всі сніданки, обіди та вечері")
			},
		},
		seed.Seed{
			Name: "CreateEvent2",
			Run: func(db *gorm.DB) error {
				return CreateEvent(db, "Екскурсійний та туристичний путівник на всю подорож")
			},
		},
		seed.Seed{
			Name: "CreateEvent3",
			Run: func(db *gorm.DB) error {
				return CreateEvent(db, "Страхування подорожей та інші надзвичайні ситуації")
			},
		},

		// Events
		seed.Seed{
			Name: "CreateToursEvent1",
			Run: func(db *gorm.DB) error {
				return CreateToursEvent(db, 1, 1)
			},
		},
		seed.Seed{
			Name: "CreateToursEvent3",
			Run: func(db *gorm.DB) error {
				return CreateToursEvent(db, 1, 3)
			},
		},

		// Comment
		seed.Seed{
			Name: "CreateComment1",
			Run: func(db *gorm.DB) error {
				return CreateComment(db, "Comment1 Всі сніданки, обіди та вечері", 4, 1, 2, time.Date(2021, time.November, 14, 0, 0, 0, 0, time.UTC))
			},
		},
		seed.Seed{
			Name: "CreateComment2",
			Run: func(db *gorm.DB) error {
				return CreateComment(db, "Comment2 Екскурсійний та туристичний путівник на всю подорож", 5, 1, 2, time.Date(2021, time.November, 15, 0, 0, 0, 0, time.UTC))
			},
		},
		seed.Seed{
			Name: "CreateComment3",
			Run: func(db *gorm.DB) error {
				return CreateComment(db, "Comment3 подорожей та інші надзвичайні ситуації", 3, 1, 2, time.Date(2021, time.November, 16, 0, 0, 0, 0, time.UTC))
			},
		},

		// Order
		seed.Seed{
			Name: "CreateOrder1",
			Run: func(db *gorm.DB) error {
				return CreateOrder(db, "dc9076e9-2fda-4019-bd2c-900a8284b9c4", true, "0953417774", "1111 lalalala", time.Date(2021, time.November, 14, 0, 0, 0, 0, time.UTC), 2)
			},
		},
		seed.Seed{
			Name: "CreateOrder2",
			Run: func(db *gorm.DB) error {
				return CreateOrder(db, "dc9076e9-2fda-4019-bd2c-900a8284b9c2", false, "0502801442", "2222 lalalala", time.Date(2021, time.November, 15, 0, 0, 0, 0, time.UTC), 2)
			},
		},

		// OrderTrips
		seed.Seed{
			Name: "CreateOrderTrip1",
			Run: func(db *gorm.DB) error {
				return CreateOrderTrip(db, 1, 1, "dc9076e9-2fda-4019-bd2c-900a8284b9c4")
			},
		},
		seed.Seed{
			Name: "CreateOrderTrip2",
			Run: func(db *gorm.DB) error {
				return CreateOrderTrip(db, 2, 2, "dc9076e9-2fda-4019-bd2c-900a8284b9c4")
			},
		},
		seed.Seed{
			Name: "CreateOrderTrip3",
			Run: func(db *gorm.DB) error {
				return CreateOrderTrip(db, 1, 1, "dc9076e9-2fda-4019-bd2c-900a8284b9c2")
			},
		},
		seed.Seed{
			Name: "CreateOrderTrip4",
			Run: func(db *gorm.DB) error {
				return CreateOrderTrip(db, 2, 2, "dc9076e9-2fda-4019-bd2c-900a8284b9c2")
			},
		},
	}
}
