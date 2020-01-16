package struct_test

import "testing"

type Movie struct {
	movie_name   string
	movie_length int
	movie_id     int
}

func TestMovie(t *testing.T) {
	var kung_fu Movie
	var zu_qiu Movie
	kung_fu.movie_name = "功夫"
	kung_fu.movie_id = 1
	kung_fu.movie_length = 120
	zu_qiu.movie_name = "足球"
	zu_qiu.movie_id = 2
	zu_qiu.movie_length = 300
	hei_ba := Movie{"嘿吧", 120, 1}

	t.Log(kung_fu)
	t.Log(zu_qiu)
	t.Log(hei_ba)
}
