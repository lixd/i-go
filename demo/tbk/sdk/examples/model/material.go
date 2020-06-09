package model

type Material struct {
	TbkDgMaterialOptionalResponse TbkDgMaterialOptionalResponse `json:"tbk_dg_material_optional_response"`
}

type MapData struct {
	CouponStartTime        string      `json:"coupon_start_time"`
	CouponEndTime          string      `json:"coupon_end_time"`
	InfoDxjh               string      `json:"info_dxjh"`
	TkTotalSales           string      `json:"tk_total_sales"`
	TkTotalCommi           string      `json:"tk_total_commi"`
	CouponID               string      `json:"coupon_id"`
	NumIid                 int64       `json:"num_iid"`
	Title                  string      `json:"title"`
	PictURL                string      `json:"pict_url"`
	SmallImages            SmallImages `json:"small_images"`
	ReservePrice           string      `json:"reserve_price"`
	ZkFinalPrice           string      `json:"zk_final_price"`
	UserType               int         `json:"user_type"`
	Provcity               string      `json:"provcity"`
	ItemURL                string      `json:"item_url"`
	IncludeMkt             string      `json:"include_mkt"`
	IncludeDxjh            string      `json:"include_dxjh"`
	CommissionRate         string      `json:"commission_rate"`
	Volume                 int         `json:"volume"`
	SellerID               int         `json:"seller_id"`
	CouponTotalCount       int         `json:"coupon_total_count"`
	CouponRemainCount      int         `json:"coupon_remain_count"`
	CouponInfo             string      `json:"coupon_info"`
	CommissionType         string      `json:"commission_type"`
	ShopTitle              string      `json:"shop_title"`
	ShopDsr                int         `json:"shop_dsr"`
	CouponShareURL         string      `json:"coupon_share_url"`
	URL                    string      `json:"url"`
	LevelOneCategoryName   string      `json:"level_one_category_name"`
	LevelOneCategoryID     int         `json:"level_one_category_id"`
	CategoryName           string      `json:"category_name"`
	CategoryID             int         `json:"category_id"`
	ShortTitle             string      `json:"short_title"`
	WhiteImage             string      `json:"white_image"`
	Oetime                 string      `json:"oetime"`
	Ostime                 string      `json:"ostime"`
	JddNum                 int         `json:"jdd_num"`
	JddPrice               string      `json:"jdd_price"`
	UvSumPreSale           int         `json:"uv_sum_pre_sale"`
	XID                    string      `json:"x_id"`
	CouponStartFee         string      `json:"coupon_start_fee"`
	CouponAmount           string      `json:"coupon_amount"`
	ItemDescription        string      `json:"item_description"`
	Nick                   string      `json:"nick"`
	OrigPrice              string      `json:"orig_price"`
	TotalStock             int         `json:"total_stock"`
	SellNum                int         `json:"sell_num"`
	Stock                  int         `json:"stock"`
	TmallPlayActivityInfo  string      `json:"tmall_play_activity_info"`
	ItemID                 int64       `json:"item_id"`
	RealPostFee            string      `json:"real_post_fee"`
	LockRate               string      `json:"lock_rate"`
	LockRateEndTime        int64       `json:"lock_rate_end_time"`
	LockRateStartTime      int64       `json:"lock_rate_start_time"`
	PresaleDiscountFeeText string      `json:"presale_discount_fee_text"`
	PresaleTailEndTime     int64       `json:"presale_tail_end_time"`
	PresaleTailStartTime   int64       `json:"presale_tail_start_time"`
	PresaleEndTime         int64       `json:"presale_end_time"`
	PresaleStartTime       int64       `json:"presale_start_time"`
	PresaleDeposit         string      `json:"presale_deposit"`
	YsylTljSendTime        string      `json:"ysyl_tlj_send_time"`
	YsylCommissionRate     string      `json:"ysyl_commission_rate"`
	YsylTljFace            string      `json:"ysyl_tlj_face"`
	YsylClickURL           string      `json:"ysyl_click_url"`
	YsylTljUseEndTime      string      `json:"ysyl_tlj_use_end_time"`
	YsylTljUseStartTime    string      `json:"ysyl_tlj_use_start_time"`
	SaleBeginTime          string      `json:"sale_begin_time"`
	SaleEndTime            string      `json:"sale_end_time"`
	Distance               string      `json:"distance"`
	UsableShopID           string      `json:"usable_shop_id"`
	UsableShopName         string      `json:"usable_shop_name"`
	SalePrice              string      `json:"sale_price"`
	KuadianPromotionInfo   string      `json:"kuadian_promotion_info"`
}
type ResultList struct {
	MapData []MapData `json:"map_data"`
}
type TbkDgMaterialOptionalResponse struct {
	TotalResults int        `json:"total_results"`
	ResultList   ResultList `json:"result_list"`
}
