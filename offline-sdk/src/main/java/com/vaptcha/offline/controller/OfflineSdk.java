package com.vaptcha.offline.controller;

import com.google.gson.Gson;
import com.vaptcha.offline.constant.Constant;
import com.vaptcha.offline.domain.*;
import com.vaptcha.offline.interfaces.StorageMediumImpl;
import com.vaptcha.offline.interfaces.VaptchaOffline;
import com.vaptcha.offline.utils.Common;
import com.vaptcha.offline.utils.HttpClientUtil;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.HttpPost;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ConcurrentHashMap;

@Controller
public class OfflineSdk implements VaptchaOffline {
    private ConcurrentHashMap<String, String> map = new ConcurrentHashMap<>();
    private StorageMediumImpl storageMedium = new StorageMediumImpl(map);

    {
        //移除过期session
        Thread thread = new Thread(new Runnable() {
            @Override
            public void run() {
             while (true){
                 storageMedium.ForEach(1,
                         (key, value) -> {
                             if (value.length() > 10) {
                                 String[] split = value.split(Constant.SpiltChallenge);
                                 if (split.length < 2) {
                                     storageMedium.Delete(value);
                                 } else {
                                     long unix = Long.parseLong(split[0]);
                                     long timeStamp = Common.GetTimeStamp();
                                     long timeSpan = timeStamp - unix;
                                     if (timeSpan > 3 * 60 || timeSpan < -60) {
                                         // 大于3分钟过期
                                         storageMedium.Delete(value);
                                     }
                                 }
                             }
                         });
                 try {
                     Thread.sleep(1000 * 10);
                 } catch (InterruptedException e) {
                     e.printStackTrace();
                 }
                 System.out.println("RemoveSession Running~");
             }
            }
        });
        thread.start();
    }

    @RequestMapping(value = "/downtime")
    @ResponseBody
    @Override
    public Object DownHandler(Offline offline) {
        String callback = offline.getCallback().trim();
        if ("".equals(callback)) {
            return new VerifyResp(Constant.Fail, Constant.Fail, "");
        }
        if (Constant.ActionGet.equals(offline.getOffline_action())) {
            String vid = offline.getVid();
            String challenge = offline.getChallenge();
            ImgResp imgResp = GetImage(challenge, vid);
            String timestamp = String.valueOf(Common.GetTimeStamp());
            String[] challengeArray = new String[]{timestamp, Constant.SpiltChallenge, imgResp.getImgid(), Constant.SpiltChallenge, imgResp.getKey()};
            String challengeValue = Common.StrAppend(challengeArray);
            storageMedium.Set(imgResp.getChallenge(), challengeValue);

            Gson gson = new Gson();
            String res = gson.toJson(imgResp);
            String[] resultArray = new String[]{callback, "(", res, ")"};
            return Common.StrAppend(resultArray);
        } else {
            String path = offline.getV();
            // val=unix#imgId#dtKey
            String val = storageMedium.Get(offline.getChallenge());
            String[] vals = val.split(Constant.SpiltChallenge);
            if (vals.length == 3) {
                String imgId = vals[1];
                VerifyResp verify = verify(imgId, path, vals[2]);
                if (Constant.Success.equals(verify.getCode())) {
                    String uuid = Common.GetUUID();
                    String timestamp = String.valueOf(Common.GetTimeStamp());
//                    String[] resTokenArray = new String[]{uuid, timestamp};
//                    String resToken = Common.StrAppend(resTokenArray);
                    String resToken = uuid + timestamp;

//                    String[] realTokenArray = new String[]{timestamp,Constant.SpiltChallenge, resToken};
//                    String realToken = Common.StrAppend(realTokenArray);
                    String realToken = timestamp + Constant.SpiltChallenge + resToken;
                    //token unix#uuid#unix
                    storageMedium.Set(offline.getChallenge(), realToken);
                    //  resToken 01-challenge-token
                    String[] resultArray = new String[]{"01", Constant.SpiltToken, offline.getChallenge(), Constant.SpiltToken, resToken};
                    String resultToken = Common.StrAppend(resultArray);
                    verify.setToken(resultToken);
                }
                Gson gson = new Gson();
                String result = gson.toJson(verify);
                String[] resultArray = new String[]{offline.getCallback(), "(", result, ")"};
                return Common.StrAppend(resultArray);
            }
        }
        return null;
    }

    @RequestMapping(value = "/verify", method = RequestMethod.POST)
    @ResponseBody
    @Override
    public Object Validate(@RequestBody Validate validate) {
        ValidateResp validateResp = new ValidateResp();

        String token = validate.getToken();
        //  resToken 01-challenge-token
        String[] reqTokens = token.split(Constant.SpiltToken);
        validateResp.setMsg("验证失败");
        if (reqTokens.length == 3) {
            // val token unix#uuid+unix
            String val = storageMedium.Get(reqTokens[1]);
            String[] vals = val.split(Constant.SpiltChallenge);
            // reqTokens[2]=reqToken vals[1]=realToken
            if (2 == vals.length && reqTokens[2].equals(vals[1])) {
                validateResp.setCode(1);
                validateResp.setMsg("验证通过");
            }
        }
        storageMedium.Delete(reqTokens[1]);
        Gson gson = new Gson();
        return gson.toJson(validateResp);
    }

    private ImgResp GetImage(String challenge, String vid) {
        String key = GetKey(vid);
        if ("".equals(key)) {
            return new ImgResp("0104", "", "", "离线key获取失败", "");
        }
        String randomStr = Common.GetRandomStr();
        String str = key + randomStr;
        String imgId = Common.MD5Encode(str);
        if ("".equals(challenge)) {
            String uuid = Common.GetUUID();
            String timestamp = String.valueOf(Common.GetTimeStamp());
            challenge = uuid + timestamp;
        }
        return new ImgResp("0103", imgId, challenge, "", key);
    }

    private String GetKey(String vid) {
        String key = storageMedium.Get(vid);
        if (key != null && !"".equals(key)) {
            return key;
        }
        try {
            List<NameValuePair> parametersBody = new ArrayList();
            HttpResp httpResp = HttpClientUtil.getRequest(Constant.BaseUrl + "config/" + vid, parametersBody);
            String replace = httpResp.getResp().replace("static(", "");
            String replace1 = replace.replace(")", "");
            Gson gson = new Gson();
            GetResp offlineKey = gson.fromJson(replace1, GetResp.class);
            String sOfflineKey = offlineKey.getDtkey();
            String unix = String.valueOf(Common.GetTimeStamp());
            String sb = unix + Constant.SpiltChallenge + sOfflineKey;
            storageMedium.Set(Constant.OfflineKey, sb);
            return sOfflineKey;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return "";
    }

    private VerifyResp verify(String imgId, String path, String offlineKey) {
        String url = Common.MD5Encode(path + imgId);
        String fullUrl = Constant.VerifyUrl + offlineKey + "/" + url;
        try {
            List<NameValuePair> parametersBody = new ArrayList();
            HttpResp httpResp = HttpClientUtil.getRequest(fullUrl, parametersBody);
            if (200 == httpResp.getCode()) {
                return new VerifyResp(Constant.Success, "", "");
            } else {
                return new VerifyResp(Constant.Fail, Constant.Fail, "");
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return new VerifyResp(Constant.Fail, Constant.Fail, "");
    }
}

