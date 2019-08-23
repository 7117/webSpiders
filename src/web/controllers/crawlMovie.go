package controllers

import (
	// "fmt"
	// "strconv"
	"time"
	"github.com/astaxie/beego"
	"web/models"
	"github.com/astaxie/beego/httplib"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	var movieInfo models.MovieInfo

	// 连接到redis
	models.ConnectRedis("127.0.0.1:6379");

	sUrl := "https://movie.douban.com/subject/25827935/"
	models.PutinQueue(sUrl)

	for{
		var length int = models.GetQueueLength()

		if length == 0{
			break;
		}


		sUrl = models.PopfromQueue()
		
		//我们应当判断sUrl是否应该被访问过
		if models.IsVisit(sUrl) {
			continue
		}

		rsp := httplib.Get(sUrl);
	
		sMovieHtml,_ := rsp.String()
		if err != nil{
			panic(err)
		}

		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		//记录电影信息
		if movieInfo.Movie_name != "" {
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieRunningTime(sMovieHtml)
			
			time.Sleep(time.Second)
			models.AddMovie(&movieInfo)
			time.Sleep(time.Second)
	
		}

		//提取该页面的所有连接
		var urls []string
		urls = models.GetMovieUrls(sMovieHtml)


		for _, url := range urls {
			models.PutinQueue(url)
			c.Ctx.WriteString("出来")
		}

		//sUrl 应当记录到 访问set中
		models.AddToSet(sUrl)
		time.Sleep(time.Second)
	}

	c.Ctx.WriteString("end of crawl!")

}
		