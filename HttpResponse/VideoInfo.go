package HttpResponse

type VideoInfoResponse struct {
	Code    int                   				`json:"code"`
	Message string                				`json:"message"`
	Ttl     int                   				`json:"ttl"`
	Data    VideoInfoResponseData 				`json:"data"`
}

type VideoInfoResponseData struct {
	Aid       int                         		`json:"aid"`
	Videos    int                         		`json:"videos"`
	Tid       int                         		`json:"tid"`
	Tname     string                      		`json:"tname"`
	Copyright int                         		`json:"copyright"`
	Pic       string                      		`json:"pic"`
	Title     string                      		`json:"title"`
	Pubdate   int                         		`json:"pubdate"`
	Ctime     int                         		`json:"ctime"`
	Desc      string                      		`json:"desc"`
	State     int                         		`json:"state"`
	Attribute int                         		`json:"attribute"`
	Duration  int                         		`json:"duration"`
	Rights    VideoInfoResponseDataRights 		`json:"rights"`
	Owner     VideoInfoResponseDataOwner  		`json:"owner"`
	Stat      VideoInfoResponseDataStat   		`json:"stat"`
	Dynamic   string 					  		`json:"dynamic"`
	Cid		  int						  		`json:"cid"`
	Dimension VideoInfoResponseDataDimension 	`json:"dimension"`
	NoCache	  bool								`json:"no_cache"`
}

type VideoInfoResponseDataDimension struct {
	Width  int 									`json:"width"`
	Height int									`json:"height"`
	Rotate int									`json:"rotate"`
}

type VideoInfoResponseDataStat struct {
	Aid      int 								`json:"aid"`
	View     int 								`json:"view"`
	Danmaku  int 								`json:"danmaku"`
	Reply    int 								`json:"reply"`
	Favorite int 								`json:"favorite"`
	Coin     int 								`json:"coin"`
	Share    int 								`json:"share"`
	NowRank  int 								`json:"now_rank"`
	HisRank  int 								`json:"his_rank"`
	Like     int 								`json:"like"`
	Dislike  int 								`json:"dislike"`
}

type VideoInfoResponseDataOwner struct {
	Mid  int    								`json:"mid"`
	Name string 								`json:"name"`
	Face string 								`json:"face"`
}

type VideoInfoResponseDataRights struct {
	Bp            int 							`json:"bp"`
	Elec          int 							`json:"elec"`
	Download      int 							`json:"download"`
	Movie         int 							`json:"movie"`
	Pay           int 							`json:"pay"`
	Hd5           int 							`json:"hd_5"`
	NoReprint     int 							`json:"no_reprint"`
	Autoplay      int 							`json:"autoplay"`
	UgcPay        int 							`json:"ugc_pay"`
	IsCooperation int 							`json:"is_cooperation"`
	UgcPayPreview int 							`json:"ugc_pay_preview"`
}
