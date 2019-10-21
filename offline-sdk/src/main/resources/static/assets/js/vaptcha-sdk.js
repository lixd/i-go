
(function (window, undefined) {
    //本地变量
    var vaptcha,
        challenge,
        siteId,
        VaptchaInterval,
        VaptchaTime, //lasted refresh time point
        VaptchaInitTime,
        VaptchaData,
        interval,
        vaptchaUrl = "api.vaptcha.com/",
        imgUrl = "static.vaptcha.com/",
        requestAmount = 0,
        validateSuccessCallback = new Function(),
        validateFailCallback = new Function(),
        validateErrorCallback = new Function(),
        getReadyCallback = new Function(),
        refreshCallback = new Function(),
        initFailCallback = new Function(),
        vaptchaProduct = "float",
        vaptchaCanvasDiv,
        canvas,
        canvasContext,
        canvasWidth,
        canvasHeight,
        clickX = new Array(),
        clickY = new Array(),
        clickDrag = new Array(),
        dragSpotData = new Array(),
        characteristics = new Array(),
        startPoint = null,
        previousPoint = null,
        paint = false,
        allowLeftPaint = false,
        isValidated = false,
        startTime,
        sample = "()*,-./0123456789:?@ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz",
        _sample = "abcdefgh234lmntuwxyz",
        isModernBrowser = false,
        passcheck = false,
        vaptchaContainer,
        protocol;
    /*辅助方法*/
    //获取符合条件的最后一个匹配值的index
    function _findLastIndex(array, predicate/*,index*/) {
        var o = Object(array);
        var len = o.length >>> 0;
        if (typeof predicate !== 'function') {
            throw new TypeError('predicate must be a function');
        }
        var thisArg = arguments[2];
        var k = len - 1;
        while (k >= 0) {
            var kValue = o[k];
            if (predicate.call(array, kValue)) {
                return k;
            }
            k--;
        }
        return -1;
    }
    //获取符合条件的匹配值个数
    function _countBy(array, predicate/*,index*/) {
        if (this == null) {
            throw new TypeError('"this" is null or not defined');
        }

        var o = Object(array);
        var len = o.length >>> 0;
        if (typeof predicate !== 'function') {
            throw new TypeError('predicate must be a function');
        }
        var thisArg = arguments[2];
        var n = 0;
        var k = 0;
        while (k < len) {
            var kValue = o[k];
            if (predicate.call(array, kValue)) {
                n++;
            }
            k++;
        }
        return n;
    }
    //获取查询字符串参数
    function _getQueryString(url) {
        var result = {},
            queryString = (url && url.indexOf("?") != -1 && url.split("?")[1]) || location.search.substring(1),
            reg = /([^?&=]+)=([^&]*)/g,
            m = reg.exec(queryString);
        if (m) {
            result[decodeURIComponent(m[1])] = decodeURIComponent(m[2]);
        }
        return result;
    }
    //以jsonp方式发送数据
    function _getJsonp(url, params, callbackFuncName, callback) {
        var paramsUrl = "",
            jsonp = _getQueryString(url)[callbackFuncName];
        for (var key in params) {
            paramsUrl += "&" + key + "=" + encodeURIComponent(params[key]);
        }
        url += paramsUrl;
        window[jsonp] = function (data) {
            window[jsonp] = undefined;
            try {
                delete window[jsonp];
            } catch (e) { }

            if (head) {
                var node = document.getElementById("vaptchaJsonpTempScript");
                head.removeChild(node);
            }
            callback(data);
        };

        var head = document.getElementsByTagName('head')[0];
        var script = document.createElement('script');
        script.setAttribute("id", "vaptchaJsonpTempScript")
        script.charset = "UTF-8";
        script.src = url;
        head.appendChild(script);
        return true;
    }
    //创建XHR对象;
    function _createXHR() {
        if (typeof XMLHttpRequest != "undefined") {
            return new XMLHttpRequest();
        } else if (typeof arguments.callee.activeXString != "string") {
            var versions = ["MSXML2. XMLHttp. 6. 0", "MSXML2. XMLHttp. 3. 0", "MSXML2. XMLHttp"],
                i,
                len;
            for (i = 0, len = versions.length; i < len; i++) {
                try {
                    new ActiveXObject(versions[i]);
                    arguments.callee.activeXString = versions[i];
                    break;
                } catch (ex) {
                    //跳过 
                }
            }
            return new ActiveXObject(arguments.callee.activeXString);
        } else {
            throw new Error("No XHR object available.");
        }
    }
    //事件处理函数绑定方法
    function _eventHandler(target, eventType, handler) {
        if (target) {
            if (target.addEventListener) {
                target.addEventListener(eventType, handler, false);
            } else if (target.attachEvent) {
                target.attachEvent("on" + eventType, handler);
            } else {
                target["on" + eventType] = handler;
            }
        }
    }
    //手动触发事件
    function _fireEvent(eventType, event, target) {
        if (target) {
            if (document.createEvent) {
                var eventInstance = document.createEvent(eventType);
                eventInstance.initEvent(event, true, true);
                target.dispatchEvent(eventInstance);
            } else if (document.createEventObject) {
                target.fireEvent("on" + event);
            }
        }
    }
    //阻止默认行为
    function _stopDefault(e) {
        e = e || window.event;
        (e.preventDefault) ? e.preventDefault() : e.returnValue = false;
    }
    //根据鼠标事件对象返回鼠标坐标
    function _getMouseCoords(e) {
        e = e || window.event;
        if (e.pageX || e.pageY) {
            return { x: e.pageX, y: e.pageY };
        }
        return {
            x: e.clientX + document.body.scrollLeft - document.body.clientLeft,
            y: e.clientY + document.body.scrollTop - document.body.clientTop
        };
    }
    //超时检测
    function _overtimeCheck() {
        var now = new Date().getTime();
        if ((now - startTime) > 5000) {
            return true;
        }
    }
    //插入style标签
    function _insertStyle(styleText) {
        var head = document.getElementsByTagName('head')[0];
        var style = document.createElement('style');
        style.innerText = styleText;
        head.appendChild(style);
    }
    /**
     * get this cookie-key's value or null
     * @param name cookie-key
     */
    function _getCookie(name) {
        var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
        if (arr = document.cookie.match(reg))
            return arr[2];
        else
            return null;
    }
    /**
     * set cookie
     * @param  name cookie-name
     * @param  value cookie-value
     * @param  expiretime duration ,unit: millisecond;
     */
    function _setCookie(name, value, expiretime) {
        var exdate = new Date();
        var pushTime = exdate.getTime();
        expiretime = expiretime + pushTime;
        exdate.setTime(expiretime);
        document.cookie = name + '=' + value + ((expiretime == null) ? "" : ";expires=" + exdate.toUTCString());
    }
    /*数据验证编码*/
    //检测两点之间的距离,大于指定值才加入采样数列中
    function _isAvailableDate(coord) {
        if (startPoint) {
            var offsetX = Math.abs(startPoint.x - coord.x);
            var offsetY = Math.abs(startPoint.y - coord.y);
        } else {
            var offsetX = Math.abs(previousPoint.x - coord.x);
            var offsetY = Math.abs(previousPoint.y - coord.y);
        }
        if (Math.sqrt(offsetX * offsetX + offsetY * offsetY) > 5) {
            return true;
        } else {
            return false;
        }
    }
    //由上一个点到当前点这两坐标的[屏幕]标点,返回该向量的角度[0,360)
    function _getAngle(offsetX, offsetY) {
        var x = offsetX;
        var y = offsetY;
        var angle = 0;
        //一,四象限
        if (x > 0) {
            if (y == 0) {
                angle = 0;
            }
            else if (y > 0) {
                angle = Math.atan(y / x) * 180 / Math.PI; //一象限
            }
            else if (y < 0) {
                angle = 360 - Math.atan((-y) / x) * 180 / Math.PI; //四象限
            }
        }
        //二,三象限
        else if (x < 0) {
            if (y == 0) {
                angle = 180;
            }
            else if (y > 0) {
                angle = 180 - Math.atan(y / (-x)) * 180 / Math.PI; //二象限
            }
            else if (y < 0) {
                angle = Math.atan(y / x) * 180 / Math.PI + 180; //三象限
            }
        }
        //x === 0
        else {
            if (y > 0) {
                //90度
                angle = 90;
            }
            else if (y < 0) {
                //270度
                angle = 270;
            }
            else {
                //(0,0)
                angle = undefined;
            }
        }
        return angle;
    }
    //计算特征值
    function _calculateCharacteristic(spotData) {
        var speed,
            acceleration,
            coordOffset,
            timeInterval,
            angle,
            angleOffset,
            distance;
        characteristics = new Array();
        characteristics.push({
            "speed": 0,
            "coordOffset": {
                "offsetX": 0,
                "offsetY": 0
            },
            "acceleration": 0,
            "angle": 0,
            "angleOffset": 0
        })
        for (var i = 1; i < spotData.length; i++) {
            timeInterval = spotData[i].time - spotData[i - 1].time;
            coordOffset = {
                "offsetX": spotData[i].x - spotData[i - 1].x,
                "offsetY": spotData[i].y - spotData[i - 1].y
            };
            distance = Math.sqrt(Math.pow(coordOffset.offsetX, 2) + Math.pow(coordOffset.offsetY, 2));
            angle = _getAngle(coordOffset.offsetX, coordOffset.offsetY);
            angleOffset = angle - characteristics[i - 1].angle;
            speed = distance / timeInterval;
            acceleration = (speed - characteristics[i - 1].speed) / timeInterval;
            characteristics.push({
                "speed": speed,
                "coordOffset": coordOffset,
                "acceleration": acceleration,
                "angle": angle,
                "angleOffset": angleOffset
            })
        }
        console.log(characteristics);
    }
    //分析角度偏移量特征值
    function _analysisAngle() {
        function _sortNumber(a, b) {
            return a.angleOffset - b.angleOffset
        }
        function isZeroOffset(element) {
            return element.angleOffset === 0;
        }
        characteristics = characteristics.sort(_sortNumber);
        var index = _findLastIndex(characteristics, isZeroOffset);
        var result = (index + 1) / characteristics.length;
        if (result > 0.95) {
            return false;
        } else {
            return true;
        }
    }
    //分析加速度特征值
    function _analysisAcceleration() {
        function isZeroAcceleration(element) {
            return element.acceleration === 0;
        }
        var num = _countBy(characteristics, isZeroAcceleration);
        var result = (num) / characteristics.length;
        if (result > 0.25) {
            return false;
        } else {
            return true;
        }
    }
    //存入有效数据并绘制绘制轨迹
    function _addDrapSpotData(coord) {
        if (!!startPoint) {
            //设置起始参考点
            if (_isAvailableDate(coord)) {
                paint = true;
                previousPoint = { 'x': coord.x, 'y': coord.y };
                dragSpotData.push({ 'x': startPoint.x, 'y': startPoint.y, "time": 0 });
                dragSpotData.push({ 'x': coord.x, 'y': coord.y, "time": new Date().getTime() - startTime });
                startPoint = null;
                if (!canvas) {
                    _prepareCanvas();
                }
                _redraw();
            }
        } else {
            if (paint == true) {
                _addClick(coord.x, coord.y, true);
                if (_isAvailableDate(coord)) {
                    previousPoint = { 'x': coord.x, 'y': coord.y };
                    dragSpotData.push({ 'x': coord.x, 'y': coord.y, "time": new Date().getTime() - startTime });
                }
                _redraw();
            }
        }
    }
    //编码
    function _convertScale(number) {
        var scale = _sample.length,
            abs = Math.abs(number),
            flag = false,//3位. true为2位
            result = "",
            first, secend, third,
            length,
            multiple = parseInt(abs / scale);
        var _stack = new Array();
        //如果倍数比样本容量大
        //todo:后面优化
        if (multiple >= scale) {
            first = parseInt(number / (scale * scale));
            var r1 = first * (scale * scale);
            _stack.push(first);
            secend = parseInt((number - r1) / scale);
            _stack.push(secend);
            var r2 = secend * scale;
            third = number - r1 - r2;
            _stack.push(third);
        } else {
            first = parseInt(number / scale);
            var r1 = first * scale;
            _stack.push(first);
            secend = number - r1;
            _stack.push(secend);
            flag = true;
        }
        if (flag) { //2位用 "_"来顶替第一位
            result = "_";
        }
        length = _stack.length;
        for (var i = 0; i < length; i++) {
            result += _sample.charAt(_stack[i]);
        }
        return result;
    }
    //解码
    function _deConvertScale(code) {
        var length = code.length,
            scale = _sample.length,
            first,
            secend,
            thrid,
            result;
        for (var i = 0; i < length; i += 3) {
            if (i == 0 && code[i] == '_') {
                result = _sample.indexOf(code[i + 1]) * scale + _sample.indexOf(code[i + 2]);
            }
            else {
                result = _sample.indexOf(code[i]) * scale * scale + _sample.indexOf(code[i + 1]) * scale + _sample.indexOf(code[i + 2]);
            }
        }
        return result;
    }

    //将坐标数组转换为字符串（利用编码)
    function _assemblyCoordDate(coords) {
        var coord,
            xArray = [],
            yArray = [],
            tArray = [];
        for (var j = 0; j < coords.length; j++) {
            xArray.push(_convertScale(coords[j].x));
            yArray.push(_convertScale(coords[j].y));
            tArray.push(_convertScale(coords[j].time));
        };
        return xArray.join("") + yArray.join("") + tArray.join("");
    }
    //解析坐标数据 
    function _decodeCoords(code) {
        var xArray = [],
            yArray = [],
            tArray = [];
        for (var i = 0; i < code.length / 3; i += 3) {
            xArray.push(_deConvertScale(code.slice(i, i + 3)));
        }
        for (var i = code.length / 3; i < 2 * code.length / 3; i += 3) {
            yArray.push(_deConvertScale(code.slice(i, i + 3)));
        }
        for (var i = 2 * code.length / 3; i < code.length; i += 3) {
            tArray.push(_deConvertScale(code.slice(i, i + 3)));
        }
        //for test：重装dragSpotData
        dragSpotData = new Array();
        for (var i = 0; i < xArray.length; i++) {
            dragSpotData.push({ 'x': xArray[i], 'y': yArray[i], 'time': tArray[i] });
        }
        console.log(tArray.length);
    }
    //偏差值检测
    function _offsetCheck(coord) {
        var array = [[1, 0], [2, 0], [1, -1], [1, 1], [0, 1], [0, -1], [3, 0], [2, -1], [2, 1]],
            sample = "stuvwxyz~";
        for (var i = 0; i < array.length; i++)
            if (coord[0] == array[i][0] && coord[1] == array[i][1])
                return sample[i];
        return 0
    };
    /*绘制图形*/
    //获取画板上下文
    function _getCanvasContext() {
        if (isModernBrowser) {
            canvasContext = canvas.getContext("2d");
        } else {
            canvas = window.G_vmlCanvasManager.initElement(canvas);
            canvasContext = canvas.getContext("2d");
        }
    }
    //设置画板元素
    function _prepareCanvas() {
        vaptchaCanvasDiv = document.getElementById("vaptchaCanvasDiv");
        if (!vaptchaCanvasDiv) {
            alert("初始化失败，请刷新页面重试")
            return;
        }
        canvas = document.createElement('canvas');
        canvasWidth = document.body.offsetWidth;
        canvasHeight = document.body.offsetHeight
        canvas.setAttribute('width', canvasWidth);
        canvas.setAttribute('height', canvasHeight);
        canvas.setAttribute("id", "canvas");

        vaptchaCanvasDiv.style.cursor = "url('images/pen.ico'),url('images/pen.png'),auto"
        vaptchaCanvasDiv.appendChild(canvas);
        _getCanvasContext();
    }
    //删除画板元素
    function _removeCanvas() {
        if (canvas) {
            if (vaptchaCanvasDiv.hasChildNodes()) {
                vaptchaCanvasDiv.removeChild(canvas);
            }
        }
    }
    //清理画板
    function _clearCanvas() {
        if (canvasContext) {
            canvasContext.clearRect(0, 0, canvasWidth, canvasHeight);
        }
    }
    //添加画板轨迹坐标
    function _addClick(x, y, dragging) {
        clickX.push(x);
        clickY.push(y);
        clickDrag.push(dragging);
    };
    //绘制图形
    function _redraw() {
        if (canvasContext) {
            _clearCanvas();
            var radius;
            var i = 0;
            for (; i < clickX.length; i++) {
                radius = 5;

                canvasContext.beginPath();
                if (clickDrag[i] && i) {
                    canvasContext.moveTo(clickX[i - 1], clickY[i - 1]);
                } else {
                    canvasContext.moveTo(clickX[i], clickY[i]);
                }
                canvasContext.lineTo(clickX[i], clickY[i]);
                canvasContext.closePath();

                canvasContext.strokeStyle = '#659b41';

                canvasContext.lineJoin = "round";
                canvasContext.lineWidth = radius;
                canvasContext.stroke();
            }
            //for ie10 -
            if (!isModernBrowser) {
                canvas = window.G_vmlCanvasManager.initElement(canvas);
            }
            if (canvas.getContext) {
                //console.log("support");
            }
            //canvasContext.restore();
        }
    }
    //for test 绘制采集的数据
    function _drawTest() {
        var vaptchaCanvasDiv = document.getElementById('vaptchaCanvasDiv');
        canvas = document.createElement('canvas');
        canvasWidth = document.body.offsetWidth;
        canvas.setAttribute('width', canvasWidth);
        canvasHeight = document.body.offsetHeight
        canvas.setAttribute('height', canvasHeight);
        vaptchaCanvasDiv.appendChild(canvas);
        canvas.setAttribute('id', 'canvas');
        canvasContext = document.getElementById('canvas').getContext("2d");

        var radius;
        for (var i = 0; i < dragSpotData.length; i++) {
            dragSpotData[i].x += 100;
            dragSpotData[i].y += 100;
            radius = 5;
            canvasContext.beginPath();
            if (dragSpotData[i] && i) {
                canvasContext.moveTo(dragSpotData[i - 1].x, dragSpotData[i - 1].y);
            } else {
                canvasContext.moveTo(dragSpotData[i].x, dragSpotData[i].y);
            }
            canvasContext.lineTo(dragSpotData[i].x, dragSpotData[i].y);
            canvasContext.closePath();

            canvasContext.strokeStyle = "red";
            canvasContext.lineJoin = "round";
            canvasContext.lineWidth = radius;
            canvasContext.stroke();
        }

    }


    /*初始化方法*/
    //腻子脚本
    function _polyfill() {
        //为es3提供Array.index0f
        if (!Array.prototype.indexOf) {
            Array.prototype.indexOf = function (searchElement, fromIndex) {
                var k;
                if (this == null) {
                    throw new TypeError('"this" is null or not defined');
                }
                var o = Object(this);
                var len = o.length >>> 0;
                if (len === 0) {
                    return -1;
                }
                var n = fromIndex | 0;
                if (n >= len) {
                    return -1;
                }
                k = Math.max(n >= 0 ? n : len - Math.abs(n), 0);

                while (k < len) {
                    if (k in o && o[k] === searchElement) {
                        return k;
                    }
                    k++;
                }
                return -1;
            };
        }

        //为es3提供Array.lastIndex0f
        if (!Array.prototype.lastIndexOf) {
            Array.prototype.lastIndexOf = function (searchElement /*, fromIndex*/) {
                'use strict';
                if (this === void 0 || this === null) {
                    throw new TypeError();
                }
                var n, k,
                    t = Object(this),
                    len = t.length >>> 0;
                if (len === 0) {
                    return -1;
                }
                n = len - 1;
                if (arguments.length > 1) {
                    n = Number(arguments[1]);
                    if (n != n) {
                        n = 0;
                    }
                    else if (n != 0 && n != (1 / 0) && n != -(1 / 0)) {
                        n = (n > 0 || -1) * Math.floor(Math.abs(n));
                    }
                }
                for (k = n >= 0 ? Math.min(n, len - 1) : len - Math.abs(n); k >= 0; k--) {
                    if (k in t && t[k] === searchElement) {
                        return k;
                    }
                }
                return -1;
            };
        }

        //为es3提供Array.findIndex
        if (!Array.prototype.findIndex) {
            Object.defineProperty(Array.prototype, 'findIndex', {
                value: function (predicate) {
                    if (this == null) {
                        throw new TypeError('"this" is null or not defined');
                    }

                    var o = Object(this);
                    var len = o.length >>> 0;
                    if (typeof predicate !== 'function') {
                        throw new TypeError('predicate must be a function');
                    }
                    var thisArg = arguments[1];
                    var k = 0;
                    while (k < len) {
                        var kValue = o[k];
                        if (predicate.call(thisArg, kValue, k, o)) {
                            return k;
                        }
                        k++;
                    }
                    return -1;
                }
            });
        }

        //google exCanvas
        if (!document.createElement('canvas').getContext) {
            (function () {

                // alias some functions to make (compiled) code shorter
                var m = Math;
                var mr = m.round;
                var ms = m.sin;
                var mc = m.cos;
                var abs = m.abs;
                var sqrt = m.sqrt;

                // this is used for sub pixel precision
                var Z = 10;
                var Z2 = Z / 2;

                var IE_VERSION = +navigator.userAgent.match(/MSIE ([\d.]+)?/)[1];

                /**
                 * This funtion is assigned to the <canvas> elements as element.getContext().
                 * @this {HTMLElement}
                 * @return {CanvasRenderingContext2D_}
                 */
                function getContext() {
                    return this.context_ ||
                        (this.context_ = new CanvasRenderingContext2D_(this));
                }

                var slice = Array.prototype.slice;

                /**
                 * Binds a function to an object. The returned function will always use the
                 * passed in {@code obj} as {@code this}.
                 *
                 * Example:
                 *
                 *   g = bind(f, obj, a, b)
                 *   g(c, d) // will do f.call(obj, a, b, c, d)
                 *
                 * @param {Function} f The function to bind the object to
                 * @param {Object} obj The object that should act as this when the function
                 *     is called
                 * @param {*} var_args Rest arguments that will be used as the initial
                 *     arguments when the function is called
                 * @return {Function} A new function that has bound this
                 */
                function bind(f, obj, var_args) {
                    var a = slice.call(arguments, 2);
                    return function () {
                        return f.apply(obj, a.concat(slice.call(arguments)));
                    };
                }

                function encodeHtmlAttribute(s) {
                    return String(s).replace(/&/g, '&amp;').replace(/"/g, '&quot;');
                }

                function addNamespace(doc, prefix, urn) {
                    if (!doc.namespaces[prefix]) {
                        doc.namespaces.add(prefix, urn, '#default#VML');
                    }
                }

                function addNamespacesAndStylesheet(doc) {
                    addNamespace(doc, 'g_vml_', 'urn:schemas-microsoft-com:vml');
                    addNamespace(doc, 'g_o_', 'urn:schemas-microsoft-com:office:office');

                    // Setup default CSS.  Only add one style sheet per document
                    if (!doc.styleSheets['ex_canvas_']) {
                        var ss = doc.createStyleSheet();
                        ss.owningElement.id = 'ex_canvas_';
                        ss.cssText = 'canvas{display:inline-block;overflow:hidden;' +
                            // default size is 300x150 in Gecko and Opera
                            'text-align:left;width:300px;height:150px}';
                    }
                }

                // Add namespaces and stylesheet at startup.
                addNamespacesAndStylesheet(document);

                var G_vmlCanvasManager_ = {
                    init: function (opt_doc) {
                        var doc = opt_doc || document;
                        // Create a dummy element so that IE will allow canvas elements to be
                        // recognized.
                        doc.createElement('canvas');
                        doc.attachEvent('onreadystatechange', bind(this.init_, this, doc));
                    },

                    init_: function (doc) {
                        // find all canvas elements
                        var els = doc.getElementsByTagName('canvas');
                        for (var i = 0; i < els.length; i++) {
                            this.initElement(els[i]);
                        }
                    },

                    /**
                     * Public initializes a canvas element so that it can be used as canvas
                     * element from now on. This is called automatically before the page is
                     * loaded but if you are creating elements using createElement you need to
                     * make sure this is called on the element.
                     * @param {HTMLElement} el The canvas element to initialize.
                     * @return {HTMLElement} the element that was created.
                     */
                    initElement: function (el) {
                        if (!el.getContext) {
                            el.getContext = getContext;

                            // Add namespaces and stylesheet to document of the element.
                            addNamespacesAndStylesheet(el.ownerDocument);

                            // Remove fallback content. There is no way to hide text nodes so we
                            // just remove all childNodes. We could hide all elements and remove
                            // text nodes but who really cares about the fallback content.
                            el.innerHTML = '';

                            // do not use inline function because that will leak memory
                            el.attachEvent('onpropertychange', onPropertyChange);
                            el.attachEvent('onresize', onResize);

                            var attrs = el.attributes;
                            if (attrs.width && attrs.width.specified) {
                                // TODO: use runtimeStyle and coordsize
                                // el.getContext().setWidth_(attrs.width.nodeValue);
                                el.style.width = attrs.width.nodeValue + 'px';
                            } else {
                                el.width = el.clientWidth;
                            }
                            if (attrs.height && attrs.height.specified) {
                                // TODO: use runtimeStyle and coordsize
                                // el.getContext().setHeight_(attrs.height.nodeValue);
                                el.style.height = attrs.height.nodeValue + 'px';
                            } else {
                                el.height = el.clientHeight;
                            }
                            //el.getContext().setCoordsize_()
                        }
                        return el;
                    }
                };

                function onPropertyChange(e) {
                    var el = e.srcElement;

                    switch (e.propertyName) {
                        case 'width':
                            el.getContext().clearRect();
                            el.style.width = el.attributes.width.nodeValue + 'px';
                            // In IE8 this does not trigger onresize.
                            el.firstChild.style.width = el.clientWidth + 'px';
                            break;
                        case 'height':
                            el.getContext().clearRect();
                            el.style.height = el.attributes.height.nodeValue + 'px';
                            el.firstChild.style.height = el.clientHeight + 'px';
                            break;
                    }
                }

                function onResize(e) {
                    var el = e.srcElement;
                    if (el.firstChild) {
                        el.firstChild.style.width = el.clientWidth + 'px';
                        el.firstChild.style.height = el.clientHeight + 'px';
                    }
                }

                G_vmlCanvasManager_.init();

                // precompute "00" to "FF"
                var decToHex = [];
                for (var i = 0; i < 16; i++) {
                    for (var j = 0; j < 16; j++) {
                        decToHex[i * 16 + j] = i.toString(16) + j.toString(16);
                    }
                }

                function createMatrixIdentity() {
                    return [
                        [1, 0, 0],
                        [0, 1, 0],
                        [0, 0, 1]
                    ];
                }

                function matrixMultiply(m1, m2) {
                    var result = createMatrixIdentity();

                    for (var x = 0; x < 3; x++) {
                        for (var y = 0; y < 3; y++) {
                            var sum = 0;
                            for (var z = 0; z < 3; z++) {
                                sum += m1[x][z] * m2[z][y];
                            }
                            result[x][y] = sum;
                        }
                    }
                    return result;
                }

                function copyState(o1, o2) {
                    o2.fillStyle = o1.fillStyle;
                    o2.lineCap = o1.lineCap;
                    o2.lineJoin = o1.lineJoin;
                    o2.lineWidth = o1.lineWidth;
                    o2.miterLimit = o1.miterLimit;
                    o2.shadowBlur = o1.shadowBlur;
                    o2.shadowColor = o1.shadowColor;
                    o2.shadowOffsetX = o1.shadowOffsetX;
                    o2.shadowOffsetY = o1.shadowOffsetY;
                    o2.strokeStyle = o1.strokeStyle;
                    o2.globalAlpha = o1.globalAlpha;
                    o2.font = o1.font;
                    o2.textAlign = o1.textAlign;
                    o2.textBaseline = o1.textBaseline;
                    o2.arcScaleX_ = o1.arcScaleX_;
                    o2.arcScaleY_ = o1.arcScaleY_;
                    o2.lineScale_ = o1.lineScale_;
                }

                //可精简
                var colorData = {
                    aliceblue: '#F0F8FF',
                    antiquewhite: '#FAEBD7',
                    aquamarine: '#7FFFD4',
                    azure: '#F0FFFF',
                    beige: '#F5F5DC',
                    bisque: '#FFE4C4',
                    black: '#000000',
                    blanchedalmond: '#FFEBCD',
                    blueviolet: '#8A2BE2',
                    brown: '#A52A2A',
                    burlywood: '#DEB887',
                    cadetblue: '#5F9EA0',
                    chartreuse: '#7FFF00',
                    chocolate: '#D2691E',
                    coral: '#FF7F50',
                    cornflowerblue: '#6495ED',
                    cornsilk: '#FFF8DC',
                    crimson: '#DC143C',
                    cyan: '#00FFFF',
                    darkblue: '#00008B',
                    darkcyan: '#008B8B',
                    darkgoldenrod: '#B8860B',
                    darkgray: '#A9A9A9',
                    darkgreen: '#006400',
                    darkgrey: '#A9A9A9',
                    darkkhaki: '#BDB76B',
                    darkmagenta: '#8B008B',
                    darkolivegreen: '#556B2F',
                    darkorange: '#FF8C00',
                    darkorchid: '#9932CC',
                    darkred: '#8B0000',
                    darksalmon: '#E9967A',
                    darkseagreen: '#8FBC8F',
                    darkslateblue: '#483D8B',
                    darkslategray: '#2F4F4F',
                    darkslategrey: '#2F4F4F',
                    darkturquoise: '#00CED1',
                    darkviolet: '#9400D3',
                    deeppink: '#FF1493',
                    deepskyblue: '#00BFFF',
                    dimgray: '#696969',
                    dimgrey: '#696969',
                    dodgerblue: '#1E90FF',
                    firebrick: '#B22222',
                    floralwhite: '#FFFAF0',
                    forestgreen: '#228B22',
                    gainsboro: '#DCDCDC',
                    ghostwhite: '#F8F8FF',
                    gold: '#FFD700',
                    goldenrod: '#DAA520',
                    grey: '#808080',
                    greenyellow: '#ADFF2F',
                    honeydew: '#F0FFF0',
                    hotpink: '#FF69B4',
                    indianred: '#CD5C5C',
                    indigo: '#4B0082',
                    ivory: '#FFFFF0',
                    khaki: '#F0E68C',
                    lavender: '#E6E6FA',
                    lavenderblush: '#FFF0F5',
                    lawngreen: '#7CFC00',
                    lemonchiffon: '#FFFACD',
                    lightblue: '#ADD8E6',
                    lightcoral: '#F08080',
                    lightcyan: '#E0FFFF',
                    lightgoldenrodyellow: '#FAFAD2',
                    lightgreen: '#90EE90',
                    lightgrey: '#D3D3D3',
                    lightpink: '#FFB6C1',
                    lightsalmon: '#FFA07A',
                    lightseagreen: '#20B2AA',
                    lightskyblue: '#87CEFA',
                    lightslategray: '#778899',
                    lightslategrey: '#778899',
                    lightsteelblue: '#B0C4DE',
                    lightyellow: '#FFFFE0',
                    limegreen: '#32CD32',
                    linen: '#FAF0E6',
                    magenta: '#FF00FF',
                    mediumaquamarine: '#66CDAA',
                    mediumblue: '#0000CD',
                    mediumorchid: '#BA55D3',
                    mediumpurple: '#9370DB',
                    mediumseagreen: '#3CB371',
                    mediumslateblue: '#7B68EE',
                    mediumspringgreen: '#00FA9A',
                    mediumturquoise: '#48D1CC',
                    mediumvioletred: '#C71585',
                    midnightblue: '#191970',
                    mintcream: '#F5FFFA',
                    mistyrose: '#FFE4E1',
                    moccasin: '#FFE4B5',
                    navajowhite: '#FFDEAD',
                    oldlace: '#FDF5E6',
                    olivedrab: '#6B8E23',
                    orange: '#FFA500',
                    orangered: '#FF4500',
                    orchid: '#DA70D6',
                    palegoldenrod: '#EEE8AA',
                    palegreen: '#98FB98',
                    paleturquoise: '#AFEEEE',
                    palevioletred: '#DB7093',
                    papayawhip: '#FFEFD5',
                    peachpuff: '#FFDAB9',
                    peru: '#CD853F',
                    pink: '#FFC0CB',
                    plum: '#DDA0DD',
                    powderblue: '#B0E0E6',
                    rosybrown: '#BC8F8F',
                    royalblue: '#4169E1',
                    saddlebrown: '#8B4513',
                    salmon: '#FA8072',
                    sandybrown: '#F4A460',
                    seagreen: '#2E8B57',
                    seashell: '#FFF5EE',
                    sienna: '#A0522D',
                    skyblue: '#87CEEB',
                    slateblue: '#6A5ACD',
                    slategray: '#708090',
                    slategrey: '#708090',
                    snow: '#FFFAFA',
                    springgreen: '#00FF7F',
                    steelblue: '#4682B4',
                    tan: '#D2B48C',
                    thistle: '#D8BFD8',
                    tomato: '#FF6347',
                    turquoise: '#40E0D0',
                    violet: '#EE82EE',
                    wheat: '#F5DEB3',
                    whitesmoke: '#F5F5F5',
                    yellowgreen: '#9ACD32'
                };


                function getRgbHslContent(styleString) {
                    var start = styleString.indexOf('(', 3);
                    var end = styleString.indexOf(')', start + 1);
                    var parts = styleString.substring(start + 1, end).split(',');
                    // add alpha if needed
                    if (parts.length != 4 || styleString.charAt(3) != 'a') {
                        parts[3] = 1;
                    }
                    return parts;
                }

                function percent(s) {
                    return parseFloat(s) / 100;
                }

                function clamp(v, min, max) {
                    return Math.min(max, Math.max(min, v));
                }

                function hslToRgb(parts) {
                    var r, g, b, h, s, l;
                    h = parseFloat(parts[0]) / 360 % 360;
                    if (h < 0)
                        h++;
                    s = clamp(percent(parts[1]), 0, 1);
                    l = clamp(percent(parts[2]), 0, 1);
                    if (s == 0) {
                        r = g = b = l; // achromatic
                    } else {
                        var q = l < 0.5 ? l * (1 + s) : l + s - l * s;
                        var p = 2 * l - q;
                        r = hueToRgb(p, q, h + 1 / 3);
                        g = hueToRgb(p, q, h);
                        b = hueToRgb(p, q, h - 1 / 3);
                    }

                    return '#' + decToHex[Math.floor(r * 255)] +
                        decToHex[Math.floor(g * 255)] +
                        decToHex[Math.floor(b * 255)];
                }

                function hueToRgb(m1, m2, h) {
                    if (h < 0)
                        h++;
                    if (h > 1)
                        h--;

                    if (6 * h < 1)
                        return m1 + (m2 - m1) * 6 * h;
                    else if (2 * h < 1)
                        return m2;
                    else if (3 * h < 2)
                        return m1 + (m2 - m1) * (2 / 3 - h) * 6;
                    else
                        return m1;
                }

                var processStyleCache = {};

                function processStyle(styleString) {
                    if (styleString in processStyleCache) {
                        return processStyleCache[styleString];
                    }

                    var str, alpha = 1;

                    styleString = String(styleString);
                    if (styleString.charAt(0) == '#') {
                        str = styleString;
                    } else if (/^rgb/.test(styleString)) {
                        var parts = getRgbHslContent(styleString);
                        var str = '#', n;
                        for (var i = 0; i < 3; i++) {
                            if (parts[i].indexOf('%') != -1) {
                                n = Math.floor(percent(parts[i]) * 255);
                            } else {
                                n = +parts[i];
                            }
                            str += decToHex[clamp(n, 0, 255)];
                        }
                        alpha = +parts[3];
                    } else if (/^hsl/.test(styleString)) {
                        var parts = getRgbHslContent(styleString);
                        str = hslToRgb(parts);
                        alpha = parts[3];
                    } else {
                        str = colorData[styleString] || styleString;
                    }
                    return processStyleCache[styleString] = { color: str, alpha: alpha };
                }

                var DEFAULT_STYLE = {
                    style: 'normal',
                    variant: 'normal',
                    weight: 'normal',
                    size: 10,
                    family: 'sans-serif'
                };

                // Internal text style cache
                var fontStyleCache = {};

                function processFontStyle(styleString) {
                    if (fontStyleCache[styleString]) {
                        return fontStyleCache[styleString];
                    }

                    var el = document.createElement('div');
                    var style = el.style;
                    try {
                        style.font = styleString;
                    } catch (ex) {
                        // Ignore failures to set to invalid font.
                    }

                    return fontStyleCache[styleString] = {
                        style: style.fontStyle || DEFAULT_STYLE.style,
                        variant: style.fontVariant || DEFAULT_STYLE.variant,
                        weight: style.fontWeight || DEFAULT_STYLE.weight,
                        size: style.fontSize || DEFAULT_STYLE.size,
                        family: style.fontFamily || DEFAULT_STYLE.family
                    };
                }

                function getComputedStyle(style, element) {
                    var computedStyle = {};

                    for (var p in style) {
                        computedStyle[p] = style[p];
                    }

                    // Compute the size
                    var canvasFontSize = parseFloat(element.currentStyle.fontSize),
                        fontSize = parseFloat(style.size);

                    if (typeof style.size == 'number') {
                        computedStyle.size = style.size;
                    } else if (style.size.indexOf('px') != -1) {
                        computedStyle.size = fontSize;
                    } else if (style.size.indexOf('em') != -1) {
                        computedStyle.size = canvasFontSize * fontSize;
                    } else if (style.size.indexOf('%') != -1) {
                        computedStyle.size = (canvasFontSize / 100) * fontSize;
                    } else if (style.size.indexOf('pt') != -1) {
                        computedStyle.size = fontSize / .75;
                    } else {
                        computedStyle.size = canvasFontSize;
                    }

                    // Different scaling between normal text and VML text. This was found using
                    // trial and error to get the same size as non VML text.
                    computedStyle.size *= 0.981;

                    return computedStyle;
                }

                function buildStyle(style) {
                    return style.style + ' ' + style.variant + ' ' + style.weight + ' ' +
                        style.size + 'px ' + style.family;
                }

                var lineCapMap = {
                    'butt': 'flat',
                    'round': 'round'
                };

                function processLineCap(lineCap) {
                    return lineCapMap[lineCap] || 'square';
                }

                /**
                 * This class implements CanvasRenderingContext2D interface as described by
                 * the WHATWG.
                 * @param {HTMLElement} canvasElement The element that the 2D canvasContext should
                 * be associated with
                 */
                function CanvasRenderingContext2D_(canvasElement) {
                    this.m_ = createMatrixIdentity();

                    this.mStack_ = [];
                    this.aStack_ = [];
                    this.currentPath_ = [];

                    // Canvas canvasContext properties
                    this.strokeStyle = '#000';
                    this.fillStyle = '#000';

                    this.lineWidth = 1;
                    this.lineJoin = 'miter';
                    this.lineCap = 'butt';
                    this.miterLimit = Z * 1;
                    this.globalAlpha = 1;
                    this.font = '10px sans-serif';
                    this.textAlign = 'left';
                    this.textBaseline = 'alphabetic';
                    this.canvas = canvasElement;

                    var cssText = 'width:' + canvasElement.clientWidth + 'px;height:' +
                        canvasElement.clientHeight + 'px;overflow:hidden;position:absolute';
                    var el = canvasElement.ownerDocument.createElement('div');
                    el.style.cssText = cssText;
                    canvasElement.appendChild(el);

                    var overlayEl = el.cloneNode(false);
                    // Use a non transparent background.
                    overlayEl.style.backgroundColor = 'red';
                    overlayEl.style.filter = 'alpha(opacity=0)';
                    canvasElement.appendChild(overlayEl);

                    this.element_ = el;
                    this.arcScaleX_ = 1;
                    this.arcScaleY_ = 1;
                    this.lineScale_ = 1;
                }

                var contextPrototype = CanvasRenderingContext2D_.prototype;
                contextPrototype.clearRect = function () {
                    if (this.textMeasureEl_) {
                        this.textMeasureEl_.removeNode(true);
                        this.textMeasureEl_ = null;
                    }
                    this.element_.innerHTML = '';
                };

                contextPrototype.beginPath = function () {
                    // TODO: Branch current matrix so that save/restore has no effect
                    //       as per safari docs.
                    this.currentPath_ = [];
                };

                contextPrototype.moveTo = function (aX, aY) {
                    var p = getCoords(this, aX, aY);
                    this.currentPath_.push({ type: 'moveTo', x: p.x, y: p.y });
                    this.currentX_ = p.x;
                    this.currentY_ = p.y;
                };

                contextPrototype.lineTo = function (aX, aY) {
                    var p = getCoords(this, aX, aY);
                    this.currentPath_.push({ type: 'lineTo', x: p.x, y: p.y });

                    this.currentX_ = p.x;
                    this.currentY_ = p.y;
                };

                contextPrototype.bezierCurveTo = function (aCP1x, aCP1y,
                    aCP2x, aCP2y,
                    aX, aY) {
                    var p = getCoords(this, aX, aY);
                    var cp1 = getCoords(this, aCP1x, aCP1y);
                    var cp2 = getCoords(this, aCP2x, aCP2y);
                    bezierCurveTo(this, cp1, cp2, p);
                };

                // Helper function that takes the already fixed cordinates.
                function bezierCurveTo(self, cp1, cp2, p) {
                    self.currentPath_.push({
                        type: 'bezierCurveTo',
                        cp1x: cp1.x,
                        cp1y: cp1.y,
                        cp2x: cp2.x,
                        cp2y: cp2.y,
                        x: p.x,
                        y: p.y
                    });
                    self.currentX_ = p.x;
                    self.currentY_ = p.y;
                }

                contextPrototype.quadraticCurveTo = function (aCPx, aCPy, aX, aY) {
                    // the following is lifted almost directly from
                    // http://developer.mozilla.org/en/docs/Canvas_tutorial:Drawing_shapes

                    var cp = getCoords(this, aCPx, aCPy);
                    var p = getCoords(this, aX, aY);

                    var cp1 = {
                        x: this.currentX_ + 2.0 / 3.0 * (cp.x - this.currentX_),
                        y: this.currentY_ + 2.0 / 3.0 * (cp.y - this.currentY_)
                    };
                    var cp2 = {
                        x: cp1.x + (p.x - this.currentX_) / 3.0,
                        y: cp1.y + (p.y - this.currentY_) / 3.0
                    };

                    bezierCurveTo(this, cp1, cp2, p);
                };

                contextPrototype.strokeRect = function (aX, aY, aWidth, aHeight) {
                    var oldPath = this.currentPath_;
                    this.beginPath();

                    this.moveTo(aX, aY);
                    this.lineTo(aX + aWidth, aY);
                    this.lineTo(aX + aWidth, aY + aHeight);
                    this.lineTo(aX, aY + aHeight);
                    this.closePath();
                    this.stroke();

                    this.currentPath_ = oldPath;
                };

                contextPrototype.createLinearGradient = function (aX0, aY0, aX1, aY1) {
                    var gradient = new CanvasGradient_('gradient');
                    gradient.x0_ = aX0;
                    gradient.y0_ = aY0;
                    gradient.x1_ = aX1;
                    gradient.y1_ = aY1;
                    return gradient;
                };

                contextPrototype.createRadialGradient = function (aX0, aY0, aR0,
                    aX1, aY1, aR1) {
                    var gradient = new CanvasGradient_('gradientradial');
                    gradient.x0_ = aX0;
                    gradient.y0_ = aY0;
                    gradient.r0_ = aR0;
                    gradient.x1_ = aX1;
                    gradient.y1_ = aY1;
                    gradient.r1_ = aR1;
                    return gradient;
                };

                contextPrototype.stroke = function (aFill) {
                    var lineStr = [];
                    var lineOpen = false;

                    var W = 10;
                    var H = 10;

                    lineStr.push('<g_vml_:shape',
                        ' filled="', !!aFill, '"',
                        ' style="position:absolute;width:', W, 'px;height:', H, 'px;"',
                        ' coordorigin="0,0"',
                        ' coordsize="', Z * W, ',', Z * H, '"',
                        ' stroked="', !aFill, '"',
                        ' path="');

                    var newSeq = false;
                    var min = { x: null, y: null };
                    var max = { x: null, y: null };

                    for (var i = 0; i < this.currentPath_.length; i++) {
                        var p = this.currentPath_[i];
                        var c;

                        switch (p.type) {
                            case 'moveTo':
                                c = p;
                                lineStr.push(' m ', mr(p.x), ',', mr(p.y));
                                break;
                            case 'lineTo':
                                lineStr.push(' l ', mr(p.x), ',', mr(p.y));
                                break;
                            case 'close':
                                lineStr.push(' x ');
                                p = null;
                                break;
                            case 'bezierCurveTo':
                                lineStr.push(' c ',
                                    mr(p.cp1x), ',', mr(p.cp1y), ',',
                                    mr(p.cp2x), ',', mr(p.cp2y), ',',
                                    mr(p.x), ',', mr(p.y));
                                break;
                            case 'at':
                            case 'wa':
                                lineStr.push(' ', p.type, ' ',
                                    mr(p.x - this.arcScaleX_ * p.radius), ',',
                                    mr(p.y - this.arcScaleY_ * p.radius), ' ',
                                    mr(p.x + this.arcScaleX_ * p.radius), ',',
                                    mr(p.y + this.arcScaleY_ * p.radius), ' ',
                                    mr(p.xStart), ',', mr(p.yStart), ' ',
                                    mr(p.xEnd), ',', mr(p.yEnd));
                                break;
                        }


                        // TODO: Following is broken for curves due to
                        //       move to proper paths.

                        // Figure out dimensions so we can do gradient fills
                        // properly
                        if (p) {
                            if (min.x == null || p.x < min.x) {
                                min.x = p.x;
                            }
                            if (max.x == null || p.x > max.x) {
                                max.x = p.x;
                            }
                            if (min.y == null || p.y < min.y) {
                                min.y = p.y;
                            }
                            if (max.y == null || p.y > max.y) {
                                max.y = p.y;
                            }
                        }
                    }
                    lineStr.push(' ">');

                    if (!aFill) {
                        appendStroke(this, lineStr);
                    } else {
                        appendFill(this, lineStr, min, max);
                    }

                    lineStr.push('</g_vml_:shape>');

                    this.element_.insertAdjacentHTML('beforeEnd', lineStr.join(''));
                };

                function appendStroke(ctx, lineStr) {
                    var a = processStyle(ctx.strokeStyle);
                    var color = a.color;
                    var opacity = a.alpha * ctx.globalAlpha;
                    var lineWidth = ctx.lineScale_ * ctx.lineWidth;

                    // VML cannot correctly render a line if the width is less than 1px.
                    // In that case, we dilute the color to make the line look thinner.
                    if (lineWidth < 1) {
                        opacity *= lineWidth;
                    }

                    lineStr.push(
                        '<g_vml_:stroke',
                        ' opacity="', opacity, '"',
                        ' joinstyle="', ctx.lineJoin, '"',
                        ' miterlimit="', ctx.miterLimit, '"',
                        ' endcap="', processLineCap(ctx.lineCap), '"',
                        ' weight="', lineWidth, 'px"',
                        ' color="', color, '" />'
                    );
                }

                function appendFill(ctx, lineStr, min, max) {
                    var fillStyle = ctx.fillStyle;
                    var arcScaleX = ctx.arcScaleX_;
                    var arcScaleY = ctx.arcScaleY_;
                    var width = max.x - min.x;
                    var height = max.y - min.y;
                    if (fillStyle instanceof CanvasGradient_) {
                        // TODO: Gradients transformed with the transformation matrix.
                        var angle = 0;
                        var focus = { x: 0, y: 0 };

                        // additional offset
                        var shift = 0;
                        // scale factor for offset
                        var expansion = 1;

                        if (fillStyle.type_ == 'gradient') {
                            var x0 = fillStyle.x0_ / arcScaleX;
                            var y0 = fillStyle.y0_ / arcScaleY;
                            var x1 = fillStyle.x1_ / arcScaleX;
                            var y1 = fillStyle.y1_ / arcScaleY;
                            var p0 = getCoords(ctx, x0, y0);
                            var p1 = getCoords(ctx, x1, y1);
                            var dx = p1.x - p0.x;
                            var dy = p1.y - p0.y;
                            angle = Math.atan2(dx, dy) * 180 / Math.PI;

                            // The angle should be a non-negative number.
                            if (angle < 0) {
                                angle += 360;
                            }

                            // Very small angles produce an unexpected result because they are
                            // converted to a scientific notation string.
                            if (angle < 1e-6) {
                                angle = 0;
                            }
                        } else {
                            var p0 = getCoords(ctx, fillStyle.x0_, fillStyle.y0_);
                            focus = {
                                x: (p0.x - min.x) / width,
                                y: (p0.y - min.y) / height
                            };

                            width /= arcScaleX * Z;
                            height /= arcScaleY * Z;
                            var dimension = m.max(width, height);
                            shift = 2 * fillStyle.r0_ / dimension;
                            expansion = 2 * fillStyle.r1_ / dimension - shift;
                        }

                        // We need to sort the color stops in ascending order by offset,
                        // otherwise IE won't interpret it correctly.
                        var stops = fillStyle.colors_;
                        stops.sort(function (cs1, cs2) {
                            return cs1.offset - cs2.offset;
                        });

                        var length = stops.length;
                        var color1 = stops[0].color;
                        var color2 = stops[length - 1].color;
                        var opacity1 = stops[0].alpha * ctx.globalAlpha;
                        var opacity2 = stops[length - 1].alpha * ctx.globalAlpha;

                        var colors = [];
                        for (var i = 0; i < length; i++) {
                            var stop = stops[i];
                            colors.push(stop.offset * expansion + shift + ' ' + stop.color);
                        }

                        // When colors attribute is used, the meanings of opacity and o:opacity2
                        // are reversed.
                        lineStr.push('<g_vml_:fill type="', fillStyle.type_, '"',
                            ' method="none" focus="100%"',
                            ' color="', color1, '"',
                            ' color2="', color2, '"',
                            ' colors="', colors.join(','), '"',
                            ' opacity="', opacity2, '"',
                            ' g_o_:opacity2="', opacity1, '"',
                            ' angle="', angle, '"',
                            ' focusposition="', focus.x, ',', focus.y, '" />');
                    } else if (fillStyle instanceof CanvasPattern_) {
                        if (width && height) {
                            var deltaLeft = -min.x;
                            var deltaTop = -min.y;
                            lineStr.push('<g_vml_:fill',
                                ' position="',
                                deltaLeft / width * arcScaleX * arcScaleX, ',',
                                deltaTop / height * arcScaleY * arcScaleY, '"',
                                ' type="tile"',
                                // TODO: Figure out the correct size to fit the scale.
                                //' size="', w, 'px ', h, 'px"',
                                ' src="', fillStyle.src_, '" />');
                        }
                    } else {
                        var a = processStyle(ctx.fillStyle);
                        var color = a.color;
                        var opacity = a.alpha * ctx.globalAlpha;
                        lineStr.push('<g_vml_:fill color="', color, '" opacity="', opacity,
                            '" />');
                    }
                }

                contextPrototype.fill = function () {
                    this.stroke(true);
                };

                contextPrototype.closePath = function () {
                    this.currentPath_.push({ type: 'close' });
                };

                function getCoords(ctx, aX, aY) {
                    var m = ctx.m_;
                    return {
                        x: Z * (aX * m[0][0] + aY * m[1][0] + m[2][0]) - Z2,
                        y: Z * (aX * m[0][1] + aY * m[1][1] + m[2][1]) - Z2
                    };
                };

                contextPrototype.save = function () {
                    var o = {};
                    copyState(this, o);
                    this.aStack_.push(o);
                    this.mStack_.push(this.m_);
                    this.m_ = matrixMultiply(createMatrixIdentity(), this.m_);
                };

                contextPrototype.restore = function () {
                    if (this.aStack_.length) {
                        copyState(this.aStack_.pop(), this);
                        this.m_ = this.mStack_.pop();
                    }
                };

                function matrixIsFinite(m) {
                    return isFinite(m[0][0]) && isFinite(m[0][1]) &&
                        isFinite(m[1][0]) && isFinite(m[1][1]) &&
                        isFinite(m[2][0]) && isFinite(m[2][1]);
                }

                function setM(ctx, m, updateLineScale) {
                    if (!matrixIsFinite(m)) {
                        return;
                    }
                    ctx.m_ = m;

                    if (updateLineScale) {
                        // Get the line scale.
                        // Determinant of this.m_ means how much the area is enlarged by the
                        // transformation. So its square root can be used as a scale factor
                        // for width.
                        var det = m[0][0] * m[1][1] - m[0][1] * m[1][0];
                        ctx.lineScale_ = sqrt(abs(det));
                    }
                }

                contextPrototype.translate = function (aX, aY) {
                    var m1 = [
                        [1, 0, 0],
                        [0, 1, 0],
                        [aX, aY, 1]
                    ];

                    setM(this, matrixMultiply(m1, this.m_), false);
                };

                contextPrototype.rotate = function (aRot) {
                    var c = mc(aRot);
                    var s = ms(aRot);

                    var m1 = [
                        [c, s, 0],
                        [-s, c, 0],
                        [0, 0, 1]
                    ];

                    setM(this, matrixMultiply(m1, this.m_), false);
                };

                contextPrototype.scale = function (aX, aY) {
                    this.arcScaleX_ *= aX;
                    this.arcScaleY_ *= aY;
                    var m1 = [
                        [aX, 0, 0],
                        [0, aY, 0],
                        [0, 0, 1]
                    ];

                    setM(this, matrixMultiply(m1, this.m_), true);
                };

                contextPrototype.transform = function (m11, m12, m21, m22, dx, dy) {
                    var m1 = [
                        [m11, m12, 0],
                        [m21, m22, 0],
                        [dx, dy, 1]
                    ];

                    setM(this, matrixMultiply(m1, this.m_), true);
                };

                contextPrototype.setTransform = function (m11, m12, m21, m22, dx, dy) {
                    var m = [
                        [m11, m12, 0],
                        [m21, m22, 0],
                        [dx, dy, 1]
                    ];

                    setM(this, m, true);
                };

                /******** STUBS ********/
                contextPrototype.clip = function () {
                    // TODO: Implement
                };

                contextPrototype.arcTo = function () {
                    // TODO: Implement
                };

                contextPrototype.createPattern = function (image, repetition) {
                    return new CanvasPattern_(image, repetition);
                };

                // Gradient / Pattern Stubs
                function CanvasGradient_(aType) {
                    this.type_ = aType;
                    this.x0_ = 0;
                    this.y0_ = 0;
                    this.r0_ = 0;
                    this.x1_ = 0;
                    this.y1_ = 0;
                    this.r1_ = 0;
                    this.colors_ = [];
                }

                CanvasGradient_.prototype.addColorStop = function (aOffset, aColor) {
                    aColor = processStyle(aColor);
                    this.colors_.push({
                        offset: aOffset,
                        color: aColor.color,
                        alpha: aColor.alpha
                    });
                };

                function CanvasPattern_(image, repetition) {
                    assertImageIsValid(image);
                    switch (repetition) {
                        case 'repeat':
                        case null:
                        case '':
                            this.repetition_ = 'repeat';
                            break
                        case 'repeat-x':
                        case 'repeat-y':
                        case 'no-repeat':
                            this.repetition_ = repetition;
                            break;
                        default:
                            throwException('SYNTAX_ERR');
                    }

                    this.src_ = image.src;
                    this.width_ = image.width;
                    this.height_ = image.height;
                }

                function throwException(s) {
                    throw new DOMException_(s);
                }

                function assertImageIsValid(img) {
                    if (!img || img.nodeType != 1 || img.tagName != 'IMG') {
                        throwException('TYPE_MISMATCH_ERR');
                    }
                    if (img.readyState != 'complete') {
                        throwException('INVALID_STATE_ERR');
                    }
                }

                function DOMException_(s) {
                    this.code = this[s];
                    this.message = s + ': DOM Exception ' + this.code;
                }
                var p = DOMException_.prototype = new Error;
                p.INDEX_SIZE_ERR = 1;
                p.DOMSTRING_SIZE_ERR = 2;
                p.HIERARCHY_REQUEST_ERR = 3;
                p.WRONG_DOCUMENT_ERR = 4;
                p.INVALID_CHARACTER_ERR = 5;
                p.NO_DATA_ALLOWED_ERR = 6;
                p.NO_MODIFICATION_ALLOWED_ERR = 7;
                p.NOT_FOUND_ERR = 8;
                p.NOT_SUPPORTED_ERR = 9;
                p.INUSE_ATTRIBUTE_ERR = 10;
                p.INVALID_STATE_ERR = 11;
                p.SYNTAX_ERR = 12;
                p.INVALID_MODIFICATION_ERR = 13;
                p.NAMESPACE_ERR = 14;
                p.INVALID_ACCESS_ERR = 15;
                p.VALIDATION_ERR = 16;
                p.TYPE_MISMATCH_ERR = 17;

                // set up externs
                G_vmlCanvasManager = G_vmlCanvasManager_;
                CanvasRenderingContext2D = CanvasRenderingContext2D_;
                CanvasGradient = CanvasGradient_;
                CanvasPattern = CanvasPattern_;
                DOMException = DOMException_;
            })();
        }

    }
    //判断canvas的支持
    function _H5SupportTest() {
        isModernBrowser = !!document.createElement('canvas').getContext
    }
    //初始化数据
    function _initData() {
        _removeCanvas();
        requestAmount = 0;
        vaptchaCanvasDiv = null;
        canvasWidth = 0;
        canvasHeight = 0;
        clickX = new Array();
        clickY = new Array();
        clickDrag = new Array();
        dragSpotData = new Array();
        characteristics = new Array();
        startPoint = null;
        previousPoint = null;
        paint = false;
        allowLeftPaint = false;
        startTime = 0;
        canvas = null;
        canvasContext = null;
    }
    //添加鼠标事件
    function _addEventHandler() {
        //在mouseup事件后触发
        _eventHandler(document, "contextmenu", function (e) {
            //只有当鼠标右键按下，松开才会触发，因为是松开，所以这里需要初始化，
            _stopDefault(e);
            _initData();
        })

        // 鼠标进入图片改变鼠标样式
        var _img = document.getElementById('vaptchaImg');
        if (_img) {
            _eventHandler(_img, "mouseenter", function (e) {
                e.stopPropagation();//阻止事件冒泡
                if (!passcheck) {
                    _img.style.cursor = "url('images/pen.ico'),url('images/pen.png'),auto";
                }
            })
            _eventHandler(_img, "mousedown", function (e) {
                if (!passcheck) {
                    e.stopPropagation();//阻止事件冒泡
                    allowLeftPaint = true;
                    _prepareCanvas();
                    var coord = _getMouseCoords(e);
                    _addClick(coord.x, coord.y, false);
                    startTime = new Date().getTime();
                    startPoint = {
                        x: coord.x,
                        y: coord.y,
                        time: new Date().getTime()
                    }
                }
            })
        }


        //鼠标按下事件
        _eventHandler(document, "mousedown", function (e) {
            if (e.which != 1) { //当鼠标不是按下的左键时，应该不允许左键绘制。
                allowLeftPaint = false;
                _removeCanvas();
            }
            //在图片外点击左键无效果
            if (e.which == 1) {
                _removeCanvas();
                allowLeftPaint = false;
            }
            //点击左键绘制后，再次左键点击
            //排除其他鼠标按键
            var case0 = e.which == 3 || e.button == 2;


            var case1 = allowLeftPaint && e.which == 1;
            if ((case0 || case1) && challenge && !isValidated) {
                //当鼠标按下时，应该初始化坐标数组
                dragSpotData = new Array();
                //获取坐标
                var coord = _getMouseCoords(e);
                _addClick(coord.x, coord.y, false);
                startTime = new Date().getTime();
                startPoint = {
                    x: coord.x,
                    y: coord.y,
                    time: new Date().getTime()
                }
            }
        })

        //检测鼠标移动事件
        _eventHandler(document, "mousemove", function (e) {
            //判断之前是否按下了鼠标右键 or  按下了左键
            if (((e.which == 3) || (allowLeftPaint && e.which == 1)) && !isValidated) {
                var coord = _getMouseCoords(e);
                if (_overtimeCheck) {
                    _addDrapSpotData(coord);  //记录鼠标移动的数据
                }
            } else {
                return;
            }
        })

        //检测鼠标抬起
        _eventHandler(document, "mouseup", function (e) {
            //鼠标抬起清除画布

            _removeCanvas();

            //排除其他鼠标按键
            if ((e.which == 3 || e.button == 2 || (allowLeftPaint && e.which == 1)) && challenge && paint) {
                var data = _assemblyCoordDate(dragSpotData);
                //for ie test

                _calculateCharacteristic(dragSpotData);
                if (_analysisAngle() && _analysisAcceleration()) {
                    _validateVaptcha({
                        "v": data,
                        "siteid": siteId,
                        "challenge": challenge,
                    });
                }
                _initData();
                //解码检测：
                //_decodeCoords(data);
            }
        })
        var vaptchaImg = document.getElementById("vaptchaImg");
        if (vaptchaImg) {
            _eventHandler(vaptchaImg, "mouseover", function (e) {
                allowLeftPaint = true;
            })
        }
        //刷新
        var vaptchaRefresh = document.getElementById("vaptchaRefresh");
        if (vaptchaRefresh) {
            _eventHandler(vaptchaRefresh, "click", function (e) {
                console.log('init');
                VaptchaInitTime = _getCookie("VaptchaInitTime");
                if (VaptchaInitTime) {
                    if ((new Date().getTime() - VaptchaInitTime) < 180000) {
                        VaptchaInterval += (1 / ((new Date().getTime() - VaptchaTime) / 1000));
                    }
                }
                else {
                    VaptchaInterval = 0;
                    VaptchaInitTime = new Date().getTime();
                    _setCookie("VaptchaInitTime", VaptchaInitTime, 180000);
                }
                VaptchaTime = new Date().getTime();
                //todo:delete alert change to show tip on banner
                if (challenge && !(challenge.length)) {
                    var tipText = "验证流水号不能为空";
                    var messageHtml = "<i></i><span class=\"change text\">" + tipText + "</span>"
                    _changeBannerMessage(messageHtml, "class", "draw fail");
                } else if (requestAmount > 10) {
                    var tipText = "操作太频繁,请先绘制再重试";
                    var messageHtml = "<i></i><span class=\"change text\">" + tipText + "</span>"
                    _changeBannerMessage(messageHtml, "class", "draw fail");
                } else if (VaptchaInterval > 10) {
                    var tipText = "操作太频繁,请刷新页面重试";
                    var messageHtml = "<i></i><span class=\"change text\">" + tipText + "</span>"
                    _changeBannerMessage(messageHtml, "class", "draw fail");
                } else {
                    _refreshVaptcha(challenge, siteId);
                }
            })
        }

        //切换左键绘制（右键取消
    }


    /*生成Vaptcha*/
    //嵌入式
    function _generateEmbedVaptcha() {
        var fragment = document.createDocumentFragment();

        var navDiv = document.createElement("div");
        navDiv.setAttribute("id", "vaptchaNavDiv");
        navDiv.innerHTML = "\n<div id=\"vaptcha\" class=\"vaptcha\">\n\t<div class=\"vaptcha-main\">\n\t\t<img id='vaptchaImg' src=\"validate.fw.png\">\n\t\t<div id=\"vaptchaOpt\" class=\"opt\">\n\t\t\t<div class=\"logo-main\">\n\t\t\t\t<div class=\"vaptcha-logo\"></div>\n\t\t\t\t<span>aptcha</span>\n\t\t\t</div>\n\t\t\t<a id=\"vaptchaRefresh\" class=\"refresh\"></a>\n\t\t\t<div id=\"mouseImg\" class=\"draw\">\n\t\t\t\t<a class=\"mouse\"><img src=\"https://static.vaptcha.com/mouse.gif\" /></a>\n\t\t\t</div>\n\t\t</div>\n\t</div>\n</div>\n";
        fragment.appendChild(navDiv);

        var TempCanvasDiv = document.createElement("div");
        TempCanvasDiv.setAttribute("id", "vaptchaCanvasDiv");
        fragment.appendChild(TempCanvasDiv);
        var styleText = "\n@charset \"UTF-8\";\nbody {\n\tfont-family: \"Microsoft YaHei\", \u5FAE\u8F6F\u96C5\u9ED1, \"MicrosoftJhengHei\", \u534E\u6587\u7EC6\u9ED1, STHeiti, MingLiu;\n}\n\n#vaptchaNavDiv {\n\tposition: relative;\n}\n\n#vaptchaCanvasDiv {\n\tposition: fixed;\n\tleft: 0px;\n\ttop: 0px;\n\tz-index: 1000;\n}\n\n.vaptcha {\n\twidth: 378px;\n\toverflow: hidden;\n\t\n}\n.vaptcha-hide{\n\tdisplay: none;\n\topacity: 0;\n}\n.vaptcha-show{\n\tdisplay: block;\n\topacity: 1;\n}\n.vaptcha-animate{\n\tdisplay: block;\n\topacity: 1;\n\ttransition: opacity 0.4s ; \n\t-webkit-transition: opacity 0.4s ; \n}\n.vaptcha .vaptcha-main {\n\twidth: 100%;\n\theight: 172px;\n\tposition: relative;\n}\n.vaptcha .vaptcha-main .opt {\n\twidth: 100%;\n\theight: 30px;\n\tbackground: rgba(255, 255, 255, 0.9);\n\tline-height: 30px;\n\tposition: absolute;\n\ttop: 0;\n\ttransition: top 0.3s ; \n\t-webkit-transition: top 0.3s ; \n}\n\n.vaptcha .vaptcha-main .opt a:hover {\n\tcursor: pointer;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main {\n\tfloat: left;\n\tfont-size: 14px;\n\tmargin: 0 0 0 6px;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main .vaptcha-logo {\n\tdisplay: inline-block;\n\twidth: 14px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 0;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main span {\n\tmargin-left: -4px;\n}\n\n.vaptcha .vaptcha-main .opt .refresh {\n\tfloat: right;\n\tmargin: 5px 6px 0 0;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 -55px;\n}\n\n.vaptcha .vaptcha-main .opt .refresh:hover {\n\tbackground-position: 0 -71px;\n}\n\n.vaptcha .vaptcha-main .opt .draw {\n\twidth: 200px;\n\tmargin: 0 auto;\n\ttext-align: center;\n\tposition: relative;\n}\n.vaptcha .vaptcha-main .opt .draw span{\n\tfont-size: 12px;\n\tcolor: #D80000;\n}\n.vaptcha .vaptcha-main .opt .draw .mouse{\n\tdisplay: inline-block;\n\tmargin: 2px 0 0 0;\n}\n.vaptcha .vaptcha-main .opt .pass,.vaptcha .vaptcha-main .opt .fail{\n\tfloat: right;\n\tmargin-right: 10px;\n}\n/*.vaptcha .vaptcha-main .opt .draw i{\n\tdisplay: inline-block;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tposition: absolute;\n\ttop: 7px;\n}*/\n/*.vaptcha .vaptcha-main .opt .draw.pass i{\n\tbackground-position: 0 -91px;\n}\n.vaptcha .vaptcha-main .opt .draw.fail i{\n\tbackground-position: 0 -109px;\n}*/\n.vaptcha .vaptcha-main .opt .pass span,.vaptcha .vaptcha-main .opt .fail span{\n\tfont-size: 12px;\n\tcolor: #666;\n\tmargin: 0;\n}\n.vaptcha .vaptcha-main .opt .fail span{\n\tcolor: #262633;\n}\n.vaptcha .vaptcha-main .opt .pass span.change.text,.vaptcha .vaptcha-main .opt .fail span.change.text{\n\tfont-size: 14px;\n\tmargin-left: 20px;\n}\n.vaptcha .vaptcha-main .opt .pass span.change{\n\tcolor: #00D8A3;\n}\n.vaptcha .vaptcha-main .opt .fail span.change{\n\tcolor: #D80000;\n}\n.vaptcha .vaptcha-main .opt .pass span.dark,.vaptcha .vaptcha-main .opt .fail span.dark{\n\tcolor: #262633;\n}\n.btn-vaptcha{\n\tposition: relative;\n\twidth: 380px;\n\tline-height: 40px;\n\tbackground: #272734;\n\tborder: none;\n\tcolor: #fff;\n\ttext-align: left;\n\ttext-indent: 32px;\n}\n.btn-vaptcha:before,.btn-vaptcha:after{\n\tcontent: '';\n\tposition: absolute;\n\ttop: 10px;\n\tdisplay: inline-block;\n\twidth: 18px;\n\theight: 20px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\t\n}\n.btn-vaptcha:before{background-position: 0 -235px;left: 10px;}\n.btn-vaptcha:after{background-position: 0 -280px;right: 10px;}\n.btn-vaptcha.pass{color:#00D9A3;}\n.btn-vaptcha.pass:before{background-position: 0 -259px;}\n\n";
        _insertStyle(styleText)


        // var poz = document.getElementById("vaptchaPoz");
        if (!vaptchaContainer) {
            initFailCallback();
            return;
        }
        var poz = vaptchaContainer;
        //var tokenInput = document.createElement("input");
        //tokenInput.setAttribute("class",)
        poz.innerHTML = "";
        poz.appendChild(fragment);
        getReadyCallback();
    }
    //浮动式
    function _generateFloatVaptcha() {
        var fragment = document.createDocumentFragment();

        var navDiv = document.createElement("div");
        navDiv.setAttribute("id", "vaptchaNavDiv");
        navDiv.innerHTML = "<div id=\"vaptcha\" class=\"vaptcha\">\n<div class=\"vaptcha-main\">\n<img id='vaptchaImg'>\n<div id=\"vaptchaOpt\" class=\"opt\">\n<div class=\"logo-main\">\n<div class=\"vaptcha-logo\"></div>\n<span>aptcha</span>\n</div>\n<a id='vaptchaRefresh' class=\"reload\">\n<i class=\"refresh\"></i>\n</a>\n<div class=\"draw\">\n<a id=\"vaptchaLeftClick\" class=\"pencil\"></a>\n<span>or</span>\n<a class=\"mouse\"><img src=\"https://static.vaptcha.com/mouse.gif\" /></a>\n";
        fragment.appendChild(navDiv);

        var TempCanvasDiv = document.createElement("div");
        TempCanvasDiv.setAttribute("id", "vaptchaCanvasDiv");
        fragment.appendChild(TempCanvasDiv);
        var styleText = "@charset \"UTF-8\";\nbody {\n\tfont-family: \"Microsoft YaHei\", \u5FAE\u8F6F\u96C5\u9ED1, \"MicrosoftJhengHei\", \u534E\u6587\u7EC6\u9ED1, STHeiti, MingLiu;\n}\n\n#vaptchaPoz {\n\tmargin: 300px;\n}\n\n#vaptchaNavDiv {\n\tposition: relative;\n\t/*border: 1px solid #ccc;*/\n\t/*cursor: url(images/pen.ico), url(images/pen.png), auto;*/\n}\n\n#vaptchaCanvasDiv {\n\tposition: fixed;\n\tleft: 0px;\n\ttop: 0px;\n\tz-index: 1000;\n}\n\n.vaptcha {\n\twidth: 378px;\n\toverflow: hidden;\n\tposition: absolute;\n\ttop: -172px;\n\t\n}\n.vaptcha-hide{\n\tdisplay: none;\n\topacity: 0;\n}\n.vaptcha-show{\n\tdisplay: block;\n\topacity: 1;\n}\n.vaptcha-animate{\n\tdisplay: block;\n\topacity: 1;\n\ttransition: opacity 0.4s ; \n\t-webkit-transition: opacity 0.4s ; \n}\n\n.vaptcha:hover .vaptcha-main .opt{\n\ttop: 0;\n}\n.vaptcha .vaptcha-main {\n\twidth: 100%;\n\theight: 172px;\n\tposition: relative;\n}\n\n.vaptcha .vaptcha-main .opt {\n\twidth: 100%;\n\theight: 30px;\n\tbackground: rgba(255, 255, 255, 0.9);\n\tline-height: 30px;\n\tposition: absolute;\n\ttop: -30px;\n\ttransition: top 0.3s ; \n\t-webkit-transition: top 0.3s ; \n}\n\n.vaptcha .vaptcha-main .opt a:hover {\n\tcursor: pointer;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main {\n\tfloat: left;\n\tfont-size: 14px;\n\tpadding-left: 6px;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main .vaptcha-logo {\n\tdisplay: inline-block;\n\twidth: 14px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 0;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main span {\n\tmargin-left: -4px;\n}\n\n.vaptcha .vaptcha-main .opt .reload {\n\tfloat: right;\n\tpadding-right: 6px;\n}\n\n.vaptcha .vaptcha-main .opt .reload .refresh {\n\tdisplay: inline-block;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 -55px;\n}\n\n.vaptcha .vaptcha-main .opt .reload .refresh:hover {\n\tbackground-position: 0 -71px;\n}\n\n.vaptcha .vaptcha-main .opt .draw {\n\twidth: 200px;\n\tmargin: 0 auto;\n\ttext-align: center;\n\tposition: relative;\n}\n\n.vaptcha .vaptcha-main .opt .draw .pencil {\n\tdisplay: inline-block;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 -19px;\n}\n\n.vaptcha .vaptcha-main .opt .draw .pencil:hover {\n\tbackground-position: 0 -38px;\n}\n\n.vaptcha .vaptcha-main .opt .draw span {\n\tfont-size: 14px;\n\tcolor: #aaa;\n\tmargin: 0 10px;\n}\n\n.vaptcha .vaptcha-main .opt .draw .mouse img {\n\tvertical-align: text-bottom;\n}\n\n.vaptcha .vaptcha-main .opt .draw i{\n\tdisplay: inline-block;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tposition: absolute;\n\ttop: 7px;\n}\n.vaptcha .vaptcha-main .opt .draw.pass i{\n\tbackground-position: 0 -91px;\n}\n.vaptcha .vaptcha-main .opt .draw.fail i{\n\tbackground-position: 0 -109px;\n}\n.vaptcha .vaptcha-main .opt .draw.pass span,.vaptcha .vaptcha-main .opt .draw.fail span{\n\tfont-size: 12px;\n\tcolor: #878787;\n\tmargin: 0;\n}\n.vaptcha .vaptcha-main .opt .draw.pass span.change.text,.vaptcha .vaptcha-main .opt .draw.fail span.change.text{\n\tfont-size: 14px;\n\tmargin-left: 20px;\n}\n.vaptcha .vaptcha-main .opt .draw.pass span.change{\n\tcolor: #00D8A3;\n}\n.vaptcha .vaptcha-main .opt .draw.fail span.change{\n\tcolor: #D90000;\n}";
        _insertStyle(styleText)


        if (!vaptchaContainer) {
            initFailCallback();
            return;
        }
        var poz = vaptchaContainer;
        poz.innerHTML = "";
        var switchVaptcha = document.createElement("span");
        switchVaptcha.innerHTML = "点击隐藏验证图片";
        switchVaptcha.style.cssText = "width:496px;border:solid 1px #c4c3c3;background-color:#f9f9f9;color:#333;padding:11px 18px; border-radius:4px;font-size:2em;text-align:center;margin:20px 0px";
        poz.appendChild(fragment);
        poz.appendChild(switchVaptcha);

        _eventHandler(switchVaptcha, "click", function () {
            if (navDiv.style.display != "none") {
                navDiv.style.display = "none";
                switchVaptcha.innerHTML = "点击显示验证图片";
            } else {
                navDiv.style.display = "block";
                switchVaptcha.innerHTML = "点击隐藏验证图片";
            }
        })
        getReadyCallback();
    };
    //弹出式
    function _generatePopupVaptcha() {
        var fragment = document.createDocumentFragment();

        var navDiv = document.createElement("div");
        navDiv.setAttribute("id", "vaptchaNavDiv");
        navDiv.innerHTML = "<div class=\"vaptcha-mask\"></div><div id=\"vaptcha\" class=\"vaptcha\">\n<div class=\"vaptcha-main\">\n<img id='vaptchaImg'>\n<div id=\"vaptchaOpt\" class=\"opt\">\n<div class=\"logo-main\">\n<div class=\"vaptcha-logo\"></div>\n<span>aptcha</span>\n</div>\n<a id='vaptchaRefresh' class=\"reload\">\n<i class=\"refresh\"></i>\n</a>\n<div class=\"draw\">\n<a id=\"vaptchaLeftClick\" class=\"pencil\"></a>\n<span>or</span>\n<a class=\"mouse\"><img src=\"https://static.vaptcha.com/mouse.gif\" /></a>\n";
        fragment.appendChild(navDiv);

        var TempCanvasDiv = document.createElement("div");
        TempCanvasDiv.setAttribute("id", "vaptchaCanvasDiv");
        fragment.appendChild(TempCanvasDiv);
        var styleText = "\n@charset \"UTF-8\";\nbody {\n\tfont-family: \"Microsoft YaHei\", \u5FAE\u8F6F\u96C5\u9ED1, \"MicrosoftJhengHei\", \u534E\u6587\u7EC6\u9ED1, STHeiti, MingLiu;\n}\n\n#vaptchaPoz {\n}\n\n.vaptcha-mask{\n\tposition: fixed;\n    width: 100%;\n    height: 100%;\n    top: 0;\n    left: 0;\n    background-color: black;\n    opacity: 0.6;\n    filter: alpha(opacity=60);\n}\n\n#vaptchaNavDiv {\n\tmargin-left: 50%;\n\tmargin-top: 50%;\n\tposition: fixed;\n\tleft: -189px;\n\ttop: -98px;\n\t/*border: 1px solid #ccc;*/\n\t/*cursor: url(images/pen.ico), url(images/pen.png), auto;*/\n}\n\n#vaptchaCanvasDiv {\n\tposition: fixed;\n\tleft: 0px;\n\ttop: 0px;\n\tz-index: 1000;\n}\n\n.vaptcha {\n\twidth: 378px;\n\toverflow: hidden;\n}\n\n.vaptcha:hover .vaptcha-main .opt{\n\t/*opacity: 1;*/\n\ttop: 0;\n}\n.vaptcha .vaptcha-main {\n\twidth: 100%;\n\theight: 172px;\n\tposition: relative;\n}\n\n.vaptcha .vaptcha-main .opt {\n\twidth: 100%;\n\theight: 30px;\n\tbackground: rgba(255, 255, 255, 0.9);\n\tline-height: 30px;\n\tposition: absolute;\n\ttop: -30px;\n\t/*opacity: 0;*/\n\ttransition: top 0.3s ; \n\t-webkit-transition: top 0.3s ; \n}\n\n\n.vaptcha .vaptcha-main .opt a:hover {\n\tcursor: pointer;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main {\n\tfloat: left;\n\tfont-size: 14px;\n\tpadding-left: 6px;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main .vaptcha-logo {\n\tdisplay: inline-block;\n\twidth: 14px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 0;\n}\n\n.vaptcha .vaptcha-main .opt .logo-main span {\n\tmargin-left: -4px;\n}\n\n.vaptcha .vaptcha-main .opt .reload {\n\tfloat: right;\n\tpadding-right: 6px;\n}\n\n.vaptcha .vaptcha-main .opt .reload .refresh {\n\tdisplay: inline-block;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 -55px;\n}\n\n.vaptcha .vaptcha-main .opt .reload .refresh:hover {\n\tbackground-position: 0 -71px;\n}\n\n.vaptcha .vaptcha-main .opt .draw {\n\twidth: 200px;\n\tmargin: 0 auto;\n\ttext-align: center;\n\tposition: relative;\n}\n\n.vaptcha .vaptcha-main .opt .draw .pencil {\n\tdisplay: inline-block;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tbackground-position: 0 -19px;\n}\n\n.vaptcha .vaptcha-main .opt .draw .pencil:hover {\n\tbackground-position: 0 -38px;\n}\n\n.vaptcha .vaptcha-main .opt .draw span {\n\tfont-size: 14px;\n\tcolor: #aaa;\n\tmargin: 0 10px;\n}\n\n.vaptcha .vaptcha-main .opt .draw .mouse img {\n\tvertical-align: text-bottom;\n}\n\n.vaptcha .vaptcha-main .opt .draw i{\n\tdisplay: inline-block;\n\twidth: 16px;\n\theight: 16px;\n\tbackground: url(https://static.vaptcha.com/validate.png) no-repeat;\n\tposition: absolute;\n\ttop: 7px;\n}\n.vaptcha .vaptcha-main .opt .draw.pass i{\n\tbackground-position: 0 -91px;\n}\n.vaptcha .vaptcha-main .opt .draw.fail i{\n\tbackground-position: 0 -109px;\n}\n.vaptcha .vaptcha-main .opt .draw.pass span,.vaptcha .vaptcha-main .opt .draw.fail span{\n\tfont-size: 12px;\n\tcolor: #878787;\n\tmargin: 0;\n}\n.vaptcha .vaptcha-main .opt .draw.pass span.change.text,.vaptcha .vaptcha-main .opt .draw.fail span.change.text{\n\tfont-size: 14px;\n\tmargin-left: 20px;\n}\n.vaptcha .vaptcha-main .opt .draw.pass span.change{\n\tcolor: #00D8A3;\n}\n.vaptcha .vaptcha-main .opt .draw.fail span.change{\n\tcolor: #D90000;\n}";
        _insertStyle(styleText)


        if (!vaptchaContainer) {
            initFailCallback();
            return;
        }
        var poz = vaptchaContainer;
        poz.innerHTML = "";
        var switchVaptcha = document.createElement("span");
        switchVaptcha.innerHTML = "点击隐藏验证图片";
        switchVaptcha.style.cssText = "width:496px;border:solid 1px #c4c3c3;background-color:#f9f9f9;color:#333;padding:11px 18px; border-radius:4px;font-size:2em;text-align:center;margin:20px 0px";
        poz.appendChild(fragment);
        poz.appendChild(switchVaptcha);

        _eventHandler(switchVaptcha, "click", function () {
            if (navDiv.style.display != "none") {
                navDiv.style.display = "none";
                switchVaptcha.innerHTML = "点击显示验证图片";
            } else {
                navDiv.style.display = "block";
                switchVaptcha.innerHTML = "点击隐藏验证图片";
            }
        })
        getReadyCallback();
    }


    /*与vaptcha服务器交互*/
    function _detectResponse(data) {
        var vaptchaMessageDiv = document.createElement("div");
        vaptchaMessageDiv.setAttribute("class", "draw fail")
        var vaptchaOpt = document.getElementById("vaptchaOpt");
        var refreshObj = document.getElementById("vaptchaRefresh");
        switch (data.code) {
            case 1:
                vaptchaOpt.removeChild(vaptchaOpt.lastChild);
                vaptchaMessageDiv.innerHTML = "<i></i><span class=\"change text\">访问被拒绝</span>";
                vaptchaOpt.appendChild(vaptchaMessageDiv);
                validateErrorCallback();
                return false;
            case 2:
                if (refreshObj) {
                    _fireEvent("MouseEvents", "click", refreshObj);
                };
                validateErrorCallback();
                return false;
            case 5:
                vaptchaOpt.removeChild(vaptchaOpt.lastChild);
                vaptchaOpt.removeChild(refreshObj);
                vaptchaMessageDiv.innerHTML = "<i></i><span class=\"change text\">刷新太快</span>";
                vaptchaOpt.appendChild(vaptchaMessageDiv);
                validateErrorCallback();
                return false;
            case 6:
                vaptchaOpt.removeChild(vaptchaOpt.lastChild);
                vaptchaOpt.removeChild(refreshObj);
                vaptchaMessageDiv.innerHTML = "<i></i><span class=\"change text\">刷新太频繁</span>";
                vaptchaOpt.appendChild(vaptchaMessageDiv);
                validateErrorCallback();
                return false;
            case 7:
                vaptchaMessageDiv.innerHTML = "绘制太频繁";
                if (refreshObj) {
                    _fireEvent("MouseEvents", "click", refreshObj);
                };
                validateErrorCallback();
                return false;
            default:
                return true;
        }
    }
    //向Vaptcha验证数据
    function _validateVaptcha(data) {
        //vaptchaUrl + "vaptcha?callback=Vaptcha" + new Date().getTime()
        _getJsonp(protocol + vaptchaUrl + "validate?callback=Vaptcha" + new Date().getTime(), data, "callback", function (result) {
            // var vaptchaMessageDiv = document.createElement("div");
            if (result.code == 4) {
                var messageHtml = "<i></i><span class=\"change text\">验证未通过 </span><span class=\"change\">" + result.similarity + "</span><span> 匹配</span>"
                _changeBannerMessage(messageHtml, "class", "draw fail");
                validateFailCallback();
            } else if (_detectResponse(result)) {
                var img = document.getElementById("vaptchaImg");
                img.setAttribute("src", protocol + imgUrl + VaptchaData.coverimg);
                if (vaptchaOpt) {
                    var messageHtml = "<i></i><span class=\"change text\">验证通过 </span><span class=\"change\">" + result.similarity + "</span><span> 匹配</span>";
                    _changeBannerMessage(messageHtml, "class", "draw pass");
                    //remove refresh <a></a> label.
                    var refreshBtn = document.getElementById("vaptchaRefresh");
                    refreshBtn && vaptchaOpt.removeChild(refreshBtn);
                }
                isValidated = true;
                passcheck = true;
                document.getElementById('vaptchaImg').style.cursor = "auto";
                //todo:赋值
                result.token = "test token";
                validateSuccessCallback(result.token);
            }
        })
    }
    //生成Vaptcha
    function _generateVaptchaImg() {
        switch (vaptchaProduct) {
            case "float": _generateFloatVaptcha();
                break;
            case "embed": _generateEmbedVaptcha();
                break;
            case "popup": _generatePopupVaptcha();
                break;
            default: _generateEmbedVaptcha();
                break;
        }
    }
    //向Vaptcha请求图片
    function _getVaptcha(challengeParam, siteIdParam) {
        requestAmount += 1;
        challenge = challengeParam;
        siteId = siteIdParam;
        //vaptchaUrl + "/get?callback=Vaptcha" + new Date().getTime()
        _getJsonp(protocol + vaptchaUrl + "get?callback=Vaptcha" + new Date().getTime(), { "challenge": challengeParam, "siteId": siteIdParam }, "callback", function (data) {
            if (_detectResponse(data)) {
                VaptchaData = data;
                var img = document.getElementById("vaptchaImg");
                img.setAttribute("src", protocol + imgUrl + VaptchaData.img);
                _addEventHandler();
                _refreshVaptchaOvertime();
            }
        })
    }
    //向Vaptcha刷新图片
    function _refreshVaptcha(challengeParam, siteIdParam) {
        requestAmount += 1;
        //vaptchaUrl + "/refresh?callback=Vaptcha" + new Date().getTime()
        _getJsonp(protocol + vaptchaUrl + "refresh?callback=Vaptcha" + new Date().getTime(), { "challenge": challengeParam, "siteId": siteIdParam }, "callback", function (data) {
            if (_detectResponse(data)) {
                VaptchaData = data;
                var img = document.getElementById("vaptchaImg")
                img.setAttribute("src", protocol + imgUrl + VaptchaData.img);
                getReadyCallback();
            }
        })
    }
    //超时自动刷新图片
    function _refreshVaptchaOvertime() {
        if (!isValidated) {
            interval = window.setInterval(() => {
                var refreshObj = document.getElementById("vaptchaRefresh");
                if (refreshObj) {
                    _fireEvent('MouseEvents', 'click', refreshObj);
                }
                if (isValidated) {
                    window.clearInterval(interval);
                }
            }, 180000);
        }

    }
    function _changeBannerMessage(message, attributeName, attributeValue) {
        var vaptchaMessageDiv = document.createElement("div");
        var vaptchaOpt = document.getElementById("vaptchaOpt");
        if (vaptchaOpt) {
            var mouseImg = document.getElementById("mouseImg");
            mouseImg && vaptchaOpt.removeChild(mouseImg);
            var lastNode = vaptchaOpt.lastChild;
            lastNode && vaptchaOpt.removeChild(lastNode);
            if (attributeName || attributeValue) {
                vaptchaMessageDiv.setAttribute(attributeName, attributeValue);
            }
            vaptchaMessageDiv.innerHTML = message;
            vaptchaOpt.appendChild(vaptchaMessageDiv);
        }
    }
    /*外部接口对象*/
    vaptcha = {
        init: function (container, options) {

            if (container) {
                vaptchaContainer = container;
            }
            else {
                if (typeof options.onInitFail === "function") {
                    initFailCallback = options.onInitFail;
                    initFailCallback();
                }
                return;
            }
            _H5SupportTest();
            _polyfill();

            if (options.ishttps) {
                protocol = "https://";
            }
            else {
                protocol = "http://";
            }
            if (typeof options.onReady === "function") {
                getReadyCallback = options.onReady;
            }
            if (typeof options.onRefresh === "function") {
                refreshCallback = options.onRefresh;
            }
            if (typeof options.onSuccess === "function") {
                validateSuccessCallback = options.onSuccess;
            }
            if (typeof options.onFail === "function") {
                validateFailCallback = options.onFail;
            }
            if (typeof options.onError === "function") {
                validateErrorCallback = options.onError;
            }
            if (options.product === "embed" || options.product === "float" || options.product === "popup") {
                vaptchaProduct = options.product;
            }
            if (options.siteId && options.challenge) {
                VaptchaInitTime = new Date().getTime();
                VaptchaTime = new Date().getTime();
                VaptchaInterval = 0;
                _setCookie("VaptchaInitTime", VaptchaInitTime, 180000);
                _generateVaptchaImg();
                _getVaptcha(options.challenge, options.siteId);
            } else {
                //do nothing
            }
        },
    };

    window.vaptcha = vaptcha;
})(window);