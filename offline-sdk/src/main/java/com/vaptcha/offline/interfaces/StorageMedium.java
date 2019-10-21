package com.vaptcha.offline.interfaces;

import java.util.function.BiConsumer;

//session存储
public interface StorageMedium {
    String Get(String key);

    String Set(String key, String value);

    String Delete(String key);

    void ForEach(long parallelismThreshold, BiConsumer<String,String> action);
}
