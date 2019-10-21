/*
 * @Author: luo 
 * @Date: 2017-07-22 16:06:24 
 * @Last Modified by: luo
 * @Last Modified time: 2017-07-24 15:42:24
 */
(function (window,factory) {
    var document = window.document;
    var ieStyles = [],isIE;
    var _clone = {
        dupNode: null,
        parentNode: null,
        init: init,
        hasChildNodes: null,
        destroy: destroy,
        style: "\n.vaptcha-loading {\n  width: 262px;\n  height: 44px;\n  display: -ms-flexbox;\n  display: flex;\n  -ms-flex-pack: center;\n  justify-content: center;\n  -ms-flex-align: center;\n  align-items: center;\n}\n\n.vaptcha-loading div {\n  width: 18px;\n  height: 18px;\n  background-color: #347eff;\n  border-radius: 50%;\n  display: inline-block;\n  animation: bounceDelay 1s infinite ease-in-out;\n  animation-fill-mode: both;\n  margin-left: 6px;\n}\n\n.vaptcha-loading div.happy1 {\n  animation-delay: -0.32s;\n}\n\n.vaptcha-loading div.happy2 {\n  animation-delay: -0.16s;\n}\n\n@keyframes bounceDelay {\n  0%,\n  80%,\n  to {\n    transform: scale(0);\n  }\n  40% {\n    transform: scale(1);\n  }\n}\n\n@-moz-keyframes bounceDelay {\n  0%,\n  80%,\n  to {\n    -moz-transform: scale(0);\n  }\n  40% {\n    -moz-transform: scale(1);\n  }\n}\n\n@-webkit-keyframes bounceDelay {\n  0%,\n  80%,\n  to {\n    -webkit-transform: scale(0);\n  }\n  40% {\n    -webkit-transform: scale(1);\n  }\n}\n\n.vaptcha-loading-content {\n  position: relative;\n  left: 50%;\n  top: 50%;\n  margin-left: -131px;\n  margin-top: -22px;\n}\n",
        content: " <div class=\"vaptcha-loading\">\n            <div class=\"happy1\"></div>\n            <div class=\"happy2\"></div>\n            <div class=\"happy3\"></div>\n        </div>"
    }
    var _source;
    var helper = {
        getParent: function (target) {
            var parent = target.parentElement;
            if (parent || target.parentNode) {
                if (!parent) {
                    parent = target.parentNode;
                }
            }
            return parent;
        },
        setStyle: function (ele, prop, value) {
            ele.style[prop] = value;
        },
        removeSelf: function (target) {
            this.getParent(target).removeChild(target);
        },
        insertStyle: function (styleText, nameAttribute) {
            if (nameAttribute && this.judgeExistStyle(styleText, nameAttribute)) {
                return;
            };
            if (document.createStyleSheet) {
                var cssStyle = document.createStyleSheet();
                cssStyle.cssText = styleText;
                ieStyles.push(nameAttribute);
                return;
            }
            var head = document.getElementsByTagName('head')[0];
            var style = document.createElement('style');
            // style.innerText = styleText;
            nameAttribute && style.setAttribute("name", nameAttribute);
            head.appendChild(style);
            style.innerText = styleText;
        },
        judgeExistStyle: function (styleText, nameAttribute) {
            if (this.judgeIE()) {
                return ieStyles.indexOf(nameAttribute) >= 0;
            }
            var head = document.getElementsByTagName('head')[0];
            var styles = document.getElementsByTagName("style");
            return styles[nameAttribute] ? true : false;
        },
        judgeIE: function () {
            var rMsie = /(msie\s|trident.*rv:)([\w.]+)/;
            var ua = navigator.userAgent.toLowerCase();
            var match = rMsie.exec(ua);
            if (match != null) {
                isIE = true;
                return true;
            } else {
                isIE = false;
                return false;
            }
        }
    }
    function init(source,child) {
        _source = source;
        var goal = document.createElement("div");
        goal.setAttribute("class", "vaptcha-loading-content");
        goal.innerHTML = _clone.content;
        helper.insertStyle(_clone.style, "v_animation");
        child ? _clone.hasChildNodes = true : _clone.hasChildNodes = false;
        _clone.dupNode = source.cloneNode(true);
        _clone.parentNode = helper.getParent(source);
        helper.setStyle(source, "display", "none");
        // _clone.parentNode.appendChild(_clone.dupNode);
        _clone.dupNode.innerHTML = "";
        if (_clone.hasChildNodes) {
            var parent = _clone.dupNode;
            var _child;
            for (var i = parent.childNodes.length - 1; i >= 0; i--) {
                if (parent.childNodes[i].isEqualNode(child)) {
                    _child = parent.childNodes[i];//找到副本中的child node
                    break;
                }
            }
            _clone.dupNode.insertBefore(goal, _child);
            helper.setStyle(_child, "display", "none");
        } else {
            _clone.dupNode.innerHTML = "";
            _clone.dupNode.appendChild(goal);
        }
        _clone.parentNode.insertBefore(_clone.dupNode, source);
    }
    function destroy() {
        helper.removeSelf(_clone.dupNode);
        helper.setStyle(_source, "display", "");
    }
    factory.clone = _clone;
})(window,factory);