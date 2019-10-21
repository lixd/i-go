(function ($) {
    $.fn.numberRock = function (options) {
        var defaults = {
            speed: 24,
            count: 100
        };
        var opts = $.extend({}, defaults, options);

        var div_by = 100,
            count = opts["count"],
            speed = Math.floor(count / div_by),
            sum = 0,
            $display = this,
            run_count = 1,
            int_speed = opts["speed"];
        var int = setInterval(function () {
            if (run_count <= div_by && speed != 0) {
                $display.text(sum = speed * run_count);
                run_count++;
            } else if (sum < count) {
                $display.text(++sum);
            } else {
                clearInterval(int);
            }
        }, int_speed);
    }

})(jQuery);


function numberRock(options) {
    if ((!options) || (options && (!options['ele'] || !options['count']))) {
        return; //options  && (element || count) not exist
    }
    var defaults = {
        speed: 24,
    };
    var opts = options;
    !options['speed'] && (opts['speed'] = defaults['speed']);
    if (options['count'] == 0) {
        opts['ele'].innerText = 0;
        return;
    }
    var div_by = 100,
        count = opts["count"],
        speed = Math.floor(count / div_by),
        sum = 0,
        $display = opts['ele'], //luo update
        run_count = 1,
        int_speed = opts["speed"];
    var int = setInterval(function () {
        if (run_count <= div_by && speed != 0) {
            // $display.innerText(sum = speed * run_count);
            sum = speed * run_count;
            $display.innerText = sum;
            run_count++;
        } else if (sum < count) {
            // $display.innerText(++sum);
            $display.innerText = ++sum;
        } else {
            clearInterval(int);
        }
    }, int_speed);
}