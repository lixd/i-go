package proxyutil

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/net/proxy"
)

/*
主要功能:维护代理IP库、根据代理IP构建http.client
15分钟替换一轮 一天大概100轮
每轮100个IP
每秒1000次推送 平均每个IP每秒10次
*/

// https://h.wandouip.com/get#api 自动生成
const (
	Scheme = "http"
	Host   = "api.wandoudl.com"
	Path   = "api/ip"
	// Key    = "9da3d1b6d847fbd1748060bcef3282c0"
	Key = "7f009c10a2ae6bf173924bdb82452d12"
	// Pack = "220650"
	Pack = "0"
)

// BuildClientWithProxy 根据代理服务器构建http.client
func BuildClientWithProxy(ps string) (*http.Client, error) {
	if ps == "" {
		return &http.Client{}, nil
	}
	targetURL := url.URL{}
	proxyServer, err := targetURL.Parse(ps)
	if err != nil {
		return nil, err
	}
	cli := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,                        // 限制连接数
			Proxy:               http.ProxyURL(proxyServer), // 设置代理服务器
		}}
	return cli, nil
}

// LoadPSWithCache 从缓存中加载IP,不足时才提取
func LoadPSWithCache(num int) ([]string, error) {
	return loadProxyServer(num, true)
}

// LoadPSWithoutCache 直接提取IP
func LoadPSWithoutCache(num int) ([]string, error) {
	return loadProxyServer(num, false)
}

// LoadSocksWithCache 从缓存中加载IP,不足时才提取
func LoadSocksWithCache(num int) ([]string, error) {
	return loadProxyServerSocks(num, true)
}

// LoadSocksWithoutCache 直接提取IP
func LoadSocksWithoutCache(num int) ([]string, error) {
	return loadProxyServerSocks(num, false)
}

// loadProxyServer 提取代理IP
/*
1.如果缓存中有则从缓存中获取
2.否则进行提取
3.isNew 则直接进行提取
*/
func loadProxyServer(num int, useCache bool) ([]string, error) {
	var (
		key  = KeyProxyServer
		list = make([]string, 0, num)
	)

	// 	1.如果缓存中有则从缓存中获取
	l := Len(key)
	for useCache && l > 0 && l > int64(len(list)) && len(list) < num {
		fmt.Println("校验缓存代理IP可用性 ing")
		ps := Pop(key)
		if ps == "" {
			break
		}
		isValidProxy := IsValidProxy(ps)
		if isValidProxy {
			list = append(list, ps)
		}
	}
	Add(list, key) // 把没过期的再次写入队列
	// 	2.否则进行提取
	proxyIP, err := loadProxyIP(num - len(list))
	if err != nil {
		return list, err
	}
	for _, v := range proxyIP.Data {
		ps := fmt.Sprintf("http://%s:%v", v.IP, v.Port)
		fmt.Printf("提取IP:%v \n", ps)
		Add([]string{ps}, key)
		list = append(list, ps)
	}
	return list, nil
}

// loadProxyServer 提取代理IP
/*
1.如果缓存中有则从缓存中获取
2.否则进行提取
3.isNew 则直接进行提取
*/
func loadProxyServerSocks(num int, useCache bool) ([]string, error) {
	var (
		key  = KeyProxyServerSocks
		list = make([]string, 0, num)
	)
	fmt.Printf("loadProxyServerSocks num:%v useCache:%v\n", num, useCache)
	// 	1.如果缓存中有则从缓存中获取
	l := Len(key)
	for useCache && l > 0 && int64(len(list)) < l && len(list) < num {
		fmt.Println("校验缓存代理IP可用性 ing")
		ps := Pop(key)
		if ps == "" { // 取不到时直接退出循环
			break
		}
		isValidProxy := IsValidProxySocks(ps)
		if isValidProxy {
			list = append(list, ps)
		} else {
			fmt.Println("代理无效:", ps)
		}
	}
	Add(list, key) // 把没过期的再次写入队列
	// 	2.否则进行提取
	proxyIP, err := loadSocks(num - len(list))
	if err != nil {
		return list, err
	}
	for _, v := range proxyIP.Data {
		ps := fmt.Sprintf("%s:%v", v.IP, v.Port)
		fmt.Printf("提取IP:%v \n", ps)
		Add([]string{ps}, key)
		// // 顺便写入MongoDB存储一下代理IP
		// err = dao.IP.Insert(v.IP)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		list = append(list, ps)
	}
	return list, nil
}

