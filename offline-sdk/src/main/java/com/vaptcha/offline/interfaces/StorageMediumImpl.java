package com.vaptcha.offline.interfaces;

import java.util.concurrent.ConcurrentHashMap;
import java.util.function.BiConsumer;

public class StorageMediumImpl implements StorageMedium {
    private ConcurrentHashMap<String, String> concurrentHashMap;

    public StorageMediumImpl(ConcurrentHashMap<String, String> concurrentHashMap) {
        this.concurrentHashMap = concurrentHashMap;
    }

    public ConcurrentHashMap<String, String> getConcurrentHashMap() {
        return concurrentHashMap;
    }

    public void setConcurrentHashMap(ConcurrentHashMap<String, String> concurrentHashMap) {
        this.concurrentHashMap = concurrentHashMap;
    }

    @Override
    public String Get(String key) {
        String value = this.concurrentHashMap.get(key);
        if (value != null) {
            return value;
        }
        return "";
    }

    @Override
    public String Set(String key, String value) {
        return this.concurrentHashMap.put(key, value);
    }

    @Override
    public String Delete(String key) {
        return this.concurrentHashMap.remove(key);
    }

    @Override
    public void ForEach(long parallelismThreshold, BiConsumer<String, String> action) {
        this.concurrentHashMap.forEach(parallelismThreshold, action);
    }

}
