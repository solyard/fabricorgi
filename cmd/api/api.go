package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fabricorgi/cmd/orgchecker"
	"github.com/fabricorgi/cmd/signer"
	"github.com/fabricorgi/config"
	"github.com/gorilla/mux"
)

type data struct {
	Data string
}

// APIEndpoint ...
const APIEndpoint = "/api/v1/"

func validateOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем из буфера тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%v", err)

	}
	_ = ioutil.WriteFile("organisation.json", body, 0777)

	// Приводим тело к структуре
	var orgStruct *orgchecker.OrganizationConfig
	_ = json.Unmarshal(body, &orgStruct)

	// Выполняем валидацию данных в структуре и отправляем запрос на добавление организации в канал
	err = config.ValidateOrgConfig(orgStruct)
	if err != nil {
		log.Printf("Error while validate request body. %v", err)
		http.Error(w, "Bad Request", 400)
	} else {
		log.Printf("Validate org config")
		// Если валидация прошла запускаем подпись и добавление данных в канала
		err = addOrganization(orgStruct)
		if err != nil {
			// Если при добавлении что-то пошло не так то возвращаем 500 и выплевывает стактрейс в  STDOUT
			http.Error(w, "Something Crashed...", 500)
		} else {
			// Если всё прошло успешно возвращаем 200 код ответа от API
			http.Error(w, "Organization Added!", 200)
		}

	}
}

func validateOrganizationRemove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем из буфера тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Приводим тело к структуре
	var orgStruct *orgchecker.OrganizationRemove
	_ = json.Unmarshal(body, &orgStruct)

	err = config.ValidateOrgRemoveConfig(orgStruct)
	if err != nil {
		log.Printf("Error while validate request body. %v", err)
		http.Error(w, "Bad Request", 400)
	} else {
		log.Printf("Validate client name for remove organistation")

		err = removeOrganization(orgStruct)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
		} else {
			http.Error(w, "Successfull", 200)
		}
	}
}

func addOrganization(org *orgchecker.OrganizationConfig) error {

	// Данные для отладки
	log.Printf("Organization Data: %v", org)

	// Вызов метода подписи данных и внесения изменения в блок
	err := signer.SignAndAdd(org)
	if err != nil {
		log.Printf("Error while additing organisation to HLF. Stacktrace: %v", err)
		return err
	} else {
		log.Print("Organization successfull added to leger")
	}

	return nil
}

func removeOrganization(org *orgchecker.OrganizationRemove) error {
	err := signer.SignAndRemove(org)
	if err != nil {
		log.Printf("Error while delete organisation from HLM. Stacktrace: %v", err)
		return err
	} else {
		log.Printf("Organisation successful removed from ledger")
	}

	return nil
}

// InitialiseAPI метод для инициализации API сервера
func InitialiseAPI() {
	// Инициализация роутера
	router := mux.NewRouter().StrictSlash(true)

	// Добавляем базовые методы в API для добавления и удаления организации
	router.HandleFunc(APIEndpoint+"addorg", validateOrganization).Methods("POST")
	router.HandleFunc(APIEndpoint+"removeorg", validateOrganizationRemove).Methods("POST")

	// Крашим приложение если при инициализации роутера возникли ошибки
	log.Fatal(http.ListenAndServe(":8081", router))
}
