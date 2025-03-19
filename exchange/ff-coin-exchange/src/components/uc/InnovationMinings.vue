<template>
  <div class="nav-rights">
    <div class="nav-right col-xs-12 col-md-10 padding-right-clear">
      <div class="bill_box rightarea padding-right-clear">
        <section class="trade-group merchant-top">
            <i class="merchant-icon tips"></i>
            <span class="tips-word">{{$t('uc.mining.title')}}</span>
        </section>
        <div class="shaow">
          <div class="money_table mining-list">
            <Row>
                <Col span="24" v-if="miningList.length == 0" style="text-align:center;margin-top: 30px;">
                  {{$t('uc.mining.empty')}}
                </Col>
                <Col :xs="24" :sm="24" :md="8" :lg="8" v-for="item in miningList" :value="item.value" :key="item.value">
                  <Card style="width:94%;position:relative; margin:5px auto;">
                    <div style="width: 100%;min-height: 58px;">
                      <div style="float:left;"><img style="width:50px;height:50px;border-radius:50px;" :src="item.image"></div>
                      <div style="float:left;text-align:left;margin-left: 15px;">
                        <h2 class="shouji">{{item.title}}</h2>
                        <p><span v-if="item.miningStatus == 0" style="font-size:12px;padding: 2px 8px;border-radius:10px;background:#FF0000;">{{$t('uc.mining.status0')}}</span></p>
                        <p><span v-if="item.miningStatus == 1" style="font-size:12px;padding: 2px 8px;border-radius:10px;background:#00b275;">{{$t('uc.mining.status1')}}</span></p>
                        <p><span v-if="item.miningStatus == 2" style="font-size:12px;padding: 2px 8px;border-radius:10px;background:#888;">{{$t('uc.mining.status2')}}</span></p>
                      </div>
                      <Button type="warning" @click="fetchProfit(item)" v-if="item.canFetchProfit > 0" style="float:right;box-shadow: 0px 0px 4px #f29100;">{{$t('uc.mining.pending')}}{{item.canFetchProfit}}</Button>
                    </div>
                    <div style="width:100%;padding: 5px 10px;background: #000;border-radius: 5px;">
                      <table class="config-table">
                        <tr>
                            <td>{{$t('uc.mining.miningUnit')}}</td><td>{{item.miningUnit}}</td>
                            <td>{{$t('uc.mining.miningProfit')}}</td>
                            <td>
                              {{item.totalProfit}}
                            </td>
                        </tr>
                        <tr>
                            <td>{{$t('uc.mining.miningDays')}}</td>
                            <td>{{item.miningTimes}}<span>{{$t('uc.mining.times')}}</span>
                            </td>
                            <td>{{$t('uc.mining.miningedDays')}}</td>
                            <td>{{item.miningedTimes}}<span>{{$t('uc.mining.times')}}</span>
                            </td>
                        </tr>
                        <tr>
                            <td>{{$t('uc.mining.miningDaysProfit')}}</td>
                            <td>{{item.miningDaysprofit}} {{item.miningUnit}}/{{item.miningDays}}
                                <span v-if="item.period==0">{{$t('uc.mining.day')}}</span>
                                <span v-if="item.period==1">{{$t('uc.mining.week')}}</span>
                                <span v-if="item.period==2">{{$t('uc.mining.month')}}</span>
                                <span v-if="item.period==3">{{$t('uc.mining.year')}}</span>
                            </td>
                            <td>{{$t('uc.mining.miningCurrentDaysProfit')}}</td>
                            <td>{{item.currentDaysprofit}} {{item.miningUnit}}/{{item.miningDays}}
                                <span v-if="item.period==0">{{$t('uc.mining.day')}}</span>
                                <span v-if="item.period==1">{{$t('uc.mining.week')}}</span>
                                <span v-if="item.period==2">{{$t('uc.mining.month')}}</span>
                                <span v-if="item.period==3">{{$t('uc.mining.year')}}</span>
                            </td>
                          </tr>
                      </table>
                    </div>
                    <div class="shoujiBtn">
                      <div style="font-size:12px;padding-top: 10px;text-align:left;color: #828ea1;" v-if="item.miningInvite > 0">
                        <Icon type="ios-information-circle" /> {{$t('uc.mining.invitetip1')}}{{item.miningInvite | percentFun}}{{$t('uc.mining.invitetip2')}}<span v-if="item.miningInvitelimit != 0">{{item.miningInvitelimit | percentFun}}%</span>
                        <span v-if="item.miningInvitelimit == 0">-{{$t('uc.mining.invitetip3')}}</span>
                      </div>
                      <div style="font-size:12px;padding-top: 10px;padding-bottom: 10px;text-align:left;color: #828ea1;">
                      <span>{{$t('uc.mining.reminder')}}<router-link to="/crowdfunding" style="color:#f29100;cursor: pointer;">{{$t('uc.mining.btn')}}</router-link></span>
                      </div>
                    </div>
                    <Button long type="warning" :to="'/mining/detail/'+ item.activityId">{{$t('activity.invite2')}}</Button>
                  </Card>
                </Col>
            </Row>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  components: {},
  data() {
    return {
      loginmsg: this.$t("common.logintip"),
      total: 0,
      pageSize: 10,
      loading: true,
      pageNo: 1,
      miningList: []
    };
  },
  methods: {
    fetchProfit(item) {
      this.loading = true;
      item.canFetchProfit = 0;
      let params = {id: item.id};
      this.$http.post(this.host + this.api.uc.myMiningsFetchProfit, params).then(response => {
        var resp = response.body;
        if (resp.code == 0) {
          this.getMyMiningList();
          this.$Message.success(this.$t("uc.mining.fetchdone"))
        } else {
          this.$Message.error(resp.message);
        }
        this.loading = false;
      });
    },
    getMyMiningList() {
      let params = {};
      params.pageNo = this.pageNo;
      params.pageSize = this.pageSize;
      this.$http.post(this.host + this.api.uc.myInnovationMinings, params).then(response => {
        var resp = response.body;
        if (resp.code == 0) {
          this.miningList = resp.data.content;
        } else {
          this.$Message.error(this.loginmsg);
        }
        this.loading = false;
      });
    },
    loadDataPage(data){
      this.pageNo = data;
      this.getMyMiningList();
    }
  },
  created() {
    this.getMyMiningList();
  },
  filters: {
    percentFun: function(value){
      var tem = value * 100;
      return tem.toFixed(0);
    }
  },
  computed: {
  }
};
</script>

