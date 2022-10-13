package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SeraS struct {
	m3                          uint16
	icindeYetisenBitkiler       string
	co2                         uint16
	günlükOrtSuIhiyaci          uint16
	günlükOrtEleketrikIhtiyaci  uint16
	enDusukSicaklik             int16
	enYuksekSicaklik            int16
	enYuksektoprakNem           uint16
	enDusuktoprakNem            uint16
	yukselenToprakSonuc         uint16
	dereceBasinaElektrikMasrafi uint16
	m3BasinaSuFaturasi          uint16
	gunlukmahsulsayısı          uint16
	toplamMahsulKg              uint16
}

func init() {
	fmt.Print("\n\n")
	fmt.Println("        ******************************")
	fmt.Println("             Akıllı sera sistemi      ")

	fmt.Println("        ******************************")
	fmt.Print("\n\n")
	fmt.Println("-->Akıllı sera sistemi seranıza ektiğiniz bitkilerin gereksinimlerini kaydeder")
	fmt.Println("-->Seranızdaki günlük co2 , sıcaklık , nem oranı , gün ışığı ve topraktaki nem oranındaki değişimleri hesaplayarak bu değerleri optimum değerlere çevirir.")
	fmt.Println("-->Sistemimizdeki bitki bilgilerini düzenleyebilir , yeni bitkiler tanımlayabilirsiniz.")
	fmt.Println("-->Sistemimze seranızın boyutlarını düzenleytebilirsiniz.")
	fmt.Println("-->Düzenlemelerinizi yaptıktan sonra sistem devreye girer ve size saat başı  raporlama işlemi yapar.")
	fmt.Println("-->Bu raporda akıllı sistemin seranızda yaptığı değişiklikler yer alır")
	fmt.Println("-->Sistemimiz gün sonunda toplam üretilen ve hasat edilen bitkilerin topnajları rapor edecektir")
	fmt.Print("\n\n")

}

