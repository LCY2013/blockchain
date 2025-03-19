var $ = require('jquery');
var WebsockFeed = function(url,coin,socket,scale){
    this._datafeedURL = url;
    this.coin = coin;
    this.socket = socket;
    this.lastBar = null;
    this.currentBar = null;
    this.subscribe = true;
    this.scale = scale;
};

WebsockFeed.prototype.onReady=function(callback){
    var config = {};
    config.exchanges = [];
    config.supported_resolutions = ["1","5","15","30","60","240","1D","1W","1M"];
    config.supports_group_request = false;
    config.supports_marks = false;
    config.supports_search = false;
    config.supports_time = true;
    config.supports_timescale_marks = false;

    $("#"+window.tvWidget.id).contents().on("click",".date-range-list>a",function(){
      if (window.tvWidget) {
        if ($(this).html() == "分时") {
          $(this).parent().addClass("real-op").removeClass("common-op");
          window.tvWidget.chart().setChartType(3);
        }else {
          $(this).parent().addClass("common-op").removeClass("real-op");
          window.tvWidget.chart().setChartType(1);
        }
      }
    });

    setTimeout(function() {
        callback(config);
    }, 0);
};

WebsockFeed.prototype.subscribeBars = function(symbolInfo, resolution, onRealtimeCallback, listenerGUID, onResetCacheNeededCallback) {
    var that = this;
    this.socket.on('/topic/swap/trade/'+symbolInfo.name, function(msg) {
        var resp = JSON.parse(msg);
        if(that.lastBar != null && resp.length > 0){
            var price = resp[resp.length -1].price;
            that.lastBar.close = price;
            if(price > that.lastBar.high){
                that.lastBar.high = price;
            }
            if(price < that.lastBar.low){
                that.lastBar.low = price;
            }
            onRealtimeCallback(that.lastBar);
        }
    });
    // function parseTime(time) {
    //     if (time) {
    //       var date = new Date(time)
    //       var year = date.getFullYear()
    //       /* 在日期格式中，月份是从0开始的，因此要加0
    //        * 使用三元表达式在小于10的前面加0，以达到格式统一  如 09:11:05
    //        * */
    //       var month = date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1
    //       var day = date.getDate() < 10 ? '0' + date.getDate() : date.getDate()
    //       var hours = date.getHours() < 10 ? '0' + date.getHours() : date.getHours()
    //       var minutes = date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()
    //       var seconds = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds()
    //       // 拼接
    //       return year + '-' + month + '-' + day + ' ' + hours + ':' + minutes + ':' + seconds
    //     } else {
    //       return ''
    //     }
    //   }
    this.socket.on('/topic/swap/kline/'+symbolInfo.name, function(msg) {
        if (resolution != "1") return;
        if (that.currentBar != null) onRealtimeCallback(that.currentBar);
        var resp = JSON.parse(msg);
        let bar = {time:resp.time*1000,open:resp.openPrice,high:resp.highestPrice,low:resp.lowestPrice,close:resp.closePrice,volume:resp.volume};
        // console.log(symbolInfo.name, "kline", bar, parseTime(that.currentBar.time), parseTime(bar.time));
        if (that.currentBar.time<bar.time) {
            that.lastBar = bar;
            that.currentBar = bar;
            onRealtimeCallback(bar);
        }
    });
};

WebsockFeed.prototype.unsubscribeBars = function(subscriberUID){
    this.subscribe = false;
}

WebsockFeed.prototype.resolveSymbol = function(symbolName, onSymbolResolvedCallback, onResolveErrorCallback){
    // var data = {"name":this.coin.symbol,"exchange-traded":"","exchange-listed":"","minmov":1,"minmov2":0,"pointvalue":1,"has_intraday":true,"has_no_volume":false,"description":this.coin.symbol,"type":"bitcoin","session":"24x7","supported_resolutions":["1","5","15","30","60","240","1D","1W","1M"],"pricescale":500,"ticker":"","timezone":"Asia/Shanghai"};
  // var data = {"name":this.coin.symbol,"exchange-traded":"","exchange-listed":"","minmov":1,"volumescale":10000,"has_daily":true,"has_weekly_and_monthly":true,"has_intraday":true,"description":this.coin.symbol,"type":"bitcoin","session":"24x7","supported_resolutions":["1","5","15","30","60","240","1D","1W","1M"],"pricescale":100,"ticker":"","timezone":"Asia/Shanghai"};
  var data = {"name":this.coin.symbol,"exchange-traded":"","exchange-listed":"","minmov":1,"volumescale":10000,"has_daily":true,"has_weekly_and_monthly":true,"has_intraday":true,"description":this.coin.symbol,"type":"bitcoin","session":"24x7","supported_resolutions":["1","5","15","30","60","1D","1W","1M"],"pricescale":Math.pow(10,this.scale || 2),"ticker":"","timezone":"Asia/Shanghai"};
    setTimeout(function() {
        onSymbolResolvedCallback(data);
    }, 0);
};

WebsockFeed.prototype._send = function(url, params) {
    var request = url;
    if (params) {
        for (var i = 0; i < Object.keys(params).length; ++i) {
            var key = Object.keys(params)[i];
            var value = encodeURIComponent(params[key]);
            request += (i === 0 ? '?' : '&') + key + '=' + value;
        }
    }

    return $.ajax({
        type: 'GET',
        url: request,
        dataType: 'json'
    });
};

WebsockFeed.prototype.getBars = function(symbolInfo, resolution, from, to, onHistoryCallback, onErrorCallback, firstDataRequest){
    var bars = [];
    var that = this;

    this._send(this._datafeedURL+'/history',{
        symbol: symbolInfo.name,
        from: from*1000,
        to: firstDataRequest ? new Date().getTime():to*1000,
        resolution: resolution
    })
    .done(function(response) {
        var data = response.data;
        for(var i = 0;i<data.length;i++){
            var item = data[i];
            bars.push({time:item[0],open:item[1],high:item[2],low:item[3],close:item[4],volume:item[5]})
        }
        that.lastBar = bars.length > 0 ? bars[bars.length-1]:null;
        that.currentBar = that.lastBar;
        var noData = bars.length == 0;
        var retBars = onHistoryCallback(bars,{noData:noData});
    })
    .fail(function(reason) {
        onErrorCallback(reason);
    });
};
WebsockFeed.prototype.periodLengthSeconds = function(resolution, requiredPeriodsCount) {
    var daysCount = 0;
    if (resolution === 'D') {
        daysCount = requiredPeriodsCount;
    } else if (resolution === 'M') {
        daysCount = 31 * requiredPeriodsCount;
    } else if (resolution === 'W') {
        daysCount = 7 * requiredPeriodsCount;
    }
    else if(resolution === 'H'){
        daysCount = requiredPeriodsCount * resolution / 24;
    }
    else {
        daysCount = requiredPeriodsCount * resolution / (24 * 60);
    }

    return daysCount * 24 * 60 * 60;
};

export default {WebsockFeed}
