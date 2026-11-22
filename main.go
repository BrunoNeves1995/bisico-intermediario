package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	URL      = "postgres://%s:%s@%s:%d/%s?sslmode=disable"
	PGUSER   = "postgres"
	PGPASS   = "postgres"
	PGPORT   = 5432
	PGSERVER = "172.26.144.1"
	PGDB     = "estudo"
)

func main() {
	//Banco
	dsn := fmt.Sprintf(URL, PGUSER, PGPASS, PGSERVER, PGPORT, PGDB)
	log.Println(dsn)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("unable to parse DSN: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	//Servidor
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var clientes []Cliente
		rows, err := db.Query("SELECT * FROM public.teste")
		if err != nil {
			log.Println("Erro ao consultar banco:", err)
			http.Error(w, "Erro ao buscar clientes", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var c Cliente
			err := rows.Scan(&c.Id, &c.NomeRazao)
			if err != nil {
				log.Println("Erro ao escanear linha:", err)
				http.Error(w, "Erro ao processar dados", http.StatusInternalServerError)
				return
			}
			clientes = append(clientes, c)
		}

		if err := rows.Err(); err != nil {
			log.Println("Erro na iteração das linhas:", err)
			http.Error(w, "Erro ao processar dados", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(clientes); err != nil {
			log.Println("Erro ao serializar JSON:", err)
			http.Error(w, "Erro ao gerar resposta", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("servidor rodando: http://localhost:8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		fmt.Println("Erro ao iniciar o servidor", err)
	}
}