// loadSocks 提取代理IP
func loadSocks(num int) (ProxyIP, error) {
	var item ProxyIP

	if num == 0 {
		return item, nil
	}
	link := payloadSocks(num)
	req, err := BuildRequest(link)
	if err != nil {
		return item, err
	}
	client, err := BuildClientWithProxy("")
	if err != nil {
		return item, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(buf, &item)
	if err != nil {
		return item, err
	}
	if item.Code == 10048 || item.Msg == "没有可用套餐" {
		fmt.Printf("提取代理:%+v\n", item)
		return item, errors.New("没有可用套餐")
	}
	return item, nil
}

// loadProxyIP 提取代理IP
func loadProxyIP(num int) (ProxyIP, error) {
	var item ProxyIP

	if num == 0 {
		return item, nil
	}
	link := payload(num)
	fmt.Println("提取IPLink:", link)
	req, err := BuildRequest(link)
	if err != nil {
		return item, err
	}
	client, err := BuildClientWithProxy("")
	if err != nil {
		return item, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(buf, &item)
	if err != nil {
		return item, err
	}
	if item.Code == 10048 || item.Msg == "没有可用套餐" {
		fmt.Printf("提取代理:%+v\n", item)
		return item, errors.New("没有可用套餐")
	}
	return item, nil
}

func payload(num int) string {
	v := url.Values{}
	v.Set("app_key", Key)
	// v.Set("pack", "0")              // 所有可用套餐
	v.Set("pack", Pack)             // 包量套餐
	v.Set("num", strconv.Itoa(num)) // 提取IP数
	v.Set("xy", "1")                // http
	v.Set("type", "2")
	v.Set("lb", `\r\n`)
	v.Set("mr", "1")
	v.Set("area_id=", "")
	l := url.URL{
		Scheme:     Scheme,
		Host:       Host,
		Path:       Path,
		ForceQuery: false,
		RawQuery:   v.Encode(),
		Fragment:   "",
	}
	return l.String()
}
func payloadSocks(num int) string {
	v := url.Values{}
	v.Set("app_key", Key)
	v.Set("pack", "0")
	v.Set("num", strconv.Itoa(num)) // 提取IP数
	v.Set("xy", "3")                // socks
	v.Set("type", "2")
	v.Set("lb", `\r\n`)
	v.Set("mr", "1")
	v.Set("area_id=", "")
	l := url.URL{
		Scheme:     Scheme,
		Host:       Host,
		Path:       Path,
		ForceQuery: false,
		RawQuery:   v.Encode(),
		Fragment:   "",
	}
	return l.String()
}

// ----------------Proxy Test------------------
func socket5Proxy(targetURLString, httpProxyServer string) {
	dialer, err := proxy.SOCKS5("tcp", httpProxyServer, nil, proxy.Direct)
	if err != nil {
		fmt.Println("can't connect to the proxy:", err)
		return
	}

	// setup a http client
	httpTransport := &http.Transport{}
	client := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial

	// 访问地址
	rqt, err := http.NewRequest("GET", targetURLString, nil)
	if err != nil {
		println("请求网站失败")
		return
	}
	// 处理返回结果
	response, err := client.Do(rqt)
	if err != nil {
		fmt.Println("Do", err)
		return
	}
	defer response.Body.Close()
	// 读取内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	// 输入响应结果
	fmt.Println("http 请求结果:", string(body))
}

func normalProxy(ps string) {
	cli, err := BuildClientWithProxy(ps)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(http.MethodGet, "http://myip.ipip.net", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*(*string)(unsafe.Pointer(&bytes)))
}

// IsValidProxy 判断当前代理是否有效
func IsValidProxy(ps string) bool {
	cli, err := BuildClientWithProxy(ps)
	if err != nil {
		return false
	}
	ctx := context.TODO()
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	req, err := http.NewRequestWithContext(timeout, http.MethodGet, "http://myip.ipip.net", nil)
	if err != nil {
		return false
	}
	resp, err := cli.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

const u = `https://qydc.vaptcha.net/validate?v=_4m_4f_32_2z_2l_2c_hm_he_gx_g3_ge_fy_fn_f2_fc_ex_eu_ex_fh_fx_h2_3g_43_lt_m4_mz_n3_nx_tc_t2_tn_tz_ue_uf_ug_uh_uh_uh_uh_uf_ud_tz_tw_t3_td_ny_nn_nf_na_m4_mb_lm_lg_lb_4t_43_4e_3w_3l_3g_2x_2m_2g_ea_dy_dx_dx_dx_dx_dx_dx_dy_dz_ea_ea_eb_eb_ed_ef_el_ey_fe_fh_ft_ga_gc_gf_g2_g4_g4_gl_gm_gn_gu_gy_hc_hh_hm_hx_2c_2h_2m_2y_3e_33_3t_4a_4f_43_4m_4x_la_le_lh_l4_lm_ln_lt_lw_ly_ma_mb_md_mf_mh_m4_ct_dc_dz_eg_eu_fc_f4_fy_gg_gt_hc_h4_hy_2h_34_4g_ly_ny_uc_u4_wt_x4_xy_yg_yt_zc_zl_zybagbatbbcbctbfcbhcb2gb4gbl4bmgbncbtgbu4bwgbxcbygbzcbz4bzycatcbccbycctcd4cegcfccfycguchlc2tc4tclycn4ctgcuy&vi=5b4d9c33a485e50410192331&k=07505TTUV2574d1b&dt=1399&ch=290&cw=460&origin_url=https%3A%2F%2Fwww.vaptcha.com%2F&lo=true&x=1627522965oeoB6Kbf716&callback=VaptchaJsonp1627523873169&d=9D5F1ADC8A4A576E`

// IsValidProxy 判断当前代理是否有效
func Validate(ps string) bool {
	cli, err := BuildClientWithProxy(ps)
	if err != nil {
		return false
	}
	ctx := context.TODO()
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	req, err := http.NewRequestWithContext(timeout, http.MethodGet, u, nil)
	if err != nil {
		return false
	}
	resp, err := cli.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func IsValidProxySocks(ps string) bool {
	// ps = strings.TrimPrefix(ps, "http://")
	// fmt.Println("IsValidProxySocks ps:", ps)
	// ps = "183.69.229.173:8081"
	dialer, err := proxy.SOCKS5("tcp", ps, nil, proxy.Direct)
	if err != nil {
		return false
	}

	// setup a http client
	httpTransport := &http.Transport{}
	client := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial

	// 请求百度 只有有响应则说明 IP 有效
	req, err := http.NewRequest("GET", "http://myip.ipip.net", nil)
	// req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		return false
	}
	// 处理返回结果
	response, err := client.Do(req)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	_, _ = io.Copy(ioutil.Discard, response.Body)
	return true
}

// FormatIP 从完整代理链接中提取IP params:http://ip:port return:ip
func FormatIP(ps string) string {
	ps = strings.ReplaceAll(ps, "http://", "")
	split := strings.Split(ps, ":")
	if len(split) != 0 {
		return split[0]
	}
	return ""
}
