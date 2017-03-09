package main

import "fmt"

var currentId int
var todos Todos

func init() {
  RepoCreateTodo(Todo{Name: "Write presentation"})
  RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
  for _, t := range todos {
    if t.Id == id {
      return t
    }
  }

  return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
  currentId += 1
  t.Id = currentId
  todos = append(todos, t)
  return t
}

func RepoDestroyTodo(id int) error {
  for _, t := range todos {
    if t.Id == id {
      // this is how we remove from a slice
      todos = append(todos[:id], todos[id+1:]...)
      return nil
    }
  }

  return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}