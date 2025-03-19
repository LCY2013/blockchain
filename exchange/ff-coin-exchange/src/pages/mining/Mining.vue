<template>
  <div class="activity">
    <img class="bannerimg banner-pc" src="../../assets/images/activity-bg_pro.jpg">
    <!-- <img class="bannerimg banner-mobile" src="../../assets/images/activity_mobile.jpg"> -->
    <div class="activity_container">
      <!-- <h1>{{$t("header.labdetail")}}</h1>
      <p style="letter-spacing: 1px;">{{$t("activity.headertitledesc")}}</p> -->
      <div class="main">
          <Tabs :animated="false" style="width:100%;">
              <div class="activity-container">
                <div class="tips-line" v-if="mining.loaded && mining.total == 0"><img src="../../assets/images/emptydata.png"></img></div>
                <div class="tips-line" v-if="!mining.loaded"><Spin size="large"></Spin></div>
                <div class="activity-item" v-for="(item, index) in mining.items" :key="index">
                  <div class="activity-type">{{item.miningName}}</div>
                  <Row>
                      <Col :xs="24" :sm="24" :md="4" :lg="4"><img :src="item.smallImageUrl"></img></Col>
                      <Col :xs="24" :sm="24" :md="20" :lg="20">
                        <div class="title">
                          <span>{{item.title}}</span>
                        </div>
                        <div style="width: 100%;padding-top: 10px;">
                          <Row>
                              <Col :xs="24" :sm="24" :md="10" :lg="10">
                                <div class="progress-text" style="margin-bottom: -3px;">
                                  <p class="gray">{{$t('uc.mining.miningUnit')}} <span style="color:red;">{{item.miningUnit}}</span></p>
                                  <p class="gray">{{$t('activity.limittimes')}}
                                    <span style="color:red;" v-if="item.limitTimes > 0">{{item.limitTimes}}</span>
                                    <span style="color:red;" v-else>{{$t('activity.unlimited')}}</span>
                                  </p>
                                </div>
                                <div class="progress-text">
                                  <span>{{item.unit}}</span>
                                </div>
                              </Col>
                              <Col :xs="24" :sm="24" :md="14" :lg="14">
                                <p class="progress-time">{{$t('uc.mining.miningDays')}}{{item.miningTimes}}{{$t('uc.mining.times')}}</p>
                                <p class="progress-time">
                                  {{$t('uc.mining.miningProfit')}} {{item.miningDaysprofit}}{{item.miningUnit}} / {{item.miningDays}}
                                    <span v-if="item.miningPeriod == 0">{{$t('uc.mining.day')}}</span>
                                    <span v-if="item.miningPeriod == 1">{{$t('uc.mining.week')}}</span>
                                    <span v-if="item.miningPeriod == 2">{{$t('uc.mining.month')}}</span>
                                    <span v-if="item.miningPeriod == 3">{{$t('uc.mining.year')}}</span>
                                </p>
                              </Col>
                          </Row>
                          <Row class="bottom-panel">
                              <Col span="24">
                                <div class="bottom">
                                    <p style="position:relative;top:15px;">{{$t('uc.mining.price')}} <span style="color:red;">{{item.price}}</span>{{item.acceptUnit}}</p>
                                    <Button style="position:relative;top:5px;" type="warning" :to="'/mining/detail/'+ item.id" target="_blank">{{$t('activity.viewdetail')}}</Button>
                                </div>
                                <div class="bottom-mobile">
                                    <p>{{$t('uc.mining.price')}} <span style="color:red;">{{item.price}}</span>{{item.acceptUnit}}</p>
                                    <Button long type="warning" :to="'/mining/detail/'+ item.id">{{$t('activity.viewdetail')}}</Button>
                                </div>
                              </Col>
                          </Row>
                        </div>
                      </Col>
                  </Row>
                </div>
              </div>
              <div class="page" v-if="mining.total > 0">
                <Page :total="mining.total" :pageSize="mining.pageSize" :current="mining.pageNo" @on-change="pageChange"></Page>
              </div>
          </Tabs>
      </div>
    </div>
    <div class="app_bottom">
      <div class="left_logo">
        <img style="float:left;" src="../../assets/images/applogo.png"></img>
        <div style="float:left;height: 40px;line-height:40px;color:#000;">{{$t("cms.downloadslogan")}}</div>
      </div>
      <div class="right_btn_wrap"><router-link to="/app" class="right_btn">{{$t("cms.download")}}</router-link></div>
    </div>
  </div>
