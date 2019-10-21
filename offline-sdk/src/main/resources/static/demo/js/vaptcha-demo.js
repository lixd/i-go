(function() {
  var vids = {
    embed: '5b4d9c33a485e50410192331',
    float: '5b4d9dfea485e5041019253f',
    invisible: '5b4d9c76a485e504101923df'
  }
  // vaptcha demo
  VD = function(element, type, cb, lang) {
      var hashtype = location.hash.substr(1)
      var el = $('[type=' + hashtype + ']')
      var it = this;
      this.vaptcha = null;
      this.token = null;
      this.reset = function() {};
      this.initVaptcha = function(vidType) {
          if (type === 'invisible') vidType = $('.type-switch .active').attr('type') || 'simple'
          if (it.vaptcha) {
              if (!vidType) return it.vaptcha.reset();
              else it.vaptcha.destroy();
          }
          el && !vidType && (vidType = hashtype)
          var options = {
              vid: vids[type],
              container: element,
              https: true,
              type: type,
              // style:type==='float'?'light':'dark',
              ai: type === 'popup' ? false : true,
              checkingAnimation: 'display',
              lang: lang || 'zh-CN',
              outage: 'http://localhost:8080/downtime',
              // color: type==='float'?'rgb(51, 51, 51)':'',
          }
          window.vaptcha(options).then(function(vaptcha_obj) {
              it.vaptcha = vaptcha_obj;
              vaptcha_obj.listen('pass', function(token, challenge) {
                  it.challenge = challenge;
                  it.token = token;
                  cb && cb()
              })
              type == 'invisible' ? it.vaptcha.validate() : it.vaptcha.render();
          })
      }
      var current = 0;
      var previousTime;
      var vpRefreshI = document.getElementById('vpRefreshI');
      $('#reset').click(function() {
          it.reset();
          if (!previousTime || new Date().getTime() - previousTime > 500) {
              current = current + 180;
              vpRefreshI.style.transform = 'rotateZ(' + current + 'deg)';
              previousTime = new Date().getTime();
          }
      });
      $('.type-switch a').on('click', function(){
          if ($(this).hasClass('active')) return 
          $('.type-switch a').removeClass('active')
          $(this).addClass('active')
          it.reset($(this).attr('type') || 'simple')
          $('.intelligent-detection').hide();
          if($(this).attr('type') === 'hardDisabledAi') {
              $('.intelligent-detection.text-grey').show();
          } else {
              $('.intelligent-detection.ai').show();
          }
          location.hash = $(this).attr('type') ? '#' + $(this).attr('type') : ''
      })
      if(el) {
          el.click()
      }
  };
  VD.prototype = {
      constructor: VD,

  }
  VD.toast = function(msg) {
      var el = $('.tip-block');
      el.html(msg);
      el.show();
      var timer = setTimeout(function() {
          el.hide();
          clearTimeout(timer);
      }, 1000)
  }
  VD.validatePhone = function(value) {
      var reg = /^[0-9]*$/;
      if (value.length == 11 && reg.test(value)) {
          return true;
      } else {
          return false;
      }
  }
  
  window.VD = VD;
})();