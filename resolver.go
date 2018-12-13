//go:generate gorunpkg github.com/99designs/gqlgen

package gqlgen_example

import (
	context "context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/aneri/gqlgen-example/models"

	"github.com/aneri/gqlgen-example/dal"
)

var ctxt context.Context

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

// MiddleWareHandler to handle db connection
func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		crConn, err := dal.Connect()
		if err != nil {
			log.Fatal(err)
		}
		ctxt = context.WithValue(request.Context(), "crConn", crConn)
		next.ServeHTTP(writer, request.WithContext(ctxt))
	})
}
func (r *mutationResolver) CreateJob(ctx context.Context, input NewJob) (Job, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	jobs := models.Job{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: input.Description,
		Location:    input.Location,
		CreatedBy:   input.CreatedBy,
	}
	crConn.Db.Create(&jobs)
	return Job{jobs.ID, jobs.Name, jobs.Description, jobs.Location, jobs.Description}, nil
}
func (r *mutationResolver) DeleteJob(ctx context.Context, id string) (string, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	crConn.Db.Where("id=?", id).Delete(&models.Job{})
	return "job deleted successfully", nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Jobs(ctx context.Context) ([]Job, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	var personal []Job
	crConn.Db.Find(&personal)
	return personal, nil
}
