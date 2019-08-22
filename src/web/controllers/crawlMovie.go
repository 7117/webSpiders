package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"github.com/astaxie/beego/httplib"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	sUrl := "https://movie.douban.com/subject/25827935/"  
	rsp := httplib.Get(sUrl);
	
	sMovieHtml,err := rsp.String()
	if err != nil{
		panic(err)
	}

	var movieInfo models.MovieInfo

	movieInfo.Movie_name            = models.GetMovieName(sMovieHtml)
	movieInfo.Movie_director        = models.GetMovieDirector(sMovieHtml)
	movieInfo.Movie_main_character  = models.GetMovieMainCharacters(sMovieHtml) 
	movieInfo.Movie_type            = models.GetMovieGenre(sMovieHtml)
	movieInfo.Movie_on_time         = models.GetMovieOnTime(sMovieHtml)
	movieInfo.Movie_grade           = models.GetMovieGrade(sMovieHtml)
	movieInfo.Movie_span            = models.GetMovieRunningTime(sMovieHtml)
  
	
  }
  