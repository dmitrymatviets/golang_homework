package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olivere/elastic"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

/*
Взять документ из ElasticSearch и сохранить его в xml фаил.
При запуске программы на вход параметром подается ссылка на картинку. Картинку нужно сохранить в БД.
Сделать метод get для скачивания картинки.
*/

const (
	elasticHost      = "http://nginx01.test.lan:9300"
	elasticIndexName = "item_index_v7"
	itemId           = "100023350026"
	filePath         = "./ch5-7/practice/homework/item.xml"
	defaultUrl       = "https://img.tsn.ua/cached/1533898791/tsn-3ad8a7940cc99f147f48233aa7502420/thumbs/585xX/ac/00/2f07e665934361372c1544e1591700ac.jpeg"
	dbPath           = "./ch5-7/practice/homework/db.db"
	port             = ":1234"
)

var dbInstance *sql.DB

func main() {
	dbInstance = connect(dbPath)
	defer dbInstance.Close()

	elasticToXml()

	url := defaultUrl
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	saveImageToDb(downloadImage(url))

	startWebServer(port)

}

func saveImageToDb(image *[]byte, contentType string) {
	_, err := dbInstance.Exec("INSERT INTO `Files` (contentType, file) values(?, ?)", contentType, &image)
	if err != nil {
		panic(fmt.Errorf("ошибка сохранения в БД: %v", err))
	}
	fmt.Println("Картинка сохранена в БД")
}

func startWebServer(port string) {
	fmt.Println("Сервер запущен", port)
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		id := (r.URL.Query()).Get("id")

		writeErr := func(status int) {
			w.WriteHeader(status)
			fmt.Fprintf(w, "bad request")
		}

		if id == "" {
			writeErr(http.StatusBadRequest)
			return
		} else {
			contentType, file, err := getImageFromDb(id)
			if err != nil {
				fmt.Println(err)
				writeErr(http.StatusBadRequest)
				return
			}

			w.Header().Add("Content-Type", contentType)
			w.Write(file)
		}

	})
	http.ListenAndServe(port, nil)
}

func getImageFromDb(id string) (contentType string, file []byte, err error) {
	rows, err := dbInstance.Query("SELECT contentType, file FROM Files WHERE id=?", id)

	if err != nil {
		return
	}

	rows.Next()

	if err = rows.Scan(&contentType, &file); err != nil {
		return
	}

	return
}

func downloadImage(url string) (*[]byte, string) {
	var buffer []byte

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		panic(fmt.Errorf("Ошибка запроса картинки: %v", err))
	}

	buffer, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Errorf("Ошибка получения картинки: %v", err))
	}

	fmt.Println("Картинка загружена")

	return &buffer, response.Header.Get("Content-Type")
}

func elasticToXml() string {
	searchResult := getFromElastic(itemId)

	item := &ElasticItem{}
	json.Unmarshal(*searchResult.Hits.Hits[0].Source, item)

	xmlItem := writeXmlToFile(item, filePath)

	result := string(xmlItem)
	fmt.Println(result)

	return result
}

func writeXmlToFile(item interface{}, fileName string) string {
	xmlItem, err := xml.MarshalIndent(item, "  ", "    ")
	if err != nil {
		fmt.Printf("ошибка маршаллинга xml: %v\n", err)
	}

	ioutil.WriteFile(fileName, xmlItem, os.FileMode(777))

	fmt.Println("Файл записан")
	return string(xmlItem)
}

func getFromElastic(id string) *elastic.SearchResult {
	client, err := elastic.NewClient(
		elastic.SetURL(elasticHost))
	if err != nil {
		panic(fmt.Errorf("ошибка создания клиента %v", err))
	}
	termQuery := elastic.NewTermQuery("_id", id)
	searchResult, err := client.Search().
		Index(elasticIndexName).
		Query(termQuery).
		Do(context.Background())
	if err != nil {
		panic(fmt.Errorf("ошибка поиска в эластике %v", err))

	}
	if searchResult == nil || searchResult.Hits == nil || searchResult.Hits.Hits == nil || len(searchResult.Hits.Hits) <= 0 {
		panic(fmt.Errorf("не найдено в elastic"))
	}

	return searchResult
}

func connect(dbPath string) *sql.DB {

	dbPath, _ = filepath.Abs(dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Errorf("ошибка коннекта к БД: %v", err))
	}
	return db

}

type ElasticItem struct {
	Id    string `json:"item_id" xml:"id"`
	Brand string `json:"brand" xml:"brand"`
	Name  string `json:"registry_name" xml:"name"`
}
