package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	fmt.Println("\n\nPretende criar um template vazio? (S|N)\n\n")
	log.SetPrefix("FATCA_GENERATOR_ERROR -> ")

	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()

	data := string(char)
	data = strings.ToLower(data)
	currentDirectory, _ := os.Getwd()

	switch data {
	case "s":
		f, err := excelize.OpenFile("fatca.xlsx")
		if err != nil {
			log.SetFlags(0)
			log.Fatal(err)

		}

		// GETTING CELLS AND VALUES
		SendingCompanyIN := f.GetCellValue("Sheet1", "A2")
		TransmittingCountry := f.GetCellValue("Sheet1", "B2")
		ReceivingCountry := f.GetCellValue("Sheet1", "C2")
		MessageType := f.GetCellValue("Sheet1", "D2")
		MessageRefId := f.GetCellValue("Sheet1", "E2")
		ReportingPeriod := f.GetCellValue("Sheet1", "F2")
		ResCountryCode := f.GetCellValue("Sheet1", "G2")
		TIN := f.GetCellValue("Sheet1", "H2")
		Name := f.GetCellValue("Sheet1", "I2")
		CountryCode := f.GetCellValue("Sheet1", "J2")
		AddressFree := f.GetCellValue("Sheet1", "K2")
		FilerCategory := f.GetCellValue("Sheet1", "L2")
		DocTypeIndic := f.GetCellValue("Sheet1", "M2")
		DocRefId := f.GetCellValue("Sheet1", "N2")
		NoAccountToReport := f.GetCellValue("Sheet1", "O2")

		Timestamp := time.Now().Format(time.ANSIC)

		dataModel := DataModel{SendingCompanyIN,
			TransmittingCountry,
			ReceivingCountry,
			MessageType,
			MessageRefId,
			ReportingPeriod,
			Timestamp,
			ResCountryCode,
			TIN,
			Name,
			CountryCode,
			AddressFree,
			FilerCategory,
			DocTypeIndic,
			DocRefId,
			NoAccountToReport}

		dataModel.Timestamp = time.Now().Format("02-01-2006")

		templateData, err := template.New("Minifin => ").Parse(
			`<?xml version="1.0" encoding= "UTF-8"?>
					 <ftc:FATCA_OECD version="2.0"
					 xsi:schemaLocation="urn:oecd:ties:fatca:v2 FatcaXML_v2.0.xsd"
					 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
					 xmlns:ftc="urn:oecd:ties:fatca:v2"
					 xmlns:sfa="urn:oecd:ties:stffatcatypes:v2">
					 <ftc:MessageSpec><sfa:SendingCompanyIN>{{.SendingCompanyIN}}</sfa:SendingCompanyIN>
					 <sfa:TransmittingCountry>{{.TransmittingCountry}}</sfa:TransmittingCountry>
					 <sfa:ReceivingCountry>{{.ReceivingCountry}}</sfa:ReceivingCountry>
					 <sfa:MessageType>FATCA</sfa:MessageType>
					 <sfa:MessageRefId>{{.MessageRefId}}</sfa:MessageRefId>
					 <sfa:ReportingPeriod>{{.ReportingPeriod}}</sfa:ReportingPeriod>
					 <sfa:Timestamp>{{.Timestamp}}</sfa:Timestamp>
					 </ftc:MessageSpec>
					 <ftc:FATCA>
					 <ftc:ReportingFI>
					 <sfa:ResCountryCode>{{.ResCountryCode}}</sfa:ResCountryCode>
					 <sfa:TIN>{{.TIN}}</sfa:TIN>
					 <sfa:Name>{{.Name}}</sfa:Name>
					 <sfa:Address>
					 <sfa:CountryCode>{{.CountryCode}}</sfa:CountryCode>
					 <sfa:AddressFree>{{.AddressFree}}</sfa:AddressFree>
					 </sfa:Address>
					 <ftc:FilerCategory>{{.FilerCategory}}</ftc:FilerCategory>
					 <ftc:DocSpec>
					 <ftc:DocTypeIndic>{{.DocTypeIndic}}</ftc:DocTypeIndic>
					 <ftc:DocRefId>{{.DocRefId}}
					 </ftc:DocRefId>
					 </ftc:DocSpec>
					 </ftc:ReportingFI>
					 <ftc:ReportingGroup>
					 <ftc:NilReport>
					 <ftc:DocSpec>
					 <ftc:DocTypeIndic>{{.DocTypeIndic}}</ftc:DocTypeIndic>
					 <ftc:DocRefId>{{.DocRefId}}</ftc:DocRefId>
					 </ftc:DocSpec>
					 <ftc:NoAccountToReport>{{.NoAccountToReport}}</ftc:NoAccountToReport>
					 </ftc:NilReport>
					 </ftc:ReportingGroup>
					 </ftc:FATCA>
					 </ftc:FATCA_OECD>
				`)

		file, e := os.OpenFile("fatca.xml", os.O_RDWR|os.O_CREATE, 0666)

		if e != nil {

			fmt.Println(e)
			os.Exit(3)

		}

		templateData.Execute(file, dataModel)

		fmt.Println("\n\nSeu ficheiro foi criado em ", currentDirectory, "/fatca.xml")
		os.Exit(3)

	case "n":
		fmt.Println("Documento não disponivel!")
		os.Exit(3)

	}

	log.Fatal("Documento não disponivel!")
	os.Exit(3)

}

type DataModel struct {
	SendingCompanyIN,
	TransmittingCountry,
	ReceivingCountry,
	MessageType,
	MessageRefId,
	ReportingPeriod,
	Timestamp,
	ResCountryCode,
	TIN,
	Name,
	CountryCode,
	AddressFree,
	FilerCategory,
	DocTypeIndic,
	DocRefId,
	NoAccountToReport string
}
