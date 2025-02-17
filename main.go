package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// Function to handle the /get-pdf API route
func getPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the LaTeX code from the request body
	latexCode, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading LaTeX code", http.StatusInternalServerError)
		return
	}

	// Write LaTeX code to a temporary .tex file
	texFile := fmt.Sprintf("document_%d.tex", time.Now().Unix())
	err = ioutil.WriteFile(texFile, latexCode, 0644)
	if err != nil {
		http.Error(w, "Error writing LaTeX file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(texFile)

	// Execute pdflatex command to generate the PDF
	cmd := exec.Command("pdflatex", texFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Error compiling LaTeX", http.StatusInternalServerError)
		return
	}

	// Serve the generated PDF file
	pdfFile := fmt.Sprintf("document_%d.pdf", time.Now().Unix())
	defer os.Remove(pdfFile)

	// Move the generated PDF to the correct file
	err = os.Rename("document.pdf", pdfFile)
	if err != nil {
		http.Error(w, "Error moving PDF file", http.StatusInternalServerError)
		return
	}

	// Set the headers to prompt for file download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", pdfFile))
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, pdfFile)
}

func main() {
	// Set up the API route
	http.HandleFunc("/get-pdf", getPDFHandler)

	// Start the HTTP server
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
