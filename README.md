# SnapPDF

SnapPDF is a lightweight, user-friendly Windows application that allows you to quickly convert JPG images to PDF files directly from your context menu. With SnapPDF, transforming your images into professional-looking PDFs is just a right-click away!

## Features

- Convert JPG images to PDF with a single click
- Integrate seamlessly with Windows Explorer via context menu
- Maintain image quality during conversion
- Support for multiple file conversion (coming soon)
- Easy installation and uninstallation process

## Installation

1. Download the latest release of SnapPDF from the [Releases](https://github.com/tungpsit/SnapPDF/releases) page.
2. Run the downloaded executable file.
3. To add SnapPDF to your context menu, run:
```
SnapPDF.exe --install
```

Note: You may need to run this command as an administrator.

## Usage

After installation, you can use SnapPDF in two ways:

1. **Via Context Menu:**
- Right-click on any JPG file in Windows Explorer.
- Select "Convert JPG to PDF" from the context menu.
- The PDF will be created in the same directory as the original image.

2. **Via Command Line:**

```
SnapPDF.exe <input.jpg>
```

This will create a PDF file in the same directory as the input JPG file.

## Uninstallation

To remove SnapPDF from your context menu:

1. Open a command prompt (run as administrator).
2. Navigate to the directory containing SnapPDF.exe.
3. Run: 
  ```
  SnapPDF.exe --uninstall
  ```
  
## Building from Source

To build SnapPDF from source:

1. Ensure you have Go installed on your system.
2. Clone this repository:
  ```
  git clone https://github.com/yourusername/SnapPDF.git
  ```
1. Navigate to the project directory: 
  ```
  cd SnapPDF
  ```
2. Install dependencies:
  ```
  go get github.com/jung-kurt/gofpdf
  go get golang.org/x/sys/windows/registry
  go build
  ```


## Contributing

Contributions to SnapPDF are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [gofpdf](https://github.com/jung-kurt/gofpdf) for PDF creation
- [golang.org/x/sys](https://golang.org/x/sys) for Windows registry manipulation

## Support

If you encounter any issues or have any questions, please file an issue on the GitHub issue tracker.

---

We hope SnapPDF makes your document conversion tasks a breeze! Happy converting!
