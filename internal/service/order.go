package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jung-kurt/gofpdf"
	qrcode "github.com/skip2/go-qrcode"
)

const (
	pathToPDF = "/static/pdfs/"
	pathToQR  = "/static/img/qr/"

	title  = "Туристичний документ"
	info_1 = "Цей документ є підставою для перебування в туристичному місці та підтвердженням оплати поїздок."
	info_2 = "Перебування в туристичному місці здійснюється за пред'явленням даного документа та докементами які засвідчують особу."
	info_3 = "Докемент втрачає чиність після закінчення кінцевого терміну поїздки(-ок)."
	info_4 = "Повернення документа та відмінна поїздки регелюється з власником(-ами) індивідуально."
	info_5 = "Перевіряйте дату, кількість та місце поїздок."
	info_6 = "Поїздки розпочинаються та закінчуються відповідно датам в 00:00."
)

type OrderService struct {
	orderRepository repository.Order
}

func NewOrderService(orderRepository repository.Order) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (s *OrderService) GetAllOrdersByUserIdTrip(userID int, role string) ([]models.OrderOutput, error) {
	return s.orderRepository.GetAllOrdersByUserIdTrip(userID, role)
}

func (s *OrderService) GetOneByID(orderID uuid.UUID) (models.OrderOutput, error) {
	return s.orderRepository.GetOneByID(orderID)
}

func (s *OrderService) GetOneOrderByUserIdTrip(userID int, orderID uuid.UUID) (models.Order, error) {
	return s.orderRepository.GetOneOrderByUserIdTrip(userID, orderID)
}

func (s *OrderService) UpdatePaid(id uuid.UUID, paid bool) error {
	return s.orderRepository.UpdatePaid(id, paid)
}

func (s *OrderService) Create(order models.Order) (uuid.UUID, error) {
	return s.orderRepository.Create(order)
}

func (s *OrderService) DownloadOrder(order models.OrderOutput, trips []models.TripsOrderOutput) (string, error) {

	pdf := newReport()

	w := []float64{55, 20, 20, 25}
	head := []string{"Номер замовлення", "Прізвище", "І'мя", "По батькові"}
	pdf = header(pdf, head, w)

	body := [][]string{{order.Id.String(), order.UserSurname, order.UserName, order.UserPatronymic}}
	pdf = table(pdf, body, w)

	pdf.Ln(25)
	tr := format(pdf, 12)
	pdf.CellFormat(180, 10, tr("Поїздки"), "", 0, "C", false, 0, "")
	pdf.Ln(10)

	w = []float64{70, 20, 20, 20, 20, 25}
	head = []string{"Тур", "Початок", "Кінець", "Кількість", "Ціна, грн", "Дата оформлення"}
	pdf = header(pdf, head, w)

	body = [][]string{}
	for _, v := range trips {
		body = append(body, []string{v.Title, fmt.Sprintf("%02d-%02d-%d", v.Start.Day(), v.Start.Month(), v.Start.Year()), fmt.Sprintf("%02d-%02d-%d", v.End.Day(), v.End.Month(), v.End.Year()), fmt.Sprintf("%v", v.Count), fmt.Sprintf("%v", v.Price), fmt.Sprintf("%02d-%02d-%d %02d:%02d", order.CreatedAt.Day(), order.CreatedAt.Month(), order.CreatedAt.Year(), order.CreatedAt.Hour(), order.CreatedAt.Minute())})
	}
	pdf = table(pdf, body, w)

	w = []float64{20}
	head = []string{"Загалом"}
	pdf = header(pdf, head, w)

	body = [][]string{{fmt.Sprintf("%v", order.GeneralPrice)}}
	pdf = table(pdf, body, w)

	pdf.Ln(5)
	tr = format(pdf, 12)
	pdf.CellFormat(180, 10, tr("Відомості для клієнта"), "", 0, "C", false, 0, "")
	pdf.Ln(10)
	tr = format(pdf, 8)
	pdf.Cell(50, 10, tr(info_5))
	pdf.Ln(5)
	pdf.Cell(50, 10, tr(info_6))

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := time.Now()

	namePDF := fmt.Sprintf("%v_%v_%v.pdf", slug.Make(order.UserSurname), slug.Make(order.UserName), fmt.Sprintf("%d_%02d_%02d_%02d_%02d_%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
	nameQR := fmt.Sprintf("%v_%v_%v.png", slug.Make(order.UserSurname), slug.Make(order.UserName), fmt.Sprintf("%d_%02d_%02d_%02d_%02d_%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))

	err = qrcode.WriteFile(order.UserEmail, qrcode.Medium, 256, pwd+pathToQR+nameQR)
	if err != nil {
		log.Println(err)
	}

	pdf.Image(pwd+pathToQR+nameQR, 140, 38, 50, 50, false, "", 0, "")

	err = pdf.OutputFileAndClose(pwd + pathToPDF + namePDF)
	if err != nil {
		log.Println(err)
	}

	return namePDF, nil
}

func newReport() *gofpdf.Fpdf {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pdf := gofpdf.New("P", "mm", "A4", pwd+"/client/src/assets/fonts/fontPDF")

	pdf.AddPage()
	tr := format(pdf, 12)
	pdf.CellFormat(180, 10, tr(title), "", 0, "C", false, 0, "")
	tr = format(pdf, 8)
	pdf.Ln(10)
	pdf.Cell(50, 10, tr(info_1))
	pdf.Ln(5)
	pdf.Cell(50, 10, tr(info_2))
	pdf.Ln(5)
	pdf.Cell(50, 10, tr(info_3))
	pdf.Ln(5)
	pdf.Cell(50, 10, tr(info_4))

	//pdf.CellFormat(40, 20, tr(info), "", 0, "", true, 0, "")
	//SplitText (txt рядок , w float64 ) (рядки [] рядок )
	pdf.Ln(10)

	return pdf
}

func format(pdf *gofpdf.Fpdf, sizeStr float64) (rep func(string) string) {
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.SetFont("Helvetica", "", sizeStr)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	return tr
}

func header(pdf *gofpdf.Fpdf, hdr []string, w []float64) *gofpdf.Fpdf {

	pdf.SetFillColor(255, 255, 255)
	tr := format(pdf, 7)
	for k, str := range hdr {
		pdf.CellFormat(w[k], 5, tr(str), "1", 0, "", true, 0, "")
	}
	pdf.Ln(-1)
	return pdf
}

func table(pdf *gofpdf.Fpdf, tbl [][]string, w []float64) *gofpdf.Fpdf {

	pdf.SetFillColor(255, 255, 255)

	//align := []string{"L", "L", "L", "R", "R", "R"}
	tr := format(pdf, 7)
	for _, line := range tbl {
		for k, str := range line {
			pdf.CellFormat(w[k], 5, tr(str), "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}
	return pdf
}
