package ds_test

import (
	"testing"

	"github.com/maestre3d/ds"
)

func TestLinkedList_Append(t *testing.T) {
	var list ds.List[string] = ds.NewDoublyLinkedList[string]()
	list.Append("foo")
	list.Append("bar")
	list.Append("baz")
	list.Append("bazinga")
	t.Log(list.Len())
	list.RemoveAt(-1)
	t.Log(list.Len())
	list.InsertAt(1, "elver")
	//t.Log(list.GetAt(1))
	//    var list ds.List[string] = ds.NewSinglyLinkedList[string]()

	//	list.Append("foo")
	//    list.Append("bar")
	//    list.Append("baz")
	//    list.InsertAt(0, "bazinga")
	//    list.InsertAt(1, "foobar")
	//    list.RemoveAt(1)
	//
	// SELECT a.full_name, a.age, a.picture_url, m.display_name, d.full_name, d.picture_url FROM actors as a
	// INNER JOIN movies as m AND directors as d WHERE a.movie_id = 101 AND a.gender = 'male' AND a.movie_id = m.movie_id
	// AND m.director_id = d.director_id;
	//
	// CREATE MATERIALIZED VIEW get_male_actors_from_movie_view AS
	// SELECT a.full_name, a.birthdate, a.picture_url, m.display_name, d.full_name, d.picture_url FROM actors as a
	// INNER JOIN movies as m AND directors as d WHERE a.gender = 'male' AND a.movie_id = m.movie_id
	// AND m.director_id = d.director_id;
	//
	// SELECT * FROM get_male_actors_from_movie_view WHERE movie_id = 101;
	for list.HasNext() {
		t.Log(list.Next())
	}
}
