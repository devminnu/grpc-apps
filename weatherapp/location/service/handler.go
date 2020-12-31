package service

import (
	"github.com/gin-gonic/gin"
)

func getLocation(c *gin.Context) {
	// Multipart form
	// form, _ := c.MultipartForm()
	// files := form.File["upload[]"]

	// for _, file := range files {
	// 	log.Println(file.Filename)

	// 	// Upload the file to specific dst.
	// 	c.SaveUploadedFile(file, dst)
	// }
	// c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

/* // Create a user
func createUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)

		var s db.User
		err := decoder.Decode(&s)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = deps.DB.Create(req.Context(), s)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	})
} */
