(function() {
	var g = void 0,
		h = !0,
		k = null,
		l = !1,
		m, n = this;

	function p(a) {
		a = a.split(".");
		for(var b = n, c; c = a.shift();)
			if(b[c] != k) b = b[c];
			else return k;
		return b
	}

	function q(a) {
		return "string" == typeof a
	}

	function r(a, b) {
		var c = a.split("."),
			d = n;
		!(c[0] in d) && d.execScript && d.execScript("var " + c[0]);
		for(var f; c.length && (f = c.shift());) !c.length && b !== g ? d[f] = b : d = d[f] ? d[f] : d[f] = {}
	};
	var u = Array.prototype,
		v = u.indexOf ? function(a, b, c) {
			return u.indexOf.call(a, b, c)
		} : function(a, b, c) {
			c = c == k ? 0 : 0 > c ? Math.max(0, a.length + c) : c;
			if(q(a)) return !q(b) || 1 != b.length ? -1 : a.indexOf(b, c);
			for(; c < a.length; c++)
				if(c in a && a[c] === b) return c;
			return -1
		},
		aa = u.forEach ? function(a, b, c) {
			u.forEach.call(a, b, c)
		} : function(a, b, c) {
			for(var d = a.length, f = q(a) ? a.split("") : a, e = 0; e < d; e++) e in f && b.call(c, f[e], e, a)
		},
		ba = u.filter ? function(a, b, c) {
			return u.filter.call(a, b, c)
		} : function(a, b, c) {
			for(var d = a.length, f = [], e = 0, w = q(a) ? a.split("") :
					a, x = 0; x < d; x++)
				if(x in w) {
					var G = w[x];
					b.call(c, G, x, a) && (f[e++] = G)
				}
			return f
		};

	function ca(a, b, c) {
		return 2 >= arguments.length ? u.slice.call(a, b) : u.slice.call(a, b, c)
	};

	function da(a) {
		a = a.className;
		return q(a) && a.match(/\S+/g) || []
	}

	function ea(a, b) {
		for(var c = da(a), d = ca(arguments, 1), f = c, e = 0; e < d.length; e++) 0 <= v(f, d[e]) || f.push(d[e]);
		a.className = c.join(" ")
	}

	function fa(a, b) {
		var c = da(a),
			d = ca(arguments, 1),
			c = ga(c, d);
		a.className = c.join(" ")
	}

	function ga(a, b) {
		return ba(a, function(a) {
			return !(0 <= v(b, a))
		})
	}

	function y(a, b, c) {
		c ? ea(a, b) : fa(a, b)
	};

	function ha(a) {
		var b = z,
			c;
		for(c in b)
			if(a.call(g, b[c], c, b)) return c
	};
	var A, B, C, D;

	function ia() {
		return n.navigator ? n.navigator.userAgent : k
	}
	D = C = B = A = l;
	var E;
	if(E = ia()) {
		var ja = n.navigator;
		A = 0 == E.indexOf("Opera");
		B = !A && -1 != E.indexOf("MSIE");
		C = !A && -1 != E.indexOf("WebKit");
		D = !A && !C && "Gecko" == ja.product
	}
	var F = B,
		H = D,
		ka = C;

	function la() {
		var a = n.document;
		return a ? a.documentMode : g
	}
	var I;
	a: {
		var J = "",
			K;
		if(A && n.opera) var L = n.opera.version,
			J = "function" == typeof L ? L() : L;
		else if(H ? K = /rv\:([^\);]+)(\)|;)/ : F ? K = /MSIE\s+([^\);]+)(\)|;)/ : ka && (K = /WebKit\/(\S+)/), K) var ma = K.exec(ia()),
			J = ma ? ma[1] : "";
		if(F) {
			var na = la();
			if(na > parseFloat(J)) {
				I = String(na);
				break a
			}
		}
		I = J
	}
	var oa = I,
		pa = {};

	function qa(a) {
		if(!pa[a]) {
			for(var b = 0, c = String(oa).replace(/^[\s\xa0]+|[\s\xa0]+$/g, "").split("."), d = String(a).replace(/^[\s\xa0]+|[\s\xa0]+$/g, "").split("."), f = Math.max(c.length, d.length), e = 0; 0 == b && e < f; e++) {
				var w = c[e] || "",
					x = d[e] || "",
					G = RegExp("(\\d*)(\\D*)", "g"),
					za = RegExp("(\\d*)(\\D*)", "g");
				do {
					var s = G.exec(w) || ["", "", ""],
						t = za.exec(x) || ["", "", ""];
					if(0 == s[0].length && 0 == t[0].length) break;
					b = ((0 == s[1].length ? 0 : parseInt(s[1], 10)) < (0 == t[1].length ? 0 : parseInt(t[1], 10)) ? -1 : (0 == s[1].length ? 0 : parseInt(s[1],
						10)) > (0 == t[1].length ? 0 : parseInt(t[1], 10)) ? 1 : 0) || ((0 == s[2].length) < (0 == t[2].length) ? -1 : (0 == s[2].length) > (0 == t[2].length) ? 1 : 0) || (s[2] < t[2] ? -1 : s[2] > t[2] ? 1 : 0)
				} while (0 == b)
			}
			pa[a] = 0 <= b
		}
	}
	var ra = n.document,
		sa = !ra || !F ? g : la() || ("CSS1Compat" == ra.compatMode ? parseInt(oa, 10) : 5);
	if(H || F) {
		var M;
		if(M = F) M = F && 9 <= sa;
		M || H && qa("1.9.1")
	}
	F && qa("9");

	function ta(a) {
		var b = document;
		return b.querySelectorAll && b.querySelector ? b.querySelectorAll("." + a) : b.getElementsByClassName ? b.getElementsByClassName(a) : ua(a)
	}

	function ua(a) {
		var b, c, d, f;
		b = document;
		if(b.querySelectorAll && b.querySelector && a) return b.querySelectorAll("" + (a ? "." + a : ""));
		if(a && b.getElementsByClassName) {
			var e = b.getElementsByClassName(a);
			return e
		}
		e = b.getElementsByTagName("*");
		if(a) {
			f = {};
			for(c = d = 0; b = e[c]; c++) {
				var w = b.className;
				"function" == typeof w.split && 0 <= v(w.split(/\s+/), a) && (f[d++] = b)
			}
			f.length = d;
			return f
		}
		return e
	}

	function va(a, b) {
		for(var c = 0; a;) {
			if(b(a)) return a;
			a = a.parentNode;
			c++
		}
		return k
	};
	var wa = {};

	function xa() {
		return wa["panel-id"] || (wa["panel-id"] = "panel-id".replace(/\-([a-z])/g, function(a, b) {
			return b.toUpperCase()
		}))
	};
	var N = p("yt.dom.getNextId_");
	if(!N) {
		N = function() {
			return ++ya
		};
		r("yt.dom.getNextId_", N);
		var ya = 0
	};

	function O(a) {
		if(a = a || window.event) {
			for(var b in a) b in Aa || (this[b] = a[b]);
			this.scale = a.scale;
			this.rotation = a.rotation;
			if((b = a.target || a.srcElement) && 3 == b.nodeType) b = b.parentNode;
			this.target = b;
			if(b = a.relatedTarget) try {
				b = b.nodeName && b
			} catch(c) {
				b = k
			};
//			} else "mouseover" == this.type ? b = a.fromElement : "mouseout" == this.type && (b = a.toElement);
			this.relatedTarget = b;
			this.clientX = a.clientX != g ? a.clientX : a.pageX;
			this.clientY = a.clientY != g ? a.clientY : a.pageY;
			if(document.body && document.documentElement) {
				b = document.body.scrollLeft +
					document.documentElement.scrollLeft;
				var d = document.body.scrollTop + document.documentElement.scrollTop;
				this.pageX = a.pageX != g ? a.pageX : a.clientX + b;
				this.pageY = a.pageY != g ? a.pageY : a.clientY + d
			}
			this.keyCode = a.keyCode ? a.keyCode : a.which;
			this.charCode = a.charCode || ("keypress" == this.type ? this.keyCode : 0);
			this.altKey = a.altKey;
			this.ctrlKey = a.ctrlKey;
			this.shiftKey = a.shiftKey;
			"MozMousePixelScroll" == this.type ? (this.wheelDeltaX = a.axis == a.HORIZONTAL_AXIS ? a.detail : 0, this.wheelDeltaY = a.axis == a.HORIZONTAL_AXIS ? 0 : a.detail) :
				window.opera ? (this.wheelDeltaX = 0, this.wheelDeltaY = a.detail) : 0 == a.wheelDelta % 120 ? "WebkitTransform" in document.documentElement.style ? window.a && 0 == navigator.platform.indexOf("Mac") ? (this.wheelDeltaX = a.wheelDeltaX / -30, this.wheelDeltaY = a.wheelDeltaY / -30) : (this.wheelDeltaX = a.wheelDeltaX / -1.2, this.wheelDeltaY = a.wheelDeltaY / -1.2) : (this.wheelDeltaX = 0, this.wheelDeltaY = a.wheelDelta / -1.6) : (this.wheelDeltaX = a.wheelDeltaX / -3, this.wheelDeltaY = a.wheelDeltaY / -3)
		}
	}
	m = O.prototype;
	m.type = "";
	m.target = k;
	m.relatedTarget = k;
	m.currentTarget = k;
	m.keyCode = 0;
	m.charCode = 0;
	m.altKey = l;
	m.ctrlKey = l;
	m.shiftKey = l;
	m.clientX = 0;
	m.clientY = 0;
	m.pageX = 0;
	m.pageY = 0;
	m.wheelDeltaX = 0;
	m.wheelDeltaY = 0;
	m.rotation = 0;
	m.scale = 1;
	m.touches = k;
	var Aa = {
		stopPropagation: 1,
		preventMouseEvent: 1,
		preventManipulation: 1,
		preventDefault: 1,
		layerX: 1,
		layerY: 1,
		scale: 1,
		rotation: 1
	};
	var z = p("yt.events.listeners_") || {};
	r("yt.events.listeners_", z);
	var Ba = p("yt.events.counter_") || {
		count: 0
	};
	r("yt.events.counter_", Ba);

	function Ca(a, b, c) {
		return ha(function(d) {
			return d[0] == a && d[1] == b && d[2] == c && d[4] == l
		})
	}

	function P(a, b, c) {
		if(a && (a.addEventListener || a.attachEvent)) {
			var d = Ca(a, b, c);
			if(!d) {
				var d = ++Ba.count + "",
					f = !(!("mouseenter" == b || "mouseleave" == b) || !a.addEventListener || "onmouseenter" in document),
					e;
				e = f ? function(d) {
					d = new O(d);
					if(!va(d.relatedTarget, function(b) {
							return b == a
						})) return d.currentTarget = a, d.type = b, c.call(a, d)
				} : function(b) {
					b = new O(b);
					b.currentTarget = a;
					return c.call(a, b)
				};
				z[d] = [a, b, c, e, l];
				a.addEventListener ? "mouseenter" == b && f ? a.addEventListener("mouseover", e, l) : "mouseleave" == b && f ? a.addEventListener("mouseout",
					e, l) : "mousewheel" == b && "MozBoxSizing" in document.documentElement.style ? a.addEventListener("MozMousePixelScroll", e, l) : a.addEventListener(b, e, l) : a.attachEvent("on" + b, e)

			}
		}
	};
	r("yt.config_", window.yt && window.yt.config_ || {});
	r("yt.tokens_", window.yt && window.yt.tokens_ || {});
	r("yt.globals_", window.yt && window.yt.globals_ || {});
	r("yt.msgs_", window.yt && window.yt.msgs_ || {});
	var Da = window.yt && window.yt.timeouts_ || [];
	r("yt.timeouts_", Da);
	var Ea = window.yt && window.yt.intervals_ || [];
	r("yt.intervals_", Ea);

	function Fa(a) {
		a = window.setTimeout(a, 0);
		Da.push(a)
	};
	var Q, R, S, T, U = 0,
		V = {
			pageX: 0,
			pageY: 0
		},
		W = l;

	function X() {
		if(!T) {
			var a = window.setInterval(Ga, 1E4);
			Ea.push(a);
			T = a
		}
	}

	function Ha() {
		T && window.clearInterval(T);
		T = k
	}

	function Ga(a) {
		Y(U, l);
		if(a && a.currentTarget) {
			if(a = a.currentTarget, a = a.dataset ? a.dataset[xa()] : a.getAttribute("data-panel-id")) {
				var b = Number(a);
				U = 0 == b && /^[\s\xa0]*$/.test(a) ? NaN : b
			}
		} else U++, R.length == U && (U = 0);
		Y(U, h)
	}

	function Y(a, b) {
		R[a].className = "ban-vaptcha-cont";
		y(R[a], "ban-vaptcha-cont-current", b);
		y(S[a], "ban-vaptcha-nav-item-current", b)
	}

	function Z(a) {
		var b = l;
		if("touchstart" == a.type) V.pageX = a.touches[0].pageX, V.pageY = a.touches[0].pageY, W = h, Ha();
		else if("touchmove" == a.type && W) {
			if(1 < a.touches.length || a.scale && 1 !== a.scale) return;
			var c = V.pageX - a.touches[0].pageX;
			30 < c ? ($(l, h), U++, U == R.length && (U = 0), $(h, l), b = h) : -30 > c && ($(l, l), U--, 0 > U && (U = R.length - 1), $(h, h), b = h)
		}
		if(b || "touchend" == a.type) W = l;
		"touchend" == a.type && (V.pageX = 0, V.pageY = 0, X())
	}

	function $(a, b) {
		var c = U;
		y(R[c], "ban-vaptcha-cont-animate", !a);
		y(R[c], "ban-vaptcha-cont-coming-left", a && b);
		y(R[c], "ban-vaptcha-cont-coming-right", a && !b);
		y(R[c], "ban-vaptcha-cont-leaving-left", !a && b);
		y(R[c], "ban-vaptcha-cont-leaving-right", !a && !b);
		a ? Fa(function() {
			y(R[c], "ban-vaptcha-cont-animate", a);
			y(R[c], "ban-vaptcha-cont-current", a)
		}) : y(R[c], "ban-vaptcha-cont-current", a);
		y(S[c], "ban-vaptcha-nav-item-current", a)
	};
	Q = q("ban-vaptcha") ? document.getElementById("ban-vaptcha") : "ban-vaptcha";
//	P(Q, "mouseenter", Ha);
//	P(Q, "mouseleave", X);
	P(Q, "touchstart", Z);
	P(Q, "touchmove", Z);
	P(Q, "touchend", Z);
	R = ta("ban-vaptcha-cont");
	S = ta("ban-vaptcha-nav-item");
	aa(S, function(a) {
		P(a, "click", Ga)
	});
	Y(0, h);
	X();
})();