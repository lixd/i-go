(function($) {
    'use strict';
    // Create the defaults once
    var pluginName = 'svgPie',
        defaults = {
            easing: 'easeOutCubic',
            dimension: 28,
            percentage: 50,
            duration: 2000,
            onStart: function() {},
            onComplete: function() {}
        };
    // The actual plugin constructor
    function Plugin(element, options) {
        this.element = element;
        this.settings = $.extend({}, defaults, options);
        this._defaults = defaults;
        this._name = pluginName;
        this.init();
    }
    // Custom easing function borrowed from jQuery-UI
    $.extend($.easing, {
        easeOutCubic: function(x, t, b, c, d) {
            return c * ((t = t / d - 1) * t * t + 1) + b;
        }
    });
    // Avoid Plugin.prototype conflicts
    $.extend(Plugin.prototype, {
        // Initialization logic
        init: function() {
            $(this.element).css({
                'width': this.settings.dimension + 'px',
                'height': this.settings.dimension + 'px'
            });
            this.createSvg();
            this.animateNumber();
            this.animateStrokeDasharray();
            $(this.element).addClass('vp-factor-rendered');
        },
        // SVG pie markup rendering
        createSvg: function() {
            var half = this.settings.dimension / 2;
            var quarter = this.settings.dimension / 4;
            var area = Math.PI * 2 * quarter;
            var svg = 
                '<svg xmlns:svg="http://www.w3.org/2000/svg"' +
                'xmlns="http://www.w3.org/2000/svg"' +
                '>' +
                '<circle r="' + half +
                '" cx="' + half +
                '" cy="' + half +
                '"/>' +
                '<circle r="' + (quarter + 0.5) + // +0.5 to debug non-webkit based browsers
                '" cx="' + half +
                '" cy="' + half + '"' +
                'style="stroke-width:' + half + 'px;' +
                'stroke-dasharray:' + '0px' + ' ' + area + ';' +
                '"/>' +
                '</svg>' +
                '<div class="vp-factor-percentage"' +
                '></div>';
            $(this.element).prepend(svg);
        },
        // Number animation
        animateNumber: function() {
            var $target = $(this.element).find('.vp-factor-percentage');
            $({
                percentageValue: 0
            }).animate({
                percentageValue: this.settings.percentage
            }, {
                duration: this.settings.duration,
                easing: this.settings.easing,
                start: this.settings.onStart,
                step: function() {
                    $target.text(Math.round(this.percentageValue) + '%');
                },
                complete: this.settings.onComplete
            });
        },
        // Pie animation
        animateStrokeDasharray: function() {
            var debug = this.settings.percentage >= 100 ? 1 : 0; // to debug non webkit browsers
            var area = 2 * Math.PI * ((this.settings.dimension / 4) + 0.4); // +0.4 to debug non webkit browsers
            var strokeEndValue = (this.settings.percentage + debug) * area / 100;
            var $target = $(this.element).find('svg circle:nth-child(2)');
            $({
                strokeValue: 0
            }).animate({
                strokeValue: strokeEndValue
            }, {
                duration: this.settings.duration,
                easing: this.settings.easing,
                step: function() {
                    $target.css('stroke-dasharray', this.strokeValue + 'px' + ' ' + area + 'px');
                }
            });
        }
    });
    // preventing against multiple instantiations
    $.fn[pluginName] = function(options) {
        return this.each(function() {
            if (!$.data(this, 'plugin_' + pluginName)) {
                $.data(this, 'plugin_' + pluginName, new Plugin(this, options));
            }
        });
    };
})(jQuery);
