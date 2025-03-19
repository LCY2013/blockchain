export default {
    common: {
        area: '/uc/support/country'
    },
    uc: { //会员中心接口
        login: '/uc/login',
        register: '/uc/register',
        getRandomly:'/uc/promotion/getRandomly',
        wallet: '/uc/asset/wallet/',
        contractWallet: '/uc/asset/contract-wallet/',
        contractWalletTransfer: '/uc/asset/contract-wallet/transfer',
        captcha: '/uc/start/captcha',
        identification: '/uc/approve/certified/business/status', //商家认证
        apply: '/uc/approve/certified/business/apply', //商家认证申请
        announcement: '/uc/announcement/page', //公告列表
        paydividends: "/uc/bonus/user/page", //持币分红
        mylistrecord: "/uc/mine/detail/", //交易挖矿详情
        activitylist: "/uc/activity/page-query", // 活动列表
        memberactivity: "/uc/activity/getmemberrecords", // 用户参与活动列表
        attendActivity: "/uc/activity/attend", // 参与活动
        buymining: "/uc/mining/buy", // 购买矿机
        mypromotion: "/uc/promotion/mypromotion", // 我的推广佣金
        toppromotion: "/uc/promotion/toprank", // 全球传递爱排行
        getfreecard: "/uc/promotion/promotioncard/getfreecard", // 获取免费推广卡
        exchangecard: "/uc/promotion/promotioncard/exchangecard", // 兑换推广卡
        mycardlist: "/uc/promotion/promotioncard/mycard", // 兑换推广卡
        toppromotionmonth: "/uc/promotion/monthtoprank", // 全球传递爱排行
        toppromotionweek: "/uc/promotion/weektoprank", // 全球传递爱排行
        memberInfo: "/uc/member/my-info", // 获取用户信息
        mypromotionrecord: "/uc/promotion/record", //获取推广邀请人记录
        myInnovationOrderList: "/uc/activity/getmyorders", // 创新实验区参与订单列表
        myInnovationMinings: "/uc/miningorder/my-minings", // 获取我的矿机列表
        myMiningsFetchProfit: "/uc/miningorder/fetch-profit", // 获取我的矿机收益
    },
    market: { //币币交易行情接口
        ws: '/market/market-ws',
        thumb: '/market/symbol-thumb',
        thumbTrend: '/market/symbol-thumb-trend',
        plate: '/market/exchange-plate', //主动查询的盘口信息
        platemini: '/market/exchange-plate-mini', //获取10条数据
        platefull: '/market/exchange-plate-full', //获取所有数据
        trade: '/market/latest-trade', //主动查询的实时成交信息
        symbolInfo: '/market/symbol-info',
        coinInfo: '/market/coin-info',
        indexData: "/market/index_info"
    },
    exchange: { //币币交易委托提交与查询接口
        orderAdd: '/exchange/order/add', //提交订单接口
        current: '/exchange/order/current', //当前委托接口
        history: '/exchange/order/history', //历史委托接口
        detail: '/exchange/order/detail/', //详细订单接口
        favorFind: '/exchange/favor/find', //查询自选
        favorAdd: '/exchange/favor/add', //添加自选
        favorDelete: '/exchange/favor/delete', //删除自选
        orderCancel: '/exchange/order/cancel' //取消委托
    },
    otc: {
        coin: '/otc/coin/all', //查询支持的币种
        advertise: '/otc/advertise/page-by-unit', //获取广告
        createAd: '/uc/ad/create', //创建广告
        adDetail: '/otc/advertise/detail'
    },
    activity: {
        activity: "/activity/page-query"
    },
    crowdfunding:{
        crowdfundingList:"/crowd/funding/enabled",
        donateList:"/crowd/funding/getdetail",
        offlineList:"/crowd/funding/enablewelfare",
        getWelfareDetail:"/crowd/funding/getWelfareDetail",
        medical:"/crowd/funding/addmedicalfunding",
        wish:"/crowd/funding/addfunding",
        publicWelfare:"/crowd/funding/addwelfare",
        donate:"/crowd/funding/addfundingrecord",
        comment:"/crowd/funding/addcommon",
        level:"/crowd/funding/getlevel",
        totalAmount:"/crowd/funding/indexAmount",
        myCrowdfundingYl: "/crowd/funding/getMemberMedicalFunding", //个人中心医疗众筹列表
        myCrowdfundingQt: "/crowd/funding/getMemberFunding", //个人中心其他众筹列表
        myCrowdfundingGy: "/crowd/funding/getMemberWelfare", //个人中心线下公益列表
        myCrowdfundingJk: "/crowd/funding/getFundingRecord", //个人中心我的捐款列表
        getCash:"/crowd/funding/getCash"
    },
    swap: {
        info: '/swap/contract-coin/info/',
        ws: '/swap/swap-ws',
        thumb: '/swap/symbol-thumb',
        plate: '/swap/exchange-plate', //主动查询的盘口信息
        platemini: '/swap/exchange-plate-mini', //获取10条数据
        platefull: '/swap/exchange-plate-full', //获取所有数据
        trade: '/swap/latest-trade', //主动查询的实时成交信息
        symbolInfo: '/swap/symbol-info',
        addOrderEntrust: '/swap/order-entrust/add', //添加委托单
        quickClose: '/swap/order-entrust/quick-close/', //闪电平仓
        current: '/swap/order-entrust/current', //当前委托接口
        history: '/swap/order-entrust/history', //历史委托接口
        position: '/swap/order/position', //当前仓位
        entrustCancel: '/swap/order-entrust/cancel', //取消委托
        leverage: '/swap/contract-leverage'
    }
}
