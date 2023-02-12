package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"nothing-behind.com/sample_gin/features/vocabulary/infra/postgres"
	"nothing-behind.com/sample_gin/features/vocabulary/infra/postgres/repository"
	"nothing-behind.com/sample_gin/features/vocabulary/usecase"
	"nothing-behind.com/sample_gin/utils/test_utils"
	"testing"
)

func TestListVocabulariesHandler(t *testing.T) {
	db, err := postgres.InitForTest()
	if err != nil {
		return
	}
	//postgres, err := postgres2.New()
	if err != nil {
		log.Fatal("fail init database")
	}

	vocabularyRepository := repository.NewVocabularyRepository(db)
	ListCategoriesUc := usecase.NewListCategories(db, vocabularyRepository)

	funcToTest := ListVocabulariesHandler(ListCategoriesUc)

	test_utils.RunHandlerTest(t, []test_utils.HandlerTestCase{
		{
			Name: "ok - get list vocabularies case non level",
			PrepareFunc: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(
					"GET",
					"/vocabularies",
					bytes.NewReader([]byte(``)),
				)
				return c, w
			},
			ExecHandlerFunc: funcToTest,
			WantCode:        http.StatusOK,
			WantFunc: func(w *httptest.ResponseRecorder) {
				var m usecase.ListOutput
				err := json.Unmarshal(w.Body.Bytes(), &m)
				if err != nil {
					return
				}
				assert.EqualValues(t, 30, len(*m.Vocabularies))
			},
		},
		{
			Name: "ok - get list vocabularies case level=1",
			PrepareFunc: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(
					"GET",
					"/vocabularies?level=1",
					bytes.NewReader([]byte(``)),
				)
				return c, w
			},
			ExecHandlerFunc: funcToTest,
			WantCode:        http.StatusOK,
			WantFunc: func(w *httptest.ResponseRecorder) {
				var m usecase.ListOutput
				err := json.Unmarshal(w.Body.Bytes(), &m)
				if err != nil {
					return
				}

				// 取得したvocのレベルが全て1
				for _, voc := range *m.Vocabularies {
					assert.EqualValues(t, "1", voc.Level)
				}
			},
		},
		{
			Name: "ok - get list vocabularies case level=5",
			PrepareFunc: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(
					"GET",
					"/vocabularies?level=5",
					bytes.NewReader([]byte(``)),
				)
				return c, w
			},
			ExecHandlerFunc: funcToTest,
			WantCode:        http.StatusOK,
			WantFunc: func(w *httptest.ResponseRecorder) {
				var m usecase.ListOutput
				err := json.Unmarshal(w.Body.Bytes(), &m)
				if err != nil {
					return
				}
				// 取得したvocのレベルが全て5
				for _, voc := range *m.Vocabularies {
					assert.EqualValues(t, "5", voc.Level)
				}
			},
		},
		{
			Name: "ok - get list vocabularies case level=10",
			PrepareFunc: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(
					"GET",
					"/vocabularies?level=10",
					bytes.NewReader([]byte(``)),
				)
				return c, w
			},
			ExecHandlerFunc: funcToTest,
			WantCode:        http.StatusOK,
			WantFunc: func(w *httptest.ResponseRecorder) {
				var m usecase.ListOutput
				err := json.Unmarshal(w.Body.Bytes(), &m)
				if err != nil {
					return
				}
				// 取得したvocのレベルが全て10
				for _, voc := range *m.Vocabularies {
					assert.EqualValues(t, "10", voc.Level)
				}
			},
		},
		{
			Name: "ok - get list vocabularies case not exist level=111",
			PrepareFunc: func() (*gin.Context, *httptest.ResponseRecorder) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(
					"GET",
					"/vocabularies?level=111",
					bytes.NewReader([]byte(``)),
				)
				return c, w
			},
			ExecHandlerFunc: funcToTest,
			WantCode:        http.StatusOK,
			WantFunc: func(w *httptest.ResponseRecorder) {
				// jsonの中身が空であること
				assert.JSONEq(t, `{"vocabularies": []}`, string(w.Body.Bytes()))
			},
		},
	})
}
