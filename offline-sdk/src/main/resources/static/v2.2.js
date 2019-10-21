(function () {
  'use strict';

  /* ie not support HTMLElement */
  /* eslint-disable */
  window.HTMLElement = window.HTMLElement || Element;
  if (!Array.prototype.map) {
    Array.prototype.map = function (callback, thisArg) {

      var T, A, k;

      if (this == null) {
        throw new TypeError(" this is null or not defined");
      }

      // 1. 将O赋值为调用map方法的数组.
      var O = Object(this);

      // 2.将len赋值为数组O的长度.
      var len = O.length >>> 0;

      // 3.如果callback不是函数,则抛出TypeError异常.
      if (Object.prototype.toString.call(callback) != "[object Function]") {
        throw new TypeError(callback + " is not a function");
      }

      // 4. 如果参数thisArg有值,则将T赋值为thisArg否则T为undefined.
      if (thisArg) {
        T = thisArg;
      }

      // 5. 创建新数组A,长度为原数组O长度len
      A = new Array(len);

      // 6. 将k赋值为0
      k = 0;

      // 7. 当 k < len 时,执行循环.
      while (k < len) {

        var kValue, mappedValue;

        //遍历O,k为原数组索引
        if (k in O) {

          //kValue为索引k对应的值.
          kValue = O[k];

          // 执行callback,this指向T,参数有三个.分别是kValue:值,k:索引,O:原数组.
          mappedValue = callback.call(T, kValue, k, O);

          // 返回值添加到新数组A中.
          A[k] = mappedValue;
        }
        // k自增1
        k++;
      }

      // 8. 返回新数组A
      return A;
    };
  }
  if (!Array.prototype.includes) {
    Array.prototype.includes = function (searchElement, fromIndex) {
      if (this == null) {
        throw new TypeError('"this" is null or not defined');
      }
      var o = Object(this);
      var len = o.length >>> 0;
      if (len === 0) {
        return false;
      }
      var n = fromIndex | 0;
      var k = Math.max(n >= 0 ? n : len - Math.abs(n), 0);
      while (k < len) {
        if (o[k] === searchElement) {
          return true;
        }
        k++;
      }
      return false;
    };
  }
  if (!Array.prototype.findIndex) {
    Array.prototype.findIndex = function (predicate) {
      // 1. Let O be ? ToObject(this value).
      if (this == null) {
        throw new TypeError('"this" is null or not defined');
      }

      var o = Object(this);

      // 2. Let len be ? ToLength(? Get(O, "length")).
      var len = o.length >>> 0;

      // 3. If IsCallable(predicate) is false, throw a TypeError exception.
      if (typeof predicate !== 'function') {
        throw new TypeError('predicate must be a function');
      }

      // 4. If thisArg was supplied, let T be thisArg; else let T be undefined.
      var thisArg = arguments[1];

      // 5. Let k be 0.
      var k = 0;

      // 6. Repeat, while k < len
      while (k < len) {
        // a. Let Pk be ! ToString(k).
        // b. Let kValue be ? Get(O, Pk).
        // c. Let testResult be ToBoolean(? Call(predicate, T, « kValue, k, O »)).
        // d. If testResult is true, return k.
        var kValue = o[k];
        if (predicate.call(thisArg, kValue, k, o)) {
          return k;
        }
        // e. Increase k by 1.
        k++;
      }

      // 7. Return -1.
      return -1;
    };
  }
  if (!Object.create) {
    Object.create = function (obj) {
      var F = function F() {};
      F.prototype = obj;
      return new F();
    };
  }

  var config = {
    vid: null,
    scene: '',
    container: null,
    type: 'float',
    style: 'dark',
    lang: 'zh-CN',
    ai: true,
    https: true,
    guide: true,
    aiAnimation: true,
    protocol: 'https://',
    // css_version: 'downtime',
    css_version: '2.2.3',
    cdn_servers: ['cdntest.vaptcha.com'],
    // cdn_servers: ['cdn.vaptcha.com'],
    api_server: 'api.vaptchatest.com/v2',
    canvas_path: '/canvas.min.js',
    outage: ''
  };

  var _typeof = typeof Symbol === "function" && typeof Symbol.iterator === "symbol" ? function (obj) { return typeof obj; } : function (obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; };

  function isUndef(v) {
    return v === undefined || v === null;
  }

  function isDef(v) {
    return v !== undefined && v !== null;
  }

  function isObject(obj) {
    return obj !== null && (typeof obj === 'undefined' ? 'undefined' : _typeof(obj)) === 'object';
  }

  // $FlowFixMe
  function isElement(ele) {
    return (typeof ele === 'undefined' ? 'undefined' : _typeof(ele)) === 'object' && ele instanceof HTMLElement;
  }

  function getFunctionName(fn) {
    var match = fn && fn.toString().match(/^\s*function (\w+)/);
    return match ? match[1] : '';
  }

  function getIEVersion() {
    var reIE = new RegExp('MSIE (\\d+\\.\\d+);');
    var userAgent = navigator.userAgent;
    reIE.test(userAgent);
    // $FlowFixMe
    return parseFloat(RegExp.$1) || Infinity;
  }

  /**
   * Mix properties into target object.
   */
  function extend(to, _from) {
    /* eslint guard-for-in: "off" */
    for (var key in _from) {
      to[key] = _from[key];
    }
    return to;
  }
  /**
   * Create a cached version of a pure function.
   */
  function cached(fn) {
    var cache = Object.create(null);
    return function cachedFn(str) {
      var hit = cache[str];
      return hit || (cache[str] = fn(str));
    };
  }
  /**
   * Get the raw type string of a value e.g. [object Object]
   */
  var _toString = Object.prototype.toString;

  function toRawType(value) {
    return _toString.call(value).slice(8, -1);
  }

  /**
   * 获取对象所对应键值数组的交集 {a: 1, b:1} , ['a'] => {a: 1}
   */
  function objectIntersect(from, it) {
    var newObject = {};
    it.map(function (v) {
      if (isDef(from[v])) newObject[v] = from[v];
      return null;
    });
    return newObject;
  }

  var getQueryString = cached(function (url) {
    var result = {};
    var queryString = url && url.indexOf('?') !== -1 && url.split('?')[1] || window.location.search.substring(1);
    var querys = queryString.split('&');
    for (var i = 0; i < querys.length; i++) {
      var r = querys[i].split('=');
      result[decodeURIComponent(r[0])] = decodeURIComponent(r[1]);
    }
    return result;
  });

  function throwError(msg) {
    throw new Error('Vaptcha error: ' + msg);
  }

  var capitalize = cached(function (str) {
    return str.charAt(0).toUpperCase() + str.slice(1);
  });

  function fnAfter(fn, after) {
    var _this = this,
        _arguments = arguments;

    return function () {
      var ret = fn.apply(_this, _arguments);
      after.apply(_this, _arguments);
      return ret;
    };
  }

  function fnBefore(fn, before) {
    var _this2 = this,
        _arguments2 = arguments;

    return function () {
      if (!before.apply(_this2, _arguments2)) return null;
      return fn.apply(_this2, _arguments2);
    };
  }

  // downtime 容错
  function isDowntime(config) {
    var mode = config.mode,
        outage = config.outage;

    if (!mode) return false;
    mode = mode.toLowerCase();
    return mode === 'downtime' && !!outage;
  }

  var Validator = function Validator(data) {
    this.data = data;
    this.valiudateFuns = [];
    this.ruleFuns = {
      required: function required(val, errorMsg) {
        if (isUndef(val) || val.length === 0) {
          return errorMsg;
        }
        return null;
      }
    };
  };

  Validator.prototype = {
    constructor: Validator,
    /**
     * 添加验证规则
     */
    addValidateRules: function addValidateRules(ruleFuns) {
      extend(this.ruleFuns, ruleFuns);
    },

    /**
     * 添加需要验证的数据项
     * @param key 需要验证的数据对应的键值
     * @param ruleString 验证规则 e.g: 'length:6'
     * @param errorMsg 验证错误返回的消息
     */
    add: function add(key, ruleString, errorMsg) {
      var _this = this;

      var arr = ruleString.split(':');
      var strategy = arr.shift();
      var ruleFun = this.ruleFuns[strategy];
      arr.unshift(this.data[key]);
      arr.push(errorMsg);
      if (ruleFun) this.valiudateFuns.push(function () {
        return ruleFun.apply(_this, arr);
      });else console.warn('Validator warning: rule ' + strategy + ' is not defined');
    },

    /**
     * 执行验证，验证失败则返回false，并抛出错误信息
     */
    validate: function validate() {
      for (var i = 0, fun; fun = this.valiudateFuns[i++];) {
        var msg = fun();
        if (msg) {
          throwError(msg);
          return false;
        }
      }
      return true;
    }
  };

  var code = {
    AccessDenied: '0101',
    RefreshAgain: '0102',
    Success: '0103',
    Fail: '0104',
    RefreshTooFast: '0105',
    RefreshTanto: '0106',
    DrawTanto: '0107',
    Attack: '0108',
    jsonpTimeOut: '0703',
    challengeExpire: '1002'
  };

  /* eslint-disable */

  function Promise(executor) {
    var _this = this;

    this.state = 'pending';
    this.value = undefined;
    this.reason = undefined;
    this.onResolveAsyncCallbacks = [];
    this.onRejectAsyncCallbacks = [];
    var resolve = function resolve(value) {
      if (_this.state === 'pending') {
        _this.state = 'fulfilled';
        _this.value = value; // 保存结果
        _this.onResolveAsyncCallbacks.map(function (fn) {
          return fn();
        }); // 执行异步回调
      }
    };
    var reject = function reject(reason) {
      if (_this.state === 'pending') {
        _this.state = 'rejected';
        _this.reason = reason;
        _this.onRejectAsyncCallbacks.map(function (fn) {
          return fn(reason);
        });
      }
    };
    try {
      executor(resolve, reject);
    } catch (err) {
      reject(err);
    }
  }

  Promise.prototype.then = function (onFulfilled) {
    var _this2 = this;

    if (this.state === 'fulfilled') {
      var result = onFulfilled(this.value);
      // 如果结果是一个Promise对象则返回结果，实现链式调用
      if (isObject(result) && getFunctionName(result.constructor) === 'Promise') {
        return result;
      }
    }
    if (this.state === 'pending') {
      // this.onResolveAsyncCallbacks.push(onFulfilled)
      // 返回Promise 避免异步返回 this，实现链式调用
      return new Promise(function (resolve) {
        _this2.onResolveAsyncCallbacks.push(function () {
          var result = onFulfilled(_this2.value);
          if (isObject(result) && getFunctionName(result.constructor) === 'Promise') {
            return result.then(resolve);
          }
          resolve(result);
        });
      });
    }
    return this;
  };

  Promise.prototype['catch'] = function (onRejected) {
    if (this.state === 'rejected') {
      onRejected(this.reason);
    }
    if (this.state === 'pending') {
      this.onRejectAsyncCallbacks.push(onRejected);
    }
    return this;
  };

  Promise.resolve = function (value) {
    return new Promise(function (resolve) {
      resolve(value);
    });
  };
  Promise.reject = function (value) {
    return new Promise(function (resolve, reject) {
      reject(value);
    });
  };

  var jsonp = function () {
    var timeout = 10000; // 超时时间
    var protocol = config.protocol;
    var apiServer = config.api_server;
    var formatParams = function formatParams(params) {
      var paramsUrl = '';
      for (var key in params) {
        if (Object.prototype.hasOwnProperty.call(params, key)) {
          paramsUrl += '&' + key + '=' + encodeURIComponent(params[key]);
        }
      }
      return paramsUrl;
    };
    var getUrl = function getUrl(url, params) {
      var paramsUrl = formatParams(params);
      var isFullUrl = url.indexOf('http://') > -1 || url.indexOf('https://') > -1;
      url.indexOf('?') < 0 && (paramsUrl = '?' + paramsUrl.slice(1));
      return isFullUrl ? '' + url + paramsUrl : '' + protocol + apiServer + url + paramsUrl;
    };

    var createJsonpScript = function createJsonpScript(url) {
      var head = document.getElementsByTagName('head')[0];
      var script = document.createElement('script');
      script.charset = 'UTF-8';
      script.src = url;
      head.appendChild(script);
      return {
        remove: function remove() {
          head.removeChild(script);
        }
      };
    };
    var _jsonp = function _jsonp(url, params, isStatic) {
      params = params || {};
      isStatic = isStatic || false;
      // apiServer = params.api_server || config.api_server
      // params.api_server && delete params.setConfig
      return new Promise(function (resolve) {
        if (isStatic) {
          // if(window['static'] == null) 
          extend(params, {
            ccc: 'static',
            t: new Date().valueOf()
          });
          url = getUrl(url, params);
          var script = createJsonpScript(url);
          var timer = setTimeout(function () {
            clearTimeout(timer);
            window['static'] = null;
            // script.remove()
            resolve();
          }, 5000);
          window['static'] = function () {
            resolve.apply(this, arguments);
            script.remove();
            window['static'] = null;
          };
        } else {
          var callbackFuncName = 'VaptchaJsonp' + new Date().valueOf();
          // 同一时间的请求，做处理，待优化
          if (window[callbackFuncName]) callbackFuncName = callbackFuncName + '1';
          extend(params, {
            callback: callbackFuncName
          });
          url = getUrl(url, params);
          var _script = createJsonpScript(url);
          var _timer = setTimeout(function () {
            clearTimeout(_timer);
            window[callbackFuncName] = null;
            _script.remove();
            resolve({
              code: '0703', /* returnCode.jsonpTimeOut */
              msg: 'Time out,Refresh Again!'
            });
          }, timeout);

          window[callbackFuncName] = function () {
            clearTimeout(_timer);
            resolve.apply(this, arguments);
            _script.remove();
            window[callbackFuncName] = null;
          };
        }
      });
    };

    _jsonp.setConfig = function (cfg) {
      protocol = cfg.protocol || protocol;
      apiServer = cfg.api_server || apiServer;
    };

    return _jsonp;
  }();

  var api = {
    getConfig: function getConfig(data) {
      var params = {
        id: data.vid,
        type: data.type,
        scene: data.scene || ''
      };
      return jsonp('/config', params);
    },
    refresh: function refresh(data) {
      return jsonp('/refresh', data);
    },
    click: function click(data) {
      return jsonp('/click', data);
    },
    get: function get(data) {
      return jsonp('/get', data);
    },
    verify: function verify(data) {
      return jsonp('/verify', data);
    },
    userbehavior: function userbehavior(data) {
      return jsonp('/userbehavior', data);
    },
    staticConfig: function staticConfig(data) {
      return jsonp(data.protocol + 'channel.vaptcha.com/config/' + data.id, {}, true);
      // return jsonp('http://cdntest1.vaptcha.com/static/'+ data, {}, true)
    }
  };

  var errorMsgs = {
    en: {
      '0201': 'id empty',
      '0202': 'id error',
      '0208': 'scene error',
      '0209': 'request used up',
      '0906': 'params error',
      '0702': 'domain does not match'
    },
    'zh-CN': {
      '0702': '\u9A8C\u8BC1\u5355\u5143\u4E0E\u57DF\u540D\u4E0D\u5339\u914D' //验证单元与域名不匹配
    }
  };
  // const getChallengeReturnCode = {
  //   IdEmpty: '0201',
  //   IdError: '0202',
  //   TimeStampError: '0203',
  //   TimeStampExpire: '0204',
  //   TimeStampEmpty: '0205',
  //   SignatureError: '0206',
  //   SignatureEmpty: '0207',
  //   SceneError: '0208',
  //   RequestUsedUp: '0209',
  // }

  var Vaptcha = function () {
    var loading = false;
    /**
     * 宕机模式配置
     */
    var downTimeConfig = function downTimeConfig(config) {
      console.log(config);
      var validator = new Validator(config);
      validator.add('outage', 'required', 'please configure outage');
      validator.validate();
      extend(config, {
        js_path: 'vaptcha-sdk-downtime.2.0.2.js',
        api_server: window.location.host,
        protocol: window.location.protocol + '//',
        mode: 'DownTime'
        // guide: true,
        // 'css_version': '1.2.7',
        // 'api_server': 'api.vaptcha.com/v2',
        // type: 'character'
      });
      jsonp.setConfig(config);
      return jsonp(config.outage, {
        action: 'get'
      }).then(function (result) {
        if (result.code !== code.Success) {
          throwError(errorMsgs[result.msg] || result.msg);
          return Promise.reject(result.code);
        }
        extend(config, result);
        return Promise.resolve();
      });
    };
    var getStaticConfig = function getStaticConfig(config) {
      return api.staticConfig({ protocol: config.protocol, id: config.vid }).then(function (serverConfig) {
        // console.log(serverConfig)
        // if(serverConfig.dt) return downTimeConfig(config)
        // else {
        //   extend(config, serverConfig)
        // }
        return Promise.resolve(serverConfig);
      });
    };
    /**
     * 拉取并合并服务端的配置
     * 服务端接口返回 challenge
     */
    var mergeServerConfig = function mergeServerConfig(config) {
      if (isDowntime(config)) {
        return downTimeConfig(config);
      }
      return getStaticConfig(config).then(function (serverConfig) {
        if (serverConfig.state) return Promise.reject('VAPTCHA cell error');
        if (serverConfig.dt) {
          if (config.outage == "") return Promise.reject('downtime not configured');
          extend(config, { mode: 'downTime' });
          return downTimeConfig(config);
        }
        extend(config, { api_server: serverConfig.api });
        jsonp.setConfig(config);
        return api.getConfig(config);
      }).then(function (serverConfig) {
        // api.getConfig(config).then((serverConfig) => {
        if (serverConfig.code !== code.Success) {
          var msgs = errorMsgs[config.lang] || errorMsgs['zh-CN'];
          serverConfig.msg === '0702' && alert('Vaptcha error: ' + msgs[serverConfig.msg]);
          throwError(msgs[serverConfig.msg] || serverConfig.msg);
          return Promise.reject(serverConfig.code);
        }

        if (serverConfig.type !== config.type) {
          serverConfig.mode = serverConfig.type;
          serverConfig.type = config.type;
        }
        extend(config, serverConfig);

        if (isDowntime(config)) {
          return downTimeConfig(config);
        }

        // extend(config, {
        //   guide: true,
        //   css_version: '1.2.7',
        //   // api_server: 'api.vaptcha.com/v2',
        //   mode: 'character'
        // })
        return Promise.resolve();
        // })
      });
    };
    var getCdnUrl = function getCdnUrl(config, path) {
      return '' + config.protocol + config.cdn_servers[0] + '/' + path;
    };
    var insertStyles = function insertStyles(url) {
      var head = document.getElementsByTagName('head')[0];
      var styleName = 'vaptcha_style';
      var style = document.getElementById(styleName);
      return new Promise(function (resolve) {
        if (isUndef(style)) {
          style = document.createElement('link');
          extend(style, {
            rel: 'stylesheet',
            type: 'text/css',
            href: url,
            id: styleName,
            onload: resolve
          });
          head && head.appendChild(style);
        } else {
          resolve();
        }
      });
    };
    var loadScript = function loadScript(url) {
      var head = document.getElementsByTagName('head')[0];
      // const scriptId = `vaptcha_script_${name}`
      // let script: any = document.getElementById(scriptId)
      var script = document.querySelector('script[src=\'' + url + '\']');
      return new Promise(function (resolve) {
        if (isDef(script)) {
          script.loaded ? resolve() : setTimeout(function () {
            return loadScript(url).then(resolve);
          });
          return;
        }
        script = document.createElement('script');
        var loadCallback = function loadCallback() {
          if (!script.readyState || script.readyState === 'loaded' || script.readyState === 'complete') {
            resolve();
            script.loaded = true;
            script.onload = null;
            script.onreadystatechange = null;
          }
        };
        extend(script, {
          async: true,
          // id: scriptId,
          charset: 'utf-8',
          src: url,
          onerror: function onerror() {
            return throwError('load sdk timeout');
          },
          onload: loadCallback,
          onreadystatechange: loadCallback
        });
        head.appendChild(script);
      });
    };
    var getCaptcha = function getCaptcha(_ref) {
      var sdkName = _ref.sdkName,
          config = _ref.config;

      // const sdkSrc = '/dist/vaptcha-sdk.js'
      // const version = '2.1.3'
      // var sdkSrc = '/dist/0.0.1/vaptcha-sdk-' + sdkName + '.0.0.1.js';
      // const sdkSrc = `/dist/${version}/vaptcha-sdk-${sdkName}.${version}.js`
      // const sdkSrc = `/dist/${version}/vaptcha-sdk-mobile.${version}.js`
      // const sdkSrc = getCdnUrl(config, 'vaptcha-sdk-mobile.2.1.4.js') // eslint-disable-line
      const sdkSrc = getCdnUrl(config, config.js_path) // eslint-disable-line
      // const sdkSrc = 'https://cdn.vaptcha.com/' + getCdnUrl(config, config.js_path).split('/')[3].split('.')[0] + '.2.5.2.js' 
      return loadScript(sdkSrc).then(function () {
        var funName = sdkName == 'downtime' ? 'DownTime' : capitalize(sdkName);
        // loadScript(config.type, getCdnUrl(config, config.js_path, () => {
        var VaptchaConstruct = window['_' + funName + 'Vaptcha'];
        var captcha = new VaptchaConstruct(config);
        // 给vaptcha对象添加重置方法
        captcha.vaptcha.resetCaptcha = function (sdkName, newConfig) {
          //eslint-disable-line
          extend(config, newConfig);
          getCaptcha({ sdkName: sdkName, config: config }).then(function (newCaptcha) {
            captcha.destroy();
            captcha.options = newCaptcha.options;
            captcha.vaptcha = newCaptcha.vaptcha;
            newCaptcha.render();
            // 如果是点击式主动触发第一次点击
            config.mode === 'character' && ['click', 'float', 'popup'].includes(config.type) && newCaptcha.vaptcha.dtClickCb({ target: newCaptcha.vaptcha.btnDiv });
          });
        };
        return Promise.resolve(captcha);
      });
    };
    var render = function render(config) {
      loading = true;
      config.https = true;
      config.protocol = 'https://'; // config.https ? 'https://' : 'http://'
      jsonp.setConfig(config);
      // 基础类型以外的值初始化为popup
      !['embed', 'popup', 'invisible'].includes(config.type) && (config.type = 'popup');
      // ie9以下加载canvas.min.js
      getIEVersion() < 9 && loadScript(getCdnUrl(config, config.canvas_path));
      // getIEVersion() < 9 && loadScript('https://cdnjs.cloudflare.com/ajax/libs/flot/0.7/excanvas.min.js')
      // 创建验证类， 用于验证传入的配置
      var validator = new Validator(config);
      validator.addValidateRules({
        elementOrSelector: function elementOrSelector(val, msg) {
          // 如果是string， 则将其实识别为selector
          if (toRawType(config.container) === 'String') {
            config.container = document.querySelector(config.container);
          }
          // jQuery对象
          if (isObject(config.container) && isElement(config.container[0])) {
            config.container = config.container[0];
          }
          if (!isElement(config.container)) {
            return msg;
          }
        }
      });
      validator.add('vid', 'required', 'please configure vid');
      // 非隐藏式container不能为空
      config.type !== 'invisible' && validator.add('container', 'elementOrSelector', 'please configure container with element or selector');
      if (validator.validate()) {
        return mergeServerConfig(config).then(function () {
          // config.css_version = '2.1.0'
          var cssName = config.https ? 'css/theme_https.' + config.css_version + '.css' : 'css/theme.' + config.css_version + '.css';
          // const cssName = `theme_https.${config.css_version}.css`
          // const url = 'https://cdn.vaptcha.com/theme_https.2.2.1.css'
          var url = getCdnUrl(config, cssName);
          console.log(config);
          // const url = 'http://192.168.0.103:4396/css/theme.2.1.8.css'
          return insertStyles(url);
        }).then(function () {
          var sdkName = config.mode || config.type;
          loading = false;
          console.log(config);
          return getCaptcha({ sdkName: sdkName, config: config });
        });
      }
    };
    return function VP(config) {
      return new Promise(function (resolve) {
        loading ? setTimeout(function () {
          VP(config).then(resolve);
        }, 1000) : render(config).then(resolve);
      })['catch'](function (e) {
        loading = false;
        throwError(e);
        return Promise.reject(e);
      });
    };
  }();

  /**
   * 自动将配置了vid的dom作为Vaptcha的conatiner 初始化
   */
  var autoRenderVaptcha = function () {
    /**
     * 获取dom上配置的json字符串
     */
    var getDomConfig = function getDomConfig(dom) {
      var configString = dom.getAttribute('data-config');
      var domConfig = {};
      if (isDef(configString)) {
        try {
          // $FlowFixMe
          domConfig = JSON.parse(configString);
        } catch (e) {
          throwError('dom config format error');
        }
      }
      return domConfig;
    };

    /**
     * 获取dom上的vid
     */
    var getDomVid = function getDomVid(dom) {
      var vid = dom.getAttribute('data-vid');
      return isDef(vid) ? { vid: vid } : {};
    };

    /**
     * 渲染验证码
     */
    var renderCaptcha = function renderCaptcha(container, domConfig) {
      var config$$1 = Object.create(config);
      config$$1.container = container;
      /* $FlowFixMe */
      extend(config$$1, domConfig);
      /**
       * 只要vid存在就初始化
       */
      if (isDef(config$$1.vid)) {
        Vaptcha(config$$1).then(function (obj) {
          console.log(obj);
          obj.renderTokenInput();
          obj.render();
        });
      }
    };

    return function () {
      var vidCtxs = document.querySelectorAll('[data-vid]');
      var configCtxs = document.querySelectorAll('[data-config]');
      // 初始化含有[data-config]的dom
      for (var i = 0; i < configCtxs.length; i++) {
        var config$$1 = getDomConfig(configCtxs[i]);
        renderCaptcha(configCtxs[i], config$$1);
      }
      // 初始化含有[data-vid]且无[data-config]的dom
      for (var _i = 0; _i < vidCtxs.length; _i++) {
        if (!Array.prototype.includes.call(configCtxs, vidCtxs[_i])) {
          var _config = getDomVid(vidCtxs[_i]);
          renderCaptcha(vidCtxs[_i], _config);
        }
      }
    };
  }();

  window.onload = autoRenderVaptcha;

  window.vaptcha = function (userConfig) {
    var config$$1 = Object.create(config);
    extend(config$$1, userConfig);
    console.log(config$$1);
    return Vaptcha(config$$1);
  };

}());