</template>

<script>
var moment = require("moment");
export default {

  data() {
    return {
      mining:{
          loaded: false,
          pageSize: 10,
          pageNo: 1,
          total: 0,
          items: []
      }
    };
  },
  created: function() {
    this.init();
    this.getData();
  },
  filters:{
    dateFormat: function(tick) {
      return moment(tick).format("YYYY-MM-DD HH:mm:ss");
    },
    fixedScale: function(value, scale) {
      return value.toFixed(scale);
    }
  },
  computed: {
    lang() {
      return this.$store.state.lang;
    },
    langPram(){
      if(this.$store.state.lang == "简体中文"){
        return "CN";
      }
      if(this.$store.state.lang == "English"){
        return "EN";
      }
      return "CN";
    }
  },
  methods: {
    init() {
      this.$store.commit("navigate", "nav-mining");
    },
    detail(aId){
      this.$router.push("/mining/detail/" + aId);
    },
    pageChange(page){
        this.mining.pageNo = page;
        this.getData();
    },
    getData() {
      let param = {};
      param["pageNo"] = this.mining.pageNo;
      param["pageSize"] = 10;
      this.$http.post(this.host + "/uc/mining/page-query", param).then(res => {
        if (res.status == 200 && res.body.code == 0) {
          let aList = res.body.data.content;
            this.mining.loaded = true;
            this.mining.items = aList;
            this.mining.total = res.body.data.totalElements;
        } else {
          this.$Message.error(res.body.message);
        }
      });
    }
  }
};
</script>

