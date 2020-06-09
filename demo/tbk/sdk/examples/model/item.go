package model

type Item struct {
	TbkItemInfoGetResponse TbkItemInfoGetResponse `json:"tbk_item_info_get_response"`
}

type NTbkItem struct {
	CatName                    string      `json:"cat_name"`
	NumIid                     int         `json:"num_iid"`
	Title                      string      `json:"title"`
	PictURL                    string      `json:"pict_url"`
	SmallImages                SmallImages `json:"small_images"`
	ReservePrice               string      `json:"reserve_price"`
	ZkFinalPrice               string      `json:"zk_final_price"`
	UserType                   int         `json:"user_type"`
	Provcity                   string      `json:"provcity"`
	ItemURL                    string      `json:"item_url"`
	SellerID                   int         `json:"seller_id"`
	Volume                     int         `json:"volume"`
	Nick                       string      `json:"nick"`
	CatLeafName                string      `json:"cat_leaf_name"`
	IsPrepay                   bool        `json:"is_prepay"`
	ShopDsr                    int         `json:"shop_dsr"`
	Ratesum                    int         `json:"ratesum"`
	IRfdRate                   bool        `json:"i_rfd_rate"`
	HGoodRate                  bool        `json:"h_good_rate"`
	HPayRate30                 bool        `json:"h_pay_rate30"`
	FreeShipment               bool        `json:"free_shipment"`
	MaterialLibType            string      `json:"material_lib_type"`
	PresaleDiscountFeeText     string      `json:"presale_discount_fee_text"`
	PresaleTailEndTime         int64       `json:"presale_tail_end_time"`
	PresaleTailStartTime       int64       `json:"presale_tail_start_time"`
	PresaleEndTime             int64       `json:"presale_end_time"`
	PresaleStartTime           int64       `json:"presale_start_time"`
	PresaleDeposit             string      `json:"presale_deposit"`
	JuPlayEndTime              int64       `json:"ju_play_end_time"`
	JuPlayStartTime            int64       `json:"ju_play_start_time"`
	PlayInfo                   string      `json:"play_info"`
	TmallPlayActivityEndTime   int64       `json:"tmall_play_activity_end_time"`
	TmallPlayActivityStartTime int64       `json:"tmall_play_activity_start_time"`
	JuOnlineStartTime          string      `json:"ju_online_start_time"`
	JuOnlineEndTime            string      `json:"ju_online_end_time"`
	JuPreShowStartTime         string      `json:"ju_pre_show_start_time"`
	JuPreShowEndTime           string      `json:"ju_pre_show_end_time"`
	SalePrice                  string      `json:"sale_price"`
	KuadianPromotionInfo       string      `json:"kuadian_promotion_info"`
}
type Results struct {
	NTbkItem []NTbkItem `json:"n_tbk_item"`
}
type TbkItemInfoGetResponse struct {
	Results Results `json:"results"`
}
