package model

type Model struct {
	Object  Object
	Details []Detail
}

type Object interface {
	ObjectID() string
}

type Object1 struct {
	ID   string `dynamo:"id"`
	Name string `dynamo:"name"`
	Age  int    `dynamo:"age"`
}

func (o *Object1) ObjectID() string {
	return o.ID
}

type Detail interface {
	DetailID() string
	ObjectID() string
}

type Detail1 struct {
	ID     string `dynamo:"id"`
	Object string `dynamo:"object_id"`
	Job    string `dynamo:"job"`
}

func (d *Detail1) DetailID() string {
	return d.ID
}
func (d *Detail1) ObjectID() string {
	return d.Object
}

type Detail2 struct {
	ID         string `dynamo:"id"`
	Object     string `dynamo:"object_id"`
	MiddleName string `dynamo:"middle_name"`
}

func (d *Detail2) DetailID() string {
	return d.ID
}

func (d *Detail2) ObjectID() string {
	return d.Object
}
