package main

import (
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"lris/excel"
	"lris/model"
	"net/http"
	"net/url"
	"regexp"
)

const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"
const cookie = "mid=YGbYigALAAH40gS4WJtW-52aA7xl; ig_did=64DFF4C1-EDE4-4B60-85BB-8EDE470BA9F1; ig_nrcb=1; csrftoken=Gka4J9xgAejtNWAx83SsK0EbIrfad4II; rur=PRN; ds_user_id=18070022789; sessionid=18070022789%3AH4KzTatge28jx3%3A18"

func main() {
	app := iris.Default()
	app.Use(myMiddleware)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		edges, done := getInsData(ctx, app)
		if done {
			return
		}
		saveDataAsFile(edges)
		ctx.SendFile("test.xlsx", "test.xlsx")
	})

	app.Run(iris.Addr(":8080"))
}

func saveDataAsFile(edges []model.EdgeOwnerToTimelineMediaEdge) {
	xlsx := excelize.NewFile()
	excel.SetTitle(xlsx, "ID", "类型", "评论数", "点赞数", "封面地址")
	for i, edge := range edges {
		node := edge.Node
		excel.AddRow(xlsx, i+2, node.ID, node.Typename, node.EdgeMediaToComment.Count, node.EdgeMediaPreviewLike.Count, node.DisplayURL)
	}
	xlsx.SaveAs("test.xlsx")
}

func getInsData(ctx iris.Context, app *iris.Application) ([]model.EdgeOwnerToTimelineMediaEdge, bool) {
	url := ctx.URLParam("url")
	data, err := model.UnmarshalSharedData(getSharedDataRes(url))
	if err != nil {
		app.Logger().Error(err)
		return nil, true
	}
	page := data.EntryData.ProfilePage[0]
	user := page.Graphql.User
	media := user.EdgeOwnerToTimelineMedia
	edges := media.Edges
	if media.PageInfo.HasNextPage {
		nextData := getNextData(user, media.PageInfo)
		edges = append(edges, nextData...)
	}
	return edges, false
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func getSharedDataRes(url string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("user-agent", UserAgent)
	request.Header.Add("cookie", cookie)
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	r, _ := regexp.Compile("window._sharedData = (.*?);</script>")
	return r.FindSubmatch(body)[1]
}

func getNextData(user model.GraphqlUser, pageInfo model.PageInfo) []model.EdgeOwnerToTimelineMediaEdge {
	id := user.ID
	queryUrl := getQueryUrl(id, pageInfo)
	data := getQueryData(queryUrl)
	edges := data.Edges
	for data.PageInfo.HasNextPage {
		queryUrl = getQueryUrl(id, data.PageInfo)
		data = getQueryData(queryUrl)
		edges = append(edges, data.Edges...)
	}
	return edges
}

func getQueryUrl(id string, pageInfo model.PageInfo) string {
	marshal, err := json.Marshal(map[string]interface{}{"id": id, "first": 50, "after": pageInfo.EndCursor})
	if err != nil {
		panic(err)
	}
	url := "https://www.instagram.com/graphql/query/?query_hash=42d2750e44dbac713ff30130659cd891&variables=" + url.QueryEscape(string(marshal))
	return url
}

func getQueryData(url string) model.EdgeOwnerToTimelineMedia {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("user-agent", UserAgent)
	request.Header.Add("cookie", cookie)
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	data, err := model.UnmarshalQueryPage(body)
	return data.Data.User.EdgeOwnerToTimelineMedia
}
