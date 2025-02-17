package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// LaTeX content
	latexCode := `
    \documentclass[a4paper,10pt]{article}
\usepackage[utf8]{inputenc}
\usepackage{titlesec}
\usepackage{hyperref}
\usepackage{xcolor}
\usepackage{multicol}
\usepackage{parskip}

% Manual margin adjustments
\setlength{\topmargin}{-0.5in}
\setlength{\evensidemargin}{0in}
\setlength{\oddsidemargin}{0in}
\setlength{\textwidth}{6.5in}
\setlength{\textheight}{9in}
\setlength{\headheight}{0in}
\setlength{\headsep}{0.25in}
\setlength{\footskip}{0.5in}
\setlength{\parskip}{1em}

% Section formatting
\titleformat{\section}{\Large\bfseries}{}{0em}{}[\titlerule]
\titleformat{\subsection}{\large\bfseries}{}{0em}{}

% Hyperlink settings
\hypersetup{
    colorlinks=true,
    linkcolor=black,
    filecolor=black,      
    urlcolor=black,
}

% Custom commands
\newcommand{\resumeItem}[2]{
  \item\textbf{#1}{: #2}
}
\newcommand{\resumeSubheading}[4]{
  \item
    \textbf{#1} \hfill #2 \\
    \textit{#3} \hfill \textit{#4}
}
\newcommand{\resumeSubSubheading}[2]{
    \item
    \textbf{#1} \hfill #2
}

\begin{document}

\begin{tabbing}
\hspace{4in} \= \kill
{\Huge \textbf{Ritesh Girase}} \\
\href{mailto:riteshgirasecode@gmail.com}{riteshgirasecode@gmail.com} \> 
\hspace{1.8in}
\href{https://www.linkedin.com/}{LinkedIn} \\
Nashik, Maharashtra, India \> \hspace{1.85in}\href{https://github.com}{GitHub} \\ 
9552566981 \\
\end{tabbing}
\vspace{-0.4cm}

\end{document}


	`

	// Write LaTeX code to a .tex file
	texFile := "document.tex"
	err := os.WriteFile(texFile, []byte(latexCode), 0644)
	if err != nil {
		fmt.Println("Error writing LaTeX file:", err)
		return
	}

	// Execute pdflatex command to generate the PDF
	cmd := exec.Command("pdflatex", texFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error compiling LaTeX:", err)
		return
	}

	fmt.Println("PDF generated successfully!")
}
