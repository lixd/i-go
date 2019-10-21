package com.vaptcha.offline.utils;

import com.vaptcha.offline.constant.Constant;
import org.springframework.util.DigestUtils;

import java.sql.Timestamp;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.Random;
import java.util.UUID;

public class Common {
    /**
     * 获取4位随机字符串
     *
     * @return
     */
    public static String GetRandomStr() {
        Random r = new Random();
        String[] split = Constant.Char.split("");
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < 4; i++) {
            String s = split[r.nextInt(16)];
            sb.append(s);
        }
        return sb.toString();
    }

    public static String MD5Encode(String str) {
        return DigestUtils.md5DigestAsHex(str.getBytes());
    }

    public static String GetUUID() {
        String uuid = UUID.randomUUID().toString();
        return uuid.replace("-", "");
    }

    public static String StrAppend(String[] s) {
        StringBuilder sb = new StringBuilder();
        for (String value : s) {
            sb.append(value);
        }
        return sb.toString();
    }

    public static long GetTimeStamp() {
        return Timestamp.valueOf(LocalDateTime.of(LocalDate.now(), LocalTime.of(0, 0))).toInstant().
                getEpochSecond();
    }
}
