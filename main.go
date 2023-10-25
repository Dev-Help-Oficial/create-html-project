package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var targetDir string
	var projectName string

	// Verifica se o usuário forneceu um nome de pasta como argumento
	if len(os.Args) > 1 {
		projectName = os.Args[1]
	} else {
		projectName = "meu-projeto" // Nome padrão caso o usuário não especifique
	}

	// Verifique se o usuário deseja criar a pasta na pasta de trabalho atual
	if projectName == "./" || projectName == "" {
		targetDir, _ = os.Getwd()
	} else {
		targetDir = projectName
	}

	// Cria a pasta do projeto
	err := os.Mkdir(targetDir, 0755)
	if err != nil {
		fmt.Println("Erro ao criar a pasta do projeto:", err)
		return
	}

	// Cria as pastas 'css' e 'js'
	cssDir := filepath.Join(targetDir, "css")
	jsDir := filepath.Join(targetDir, "js")

	err = os.Mkdir(cssDir, 0755)
	if err != nil {
		fmt.Println("Erro ao criar a pasta 'css':", err)
		return
	}

	err = os.Mkdir(jsDir, 0755)
	if err != nil {
		fmt.Println("Erro ao criar a pasta 'js':", err)
		return
	}

	// Cria o arquivo index.html
	indexHTML := filepath.Join(targetDir, "index.html")
	file, err := os.Create(indexHTML)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo index.html:", err)
		return
	}
	defer file.Close()

	// Escreve o conteúdo HTML no arquivo index.html
	htmlContent := `

	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css">
		<link rel="stylesheet" href="./css/style.css">
		<script src="./js/script.js"></script>
		<title>create-html-bootstrap</title>
	</head>
	<body>
		<h2 class="text-center h3 p-5">Tema boilerplate auto-gerado por create-html-bootstrap</h2>
	
		<main class="container flex min-h-screen flex-col items-center justify-between p-5 m-5">
			<div class="mt-5 row">
				<div class="col-md-6 col-lg-6">
					<a class="text-reset text-decoration-none" href="https://github.com/Dev-Help-Oficial/blob/main/CONTRIBUTING.md" target="_blank">
					<div class="card mb-3 info-card">
						<div class="card-body">
							<h2 class="card-title h3 fw-semibold user-select-none">Contribua <i class="arrow fas fa-arrow-right"></i></h2>
							<p class="text-muted py-2 user-select-none card-text">Contribua com o projeto com Issues e Pull Requests em nosso Github.</p>
						</div>
					</div>
				</a>
				</div>
				<div class="col-md-6 col-lg-6">
					<a class="text-reset text-decoration-none" href="https://github.com/Dev-Help-Oficial/" target="_blank">
					<div class="card mb-3 info-card">
						<div class="card-body">
							<h2 class="card-title h3 fw-semibold user-select-none">Visite nosso GitHub <i class="arrow fas fa-arrow-right arrow"></i></h2>
							<p class="text-muted py-2 user-select-none card-text">Clique aqui para acessar a nossa organização no Github e ver outros projetos.</p>
						</div>
					</div>
					</a>
				</div>
			</div>
		</main>
	
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js"></script>
	</body>
	</html>
	

	
`
	_, err = file.WriteString(htmlContent)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo index.html:", err)
	}

	// Crie o arquivo style.css e escreva seu conteúdo
	styleCSS := filepath.Join(cssDir, "style.css")
	cssFile, err := os.Create(styleCSS)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo style.css:", err)
		return
	}
	defer cssFile.Close()

	cssContent := `
	body {
		background: linear-gradient(to bottom, #fffefe, #b6b6ff);
		background-attachment: fixed;
	}

   .info-card {
		background: transparent;
		cursor: pointer;

	}

	.info-card:hover {
		border: 1px solid;
	}

	.info-card .arrow {
		display: inline-block;
		margin-left: 5px;
		transition: transform 0.3s;
	}

	.info-card:hover .arrow {
		transform: translateX(5px);
	}
`
	_, err = cssFile.WriteString(cssContent)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo style.css:", err)
	}

	scriptJS := filepath.Join(jsDir, "script.js")
	jsFile, err := os.Create(scriptJS)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo script.js:", err)
		return
	}
	defer jsFile.Close()

	fmt.Printf("Projeto HTML gerado com sucesso em: %s\n", targetDir)
}
