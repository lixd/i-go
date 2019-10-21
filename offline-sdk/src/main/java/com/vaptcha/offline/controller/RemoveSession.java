package com.vaptcha.offline.controller;

import com.vaptcha.offline.constant.Constant;
import com.vaptcha.offline.interfaces.StorageMediumImpl;
import com.vaptcha.offline.utils.Common;

public class RemoveSession implements Runnable {
    private StorageMediumImpl storageMedium;

    public RemoveSession(StorageMediumImpl storageMedium) {
        this.storageMedium = storageMedium;
    }

    @Override
    public void run() {
        while (true) {
            this.storageMedium.ForEach(1,
                    (key, value) -> {
                        if (value.length() > 10) {
                            String[] split = value.split(Constant.SpiltChallenge);
                            if (split.length < 2) {
                                this.storageMedium.Delete(value);
                            } else {
                                long unix = Long.parseLong(split[0]);
                                long timeStamp = Common.GetTimeStamp();
                                long timeSpan = timeStamp - unix;
                                if (timeSpan > 3 * 60 || timeSpan < -60) {
                                    // 大于3分钟过期
                                    this.storageMedium.Delete(value);
                                }
                            }
                        }
                    });
            try {
                Thread.sleep(10000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println("RemoveSession Running~");
        }
    }
}
