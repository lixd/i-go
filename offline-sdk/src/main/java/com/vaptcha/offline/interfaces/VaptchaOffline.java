package com.vaptcha.offline.interfaces;

import com.vaptcha.offline.domain.Offline;
import com.vaptcha.offline.domain.Validate;
import com.vaptcha.offline.domain.VerifyResp;
import org.springframework.ui.Model;

//sdk
public interface VaptchaOffline {
    Object DownHandler(Offline offline);

    Object Validate(Validate validate);
}
