package key_value

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/cemezgn/keyValueApp/pkg/server"
)

var (
	listKeysRequest   = regexp.MustCompile(`^\/keys[\/]*$`)
	getKeysRequest    = regexp.MustCompile(`^\/keys\/(\d+)$`)
	createKeysRequest = regexp.MustCompile(`^\/keys[\/]*$`)
)

type Service struct {
	Repository KeyValueRepository
}

type KeyValueRepository interface {
	List() ([]byte, error)
	Get(value string) (Item, bool)
	Create(i Item) ([]byte, error)
}

func NewService(repository KeyValueRepository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listKeysRequest.MatchString(r.URL.Path):
		s.List(w, r)
		return
	case r.Method == http.MethodGet && getKeysRequest.MatchString(r.URL.Path):
		s.Get(w, r)
		return
	case r.Method == http.MethodPost && createKeysRequest.MatchString(r.URL.Path):
		s.Create(w, r)
		return
	default:
		server.NotFound(w)
		return
	}
}

func (s *Service) List(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := s.Repository.List()

	if err != nil {
		server.NotFound(w)
		return
	}

	fmt.Println(r.Method, r.Host ,r.RequestURI)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	matches := getKeysRequest.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		server.NotFound(w)
		return
	}

	item, ok := s.Repository.Get(matches[1])

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("item with key not found"))
		return
	}
	jsonBytes, err := json.Marshal(item)
	if err != nil {
		server.InternalServerError(w)
		return
	}

	fmt.Println(r.Method, r.Host ,r.RequestURI)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	var i Item

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		server.InternalServerError(w)
		return
	}

	jsonBytes, err := s.Repository.Create(i)
	if err != nil {
		server.InternalServerError(w)
		return
	}

	fmt.Println(r.Method, r.Host ,r.RequestURI)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
