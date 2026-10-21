package graphql

import (
	"employee-management/model"
	"employee-management/service"
	"fmt"
	gql "github.com/graphql-go/graphql"
	"strings"
)

type Server struct {
	service service.EmployeeService
	schema  gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	x := &Server{service: s}
	sc, err := x.buildSchema()
	if err != nil {
		return nil, err
	}
	x.schema = sc
	return x, nil
}
func (x *Server) buildSchema() (gql.Schema, error) {
	emp := gql.NewObject(gql.ObjectConfig{Name: "Employee", Fields: gql.Fields{
		"id": &gql.Field{Type: gql.NewNonNull(gql.Int)}, "name": &gql.Field{Type: gql.NewNonNull(gql.String)},
		"email": &gql.Field{Type: gql.NewNonNull(gql.String)}, "department": &gql.Field{Type: gql.NewNonNull(gql.String)},
		"salary": &gql.Field{Type: gql.NewNonNull(gql.Float)},
	}})
	input := gql.NewInputObject(gql.InputObjectConfig{Name: "EmployeeInput", Fields: gql.InputObjectConfigFieldMap{
		"name": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)}, "email": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)}, "salary": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
	}})
	query := gql.NewObject(gql.ObjectConfig{Name: "Query", Fields: gql.Fields{
		"employee": &gql.Field{Type: emp, Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			e, err := x.service.GetByID(uint(p.Args["id"].(int)))
			if service.IsNotFound(err) {
				return nil, fmt.Errorf("employee not found")
			}
			return e, err
		}},
		"employees": &gql.Field{Type: gql.NewList(emp), Resolve: func(p gql.ResolveParams) (interface{}, error) { return x.service.GetAll() }},
	}})
	mutation := gql.NewObject(gql.ObjectConfig{Name: "Mutation", Fields: gql.Fields{
		"createEmployee": &gql.Field{Type: emp, Args: gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(input)}}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			e := fromInput(p.Args["input"].(map[string]interface{}))
			if err := x.service.Create(&e); err != nil {
				return nil, err
			}
			return e, nil
		}},
		"updateEmployee": &gql.Field{Type: emp, Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}, "input": &gql.ArgumentConfig{Type: gql.NewNonNull(input)}}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			e := fromInput(p.Args["input"].(map[string]interface{}))
			return x.service.Update(uint(p.Args["id"].(int)), &e)
		}},
		"deleteEmployee": &gql.Field{Type: gql.NewNonNull(gql.Boolean), Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if err := x.service.Delete(uint(p.Args["id"].(int))); err != nil {
				if service.IsNotFound(err) {
					return false, fmt.Errorf("employee not found")
				}
				return false, err
			}
			return true, nil
		}},
	}})
	return gql.NewSchema(gql.SchemaConfig{Query: query, Mutation: mutation})
}
func fromInput(m map[string]interface{}) model.Employee {
	return model.Employee{Name: strings.TrimSpace(fmt.Sprint(m["name"])), Email: strings.TrimSpace(fmt.Sprint(m["email"])), Department: strings.TrimSpace(fmt.Sprint(m["department"])), Salary: m["salary"].(float64)}
}
