package models

import (
	"encoding/json"
	"fmt"
	redisClient "fyoukuApi/services/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int
	Comment            int
	UserId             int
	IsRecommend             int
}
type VideoData struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
	Comment       int
}

type Episodes struct {
	Id      int
	Title   string
	AddTime int64
	Num     int
	PlayUrl string
	Comment int
}

func init() {
	orm.RegisterModel(new(Video))
}

func GetChannelHotList(ChannelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time,img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_hot=1 AND channel_id=? ORDER BY episodes_update_time DESC LIMIT 9", ChannelId).QueryRows(&videos)

	return num, videos, err
}

func GetChannelRecommendRegionList(channelId, regionId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time,img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_recommend=1 AND region_id=? AND channel_id=? ORDER BY episodes_update_time DESC LIMIT 9", regionId, channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelRecommendTypeList(channelId, typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time,img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_recommend=1 AND type_id=? AND channel_id=? ORDER BY episodes_update_time DESC LIMIT 9", typeId, channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelVideoList(channelId, regionId, typeId int, end, sort string, offset, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params

	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" {
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		qs = qs.Filter("is_end", 1)
	}

	if sort == "episodesUpdateTime" {
		qs = qs.OrderBy("-episodes_update_time")
	} else if sort == "comment" {
		qs = qs.OrderBy("-comment")
	} else if sort == "addTime" {
		qs = qs.OrderBy("-add_time")
	} else {
		qs = qs.OrderBy("-add_time")
	}
	nums, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit, offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	return nums, videos, err
}

func GetVideoInfo(videoId int) (Video, error) {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("SELECT * FROM video WHERE id=? LIMIT 1", videoId).QueryRow(&video)
	return video, err
}

// RedisGetVideoInfo 增加redis缓存 - 获取视频详情
func RedisGetVideoInfo(videoId int) (Video, error) {
	var video Video
	conn := redisClient.PoolConnect()
	defer conn.Close()
	//定义redis key
	redisKey := "video:id:" + strconv.Itoa(videoId)
	//判断key是否存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(res, &video)
	} else {
		o := orm.NewOrm()
		err := o.Raw("SELECT * FROM video WHERE id=? LIMIT 1", videoId).QueryRow(&video)
		if err == nil {
			//保存redis
			_, err := conn.Do("hmset", redis.Args{redisKey}.AddFlat(video)...)
			if err == nil {
				conn.Do("expire", redisKey, 86400)
			}
		}
	}
	return video, err
}

// GetVideoEpisodesList 获取视频剧集列表
func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.Raw("SELECT id,title,add_time,num,play_url,comment FROM video_episodes WHERE video_id=? AND status=1 ORDER BY num ASC", videoId).QueryRows(&episodes)
	return num, episodes, err
}

// RedisGetVideoEpisodesList 增加redis缓存 - 获取视频剧集列表
func RedisGetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	var (
		episodes []Episodes
		num      int64
		err      error
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()
	redisKey := "video:episodes:videoId:" + strconv.Itoa(videoId)
	//判断redisKey是否已经存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		num, err = redis.Int64(conn.Do("llen", redisKey))
		if err == nil {
			values, _ := redis.Values(conn.Do("lrange", redisKey, "0", "-1"))
			var episodesInfo Episodes
			for _, v := range values {
				err = json.Unmarshal(v.([]byte), &episodesInfo)
				if err == nil {
					episodes = append(episodes, episodesInfo)
				}
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("SELECT id,title,add_time,num,play_url,comment FROM video_episodes WHERE video_id=? AND status=1 ORDER BY num ASC", videoId).QueryRows(&episodes)
		if err == nil {
			//遍历获取到到信息，把信息json化保存
			for _, v := range episodes {
				jsonValue, err := json.Marshal(v)
				if err == nil {
					//保存redis中
					conn.Do("rpush", redisKey, jsonValue)
				}
			}
			conn.Do("expire", redisKey, 86400)
		}
	}
	return num, episodes, err
}

// GetChannelTop 频道排行榜
func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end FROM video WHERE status=1 AND channel_id=? ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
	return num, videos, err
}

