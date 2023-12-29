package elastic

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

var ElasticClient *elasticsearch.Client

func ConnectToElasticsearch() {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}

	cert := envFile["http_ca_cert"]

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username:  "elastic",
		Password:  "beukCnNalVU=AWkBd9r*",
		CACert:    []byte(cert),
		Transport: tr,
	}

	ElasticClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect elasticsearch... OK")
}