func main() {

	var (
		i                      uint16
		k                      uint16
		bitkiCesisiSayisi      uint8
		sera                   SeraS
		yukseklik              uint16
		m2                     uint16
		saatSayisi             uint16
		seraGunSayisi          []uint16
		bitkiCesidi            []string
		elektrikFatura         uint16
		toplamMahsul           uint16
		bitki                  string
		havaCo2Ppm             int16
		havaSicakligi          int16
		sicaklikDegisimi       int16
		bitkigun               uint16
		yukselecekSıcaklık     int16
		dusecekSıcaklık        int16
		yukselenToprakNemi     uint16
		dusenToprakNemi        uint16
		gun                    uint16
		toplamSuFaturası       uint16
		gunlukSuFaturası       uint16
		gunlukElektrikFaturası uint16
		toplamGunlukMahsul     uint16
		toprakNem              uint16
	)
	var elektrikFaturaları [5]uint16 = [5]uint16{0, 0, 0, 0, 0}
	var suFaturaları [5]uint16 = [5]uint16{0, 0, 0, 0, 0}
	var mahsulKiloları [5]uint16 = [5]uint16{0, 0, 0, 0, 0}
	havaSicakligi = 20
	sera.dereceBasinaElektrikMasrafi = 50
	sera.m3BasinaSuFaturasi = 2

	fmt.Println("sera yükseklik belirleyin")
	fmt.Println("lütfen bir tam sayi yaziniz")
	fmt.Scan(&yukseklik)

	fmt.Println("sera m2 belirleyin")
	fmt.Println("lütfen bir tam sayi yaziniz")
	fmt.Scan(&m2)

	sera.m3 = yukseklik * m2
	fmt.Println("Sera yüksekliği = ", sera.m3, " m3")

	fmt.Println("Sistemde özellikleri bulunan bitkliler = marul , havuç , domates , soğan , patates , salatalık")
	fmt.Println("sera da yetiştireceğiniz bitki çeşit sayısını belirleyiniz")
	fmt.Scan(&bitkiCesisiSayisi)
	fmt.Println("Bitki çesidi saysı =", bitkiCesisiSayisi)
	fmt.Println("Serada yetiştireceğiniz bitkileri belirleyiniz")
	for k = 0; k < uint16(bitkiCesisiSayisi); k++ { // kullanıcıdan bitkileri alıyoruz

		fmt.Scan(&bitki)

		bitkiCesidi = append(bitkiCesidi, bitki)

	}

	for t := 0; t < len(bitkiCesidi); t++ { // bitkilerin günlerini alıyoruz
		fmt.Println(t, ". Bitki =", bitkiCesidi[t])
		seraGunu := sera.bitkiOzelligi(bitkiCesidi[t])
		seraGunSayisi = append(seraGunSayisi, seraGunu)
		fmt.Println(t, ". Bitkinin gun sayısı", seraGunu)

	}

	gun = yuksekGunSayisiHesapla(seraGunSayisi) // en uzun güne sahip bitkinin gününü alıyoruz
	for i = 0; i < uint16(gun); i++ {

		time.Sleep(2 * time.Second)
		saatSayisi++
		fmt.Printf("***********************\n")
		fmt.Printf("|    Gun sayisi = %d   |\n", i+1)

		rand.Seed(time.Now().UnixNano())

		sicaklikDegisimi = int16(rand.Intn(5 - (-5)))
		toprakNem = uint16(rand.Intn(15 - 0))
		havaSicakligi += sicaklikDegisimi
		fmt.Printf("|    Hava = %d derece |", havaSicakligi)
		fmt.Print("\n")
		fmt.Println("***********************")
		for v := 0; v < len(bitkiCesidi); v++ {

			bitkigun = sera.bitkiOzelligi(bitkiCesidi[v])
			if i >= bitkigun {
				continue
			}
			fmt.Printf("\n\n\n")
			fmt.Printf("%s serası için değerler", bitkiCesidi[v])
			fmt.Println("")

			if toprakNem < sera.enDusuktoprakNem {
				yukselenToprakNemi = sera.enDusuktoprakNem - toprakNem
				fmt.Printf("-->Toprak nemi yuzde ,%d kadar arttırıldı\n", yukselenToprakNemi)
				gunlukSuFaturası = sera.m3 * sera.m3BasinaSuFaturasi * yukselenToprakNemi
				toplamSuFaturası = gunlukSuFaturası + toplamSuFaturası
				suFaturaları[v] = toplamSuFaturası
				fmt.Printf("-->Günlük su faturası = %d\n", gunlukSuFaturası)
				fmt.Printf("-->Toplam Su faturası %d\n", suFaturaları[(v)])

			} else if toprakNem > sera.enYuksektoprakNem {
				dusenToprakNemi = toprakNem - uint16(sera.enYuksektoprakNem)

				fmt.Printf("-->Toprak nemi yuzde ,%d kadar azaltıldı", dusenToprakNemi)
				fmt.Println()

				gunlukSuFaturası = sera.m3 * sera.m3BasinaSuFaturasi * dusenToprakNemi

				toplamSuFaturası = gunlukSuFaturası + toplamSuFaturası
				suFaturaları[v] = toplamSuFaturası
				fmt.Printf("-->Günlük su faturası = %d\n", gunlukSuFaturası)
				fmt.Printf("-->Toplam Su faturası %d\n", suFaturaları[v])

			} else {

				fmt.Println("-->Nem normal durumda")

			}

			if havaSicakligi < sera.enDusukSicaklik { // ortamdaki co2 seviyesini  sıcaklığa göre belirledik

				yukselecekSıcaklık = sera.enDusukSicaklik - havaSicakligi
				fmt.Printf("-->Hava sıcaklıgı olması gerekenden %d kadar düşük\n", yukselecekSıcaklık)
				havaCo2Ppm = yukselecekSıcaklık * 45
				fmt.Printf("-->Sıcaklık %d°  arttırıldı\n", yukselecekSıcaklık)

				fmt.Printf("-->CO2 %d  kadar arttırıldı\n", havaCo2Ppm)

				gunlukElektrikFaturası = sera.m3 * sera.dereceBasinaElektrikMasrafi * uint16(yukselecekSıcaklık)
				fmt.Printf("-->Gunluk elektrik faturası = %d\n", gunlukElektrikFaturası)
				elektrikFatura = gunlukElektrikFaturası + elektrikFatura
				elektrikFaturaları[v] = elektrikFatura

				fmt.Printf("\n\n\n")

			} else if havaSicakligi > sera.enYuksekSicaklik {
				dusecekSıcaklık = havaSicakligi - sera.enYuksekSicaklik
				fmt.Printf("-->Hava sıcaklıgı olması gerekenden %d kadar yüksek\n", dusecekSıcaklık)

				havaCo2Ppm = dusecekSıcaklık * 45
				fmt.Printf("-->Sıcaklık %d°  azaltıldı\n", dusecekSıcaklık)

				fmt.Printf("-->CO2 %d  kadar azaltıldı\n", havaCo2Ppm)

				gunlukElektrikFaturası = sera.m3 * sera.dereceBasinaElektrikMasrafi * uint16(dusecekSıcaklık)
				elektrikFaturaları[v] = elektrikFatura
				fmt.Printf("-->Gunluk elektrik faturası = %d\n", gunlukElektrikFaturası)

			} else {

				fmt.Println("-->Sıcaklık ve CO2 normal durumda")

			}
			fmt.Printf("-->Toplam Elektrik Faturası = %d tl\n", elektrikFaturaları[v])
			toplamGunlukMahsul = sera.gunlukmahsulsayısı * m2
			toplamMahsul = toplamGunlukMahsul + toplamMahsul

			mahsulKiloları[v] = toplamMahsul

			if (i+1)%4 == 0 {

				fmt.Printf("-->şuana kadar ki toplam %d mahsulu", mahsulKiloları[v])
				fmt.Printf("\n\n\n")
			}

		}
		havaSicakligi = 20
	}
	fmt.Printf("\n\n\n")
	fmt.Printf("\n\n\n")
	for v := 0; v < len(bitkiCesidi); v++ {
		fmt.Printf("-->%s icin Odenmesi gereken toplam elektrik farurası %d tl\n", bitkiCesidi[v], elektrikFaturaları[(v)])
		fmt.Printf("-->%s icin Odenmesi gereken toplam su farurası %d tl\n", bitkiCesidi[v], suFaturaları[(v)])
		fmt.Printf("-->%s icin toplam mahsul %d kg\n", bitkiCesidi[v], mahsulKiloları[(v)])
		fmt.Printf("\n\n\n")
	}
}
func (a *SeraS) bitkiOzelligi(bitki string) uint16 {
	var gun uint16
	var bitkiGunu map[string]uint16
	bitkiGunu = make(map[string]uint16)

	bitkiGunu["marul"] = 90
	bitkiGunu["havuç"] = 120
	bitkiGunu["domates"] = 100
	bitkiGunu["patates"] = 90
	bitkiGunu["soğan"] = 85
	bitkiGunu["salatalık"] = 110

	for k, v := range bitkiGunu {
		switch {

		case k == "marul" && bitki == "marul":
			a.co2 = 1200
			a.günlükOrtSuIhiyaci = 3
			a.günlükOrtEleketrikIhtiyaci = 5
			gun = v
			a.enDusukSicaklik = 22
			a.enYuksekSicaklik = 27
			a.enDusuktoprakNem = 8
			a.enYuksektoprakNem = 11
			a.gunlukmahsulsayısı = 10
			a.toplamMahsulKg = 0

		case k == "havuç" && bitki == "havuç":
			a.co2 = 1100
			a.günlükOrtSuIhiyaci = 2
			a.günlükOrtEleketrikIhtiyaci = 8
			gun = v
			a.enDusukSicaklik = 23
			a.enYuksekSicaklik = 26
			a.enDusuktoprakNem = 9
			a.enYuksektoprakNem = 14
			a.gunlukmahsulsayısı = 8
			a.toplamMahsulKg = 0

		case k == "domates" && bitki == "domates":
			a.co2 = 1500
			a.günlükOrtSuIhiyaci = 1
			a.günlükOrtEleketrikIhtiyaci = 3
			gun = v
			a.enDusukSicaklik = 24
			a.enYuksekSicaklik = 29
			a.enDusuktoprakNem = 10
			a.enYuksektoprakNem = 14
			a.gunlukmahsulsayısı = 6
			a.toplamMahsulKg = 0

		case k == "patates" && bitki == "patates":
			a.co2 = 900
			a.günlükOrtSuIhiyaci = 1
			a.günlükOrtEleketrikIhtiyaci = 4
			gun = v
			a.enDusukSicaklik = 21
			a.enYuksekSicaklik = 28
			a.enDusuktoprakNem = 11
			a.enYuksektoprakNem = 13
			a.gunlukmahsulsayısı = 3
			a.toplamMahsulKg = 0

		case k == "soğan" && bitki == "soğan":
			a.co2 = 1000
			a.günlükOrtSuIhiyaci = 3
			a.günlükOrtEleketrikIhtiyaci = 7
			gun = v
			a.enDusukSicaklik = 21
			a.enYuksekSicaklik = 27
			a.enDusuktoprakNem = 9
			a.enYuksektoprakNem = 11
			a.gunlukmahsulsayısı = 12
			a.toplamMahsulKg = 0

		case k == "salatalık" && bitki == "salatalık":
			a.co2 = 1100
			a.günlükOrtSuIhiyaci = 2
			a.günlükOrtEleketrikIhtiyaci = 5
			gun = v
			a.enDusukSicaklik = 21
			a.enYuksekSicaklik = 23
			a.enDusuktoprakNem = 5
			a.enYuksektoprakNem = 10
			a.gunlukmahsulsayısı = 14
			a.toplamMahsulKg = 0

		}
	}
	return gun
}
func yuksekGunSayisiHesapla(slice []uint16) uint16 {
	var (
		a uint16
		i uint16
	)
	for i = 0; i < uint16(len(slice)); i++ {

		if slice[i] > a {
			a = slice[i]
		}

	}

	return a

}
