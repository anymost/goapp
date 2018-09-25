package main

import (
	"runtime"
	"os"
	"bufio"
	"net/http"
	"strconv"
	"io"
	"math"
	"strings"
	"fmt"
	"regexp"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"log"
)

var (
	GoroutineCount = runtime.NumCPU()
	FollowerMap    = make(map[int]int)
	StarMap        = make(map[int]int)
	Repository     = "https://github.com/vuejs/vue/stargazers?page="
	FollowerArrayFileName  = "./vue-followers.json"
	StarArrayFileName      = "./vue-stars.json"
)

type User struct {
	Name      string
	Followers int
	Stars     int
}

type Account struct {
	Follower int
	Count int
}

func (user *User) sayName()  {
	fmt.Print(user.Name)
}


func main() {
	runtime.GOMAXPROCS(GoroutineCount)
	fmt.Println("running on " + strconv.Itoa(GoroutineCount) + " goroutine")

    pageIndex := 1
    for {
		users, isEnd := RequestPage(pageIndex)
		if isEnd == false {
			log.Fatal("download complete")
			break
		} else {
			CollectUsers(users)
		}
		log.Fatal("page " + strconv.Itoa(pageIndex))
		pageIndex++
	}
	WriteFile()
}



func RequestPage(pageIndex int)([]User, bool) {
	resp, err := http.Get(Repository + strconv.Itoa(pageIndex))
	if resp.StatusCode != http.StatusOK || err != nil {
		return nil, false
	}
	defer resp.Body.Close()
	usersInfo := ExtractUserInfo(resp.Body)
	usersMatrix := SliceUsers(usersInfo)
	users := make([]*[]User, 0)
	ch := make(chan *[]User, GoroutineCount)
	for _, userSlice := range usersMatrix{
		go ExtractUsers(userSlice, ch)
	}
	for i := 0; i < GoroutineCount; i++ {
		followerSlice := <- ch
		log.Fatal(followerSlice)
		users = append(users, followerSlice)
	}
	log.Fatal(users)
	return platMatrix(users), true
}

func platMatrix(usersMatrix []*[]User)(users []User) {
	users = make([]User, 0)
	for _, userSlice := range usersMatrix {
		for _, user := range *userSlice {
			users = append(users, user)
		}
	}
	return
}

func CollectUsers(followers []User)  {
	for _, follower := range followers {
		if follower.Name != "" {
			followerCount := follower.Followers
			starCount     := follower.Stars
			if FollowerMap[followerCount] == 0 {
				FollowerMap[followerCount] = 1
			} else {
				FollowerMap[followerCount] += 1
			}

			if StarMap[starCount] == 0 {
				StarMap[starCount] = 1
			} else {
				StarMap[starCount] += 1
			}
		}
	}
}

func SplitUsers()(followersMap map[string][]int, starMap map[string][]int)  {
	log.Fatal("start split users")
	followersMap = map[string][]int{"follower": make([]int, 0), "count": make([]int, 0)}
	starMap      = map[string][]int{"star": make([]int, 0), "count": make([]int, 0)}
	for follower, count := range FollowerMap {
		followersMap["follower"] = append(followersMap["follower"], follower)
		followersMap["count"] = append(followersMap["count"], count)
	}
	for star, count := range FollowerMap {
		starMap["star"] = append(starMap["star"], star)
		starMap["count"] = append(starMap["count"], count)
	}
	return
}

func WriteFile()  {
	log.Fatal("start writer files")
	followersMap, starsMap := SplitUsers()

	followerFile, _ := os.OpenFile(FollowerArrayFileName, os.O_WRONLY | os.O_CREATE, 0666)
	defer followerFile.Close()
	followerWriter := bufio.NewWriter(followerFile)


	starFile, _ := os.OpenFile(StarArrayFileName, os.O_WRONLY | os.O_CREATE, 0666)
	defer starFile.Close()
	starWriter := bufio.NewWriter(starFile)

	byteValue, _ := json.Marshal(followersMap)
	followerWriter.WriteString(string(byteValue))

	byteValue, _ = json.Marshal(starsMap)
	starWriter.WriteString(string(byteValue))


}

func ExtractUserInfo(body io.ReadCloser)[]string  {
	content, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return make([]string, 0)
	}
	linkElements := content.Find(".css-truncate > a")
	linkHrefs := make([]string, len(linkElements.Nodes))
	linkElements.Each(func(i int, selection *goquery.Selection) {
		value, ok := selection.Attr("href")
		if ok {
			linkHrefs[i] = value
		}
	})
	return linkHrefs
}

func SliceUsers(users []string)[][]string  {
	userMatrix := make([][]string, GoroutineCount)
	singleLength := int(math.Ceil(float64(len(users)) / float64(GoroutineCount)))
	for i := range userMatrix {
		userMatrix[i] = make([]string, singleLength)
	}
	for i := 0; i < GoroutineCount; i++ {
		arrayLength := (i + 1) * singleLength
		if arrayLength > len(users) - 1{
			arrayLength = len(users) - 1
		}
		userMatrix[i] = users[i * singleLength: arrayLength ]
	}
	return userMatrix
}



func ExtractUsers(links []string, ch chan<- *[]User)  {
	followers := make([]User, len(links))
	for index, link := range links {
		resp, err := http.Get("https://github.com" + link)
		if err != nil {
			resp.Body.Close()
			followers[index] = User{}
			continue
		}
		content, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			resp.Body.Close()
			followers[index] = User{}
			continue
		}

		follower := trimFollower(link, content)
		followers[index] = follower
		resp.Body.Close()
	}
	ch <- &followers
}

func trimFollower(link string, content *goquery.Document) User {
	followElement := content.Find("a[title='Followers'] > .Counter")
	startElement := content.Find("a[title='Stars'] > .Counter")
	followStr, _ := (*followElement).Html()
	startStr, _ := (*startElement).Html()

	r, _ := regexp.Compile("([0-9]+)")
	followMatch := r.FindString(followStr)
	follow, _ := strconv.Atoi(followMatch)
	starMatch := r.FindString(startStr)
	star, _ := strconv.Atoi(starMatch)

	name := strings.Trim(link, "/")
	return User{name, follow, star}
}