<style lang="scss">
.nav-right {
  .rightarea.bill_box {
    .shaow {
      padding: 5px;
    }
    .money_table {
      .search{
        width: 200px;
        margin-bottom: 10px;
      }
      .ivu-table-wrapper {
        .ivu-table-header{
          background: #27313e;
          th{
            color: #fff;
          }
        }
        .ivu-table-body {
          td {
            color: #fff;
            .ivu-table-cell {
              padding: 10px 10px;
              .ivu-btn {
                background: transparent;
                height: 25px;
                padding: 0 0px;
                border-radius: 0;
                span {
                  display: inline-block;
                  line-height: 20px;
                  font-size: 12px;
                  padding: 0 15px;
                  letter-spacing: 1px;
                }
              }
              .ivu-btn.ivu-btn-info {
                border: 1px solid #f0ac19;
                span {
                  color: #f0ac19;
                }
              }
              .ivu-btn.ivu-btn-error {
                border: 1px solid #f15057;
                span {
                  color: #f15057;
                }
              }
              .ivu-btn.ivu-btn-primary {
                border: 1px solid #00b275;
                border: 1px solid #00b275;
                span {
                  color: #00b275;
                }
              }
              .ivu-btn.ivu-btn-default {
                border: 1px solid #2c384f;
                background: #222c3e;
                span {
                  color: #54637a;
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>

<style scoped lang="scss">
.nav-right {
  height: auto;
  overflow: hidden;
  padding: 0 0 0 15px;
  .rightarea.bill_box {
    padding-left: 15px;
    padding-right: 15px;
    width: 100%;
    height: auto;
    overflow: hidden;
  }
}

.demo-spin-icon-load{
  animation: ani-demo-spin 1s linear infinite;
}

.header-btn{
  float:right;padding: 5px 15px;border: 1px solid #f0ac19;color: #f0ac19;
  margin-left: 20px;
  &:hover{
    background: #f0ac19;
    color: #000;
    cursor: pointer;
  }
}
.mining-list .ivu-row .ivu-col .ivu-card{
  background: #2f3e59;
}

.mining-list .ivu-row .ivu-col .ivu-card-bordered{
  border: none!important;
}
.config-table{
  width:100%;
  tr{
    td{
      color: #828ea1;
      font-size: 12px;
      &:nth-child(1){
        text-align:left;
      }
      &:nth-child(2){
        text-align:right;
        padding-right: 10px;
        color: #EEE;
      }
      &:nth-child(3){
        text-align:left;
        padding-left: 10px;
      }
      &:nth-child(4){
        text-align:right;
        color: #EEE;
      }
    }
  }
}

.merchant-top {
    height: 50px;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    padding: 0 15px;
}

.trade-group {
    margin-bottom: 20px;
    font-size: 14px;
}

.merchant-icon {
    display: inline-block;
    margin-left: 4px;
    background-size: 100% 100%;
}

.merchant-top .tips-word {
    -webkit-box-flex: 2;
    -ms-flex-positive: 2;
    flex-grow: 2;
    text-align: left;
}

.merchant-icon.tips {
    width: 4px;
    height: 22px;
    margin-right: 10px;
    background: #f0a70a;
}

.bill_box {
    width: 100%;
    height: auto;
    overflow: hidden;
}
.shoujiBtn{
  width: 100%;
  height: 95px;
}
@media screen and (max-width:1200px) {
.shouji{
  font-size: 15px;
}
.shoujiBtn{
  width: 100%;
  height: auto;
}
}
</style>
