package Twitter

import "fmt"

type Twitte struct {
	Uid  int
	Id   int
	Next *Twitte
}

type User struct {
	follow map[int]bool
}

type Twitter struct {
	UserFollow map[int]*User
	Last       *Twitte
}

func Constructor() Twitter {
	t := new(Twitter)
	t.Last = new(Twitte)
	t.Last.Id = -1
	t.Last.Next = nil
	t.UserFollow = make(map[int]*User)
	return *t
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	newT := new(Twitte)
	newT.Id = tweetId
	newT.Uid = userId
	newT.Next = this.Last
	this.Last = newT
	if this.UserFollow[userId] == nil {
		newU := new(User)
		newU.follow = make(map[int]bool)
		newU.follow[userId] = true
		this.UserFollow[userId] = newU
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {

	head, result := this.Last, make([]int, 0)
	if this.UserFollow[userId] == nil {
		return result
	}

	for head.Next != nil || len(result) == 10 {
		fmt.Println(head.Id)
		if this.UserFollow[userId].follow[head.Uid] {
			result = append(result, head.Id)
		}
		head = head.Next
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.UserFollow[followerId] == nil {
		newU := new(User)
		newU.follow = make(map[int]bool)
		newU.follow[followeeId] = true
		newU.follow[followerId] = true
		this.UserFollow[followerId] = newU
	}
	this.UserFollow[followerId].follow[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.UserFollow[followerId] != nil {
		this.UserFollow[followerId].follow[followeeId] = false
	}
}
