package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type Tweet struct {
	Id		int
	Timestamp	int
}

type Twitter struct {
	// key为UID, value为map[int]bool作为Set存储某个UID的追随者
	Follower	map[int]map[int]bool
	// UID所发的所有Tweet
	Tweet		map[int][]*Tweet
	// 序列号，模拟timestamp打到Tweet上
	SerialNo	int
	Limit		int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Follower:	make(map[int]map[int]bool),
		Tweet:		make(map[int][]*Tweet),
		SerialNo:	0,
		Limit:		10,
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	t := &Tweet{
		Id: tweetId,
		Timestamp: this.SerialNo,
	}
	this.SerialNo += 1

	if _, ok := this.Tweet[userId]; !ok {
		// 第一次发推
		this.Tweet[userId] = []*Tweet{}
		// 自己需要能够看到自己发的推，所以需要follow一下自己
		this.Follow(userId, userId)
	}
	this.Tweet[userId] = append(this.Tweet[userId], t)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	pq := datastruct.NewPQ([]int{}, []int{}, true)
	// 遍历userId关注的用户
	for k, _ := range this.Follower[userId] {
		tweets := this.Tweet[k]
		//将该用户的发推记录加入PQ，只保留最近10条
		for i := len(tweets) - 1; i >= 0; i-- {
			pq.Push(tweets[i].Id, tweets[i].Timestamp)
			if pq.Size() > this.Limit {
				pq.Pop()
			}
		}
	}

	// 从PQ中提取内容并返回
	rslt := []int{}
	for pq.Size() != 0 {
		rslt = append([]int{pq.Pop()}, rslt...)
	}
	return rslt
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if _, ok := this.Follower[followerId]; !ok {
		this.Follower[followerId] = make(map[int]bool)
	}
	this.Follower[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if _, ok := this.Follower[followerId][followeeId]; !ok {
		return
	}

	if followerId == followeeId {
		return
	}

	delete(this.Follower[followerId], followeeId)
}

func main() {
	obj := Constructor();
	obj.PostTweet(1, 5);
	fmt.Println(obj.GetNewsFeed(1))
	obj.Follow(1, 2)
	obj.PostTweet(2, 6);
	obj.PostTweet(2, 8);
	obj.PostTweet(1, 15);
	fmt.Println(obj.GetNewsFeed(1))
	obj.Unfollow(1, 2);
	fmt.Println(obj.GetNewsFeed(1))
}
