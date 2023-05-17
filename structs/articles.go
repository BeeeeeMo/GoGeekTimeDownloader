package structs

type Article struct {
	Error []any `json:"error"`
	Extra []any `json:"extra"`
	Data  struct {
		TextReadVersion int    `json:"text_read_version"`
		AudioSize       int    `json:"audio_size"`
		ArticleCover    string `json:"article_cover"`
		Subtitles       []any  `json:"subtitles"`
		ProductType     string `json:"product_type"`
		ReadingTime     int    `json:"reading_time"`
		IsFinished      bool   `json:"is_finished"`
		Like            struct {
			HadDone bool `json:"had_done"`
			Count   int  `json:"count"`
		} `json:"like"`
		AudioTime            string `json:"audio_time"`
		VideoHeight          int    `json:"video_height"`
		ArticleContent       string `json:"article_content"`
		FloatQrcode          string `json:"float_qrcode"`
		ArticleCoverHidden   bool   `json:"article_cover_hidden"`
		IsRequired           bool   `json:"is_required"`
		Score                string `json:"score"`
		LikeCount            int    `json:"like_count"`
		ArticleSubtitle      string `json:"article_subtitle"`
		AudioDownloadURL     string `json:"audio_download_url"`
		HadViewed            bool   `json:"had_viewed"`
		ArticleTitle         string `json:"article_title"`
		ColumnBgcolor        string `json:"column_bgcolor"`
		OfflinePackage       string `json:"offline_package"`
		AudioTitle           string `json:"audio_title"`
		VideoSize            int    `json:"video_size"`
		TextReadPercent      int    `json:"text_read_percent"`
		Cid                  int    `json:"cid"`
		ArticleCshort        string `json:"article_cshort"`
		VideoWidth           int    `json:"video_width"`
		ColumnCouldSub       bool   `json:"column_could_sub"`
		VideoID              string `json:"video_id"`
		Sku                  string `json:"sku"`
		VideoCover           string `json:"video_cover"`
		AuthorName           string `json:"author_name"`
		ColumnIsOnboard      bool   `json:"column_is_onboard"`
		InlineVideoSubtitles []any  `json:"inline_video_subtitles"`
		AudioURL             string `json:"audio_url"`
		ChapterID            string `json:"chapter_id"`
		ColumnHadSub         bool   `json:"column_had_sub"`
		ColumnCover          string `json:"column_cover"`
		Neighbors            struct {
			Left struct {
				ArticleTitle string `json:"article_title"`
				ID           int    `json:"id"`
			} `json:"left"`
			Right struct {
				ArticleTitle string `json:"article_title"`
				ID           int    `json:"id"`
			} `json:"right"`
		} `json:"neighbors"`
		RatePercent     int `json:"rate_percent"`
		FooterCoverData struct {
			ImgURL  string `json:"img_url"`
			MpURL   string `json:"mp_url"`
			LinkURL string `json:"link_url"`
		} `json:"footer_cover_data"`
		FloatAppQrcode     string `json:"float_app_qrcode"`
		ColumnIsExperience bool   `json:"column_is_experience"`
		Rate               struct {
			Num1 struct {
				CurVersion     int  `json:"cur_version"`
				MaxRate        int  `json:"max_rate"`
				CurRate        int  `json:"cur_rate"`
				IsFinished     bool `json:"is_finished"`
				TotalRate      int  `json:"total_rate"`
				LearnedSeconds int  `json:"learned_seconds"`
			} `json:"1"`
			Num2 struct {
				CurVersion     int  `json:"cur_version"`
				MaxRate        int  `json:"max_rate"`
				CurRate        int  `json:"cur_rate"`
				IsFinished     bool `json:"is_finished"`
				TotalRate      int  `json:"total_rate"`
				LearnedSeconds int  `json:"learned_seconds"`
			} `json:"2"`
			Num3 struct {
				CurVersion     int  `json:"cur_version"`
				MaxRate        int  `json:"max_rate"`
				CurRate        int  `json:"cur_rate"`
				IsFinished     bool `json:"is_finished"`
				TotalRate      int  `json:"total_rate"`
				LearnedSeconds int  `json:"learned_seconds"`
			} `json:"3"`
		} `json:"rate"`
		ProductID       int    `json:"product_id"`
		HadLiked        bool   `json:"had_liked"`
		ID              int    `json:"id"`
		FreeGet         bool   `json:"free_get"`
		IsVideoPreview  bool   `json:"is_video_preview"`
		ArticleSummary  string `json:"article_summary"`
		ColumnSaleType  int    `json:"column_sale_type"`
		FloatQrcodeJump string `json:"float_qrcode_jump"`
		IPAddress       string `json:"ip_address"`
		ColumnID        int    `json:"column_id"`
		Offline         struct {
			Size        int    `json:"size"`
			FileName    string `json:"file_name"`
			DownloadURL string `json:"download_url"`
		} `json:"offline"`
		VideoTime string `json:"video_time"`
		Share     struct {
			Content string `json:"content"`
			Title   string `json:"title"`
			Poster  string `json:"poster"`
			Cover   string `json:"cover"`
		} `json:"share"`
		ArticleCouldPreview bool   `json:"article_could_preview"`
		ArticlePosterWxlite string `json:"article_poster_wxlite"`
		ArticleFeatures     int    `json:"article_features"`
		CommentCount        int    `json:"comment_count"`
		AudioMd5            string `json:"audio_md5"`
		ChapterSourceID     string `json:"chapter_source_id"`
		AudioTimeArr        struct {
			M string `json:"m"`
			S string `json:"s"`
			H string `json:"h"`
		} `json:"audio_time_arr"`
		HlsVideos         []any  `json:"hls_videos"`
		InPvip            int    `json:"in_pvip"`
		ArticleSharetitle string `json:"article_sharetitle"`
		ArticleCtime      int    `json:"article_ctime"`
		AudioDubber       string `json:"audio_dubber"`
	} `json:"data"`
	Code int `json:"code"`
}
