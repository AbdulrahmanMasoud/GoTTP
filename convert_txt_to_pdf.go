package main

import ( 
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
)

func main() {

	// Declaring Variabels
	var file, pdfFile string // This For File txt name , This for What You Ned PDF Flil Bee
	var fontSize float64 // This For Font Size

	// Get File Name From User
	fmt.Print("Plesa Enter Your File Name: ")
	fmt.Scanln(&file)

	// Get Font Size From User
	fmt.Print("Plesa Enter Your Fonts size: ")
	fmt.Scanln(&fontSize)

	// Get New Name For PDF File
	fmt.Print("Plesa Enter Your New File Name: ")
	fmt.Scanln(&pdfFile)
	
	// Read .txt File
	content, err := ioutil.ReadFile(file + ".txt")

	// If has No File Like this Name Just Get Err msg
	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	// Make PDF Config File
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage() // Add Data Into PDF
	pdf.SetFont("Arial","B", fontSize) // This Configration With Font

	pdf.MultiCell(190, 5, string(content), "0", "0", false) // Configration With Page Like widt and Height

	_ = pdf.OutputFileAndClose(pdfFile + ".pdf") // Make This PDF File

	fmt.Println("PDF Created") // Just Print Success Message
}