<style scoped>
    .activity .ivu-tabs-bar{
        border-bottom: 1px solid #dcdee2;
    }
    .activity .ivu-tabs-nav .ivu-tabs-tab:hover{
        color: #f0a70a;
    }
    .activity .ivu-tabs-nav .ivu-tabs-tab:hover, .activity .ivu-tabs-nav .ivu-tabs-tab-active{
        color: #f0a70a;
    }
    .activity .ivu-tabs-ink-bar{
        background-color: #f0a70a;
    }
    .app_bottom{
      display: none;
      z-index: 999;
      position: fixed;
      bottom: 0;
      left: 0;
      width: 100%;
      height: 40px;
      background: rgba(242,246,250,1) !important;
    }
    .app_bottom .left_logo img{
      height: 30px;margin-top:5px;border-radius: 5px;margin-left: 5px;margin-right: 5px;
    }
    .app_bottom .right_btn_wrap{
      float: right;height: 40px;line-height: 40px;margin-right: 5px;
    }
    .app_bottom .right_btn{
      color: #FFF;
      border-radius: 3px;
      padding: 0px 10px;
      line-height: 26px;
      height: 26px;
      display: block;
      margin-top: 7px;
      background: linear-gradient(200deg, #ff9900, #ffbe2b, #ffa500);
    }
    @media screen and (max-width:768px){
      .activity{
          padding-top: 45px!important;
      }
      .activity_container {
        padding: 0 2%!important;
      }
      .activity .main {
        margin-top: 20px !important;
      }
      .progress-time{
        text-align: left!important;
        padding-right: 0px!important;
        margin-top: 10px!important;
      }
      .activity-item .title{
        text-align: left!important;
      }
      .activity-item .title div{
        display: none!important;
      }
      .banner-mobile{
        display: block!important;
      }
      .activity_container > h1{
        font-size: 20px!important;
        margin-top: -170px!important;
      }
      .activity-item img {
          width: 100px!important;
          height: 100px!important;
      }
      .activity-item .title span{
        font-size:18px!important;
      }
      .app_bottom{
        display: block!important;
      }
      .bottom-panel .bottom{
        display: none!important;
      }
      .bottom-panel .bottom-mobile{
        display: block!important;
      }
      .bottom-panel .bottom-mobile p{
        margin: 10px 0 10px 0;
      }
      .bottom-panel .bottom-mobile p span{
          font-size: 12px;
          color: #a7a7a7;
          margin-top:15px;
      }
    }
    .banner-pc{
      display: block;
    }
    .banner-mobile{
      display: none;
    }
</style>
<style lang="scss" scoped>
  .activity {
    background: rgba(242,246,250,1) !important;
    height: 100%;
    background-size: cover;
    position: relative;
    overflow: hidden;
    padding-bottom: 50px;
    padding-top: 60px;
    color: #fff;
  }
  .activity .bannerimg {
    width: 100%;
  }
  .activity_container {
    padding: 0 12%;
    text-align: center;
    height: 100%;
    min-height: 600px;
    > h1 {
      margin-top: -200px;
      font-size: 32px;
      line-height: 1;
      padding: 50px 0 20px 0;
      letter-spacing: 3px;
    }
  }
  .activity .main {
    margin-top: 20px;
    display: flex;
    justify-content: space-between;
    flex-direction: row;
    flex-wrap: wrap;
  }
  .activity-container{
    min-height: 200px;
  }
  .tips-line{
    height: 100px;
    line-height: 100px;
    text-align: center;
    position: relative;
    display: inline-block;
    img{
      width: 100px;
      margin-top: 50px;
    }
    .ivu-spin{
      margin-top: 50px;
    }
  }
  .activity-item{
    position:relative;
    overflow: hidden;
    padding: 15px 20px;
    margin-top: 25px;
    margin-bottom: 25px;
    width: 98%;
    margin-left: 1%;
    min-height: 20px;
    background:#FFF;
    border-radius: 5px;
    transition: transform 0.2s ease;
    &:hover{
      box-shadow: 0 0 25px #DDD;
      transform: scale(1.01,1.01);
    }
    .activity-type{
      position: absolute;
      width: 140px;
      height: 25px;
      line-height: 25px;
      background-color: #000;
      color: white;
      text-align: center;
      transform: rotate(-40deg);
      top: 23px;
      left: -30px;
      z-index: 99;
      box-shadow: 1px 1px 4px #000;
      z-index: 99;
    }
    .title{
      width: 100%;padding-top: 10px;display:flex;flex-direction: row;
      span{
        font-size: 22px;color:rgba(49,54,62,1) !important;
      }
      div{
        height: 30px;
        line-height: 30px;
        padding: 0 15px 0 25px;
        border-top-left-radius: 15px;
        border-bottom-left-radius: 15px;
        border-top-right-radius: 3px;
        border-bottom-right-radius: 3px;
        position: relative;
        &:before{
          content: "●";
          position:absolute;
          top: -1px;
          left: 5px;
        }
      }
      div.step0{
        margin-left: 15px;
        color: #FFF;
        border: 1px solid #FFF;
        background: #5bacff;
        background-image: repeating-linear-gradient(60deg, hsla(0,0%,100%,.1), hsla(0,0%,100%,.1) 10px, transparent 0, transparent 20px);
      }
    }
    img{
      width: 160px;
      height: 160px;
    }
  }
  .progress-text{
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    span{
      font-size: 12px;
    }
  }
  .bottom-panel{
      border-top: 1px solid rgb(237, 237, 237);margin-top: 15px;
      .bottom-mobile{
        display: none;
      }
      .bottom{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        span{
          font-size: 12px;
          color: #a7a7a7;
          margin-top:15px;
        }
        button, a{
          margin-top: 11px;
        }
        a.ivu-btn-primary{
          background:#0095ff;
        }
        a.ivu-btn-primary:hover{
          background: #2ba7ff;
        }
      }
  }
  .progress-time{
    font-size: 13px;
    letter-spacing: 1px;
    text-align: right;
  }
  .right{
    float: right;
  }
  .left{
    float: left;
  }
  .gray{
    color: #000;
  }
</style>
