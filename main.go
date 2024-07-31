package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
	"golang.org/x/sys/windows/registry"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  jpg-to-pdf <input.jpg>")
		fmt.Println("  jpg-to-pdf --install")
		fmt.Println("  jpg-to-pdf --uninstall")
		return
	}

	switch os.Args[1] {
	case "--install":
		err := installContextMenu()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Context menu entry installed successfully.")
		return
	case "--uninstall":
		err := uninstallContextMenu()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Context menu entry uninstalled successfully.")
		return
	}

	// Existing JPG to PDF conversion code...
	inputFile := os.Args[1]
	outputFile := filepath.Base(inputFile[:len(inputFile)-len(filepath.Ext(inputFile))]) + ".pdf"

	// Open the JPG file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the JPG image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Get image dimensions
	bounds := img.Bounds()
	width := float64(bounds.Max.X)
	height := float64(bounds.Max.Y)

	// Create a new PDF document
	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.AddPage()

	// Calculate scaling factors to fit the image on the page
	pageWidth, pageHeight := pdf.GetPageSize()
	scaleX := pageWidth / width
	scaleY := pageHeight / height
	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	// Add the image to the PDF
	pdf.Image(inputFile, 0, 0, width*scale, height*scale, false, "", 0, "")

	// Save the PDF
	err = pdf.OutputFileAndClose(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully converted %s to %s\n", inputFile, outputFile)
}

func installContextMenu() error {
	key, _, err := registry.CreateKey(registry.CLASSES_ROOT, `SystemFileAssociations\.jpg\shell\Convert to PDF`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.SetStringValue("", "Convert JPG to PDF")
	if err != nil {
		return err
	}

	commandKey, _, err := registry.CreateKey(key, "command", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer commandKey.Close()

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	err = commandKey.SetStringValue("", fmt.Sprintf("\"%s\" \"%%1\"", exePath))
	if err != nil {
		return err
	}

	return nil
}

func uninstallContextMenu() error {
	err := registry.DeleteKey(registry.CLASSES_ROOT, `SystemFileAssociations\.jpg\shell\Convert to PDF\command`)
	if err != nil && err != registry.ErrNotExist {
		return err
	}

	err = registry.DeleteKey(registry.CLASSES_ROOT, `SystemFileAssociations\.jpg\shell\Convert to PDF`)
	if err != nil && err != registry.ErrNotExist {
		return err
	}

	return nil
}
