$(document).ready(function(){


    //跳转
    let toLogin = $('.to-login')
    let setNewCount = $('#newCount')
    let forgotKey = $('.forgot')
    // let nextStep = $('.next-step')
    // 密码强度
    let mm = $('#mm');
    let resetmm = $('#resetmm');
    let oLevel = $('#level');
    let oLevel1 = $('#level1');
    let canBeSubmit = false;//是否能够提交注册
    let canSubmitReset = false;
    let canBeLogin = false;
    //手机格式验证码
    let vaptchaContent = $('#vaptcha-content');
    let getCode = $('#getCode');
    let getResetCode = $('#getResetCode');
    let btnLogin = $('#btn_login');
    let btnRegist = $('#btn_regist');
    let btnNext = $('#btn_next');
    let btnComfirm = $('#btn-comfirm');
    let resetTip = $('#resetTip')
    let time = 120;
    let timer1,timer2, webToken,codeFlag = false;
    // const baseUrl = 'http://easymock.wlinno.com/mock/5cd38a16d67ed2719c2a78b7/api/';
    const baseUrl = 'https://management.vaptcha.com/api/v2/';
    // const baseUrl = 'http://192.168.0.142:5559/api/v2/';
    // const baseUrl = 'http://localhost:8080/api/v2/';

    let code = $('#code')

    let serverClose = $('#serverClose');
    let agreementLink = $('.agreement-link')
    let userModal = $('.user-modal')
    let agreenBtn = $('#agreenBtn');

    let inputEle = $('.input-style')
    let times1  = $('#times1')
    let isAgreement = true
    let agreementCheckbox = $('#agreementCheckbox')
    let keepLogin = $('#keepLogin')
    let remembermm  = '1'


    //请求是否自动登录
    $.ajax({
      type: "get",
      url: baseUrl + "account/loginstatus",
      // data: data,
      dataType: "json",
      success: function (res) {
        if(res.code !== 401){
          console.log(res.msg,res.data)
          //跳转到主页面
          window.location.href = '/manage'
        }else{
          //跳转到登录页面
          // window.location.href = '/login.html'
        }
      }
    })

    /**
     * 密码强度判断
     * @param {String} id Dom的id
     * @param {String} changeLevel 强度条的id
     * @param {String} canSubmit 标识是否能够提交
     * @param {*} e event
     */
    let strength = (id, changeLevel, canSubmit, e) => {
        e = e || window.e
        var strongRegex = new RegExp("^(?=.{8,})(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*\\W).*$", "g");
        var mediumRegex = new RegExp("^(?=.{7,})(((?=.*[A-Z])(?=.*[a-z]))|((?=.*[A-Z])(?=.*[0-9]))|((?=.*[a-z])(?=.*[0-9]))).*$", "g");
        var enoughRegex = new RegExp("(?=.{6,}).*", "g");
        var allWordsRegex = new RegExp('[a-zA-Z]');
        // mm.removeClass("error");
        if (id.val().length >= 6) {
            // step3.removeClass('error');
            // oErrTip.html('');
            if (strongRegex.test(id.val())) {
                //大写、小写、数字、特殊 包含4种 而且至少8位
                changeLevel.removeClass('pw-medium');
                changeLevel.removeClass('pw-weak');
                changeLevel.addClass('pw-strong');
                canSubmit = true;
            } else if (mediumRegex.test(id.val()) || allWordsRegex.test(id.val())) {
                // 大写、小写、数字、特殊 包含2种或者3种、而且至少7位,或者全为字母 至少大于6位。
                // this.passWordLevelEvent.emit({ strong: false, common: true, weak: false });
                // this.resultEvent.emit(true);
                changeLevel.removeClass('pw-strong');
                changeLevel.removeClass('pw-weak');
                changeLevel.addClass('pw-medium');
                canSubmit = true;
            } else {
                // this.passWordLevelEvent.emit({ strong: false, common: false, weak: true });
                // this.resultEvent.emit(false);
                changeLevel.removeClass('pw-strong');
                changeLevel.removeClass('pw-medium');
                changeLevel.addClass('pw-weak');
                canSubmit = false;
            }
        } else if (id.val().length > 0 && id.val().length < 6) {
        changeLevel.removeClass('pw-strong');
        changeLevel.removeClass('pw-medium');
        changeLevel.addClass('pw-weak');
        changeLevel = false;
            // step3.removeClass('error');
            // oErrTip.html('');
        } else {
        changeLevel.removeClass('pw-strong');
        changeLevel.removeClass('pw-medium');
        changeLevel.removeClass('pw-weak');
            canSubmit = false;
        }
        if (e.keyCode == 13) {
        // 提交
           if(id==mm){
            register()
           }else if(id==resetmm){
            comfirm()
           }
        }
    }

    /* 验证码格式判断 */
    let validateCodeFormat = (value) => {
        // let  reg = /^[A-Za-z0-9]+$/;
        var reg = /^\d{6}$/g;
        if (value.length = 6 && reg.test(value)) {
            return true;
        }
        else {
            return false;
        }
    }
    /* 验证码格式错误的样式改变 */
   
    let erroCodeCss = (id, submit) => {
        let isRightTCode = validateCodeFormat(trim($('#' + id).val()));
        if(!isRightTCode) {
          submit = false;
          $('#'+ id).addClass('error');
        } else {
          submit = true;
          $('#' + id).removeClass('error');
      }
    }

    /*  手机格式判断 */

    // 手机格式
    let validatePhoneFormat = (value) => {
        // let  reg = /^[A-Za-z0-9]+$/;
        var reg = /^1\d{10}$/g;
        if (value.length >= 6 && value.length <= 11 && reg.test(value)) {
            return true;
        }
        else {
            return false;
        }
    }

    /**
     * 删除左右两端的空格 
     * @param {String} str 字符串
     */
    let trim = (str) => {
        return str.replace(/(^\s*)|(\s*$)/g, "");
    }

    /**
     * 手机格式错误的样式改变
     * @param {String} id id
     * @param {*} submit 
     */
    let erroTelCss = (id, submit) => {
        let isRightTel = validatePhoneFormat(trim($('#' + id).val()));
        if(!isRightTel) {
          submit = false;
          $('#'+ id).addClass('error');
        } else {
          submit = true;
          $('#' + id).removeClass('error');
         if(id=='s1tel'){
           getResetCode.addClass('active')
         }else if(id='rtel'){
           getCode.addClass('active')
        }
      }
    }

    /**
     * 空值或长度错误
     * @param {String} id Dom的id
     * @param {String} submit 标识是否能够提交
     * @param {Number} minLength  允许的最小长度
     * @param {Number} maxLength 允许的最大长度
     */
    let emptyOrLengthError = (id,submit, minLength, maxLength) => {
        if(minLength && maxLength){
        if(trim($('#' + id).val()) == '' || $('#' + id).val().length < minLength || $('#' + id).val().length > maxLength) {
            submit = false;
            $('#' + id).addClass('error');
        } else {
            submit = true;
            $('#' + id).removeClass('error');
        }
        } else {
        if(trim($('#' + id).val()) == '') {
            submit = false;
            $('#' + id).addClass('error');
        } else {
            submit = true;
            $('#' + id).removeClass('error');
        }
        }
    }

    /**
     * 登录
     * @param {String} token 人机验证token
     */
    let login = token => {
        let data = {
          phone: $('#ltel').val(),
          password: $('#lpwd').val(),
          validatetoken: token,
          remember:remembermm
        }
        $.ajax({
        type: "post",
        url: baseUrl + "account/login",
        data: JSON.stringify(data),
        dataType: "json",
        contentType: 'application/json',
        success: function (res) {
            if(res.code === 200) {
              $('#ltel, #lpwd').val('')
              window.location.href = '/manage'
            } else {
              $('#ltel,#lpwd').addClass('error');
            }
        }
        });
    }
  
  /**
   * 发送短信
   * @param {String} type register:注册 forgetpassword:忘记密码
   * @param {String} token 人机验证token
   */
  let sendSms = (type,token) => {
    let data = {
      phone:type =='register'? $('#rtel').val():$('#s1tel').val(),
      type: type,
      validatetoken: token
    }
    $.ajax({
      type: "get",
      url: baseUrl + "sms/send",
      data:  data,
      dataType: "json",
      // contentType: 'application/json',
      success: function (res) {
        if(res.code === 200) {
          time = res.data.time
          if(type === 'register'){
            getCode.css('display', 'none');
            $('#times1').css('display', 'block');
            $('#registTime').html(time);
            timer1 = setInterval(() => {
              time --;
              if(time == 0){ 
                clearInterval(timer1);
                time = 120
                getCode.css('display', 'block');
                $('#times1').css('display', 'none');
              }
              $('#registTime').html(time);
            },1000)
          } else {
            getResetCode.css('display', 'none');
            $('#times2').css('display', 'block');
            $('#loginResetTime').html(time);
            timer2 = setInterval(() => {
              time --;
              if(time == 0){ 
                clearInterval(timer2);
                time = 120
                getResetCode.css('display', 'block');
                $('#times2').css('display', 'none');
              }
              $('#loginResetTime').html(time);
            },1000)
          }
        } else {
          if(type === 'register'){
            $('#ltel').addClass('error');
          }else{
            $('#s1tel').addClass('error');
          }
        }
      }
    });
  }
  /**
   * 手机是否存在
   * @param {Function} cb 
   */
  let singlePhone = (cb) => {
    let data = {
      phone: $('#rtel').val(),
    }
    $.ajax({
      type: "get",
      url: baseUrl + "user/phone/presence",
      data:  data,
      dataType: "json",
      // contentType: 'application/json',
      success: function (res) {
        if(!res.data) {
          cb()
        } else {
          $('#rtel').addClass('error');
        }
      }
    });
  }
  
  /**
   * vaptcha 验证模块
   * @param {Number} type 0:登录验证 1:发送短信
   * @param {String} val register: 注册   forgetpassword: 忘记密码
   */
  let vaptcha = (type,val) => {
    if(codeFlag) return
    codeFlag = true
    var options = {
      vid: '59b252ed57f5a21114866a5d',
      scene: '01',
      // https: true,
      type: 'invisible', 
    }

  window.vaptcha(options).then(function (vaptcha_obj) {
      vaptcha_obj.listen('pass', function() {
        codeFlag = false;
        webToken = vaptcha_obj.getToken();
        if(type === 0) {
          login(webToken)
        } else if(type === 1) {
          let tel = val == 'register'? $('#rtel').val()  : $('#s1tel').val()
          // if(val=='register')
          // {
            let data = {
              phone: tel,
            }
            $.ajax({
              type: "get",
              url: baseUrl + "user/phone/presence",
              data:  data,
              dataType: "json",
              // contentType: 'application/json',
              success: function (res) {
                if( val == 'register'){
                  if(!res.data) {
                    sendSms(val, webToken)
                  } else {
                    // $(selector).css(propertyName, value);
                    $('#tel-exit').css('display', 'inline-block')
                    $('#rtel').addClass('error');
                  }
                }
                else {
                  if(res.data) {
                    sendSms(val, webToken)
                  } else {
                    $('#s1tel').addClass('error')
                  }
                } 
              }
            });
          // }
          // else sendSms(val, webToken)
        }
        vaptcha_obj.destroy();
      })
      vaptcha_obj.validate();
    })
  }
  
  /**
   * 判断是否能够提交
   * @param {Array} arr 需要验证的id数组
   */
  let isCanSubmit = (arr) => {
    let flag = true
    arr.forEach(element => {
      if($('#' + element).hasClass('error')){
        flag = false
        return;
      } else if($('#' + element).val() == '') {
        flag = false
        return;
      }
    });
    return flag;
  }

    /* 
      图标颜色转化,字体图标添加一个fa-active的类
    */
    inputEle.on('click',function(){
        inputEle.prev().children(".fa").removeClass('fa-active')
        $(this).prev().children(".fa").addClass('fa-active')
      })
    /**
     * 判断是否记住密码
     */
    function isRemenber ( ) { 
      let isChecked = keepLogin.attr('checked')
      if(!isChecked) {
        remembermm = '0'
      }
     }
    
    /* 控制跳转路径 */
    // 显示登录
    toLogin.on('click',function(e){
        e.preventDefault();
        window.location.href = '/login.html'
    })

    //返回注册界面
    setNewCount.on('click',function(e){
        e.preventDefault();
        window.location.href ="/register.html"
    })
    //重置密码
    forgotKey.on('click', function (e) { 
        e.preventDefault();
        window.location.href = '/reset.html'
    });


    // 密码强度 注册
    mm.on('keyup', function (e) { 
        e.preventDefault();
        strength(mm, oLevel, canBeSubmit, e);
        emptyOrLengthError('mm', canBeSubmit, 6, 24)
    });

    mm.on('blur', function (e) { 
        e.preventDefault();
        emptyOrLengthError('mm', canBeSubmit, 6, 24)
    });

    // 密码强度 重置
    resetmm.on('keyup', function (e) {
        e.preventDefault();
        strength(resetmm, oLevel1, canSubmitReset,e);
        emptyOrLengthError('resetmm', canSubmitReset, 6, 24)
    })

    resetmm.on('blur', function (e) {
        e.preventDefault();
        emptyOrLengthError('resetmm', canSubmitReset, 6, 24)
    })

    // 第一步  手机格式验证
    $('#s1tel').on('blur', function (e) { 
        e.preventDefault();
        erroTelCss('s1tel', canSubmitReset);
    });

    $('#s1tel').on('keyup', function (e) { 
        e.preventDefault();
        erroTelCss('s1tel', canSubmitReset);
    });

    // 登陆 手机格式验证
    $('#ltel').on('blur', function (e) { 
        e.preventDefault();
        erroTelCss('ltel', canBeLogin);
    });

    $('#ltel').on('keyup', function (e) {
        e.preventDefault();
        erroTelCss('ltel', canBeLogin)
        $('#lpwd').removeClass('error')
    })

    // 注册 手机格式验证
    $('#rtel').on('blur', function (e) { 
        e.preventDefault();
        $('#tel-exit').css('display', 'none')
        erroTelCss('rtel', canBeSubmit);
    });

    $('#rtel').on('keyup', function (e) {
        e.preventDefault();
        erroTelCss('rtel', canBeSubmit);
    })
      // 登陆 密码格式验证
    $('#lpwd').on('blur', function (e) { 
        e.preventDefault();
        emptyOrLengthError('lpwd',canBeLogin);
    });

    $('#lpwd').on('keyup', function (e) {
        e.preventDefault()
        emptyOrLengthError('lpwd',canBeLogin)
        //检测手机号
        erroTelCss('ltel', canBeLogin)
        let flag = isCanSubmit(['lpwd','ltel'])
        if(e.keyCode == 13 && flag) {
            isRemenber()
            vaptcha(0)
        }
    })

    //注册 短信验证码格式

    $('#code').on('blur', function (e) { 
        e.preventDefault();
        erroCodeCss('code',canBeSubmit)
    });

    $('#code').on('keyup', function (e) {
        e.preventDefault();
        erroCodeCss('code',canBeSubmit)
        //检测手机号
        erroTelCss('rtel', canBeSubmit)
    })

    //重置 短信验证码格式
    $('#recode').on('blur', function (e) { 
        e.preventDefault();
        erroCodeCss('recode',canBeSubmit)
    });

    $('#recode').on('keyup', function (e) {
        e.preventDefault();
        erroCodeCss('recode',canBeSubmit)
        //检测手机号
        erroTelCss('s1tel', canBeSubmit)
    })
    

    // 验证码
    getResetCode.on('click', function (e) {
      e.preventDefault();
      let isRightTel = validatePhoneFormat(trim($('#s1tel').val()));
      if(isRightTel) {
      // 请求验证码
      vaptcha(1, 'forgetpassword');

      } else {
      $('#s1tel').addClass('error');
      }
    })

    getCode.on('click', function (e) {
      e.preventDefault();
      let isRightTel = validatePhoneFormat(trim($('#rtel').val()));
      if(isRightTel) {
        // 请求验证码
        vaptcha(1, 'register');
      } else {
        $('#rtel').addClass('error');
      }
    })


    //用户协议显示
    agreementLink.on('click',function(e){
      e.preventDefault();
      userModal.addClass('show')
      })
    serverClose.on('click', function (e) { 
      e.preventDefault();
      userModal.removeClass('show')
      // vaptchaContent.css('display','block')
    });

    agreenBtn.on('click', function (e) { 
      e.preventDefault();
      userModal.removeClass('show')

    });

    btnLogin.on('click', function (e) {
      e.preventDefault();
      isRemenber()
      let flag = isCanSubmit(['lpwd','ltel'])
      if(flag){
        vaptcha(0)
      }

    })



    //注册提交函数
    function register () {
      let flag = isCanSubmit(['rtel','code','mm'])
      let flag2 = agreementCheckbox.prop("checked")
      if(flag2){ 
        isAgreement = true
      } else{
        alert('请先VAPTCHA用户服务协议')
      }
      if(flag && isAgreement) {
        clearInterval(timer1);
        // oLevel.removeClass('pw-weak pw-medium pw-strong')
        times1.css('display','none')
        getCode.css('display','block')
        getCode.removeClass('active')
        let data = {
          phone: $('#rtel').val(),
          code: $('#code').val(),
          password: mm.val()
        }
        $.ajax({
          type: "post",
          url:  baseUrl + 'account/register',
          data: JSON.stringify(data),
          dataType: "json",
          contentType: 'application/json',
          success: function (res) {
            if(res.code === 200){
              // console.log('成功')
              $('#rtel').val('')
              $('#code').val('')
              mm.val('')
              //跳转到登录页面
              window.location.href = '/login.html'
            }else{
              $('#rtel,#code').addClass('error')
            }
          }
        });
      }
    }

  //注册新账户接口
  btnRegist.on('click', function (e) {
    e.preventDefault();
    register()
  })


  function canGoStep2() { 

    let flag = isCanSubmit(['s1tel','recode'])
    if(flag) {
      oLevel1.removeClass('pw-weak pw-medium pw-strong')
      $('#times2').css('display','none')
      // getResetCode.css('display','block')
      // getResetCode.removeClass('active')
      let data = {
         phone: $('#s1tel').val(),
         code: $('#recode').val(),
      }
      $.ajax({
        type: "post",
        url:  baseUrl + 'account/password/forget',
        data: JSON.stringify(data),
        contentType: 'application/json',
        dataType: "json",
        success: function (res) {
          if(res.code === 200){
            $('#s1tel').val('')
            $('#recode').val('')
            $('#step1').addClass('hide')
            $('#step2').removeClass('hide')
            return token = res.data.token;
          }else{
            $('#s1tel,#recode').addClass('error')
          }
        }
      });
    }
   }

  //忘记密码接口
  btnNext.on('click', function (e) {

    e.preventDefault();
    canGoStep2()
  })
  
  //忘记密码回车跳转
  $('#recode').on('keyup', function (e) {
    if(e.keyCode == 13) {
      canGoStep2()
    }
  })

  //重置密码确认

  function comfirm() {
    let flag = isCanSubmit(['resetmm'])
    if(flag) {
      let data = {
         token: token,
         password: $('#resetmm').val(),
      }
      $.ajax({
        type: "post",
        url:  baseUrl + 'account/password/reset',
        data: JSON.stringify(data),
        dataType: "json",
        contentType: 'application/json',
        success: function (res) {
          if(res.code === 200){
            $('#resetTip').addClass('show')
            resetmm.val('')
            setTimeout(function() {
              window.location.href = '/login.html'
              $('#resetTip').addClass('hide')
            },2000)
          }else{
            // resetmm.addClass('error')
          }
        }
      });
    }
  }

  btnComfirm.on('click', function (e) {
    e.preventDefault();
    comfirm()
  })


});