// RedisGetChannelTop 增加redis缓存 - 频道排行榜
func RedisGetChannelTop(channelId int) (int64, []VideoData, error) {
	var (
		videos []VideoData
		num    int64
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()
	redisKey := "video:top:channel:channelId:" + strconv.Itoa(channelId)
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		num = 0
		res, _ := redis.Values(conn.Do("zrevrange", redisKey, "0", "10", "WITHSCORES"))
		for k, v := range res {
			fmt.Println(string(v.([]byte)))
			if k%2 == 0 {
				videoId, err := strconv.Atoi(string(v.([]byte)))
				videoInfo, err := RedisGetVideoInfo(videoId)
				if err == nil {
					var videoDataInfo VideoData
					videoDataInfo.Id = videoInfo.Id
					videoDataInfo.Img = videoInfo.Img
					videoDataInfo.Img1 = videoInfo.Img1
					videoDataInfo.IsEnd = videoInfo.IsEnd
					videoDataInfo.SubTitle = videoInfo.SubTitle
					videoDataInfo.Title = videoInfo.Title
					videoDataInfo.AddTime = videoInfo.AddTime
					videoDataInfo.Comment = videoInfo.Comment
					videoDataInfo.EpisodesCount = videoInfo.EpisodesCount
					videos = append(videos, videoDataInfo)
					num++
				}
			}

		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end FROM video WHERE status=1 AND channel_id=? ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
		if err == nil {
			for _, v := range videos {
				conn.Do("zadd", redisKey, v.Comment, v.Id)
			}
			conn.Do("expire", redisKey, 86400*30)
		}
	}
	return num, videos, err
}

// GetTypeTop 类型排行榜
func GetTypeTop(typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end FROM video WHERE status=1 AND type_id=? ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
	return num, videos, err
}

// RedisGetTypeTop 增加redis缓存 - 类型排行榜
func RedisGetTypeTop(typeId int) (int64, []VideoData, error) {
	var (
		videos []VideoData
		num    int64
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()
	redisKey := "video:top:type:typeId:" + strconv.Itoa(typeId)
	exsits, err := redis.Bool(conn.Do("exists", redisKey))
	if exsits {
		num = 0
		res, _ := redis.Values(conn.Do("zrevrange", redisKey, "0", "10", "WITHSCORES"))
		for k, v := range res {
			if k%2 == 0 {
				videoId, err := strconv.Atoi(string(v.([]byte)))
				videoInfo, err := RedisGetVideoInfo(videoId)
				if err == nil {
					var videoDataInfo VideoData
					videoDataInfo.Id = videoInfo.Id
					videoDataInfo.Img = videoInfo.Img
					videoDataInfo.Img1 = videoInfo.Img1
					videoDataInfo.IsEnd = videoInfo.IsEnd
					videoDataInfo.SubTitle = videoInfo.SubTitle
					videoDataInfo.Title = videoInfo.Title
					videoDataInfo.AddTime = videoInfo.AddTime
					videoDataInfo.Comment = videoInfo.Comment
					videoDataInfo.EpisodesCount = videoInfo.EpisodesCount
					videos = append(videos, videoDataInfo)
					num++
				}
			}

		}

	} else {
		o := orm.NewOrm()
		num, err = o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end FROM video WHERE status=1 AND type_id=? ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
		if err == nil {
			for _, v := range videos {
				conn.Do("zadd", redisKey, v.Comment, v.Id)
			}
			conn.Do("expire", redisKey, 86400*30)
		}
	}
	return num, videos, err
}

func GetUserVideo(uid int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end FROM video WHERE user_id=? ORDER BY add_time DESC", uid).QueryRows(&videos)
	return num, videos, err

}

func SaveVideo(title, subTitle string, channelId, regionId, typeId int, playUrl string, userId int) error {
	o := orm.NewOrm()
	var video Video
	videoTime := time.Now().Unix()
	video.Title = title
	video.SubTitle = subTitle
	video.AddTime = videoTime
	video.Img = ""
	video.Img1 = ""
	video.EpisodesCount = 1
	video.IsEnd = 1
	video.ChannelId = channelId
	video.Status = 1
	video.RegionId = regionId
	video.TypeId = typeId
	video.EpisodesUpdateTime = int(videoTime)
	video.Comment = 0
	video.UserId = userId
	videoId, err := o.Insert(&video)
	if err == nil {
		o.Raw("INSERT INTO video_episodes (title,add_time,num,video_id,play_url,status,comment) VALUE (?,?,?,?,?,?,?)", subTitle, videoTime, 1, videoId, playUrl, 1, 0).Exec()

	}
	return err
}

// GetAllList 获取所以视频数据
func GetAllList() (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	num, err := o.Raw("SELECT id,title,sub_title,status,add_time,img,img1,channel_id,type_id,region_id,user_id,episodes_count,episodes_update_time,is_end,is_hot,is_recommend,comment FROM video").QueryRows(&videos)
	return num, videos, err
}
