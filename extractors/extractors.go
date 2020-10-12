package extractors

import (
	"net/url"
	"strings"

	"github.com/maseer/annie/extractors/bcy"
	"github.com/maseer/annie/extractors/bilibili"
	"github.com/maseer/annie/extractors/douyin"
	"github.com/maseer/annie/extractors/douyu"
	"github.com/maseer/annie/extractors/facebook"
	"github.com/maseer/annie/extractors/geekbang"
	"github.com/maseer/annie/extractors/haokan"
	"github.com/maseer/annie/extractors/instagram"
	"github.com/maseer/annie/extractors/iqiyi"
	"github.com/maseer/annie/extractors/mgtv"
	"github.com/maseer/annie/extractors/miaopai"
	"github.com/maseer/annie/extractors/netease"
	"github.com/maseer/annie/extractors/pixivision"
	"github.com/maseer/annie/extractors/pornhub"
	"github.com/maseer/annie/extractors/qq"
	"github.com/maseer/annie/extractors/tangdou"
	"github.com/maseer/annie/extractors/tiktok"
	"github.com/maseer/annie/extractors/tumblr"
	"github.com/maseer/annie/extractors/twitter"
	"github.com/maseer/annie/extractors/types"
	"github.com/maseer/annie/extractors/udn"
	"github.com/maseer/annie/extractors/universal"
	"github.com/maseer/annie/extractors/vimeo"
	"github.com/maseer/annie/extractors/weibo"
	"github.com/maseer/annie/extractors/xvideos"
	"github.com/maseer/annie/extractors/yinyuetai"
	"github.com/maseer/annie/extractors/youku"
	"github.com/maseer/annie/extractors/youtube"
	"github.com/maseer/annie/utils"
)

var extractorMap map[string]types.Extractor

func init() {
	douyinExtractor := douyin.New()
	youtubeExtractor := youtube.New()

	extractorMap = map[string]types.Extractor{
		"": universal.New(), // universal extractor

		"douyin":     douyinExtractor,
		"iesdouyin":  douyinExtractor,
		"bilibili":   bilibili.New(),
		"bcy":        bcy.New(),
		"pixivision": pixivision.New(),
		"youku":      youku.New(),
		"youtube":    youtubeExtractor,
		"youtu":      youtubeExtractor, // youtu.be
		"iqiyi":      iqiyi.New(),
		"mgtv":       mgtv.New(),
		"tangdou":    tangdou.New(),
		"tumblr":     tumblr.New(),
		"vimeo":      vimeo.New(),
		"facebook":   facebook.New(),
		"douyu":      douyu.New(),
		"miaopai":    miaopai.New(),
		"163":        netease.New(),
		"weibo":      weibo.New(),
		"instagram":  instagram.New(),
		"twitter":    twitter.New(),
		"qq":         qq.New(),
		"yinyuetai":  yinyuetai.New(),
		"geekbang":   geekbang.New(),
		"pornhub":    pornhub.New(),
		"xvideos":    xvideos.New(),
		"udn":        udn.New(),
		"tiktok":     tiktok.New(),
		"haokan":     haokan.New(),
	}
}

// Extract is the main function to extract the data.
func Extract(u string, option types.Options) ([]*types.Data, error) {
	u = strings.TrimSpace(u)
	var domain string

	bilibiliShortLink := utils.MatchOneOf(u, `^(av|BV|ep)\w+`)
	if len(bilibiliShortLink) > 1 {
		bilibiliURL := map[string]string{
			"av": "https://www.bilibili.com/video/",
			"BV": "https://www.bilibili.com/video/",
			"ep": "https://www.bilibili.com/bangumi/play/",
		}
		domain = "bilibili"
		u = bilibiliURL[bilibiliShortLink[1]] + u
	} else {
		u, err := url.ParseRequestURI(u)
		if err != nil {
			return nil, err
		}
		if u.Host == "haokan.baidu.com" {
			domain = "haokan"
		} else {
			domain = utils.Domain(u.Host)
		}
	}
	extractor := extractorMap[domain]
	videos, err := extractor.Extract(u, option)
	if err != nil {
		return nil, err
	}
	for _, v := range videos {
		v.FillUpStreamsData()
	}
	return videos, nil
}
